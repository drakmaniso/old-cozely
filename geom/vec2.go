// Copyright (c) 2013-2016 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package geom

import "github.com/drakmaniso/glam/math"

//------------------------------------------------------------------------------

// Vec2 is a single-precision vector with 2 components..
type Vec2 struct {
	X float32
	Y float32
}

//------------------------------------------------------------------------------

// Homogenized returns the homogeneous coordinates of the vector.
func (v Vec2) Homogenized() Vec3 {
	return Vec3{v.X, v.Y, 1.0}
}

// HomogenizedAsDirection returns the homogeneous coordinates
// of a point at infinity in the direction of the vector.
func (v Vec2) HomogenizedAsDirection() Vec3 {
	return Vec3{v.X, v.Y, 0.0}
}

//------------------------------------------------------------------------------

// Plus returns the sum with another vector.
func (v Vec2) Plus(b Vec2) Vec2 {
	return Vec2{v.X + b.X, v.Y + b.Y}
}

// Minus returns the difference with another vector.
func (v Vec2) Minus(b Vec2) Vec2 {
	return Vec2{v.X - b.X, v.Y - b.Y}
}

// Inverse return the inverse of the vector.
func (v Vec2) Inverse() Vec2 {
	return Vec2{-v.X, -v.Y}
}

// Times returns the product with a scalar.
func (v Vec2) Times(s float32) Vec2 {
	return Vec2{v.X * s, v.Y * s}
}

// Slash returns the division by a scalar (which must be non-zero).
func (v Vec2) Slash(s float32) Vec2 {
	return Vec2{v.X / s, v.Y / s}
}

//------------------------------------------------------------------------------

// Dot returns the dot product with another vector.
func (v Vec2) Dot(b Vec2) float32 {
	return v.X*b.X + v.Y*b.Y
}

//------------------------------------------------------------------------------

// Length returns the euclidian length of the vector.
func (v Vec2) Length() float32 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// Normalized return the normalization of the vector (i.e. the vector divided
// by its length).
//
// Important: Length() must be non-zero.
func (v Vec2) Normalized() Vec2 {
	length := math.Sqrt(v.X*v.X + v.Y*v.Y)
	return Vec2{v.X / length, v.Y / length}
}

//------------------------------------------------------------------------------

// IsAlmostEqual returns true if the difference between the two vectors is less
// than the specified ULPs (Unit in the Last Place).
//
// Handle special cases: zero, infinites, denormals.
//
// See also IsNearlyEqual and IsRoughlyEqual.
func (v Vec2) IsAlmostEqual(o Vec2, ulps uint32) bool {
	return math.IsAlmostEqual(v.X, o.X, ulps) &&
		math.IsAlmostEqual(v.Y, o.Y, ulps)
}

// IsNearlyEqual Returns true if the relative error between the two vectors is
// less than epsilon.
//
// Handles special cases: zero, infinites, denormals.
//
// See also IsAlmostEqual and IsRoughlyEqual.
func (v Vec2) IsNearlyEqual(o Vec2, epsilon float32) bool {
	return math.IsNearlyEqual(v.X, o.X, epsilon) &&
		math.IsNearlyEqual(v.Y, o.Y, epsilon)
}

// IsRoughlyEqual Returns true if the absolute error between the two vectors is
// less than epsilon.
//
// See also IsNearlyEqual and IsAlmostEqual.
func (v Vec2) IsRoughlyEqual(o Vec2, epsilon float32) bool {
	return math.IsRoughlyEqual(v.X, o.X, epsilon) &&
		math.IsRoughlyEqual(v.Y, o.Y, epsilon)
}

//------------------------------------------------------------------------------
