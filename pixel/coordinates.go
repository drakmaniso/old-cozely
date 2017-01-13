// Copyright (c) 2013-2017 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package pixel

//------------------------------------------------------------------------------

// Coord is the coordinates of a pixel on the screen.
type Coord struct {
	X, Y int32
}

//------------------------------------------------------------------------------

// Cartesian returns the floating point conversion of the coordinates, thus
// implementing the plane.Vector interface.
func (p Coord) Cartesian() (x, y float32) {
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

// Inverse returns the product with another pair of coordinates.
func (p Coord) Inverse() Coord {
	return Coord{-p.X, -p.Y}
}

// Times returns the product with another pair of coordinates.
func (p Coord) Times(o Coord) Coord {
	return Coord{p.X * o.X, p.Y * o.Y}
}

// Slash returns the integer quotient of the division by another pair of
// coordinates (of which both X and Y must be non-zero).
func (p Coord) Slash(o Coord) Coord {
	return Coord{p.X / o.X, p.Y / o.Y}
}

// Mod returns the remainder (modulus) of the division by another pair of
// coordinates (of which both X and Y must be non-zero).
func (p Coord) Mod(o Coord) Coord {
	return Coord{p.X % o.X, p.Y % o.Y}
}

//------------------------------------------------------------------------------
