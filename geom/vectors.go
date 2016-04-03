// Copyright (c) 2013-2016 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package geom

//------------------------------------------------------------------------------

import (
	math64 "math"

	math32 "github.com/drakmaniso/glam/math"
)

//------------------------------------------------------------------------------

// Vec4 is a 32-bit float vector with 4 components.
type Vec4 struct {
	X float32
	Y float32
	Z float32
	W float32
}

// Vec3 is a 32-bit float vector with 3 components.
type Vec3 struct {
	X float32
	Y float32
	Z float32
}

// Vec2 is a 32-bit float vector with 2 components..
type Vec2 struct {
	X float32
	Y float32
}

// DVec4 is a 64-bit float vector with 4 components.
type DVec4 struct {
	X float64
	Y float64
	Z float64
	W float64
}

// DVec3 is a 64-bit float vector with 3 components.
type DVec3 struct {
	X float64
	Y float64
	Z float64
}

// DVec2 is a 64-bit float vector with 2 components..
type DVec2 struct {
	X float64
	Y float64
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

// Dehomogenized returns the dehomogenization of the vector (i.e. perspective
// divide).
//
// Important: W must be non-zero.
func (v DVec4) Dehomogenized() DVec3 {
	return DVec3{v.X / v.W, v.Y / v.W, v.Z / v.W}
}

// Homogenized returns the homogeneous coordinates of the vector.
func (v DVec3) Homogenized() DVec4 {
	return DVec4{v.X, v.Y, v.Z, 1.0}
}

// Dehomogenized returns the dehomogenization of the vector (i.e. perspective
// divide).
//
// Important: Z must be non-zero.
func (v DVec3) Dehomogenized() DVec2 {
	return DVec2{v.X / v.Z, v.Y / v.Z}
}

// Homogenized returns the homogeneous coordinates of the vector.
func (v DVec2) Homogenized() DVec3 {
	return DVec3{v.X, v.Y, 1.0}
}

//------------------------------------------------------------------------------

// Vec3 returns a 3 components vector made of X, Y and Z.
func (v Vec4) Vec3() Vec3 {
	return Vec3{v.X, v.Y, v.Z}
}

// Vec2 returns a 2 components vector made of X, and Y.
func (v Vec4) Vec2() Vec2 {
	return Vec2{v.X, v.Y}
}

// Vec2 returns a 2 components vector made of X, and Y.
func (v Vec3) Vec2() Vec2 {
	return Vec2{v.X, v.Y}
}

// Vec4 returns a 4 components vector made of X, Y, Z and w.
func (v Vec3) Vec4(w float32) Vec4 {
	return Vec4{v.X, v.Y, v.Z, w}
}

// Vec3 returns a 3 components vector made of X, Y, and z.
func (v Vec2) Vec3(z float32) Vec3 {
	return Vec3{v.X, v.Y, z}
}

// Vec4 returns a 4 components vector made of X, Y, z, and w.
func (v Vec2) Vec4(z, w float32) Vec4 {
	return Vec4{v.X, v.Y, z, w}
}

// DVec3 returns a 3 components vector made of X, Y and Z.
func (v DVec4) DVec3() DVec3 {
	return DVec3{v.X, v.Y, v.Z}
}

// DVec2 returns a 2 components vector made of X, and Y.
func (v DVec4) DVec2() DVec2 {
	return DVec2{v.X, v.Y}
}

// DVec2 returns a 2 components vector made of X, and Y.
func (v DVec3) DVec2() DVec2 {
	return DVec2{v.X, v.Y}
}

// DVec4 returns a 4 components vector made of X, Y, Z and w.
func (v DVec3) DVec4(w float64) DVec4 {
	return DVec4{v.X, v.Y, v.Z, w}
}

// DVec3 returns a 3 components vector made of X, Y, and z.
func (v DVec2) DVec3(z float64) DVec3 {
	return DVec3{v.X, v.Y, z}
}

// DVec4 returns a 4 components vector made of X, Y, z, and w.
func (v DVec2) DVec4(z, w float64) DVec4 {
	return DVec4{v.X, v.Y, z, w}
}

//------------------------------------------------------------------------------

// DVec4 converts the vector to 64-bit.
func (v Vec4) DVec4() DVec4 {
	return DVec4{float64(v.X), float64(v.Y), float64(v.Z), float64(v.W)}
}

// Vec4 converts the vector to 64-bit.
func (v DVec4) Vec4() Vec4 {
	return Vec4{float32(v.X), float32(v.Y), float32(v.Z), float32(v.W)}
}

// DVec3 converts the vector to 64-bit.
func (v Vec3) DVec3() DVec3 {
	return DVec3{float64(v.X), float64(v.Y), float64(v.Z)}
}

// Vec3 converts the vector to 64-bit.
func (v DVec3) Vec3() Vec3 {
	return Vec3{float32(v.X), float32(v.Y), float32(v.Z)}
}

// DVec2 converts the vector to 64-bit.
func (v Vec2) DVec2() DVec2 {
	return DVec2{float64(v.X), float64(v.Y)}
}

// Vec2 converts the vector to 64-bit.
func (v DVec2) Vec2() Vec2 {
	return Vec2{float32(v.X), float32(v.Y)}
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

// Plus returns the sum with another vector.
func (v DVec4) Plus(o DVec4) DVec4 {
	return DVec4{v.X + o.X, v.Y + o.Y, v.Z + o.Z, v.W + o.W}
}

// Plus returns the sum with another vector.
func (v DVec3) Plus(b DVec3) DVec3 {
	return DVec3{v.X + b.X, v.Y + b.Y, v.Z + b.Z}
}

// Plus returns the sum with another vector.
func (v DVec2) Plus(b DVec2) DVec2 {
	return DVec2{v.X + b.X, v.Y + b.Y}
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

// Minus returns the difference with another vector.
func (v DVec4) Minus(o DVec4) DVec4 {
	return DVec4{v.X - o.X, v.Y - o.Y, v.Z - o.Z, v.W - o.W}
}

// Minus returns the difference with another vector.
func (v DVec3) Minus(b DVec3) DVec3 {
	return DVec3{v.X - b.X, v.Y - b.Y, v.Z - b.Z}
}

// Minus returns the difference with another vector.
func (v DVec2) Minus(b DVec2) DVec2 {
	return DVec2{v.X - b.X, v.Y - b.Y}
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

// Inverse return the inverse of the vector.
func (v DVec4) Inverse() DVec4 {
	return DVec4{-v.X, -v.Y, -v.Z, -v.W}
}

// Inverse return the inverse of the vector.
func (v DVec3) Inverse() DVec3 {
	return DVec3{-v.X, -v.Y, -v.Z}
}

// Inverse return the inverse of the vector.
func (v DVec2) Inverse() DVec2 {
	return DVec2{-v.X, -v.Y}
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

// Times returns the product with a scalar.
func (v DVec4) Times(s float64) DVec4 {
	return DVec4{v.X * s, v.Y * s, v.Z * s, v.W * s}
}

// Times returns the product with a scalar.
func (v DVec3) Times(s float64) DVec3 {
	return DVec3{v.X * s, v.Y * s, v.Z * s}
}

// Times returns the product with a scalar.
func (v DVec2) Times(s float64) DVec2 {
	return DVec2{v.X * s, v.Y * s}
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

// Slash returns the division by a scalar (which must be non-zero).
func (v DVec4) Slash(s float64) DVec4 {
	return DVec4{v.X / s, v.Y / s, v.Z / s, v.W / s}
}

// Slash returns the division by a scalar.
//
// Important: the scalar must be non-zero.
func (v DVec3) Slash(s float64) DVec3 {
	return DVec3{v.X / s, v.Y / s, v.Z / s}
}

// Slash returns the division by a scalar (which must be non-zero).
func (v DVec2) Slash(s float64) DVec2 {
	return DVec2{v.X / s, v.Y / s}
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

// Dot returns the dot product with another vector.
func (v DVec4) Dot(o DVec4) float64 {
	return v.X*o.X + v.Y*o.Y + v.Z*o.Z + v.W*o.W
}

// Dot returns the dot product with another vector.
func (v DVec3) Dot(b DVec3) float64 {
	return v.X*b.X + v.Y*b.Y + v.Z*b.Z
}

// Dot returns the dot product with another vector.
func (v DVec2) Dot(b DVec2) float64 {
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

// Cross returns the cross product with another vector.
func (v DVec3) Cross(b DVec3) DVec3 {
	return DVec3{
		v.Y*b.Z - v.Z*b.Y,
		v.Z*b.X - v.X*b.Z,
		v.X*b.Y - v.Y*b.X,
	}
}

//------------------------------------------------------------------------------

// Length returns the euclidian length of the vector.
func (v Vec4) Length() float32 {
	return math32.Sqrt(v.X*v.X + v.Y*v.Y + v.Z*v.Z + v.W*v.W)
}

// Length returns the euclidian length.
func (v Vec3) Length() float32 {
	return math32.Sqrt(v.X*v.X + v.Y*v.Y + v.Z*v.Z)
}

// Length returns the euclidian length of the vector.
func (v Vec2) Length() float32 {
	return math32.Sqrt(v.X*v.X + v.Y*v.Y)
}

// Length returns the euclidian length of the vector.
func (v DVec4) Length() float64 {
	return math64.Sqrt(v.X*v.X + v.Y*v.Y + v.Z*v.Z + v.W*v.W)
}

// Length returns the euclidian length.
func (v DVec3) Length() float64 {
	return math64.Sqrt(v.X*v.X + v.Y*v.Y + v.Z*v.Z)
}

// Length returns the euclidian length of the vector.
func (v DVec2) Length() float64 {
	return math64.Sqrt(v.X*v.X + v.Y*v.Y)
}

//------------------------------------------------------------------------------

// Normalized returns the normalization of the vector (i.e. the vector divided
// by its length).
//
// Important: Length() must be non-zero.
func (v Vec4) Normalized() Vec4 {
	length := math32.Sqrt(v.X*v.X + v.Y*v.Y + v.Z*v.Z + v.W*v.W)
	return Vec4{v.X / length, v.Y / length, v.Z / length, v.W / length}
}

// Normalized return the normalization of the vector (i.e. the vector divided
// by its length).
//
// Important: Length() must be non-zero.
func (v Vec3) Normalized() Vec3 {
	length := math32.Sqrt(v.X*v.X + v.Y*v.Y + v.Z*v.Z)
	return Vec3{v.X / length, v.Y / length, v.Z / length}
}

// Normalized return the normalization of the vector (i.e. the vector divided
// by its length).
//
// Important: Length() must be non-zero.
func (v Vec2) Normalized() Vec2 {
	length := math32.Sqrt(v.X*v.X + v.Y*v.Y)
	return Vec2{v.X / length, v.Y / length}
}

// Normalized returns the normalization of the vector (i.e. the vector divided
// by its length).
//
// Important: Length() must be non-zero.
func (v DVec4) Normalized() DVec4 {
	length := math64.Sqrt(v.X*v.X + v.Y*v.Y + v.Z*v.Z + v.W*v.W)
	return DVec4{v.X / length, v.Y / length, v.Z / length, v.W / length}
}

// Normalized return the normalization of the vector (i.e. the vector divided
// by its length).
//
// Important: Length() must be non-zero.
func (v DVec3) Normalized() DVec3 {
	length := math64.Sqrt(v.X*v.X + v.Y*v.Y + v.Z*v.Z)
	return DVec3{v.X / length, v.Y / length, v.Z / length}
}

// Normalized return the normalization of the vector (i.e. the vector divided
// by its length).
//
// Important: Length() must be non-zero.
func (v DVec2) Normalized() DVec2 {
	length := math64.Sqrt(v.X*v.X + v.Y*v.Y)
	return DVec2{v.X / length, v.Y / length}
}

//------------------------------------------------------------------------------

// IsAlmostEqual returns true if the difference between the two vectors is less
// than the specified ULPs (Unit in the Last Place).
//
// Handle special cases: zero, infinites, denormals.
//
// See also IsNearlyEqual and IsRoughlyEqual.
func (v Vec4) IsAlmostEqual(o Vec4, ulps uint32) bool {
	return math32.IsAlmostEqual(v.X, o.X, ulps) &&
		math32.IsAlmostEqual(v.Y, o.Y, ulps) &&
		math32.IsAlmostEqual(v.Z, o.Z, ulps) &&
		math32.IsAlmostEqual(v.W, o.W, ulps)
}

// IsAlmostEqual returns true if the difference between the two vectors is less
// than the specified ULPs (Unit in the Last Place).
//
// Handle special cases: zero, infinites, denormals.
//
// See also IsNearlyEqual and IsRoughlyEqual.
func (v Vec3) IsAlmostEqual(o Vec3, ulps uint32) bool {
	return math32.IsAlmostEqual(v.X, o.X, ulps) &&
		math32.IsAlmostEqual(v.Y, o.Y, ulps) &&
		math32.IsAlmostEqual(v.Z, o.Z, ulps)
}

// IsAlmostEqual returns true if the difference between the two vectors is less
// than the specified ULPs (Unit in the Last Place).
//
// Handle special cases: zero, infinites, denormals.
//
// See also IsNearlyEqual and IsRoughlyEqual.
func (v Vec2) IsAlmostEqual(o Vec2, ulps uint32) bool {
	return math32.IsAlmostEqual(v.X, o.X, ulps) &&
		math32.IsAlmostEqual(v.Y, o.Y, ulps)
}

//------------------------------------------------------------------------------

// IsNearlyEqual Returns true if the relative error between the two vectors is
// less than epsilon.
//
// Handles special cases: zero, infinites, denormals.
//
// See also IsAlmostEqual and IsRoughlyEqual.
func (v Vec4) IsNearlyEqual(o Vec4, epsilon float32) bool {
	return math32.IsNearlyEqual(v.X, o.X, epsilon) &&
		math32.IsNearlyEqual(v.Y, o.Y, epsilon) &&
		math32.IsNearlyEqual(v.Z, o.Z, epsilon) &&
		math32.IsNearlyEqual(v.W, o.W, epsilon)
}

// IsNearlyEqual Returns true if the relative error between the two vectors is
// less than epsilon.
//
// Handles special cases: zero, infinites, denormals.
//
// See also IsAlmostEqual and IsRoughlyEqual.
func (v Vec3) IsNearlyEqual(o Vec3, epsilon float32) bool {
	return math32.IsNearlyEqual(v.X, o.X, epsilon) &&
		math32.IsNearlyEqual(v.Y, o.Y, epsilon) &&
		math32.IsNearlyEqual(v.Z, o.Z, epsilon)
}

// IsNearlyEqual Returns true if the relative error between the two vectors is
// less than epsilon.
//
// Handles special cases: zero, infinites, denormals.
//
// See also IsAlmostEqual and IsRoughlyEqual.
func (v Vec2) IsNearlyEqual(o Vec2, epsilon float32) bool {
	return math32.IsNearlyEqual(v.X, o.X, epsilon) &&
		math32.IsNearlyEqual(v.Y, o.Y, epsilon)
}

//------------------------------------------------------------------------------

// IsRoughlyEqual Returns true if the absolute error between the two vectors is
// less than epsilon.
//
// See also IsNearlyEqual and IsAlmostEqual.
func (v Vec4) IsRoughlyEqual(o Vec4, epsilon float32) bool {
	return math32.IsRoughlyEqual(v.X, o.X, epsilon) &&
		math32.IsRoughlyEqual(v.Y, o.Y, epsilon) &&
		math32.IsRoughlyEqual(v.Z, o.Z, epsilon) &&
		math32.IsRoughlyEqual(v.W, o.W, epsilon)
}

// IsRoughlyEqual Returns true if the absolute error between the two vectors is
// less than epsilon.
//
// See also IsNearlyEqual and IsAlmostEqual.
func (v Vec3) IsRoughlyEqual(o Vec3, epsilon float32) bool {
	return math32.IsRoughlyEqual(v.X, o.X, epsilon) &&
		math32.IsRoughlyEqual(v.Y, o.Y, epsilon) &&
		math32.IsRoughlyEqual(v.Z, o.Z, epsilon)
}

// IsRoughlyEqual Returns true if the absolute error between the two vectors is
// less than epsilon.
//
// See also IsNearlyEqual and IsAlmostEqual.
func (v Vec2) IsRoughlyEqual(o Vec2, epsilon float32) bool {
	return math32.IsRoughlyEqual(v.X, o.X, epsilon) &&
		math32.IsRoughlyEqual(v.Y, o.Y, epsilon)
}

//------------------------------------------------------------------------------

//TODO: Implement Is..Equal for 64-bit.

//------------------------------------------------------------------------------
