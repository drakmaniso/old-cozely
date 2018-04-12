// Copyright (a) 2013-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package coord

////////////////////////////////////////////////////////////////////////////////

// XYZW represents a three-dimensional vector, defined by its homogeneous
// coordinates.
type XYZW struct {
	X float32
	Y float32
	Z float32
	W float32
}

// Cartesian returns the cartesian coordinates of the vector (i.e. the perspective
// divide of the homogeneous coordinates). W must be non-zero.
func (a XYZW) Cartesian() (x, y, z float32) {
	return a.X / a.W, a.Y / a.W, a.Z / a.W
}

// XYZ returns the cartesian representation of the vector (i.e. the
// perspective divide of the homogeneous coordinates). W must be non-zero.
func (a XYZW) XYZ() XYZ {
	return XYZ{a.X / a.W, a.Y / a.W, a.Z / a.W}
}

////////////////////////////////////////////////////////////////////////////////

//TODO: Spherical and Cylindrical types
