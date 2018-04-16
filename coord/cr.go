// Copyright (c) 2013-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package coord

////////////////////////////////////////////////////////////////////////////////

// CR represents the position of a cell on a square grid, defined by column C
// and row R (i.e., a pair of integer coordinates).
type CR struct {
	C, R int16
}

// Cartesian returns the cartesian coordinates of the grid cell.
func (a CR) Cartesian() (x, y, z float32) {
	return float32(a.C), float32(a.R), 0
}

// XY returns the cartesian coordinates of the grid cell.
func (a CR) XY() XY {
	return XY{float32(a.C), float32(a.R)}
}

// CRD returns the cubic grid cell (C, R, d)
func (a CR) CRD(d int16) CRD {
	return CRD{a.C, a.R, d}
}

////////////////////////////////////////////////////////////////////////////////

// Plus returns the sum with another pair of coordinates.
func (a CR) Plus(b CR) CR {
	return CR{a.C + b.C, a.R + b.R}
}

// Pluss returns the sum with a scalar.
func (a CR) Pluss(s int16) CR {
	return CR{a.C + s, a.R + s}
}

// Minus returns the difference with another pair of coordinates.
func (a CR) Minus(b CR) CR {
	return CR{a.C - b.C, a.R - b.R}
}

// Minuss returns the difference with a scalar.
func (a CR) Minuss(s int16) CR {
	return CR{a.C - s, a.R - s}
}

// Opposite returns the opposite pair of coordinates.
func (a CR) Opposite() CR {
	return CR{-a.C, -a.R}
}

// Times returns the product with a scalar.
func (a CR) Times(s int16) CR {
	return CR{a.C * s, a.R * s}
}

// Timescr returns the component-wise product with another pair of coordinates.
func (a CR) Timescr(b CR) CR {
	return CR{a.C * b.C, a.R * b.R}
}

// Slash returns the integer quotient of the division by a scalar (which must be
// non-zero).
func (a CR) Slash(s int16) CR {
	return CR{a.C / s, a.R / s}
}

// Slashcr returns the integer quotients of the component-wise division by
// another pair of coordinates (of which both C and R must be non-zero).
func (a CR) Slashcr(b CR) CR {
	return CR{a.C / b.C, a.R / b.R}
}

// Mod returns the remainder (modulus) of the division by a scalar (which must
// be non-zero).
func (a CR) Mod(s int16) CR {
	return CR{a.C % s, a.R % s}
}

// Modcr returns the remainder (modulus) of the component-wise division by
// another pair of coordinates (of which both C and R must be non-zero).
func (a CR) Modcr(b CR) CR {
	return CR{a.C % b.C, a.R % b.R}
}

////////////////////////////////////////////////////////////////////////////////

// FlipC returns the coordinates with the signe of C flipped.
func (a CR) FlipC() CR {
	return CR{-a.C, a.R}
}

// FlipR returns the coordinates with the signe of R flipped.
func (a CR) FlipR() CR {
	return CR{a.C, -a.R}
}

// Col returns the coordinates projected on the C axis (i.e. with R nulled).
func (a CR) Col() CR {
	return CR{a.C, 0}
}

// Row returns the coordinates projected on the R axis (i.e. with C nulled).
func (a CR) Row() CR {
	return CR{0, a.R}
}

// RC returns the coordinates with C and R swapped.
func (a CR) RC() CR {
	return CR{a.R, a.C}
}

// Perp returns the coordinates rotated by 90 in counter-clockwise direction.
func (a CR) Perp() CR {
	return CR{-a.R, a.C}
}

////////////////////////////////////////////////////////////////////////////////
