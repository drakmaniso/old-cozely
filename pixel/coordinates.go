// Copyright (c) 2013-2017 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package pixel

//------------------------------------------------------------------------------

// Coord represents a pair of coordinates on the virtual screen, in (virtual)
// pixels.
type Coord struct {
	X, Y int16
}

// XY returns the cartesian coordinates of the vector.
//
// This function implements the plane.Vector interface.
func (p Coord) XY() (x, y float32) {
	return float32(p.X), float32(p.Y)
}

//------------------------------------------------------------------------------

// Plus returns the sum with another pair of coordinates.
func (p Coord) Plus(o Coord) Coord {
	return Coord{p.X + o.X, p.Y + o.Y}
}

// Minus returns the difference with another pair of coordinates.
func (p Coord) Minus(o Coord) Coord {
	return Coord{p.X - o.X, p.Y - o.Y}
}

// Opposite returns the opposite pair of coordinates.
func (p Coord) Opposite() Coord {
	return Coord{-p.X, -p.Y}
}

// Times returns the product with a scalar.
func (p Coord) Times(s int16) Coord {
	return Coord{p.X * s, p.Y * s}
}

// TimesCW returns the component-wise product with another pair of coordinates.
func (p Coord) TimesCW(o Coord) Coord {
	return Coord{p.X * o.X, p.Y * o.Y}
}

// Slash returns the integer quotient of the division by a scalar (which must be
// non-zero).
func (p Coord) Slash(s int16) Coord {
	return Coord{p.X / s, p.Y / s}
}

// SlashCW returns the integer quotients of the component-wise division by
// another pair of coordinates (of which both X and Y must be non-zero).
func (p Coord) SlashCW(o Coord) Coord {
	return Coord{p.X / o.X, p.Y / o.Y}
}

// Mod returns the remainder (modulus) of the division by a scalar (which must
// be non-zero).
func (p Coord) Mod(s int16) Coord {
	return Coord{p.X % s, p.Y % s}
}

// ModCW returns the remainder (modulus) of the component-wise division by
// another pair of coordinates (of which both X and Y must be non-zero).
func (p Coord) ModCW(o Coord) Coord {
	return Coord{p.X % o.X, p.Y % o.Y}
}

//------------------------------------------------------------------------------
