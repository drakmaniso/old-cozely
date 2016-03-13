// Copyright (c) 2013-2016 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package geom

//------------------------------------------------------------------------------

// IVec3 is an integer vector with 3 components.
type IVec3 struct {
	X int32
	Y int32
	Z int32
}

//------------------------------------------------------------------------------

// Float converts to a float vector.
func (v IVec3) Float() Vec3 {
	return Vec3{float32(v.X), float32(v.Y), float32(v.Z)}
}

// XY returns a 2D vector made of X, and Y.
func (v IVec3) XY() IVec2 {
	return IVec2{v.X, v.Y}
}

//------------------------------------------------------------------------------

// Plus returns the sum with another vector.
func (v IVec3) Plus(o IVec3) IVec3 {
	return IVec3{v.X + o.X, v.Y + o.Y, v.Z + o.Z}
}

// Minus returns the difference with another vector.
func (v IVec3) Minus(o IVec3) IVec3 {
	return IVec3{v.X - o.X, v.Y - o.Y, v.Z - o.Z}
}

// Inverse return the inverse of the vector.
func (v IVec3) Inverse() IVec3 {
	return IVec3{-v.X, -v.Y, -v.Z}
}

// Times returns the product with a scalar.
func (v IVec3) Times(s int32) IVec3 {
	return IVec3{v.X * s, v.Y * s, v.Z * s}
}

// Slash returns the division by a scalar (which must be non-zero).
func (v IVec3) Slash(s int32) IVec3 {
	return IVec3{v.X / s, v.Y / s, v.Z / s}
}

//------------------------------------------------------------------------------
