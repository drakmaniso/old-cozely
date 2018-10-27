// Copyright (c) 2013-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package coord

import (
	"math"

	"github.com/cozely/cozely/x/math32"
)

////////////////////////////////////////////////////////////////////////////////

// XYZ represents either 3D cartesian coordinates, or 2D projective
// (homogeneous) coordinates.
type XYZ struct {
	X float32
	Y float32
	Z float32
}

// Cartesian returns the three coordinates.
func (a XYZ) Cartesian() (x, y, z float32) {
	return a.X, a.Y, a.Z
}

// XYZW returns the homogenous coordinates of the vector, with W set to w.
func (a XYZ) XYZW(w float32) XYZW {
	return XYZW{a.X, a.Y, a.Z, w}
}

// Plus returns the sum with another vector.
func (a XYZ) Plus(b XYZ) XYZ {
	return XYZ{a.X + b.X, a.Y + b.Y, a.Z + b.Z}
}

// Pluss returns the sum with a scalar.
func (a XYZ) Pluss(s float32) XYZ {
	return XYZ{a.X + s, a.Y + s, a.Z + s}
}

// Minus returns the difference with another vector.
func (a XYZ) Minus(b XYZ) XYZ {
	return XYZ{a.X - b.X, a.Y - b.Y, a.Z - b.Z}
}

// Minuss returns the sum with a scalar.
func (a XYZ) Minuss(s float32) XYZ {
	return XYZ{a.X - s, a.Y - s, a.Z - s}
}

// Opposite returns the opposite of the vector.
func (a XYZ) Opposite() XYZ {
	return XYZ{-a.X, -a.Y, -a.Z}
}

// Timess returns the product with a scalar.
func (a XYZ) Timess(s float32) XYZ {
	return XYZ{a.X * s, a.Y * s, a.Z * s}
}

// Times returns the component-wise product with another vector.
func (a XYZ) Times(b XYZ) XYZ {
	return XYZ{a.X * b.X, a.Y * b.Y, a.Z * b.Z}
}

// Slashs returns the division by a scalar (which must be non-zero).
func (a XYZ) Slashs(s float32) XYZ {
	return XYZ{a.X / s, a.Y / s, a.Z / s}
}

// Slash returns the component-wise division by another vector (of which X, Y
// and Z must be non-zero).
func (a XYZ) Slash(b XYZ) XYZ {
	return XYZ{a.X / b.X, a.Y / b.Y, a.Z / b.Z}
}

// Mods returns the remainder (modulus) of the division by a scalar (which must
// be non-zero).
func (a XYZ) Mods(s float32) XYZ {
	return XYZ{math32.Mod(a.X, s), math32.Mod(a.Y, s), math32.Mod(a.Z, s)}
}

// Mod returns the remainders (modulus) of the component-wise division by
// another vector (of which X, Y and Z must be non-zero).
func (a XYZ) Mod(b XYZ) XYZ {
	return XYZ{math32.Mod(a.X, b.X), math32.Mod(a.Y, b.Y), math32.Mod(a.Z, b.Z)}
}

// Modf returns the integer part and the fractional part of (each component of)
// the vector.
func (a XYZ) Modf() (intg, frac XYZ) {
	xintg, xfrac := math32.Modf(a.X)
	yintg, yfrac := math32.Modf(a.Y)
	zintg, zfrac := math32.Modf(a.Z)
	return XYZ{xintg, yintg, zintg}, XYZ{xfrac, yfrac, zfrac}
}

// Dot returns the dot product with another vector.
func (a XYZ) Dot(b XYZ) float32 {
	return a.X*b.X + a.Y*b.Y + a.Z*b.Z
}

// Cross returns the cross product with another vector.
func (a XYZ) Cross(b XYZ) XYZ {
	return XYZ{
		a.Y*b.Z - a.Z*b.Y,
		a.Z*b.X - a.X*b.Z,
		a.X*b.Y - a.Y*b.X,
	}
}

// Length returns the euclidian length of the vector.
func (a XYZ) Length() float32 {
	// Double conversion is faster than math32.Sqrt because the Go compiler
	// optimizes it.
	return float32(math.Sqrt(float64(a.X*a.X + a.Y*a.Y + a.Z*a.Z)))
}

// Length2 returns the square of the euclidian length of the vector.
func (a XYZ) Length2() float32 {
	return a.X*a.X + a.Y*a.Y + a.Z*a.Z
}

// Distance returns the distance with another vector.
func (a XYZ) Distance(b XYZ) float32 {
	d := XYZ{a.X - b.X, a.Y - b.Y, a.Z - b.Z}
	return float32(math.Sqrt(float64(d.X*d.X + d.Y*d.Y + d.Z*d.Z)))
}

// Distance2 returns the square of the distance with another vector.
func (a XYZ) Distance2(b XYZ) float32 {
	d := XYZ{a.X - b.X, a.Y - b.Y, a.Z - b.Z}
	return d.X*d.X + d.Y*d.Y + d.Z*d.Z
}

// Normalized return the normalization of the vector (i.e. the vector divided
// by its length, which must be non-zero).
func (a XYZ) Normalized() XYZ {
	// Double conversion is faster than math32.Sqrt because the Go compiler
	// optimizes it.
	l := float32(math.Sqrt(float64(a.X*a.X + a.Y*a.Y + a.Z*a.Z)))
	return XYZ{a.X / l, a.Y / l, a.Z / l}
}

// IsAlmostEqual returns true if the difference between the two vectors is less
// than the specified ULPs (Unit in the Last Place).
//
// Handle special cases: zero, infinites, denormals.
//
// See also IsNearlyEqual and IsRoughlyEqual.
func (a XYZ) IsAlmostEqual(b XYZ, ulps uint32) bool {
	return math32.IsAlmostEqual(a.X, b.X, ulps) &&
		math32.IsAlmostEqual(a.Y, b.Y, ulps) &&
		math32.IsAlmostEqual(a.Z, b.Z, ulps)
}

// IsNearlyEqual Returns true if the relative error between the two vectors is
// less than epsilon.
//
// Handles special cases: zero, infinites, denormals.
//
// See also IsAlmostEqual and IsRoughlyEqual.
func (a XYZ) IsNearlyEqual(b XYZ, epsilon float32) bool {
	return math32.IsNearlyEqual(a.X, b.X, epsilon) &&
		math32.IsNearlyEqual(a.Y, b.Y, epsilon) &&
		math32.IsNearlyEqual(a.Z, b.Z, epsilon)
}

// IsRoughlyEqual Returns true if the absolute error between the two vectors is
// less than epsilon.
//
// See also IsNearlyEqual and IsAlmostEqual.
func (a XYZ) IsRoughlyEqual(b XYZ, epsilon float32) bool {
	return math32.IsRoughlyEqual(a.X, b.X, epsilon) &&
		math32.IsRoughlyEqual(a.Y, b.Y, epsilon) &&
		math32.IsRoughlyEqual(a.Z, b.Z, epsilon)
}

////////////////////////////////////////////////////////////////////////////////

// XYproj returns the cartesian representation of the vector (i.e. the perspective
// divide of the homogeneous coordinates). Z must be non-zero.
func (a XYZ) XYproj() XY {
	return XY{a.X / a.Z, a.Y / a.Z}
}

// XY returns the planar coordinates {X, Y}.
func (a XYZ) XY() XY {
	return XY{a.X, a.Y}
}

// XZ returns the planar coordinates {X, Z}.
func (a XYZ) XZ() XY {
	return XY{a.X, a.Z}
}

// YX returns the planar coordinates {Y, X}.
func (a XYZ) YX() XY {
	return XY{a.Y, a.X}
}

// YZ returns the planar coordinates {Y, Z}.
func (a XYZ) YZ() XY {
	return XY{a.Y, a.Z}
}

// ZX returns the planar coordinates {Z, X}.
func (a XYZ) ZX() XY {
	return XY{a.Z, a.X}
}

// ZY returns the planar coordinates {Z, Y}.
func (a XYZ) ZY() XY {
	return XY{a.Z, a.Y}
}

////////////////////////////////////////////////////////////////////////////////
