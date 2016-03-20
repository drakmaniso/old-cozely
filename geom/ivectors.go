// Copyright (c) 2013-2016 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package geom

//------------------------------------------------------------------------------

// IVec4 is an integer vector with 4 components.
type IVec4 struct {
	X int32
	Y int32
	Z int32
	W int32
}

// IVec3 is an integer vector with 3 components.
type IVec3 struct {
	X int32
	Y int32
	Z int32
}

// IVec2 is an integer vector with 2 components.
type IVec2 struct {
	X int32
	Y int32
}

//------------------------------------------------------------------------------

// Float converts to a float vector.
func (v IVec4) Float() Vec4 {
	return Vec4{float32(v.X), float32(v.Y), float32(v.Z), float32(v.W)}
}

// Float converts to a float vector.
func (v IVec3) Float() Vec3 {
	return Vec3{float32(v.X), float32(v.Y), float32(v.Z)}
}

// Float converts to a float vector.
func (v IVec2) Float() Vec2 {
	return Vec2{float32(v.X), float32(v.Y)}
}

//------------------------------------------------------------------------------

// XYZ returns a 3D vector made of X, Y and Z.
func (v IVec4) XYZ() IVec3 {
	return IVec3{v.X, v.Y, v.Z}
}

// XY returns a 2D vector made of X, and Y.
func (v IVec4) XY() IVec2 {
	return IVec2{v.X, v.Y}
}

// XY returns a 2D vector made of X, and Y.
func (v IVec3) XY() IVec2 {
	return IVec2{v.X, v.Y}
}

//------------------------------------------------------------------------------

// Plus returns the sum with another vector.
func (v IVec4) Plus(o IVec4) IVec4 {
	return IVec4{v.X + o.X, v.Y + o.Y, v.Z + o.Z, v.W + o.W}
}

// Plus returns the sum with another vector.
func (v IVec3) Plus(o IVec3) IVec3 {
	return IVec3{v.X + o.X, v.Y + o.Y, v.Z + o.Z}
}

// Plus returns the sum with another vector.
func (v IVec2) Plus(o IVec2) IVec2 {
	return IVec2{v.X + o.X, v.Y + o.Y}
}

//------------------------------------------------------------------------------

// Minus returns the difference with another vector.
func (v IVec4) Minus(o IVec4) IVec4 {
	return IVec4{v.X - o.X, v.Y - o.Y, v.Z - o.Z, v.W - o.W}
}

// Minus returns the difference with another vector.
func (v IVec3) Minus(o IVec3) IVec3 {
	return IVec3{v.X - o.X, v.Y - o.Y, v.Z - o.Z}
}

// Minus returns the difference with another vector.
func (v IVec2) Minus(o IVec2) IVec2 {
	return IVec2{v.X - o.X, v.Y - o.Y}
}

//------------------------------------------------------------------------------

// Inverse return the inverse of the vector.
func (v IVec4) Inverse() IVec4 {
	return IVec4{-v.X, -v.Y, -v.Z, -v.W}
}

// Inverse return the inverse of the vector.
func (v IVec3) Inverse() IVec3 {
	return IVec3{-v.X, -v.Y, -v.Z}
}

// Inverse return the inverse of the vector.
func (v IVec2) Inverse() IVec2 {
	return IVec2{-v.X, -v.Y}
}

//------------------------------------------------------------------------------

// Times returns the product with a scalar.
func (v IVec4) Times(s int32) IVec4 {
	return IVec4{v.X * s, v.Y * s, v.Z * s, v.W * s}
}

// Times returns the product with a scalar.
func (v IVec3) Times(s int32) IVec3 {
	return IVec3{v.X * s, v.Y * s, v.Z * s}
}

// Times returns the product with a scalar.
func (v IVec2) Times(s int32) IVec2 {
	return IVec2{v.X * s, v.Y * s}
}

//------------------------------------------------------------------------------

// Slash returns the division by a scalar (which must be non-zero).
func (v IVec4) Slash(s int32) IVec4 {
	return IVec4{v.X / s, v.Y / s, v.Z / s, v.W / s}
}

// Slash returns the division by a scalar (which must be non-zero).
func (v IVec3) Slash(s int32) IVec3 {
	return IVec3{v.X / s, v.Y / s, v.Z / s}
}

// Slash returns the division by a scalar (which must be non-zero).
func (v IVec2) Slash(s int32) IVec2 {
	return IVec2{v.X / s, v.Y / s}
}

//------------------------------------------------------------------------------
