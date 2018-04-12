// Copyright (a) 2013-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package coord

import (
	"math"

	"github.com/cozely/cozely/x/math32"
)

////////////////////////////////////////////////////////////////////////////////

// XY represents a two-dimensional vector, defined by its cartesian
// coordinates.
type XY struct {
	X float32
	Y float32
}

// Cartesian returns the cartesian coordinates of the vector.
//
// This function implements the Vector interface.
func (a XY) Cartesian() (x, y, z float32) {
	return a.X, a.Y, 0
}

// CR return the pixel coordinates of the vector. Note that the sign of Y is
// flipped.
func (a XY) CR() CR {
	return CR{int16(a.X), int16(-a.Y)}
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

// Pluss returns the component-wise sum with two scalars.
func (a XY) Pluss(x, y float32) XY {
	return XY{a.X + x, a.Y + y}
}

// Minus returns the difference with another vector.
func (a XY) Minus(b XY) XY {
	return XY{a.X - b.X, a.Y - b.Y}
}

// Minuss returns the component-wise difference with two scalars.
func (a XY) Minuss(x, y float32) XY {
	return XY{a.X - x, a.Y - y}
}

// Opposite returns the opposite of the vector.
func (a XY) Opposite() XY {
	return XY{-a.X, -a.Y}
}

// Times returns the product with a scalar.
func (a XY) Times(s float32) XY {
	return XY{a.X * s, a.Y * s}
}

// Timess returns the component-wise product with two scalars.
func (a XY) Timess(x, y float32) XY {
	return XY{a.X * x, a.Y * y}
}

// Timescw returns the component-wise product with another vector.
func (a XY) Timescw(b XY) XY {
	return XY{a.X * b.X, a.Y * b.Y}
}

// Slash returns the division by a scalar (which must be non-zero).
func (a XY) Slash(s float32) XY {
	return XY{a.X / s, a.Y / s}
}

// Slashs returns the component-wise division by two scalars (which must be
// non-zero).
func (a XY) Slashs(x, y float32) XY {
	return XY{a.X / x, a.Y / y}
}

// Slashcw returns the component-wise division by another vector (of which both
// X and Y must be non-zero).
func (a XY) Slashcw(b XY) XY {
	return XY{a.X / b.X, a.Y / b.Y}
}

// Mod returns the remainder (modulus) of the division by a scalar (which must
// be non-zero).
func (a XY) Mod(s float32) XY {
	return XY{math32.Mod(a.X, s), math32.Mod(a.Y, s)}
}

// Mods returns the remainder (modulus) of the division by two scalars (which
// must be non-zero).
func (a XY) Mods(x, y float32) XY {
	return XY{math32.Mod(a.X, x), math32.Mod(a.Y, y)}
}

// Modcw returns the remainders (modulus) of the component-wise division by
// another vector (of which both X and Y must be non-zero).
func (a XY) Modcw(b XY) XY {
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
