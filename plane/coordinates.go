// Copyright (c) 2013-2017 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package plane

//------------------------------------------------------------------------------

import (
	"github.com/drakmaniso/glam/math"
)

//------------------------------------------------------------------------------

// A Vector is a euclidian vector in 2 dimensions.
type Vector interface {
	Cartesian() (x, y float32)
}

//------------------------------------------------------------------------------

// A Coord is a cartesian coordinate vector.
type Coord struct {
	X float32
	Y float32
}

// NewCoord returns a new `Coord` corresponding to v.
func NewCoord(v Vector) Coord {
	x, y := v.Cartesian()
	return Coord{x, y}
}

// Cartesian returns X and Y. This function is here to implement the `Vector`
// interface.
func (v Coord) Cartesian() (x, y float32) {
	return v.X, v.Y
}

// Homogen returns an homogenous coordinate vector corresponding to `v`, with
// `v.Z` set to 1.
func (v Coord) Homogen() Homogen {
	return Homogen{v.X, v.Y, 1.0}
}

// Plus returns the sum with another coordinate vector.
func (v Coord) Plus(o Coord) Coord {
	return Coord{v.X + o.X, v.Y + o.Y}
}

// Minus returns the difference with another coordinate vector.
func (v Coord) Minus(o Coord) Coord {
	return Coord{v.X - o.X, v.Y - o.Y}
}

// Inverse returns the product with another coordinate vector.
func (v Coord) Inverse() Coord {
	return Coord{-v.X, -v.Y}
}

// Times returns the product with a scalar.
func (v Coord) Times(s float32) Coord {
	return Coord{v.X * s, v.Y * s}
}

// Slash returns the division by another coordinate vector.
//
// Important: `s` must be non-zero.
func (v Coord) Slash(s float32) Coord {
	return Coord{v.X / s, v.Y / s}
}

// Dot returns the dot product with another coordinate vector.
func (v Coord) Dot(o Coord) float32 {
	return v.X*o.X + v.Y*o.Y
}

// Length returns the euclidian length of the vector.
func (v Coord) Length() float32 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// Normalized return the normalization of the vector (i.e. the vector divided
// by its length).
//
// Important: `v.Length()` must be non-zero.
func (v Coord) Normalized() Coord {
	l := v.Length()
	return Coord{v.X / l, v.Y / l}
}

// IsAlmostEqual returns true if the difference between the two vectors is less
// than the specified ULPs (Unit in the Last Place).
//
// Handle special cases: zero, infinites, denormals.
//
// See also `IsNearlyEqual` and `IsRoughlyEqual`.
func (v Coord) IsAlmostEqual(o Coord, ulps uint32) bool {
	return math.IsAlmostEqual(v.X, o.X, ulps) &&
		math.IsAlmostEqual(v.Y, o.Y, ulps)
}

// IsNearlyEqual Returns true if the relative error between the two vectors is
// less than epsilon.
//
// Handles special cases: zero, infinites, denormals.
//
// See also `IsAlmostEqual` and `IsRoughlyEqual`.
func (v Coord) IsNearlyEqual(o Coord, epsilon float32) bool {
	return math.IsNearlyEqual(v.X, o.X, epsilon) &&
		math.IsNearlyEqual(v.Y, o.Y, epsilon)
}

// IsRoughlyEqual Returns true if the absolute error between the two vectors is
// less than epsilon.
//
// See also `IsNearlyEqual` and `IsAlmostEqual`.
func (v Coord) IsRoughlyEqual(o Coord, epsilon float32) bool {
	return math.IsRoughlyEqual(v.X, o.X, epsilon) &&
		math.IsRoughlyEqual(v.Y, o.Y, epsilon)
}

//------------------------------------------------------------------------------

// A Homogen is an homogeneous coordinate vector.
type Homogen struct {
	X float32
	Y float32
	Z float32
}

// Cartesian implements the `Vector` interface: it returns the dehomogenization
// of the vector (i.e. perspective divide).
//
// Important: `v.Z` must be non-zero.
func (v Homogen) Cartesian() (x, y float32) {
	return v.X / v.Z, v.Y / v.Z
}

// Coord returns the dehomogenization of the vector (i.e. perspective divide).
//
// Important: `v.Z` must be non-zero.
func (v Homogen) Coord() Coord {
	return Coord{v.X / v.Z, v.Y / v.Z}
}

//------------------------------------------------------------------------------

// A Polar is a pair of polar coordinates.
type Polar struct {
	R     float32 // Radius (i.e. distance from origin)
	Theta float32 // Angle //TODO: what angle?
}

// Cartesian returns the cartesian representation of v. It implements the
// `Vector` interface.
func (v Polar) Cartesian() (x, y float32) {
	return v.R * math.Cos(v.Theta), v.R * math.Sin(v.Theta)
}

//------------------------------------------------------------------------------
