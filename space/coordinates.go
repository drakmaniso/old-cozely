// Copyright (c) 2013-2017 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package space

//------------------------------------------------------------------------------

import (
	"github.com/drakmaniso/glam/math"
)

//------------------------------------------------------------------------------

// Vector encapsulates any type that can represent an euclidian vector (ie can
// represent a point in space).
type Vector interface {
	Cartesian() (x, y, z float32)
}

//------------------------------------------------------------------------------

// Coord is a cartesian coordinate vector.
type Coord struct {
	X float32
	Y float32
	Z float32
}

// NewCoord returns a new Coord corresponding to v.
func NewCoord(v Vector) Coord {
	x, y, z := v.Cartesian()
	return Coord{x, y, z}
}

// Cartesian returns X, Y and Z. This function is here to implement the Vector
// interface.
func (v Coord) Cartesian() (x, y, z float32) {
	return v.X, v.Y, v.Z
}

// Homogen returns an homogenous coordinate vector corresponding to v, with W
// set to 1.
func (v Coord) Homogen() Homogen {
	return Homogen{v.X, v.Y, v.Z, 1.0}
}

// Plus returns the sum with another coordinate vector.
func (v Coord) Plus(o Coord) Coord {
	return Coord{v.X + o.X, v.Y + o.Y, v.Z + o.Z}
}

// Minus returns the difference with another coordinate vector.
func (v Coord) Minus(o Coord) Coord {
	return Coord{v.X - o.X, v.Y - o.Y, v.Z - o.Z}
}

// Inverse returns the product with another coordinate vector.
func (v Coord) Inverse() Coord {
	return Coord{-v.X, -v.Y, -v.Z}
}

// Times returns the product with another coordinate vector.
func (v Coord) Times(o Coord) Coord {
	return Coord{v.X * o.X, v.Y * o.Y, v.Z * o.Z}
}

// Slash returns the division by another coordinate vector (of which both X and
// Y must be non-zero).
func (v Coord) Slash(o Coord) Coord {
	return Coord{v.X / o.X, v.Y / o.Y, v.Z / o.Z}
}

// Dot returns the dot product with another coordinate vector.
func (v Coord) Dot(o Coord) float32 {
	return v.X*o.X + v.Y*o.Y + v.Z*o.Z
}

// Cross returns the cross product with another vector.
func (v Coord) Cross(o Coord) Coord {
	return Coord{
		v.Y*o.Z - v.Z*o.Y,
		v.Z*o.X - v.X*o.Z,
		v.X*o.Y - v.Y*o.X,
	}
}

// Length returns the euclidian length of the vector.
func (v Coord) Length() float32 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y + v.Z*v.Z)
}

// Normalized return the normalization of the vector (i.e. the vector divided
// by its length).
//
// Important: Length() must be non-zero.
func (v Coord) Normalized() Coord {
	l := v.Length()
	return Coord{v.X / l, v.Y / l, v.Z / l}
}

// IsAlmostEqual returns true if the difference between the two vectors is less
// than the specified ULPs (Unit in the Last Place).
//
// Handle special cases: zero, infinites, denormals.
//
// See also IsNearlyEqual and IsRoughlyEqual.
func (v Coord) IsAlmostEqual(o Coord, ulps uint32) bool {
	return math.IsAlmostEqual(v.X, o.X, ulps) &&
		math.IsAlmostEqual(v.Y, o.Y, ulps) &&
		math.IsAlmostEqual(v.Z, o.Z, ulps)
}

// IsNearlyEqual Returns true if the relative error between the two vectors is
// less than epsilon.
//
// Handles special cases: zero, infinites, denormals.
//
// See also IsAlmostEqual and IsRoughlyEqual.
func (v Coord) IsNearlyEqual(o Coord, epsilon float32) bool {
	return math.IsNearlyEqual(v.X, o.X, epsilon) &&
		math.IsNearlyEqual(v.Y, o.Y, epsilon) &&
		math.IsNearlyEqual(v.Z, o.Z, epsilon)
}

// IsRoughlyEqual Returns true if the absolute error between the two vectors is
// less than epsilon.
//
// See also IsNearlyEqual and IsAlmostEqual.
func (v Coord) IsRoughlyEqual(o Coord, epsilon float32) bool {
	return math.IsRoughlyEqual(v.X, o.X, epsilon) &&
		math.IsRoughlyEqual(v.Y, o.Y, epsilon) &&
		math.IsRoughlyEqual(v.Z, o.Z, epsilon)
}

//------------------------------------------------------------------------------

// Homogen is an homogeneous coordinate vector.
type Homogen struct {
	X float32
	Y float32
	Z float32
	W float32
}

// Cartesian implements the Vector interface: it returns the dehomogenization of
// the vector (i.e. perspective divide).
//
// Important: v.W must be non-zero.
func (v Homogen) Cartesian() (x, y, z float32) {
	return v.X / v.W, v.Y / v.W, v.Z / v.W
}

// Coord returns the dehomogenization of the vector (i.e. perspective divide).
//
// Important: v.W must be non-zero.
func (v Homogen) Coord() Coord {
	return Coord{v.X / v.W, v.Y / v.W, v.Z / v.W}
}

// Plus returns the sum with another coordinate vector.
func (v Homogen) Plus(o Homogen) Homogen {
	return Homogen{v.X + o.X, v.Y + o.Y, v.Z + o.Z, v.W + o.W}
}

// Minus returns the difference with another coordinate vector.
func (v Homogen) Minus(o Homogen) Homogen {
	return Homogen{v.X - o.X, v.Y - o.Y, v.Z - o.Z, v.W - o.W}
}

// Inverse returns the product with another coordinate vector.
func (v Homogen) Inverse() Homogen {
	return Homogen{-v.X, -v.Y, -v.Z, -v.W}
}

// Times returns the product with another coordinate vector.
func (v Homogen) Times(o Homogen) Homogen {
	return Homogen{v.X * o.X, v.Y * o.Y, v.Z * o.Z, v.W * o.W}
}

// Slash returns the division by another coordinate vector (of which both X and
// Y must be non-zero).
func (v Homogen) Slash(o Homogen) Homogen {
	return Homogen{v.X / o.X, v.Y / o.Y, v.Z / o.Z, v.W / o.W}
}

// Dot returns the dot product with another coordinate vector.
func (v Homogen) Dot(o Homogen) float32 {
	return v.X*o.X + v.Y*o.Y + v.Z*o.Z + v.W*o.W
}

// Length returns the euclidian length of the vector.
func (v Homogen) Length() float32 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y + v.Z*v.Z + v.W*v.W)
}

// Normalized return the normalization of the vector (i.e. the vector divided
// by its length).
//
// Important: Length() must be non-zero.
func (v Homogen) Normalized() Homogen {
	l := v.Length()
	return Homogen{v.X / l, v.Y / l, v.Z / l, v.W / l}
}

//------------------------------------------------------------------------------

//TODO: Spherical and Cylindrical types

//------------------------------------------------------------------------------
