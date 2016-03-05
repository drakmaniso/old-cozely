// Copyright (c) 2013-2016 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package geom

import "github.com/drakmaniso/glam/math"

//------------------------------------------------------------------------------

// Vec3 is a single-precision vector with 3 components.
type Vec3 struct {
	X float32
	Y float32
	Z float32
}

//------------------------------------------------------------------------------

// Homogenized returns the homogeneous coordinates of the vector.
func (v Vec3) Homogenized() Vec4 {
	return Vec4{v.X, v.Y, v.Z, 1.0}
}

// HomogenizedAsDirection returns the homogeneous coordinates
// of a point at infinity in the direction of the vector.
func (v Vec3) HomogenizedAsDirection() Vec4 {
	return Vec4{v.X, v.Y, v.Z, 0.0}
}

// Dehomogenized returns the dehomogenization of the vector (i.e. perspective
// divide).
//
// Important: Z must be non-zero.
func (v Vec3) Dehomogenized() Vec2 {
	return Vec2{v.X / v.Z, v.Y / v.Z}
}

// XY returns a 2D vector made of X, and Y.
func (v Vec3) XY() Vec2 {
	return Vec2{v.X, v.Y}
}

//------------------------------------------------------------------------------

// Plus returns the sum with another vector.
func (v Vec3) Plus(b Vec3) Vec3 {
	return Vec3{v.X + b.X, v.Y + b.Y, v.Z + b.Z}
}

// Minus returns the difference with another vector.
func (v Vec3) Minus(b Vec3) Vec3 {
	return Vec3{v.X - b.X, v.Y - b.Y, v.Z - b.Z}
}

// Inverse return the inverse of the vector.
func (v Vec3) Inverse() Vec3 {
	return Vec3{-v.X, -v.Y, -v.Z}
}

// Times returns the product with a scalar.
func (v Vec3) Times(s float32) Vec3 {
	return Vec3{v.X * s, v.Y * s, v.Z * s}
}

// Slash returns the division by a scalar.
//
// Important: the scalar must be non-zero.
func (v Vec3) Slash(s float32) Vec3 {
	return Vec3{v.X / s, v.Y / s, v.Z / s}
}

//------------------------------------------------------------------------------

// Cross returns the cross product with another vector.
func (v Vec3) Cross(b Vec3) Vec3 {
	return Vec3{
		v.Y*b.Z - v.Z*b.Y,
		v.Z*b.X - v.X*b.Z,
		v.X*b.Y - v.Y*b.X,
	}
}

// Dot returns the dot product with another vector.
func (v Vec3) Dot(b Vec3) float32 {
	return v.X*b.X + v.Y*b.Y + v.Z*b.Z
}

//------------------------------------------------------------------------------

// Length returns the euclidian length.
func (v Vec3) Length() float32 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y + v.Z*v.Z)
}

// Normalized return the normalization of the vector (i.e. the vector divided
// by its length).
//
// Important: Length() must be non-zero.
func (v Vec3) Normalized() Vec3 {
	length := math.Sqrt(v.X*v.X + v.Y*v.Y + v.Z*v.Z)
	return Vec3{v.X / length, v.Y / length, v.Z / length}
}

//------------------------------------------------------------------------------

// IsAlmostEqual returns true if the difference between the two vectors is less
// than the specified ULPs (Unit in the Last Place).
//
// Handle special cases: zero, infinites, denormals.
//
// See also IsNearlyEqual and IsRoughlyEqual.
func (v Vec3) IsAlmostEqual(o Vec3, ulps uint32) bool {
	return math.IsAlmostEqual(v.X, o.X, ulps) &&
		math.IsAlmostEqual(v.Y, o.Y, ulps) &&
		math.IsAlmostEqual(v.Z, o.Z, ulps)
}

// IsNearlyEqual Returns true if the relative error between the two vectors is
// less than epsilon.
//
// Handles special cases: zero, infinites, denormals.
//
// See also IsAlmostEqual and IsRoughlyEqual.
func (v Vec3) IsNearlyEqual(o Vec3, epsilon float32) bool {
	return math.IsNearlyEqual(v.X, o.X, epsilon) &&
		math.IsNearlyEqual(v.Y, o.Y, epsilon) &&
		math.IsNearlyEqual(v.Z, o.Z, epsilon)
}

// IsRoughlyEqual Returns true if the absolute error between the two vectors is
// less than epsilon.
//
// See also IsNearlyEqual and IsAlmostEqual.
func (v Vec3) IsRoughlyEqual(o Vec3, epsilon float32) bool {
	return math.IsRoughlyEqual(v.X, o.X, epsilon) &&
		math.IsRoughlyEqual(v.Y, o.Y, epsilon) &&
		math.IsRoughlyEqual(v.Z, o.Z, epsilon)
}

//------------------------------------------------------------------------------
