// Copyright (c) 2013-2017 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package pixel

//------------------------------------------------------------------------------

// XY is the cartesian coordinates of a pixel on the screen.
type XY struct {
	X, Y int32
}

//------------------------------------------------------------------------------

// Cartesian implements the plane.Coordinates interface.
func (p XY) Cartesian() (x, y float32) {
	return float32(p.X), float32(p.Y)
}

//------------------------------------------------------------------------------

// Plus returns the sum with other coordinates.
func (p XY) Plus(o XY) XY {
	return XY{p.X + o.X, p.Y + o.Y}
}

// Minus returns the difference with other coordinates.
func (p XY) Minus(o XY) XY {
	return XY{p.X - o.X, p.Y - o.Y}
}

// Inverse returns the product with other coordinates.
func (p XY) Inverse() XY {
	return XY{-p.X, -p.Y}
}

// Times returns the product with other coordinates.
func (p XY) Times(o XY) XY {
	return XY{p.X * o.X, p.Y * o.Y}
}

// Slash returns the integer quotient of the division by other coordinates (of
// which both X and Y must be non-zero).
func (p XY) Slash(o XY) XY {
	return XY{p.X / o.X, p.Y / o.Y}
}

// Mod returns the remainder (modulus) of the division by other coordinates (of
// which both X and Y must be non-zero).
func (p XY) Mod(o XY) XY {
	return XY{p.X % o.X, p.Y % o.Y}
}

//------------------------------------------------------------------------------
