// Copyright (c) 2013-2016 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package geom

//------------------------------------------------------------------------------

// IVec4 is a 32-bit integer vector with 4 components.
type IVec4 struct {
	X int32
	Y int32
	Z int32
	W int32
}

// IVec3 is a 32-bit integer vector with 3 components.
type IVec3 struct {
	X int32
	Y int32
	Z int32
}

// IVec2 is a 32-bit integer vector with 2 components.
type IVec2 struct {
	X int32
	Y int32
}

// UVec4 is an unsigned 32-bit integer vector with 4 components.
type UVec4 struct {
	X uint32
	Y uint32
	Z uint32
	W uint32
}

// UVec3 is an unsigned 32-bit integer vector with 3 components.
type UVec3 struct {
	X uint32
	Y uint32
	Z uint32
}

// UVec2 is an unsigned 32-bit integer vector with 2 components.
type UVec2 struct {
	X uint32
	Y uint32
}

// I8Vec4 is a 8-bit integer vector with 4 components.
type I8Vec4 struct {
	X int8
	Y int8
	Z int8
	W int8
}

// I8Vec3 is a 8-bit integer vector with 3 components.
type I8Vec3 struct {
	X int8
	Y int8
	Z int8
}

// I8Vec2 is a 8-bit integer vector with 2 components.
type I8Vec2 struct {
	X int8
	Y int8
}

// U8Vec4 is an unsigned 8-bit integer vector with 4 components.
type U8Vec4 struct {
	X uint8
	Y uint8
	Z uint8
	W uint8
}

// U8Vec3 is an unsigned 8-bit integer vector with 3 components.
type U8Vec3 struct {
	X uint8
	Y uint8
	Z uint8
}

// U8Vec2 is an unsigned 8-bit integer vector with 2 components.
type U8Vec2 struct {
	X uint8
	Y uint8
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

// Float converts to a float vector.
func (v UVec4) Float() Vec4 {
	return Vec4{float32(v.X), float32(v.Y), float32(v.Z), float32(v.W)}
}

// Float converts to a float vector.
func (v UVec3) Float() Vec3 {
	return Vec3{float32(v.X), float32(v.Y), float32(v.Z)}
}

// Float converts to a float vector.
func (v UVec2) Float() Vec2 {
	return Vec2{float32(v.X), float32(v.Y)}
}

// Float converts to a float vector.
func (v I8Vec4) Float() Vec4 {
	return Vec4{float32(v.X), float32(v.Y), float32(v.Z), float32(v.W)}
}

// Float converts to a float vector.
func (v I8Vec3) Float() Vec3 {
	return Vec3{float32(v.X), float32(v.Y), float32(v.Z)}
}

// Float converts to a float vector.
func (v I8Vec2) Float() Vec2 {
	return Vec2{float32(v.X), float32(v.Y)}
}

// Float converts to a float vector.
func (v U8Vec4) Float() Vec4 {
	return Vec4{float32(v.X), float32(v.Y), float32(v.Z), float32(v.W)}
}

// Float converts to a float vector.
func (v U8Vec3) Float() Vec3 {
	return Vec3{float32(v.X), float32(v.Y), float32(v.Z)}
}

// Float converts to a float vector.
func (v U8Vec2) Float() Vec2 {
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

// XYZw returns a 4D vector made of X, Y, Z and w.
func (v IVec3) XYZw(w int32) IVec4 {
	return IVec4{v.X, v.Y, v.Z, w}
}

// XYz returns a 3D vector made of X, Y, and z.
func (v IVec2) XYz(z int32) IVec3 {
	return IVec3{v.X, v.Y, z}
}

// XYzw returns a 4D vector made of X, Y, z, and w.
func (v IVec2) XYzw(z, w int32) IVec4 {
	return IVec4{v.X, v.Y, z, w}
}

// XYZ returns a 3D vector made of X, Y and Z.
func (v UVec4) XYZ() UVec3 {
	return UVec3{v.X, v.Y, v.Z}
}

// XY returns a 2D vector made of X, and Y.
func (v UVec4) XY() UVec2 {
	return UVec2{v.X, v.Y}
}

// XY returns a 2D vector made of X, and Y.
func (v UVec3) XY() UVec2 {
	return UVec2{v.X, v.Y}
}

// XYZw returns a 4D vector made of X, Y, Z and w.
func (v UVec3) XYZw(w uint32) UVec4 {
	return UVec4{v.X, v.Y, v.Z, w}
}

// XYz returns a 3D vector made of X, Y, and z.
func (v UVec2) XYz(z uint32) UVec3 {
	return UVec3{v.X, v.Y, z}
}

// XYzw returns a 4D vector made of X, Y, z, and w.
func (v UVec2) XYzw(z, w uint32) UVec4 {
	return UVec4{v.X, v.Y, z, w}
}

// XYZ returns a 3D vector made of X, Y and Z.
func (v I8Vec4) XYZ() I8Vec3 {
	return I8Vec3{v.X, v.Y, v.Z}
}

// XY returns a 2D vector made of X, and Y.
func (v I8Vec4) XY() I8Vec2 {
	return I8Vec2{v.X, v.Y}
}

// XY returns a 2D vector made of X, and Y.
func (v I8Vec3) XY() I8Vec2 {
	return I8Vec2{v.X, v.Y}
}

// XYZw returns a 4D vector made of X, Y, Z and w.
func (v I8Vec3) XYZw(w int8) I8Vec4 {
	return I8Vec4{v.X, v.Y, v.Z, w}
}

// XYz returns a 3D vector made of X, Y, and z.
func (v I8Vec2) XYz(z int8) I8Vec3 {
	return I8Vec3{v.X, v.Y, z}
}

// XYzw returns a 4D vector made of X, Y, z, and w.
func (v I8Vec2) XYzw(z, w int8) I8Vec4 {
	return I8Vec4{v.X, v.Y, z, w}
}

// XYZ returns a 3D vector made of X, Y and Z.
func (v U8Vec4) XYZ() U8Vec3 {
	return U8Vec3{v.X, v.Y, v.Z}
}

// XY returns a 2D vector made of X, and Y.
func (v U8Vec4) XY() U8Vec2 {
	return U8Vec2{v.X, v.Y}
}

// XY returns a 2D vector made of X, and Y.
func (v U8Vec3) XY() U8Vec2 {
	return U8Vec2{v.X, v.Y}
}

// XYZw returns a 4D vector made of X, Y, Z and w.
func (v U8Vec3) XYZw(w uint8) U8Vec4 {
	return U8Vec4{v.X, v.Y, v.Z, w}
}

// XYz returns a 3D vector made of X, Y, and z.
func (v U8Vec2) XYz(z uint8) U8Vec3 {
	return U8Vec3{v.X, v.Y, z}
}

// XYzw returns a 4D vector made of X, Y, z, and w.
func (v U8Vec2) XYzw(z, w uint8) U8Vec4 {
	return U8Vec4{v.X, v.Y, z, w}
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

// Plus returns the sum with another vector.
func (v UVec4) Plus(o UVec4) UVec4 {
	return UVec4{v.X + o.X, v.Y + o.Y, v.Z + o.Z, v.W + o.W}
}

// Plus returns the sum with another vector.
func (v UVec3) Plus(o UVec3) UVec3 {
	return UVec3{v.X + o.X, v.Y + o.Y, v.Z + o.Z}
}

// Plus returns the sum with another vector.
func (v UVec2) Plus(o UVec2) UVec2 {
	return UVec2{v.X + o.X, v.Y + o.Y}
}

// Plus returns the sum with another vector.
func (v I8Vec4) Plus(o I8Vec4) I8Vec4 {
	return I8Vec4{v.X + o.X, v.Y + o.Y, v.Z + o.Z, v.W + o.W}
}

// Plus returns the sum with another vector.
func (v I8Vec3) Plus(o I8Vec3) I8Vec3 {
	return I8Vec3{v.X + o.X, v.Y + o.Y, v.Z + o.Z}
}

// Plus returns the sum with another vector.
func (v I8Vec2) Plus(o I8Vec2) I8Vec2 {
	return I8Vec2{v.X + o.X, v.Y + o.Y}
}

// Plus returns the sum with another vector.
func (v U8Vec4) Plus(o U8Vec4) U8Vec4 {
	return U8Vec4{v.X + o.X, v.Y + o.Y, v.Z + o.Z, v.W + o.W}
}

// Plus returns the sum with another vector.
func (v U8Vec3) Plus(o U8Vec3) U8Vec3 {
	return U8Vec3{v.X + o.X, v.Y + o.Y, v.Z + o.Z}
}

// Plus returns the sum with another vector.
func (v U8Vec2) Plus(o U8Vec2) U8Vec2 {
	return U8Vec2{v.X + o.X, v.Y + o.Y}
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

// Minus returns the difference with another vector.
func (v UVec4) Minus(o UVec4) UVec4 {
	return UVec4{v.X - o.X, v.Y - o.Y, v.Z - o.Z, v.W - o.W}
}

// Minus returns the difference with another vector.
func (v UVec3) Minus(o UVec3) UVec3 {
	return UVec3{v.X - o.X, v.Y - o.Y, v.Z - o.Z}
}

// Minus returns the difference with another vector.
func (v UVec2) Minus(o UVec2) UVec2 {
	return UVec2{v.X - o.X, v.Y - o.Y}
}

// Minus returns the difference with another vector.
func (v I8Vec4) Minus(o I8Vec4) I8Vec4 {
	return I8Vec4{v.X - o.X, v.Y - o.Y, v.Z - o.Z, v.W - o.W}
}

// Minus returns the difference with another vector.
func (v I8Vec3) Minus(o I8Vec3) I8Vec3 {
	return I8Vec3{v.X - o.X, v.Y - o.Y, v.Z - o.Z}
}

// Minus returns the difference with another vector.
func (v I8Vec2) Minus(o I8Vec2) I8Vec2 {
	return I8Vec2{v.X - o.X, v.Y - o.Y}
}

// Minus returns the difference with another vector.
func (v U8Vec4) Minus(o U8Vec4) U8Vec4 {
	return U8Vec4{v.X - o.X, v.Y - o.Y, v.Z - o.Z, v.W - o.W}
}

// Minus returns the difference with another vector.
func (v U8Vec3) Minus(o U8Vec3) U8Vec3 {
	return U8Vec3{v.X - o.X, v.Y - o.Y, v.Z - o.Z}
}

// Minus returns the difference with another vector.
func (v U8Vec2) Minus(o U8Vec2) U8Vec2 {
	return U8Vec2{v.X - o.X, v.Y - o.Y}
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

// Inverse return the inverse of the vector.
func (v UVec4) Inverse() UVec4 {
	return UVec4{-v.X, -v.Y, -v.Z, -v.W}
}

// Inverse return the inverse of the vector.
func (v UVec3) Inverse() UVec3 {
	return UVec3{-v.X, -v.Y, -v.Z}
}

// Inverse return the inverse of the vector.
func (v UVec2) Inverse() UVec2 {
	return UVec2{-v.X, -v.Y}
}

// Inverse return the inverse of the vector.
func (v I8Vec4) Inverse() I8Vec4 {
	return I8Vec4{-v.X, -v.Y, -v.Z, -v.W}
}

// Inverse return the inverse of the vector.
func (v I8Vec3) Inverse() I8Vec3 {
	return I8Vec3{-v.X, -v.Y, -v.Z}
}

// Inverse return the inverse of the vector.
func (v I8Vec2) Inverse() I8Vec2 {
	return I8Vec2{-v.X, -v.Y}
}

// Inverse return the inverse of the vector.
func (v U8Vec4) Inverse() U8Vec4 {
	return U8Vec4{-v.X, -v.Y, -v.Z, -v.W}
}

// Inverse return the inverse of the vector.
func (v U8Vec3) Inverse() U8Vec3 {
	return U8Vec3{-v.X, -v.Y, -v.Z}
}

// Inverse return the inverse of the vector.
func (v U8Vec2) Inverse() U8Vec2 {
	return U8Vec2{-v.X, -v.Y}
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

// Times returns the product with a scalar.
func (v UVec4) Times(s uint32) UVec4 {
	return UVec4{v.X * s, v.Y * s, v.Z * s, v.W * s}
}

// Times returns the product with a scalar.
func (v UVec3) Times(s uint32) UVec3 {
	return UVec3{v.X * s, v.Y * s, v.Z * s}
}

// Times returns the product with a scalar.
func (v UVec2) Times(s uint32) UVec2 {
	return UVec2{v.X * s, v.Y * s}
}

// Times returns the product with a scalar.
func (v I8Vec4) Times(s int8) I8Vec4 {
	return I8Vec4{v.X * s, v.Y * s, v.Z * s, v.W * s}
}

// Times returns the product with a scalar.
func (v I8Vec3) Times(s int8) I8Vec3 {
	return I8Vec3{v.X * s, v.Y * s, v.Z * s}
}

// Times returns the product with a scalar.
func (v I8Vec2) Times(s int8) I8Vec2 {
	return I8Vec2{v.X * s, v.Y * s}
}

// Times returns the product with a scalar.
func (v U8Vec4) Times(s uint8) U8Vec4 {
	return U8Vec4{v.X * s, v.Y * s, v.Z * s, v.W * s}
}

// Times returns the product with a scalar.
func (v U8Vec3) Times(s uint8) U8Vec3 {
	return U8Vec3{v.X * s, v.Y * s, v.Z * s}
}

// Times returns the product with a scalar.
func (v U8Vec2) Times(s uint8) U8Vec2 {
	return U8Vec2{v.X * s, v.Y * s}
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

// Slash returns the division by a scalar (which must be non-zero).
func (v UVec4) Slash(s uint32) UVec4 {
	return UVec4{v.X / s, v.Y / s, v.Z / s, v.W / s}
}

// Slash returns the division by a scalar (which must be non-zero).
func (v UVec3) Slash(s uint32) UVec3 {
	return UVec3{v.X / s, v.Y / s, v.Z / s}
}

// Slash returns the division by a scalar (which must be non-zero).
func (v UVec2) Slash(s uint32) UVec2 {
	return UVec2{v.X / s, v.Y / s}
}

// Slash returns the division by a scalar (which must be non-zero).
func (v I8Vec4) Slash(s int8) I8Vec4 {
	return I8Vec4{v.X / s, v.Y / s, v.Z / s, v.W / s}
}

// Slash returns the division by a scalar (which must be non-zero).
func (v I8Vec3) Slash(s int8) I8Vec3 {
	return I8Vec3{v.X / s, v.Y / s, v.Z / s}
}

// Slash returns the division by a scalar (which must be non-zero).
func (v I8Vec2) Slash(s int8) I8Vec2 {
	return I8Vec2{v.X / s, v.Y / s}
}

// Slash returns the division by a scalar (which must be non-zero).
func (v U8Vec4) Slash(s uint8) U8Vec4 {
	return U8Vec4{v.X / s, v.Y / s, v.Z / s, v.W / s}
}

// Slash returns the division by a scalar (which must be non-zero).
func (v U8Vec3) Slash(s uint8) U8Vec3 {
	return U8Vec3{v.X / s, v.Y / s, v.Z / s}
}

// Slash returns the division by a scalar (which must be non-zero).
func (v U8Vec2) Slash(s uint8) U8Vec2 {
	return U8Vec2{v.X / s, v.Y / s}
}

//------------------------------------------------------------------------------
