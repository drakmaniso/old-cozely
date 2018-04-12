// Copyright (c) 2013-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

/*
Package coord provides vector types for various coordinates systems.

Floating-point Vectors

Three structs with float32 fields give access to geometric transformations, and
other useful mathematical operations (see subpackages plane and space):

  // 2D cartesian coordinates
  var a = XY{1.0, 2.0}

  // 3D cartesian coordinates (or 2D projective coordinates)
  var b = XYZ{1.0, 2.0, 3.0}

  // 3D projective coordinates
  var c = XYZW{1.0, 2.0, 3.0, 4.0}

Integer Vectors

Two structs with int16 fields are used both for on-screen coordinates and
in-game grids:

  // 2D cartesian grid coordinates (e.g. pixel coordinates)
  var a = CR{1, 2}

  // CRD for 3D cartesian grid coordinates (e.g. voxel coordinates)
  var b = CRD{1, 2, 3}

Other Coordinate systems

Various structs for hexagonal and triangular grids:

	// Hexagonal grid coordinates
	var a = QR{1, 2}
	// Axial (hex) coordinates
	var b = AL{1.0, 2.0}
	// triangular grid coordinates
	var c = QRT{1, 2, 3}

And finally structs to manipulate angles:

	// Polar coordinates
	var a = RA{1.0, 3.14}
	// Cylindrical coordinates
	var b = RAZ{1.0, 3.14, 2.0}
	// Spherical coordinates
	var c = RAS{1.0, 3.14, 6.28}
*/
package coord
