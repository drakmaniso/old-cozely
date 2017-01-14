// Copyright (c) 2013-2017 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package plane

//------------------------------------------------------------------------------

import "github.com/drakmaniso/glam/math"

//------------------------------------------------------------------------------

// Identity matrix.
func Identity() Matrix {
	return Matrix{
		{1, 0, 0},
		{0, 1, 0},
		{0, 0, 1},
	}
}

//------------------------------------------------------------------------------

// Transform applies the tranformation matrix to a vector (i.e. returns matrix
// multiplied by column vector).
func (m Matrix) Transform(v Homogen) Homogen {
	return Homogen{
		X: m[0][0]*v.X + m[1][0]*v.Y + m[2][0]*v.Z,
		Y: m[0][1]*v.X + m[1][1]*v.Y + m[2][1]*v.Z,
		Z: m[0][2]*v.X + m[1][2]*v.Y + m[2][2]*v.Z,
	}
}

//------------------------------------------------------------------------------

// Translation by a vector.
func Translation(t Vector) Matrix {
	x, y := t.Cartesian()

	return Matrix{
		{1, 0, 0},
		{0, 1, 0},
		{x, y, 1},
	}
}

//------------------------------------------------------------------------------

// Rotation by an angle.
func Rotation(angle float32) Matrix {
	c := math.Cos(angle)
	s := math.Sin(angle)

	return Matrix{
		{c, -s, 0},
		{s, c, 0},
		{0, 0, 1},
	}
}

//------------------------------------------------------------------------------

// RotationAround a point.
func RotationAround(angle float32, center Vector) Matrix {
	c := math.Cos(angle)
	s := math.Sin(angle)

	x, y := center.Cartesian()

	return Matrix{
		{c, -s, 0},
		{s, c, 0},
		{x - c*x - s*y, y + s*x - c*y, 1},
	}
}

//------------------------------------------------------------------------------

// Scaling and/or mirror along both axis.
func Scaling(s Vector) Matrix {
	x, y := s.Cartesian()

	return Matrix{
		{x, 0, 0},
		{0, y, 0},
		{0, 0, 1},
	}
}

//------------------------------------------------------------------------------

// ScalingAround a point (and/or mirror).
func ScalingAround(s Vector, center Vector) Matrix {
	cx, cy := center.Cartesian()
	sx, sy := s.Cartesian()

	return Matrix{
		{sx, 0, 0},
		{0, sy, 0},
		{cx - cx*sx, cy - cy*sy, 1},
	}
}

//------------------------------------------------------------------------------

// Shearing along both axis.
func Shearing(s Vector) Matrix {
	x, y := s.Cartesian()

	return Matrix{
		{1, y, 0},
		{x, 1, 0},
		{0, 0, 1},
	}
}

//------------------------------------------------------------------------------

// Viewport returns a matrix that scale to an aspect ratio and zoom.
func Viewport(zoom, aspectRatio float32) Matrix {
	height := zoom / 2
	width := height * aspectRatio
	return Scaling(Coord{X: width, Y: height})
}

//------------------------------------------------------------------------------
