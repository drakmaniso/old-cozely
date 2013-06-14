package glm

//------------------------------------------------------------------------------

// Vec4 is a 4D vector of single-precision floats.
type Vec4 struct {
	X float32
	Y float32
	Z float32
	W float32
}

//------------------------------------------------------------------------------

// Returns the dehomogenization of a (perspective divide).
// a.W must be non-zero.
func (a Vec4) Dehomogenized() Vec3 {
	return Vec3{a.X / a.W, a.Y / a.W, a.Z / a.W}
}

//------------------------------------------------------------------------------

// Plus returns the sum a + b.
// See also Add.
func (a Vec4) Plus(b Vec4) Vec4 {
	return Vec4{a.X + b.X, a.Y + b.Y, a.Z + b.Z, a.W + b.W}
}

// Add sets a to the sum a + b.
// More efficient than Plus.
func (a *Vec4) Add(b Vec4) {
	a.X += b.X
	a.Y += b.Y
	a.Z += b.Z
	a.W += b.W
}

//------------------------------------------------------------------------------

// Minus returns the difference a - b.
// See also Subtract.
func (a Vec4) Minus(b Vec4) Vec4 {
	return Vec4{a.X - b.X, a.Y - b.Y, a.Z - b.Z, a.W - b.W}
}

// Subtract sets a to the difference a - b.
// More efficient than Minus.
func (a *Vec4) Subtract(b Vec4) {
	a.X -= b.X
	a.Y -= b.Y
	a.Z -= b.Z
	a.W -= b.W
}

//------------------------------------------------------------------------------

// Inverse return the inverse of a.
// See also Invert.
func (a Vec4) Inverse() Vec4 {
	return Vec4{-a.X, -a.Y, -a.Z, -a.W}
}

// Invert sets a to its inverse.
// More efficient than Inverse.
func (a *Vec4) Invert() {
	a.X = -a.X
	a.Y = -a.Y
	a.Z = -a.Z
	a.W = -a.W
}

//------------------------------------------------------------------------------

// Times returns the product of a with the scalar s.
// See also MultiplyBy.
func (a Vec4) Times(s float32) Vec4 {
	return Vec4{a.X * s, a.Y * s, a.Z * s, a.W * s}
}

// MultiplyBy sets a to the product of a with the scalar s.
// More efficient than Times.
func (a *Vec4) MultiplyBy(s float32) {
	a.X *= s
	a.Y *= s
	a.Z *= s
	a.W *= s
}

//------------------------------------------------------------------------------

// Slash returns the division of a by the scalar s.
// s must be non-zero.
// See also DivideBy.
func (a Vec4) Slash(s float32) Vec4 {
	return Vec4{a.X / s, a.Y / s, a.Z / s, a.W / s}
}

// DivideBy sets a to the division of a by the scalar s.
// s must be non-zero.
// More efficient than Slash.
func (a *Vec4) DivideBy(s float32) {
	a.X /= s
	a.Y /= s
	a.Z /= s
	a.W /= s
}

//------------------------------------------------------------------------------

// Dot returns the dot product of a and b.
func (a Vec4) Dot(b Vec4) float32 {
	return a.X*b.X + a.Y*b.Y + a.Z*b.Z + a.W*b.W
}

//------------------------------------------------------------------------------

// Returns |a| (the euclidian length of a).
func (a Vec4) Length() float32 {
	return Sqrt(a.X*a.X + a.Y*a.Y + a.Z*a.Z + a.W*a.W)
}

// Normalized return a/|a| (i.e. the normalization of a).
// a must be non-zero.
// See also Normalize.
func (a Vec4) Normalized() Vec4 {
	length := Sqrt(a.X*a.X + a.Y*a.Y + a.Z*a.Z + a.W*a.W)
	return Vec4{a.X / length, a.Y / length, a.Z / length, a.W / length}
}

// Normalize sets a to a/|a| (i.e. normalizes a).
// a must be non-zero.
// More efficitent than Normalized.
func (a *Vec4) Normalize() {
	length := Sqrt(a.X*a.X + a.Y*a.Y + a.Z*a.Z + a.W*a.W)
	a.X /= length
	a.Y /= length
	a.Z /= length
}

//------------------------------------------------------------------------------
// Copyright (c) 2013 - Laurent Moussault <moussault.laurent@gmail.com>
