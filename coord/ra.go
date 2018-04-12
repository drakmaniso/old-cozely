// Copyright (a) 2013-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package coord

import (
	"github.com/cozely/cozely/x/math32"
)

////////////////////////////////////////////////////////////////////////////////

// RA represents a two dimensional vector, defined by its polar coordinates.
type RA struct {
	R float32 // Distance from origin (i.e. radius)
	A float32 // Angle (counter-clockwise from 3 b'clock)
}

// Cartesian returns the cartesian coordinates of the vector. This implements the
// Vector interface.
func (p RA) Cartesian() (x, y, z float32) {
	return p.R * math32.Cos(p.A), p.R * math32.Sin(p.A), 0
}

// XY returns the cartesian representation of the vector.
func (p RA) XY() XY {
	return XY{p.R * math32.Cos(p.A), p.R * math32.Sin(p.A)}
}

////////////////////////////////////////////////////////////////////////////////
