// Copyright (c) 2013-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package coord

import (
	"math"

	"github.com/cozely/cozely/x/math32"
)

////////////////////////////////////////////////////////////////////////////////

// Vector3D represents any three-dimensional vector.
type Vector3D interface {
	// Cartesian returns the cartesian coordinates of the vector.
	Cartesian() (x, y, z float32)
}

////////////////////////////////////////////////////////////////////////////////

// XYZ represents a three-dimensional vector, defined by its cartesian
// coordinates.
type XYZ struct {
	X float32
	Y float32
	Z float32
}

// Cartesian returns the cartesian coordinates of the vector.
//
// This implements the Vector interface.
func (v XYZ) Cartesian() (x, y, z float32) {
	return v.X, v.Y, v.Z
}

// XYZW returns the homogenous coordinates of the vector, with W set to w.
func (v XYZ) XYZW(w float32) XYZW {
	return XYZW{v.X, v.Y, v.Z, w}
}

// Plus returns the sum with another vector.
func (v XYZ) Plus(o XYZ) XYZ {
	return XYZ{v.X + o.X, v.Y + o.Y, v.Z + o.Z}
}

// Minus returns the difference with another vector.
func (v XYZ) Minus(o XYZ) XYZ {
	return XYZ{v.X - o.X, v.Y - o.Y, v.Z - o.Z}
}

// Opposite returns the opposite of the vector.
func (v XYZ) Opposite() XYZ {
	return XYZ{-v.X, -v.Y, -v.Z}
}

// Times returns the product with a scalar.
func (v XYZ) Times(s float32) XYZ {
	return XYZ{v.X * s, v.Y * s, v.Z * s}
}

// TimesCW returns the component-wise product with another vector.
func (v XYZ) TimesCW(o XYZ) XYZ {
	return XYZ{v.X * o.X, v.Y * o.Y, v.Z * o.Z}
}

// Slash returns the division by a scalar (which must be non-zero).
func (v XYZ) Slash(s float32) XYZ {
	return XYZ{v.X / s, v.Y / s, v.Z / s}
}

// SlashCW returns the component-wise division by another vector (of which X, Y
// and Z must be non-zero).
func (v XYZ) SlashCW(o XYZ) XYZ {
	return XYZ{v.X / o.X, v.Y / o.Y, v.Z / o.Z}
}

// Mod returns the remainder (modulus) of the division by a scalar (which must
// be non-zero).
func (v XYZ) Mod(s float32) XYZ {
	return XYZ{math32.Mod(v.X, s), math32.Mod(v.Y, s), math32.Mod(v.Z, s)}
}

// ModCW returns the remainders (modulus) of the component-wise division by
// another vector (of which X, Y and Z must be non-zero).
func (v XYZ) ModCW(o XYZ) XYZ {
	return XYZ{math32.Mod(v.X, o.X), math32.Mod(v.Y, o.Y), math32.Mod(v.Z, o.Z)}
}

// Modf returns the integer part and the fractional part of (each component of)
// the vector.
func (v XYZ) Modf() (intg, frac XYZ) {
	xintg, xfrac := math32.Modf(v.X)
	yintg, yfrac := math32.Modf(v.Y)
	zintg, zfrac := math32.Modf(v.Z)
	return XYZ{xintg, yintg, zintg}, XYZ{xfrac, yfrac, zfrac}
}

// Dot returns the dot product with another vector.
func (v XYZ) Dot(o XYZ) float32 {
	return v.X*o.X + v.Y*o.Y + v.Z*o.Z
}

// Cross returns the cross product with another vector.
func (v XYZ) Cross(o XYZ) XYZ {
	return XYZ{
		v.Y*o.Z - v.Z*o.Y,
		v.Z*o.X - v.X*o.Z,
		v.X*o.Y - v.Y*o.X,
	}
}

// Length returns the euclidian length of the vector.
func (v XYZ) Length() float32 {
	// Double conversion is faster than math32.Sqrt because the Go compiler
	// optimizes it.
	return float32(math.Sqrt(float64(v.X*v.X + v.Y*v.Y + v.Z*v.Z)))
}

// Length2 returns the square of the euclidian length of the vector.
func (v XYZ) Length2() float32 {
	return v.X*v.X + v.Y*v.Y + v.Z*v.Z
}

// Distance returns the distance with another vector.
func (v XYZ) Distance(o XYZ) float32 {
	d := XYZ{v.X - o.X, v.Y - o.Y, v.Z - o.Z}
	return float32(math.Sqrt(float64(d.X*d.X + d.Y*d.Y + d.Z*d.Z)))
}

// Distance2 returns the square of the distance with another vector.
func (v XYZ) Distance2(o XYZ) float32 {
	d := XYZ{v.X - o.X, v.Y - o.Y, v.Z - o.Z}
	return d.X*d.X + d.Y*d.Y + d.Z*d.Z
}

// Normalized return the normalization of the vector (i.e. the vector divided
// by its length, which must be non-zero).
func (v XYZ) Normalized() XYZ {
	// Double conversion is faster than math32.Sqrt because the Go compiler
	// optimizes it.
	l := float32(math.Sqrt(float64(v.X*v.X + v.Y*v.Y + v.Z*v.Z)))
	return XYZ{v.X / l, v.Y / l, v.Z / l}
}

// IsAlmostEqual returns true if the difference between the two vectors is less
// than the specified ULPs (Unit in the Last Place).
//
// Handle special cases: zero, infinites, denormals.
//
// See also IsNearlyEqual and IsRoughlyEqual.
func (v XYZ) IsAlmostEqual(o XYZ, ulps uint32) bool {
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
func (v XYZ) IsNearlyEqual(o XYZ, epsilon float32) bool {
	return math32.IsNearlyEqual(v.X, o.X, epsilon) &&
		math32.IsNearlyEqual(v.Y, o.Y, epsilon) &&
		math32.IsNearlyEqual(v.Z, o.Z, epsilon)
}

// IsRoughlyEqual Returns true if the absolute error between the two vectors is
// less than epsilon.
//
// See also IsNearlyEqual and IsAlmostEqual.
func (v XYZ) IsRoughlyEqual(o XYZ, epsilon float32) bool {
	return math32.IsRoughlyEqual(v.X, o.X, epsilon) &&
		math32.IsRoughlyEqual(v.Y, o.Y, epsilon) &&
		math32.IsRoughlyEqual(v.Z, o.Z, epsilon)
}

////////////////////////////////////////////////////////////////////////////////

// XYproj returns the cartesian representation of the vector (i.e. the perspective
// divide of the homogeneous coordinates). Z must be non-zero.
func (h XYZ) XYproj() XY {
	return XY{h.X / h.Z, h.Y / h.Z}
}

// XY returns the planar coordinates {X, Y}.
func (h XYZ) XY() XY {
	return XY{h.X, h.Y}
}

// XZ returns the planar coordinates {X, Z}.
func (h XYZ) XZ() XY {
	return XY{h.X, h.Z}
}

// YX returns the planar coordinates {Y, X}.
func (h XYZ) YX() XY {
	return XY{h.Y, h.X}
}

// YZ returns the planar coordinates {Y, Z}.
func (h XYZ) YZ() XY {
	return XY{h.Y, h.Z}
}

// ZX returns the planar coordinates {Z, X}.
func (h XYZ) ZX() XY {
	return XY{h.Z, h.X}
}

// ZY returns the planar coordinates {Z, Y}.
func (h XYZ) ZY() XY {
	return XY{h.Z, h.Y}
}

////////////////////////////////////////////////////////////////////////////////

// XYZW represents a three-dimensional vector, defined by its homogeneous
// coordinates.
type XYZW struct {
	X float32
	Y float32
	Z float32
	W float32
}

// Cartesian returns the cartesian coordinates of the vector (i.e. the perspective
// divide of the homogeneous coordinates). W must be non-zero.
func (v XYZW) Cartesian() (x, y, z float32) {
	return v.X / v.W, v.Y / v.W, v.Z / v.W
}

// XYZ returns the cartesian representation of the vector (i.e. the
// perspective divide of the homogeneous coordinates). W must be non-zero.
func (v XYZW) XYZ() XYZ {
	return XYZ{v.X / v.W, v.Y / v.W, v.Z / v.W}
}

////////////////////////////////////////////////////////////////////////////////

//TODO: Spherical and Cylindrical types

////////////////////////////////////////////////////////////////////////////////
