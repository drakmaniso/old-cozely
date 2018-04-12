// Copyright (c) 2013-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package coord

import (
	"math"
)

////////////////////////////////////////////////////////////////////////////////

// XY64 represents a two-dimensional vector, defined by its cartesian
// coordinates in 64-bit float.
type XY64 struct {
	X float64
	Y float64
}

// XY returns the 32-bit float version of a.
func (a XY64) XY() XY {
	return XY{float32(a.X), float32(a.Y)}
}

// Cartesian returns the cartesian coordinates of the vector.
//
// This function implements the Vector interface.
func (a XY64) Cartesian() (x, y, z float32) {
	return float32(a.X), float32(a.Y), 0
}

// XYZ64 returns the homogenous coordinates of the vector, with Z set to 1.
func (a XY64) XYZ64() XYZ64 {
	return XYZ64{a.X, a.Y, 1.0}
}

// Plus returns the sum with another vector.
func (a XY64) Plus(b XY64) XY64 {
	return XY64{a.X + b.X, a.Y + b.Y}
}

// Pluss returns the component-wise sum with two scalars.
func (a XY64) Pluss(x, y float64) XY64 {
	return XY64{a.X + x, a.Y + y}
}

// Minus returns the difference with another vector.
func (a XY64) Minus(b XY64) XY64 {
	return XY64{a.X - b.X, a.Y - b.Y}
}

// Minuss returns the component-wise difference with two scalars.
func (a XY64) Minuss(x, y float64) XY64 {
	return XY64{a.X - x, a.Y - y}
}

// Opposite returns the opposite of the vector.
func (a XY64) Opposite() XY64 {
	return XY64{-a.X, -a.Y}
}

// Times returns the product with a scalar.
func (a XY64) Times(s float64) XY64 {
	return XY64{a.X * s, a.Y * s}
}

// Timess returns the component-wise product with two scalars.
func (a XY64) Timess(x, y float64) XY64 {
	return XY64{a.X * x, a.Y * y}
}

// Timescw returns the component-wise product with another vector.
func (a XY64) Timescw(b XY64) XY64 {
	return XY64{a.X * b.X, a.Y * b.Y}
}

// Slash returns the division by a scalar (which must be non-zero).
func (a XY64) Slash(s float64) XY64 {
	return XY64{a.X / s, a.Y / s}
}

// Slashs returns the component-wise division by two scalars (which must be
// non-zero).
func (a XY64) Slashs(x, y float64) XY64 {
	return XY64{a.X / x, a.Y / y}
}

// Slashcw returns the component-wise division by another vector (of which both
// X and Y must be non-zero).
func (a XY64) Slashcw(b XY64) XY64 {
	return XY64{a.X / b.X, a.Y / b.Y}
}

// Mod returns the remainder (modulus) of the division by a scalar (which must
// be non-zero).
func (a XY64) Mod(s float64) XY64 {
	return XY64{math.Mod(a.X, s), math.Mod(a.Y, s)}
}

// Mods returns the remainder (modulus) of the component-wise division by two
// scalars (which must be non-zero).
func (a XY64) Mods(x, y float64) XY64 {
	return XY64{math.Mod(a.X, x), math.Mod(a.Y, y)}
}

// Modcw returns the remainders (modulus) of the component-wise division by
// another vector (of which both X and Y must be non-zero).
func (a XY64) Modcw(b XY64) XY64 {
	return XY64{math.Mod(a.X, b.X), math.Mod(a.Y, b.Y)}
}

// Modf returns the integer part and the fractional part of (each component of)
// the vector.
func (a XY64) Modf() (intg, frac XY64) {
	xintg, xfrac := math.Modf(a.X)
	yintg, yfrac := math.Modf(a.Y)
	return XY64{xintg, yintg}, XY64{xfrac, yfrac}
}

// Dot returns the dot product with another vector.
func (a XY64) Dot(b XY64) float64 {
	return a.X*b.X + a.Y*b.Y
}

// PerpDot returns the dot product with the perpendicular of a and another
// vector.
func (a XY64) PerpDot(b XY64) float64 {
	return a.X*b.Y - a.Y*b.X
}

// Length returns the euclidian length of the vector.
func (a XY64) Length() float64 {
	// Double conversion is faster than math.Sqrt because the Go compiler
	// optimizes it.
	return float64(math.Sqrt(float64(a.X*a.X + a.Y*a.Y)))
}

// Length2 returns the square of the euclidian length of the vector.
func (a XY64) Length2() float64 {
	return a.X*a.X + a.Y*a.Y
}

// Distance returns the distance with another vector.
func (a XY64) Distance(b XY64) float64 {
	d := XY64{a.X - b.X, a.Y - b.Y}
	// Double conversion is faster than math.Sqrt because the Go compiler
	// optimizes it.
	return float64(math.Sqrt(float64(d.X*d.X + d.Y*d.Y)))
}

// Distance2 returns the square of the distance with another vector.
func (a XY64) Distance2(b XY64) float64 {
	d := XY64{a.X - b.X, a.Y - b.Y}
	return d.X*d.X + d.Y*d.Y
}

// Normalized return the normalization of the vector (i.e. the vector divided
// by its length, which must be non-zero).
func (a XY64) Normalized() XY64 {
	// Double conversion is faster than math.Sqrt because the Go compiler
	// optimizes it.
	l := float64(math.Sqrt(float64(a.X*a.X + a.Y*a.Y)))
	return XY64{a.X / l, a.Y / l}
}

// IsAlmostEqual returns true if the difference between the two vectors is less
// than the specified ULPs (Unit in the Last Place).
//
// Handle special cases: zero, infinites, denormals.
//
// See also IsNearlyEqual and IsRoughlyEqual.
//TODO:
// func (a Coord64) IsAlmostEqual(b Coord64, ulps uint32) bool {
// 	return math.IsAlmostEqual(a.X, b.X, ulps) &&
// 		math.IsAlmostEqual(a.Y, b.Y, ulps)
// }

// IsNearlyEqual Returns true if the relative error between the two vectors is
// less than epsilon.
//
// Handles special cases: zero, infinites, denormals.
//
// See also IsAlmostEqual and IsRoughlyEqual.
//TODO:
// func (a Coord64) IsNearlyEqual(b Coord64, epsilon float64) bool {
// 	return math.IsNearlyEqual(a.X, b.X, epsilon) &&
// 		math.IsNearlyEqual(a.Y, b.Y, epsilon)
// }

// IsRoughlyEqual Returns true if the absolute error between the two vectors is
// less than epsilon.
//
// See also IsNearlyEqual and IsAlmostEqual.
//TODO:
// func (a Coord64) IsRoughlyEqual(b Coord64, epsilon float64) bool {
// 	return math.IsRoughlyEqual(a.X, b.X, epsilon) &&
// 		math.IsRoughlyEqual(a.Y, b.Y, epsilon)
// }

////////////////////////////////////////////////////////////////////////////////

// XYZ64 represents a two-dimensional vector, defined by its homogeneous
// coordinates.
type XYZ64 struct {
	X float64
	Y float64
	Z float64
}

// Cartesian returns the cartesian coordinates of the vector (i.e. the perspective
// divide of the homogeneous coordinates). Z must be non-zero.
func (a XYZ64) Cartesian() (x, y float32) {
	return float32(a.X / a.Z), float32(a.Y / a.Z)
}

////////////////////////////////////////////////////////////////////////////////

// XYZ64 represents a two-dimensional vector, defined by its homogeneous
// coordinates.
type XYZW64 struct {
	X float64
	Y float64
	Z float64
	W float64
}

// Cartesian returns the cartesian coordinates of the vector (i.e. the perspective
// divide of the homogeneous coordinates).
func (a XYZW64) Cartesian() (x, y, z float32) {
	return float32(a.X / a.W), float32(a.Y / a.W), float32(a.Z / a.W)
}

////////////////////////////////////////////////////////////////////////////////
