// Copyright (c) 2013-2016 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package geom

//------------------------------------------------------------------------------

// IVec2 is an integer vector with 2 components.
type IVec2 struct {
	X int32
	Y int32
}

//------------------------------------------------------------------------------

// Float converts to a float vector.
func (v IVec2) Float() Vec2 {
	return Vec2{float32(v.X), float32(v.Y)}
}

//------------------------------------------------------------------------------

// Plus returns the sum with another vector.
func (v IVec2) Plus(o IVec2) IVec2 {
	return IVec2{v.X + o.X, v.Y + o.Y}
}

// Minus returns the difference with another vector.
func (v IVec2) Minus(o IVec2) IVec2 {
	return IVec2{v.X - o.X, v.Y - o.Y}
}

// Inverse return the inverse of the vector.
func (v IVec2) Inverse() IVec2 {
	return IVec2{-v.X, -v.Y}
}

// Times returns the product with a scalar.
func (v IVec2) Times(s int32) IVec2 {
	return IVec2{v.X * s, v.Y * s}
}

// Slash returns the division by a scalar (which must be non-zero).
func (v IVec2) Slash(s int32) IVec2 {
	return IVec2{v.X / s, v.Y / s}
}

//------------------------------------------------------------------------------
