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
func (a Vec2) Homogenized() Vec3 {
	return Vec3{a.X, a.Y, 1.0}
}

// HomogenizedAsDirection returns the homogeneous coordinates
// of a point at infinity in the direction of the vector.
func (a Vec2) HomogenizedAsDirection() Vec3 {
	return Vec3{a.X, a.Y, 0.0}
}

//------------------------------------------------------------------------------

// Plus returns the sum with another vector.
func (a Vec2) Plus(b Vec2) Vec2 {
	return Vec2{a.X + b.X, a.Y + b.Y}
}

// Minus returns the difference with another vector.
func (a Vec2) Minus(b Vec2) Vec2 {
	return Vec2{a.X - b.X, a.Y - b.Y}
}

// Inverse return the inverse of the vector.
func (a Vec2) Inverse() Vec2 {
	return Vec2{-a.X, -a.Y}
}

// Times returns the product with a scalar.
func (a Vec2) Times(s float32) Vec2 {
	return Vec2{a.X * s, a.Y * s}
}

// Slash returns the division by a scalar (which must be non-zero).
func (a Vec2) Slash(s float32) Vec2 {
	return Vec2{a.X / s, a.Y / s}
}

//------------------------------------------------------------------------------

// Dot returns the dot product with another vector.
func (a Vec2) Dot(b Vec2) float32 {
	return a.X*b.X + a.Y*b.Y
}

//------------------------------------------------------------------------------

// Length returns the euclidian length of the vector.
func (a Vec2) Length() float32 {
	return math.Sqrt(a.X*a.X + a.Y*a.Y)
}

// Normalized return the normalization of the vector (i.e. the vector divided
// by its length).
//
// Important: Length() must be non-zero.
func (a Vec2) Normalized() Vec2 {
	length := math.Sqrt(a.X*a.X + a.Y*a.Y)
	return Vec2{a.X / length, a.Y / length}
}

//------------------------------------------------------------------------------
