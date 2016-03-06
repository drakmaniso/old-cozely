// Copyright (c) 2013-2016 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package geom

import "github.com/drakmaniso/glam/math"

//------------------------------------------------------------------------------

// Vec4 is a single-precision vector with 4 components.
//
// Note that in order to oblige the "go vet" tool, you should write vectors
// literals like this:
//  v := geom.Vec4{X: 1, Y: 2, Z: 3, W: 4} 
type Vec4 struct {
	X float32
	Y float32
	Z float32
	W float32
}

//------------------------------------------------------------------------------

// Dehomogenized returns the dehomogenization of the vector (i.e. perspective
// divide).
//
// Important: W must be non-zero.
func (v Vec4) Dehomogenized() Vec3 {
	return Vec3{v.X / v.W, v.Y / v.W, v.Z / v.W}
}

// XYZ returns a 3D vector made of X, Y and Z.
func (v Vec4) XYZ() Vec3 {
	return Vec3{v.X, v.Y, v.Z}
}

// XY returns a 2D vector made of X, and Y.
func (v Vec4) XY() Vec2 {
	return Vec2{v.X, v.Y}
}

//------------------------------------------------------------------------------

// Plus returns the sum with another vector.
func (v Vec4) Plus(o Vec4) Vec4 {
	return Vec4{v.X + o.X, v.Y + o.Y, v.Z + o.Z, v.W + o.W}
}

// Minus returns the difference with another vector.
func (v Vec4) Minus(o Vec4) Vec4 {
	return Vec4{v.X - o.X, v.Y - o.Y, v.Z - o.Z, v.W - o.W}
}

// Inverse return the inverse of the vector.
func (v Vec4) Inverse() Vec4 {
	return Vec4{-v.X, -v.Y, -v.Z, -v.W}
}

// Times returns the product with a scalar.
func (v Vec4) Times(s float32) Vec4 {
	return Vec4{v.X * s, v.Y * s, v.Z * s, v.W * s}
}

// Slash returns the division by a scalar (which must be non-zero).
func (v Vec4) Slash(s float32) Vec4 {
	return Vec4{v.X / s, v.Y / s, v.Z / s, v.W / s}
}

//------------------------------------------------------------------------------

// Dot returns the dot product with another vector.
func (v Vec4) Dot(o Vec4) float32 {
	return v.X*o.X + v.Y*o.Y + v.Z*o.Z + v.W*o.W
}

//------------------------------------------------------------------------------

// Length returns the euclidian length of the vector.
func (v Vec4) Length() float32 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y + v.Z*v.Z + v.W*v.W)
}

// Normalized returns the normalization of the vector (i.e. the vector divided
// by its length).
//
// Important: Length() must be non-zero.
func (v Vec4) Normalized() Vec4 {
	length := math.Sqrt(v.X*v.X + v.Y*v.Y + v.Z*v.Z + v.W*v.W)
	return Vec4{v.X / length, v.Y / length, v.Z / length, v.W / length}
}

//------------------------------------------------------------------------------

// IsAlmostEqual returns true if the difference between the two vectors is less
// than the specified ULPs (Unit in the Last Place).
//
// Handle special cases: zero, infinites, denormals.
//
// See also IsNearlyEqual and IsRoughlyEqual.
func (v Vec4) IsAlmostEqual(o Vec4, ulps uint32) bool {
	return math.IsAlmostEqual(v.X, o.X, ulps) &&
		math.IsAlmostEqual(v.Y, o.Y, ulps) &&
		math.IsAlmostEqual(v.Z, o.Z, ulps) &&
		math.IsAlmostEqual(v.W, o.W, ulps)
}

// IsNearlyEqual Returns true if the relative error between the two vectors is
// less than epsilon.
//
// Handles special cases: zero, infinites, denormals.
//
// See also IsAlmostEqual and IsRoughlyEqual.
func (v Vec4) IsNearlyEqual(o Vec4, epsilon float32) bool {
	return math.IsNearlyEqual(v.X, o.X, epsilon) &&
		math.IsNearlyEqual(v.Y, o.Y, epsilon) &&
		math.IsNearlyEqual(v.Z, o.Z, epsilon) &&
		math.IsNearlyEqual(v.W, o.W, epsilon)
}

// IsRoughlyEqual Returns true if the absolute error between the two vectors is
// less than epsilon.
//
// See also IsNearlyEqual and IsAlmostEqual.
func (v Vec4) IsRoughlyEqual(o Vec4, epsilon float32) bool {
	return math.IsRoughlyEqual(v.X, o.X, epsilon) &&
		math.IsRoughlyEqual(v.Y, o.Y, epsilon) &&
		math.IsRoughlyEqual(v.Z, o.Z, epsilon) &&
		math.IsRoughlyEqual(v.W, o.W, epsilon)
}

//------------------------------------------------------------------------------
