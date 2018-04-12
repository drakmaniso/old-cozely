// Copyright (c) 2013-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package noise

import (
	"github.com/cozely/cozely/coord"
	"github.com/cozely/cozely/x/math32"
)

////////////////////////////////////////////////////////////////////////////////

// Skewing and unskewing factors for 2, 3, and 4 dimensions
const sqrt3 = 1.732050807568877
const sqrt5 = 2.23606797749979
const f2 = 0.5 * (sqrt3 - 1.0)
const g2 = (3.0 - sqrt3) / 6.0
const f3 = 1.0 / 3.0
const g3 = 1.0 / 6.0
const f4 = (sqrt5 - 1.0) / 4.0
const g4 = (5.0 - sqrt5) / 20.0

const simplexNormalization = 99.204334582718712976990005025589

////////////////////////////////////////////////////////////////////////////////

// Simplex2D returns the 2D simplex noise at position p.
func Simplex2D(p coord.XY, grad []coord.XY) float32 {
	// Source: "Simplex Noise Demystified" by Stefan Gustavson
	// http://www.itn.liu.se/~stegu/simplexnoise/simplexnoise.pdf
	// and
	// http://webstaff.itn.liu.se/~stegu/simplexnoise/SimplexNoise.java

	// p = p.Times(0.70710678118654752440084436210485)
	p = p.Times(0.816496580928)

	// Noise contributions from the three corners
	var n0, n1, n2 float32

	// Skew the input space to determine which simplex cell we're in.
	var s = (p.X + p.Y) * f2
	var i = float32(math32.FastFloor(p.X + s))
	var j = float32(math32.FastFloor(p.Y + s))

	// Unskew the cell origin back to (x,y) space.
	var t = (i + j) * g2
	var x0 = i - t
	var y0 = j - t
	var dx0 = p.X - x0 // The x,y distances from the cell origin
	var dy0 = p.Y - y0

	// For the 2D case, the simplex shape is an equilateral triangle.
	// Determine which simplex we are in.
	var i1, j1 int32 // Offsets for second (middle) corner of simplex in (i,j) coords
	if dx0 > dy0 {
		// lower triangle, XY order: (0,0)->(1,0)->(1,1)
		i1 = 1
		j1 = 0
	} else {
		// upper triangle, YX order: (0,0)->(0,1)->(1,1)
		i1 = 0
		j1 = 1
	}

	// A step of (1,0) in (i,j) means a step of (1-c,-c) in (x,y), and
	// a step of (0,1) in (i,j) means a step of (-c,1-c) in (x,y), where
	// c = (3-sqrt(3))/6
	var x1 = dx0 - float32(i1) + g2 // Offsets for middle corner in (x,y) unskewed coords
	var y1 = dy0 - float32(j1) + g2
	var x2 = dx0 - 1.0 + 2.0*g2 // Offsets for last corner in (x,y) unskewed coords
	var y2 = dy0 - 1.0 + 2.0*g2

	// Work out the hashed gradient indices of the three simplex corners
	var ii = int32(i) & 255
	var jj = int32(j) & 255
	var gl = int32(len(grad))
	var gi0 = permutation[ii+permutation[jj]] % gl
	var gi1 = permutation[ii+i1+permutation[jj+j1]] % gl
	var gi2 = permutation[ii+1+permutation[jj+1]] % gl

	// Calculate the contribution from the three corners

	var t0 = 0.5 - dx0*dx0 - dy0*dy0
	if t0 < 0 {
		n0 = 0.0
	} else {
		t0 *= t0
		n0 = t0 * t0 * (grad[gi0].Dot(coord.XY{dx0, dy0}))
	}

	var t1 = 0.5 - x1*x1 - y1*y1
	if t1 < 0 {
		n1 = 0.0
	} else {
		t1 *= t1
		n1 = t1 * t1 * (grad[gi1].Dot(coord.XY{x1, y1}))
	}

	var t2 = 0.5 - x2*x2 - y2*y2
	if t2 < 0 {
		n2 = 0.0
	} else {
		t2 *= t2
		n2 = t2 * t2 * (grad[gi2].Dot(coord.XY{x2, y2}))
	}

	// Add contributions from each corner to get the final noise value.
	// The result is scaled to return values in the interval [-1,1].

	return simplexNormalization * (n0 + n1 + n2)
}

////////////////////////////////////////////////////////////////////////////////

// SimplexAxial returns the 2D simplex noise at position (q, r), expressed
// in axial coordinates.
func SimplexAxial(q, r float32, grad []coord.XY) float32 {
	// Source: "Simplex Noise Demystified" by Stefan Gustavson
	// http://www.itn.liu.se/~stegu/simplexnoise/simplexnoise.pdf
	// and
	// http://webstaff.itn.liu.se/~stegu/simplexnoise/SimplexNoise.java

	// Determine origin of simplex cell
	var q0 = math32.Floor(q + r)
	var r0 = math32.Floor(r)

	// Calculate the cartesian distance to cell origin
	var x = q + r*0.5
	var y = r * 0.5 * sqrt3
	var x0 = q0 - r0*0.5
	var y0 = r0 * 0.5 * sqrt3
	var dx0 = x - x0
	var dy0 = y - y0

	// Determine which simplex (i.e. triangle) we are in
	var q1, r1 int32 // Offsets for the second corner of triangle
	if dx0*sqrt3 > dy0 {
		// lower triangle: (0,0)->(1,0)->(1,1)
		q1 = 1
		r1 = 0
	} else {
		// upper triangle: (0,0)->(0,1)->(1,1)
		q1 = 0
		r1 = 1
	}

	var x1 = dx0 + 0.5 - 1.5*float32(q1)
	var y1 = dy0 - 0.5*sqrt3*float32(r1)
	var x2 = dx0 - 0.5
	var y2 = dy0 - 0.5*sqrt3

	dx0 *= 0.816496580928
	dy0 *= 0.816496580928
	x1 *= 0.816496580928
	y1 *= 0.816496580928
	x2 *= 0.816496580928
	y2 *= 0.816496580928

	// Work out the hashed gradient indices of the three simplex corners
	var qq = int32(q0) & 255
	var rr = int32(r0) & 255
	var gl = int32(len(grad))
	var gi0 = permutation[qq+permutation[rr]] % gl
	var gi1 = permutation[qq+q1+permutation[rr+r1]] % gl
	var gi2 = permutation[qq+1+permutation[rr+1]] % gl

	// Calculate the noise contribution from the three corners
	var n0, n1, n2 float32

	var t0 = 0.5 - dx0*dx0 - dy0*dy0
	if t0 < 0 {
		n0 = 0.0
	} else {
		t0 *= t0
		n0 = t0 * t0 * (grad[gi0].Dot(coord.XY{dx0, dy0}))
	}

	var t1 = 0.5 - x1*x1 - y1*y1
	if t1 < 0 {
		n1 = 0.0
	} else {
		t1 *= t1
		n1 = t1 * t1 * (grad[gi1].Dot(coord.XY{x1, y1}))
	}

	var t2 = 0.5 - x2*x2 - y2*y2
	if t2 < 0 {
		n2 = 0.0
	} else {
		t2 *= t2
		n2 = t2 * t2 * (grad[gi2].Dot(coord.XY{x2, y2}))
	}

	// Add contributions from each corner to get the final noise value.
	// The result is scaled to return values in the interval [-1,1].

	return 99.2043 * (n0 + n1 + n2)
}

////////////////////////////////////////////////////////////////////////////////
