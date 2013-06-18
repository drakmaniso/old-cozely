// Copyright (c) 2013 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package noise

import (
	"github.com/drakmaniso/glam"
	"github.com/drakmaniso/glam/math"
)

//------------------------------------------------------------------------------

var perlinGradient = [12]glam.Vec3{
	glam.Vec3{+1, +1, 0}, glam.Vec3{-1, +1, 0}, glam.Vec3{+1, -1, 0}, glam.Vec3{-1, -1, 0},
	glam.Vec3{+1, 0, +1}, glam.Vec3{-1, 0, +1}, glam.Vec3{+1, 0, -1}, glam.Vec3{-1, 0, -1},
	glam.Vec3{0, +1, +1}, glam.Vec3{0, -1, +1}, glam.Vec3{0, +1, -1}, glam.Vec3{0, -1, -1},
}

//------------------------------------------------------------------------------

var perlinPermutation = [512]int32{
	78, 24, 51, 138, 76, 67, 238, 181, 103, 5, 32, 79, 52, 43, 174, 58,
	242, 37, 249, 233, 84, 122, 102, 217, 147, 189, 12, 115, 218, 142, 30, 197,
	136, 50, 11, 172, 100, 33, 98, 131, 232, 169, 210, 63, 213, 129, 56, 200,
	173, 19, 44, 61, 47, 178, 46, 204, 80, 66, 69, 25, 183, 119, 48, 170,
	101, 251, 215, 109, 237, 65, 94, 93, 6, 194, 34, 185, 1, 87, 45, 96,
	112, 39, 38, 220, 206, 104, 214, 7, 82, 72, 175, 231, 99, 164, 83, 235,
	114, 10, 125, 86, 139, 35, 105, 57, 130, 160, 123, 180, 126, 74, 54, 184,
	201, 23, 223, 151, 21, 244, 18, 140, 216, 143, 134, 132, 179, 64, 4, 246,
	195, 196, 91, 150, 187, 255, 158, 118, 228, 219, 252, 117, 163, 247, 49, 128,
	120, 166, 148, 0, 22, 182, 111, 222, 209, 15, 60, 239, 53, 207, 17, 89,
	90, 159, 144, 212, 234, 243, 157, 20, 205, 250, 8, 248, 149, 199, 26, 75,
	42, 188, 133, 71, 224, 40, 110, 254, 146, 225, 192, 165, 253, 230, 221, 167,
	9, 127, 168, 14, 97, 153, 198, 171, 152, 116, 176, 113, 154, 162, 29, 108,
	3, 211, 41, 77, 31, 68, 107, 13, 2, 92, 88, 245, 208, 59, 177, 137,
	55, 155, 16, 190, 95, 36, 226, 73, 28, 85, 240, 124, 227, 202, 27, 193,
	156, 229, 62, 161, 203, 141, 186, 106, 81, 70, 135, 191, 145, 236, 241, 121,
	// Repeat the table once, so that we don't need to wrap indexes
	78, 24, 51, 138, 76, 67, 238, 181, 103, 5, 32, 79, 52, 43, 174, 58,
	242, 37, 249, 233, 84, 122, 102, 217, 147, 189, 12, 115, 218, 142, 30, 197,
	136, 50, 11, 172, 100, 33, 98, 131, 232, 169, 210, 63, 213, 129, 56, 200,
	173, 19, 44, 61, 47, 178, 46, 204, 80, 66, 69, 25, 183, 119, 48, 170,
	101, 251, 215, 109, 237, 65, 94, 93, 6, 194, 34, 185, 1, 87, 45, 96,
	112, 39, 38, 220, 206, 104, 214, 7, 82, 72, 175, 231, 99, 164, 83, 235,
	114, 10, 125, 86, 139, 35, 105, 57, 130, 160, 123, 180, 126, 74, 54, 184,
	201, 23, 223, 151, 21, 244, 18, 140, 216, 143, 134, 132, 179, 64, 4, 246,
	195, 196, 91, 150, 187, 255, 158, 118, 228, 219, 252, 117, 163, 247, 49, 128,
	120, 166, 148, 0, 22, 182, 111, 222, 209, 15, 60, 239, 53, 207, 17, 89,
	90, 159, 144, 212, 234, 243, 157, 20, 205, 250, 8, 248, 149, 199, 26, 75,
	42, 188, 133, 71, 224, 40, 110, 254, 146, 225, 192, 165, 253, 230, 221, 167,
	9, 127, 168, 14, 97, 153, 198, 171, 152, 116, 176, 113, 154, 162, 29, 108,
	3, 211, 41, 77, 31, 68, 107, 13, 2, 92, 88, 245, 208, 59, 177, 137,
	55, 155, 16, 190, 95, 36, 226, 73, 28, 85, 240, 124, 227, 202, 27, 193,
	156, 229, 62, 161, 203, 141, 186, 106, 81, 70, 135, 191, 145, 236, 241, 121,
}

//------------------------------------------------------------------------------

func perlinFade(x float32) float32 {
	return x * x * x * (x*(x*6.0-15.0) + 10.0)
}

//------------------------------------------------------------------------------

// Perlin3DAt returns the value of a 3D Perlin noise function at position p.
func Perlin3DAt(p glam.Vec3) float32 {
	// Source: "Simplex Noise Demystified" by Stefan Gustavson
	// http://www.itn.liu.se/~stegu/simplexnoise/simplexnoise.pdf

	// Unit grid cell containing point
	ix := int32(math.Floor(p.X))
	iy := int32(math.Floor(p.Y))
	iz := int32(math.Floor(p.Z))

	// Relative coordinates of point within that cell
	rx := p.X - float32(ix)
	ry := p.Y - float32(iy)
	rz := p.Z - float32(iz)

	// Wrap grid cell coordinates at 255
	ix &= 0xFF
	iy &= 0xFF
	iz &= 0xFF

	// Set of gradient indices
	g000 := perlinPermutation[ix+perlinPermutation[iy+perlinPermutation[iz]]] % 12
	g001 := perlinPermutation[ix+perlinPermutation[iy+perlinPermutation[iz+1]]] % 12
	g010 := perlinPermutation[ix+perlinPermutation[iy+1+perlinPermutation[iz]]] % 12
	g011 := perlinPermutation[ix+perlinPermutation[iy+1+perlinPermutation[iz+1]]] % 12
	g100 := perlinPermutation[ix+1+perlinPermutation[iy+perlinPermutation[iz]]] % 12
	g101 := perlinPermutation[ix+1+perlinPermutation[iy+perlinPermutation[iz+1]]] % 12
	g110 := perlinPermutation[ix+1+perlinPermutation[iy+1+perlinPermutation[iz]]] % 12
	g111 := perlinPermutation[ix+1+perlinPermutation[iy+1+perlinPermutation[iz+1]]] % 12

	// Noise contribution for each corner
	n000 := perlinGradient[g000].Dot(glam.Vec3{rx, ry, rz})
	n100 := perlinGradient[g100].Dot(glam.Vec3{rx - 1, ry, rz})
	n010 := perlinGradient[g010].Dot(glam.Vec3{rx, ry - 1, rz})
	n110 := perlinGradient[g110].Dot(glam.Vec3{rx - 1, ry - 1, rz})
	n001 := perlinGradient[g001].Dot(glam.Vec3{rx, ry, rz - 1})
	n101 := perlinGradient[g101].Dot(glam.Vec3{rx - 1, ry, rz - 1})
	n011 := perlinGradient[g011].Dot(glam.Vec3{rx, ry - 1, rz - 1})
	n111 := perlinGradient[g111].Dot(glam.Vec3{rx - 1, ry - 1, rz - 1})

	// Fade courbe
	u := perlinFade(rx)
	v := perlinFade(ry)
	w := perlinFade(rz)

	// Interpolate along x the contributions from each of the corners
	nx00 := math.Mix(n000, n100, u)
	nx01 := math.Mix(n001, n101, u)
	nx10 := math.Mix(n010, n110, u)
	nx11 := math.Mix(n011, n111, u)

	// Interpolate the four results along y
	nxy0 := math.Mix(nx00, nx10, v)
	nxy1 := math.Mix(nx01, nx11, v)

	// Interpolate the two last results along z
	nxyz := math.Mix(nxy0, nxy1, w)

	//return nxyz
	_ = nxyz
	return nxyz
}

//------------------------------------------------------------------------------
