// Copyright (c) 2013-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package plane

////////////////////////////////////////////////////////////////////////////////

// CR represents the position of a cell on a square grid, defined by column C
// and row R (i.e., a pair of integer coordinates).
//
// Note that the R axis is in the opposition direction than the cartesian
// coordinates (Coord).
type CR struct {
	C, R int16
}

// Cartesian returns the cartesian coordinates of the grid cell.
//
// Note that the sign of R is flipped. This function implements the Vector
// interface.
func (p CR) Cartesian() (x, y float32) {
	return float32(p.C), float32(-p.R)
}

// XY returns the cartesian coordinates of the grid cell.
//
// Note that the sign of R is flipped.
func (p CR) XY() XY {
	return XY{float32(p.C), float32(-p.R)}
}

////////////////////////////////////////////////////////////////////////////////

// Plus returns the sum with another pair of coordinates.
func (p CR) Plus(o CR) CR {
	return CR{p.C + o.C, p.R + o.R}
}

// Pluss returns the sum with another pair of coordinates.
func (p CR) Pluss(c, r int16) CR {
	return CR{p.C + c, p.R + r}
}

// Minus returns the difference with another pair of coordinates.
func (p CR) Minus(o CR) CR {
	return CR{p.C - o.C, p.R - o.R}
}

// Minuss returns the difference with another pair of coordinates.
func (p CR) Minuss(c, r int16) CR {
	return CR{p.C - c, p.R - r}
}

// Opposite returns the opposite pair of coordinates.
func (p CR) Opposite() CR {
	return CR{-p.C, -p.R}
}

// Times returns the product with a scalar.
func (p CR) Times(s int16) CR {
	return CR{p.C * s, p.R * s}
}

// Timess returns the component-wise product with two scalars.
func (p CR) Timess(c, r int16) CR {
	return CR{p.C * c, p.R * r}
}

// Timescw returns the component-wise product with another pair of coordinates.
func (p CR) Timescw(o CR) CR {
	return CR{p.C * o.C, p.R * o.R}
}

// Slash returns the integer quotient of the division by a scalar (which must be
// non-zero).
func (p CR) Slash(s int16) CR {
	return CR{p.C / s, p.R / s}
}

// Slashs returns the component-wise integer quotient of the division by two
// scalars (which must be non-zero).
func (p CR) Slashs(c, r int16) CR {
	return CR{p.C / c, p.R / r}
}

// Slashcw returns the integer quotients of the component-wise division by
// another pair of coordinates (of which both C and R must be non-zero).
func (p CR) Slashcw(o CR) CR {
	return CR{p.C / o.C, p.R / o.R}
}

// Mod returns the remainder (modulus) of the division by a scalar (which must
// be non-zero).
func (p CR) Mod(s int16) CR {
	return CR{p.C % s, p.R % s}
}

// Mods returns the remainders (modulus) of the component-wise division by two
// scalars (which must be non-zero).
func (p CR) Mods(c, r int16) CR {
	return CR{p.C % c, p.R % r}
}

// Modcw returns the remainder (modulus) of the component-wise division by
// another pair of coordinates (of which both C and R must be non-zero).
func (p CR) Modcw(o CR) CR {
	return CR{p.C % o.C, p.R % o.R}
}

////////////////////////////////////////////////////////////////////////////////

// FlipC returns the coordinates with the signe of C flipped.
func (p CR) FlipC() CR {
	return CR{-p.C, p.R}
}

// FlipR returns the coordinates with the signe of R flipped.
func (p CR) FlipR() CR {
	return CR{p.C, -p.R}
}

// Col returns the coordinates projected on the C axis (i.e. with R nulled).
func (p CR) Col() CR {
	return CR{p.C, 0}
}

// Row returns the coordinates projected on the R axis (i.e. with C nulled).
func (p CR) Row() CR {
	return CR{0, p.R}
}

// RC returns the coordinates with C and R swapped.
func (p CR) RC() CR {
	return CR{p.R, p.C}
}

// Perp returns the coordinates rotated by 90 in counter-clockwise direction.
func (p CR) Perp() CR {
	return CR{-p.R, p.C}
}

////////////////////////////////////////////////////////////////////////////////
