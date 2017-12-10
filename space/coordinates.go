// Copyright (c) 2013-2017 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package space

//------------------------------------------------------------------------------

import (
	"math"

	"github.com/drakmaniso/glam/math32"
)

//------------------------------------------------------------------------------

// Vector represents any three-dimensional vector.
type Vector interface {
	// XYZ returns the cartesian coordinates of the vector.
	XYZ() (x, y, z float32)
}

//------------------------------------------------------------------------------

// Coord represents a three-dimensional vector, defined by its cartesian
// coordinates.
type Coord struct {
	X float32
	Y float32
	Z float32
}

// XYZ returns the cartesian coordinates of the vector.
//
// This implements the Vector interface.
func (v Coord) XYZ() (x, y, z float32) {
	return v.X, v.Y, v.Z
}

// CoordOf returns the cartesian representation of v.
func CoordOf(v Vector) Coord {
	x, y, z := v.XYZ()
	return Coord{x, y, z}
}

// Homogen returns the homogenous coordinates of the vector, with W set to 1.
func (v Coord) Homogen() Homogen {
	return Homogen{v.X, v.Y, v.Z, 1.0}
}

// Plus returns the sum with another vector.
func (v Coord) Plus(o Coord) Coord {
	return Coord{v.X + o.X, v.Y + o.Y, v.Z + o.Z}
}

// Minus returns the difference with another vector.
func (v Coord) Minus(o Coord) Coord {
	return Coord{v.X - o.X, v.Y - o.Y, v.Z - o.Z}
}

// Opposite returns the opposite of the vector.
func (v Coord) Opposite() Coord {
	return Coord{-v.X, -v.Y, -v.Z}
}

// Times returns the product with a scalar.
func (v Coord) Times(s float32) Coord {
	return Coord{v.X * s, v.Y * s, v.Z * s}
}

// TimesCW returns the component-wise product with another vector.
func (v Coord) TimesCW(o Coord) Coord {
	return Coord{v.X * o.X, v.Y * o.Y, v.Z * o.Z}
}

// Slash returns the division by a scalar (which must be non-zero).
func (v Coord) Slash(s float32) Coord {
	return Coord{v.X / s, v.Y / s, v.Z / s}
}

// SlashCW returns the component-wise division by another vector (of which X, Y
// and Z must be non-zero).
func (v Coord) SlashCW(o Coord) Coord {
	return Coord{v.X / o.X, v.Y / o.Y, v.Z / o.Z}
}

// Mod returns the remainder (modulus) of the division by a scalar (which must
// be non-zero).
func (v Coord) Mod(s float32) Coord {
	return Coord{math32.Mod(v.X, s), math32.Mod(v.Y, s), math32.Mod(v.Z, s)}
}

// ModCW returns the remainders (modulus) of the component-wise division by
// another vector (of which X, Y and Z must be non-zero).
func (v Coord) ModCW(o Coord) Coord {
	return Coord{math32.Mod(v.X, o.X), math32.Mod(v.Y, o.Y), math32.Mod(v.Z, o.Z)}
}

// Modf returns the integer part and the fractional part of (each component of)
// the vector.
func (v Coord) Modf() (intg, frac Coord) {
	xintg, xfrac := math32.Modf(v.X)
	yintg, yfrac := math32.Modf(v.Y)
	zintg, zfrac := math32.Modf(v.Z)
	return Coord{xintg, yintg, zintg}, Coord{xfrac, yfrac, zfrac}
}

// Dot returns the dot product with another vector.
func (v Coord) Dot(o Coord) float32 {
	return v.X*o.X + v.Y*o.Y + v.Z*o.Z
}

// Cross returns the cross product with another vector.
func (v Coord) Cross(o Coord) Coord {
	return Coord{
		v.Y*o.Z - v.Z*o.Y,
		v.Z*o.X - v.X*o.Z,
		v.X*o.Y - v.Y*o.X,
	}
}

// Length returns the euclidian length of the vector.
func (v Coord) Length() float32 {
	// Double conversion is faster than math32.Sqrt because the Go compiler
	// optimizes it.
	return float32(math.Sqrt(float64(v.X*v.X + v.Y*v.Y + v.Z*v.Z)))
}

// Length2 returns the square of the euclidian length of the vector.
func (v Coord) Length2() float32 {
	return v.X*v.X + v.Y*v.Y + v.Z*v.Z
}

// Distance returns the distance with another vector.
func (v Coord) Distance(o Coord) float32 {
	d := Coord{v.X - o.X, v.Y - o.Y, v.Z - o.Z}
	return float32(math.Sqrt(float64(d.X*d.X + d.Y*d.Y + d.Z*d.Z)))
}

// Distance2 returns the square of the distance with another vector.
func (v Coord) Distance2(o Coord) float32 {
	d := Coord{v.X - o.X, v.Y - o.Y, v.Z - o.Z}
	return d.X*d.X + d.Y*d.Y + d.Z*d.Z
}

// Normalized return the normalization of the vector (i.e. the vector divided
// by its length, which must be non-zero).
func (v Coord) Normalized() Coord {
	// Double conversion is faster than math32.Sqrt because the Go compiler
	// optimizes it.
	l := float32(math.Sqrt(float64(v.X*v.X + v.Y*v.Y + v.Z*v.Z)))
	return Coord{v.X / l, v.Y / l, v.Z / l}
}

// IsAlmostEqual returns true if the difference between the two vectors is less
// than the specified ULPs (Unit in the Last Place).
//
// Handle special cases: zero, infinites, denormals.
//
// See also IsNearlyEqual and IsRoughlyEqual.
func (v Coord) IsAlmostEqual(o Coord, ulps uint32) bool {
	return math32.IsAlmostEqual(v.X, o.X, ulps) &&
		math32.IsAlmostEqual(v.Y, o.Y, ulps) &&
		math32.IsAlmostEqual(v.Z, o.Z, ulps)
}

// IsNearlyEqual Returns true if the relative error between the two vectors is
// less than epsilon.
//
// Handles special cases: zero, infinites, denormals.
//
// See also IsAlmostEqual and IsRoughlyEqual.
func (v Coord) IsNearlyEqual(o Coord, epsilon float32) bool {
	return math32.IsNearlyEqual(v.X, o.X, epsilon) &&
		math32.IsNearlyEqual(v.Y, o.Y, epsilon) &&
		math32.IsNearlyEqual(v.Z, o.Z, epsilon)
}

// IsRoughlyEqual Returns true if the absolute error between the two vectors is
// less than epsilon.
//
// See also IsNearlyEqual and IsAlmostEqual.
func (v Coord) IsRoughlyEqual(o Coord, epsilon float32) bool {
	return math32.IsRoughlyEqual(v.X, o.X, epsilon) &&
		math32.IsRoughlyEqual(v.Y, o.Y, epsilon) &&
		math32.IsRoughlyEqual(v.Z, o.Z, epsilon)
}

//------------------------------------------------------------------------------

// Homogen represents a three-dimensional vector, defined by its homogeneous
// coordinates.
type Homogen struct {
	X float32
	Y float32
	Z float32
	W float32
}

// XYZ returns the cartesian coordinates of the vector (i.e. the perspective
// divide of the homogeneous coordinates). W must be non-zero.
func (v Homogen) XYZ() (x, y, z float32) {
	return v.X / v.W, v.Y / v.W, v.Z / v.W
}

// Coord returns the cartesian representation of the vector (i.e. the
// perspective divide of the homogeneous coordinates). W must be non-zero.
func (v Homogen) Coord() Coord {
	return Coord{v.X / v.W, v.Y / v.W, v.Z / v.W}
}

//------------------------------------------------------------------------------

//TODO: Spherical and Cylindrical types

//------------------------------------------------------------------------------
