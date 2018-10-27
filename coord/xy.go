// Copyright (a) 2013-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package coord

import (
	"math"

	"github.com/cozely/cozely/x/math32"
)

////////////////////////////////////////////////////////////////////////////////

// XY represents 2D cartesian coordinates.
type XY struct {
	X float32
	Y float32
}

// Cartesian returns the two coordinates; z is always 0.
func (a XY) Cartesian() (x, y, z float32) {
	return a.X, a.Y, 0
}

// XYof converts the first two coordinates of c to XY.
func XYof(c Coordinates) XY {
	x, y, _ := c.Cartesian()
	return XY{x, y}
}

// CR return the integer coordinates of the vector.
func (a XY) CR() CR {
	return CR{int16(a.X), int16(a.Y)}
}

// RoundCR return the rounded interger coordinates of the vector.
func (a XY) RoundCR() CR {
	return CR{
		int16(math.Round(float64(a.X))),
		int16(math.Round(float64(a.Y))),
	}
}

// XY64 returns the 64-bit float representation of a.
func (a XY) XY64() XY64 {
	return XY64{float64(a.X), float64(a.Y)}
}

// XYZ returns the coordinates {X, Y, z}.
func (a XY) XYZ(z float32) XYZ {
	return XYZ{a.X, a.Y, z}
}

// Plus returns the sum with another vector.
func (a XY) Plus(b XY) XY {
	return XY{a.X + b.X, a.Y + b.Y}
}

// Pluss returns the sum with a scalar.
func (a XY) Pluss(s float32) XY {
	return XY{a.X + s, a.Y + s}
}

// Minus returns the difference with another vector.
func (a XY) Minus(b XY) XY {
	return XY{a.X - b.X, a.Y - b.Y}
}

// Minuss returns the difference with a scalar.
func (a XY) Minuss(s float32) XY {
	return XY{a.X - s, a.Y - s}
}

// Opposite returns the opposite of the vector.
func (a XY) Opposite() XY {
	return XY{-a.X, -a.Y}
}

// Timess returns the product with a scalar.
func (a XY) Timess(s float32) XY {
	return XY{a.X * s, a.Y * s}
}

// Times returns the component-wise product with another vector.
func (a XY) Times(b XY) XY {
	return XY{a.X * b.X, a.Y * b.Y}
}

// Slashs returns the division by a scalar (which must be non-zero).
func (a XY) Slashs(s float32) XY {
	return XY{a.X / s, a.Y / s}
}

// Slash returns the component-wise division by another vector (of which both
// X and Y must be non-zero).
func (a XY) Slash(b XY) XY {
	return XY{a.X / b.X, a.Y / b.Y}
}

// Mods returns the remainder (modulus) of the division by a scalar (which must
// be non-zero).
func (a XY) Mods(s float32) XY {
	return XY{math32.Mod(a.X, s), math32.Mod(a.Y, s)}
}

// Mod returns the remainders (modulus) of the component-wise division by
// another vector (of which both X and Y must be non-zero).
func (a XY) Mod(b XY) XY {
	return XY{math32.Mod(a.X, b.X), math32.Mod(a.Y, b.Y)}
}

// Modf returns the integer part and the fractional part of (each component of)
// the vector.
func (a XY) Modf() (intg, frac XY) {
	xintg, xfrac := math32.Modf(a.X)
	yintg, yfrac := math32.Modf(a.Y)
	return XY{xintg, yintg}, XY{xfrac, yfrac}
}

// Dot returns the dot product with another vector.
func (a XY) Dot(b XY) float32 {
	return a.X*b.X + a.Y*b.Y
}

// PerpDot returns the dot product with the perpendicular of v and another
// vector.
func (a XY) PerpDot(b XY) float32 {
	return a.X*b.Y - a.Y*b.X
}

// Length returns the euclidian length of the vector.
func (a XY) Length() float32 {
	// Double conversion is faster than math32.Sqrt because the Go compiler
	// optimizes it.
	return float32(math.Sqrt(float64(a.X*a.X + a.Y*a.Y)))
}

// Length2 returns the square of the euclidian length of the vector.
func (a XY) Length2() float32 {
	return a.X*a.X + a.Y*a.Y
}

// Distance returns the distance with another vector.
func (a XY) Distance(b XY) float32 {
	d := XY{a.X - b.X, a.Y - b.Y}
	// Double conversion is faster than math32.Sqrt because the Go compiler
	// optimizes it.
	return float32(math.Sqrt(float64(d.X*d.X + d.Y*d.Y)))
}

// Distance2 returns the square of the distance with another vector.
func (a XY) Distance2(b XY) float32 {
	d := XY{a.X - b.X, a.Y - b.Y}
	return d.X*d.X + d.Y*d.Y
}

// Normalized return the normalization of the vector (i.e. the vector divided
// by its length, which must be non-zero).
func (a XY) Normalized() XY {
	// Double conversion is faster than math32.Sqrt because the Go compiler
	// optimizes it.
	l := float32(math.Sqrt(float64(a.X*a.X + a.Y*a.Y)))
	return XY{a.X / l, a.Y / l}
}

// IsAlmostEqual returns true if the difference between the two vectors is less
// than the specified ULPs (Unit in the Last Place).
//
// Handle special cases: zero, infinites, denormals.
//
// See also IsNearlyEqual and IsRoughlyEqual.
func (a XY) IsAlmostEqual(b XY, ulps uint32) bool {
	return math32.IsAlmostEqual(a.X, b.X, ulps) &&
		math32.IsAlmostEqual(a.Y, b.Y, ulps)
}

// IsNearlyEqual Returns true if the relative error between the two vectors is
// less than epsilon.
//
// Handles special cases: zero, infinites, denormals.
//
// See also IsAlmostEqual and IsRoughlyEqual.
func (a XY) IsNearlyEqual(b XY, epsilon float32) bool {
	return math32.IsNearlyEqual(a.X, b.X, epsilon) &&
		math32.IsNearlyEqual(a.Y, b.Y, epsilon)
}

// IsRoughlyEqual Returns true if the absolute error between the two vectors is
// less than epsilon.
//
// See also IsNearlyEqual and IsAlmostEqual.
func (a XY) IsRoughlyEqual(b XY, epsilon float32) bool {
	return math32.IsRoughlyEqual(a.X, b.X, epsilon) &&
		math32.IsRoughlyEqual(a.Y, b.Y, epsilon)
}

////////////////////////////////////////////////////////////////////////////////

// FlipX returns the coordinates with the signe of X flipped.
func (a XY) FlipX() XY {
	return XY{-a.X, a.Y}
}

// FlipY returns the coordinates with the signe of Y flipped.
func (a XY) FlipY() XY {
	return XY{a.X, -a.Y}
}

////////////////////////////////////////////////////////////////////////////////
