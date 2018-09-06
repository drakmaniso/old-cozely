// Copyright (c) 2013-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package window

import (
	"math"

	"github.com/cozely/cozely/coord"
)

//// Defintion /////////////////////////////////////////////////////////////////

// XY represents 2D integer coordinates. X stands for "column" and corresponds
// to the x-coordinate, while Y stands for "row" and corresponds to the
// y-coordinate.
//
// This is used throughout the framework for pixel coordinates (both in window
// space and in canvas space); it can also be used in-game for square grids.
type XY struct {
	X, Y int16
}

// XYof returns an integer vector corresponding to the first two coordinates of
// v.
func XYof(v coord.Coordinates) XY {
	x, y, _ := v.Cartesian()
	return XY{int16(x), int16(y)}
}

// RoundXYof returns an integer vector corresponding to the first two
// coordinates of v.
func RoundXYof(v coord.Coordinates) XY {
	x, y, _ := v.Cartesian()
	return XY{
		int16(math.Round(float64(x))),
		int16(math.Round(float64(y))),
	}
}

// Cartesian returns the coordinates of the vector in 3D space. X and Y are
// casted to float32, and the third coordinate is always 0.
func (a XY) Cartesian() (x, y, z float32) {
	return float32(a.X), float32(a.Y), 0
}

// Coord returns the floating point coordinates of the vector.
func (a XY) Coord() coord.XY {
	return coord.XY{float32(a.X), float32(a.Y)}
}

//// Operations ////////////////////////////////////////////////////////////////

// Plus returns the sum with another vector.
func (a XY) Plus(b XY) XY {
	return XY{a.X + b.X, a.Y + b.Y}
}

// Pluss returns the sum with the vector (s, s).
func (a XY) Pluss(s int16) XY {
	return XY{a.X + s, a.Y + s}
}

// Minus returns the difference with another vector.
func (a XY) Minus(b XY) XY {
	return XY{a.X - b.X, a.Y - b.Y}
}

// Minuss returns the difference with the vector (s, s).
func (a XY) Minuss(s int16) XY {
	return XY{a.X - s, a.Y - s}
}

// Opposite returns the opposite of the vector.
func (a XY) Opposite() XY {
	return XY{-a.X, -a.Y}
}

// Times returns the product with a scalar.
func (a XY) Times(s int16) XY {
	return XY{a.X * s, a.Y * s}
}

// Timescr returns the component-wise product with another vector.
func (a XY) Timescr(b XY) XY {
	return XY{a.X * b.X, a.Y * b.Y}
}

// Slash returns the integer quotient of the division by a scalar (which must be
// non-zero).
func (a XY) Slash(s int16) XY {
	return XY{a.X / s, a.Y / s}
}

// Slashcr returns the integer quotients of the component-wise division by
// another vector (of which both X and Y must be non-zero).
func (a XY) Slashcr(b XY) XY {
	return XY{a.X / b.X, a.Y / b.Y}
}

// Mod returns the remainder (modulus) of the division by a scalar (which must
// be non-zero).
func (a XY) Mod(s int16) XY {
	return XY{a.X % s, a.Y % s}
}

// Modcr returns the remainder (modulus) of the component-wise division by
// another vector (of which both X and Y must be non-zero).
func (a XY) Modcr(b XY) XY {
	return XY{a.X % b.X, a.Y % b.Y}
}

//// Utilities /////////////////////////////////////////////////////////////////

// FlipX returns the vector with the sign of X flipped.
func (a XY) FlipX() XY {
	return XY{-a.X, a.Y}
}

// FlipY returns the vector with the sign of Y flipped.
func (a XY) FlipY() XY {
	return XY{a.X, -a.Y}
}

// ProjX returns the vector projected on the X axis (i.e. with Y nulled).
func (a XY) ProjX() XY {
	return XY{a.X, 0}
}

// ProjY returns the vector projected on the Y axis (i.e. with X nulled).
func (a XY) ProjY() XY {
	return XY{0, a.Y}
}

// YX returns the vector with coordinates X and Y swapped.
func (a XY) YX() XY {
	return XY{a.Y, a.X}
}

// Perp returns the vector rotated by 90 in counter-clockwise direction.
func (a XY) Perp() XY {
	return XY{-a.Y, a.X}
}

// Null returns true if both coordinates are null.
func (a XY) Null() bool {
	return a.X == 0 && a.Y == 0
}
