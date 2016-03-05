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
func (a Vec3) Homogenized() Vec4 {
	return Vec4{a.X, a.Y, a.Z, 1.0}
}

// HomogenizedAsDirection returns the homogeneous coordinates
// of a point at infinity in the direction of the vector.
func (a Vec3) HomogenizedAsDirection() Vec4 {
	return Vec4{a.X, a.Y, a.Z, 0.0}
}

// Dehomogenized returns the dehomogenization of the vector (i.e. perspective
// divide).
//
// Important: Z must be non-zero.
func (a Vec3) Dehomogenized() Vec2 {
	return Vec2{a.X / a.Z, a.Y / a.Z}
}

//------------------------------------------------------------------------------

// Plus returns the sum with another vector.
func (a Vec3) Plus(b Vec3) Vec3 {
	return Vec3{a.X + b.X, a.Y + b.Y, a.Z + b.Z}
}

// Minus returns the difference with another vector.
func (a Vec3) Minus(b Vec3) Vec3 {
	return Vec3{a.X - b.X, a.Y - b.Y, a.Z - b.Z}
}

// Inverse return the inverse of the vector.
func (a Vec3) Inverse() Vec3 {
	return Vec3{-a.X, -a.Y, -a.Z}
}

// Times returns the product with a scalar.
func (a Vec3) Times(s float32) Vec3 {
	return Vec3{a.X * s, a.Y * s, a.Z * s}
}

// Slash returns the division by a scalar.
//
// Important: the scalar must be non-zero.
func (a Vec3) Slash(s float32) Vec3 {
	return Vec3{a.X / s, a.Y / s, a.Z / s}
}

//------------------------------------------------------------------------------

// Cross returns the cross product with another vector.
func (a Vec3) Cross(b Vec3) Vec3 {
	return Vec3{
		a.Y*b.Z - a.Z*b.Y,
		a.Z*b.X - a.X*b.Z,
		a.X*b.Y - a.Y*b.X,
	}
}

// Dot returns the dot product with another vector.
func (a Vec3) Dot(b Vec3) float32 {
	return a.X*b.X + a.Y*b.Y + a.Z*b.Z
}

//------------------------------------------------------------------------------

// Length returns the euclidian length.
func (a Vec3) Length() float32 {
	return math.Sqrt(a.X*a.X + a.Y*a.Y + a.Z*a.Z)
}

// Normalized return the normalization of the vector (i.e. the vector divided
// by its length).
//
// Important: Length() must be non-zero.
func (a Vec3) Normalized() Vec3 {
	length := math.Sqrt(a.X*a.X + a.Y*a.Y + a.Z*a.Z)
	return Vec3{a.X / length, a.Y / length, a.Z / length}
}

//------------------------------------------------------------------------------
