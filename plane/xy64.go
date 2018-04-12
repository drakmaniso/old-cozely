// Copyright (c) 2013-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package plane

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

// XY32 returns the 32-bit float version of v.
func (v XY64) XY32() XY {
	return XY{float32(v.X), float32(v.Y)}
}

// Cartesian returns the cartesian coordinates of the vector.
//
// This function implements the Vector interface.
func (v XY64) Cartesian() (x, y float32) {
	return float32(v.X), float32(v.Y)
}

// XYZ64 returns the homogenous coordinates of the vector, with Z set to 1.
func (v XY64) XYZ64() XYZ64 {
	return XYZ64{v.X, v.Y, 1.0}
}

// Plus returns the sum with another vector.
func (v XY64) Plus(o XY64) XY64 {
	return XY64{v.X + o.X, v.Y + o.Y}
}

// Pluss returns the component-wise sum with two scalars.
func (v XY64) Pluss(x, y float64) XY64 {
	return XY64{v.X + x, v.Y + y}
}

// Minus returns the difference with another vector.
func (v XY64) Minus(o XY64) XY64 {
	return XY64{v.X - o.X, v.Y - o.Y}
}

// Minuss returns the component-wise difference with two scalars.
func (v XY64) Minuss(x, y float64) XY64 {
	return XY64{v.X - x, v.Y - y}
}

// Opposite returns the opposite of the vector.
func (v XY64) Opposite() XY64 {
	return XY64{-v.X, -v.Y}
}

// Times returns the product with a scalar.
func (v XY64) Times(s float64) XY64 {
	return XY64{v.X * s, v.Y * s}
}

// Timess returns the component-wise product with two scalars.
func (v XY64) Timess(x, y float64) XY64 {
	return XY64{v.X * x, v.Y * y}
}

// Timescw returns the component-wise product with another vector.
func (v XY64) Timescw(o XY64) XY64 {
	return XY64{v.X * o.X, v.Y * o.Y}
}

// Slash returns the division by a scalar (which must be non-zero).
func (v XY64) Slash(s float64) XY64 {
	return XY64{v.X / s, v.Y / s}
}

// Slashs returns the component-wise division by two scalars (which must be
// non-zero).
func (v XY64) Slashs(x, y float64) XY64 {
	return XY64{v.X / x, v.Y / y}
}

// Slashcw returns the component-wise division by another vector (of which both
// X and Y must be non-zero).
func (v XY64) Slashcw(o XY64) XY64 {
	return XY64{v.X / o.X, v.Y / o.Y}
}

// Mod returns the remainder (modulus) of the division by a scalar (which must
// be non-zero).
func (v XY64) Mod(s float64) XY64 {
	return XY64{math.Mod(v.X, s), math.Mod(v.Y, s)}
}

// Mods returns the remainder (modulus) of the component-wise division by two
// scalars (which must be non-zero).
func (v XY64) Mods(x, y float64) XY64 {
	return XY64{math.Mod(v.X, x), math.Mod(v.Y, y)}
}

// Modcw returns the remainders (modulus) of the component-wise division by
// another vector (of which both X and Y must be non-zero).
func (v XY64) Modcw(o XY64) XY64 {
	return XY64{math.Mod(v.X, o.X), math.Mod(v.Y, o.Y)}
}

// Modf returns the integer part and the fractional part of (each component of)
// the vector.
func (v XY64) Modf() (intg, frac XY64) {
	xintg, xfrac := math.Modf(v.X)
	yintg, yfrac := math.Modf(v.Y)
	return XY64{xintg, yintg}, XY64{xfrac, yfrac}
}

// Dot returns the dot product with another vector.
func (v XY64) Dot(o XY64) float64 {
	return v.X*o.X + v.Y*o.Y
}

// PerpDot returns the dot product with the perpendicular of v and another
// vector.
func (v XY64) PerpDot(o XY64) float64 {
	return v.X*o.Y - v.Y*o.X
}

// Length returns the euclidian length of the vector.
func (v XY64) Length() float64 {
	// Double conversion is faster than math.Sqrt because the Go compiler
	// optimizes it.
	return float64(math.Sqrt(float64(v.X*v.X + v.Y*v.Y)))
}

// Length2 returns the square of the euclidian length of the vector.
func (v XY64) Length2() float64 {
	return v.X*v.X + v.Y*v.Y
}

// Distance returns the distance with another vector.
func (v XY64) Distance(o XY64) float64 {
	d := XY64{v.X - o.X, v.Y - o.Y}
	// Double conversion is faster than math.Sqrt because the Go compiler
	// optimizes it.
	return float64(math.Sqrt(float64(d.X*d.X + d.Y*d.Y)))
}

// Distance2 returns the square of the distance with another vector.
func (v XY64) Distance2(o XY64) float64 {
	d := XY64{v.X - o.X, v.Y - o.Y}
	return d.X*d.X + d.Y*d.Y
}

// Normalized return the normalization of the vector (i.e. the vector divided
// by its length, which must be non-zero).
func (v XY64) Normalized() XY64 {
	// Double conversion is faster than math.Sqrt because the Go compiler
	// optimizes it.
	l := float64(math.Sqrt(float64(v.X*v.X + v.Y*v.Y)))
	return XY64{v.X / l, v.Y / l}
}

// IsAlmostEqual returns true if the difference between the two vectors is less
// than the specified ULPs (Unit in the Last Place).
//
// Handle special cases: zero, infinites, denormals.
//
// See also IsNearlyEqual and IsRoughlyEqual.
//TODO:
// func (v Coord64) IsAlmostEqual(o Coord64, ulps uint32) bool {
// 	return math.IsAlmostEqual(v.X, o.X, ulps) &&
// 		math.IsAlmostEqual(v.Y, o.Y, ulps)
// }

// IsNearlyEqual Returns true if the relative error between the two vectors is
// less than epsilon.
//
// Handles special cases: zero, infinites, denormals.
//
// See also IsAlmostEqual and IsRoughlyEqual.
//TODO:
// func (v Coord64) IsNearlyEqual(o Coord64, epsilon float64) bool {
// 	return math.IsNearlyEqual(v.X, o.X, epsilon) &&
// 		math.IsNearlyEqual(v.Y, o.Y, epsilon)
// }

// IsRoughlyEqual Returns true if the absolute error between the two vectors is
// less than epsilon.
//
// See also IsNearlyEqual and IsAlmostEqual.
//TODO:
// func (v Coord64) IsRoughlyEqual(o Coord64, epsilon float64) bool {
// 	return math.IsRoughlyEqual(v.X, o.X, epsilon) &&
// 		math.IsRoughlyEqual(v.Y, o.Y, epsilon)
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
func (v XYZ64) Cartesian() (x, y float32) {
	return float32(v.X / v.Z), float32(v.Y / v.Z)
}

// XY64 returns the cartesian representation of the vector (i.e. the
// perspective divide of the homogeneous coordinates). Z must be non-zero.
func (v XYZ64) XY64() XY64 {
	return XY64{v.X / v.Z, v.Y / v.Z}
}

////////////////////////////////////////////////////////////////////////////////

// DA64 represents a two dimensional vector, defined by its polar coordinates.
type DA64 struct {
	D float64 // Distance from origin (i.e. radius)
	A float64 // Angle //TODO: what angle?
}

// Cartesian returns the cartesian coordinates of the vector. This implements the
// Vector interface.
func (v DA64) Cartesian() (x, y float32) {
	return float32(v.D * math.Cos(v.A)), float32(v.D * math.Sin(v.A))
}

// XY64 returns the cartesian representation of the vector.
func (v DA64) XY64() XY64 {
	return XY64{v.D * math.Cos(v.A), v.D * math.Sin(v.A)}
}

////////////////////////////////////////////////////////////////////////////////
