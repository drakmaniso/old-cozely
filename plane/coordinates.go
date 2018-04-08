// Copyright (c) 2013-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package plane

import (
	"math"

	"github.com/drakmaniso/glam/x/math32"
)

//------------------------------------------------------------------------------

// Vector represents any two-dimensional vector.
type Vector interface {
	// XY returns the cartesian coordinates of the vector.
	XY() (x, y float32)
}

//------------------------------------------------------------------------------

// Coord represents a two-dimensional vector, defined by its cartesian
// coordinates.
type Coord struct {
	X float32
	Y float32
}

// XY returns the cartesian coordinates of the vector.
//
// This function implements the Vector interface.
func (c Coord) XY() (x, y float32) {
	return c.X, c.Y
}

// Pixel return the pixel coordinates of the vector. Note that the sign of Y is
// flipped.
func (c Coord) Pixel() Pixel {
	return Pixel{int16(c.X), int16(-c.Y)}
}

// Cartesian returns the cartesian representation of c.
func Cartesian(v Vector) Coord {
	x, y := v.XY()
	return Coord{x, y}
}

// Coord64 returns the 64-bit float representation of c.
func (c Coord) Coord64() Coord64 {
	return Coord64{float64(c.X), float64(c.Y)}
}

// Homogen returns the homogenous coordinates of the vector, with Z set to 1.
func (c Coord) Homogen() Homogen {
	return Homogen{c.X, c.Y, 1.0}
}

// Plus returns the sum with another vector.
func (c Coord) Plus(o Coord) Coord {
	return Coord{c.X + o.X, c.Y + o.Y}
}

// Pluss returns the component-wise sum with two scalars.
func (c Coord) Pluss(x, y float32) Coord {
	return Coord{c.X + x, c.Y + y}
}

// Minus returns the difference with another vector.
func (c Coord) Minus(o Coord) Coord {
	return Coord{c.X - o.X, c.Y - o.Y}
}

// Minuss returns the component-wise difference with two scalars.
func (c Coord) Minuss(x, y float32) Coord {
	return Coord{c.X - x, c.Y - y}
}

// Opposite returns the opposite of the vector.
func (c Coord) Opposite() Coord {
	return Coord{-c.X, -c.Y}
}

// Times returns the product with a scalar.
func (c Coord) Times(s float32) Coord {
	return Coord{c.X * s, c.Y * s}
}

// Timess returns the component-wise product with two scalars.
func (c Coord) Timess(x, y float32) Coord {
	return Coord{c.X * x, c.Y * y}
}

// Timescw returns the component-wise product with another vector.
func (c Coord) Timescw(o Coord) Coord {
	return Coord{c.X * o.X, c.Y * o.Y}
}

// Slash returns the division by a scalar (which must be non-zero).
func (c Coord) Slash(s float32) Coord {
	return Coord{c.X / s, c.Y / s}
}

// Slashs returns the component-wise division by two scalars (which must be
// non-zero).
func (c Coord) Slashs(x, y float32) Coord {
	return Coord{c.X / x, c.Y / y}
}

// Slashcw returns the component-wise division by another vector (of which both
// X and Y must be non-zero).
func (c Coord) Slashcw(o Coord) Coord {
	return Coord{c.X / o.X, c.Y / o.Y}
}

// Mod returns the remainder (modulus) of the division by a scalar (which must
// be non-zero).
func (c Coord) Mod(s float32) Coord {
	return Coord{math32.Mod(c.X, s), math32.Mod(c.Y, s)}
}

// Mods returns the remainder (modulus) of the division by two scalars (which
// must be non-zero).
func (c Coord) Mods(x, y float32) Coord {
	return Coord{math32.Mod(c.X, x), math32.Mod(c.Y, y)}
}

// Modcw returns the remainders (modulus) of the component-wise division by
// another vector (of which both X and Y must be non-zero).
func (c Coord) Modcw(o Coord) Coord {
	return Coord{math32.Mod(c.X, o.X), math32.Mod(c.Y, o.Y)}
}

// Modf returns the integer part and the fractional part of (each component of)
// the vector.
func (c Coord) Modf() (intg, frac Coord) {
	xintg, xfrac := math32.Modf(c.X)
	yintg, yfrac := math32.Modf(c.Y)
	return Coord{xintg, yintg}, Coord{xfrac, yfrac}
}

// Dot returns the dot product with another vector.
func (c Coord) Dot(o Coord) float32 {
	return c.X*o.X + c.Y*o.Y
}

// PerpDot returns the dot product with the perpendicular of v and another
// vector.
func (c Coord) PerpDot(o Coord) float32 {
	return c.X*o.Y - c.Y*o.X
}

// Length returns the euclidian length of the vector.
func (c Coord) Length() float32 {
	// Double conversion is faster than math32.Sqrt because the Go compiler
	// optimizes it.
	return float32(math.Sqrt(float64(c.X*c.X + c.Y*c.Y)))
}

// Length2 returns the square of the euclidian length of the vector.
func (c Coord) Length2() float32 {
	return c.X*c.X + c.Y*c.Y
}

// Distance returns the distance with another vector.
func (c Coord) Distance(o Coord) float32 {
	d := Coord{c.X - o.X, c.Y - o.Y}
	// Double conversion is faster than math32.Sqrt because the Go compiler
	// optimizes it.
	return float32(math.Sqrt(float64(d.X*d.X + d.Y*d.Y)))
}

// Distance2 returns the square of the distance with another vector.
func (c Coord) Distance2(o Coord) float32 {
	d := Coord{c.X - o.X, c.Y - o.Y}
	return d.X*d.X + d.Y*d.Y
}

// Normalized return the normalization of the vector (i.e. the vector divided
// by its length, which must be non-zero).
func (c Coord) Normalized() Coord {
	// Double conversion is faster than math32.Sqrt because the Go compiler
	// optimizes it.
	l := float32(math.Sqrt(float64(c.X*c.X + c.Y*c.Y)))
	return Coord{c.X / l, c.Y / l}
}

// IsAlmostEqual returns true if the difference between the two vectors is less
// than the specified ULPs (Unit in the Last Place).
//
// Handle special cases: zero, infinites, denormals.
//
// See also IsNearlyEqual and IsRoughlyEqual.
func (c Coord) IsAlmostEqual(o Coord, ulps uint32) bool {
	return math32.IsAlmostEqual(c.X, o.X, ulps) &&
		math32.IsAlmostEqual(c.Y, o.Y, ulps)
}

// IsNearlyEqual Returns true if the relative error between the two vectors is
// less than epsilon.
//
// Handles special cases: zero, infinites, denormals.
//
// See also IsAlmostEqual and IsRoughlyEqual.
func (c Coord) IsNearlyEqual(o Coord, epsilon float32) bool {
	return math32.IsNearlyEqual(c.X, o.X, epsilon) &&
		math32.IsNearlyEqual(c.Y, o.Y, epsilon)
}

// IsRoughlyEqual Returns true if the absolute error between the two vectors is
// less than epsilon.
//
// See also IsNearlyEqual and IsAlmostEqual.
func (c Coord) IsRoughlyEqual(o Coord, epsilon float32) bool {
	return math32.IsRoughlyEqual(c.X, o.X, epsilon) &&
		math32.IsRoughlyEqual(c.Y, o.Y, epsilon)
}

//------------------------------------------------------------------------------

// Homogen represents a two-dimensional vector, defined by its homogeneous
// coordinates.
type Homogen struct {
	X float32
	Y float32
	Z float32
}

// XY returns the cartesian coordinates of the vector (i.e. the perspective
// divide of the homogeneous coordinates). Z must be non-zero.
func (h Homogen) XY() (x, y float32) {
	return h.X / h.Z, h.Y / h.Z
}

// Cartesian returns the cartesian representation of the vector (i.e. the
// perspective divide of the homogeneous coordinates). Z must be non-zero.
func (h Homogen) Cartesian() Coord {
	return Coord{h.X / h.Z, h.Y / h.Z}
}

//------------------------------------------------------------------------------

// Polar represents a two dimensional vector, defined by its polar coordinates.
type Polar struct {
	R     float32 // Radius (i.e. distance from origin)
	Theta float32 // Angle (counter-clockwise from 3 o'clock)
}

// XY returns the cartesian coordinates of the vector. This implements the
// Vector interface.
func (p Polar) XY() (x, y float32) {
	return p.R * math32.Cos(p.Theta), p.R * math32.Sin(p.Theta)
}

// Cartesian returns the cartesian representation of the vector.
func (p Polar) Cartesian() Coord {
	return Coord{p.R * math32.Cos(p.Theta), p.R * math32.Sin(p.Theta)}
}

//------------------------------------------------------------------------------
