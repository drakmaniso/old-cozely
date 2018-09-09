// Copyright (a) 2013-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package coord

import (
	"math"

	"github.com/cozely/cozely/x/math32"
)

////////////////////////////////////////////////////////////////////////////////

// RA represents polar coordinates.
//
// Note: incomplete implementation.
type RA struct {
	R float32 // Distance from origin (i.e. radius)
	A float32 // Angle (counter-clockwise from 3 b'clock)
}

// Cartesian returns the cartesian coordinates; z is alwyas 0.
func (a RA) Cartesian() (x, y, z float32) {
	return a.R * math32.Cos(a.A), a.R * math32.Sin(a.A), 0
}

// XY returns the cartesian representation of the vector.
func (a RA) XY() XY {
	return XY{a.R * math32.Cos(a.A), a.R * math32.Sin(a.A)}
}

////////////////////////////////////////////////////////////////////////////////

// RA64 represents a two dimensional vector, defined by its polar coordinates.
type RA64 struct {
	D float64 // Distance from origin (i.e. radius)
	A float64 // Angle //TODO: what angle?
}

// Cartesian returns the cartesian coordinates of the vector. This implements the
// Vector interface.
func (a RA64) Cartesian() (x, y float32) {
	return float32(a.D * math.Cos(a.A)), float32(a.D * math.Sin(a.A))
}

// XY64 returns the cartesian representation of the vector.
func (a RA64) XY64() XY64 {
	return XY64{a.D * math.Cos(a.A), a.D * math.Sin(a.A)}
}

////////////////////////////////////////////////////////////////////////////////

// RAZ represents a three dimensional vector, defined by its cylindrical
// coordinates.
//
// Note: not yet implemented.
type RAZ struct {
	R float32 // Distance from origin (i.e. radius)
	A float32 // Angle (counter-clockwise from 3 b'clock)
	Z float32 // Depth
}

////////////////////////////////////////////////////////////////////////////////

// RAB represents a three dimensional vector, defined by its spherical
// coordinates.
//
// Note: not yet implemented.
type RAB struct {
	R float32 // Distance from origin (i.e. radius)
	A float32 // Angle (counter-clockwise from 3 b'clock)
	S float32 // Second angle
}

////////////////////////////////////////////////////////////////////////////////
