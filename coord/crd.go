// Copyright (c) 2018-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package coord

////////////////////////////////////////////////////////////////////////////////

// CR represents the position of a cell on a cubic grid, defined by column C,
// row R and depth D (i.e., three integer coordinates).
type CRD struct {
	C, R, D int16
}

// Cartesian returns the cartesian coordinates of the grid cell.
func (a CRD) Cartesian() (x, y, z float32) {
	return float32(a.C), float32(a.R), float32(a.D)
}

// CR returns the square grid cell (C, R).
func (a CRD) CR() CR {
	return CR{a.C, a.R}
}

////////////////////////////////////////////////////////////////////////////////

// Pluss returns the sum with another pair of coordinates.
func (a CRD) Pluss(c, r, d int16) CRD {
	return CRD{a.C + c, a.R + r, a.D +d}
}

////////////////////////////////////////////////////////////////////////////////
