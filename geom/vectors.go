// Copyright (c) 2013-2016 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package geom

import "github.com/drakmaniso/glam/math"

//------------------------------------------------------------------------------

// Vec4 is a single-precision vector with 4 components.
type Vec4 struct {
	X float32
	Y float32
	Z float32
	W float32
}

// Vec3 is a single-precision vector with 3 components.
type Vec3 struct {
	X float32
	Y float32
	Z float32
}

// Vec2 is a single-precision vector with 2 components..
type Vec2 struct {
	X float32
	Y float32
}

//------------------------------------------------------------------------------

// Dehomogenized returns the dehomogenization of the vector (i.e. perspective
// divide).
//
// Important: W must be non-zero.
func (v Vec4) Dehomogenized() Vec3 {
	return Vec3{v.X / v.W, v.Y / v.W, v.Z / v.W}
}

// Homogenized returns the homogeneous coordinates of the vector.
func (v Vec3) Homogenized() Vec4 {
	return Vec4{v.X, v.Y, v.Z, 1.0}
}

// Dehomogenized returns the dehomogenization of the vector (i.e. perspective
// divide).
//
// Important: Z must be non-zero.
func (v Vec3) Dehomogenized() Vec2 {
	return Vec2{v.X / v.Z, v.Y / v.Z}
}

// Homogenized returns the homogeneous coordinates of the vector.
func (v Vec2) Homogenized() Vec3 {
	return Vec3{v.X, v.Y, 1.0}
}

//------------------------------------------------------------------------------

// XYZ returns a 3D vector made of X, Y and Z.
func (v Vec4) XYZ() Vec3 {
	return Vec3{v.X, v.Y, v.Z}
}

// XY returns a 2D vector made of X, and Y.
func (v Vec4) XY() Vec2 {
	return Vec2{v.X, v.Y}
}

// XY returns a 2D vector made of X, and Y.
func (v Vec3) XY() Vec2 {
	return Vec2{v.X, v.Y}
}

// XYZw returns a 4D vector made of X, Y, Z and w.
func (v Vec3) XYZw(w float32) Vec4 {
	return Vec4{v.X, v.Y, v.Z, w}
}

// XYz returns a 3D vector made of X, Y, and z.
func (v Vec2) XYz(z float32) Vec3 {
	return Vec3{v.X, v.Y, z}
}

// XYzw returns a 4D vector made of X, Y, z, and w.
func (v Vec2) XYzw(z, w float32) Vec4 {
	return Vec4{v.X, v.Y, z, w}
}

//------------------------------------------------------------------------------

// Plus returns the sum with another vector.
func (v Vec4) Plus(o Vec4) Vec4 {
	return Vec4{v.X + o.X, v.Y + o.Y, v.Z + o.Z, v.W + o.W}
}

// Plus returns the sum with another vector.
func (v Vec3) Plus(b Vec3) Vec3 {
	return Vec3{v.X + b.X, v.Y + b.Y, v.Z + b.Z}
}

// Plus returns the sum with another vector.
func (v Vec2) Plus(b Vec2) Vec2 {
	return Vec2{v.X + b.X, v.Y + b.Y}
}

//------------------------------------------------------------------------------

// Minus returns the difference with another vector.
func (v Vec4) Minus(o Vec4) Vec4 {
	return Vec4{v.X - o.X, v.Y - o.Y, v.Z - o.Z, v.W - o.W}
}

// Minus returns the difference with another vector.
func (v Vec3) Minus(b Vec3) Vec3 {
	return Vec3{v.X - b.X, v.Y - b.Y, v.Z - b.Z}
}

// Minus returns the difference with another vector.
func (v Vec2) Minus(b Vec2) Vec2 {
	return Vec2{v.X - b.X, v.Y - b.Y}
}

//------------------------------------------------------------------------------

// Inverse return the inverse of the vector.
func (v Vec4) Inverse() Vec4 {
	return Vec4{-v.X, -v.Y, -v.Z, -v.W}
}

// Inverse return the inverse of the vector.
func (v Vec3) Inverse() Vec3 {
	return Vec3{-v.X, -v.Y, -v.Z}
}

// Inverse return the inverse of the vector.
func (v Vec2) Inverse() Vec2 {
	return Vec2{-v.X, -v.Y}
}

//------------------------------------------------------------------------------

// Times returns the product with a scalar.
func (v Vec4) Times(s float32) Vec4 {
	return Vec4{v.X * s, v.Y * s, v.Z * s, v.W * s}
}

// Times returns the product with a scalar.
func (v Vec3) Times(s float32) Vec3 {
	return Vec3{v.X * s, v.Y * s, v.Z * s}
}

// Times returns the product with a scalar.
func (v Vec2) Times(s float32) Vec2 {
	return Vec2{v.X * s, v.Y * s}
}

//------------------------------------------------------------------------------

// Slash returns the division by a scalar (which must be non-zero).
func (v Vec4) Slash(s float32) Vec4 {
	return Vec4{v.X / s, v.Y / s, v.Z / s, v.W / s}
}

// Slash returns the division by a scalar.
//
// Important: the scalar must be non-zero.
func (v Vec3) Slash(s float32) Vec3 {
	return Vec3{v.X / s, v.Y / s, v.Z / s}
}

// Slash returns the division by a scalar (which must be non-zero).
func (v Vec2) Slash(s float32) Vec2 {
	return Vec2{v.X / s, v.Y / s}
}

//------------------------------------------------------------------------------

// Dot returns the dot product with another vector.
func (v Vec4) Dot(o Vec4) float32 {
	return v.X*o.X + v.Y*o.Y + v.Z*o.Z + v.W*o.W
}

// Dot returns the dot product with another vector.
func (v Vec3) Dot(b Vec3) float32 {
	return v.X*b.X + v.Y*b.Y + v.Z*b.Z
}

// Dot returns the dot product with another vector.
func (v Vec2) Dot(b Vec2) float32 {
	return v.X*b.X + v.Y*b.Y
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

//------------------------------------------------------------------------------

// Length returns the euclidian length of the vector.
func (v Vec4) Length() float32 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y + v.Z*v.Z + v.W*v.W)
}

// Length returns the euclidian length.
func (v Vec3) Length() float32 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y + v.Z*v.Z)
}

// Length returns the euclidian length of the vector.
func (v Vec2) Length() float32 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

//------------------------------------------------------------------------------

// Normalized returns the normalization of the vector (i.e. the vector divided
// by its length).
//
// Important: Length() must be non-zero.
func (v Vec4) Normalized() Vec4 {
	length := math.Sqrt(v.X*v.X + v.Y*v.Y + v.Z*v.Z + v.W*v.W)
	return Vec4{v.X / length, v.Y / length, v.Z / length, v.W / length}
}

// Normalized return the normalization of the vector (i.e. the vector divided
// by its length).
//
// Important: Length() must be non-zero.
func (v Vec3) Normalized() Vec3 {
	length := math.Sqrt(v.X*v.X + v.Y*v.Y + v.Z*v.Z)
	return Vec3{v.X / length, v.Y / length, v.Z / length}
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
func (v Vec4) IsAlmostEqual(o Vec4, ulps uint32) bool {
	return math.IsAlmostEqual(v.X, o.X, ulps) &&
		math.IsAlmostEqual(v.Y, o.Y, ulps) &&
		math.IsAlmostEqual(v.Z, o.Z, ulps) &&
		math.IsAlmostEqual(v.W, o.W, ulps)
}

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

//------------------------------------------------------------------------------

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

//------------------------------------------------------------------------------

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

// IsRoughlyEqual Returns true if the absolute error between the two vectors is
// less than epsilon.
//
// See also IsNearlyEqual and IsAlmostEqual.
func (v Vec3) IsRoughlyEqual(o Vec3, epsilon float32) bool {
	return math.IsRoughlyEqual(v.X, o.X, epsilon) &&
		math.IsRoughlyEqual(v.Y, o.Y, epsilon) &&
		math.IsRoughlyEqual(v.Z, o.Z, epsilon)
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
