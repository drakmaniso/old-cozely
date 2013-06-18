// Copyright (c) 2013 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package glm

import "github.com/drakmaniso/glam/math"

//------------------------------------------------------------------------------

// Vec2 is a 2D vector of single-precision floats.
type Vec2 struct {
	X float32
	Y float32
}

//------------------------------------------------------------------------------

// Homogenized returns the homogeneous coordinates of a.
func (a Vec2) Homogenized() Vec3 {
	return Vec3{a.X, a.Y, 1.0}
}

// HomogenizedAsDirection returns the homogeneous coordinates
// of a point at infinity in the direction of a.
func (a Vec2) HomogenizedAsDirection() Vec3 {
	return Vec3{a.X, a.Y, 0.0}
}

//------------------------------------------------------------------------------

// Plus returns the sum a + b.
// See also Add.
func (a Vec2) Plus(b Vec2) Vec2 {
	return Vec2{a.X + b.X, a.Y + b.Y}
}

// Add sets a to the sum a + b.
// More efficient than Plus.
func (a *Vec2) Add(b Vec2) {
	a.X += b.X
	a.Y += b.Y
}

//------------------------------------------------------------------------------

// Minus returns the difference a - b.
// See also Subtract.
func (a Vec2) Minus(b Vec2) Vec2 {
	return Vec2{a.X - b.X, a.Y - b.Y}
}

// Subtract sets a to the difference a - b.
// More efficient than Minus.
func (a *Vec2) Subtract(b Vec2) {
	a.X -= b.X
	a.Y -= b.Y
}

//------------------------------------------------------------------------------

// Inverse return the inverse of a.
// See also Invert.
func (a Vec2) Inverse() Vec2 {
	return Vec2{-a.X, -a.Y}
}

// Invert sets a to its inverse.
// More efficient than Inverse.
func (a *Vec2) Invert() {
	a.X = -a.X
	a.Y = -a.Y
}

//------------------------------------------------------------------------------

// Times returns the product of a with the scalar s.
// See also Multiply.
func (a Vec2) Times(s float32) Vec2 {
	return Vec2{a.X * s, a.Y * s}
}

// Multiply sets a to the product of a with the scalar s.
// More efficient than Times.
func (a *Vec2) Multiply(s float32) {
	a.X *= s
	a.Y *= s
}

//------------------------------------------------------------------------------

// Slash returns the division of a by the scalar s.
// s must be non-zero.
// See also Divide.
func (a Vec2) Slash(s float32) Vec2 {
	return Vec2{a.X / s, a.Y / s}
}

// Divide sets a to the division of a by the scalar s.
// s must be non-zero.
// More efficient than Slash.
func (a *Vec2) Divide(s float32) {
	a.X /= s
	a.Y /= s
}

//------------------------------------------------------------------------------

// Dot returns the dot product of a and b.
func (a Vec2) Dot(b Vec2) float32 {
	return a.X*b.X + a.Y*b.Y
}

//------------------------------------------------------------------------------

// Returns |a| (the euclidian length of a).
func (a Vec2) Length() float32 {
	return math.Sqrt(a.X*a.X + a.Y*a.Y)
}

// Normalized return a/|a| (i.e. the normalization of a).
// a must be non-zero.
// See also Normalize.
func (a Vec2) Normalized() Vec2 {
	length := math.Sqrt(a.X*a.X + a.Y*a.Y)
	return Vec2{a.X / length, a.Y / length}
}

// Normalize sets a to a/|a| (i.e. normalizes a).
// a must be non-zero.
// More efficitent than Normalized.
func (a *Vec2) Normalize() {
	length := math.Sqrt(a.X*a.X + a.Y*a.Y)
	a.X /= length
	a.Y /= length
}

//------------------------------------------------------------------------------
