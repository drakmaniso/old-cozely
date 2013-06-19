// Copyright (c) 2013 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package glam

import "github.com/drakmaniso/glam/math"

//------------------------------------------------------------------------------

// Vec3 is a single-precision vector with 3 components.
type Vec3 struct {
	X float32
	Y float32
	Z float32
}

//------------------------------------------------------------------------------

// Homogenized returns the homogeneous coordinates of a.
func (a Vec3) Homogenized() Vec4 {
	return Vec4{a.X, a.Y, a.Z, 1.0}
}

// HomogenizedAsDirection returns the homogeneous coordinates
// of a point at infinity in the direction of a.
func (a Vec3) HomogenizedAsDirection() Vec4 {
	return Vec4{a.X, a.Y, a.Z, 0.0}
}

// Returns the dehomogenization of a (perspective divide).
// a.Z must be non-zero.
func (a Vec3) Dehomogenized() Vec2 {
	return Vec2{a.X / a.Z, a.Y / a.Z}
}

//------------------------------------------------------------------------------

// Plus returns the sum a + b.
// See also Add.
func (a Vec3) Plus(b Vec3) Vec3 {
	return Vec3{a.X + b.X, a.Y + b.Y, a.Z + b.Z}
}

// Add sets a to the sum a + b.
// More efficient than Plus.
func (a *Vec3) Add(b Vec3) {
	a.X += b.X
	a.Y += b.Y
	a.Z += b.Z
}

//------------------------------------------------------------------------------

// Minus returns the difference a - b.
// See also Subtract.
func (a Vec3) Minus(b Vec3) Vec3 {
	return Vec3{a.X - b.X, a.Y - b.Y, a.Z - b.Z}
}

// Subtract sets a to the difference a - b.
// More efficient than Minus.
func (a *Vec3) Subtract(b Vec3) {
	a.X -= b.X
	a.Y -= b.Y
	a.Z -= b.Z
}

//------------------------------------------------------------------------------

// Inverse return the inverse of a.
// See also Invert.
func (a Vec3) Inverse() Vec3 {
	return Vec3{-a.X, -a.Y, -a.Z}
}

// Invert sets a to its inverse.
// More efficient than Inverse.
func (a *Vec3) Invert() {
	a.X = -a.X
	a.Y = -a.Y
	a.Z = -a.Z
}

//------------------------------------------------------------------------------

// Times returns the product of a with the scalar s.
// See also Multiply.
func (a Vec3) Times(s float32) Vec3 {
	return Vec3{a.X * s, a.Y * s, a.Z * s}
}

// Multiply sets a to the product of a with the scalar s.
// More efficient than Times.
func (a *Vec3) Multiply(s float32) {
	a.X *= s
	a.Y *= s
	a.Z *= s
}

//------------------------------------------------------------------------------

// Slash returns the division of a by the scalar s.
// s must be non-zero.
// See also Divide.
func (a Vec3) Slash(s float32) Vec3 {
	return Vec3{a.X / s, a.Y / s, a.Z / s}
}

// Divide sets a to the division of a by the scalar s.
// s must be non-zero.
// More efficient than Slash.
func (a *Vec3) Divide(s float32) {
	a.X /= s
	a.Y /= s
	a.Z /= s
}

//------------------------------------------------------------------------------

// Cross returns the cross product of a and b.
func (a Vec3) Cross(b Vec3) Vec3 {
	return Vec3{
		a.Y*b.Z - a.Z*b.Y,
		a.Z*b.X - a.X*b.Z,
		a.X*b.Y - a.Y*b.X,
	}
}

//------------------------------------------------------------------------------

// Dot returns the dot product of a and b.
func (a Vec3) Dot(b Vec3) float32 {
	return a.X*b.X + a.Y*b.Y + a.Z*b.Z
}

//------------------------------------------------------------------------------

// Returns |a| (the euclidian length of a).
func (a Vec3) Length() float32 {
	return math.Sqrt(a.X*a.X + a.Y*a.Y + a.Z*a.Z)
}

// Normalized return a/|a| (i.e. the normalization of a).
// a must be non-zero.
// See also Normalize.
func (a Vec3) Normalized() Vec3 {
	length := math.Sqrt(a.X*a.X + a.Y*a.Y + a.Z*a.Z)
	return Vec3{a.X / length, a.Y / length, a.Z / length}
}

// Normalize sets a to a/|a| (i.e. normalizes a).
// a must be non-zero.
// More efficitent than Normalized.
func (a *Vec3) Normalize() {
	length := math.Sqrt(a.X*a.X + a.Y*a.Y + a.Z*a.Z)
	a.X /= length
	a.Y /= length
	a.Z /= length
}

//------------------------------------------------------------------------------
