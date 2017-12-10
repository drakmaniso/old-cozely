// Copyright (c) 2013-2017 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package plane

//------------------------------------------------------------------------------

import (
	"math"

	"github.com/drakmaniso/carol/core/math32"
)

//------------------------------------------------------------------------------

// Vector represents any two-dimensional vector.
type Vector interface {
	// XY returns the cartesian coordinates of the vector.
	XY() (x, y float32)
}

//------------------------------------------------------------------------------

// Coord represents a two-dimensional vector, defined by its cartesian
// coordinates.
type Coord struct {
	X float32
	Y float32
}

// XY returns the cartesian coordinates of the vector.
//
// This function implements the Vector interface.
func (v Coord) XY() (x, y float32) {
	return v.X, v.Y
}

// CoordOf returns the cartesian representation of v.
func CoordOf(v Vector) Coord {
	x, y := v.XY()
	return Coord{x, y}
}

// Homogen returns the homogenous coordinates of the vector, with Z set to 1.
func (v Coord) Homogen() Homogen {
	return Homogen{v.X, v.Y, 1.0}
}

// Plus returns the sum with another vector.
func (v Coord) Plus(o Coord) Coord {
	return Coord{v.X + o.X, v.Y + o.Y}
}

// Minus returns the difference with another vector.
func (v Coord) Minus(o Coord) Coord {
	return Coord{v.X - o.X, v.Y - o.Y}
}

// Opposite returns the opposite of the vector.
func (v Coord) Opposite() Coord {
	return Coord{-v.X, -v.Y}
}

// Times returns the product with a scalar.
func (v Coord) Times(s float32) Coord {
	return Coord{v.X * s, v.Y * s}
}

// TimesCW returns the component-wise product with another vector.
func (v Coord) TimesCW(o Coord) Coord {
	return Coord{v.X * o.X, v.Y * o.Y}
}

// Slash returns the division by a scalar (which must be non-zero).
func (v Coord) Slash(s float32) Coord {
	return Coord{v.X / s, v.Y / s}
}

// SlashCW returns the component-wise division by another vector (of which both
// X and Y must be non-zero).
func (v Coord) SlashCW(o Coord) Coord {
	return Coord{v.X / o.X, v.Y / o.Y}
}

// Mod returns the remainder (modulus) of the division by a scalar (which must
// be non-zero).
func (v Coord) Mod(s float32) Coord {
	return Coord{math32.Mod(v.X, s), math32.Mod(v.Y, s)}
}

// ModCW returns the remainders (modulus) of the component-wise division by
// another vector (of which both X and Y must be non-zero).
func (v Coord) ModCW(o Coord) Coord {
	return Coord{math32.Mod(v.X, o.X), math32.Mod(v.Y, o.Y)}
}

// Modf returns the integer part and the fractional part of (each component of)
// the vector.
func (v Coord) Modf() (intg, frac Coord) {
	xintg, xfrac := math32.Modf(v.X)
	yintg, yfrac := math32.Modf(v.Y)
	return Coord{xintg, yintg}, Coord{xfrac, yfrac}
}

// Dot returns the dot product with another vector.
func (v Coord) Dot(o Coord) float32 {
	return v.X*o.X + v.Y*o.Y
}

// Length returns the euclidian length of the vector.
func (v Coord) Length() float32 {
	// Double conversion is faster than math32.Sqrt because the Go compiler
	// optimizes it.
	return float32(math.Sqrt(float64(v.X*v.X + v.Y*v.Y)))
}

// Length2 returns the square of the euclidian length of the vector.
func (v Coord) Length2() float32 {
	return v.X*v.X + v.Y*v.Y
}

// Distance returns the distance with another vector.
func (v Coord) Distance(o Coord) float32 {
	d := Coord{v.X - o.X, v.Y - o.Y}
	// Double conversion is faster than math32.Sqrt because the Go compiler
	// optimizes it.
	return float32(math.Sqrt(float64(d.X*d.X + d.Y*d.Y)))
}

// Distance2 returns the square of the distance with another vector.
func (v Coord) Distance2(o Coord) float32 {
	d := Coord{v.X - o.X, v.Y - o.Y}
	return d.X*d.X + d.Y*d.Y
}

// Normalized return the normalization of the vector (i.e. the vector divided
// by its length, which must be non-zero).
func (v Coord) Normalized() Coord {
	// Double conversion is faster than math32.Sqrt because the Go compiler
	// optimizes it.
	l := float32(math.Sqrt(float64(v.X*v.X + v.Y*v.Y)))
	return Coord{v.X / l, v.Y / l}
}

// IsAlmostEqual returns true if the difference between the two vectors is less
// than the specified ULPs (Unit in the Last Place).
//
// Handle special cases: zero, infinites, denormals.
//
// See also IsNearlyEqual and IsRoughlyEqual.
func (v Coord) IsAlmostEqual(o Coord, ulps uint32) bool {
	return math32.IsAlmostEqual(v.X, o.X, ulps) &&
		math32.IsAlmostEqual(v.Y, o.Y, ulps)
}

// IsNearlyEqual Returns true if the relative error between the two vectors is
// less than epsilon.
//
// Handles special cases: zero, infinites, denormals.
//
// See also IsAlmostEqual and IsRoughlyEqual.
func (v Coord) IsNearlyEqual(o Coord, epsilon float32) bool {
	return math32.IsNearlyEqual(v.X, o.X, epsilon) &&
		math32.IsNearlyEqual(v.Y, o.Y, epsilon)
}

// IsRoughlyEqual Returns true if the absolute error between the two vectors is
// less than epsilon.
//
// See also IsNearlyEqual and IsAlmostEqual.
func (v Coord) IsRoughlyEqual(o Coord, epsilon float32) bool {
	return math32.IsRoughlyEqual(v.X, o.X, epsilon) &&
		math32.IsRoughlyEqual(v.Y, o.Y, epsilon)
}

//------------------------------------------------------------------------------

// Homogen represents a two-dimensional vector, defined by its homogeneous
// coordinates.
type Homogen struct {
	X float32
	Y float32
	Z float32
}

// XY returns the cartesian coordinates of the vector (i.e. the perspective
// divide of the homogeneous coordinates). Z must be non-zero.
func (v Homogen) XY() (x, y float32) {
	return v.X / v.Z, v.Y / v.Z
}

// Coord returns the cartesian representation of the vector (i.e. the
// perspective divide of the homogeneous coordinates). Z must be non-zero.
func (v Homogen) Coord() Coord {
	return Coord{v.X / v.Z, v.Y / v.Z}
}

//------------------------------------------------------------------------------

// Polar represents a two dimensional vector, defined by its polar coordinates.
type Polar struct {
	R     float32 // Radius (i.e. distance from origin)
	Theta float32 // Angle //TODO: what angle?
}

// XY returns the cartesian coordinates of the vector. This implements the
// Vector interface.
func (v Polar) XY() (x, y float32) {
	return v.R * math32.Cos(v.Theta), v.R * math32.Sin(v.Theta)
}

// Coord returns the cartesian representation of the vector.
func (v Polar) Coord() Coord {
	return Coord{v.R * math32.Cos(v.Theta), v.R * math32.Sin(v.Theta)}
}

//------------------------------------------------------------------------------
