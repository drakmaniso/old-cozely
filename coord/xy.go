// Copyright (c) 2013-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package coord

import (
	"math"

	"github.com/cozely/cozely/x/math32"
)

////////////////////////////////////////////////////////////////////////////////

// Vector represents any two-dimensional vector.
type Vector interface {
	// Cartesian returns the cartesian coordinates of the vector.
	Cartesian() (x, y float32)
}

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
func (c XY) Cartesian() (x, y float32) {
	return c.X, c.Y
}

// CR return the pixel coordinates of the vector. Note that the sign of Y is
// flipped.
func (c XY) CR() CR {
	return CR{int16(c.X), int16(-c.Y)}
}

// XY64 returns the 64-bit float representation of c.
func (c XY) XY64() XY64 {
	return XY64{float64(c.X), float64(c.Y)}
}

// XYZ returns the coordinates {X, Y, z}.
func (c XY) XYZ(z float32) XYZ {
	return XYZ{c.X, c.Y, z}
}

// Plus returns the sum with another vector.
func (c XY) Plus(o XY) XY {
	return XY{c.X + o.X, c.Y + o.Y}
}

// Pluss returns the component-wise sum with two scalars.
func (c XY) Pluss(x, y float32) XY {
	return XY{c.X + x, c.Y + y}
}

// Minus returns the difference with another vector.
func (c XY) Minus(o XY) XY {
	return XY{c.X - o.X, c.Y - o.Y}
}

// Minuss returns the component-wise difference with two scalars.
func (c XY) Minuss(x, y float32) XY {
	return XY{c.X - x, c.Y - y}
}

// Opposite returns the opposite of the vector.
func (c XY) Opposite() XY {
	return XY{-c.X, -c.Y}
}

// Times returns the product with a scalar.
func (c XY) Times(s float32) XY {
	return XY{c.X * s, c.Y * s}
}

// Timess returns the component-wise product with two scalars.
func (c XY) Timess(x, y float32) XY {
	return XY{c.X * x, c.Y * y}
}

// Timescw returns the component-wise product with another vector.
func (c XY) Timescw(o XY) XY {
	return XY{c.X * o.X, c.Y * o.Y}
}

// Slash returns the division by a scalar (which must be non-zero).
func (c XY) Slash(s float32) XY {
	return XY{c.X / s, c.Y / s}
}

// Slashs returns the component-wise division by two scalars (which must be
// non-zero).
func (c XY) Slashs(x, y float32) XY {
	return XY{c.X / x, c.Y / y}
}

// Slashcw returns the component-wise division by another vector (of which both
// X and Y must be non-zero).
func (c XY) Slashcw(o XY) XY {
	return XY{c.X / o.X, c.Y / o.Y}
}

// Mod returns the remainder (modulus) of the division by a scalar (which must
// be non-zero).
func (c XY) Mod(s float32) XY {
	return XY{math32.Mod(c.X, s), math32.Mod(c.Y, s)}
}

// Mods returns the remainder (modulus) of the division by two scalars (which
// must be non-zero).
func (c XY) Mods(x, y float32) XY {
	return XY{math32.Mod(c.X, x), math32.Mod(c.Y, y)}
}

// Modcw returns the remainders (modulus) of the component-wise division by
// another vector (of which both X and Y must be non-zero).
func (c XY) Modcw(o XY) XY {
	return XY{math32.Mod(c.X, o.X), math32.Mod(c.Y, o.Y)}
}

// Modf returns the integer part and the fractional part of (each component of)
// the vector.
func (c XY) Modf() (intg, frac XY) {
	xintg, xfrac := math32.Modf(c.X)
	yintg, yfrac := math32.Modf(c.Y)
	return XY{xintg, yintg}, XY{xfrac, yfrac}
}

// Dot returns the dot product with another vector.
func (c XY) Dot(o XY) float32 {
	return c.X*o.X + c.Y*o.Y
}

// PerpDot returns the dot product with the perpendicular of v and another
// vector.
func (c XY) PerpDot(o XY) float32 {
	return c.X*o.Y - c.Y*o.X
}

// Length returns the euclidian length of the vector.
func (c XY) Length() float32 {
	// Double conversion is faster than math32.Sqrt because the Go compiler
	// optimizes it.
	return float32(math.Sqrt(float64(c.X*c.X + c.Y*c.Y)))
}

// Length2 returns the square of the euclidian length of the vector.
func (c XY) Length2() float32 {
	return c.X*c.X + c.Y*c.Y
}

// Distance returns the distance with another vector.
func (c XY) Distance(o XY) float32 {
	d := XY{c.X - o.X, c.Y - o.Y}
	// Double conversion is faster than math32.Sqrt because the Go compiler
	// optimizes it.
	return float32(math.Sqrt(float64(d.X*d.X + d.Y*d.Y)))
}

// Distance2 returns the square of the distance with another vector.
func (c XY) Distance2(o XY) float32 {
	d := XY{c.X - o.X, c.Y - o.Y}
	return d.X*d.X + d.Y*d.Y
}

// Normalized return the normalization of the vector (i.e. the vector divided
// by its length, which must be non-zero).
func (c XY) Normalized() XY {
	// Double conversion is faster than math32.Sqrt because the Go compiler
	// optimizes it.
	l := float32(math.Sqrt(float64(c.X*c.X + c.Y*c.Y)))
	return XY{c.X / l, c.Y / l}
}

// IsAlmostEqual returns true if the difference between the two vectors is less
// than the specified ULPs (Unit in the Last Place).
//
// Handle special cases: zero, infinites, denormals.
//
// See also IsNearlyEqual and IsRoughlyEqual.
func (c XY) IsAlmostEqual(o XY, ulps uint32) bool {
	return math32.IsAlmostEqual(c.X, o.X, ulps) &&
		math32.IsAlmostEqual(c.Y, o.Y, ulps)
}

// IsNearlyEqual Returns true if the relative error between the two vectors is
// less than epsilon.
//
// Handles special cases: zero, infinites, denormals.
//
// See also IsAlmostEqual and IsRoughlyEqual.
func (c XY) IsNearlyEqual(o XY, epsilon float32) bool {
	return math32.IsNearlyEqual(c.X, o.X, epsilon) &&
		math32.IsNearlyEqual(c.Y, o.Y, epsilon)
}

// IsRoughlyEqual Returns true if the absolute error between the two vectors is
// less than epsilon.
//
// See also IsNearlyEqual and IsAlmostEqual.
func (c XY) IsRoughlyEqual(o XY, epsilon float32) bool {
	return math32.IsRoughlyEqual(c.X, o.X, epsilon) &&
		math32.IsRoughlyEqual(c.Y, o.Y, epsilon)
}

////////////////////////////////////////////////////////////////////////////////

// DA represents a two dimensional vector, defined by its polar coordinates.
type DA struct {
	D float32 // Distance from origin (i.e. radius)
	A float32 // Angle (counter-clockwise from 3 o'clock)
}

// Cartesian returns the cartesian coordinates of the vector. This implements the
// Vector interface.
func (p DA) Cartesian() (x, y float32) {
	return p.D * math32.Cos(p.A), p.D * math32.Sin(p.A)
}

// XY returns the cartesian representation of the vector.
func (p DA) XY() XY {
	return XY{p.D * math32.Cos(p.A), p.D * math32.Sin(p.A)}
}

////////////////////////////////////////////////////////////////////////////////
