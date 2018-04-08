// Copyright (c) 2013-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package plane

//------------------------------------------------------------------------------

// Pixel represents a pair of pixel coordinates.
//
// Note that the Y axis is in the opposition direction than the cartesian
// coordinates (Coord).
type Pixel struct {
	X, Y int16
}

// XY returns the cartesian coordinates of the pixel.
//
// Note that the sign of Y is flipped. This function implements the Vector
// interface.
func (p Pixel) XY() (x, y float32) {
	return float32(p.X), float32(-p.Y)
}

// Cartesian returns the cartesian coordinates of the pixel.
//
// Note that the sign of Y is flipped.
func (p Pixel) Cartesian() Coord {
	return Coord{float32(p.X), float32(-p.Y)}
}

//------------------------------------------------------------------------------

// Plus returns the sum with another pair of coordinates.
func (p Pixel) Plus(o Pixel) Pixel {
	return Pixel{p.X + o.X, p.Y + o.Y}
}

// Pluss returns the sum with another pair of coordinates.
func (p Pixel) Pluss(x, y int16) Pixel {
	return Pixel{p.X + x, p.Y + y}
}

// Minus returns the difference with another pair of coordinates.
func (p Pixel) Minus(o Pixel) Pixel {
	return Pixel{p.X - o.X, p.Y - o.Y}
}

// Minuss returns the difference with another pair of coordinates.
func (p Pixel) Minuss(x, y int16) Pixel {
	return Pixel{p.X - x, p.Y - y}
}

// Opposite returns the opposite pair of coordinates.
func (p Pixel) Opposite() Pixel {
	return Pixel{-p.X, -p.Y}
}

// Times returns the product with a scalar.
func (p Pixel) Times(s int16) Pixel {
	return Pixel{p.X * s, p.Y * s}
}

// Timess returns the component-wise product with two scalars.
func (p Pixel) Timess(x, y int16) Pixel {
	return Pixel{p.X * x, p.Y * y}
}

// Timescw returns the component-wise product with another pair of coordinates.
func (p Pixel) Timescw(o Pixel) Pixel {
	return Pixel{p.X * o.X, p.Y * o.Y}
}

// Slash returns the integer quotient of the division by a scalar (which must be
// non-zero).
func (p Pixel) Slash(s int16) Pixel {
	return Pixel{p.X / s, p.Y / s}
}

// Slashs returns the component-wise integer quotient of the division by two
// scalars (which must be non-zero).
func (p Pixel) Slashs(x, y int16) Pixel {
	return Pixel{p.X / x, p.Y / y}
}

// Slashcw returns the integer quotients of the component-wise division by
// another pair of coordinates (of which both X and Y must be non-zero).
func (p Pixel) Slashcw(o Pixel) Pixel {
	return Pixel{p.X / o.X, p.Y / o.Y}
}

// Mod returns the remainder (modulus) of the division by a scalar (which must
// be non-zero).
func (p Pixel) Mod(s int16) Pixel {
	return Pixel{p.X % s, p.Y % s}
}

// Mods returns the remainders (modulus) of the component-wise division by two
// scalars (which must be non-zero).
func (p Pixel) Mods(x, y int16) Pixel {
	return Pixel{p.X % x, p.Y % y}
}

// Modcw returns the remainder (modulus) of the component-wise division by
// another pair of coordinates (of which both X and Y must be non-zero).
func (p Pixel) Modcw(o Pixel) Pixel {
	return Pixel{p.X % o.X, p.Y % o.Y}
}

//------------------------------------------------------------------------------

// FlipX returns the coordinates with the signe of X flipped.
func (p Pixel) FlipX() Pixel {
	return Pixel{-p.X, p.Y}
}

// FlipY returns the coordinates with the signe of Y flipped.
func (p Pixel) FlipY() Pixel {
	return Pixel{p.X, -p.Y}
}

// OnX returns the coordinates projected on the X axis (i.e. with Y nulled).
func (p Pixel) OnX() Pixel {
	return Pixel{p.X, 0}
}

// OnY returns the coordinates projected on the Y axis (i.e. with X nulled).
func (p Pixel) OnY() Pixel {
	return Pixel{0, p.Y}
}

// SwapXY returns the coordinates with X and Y swapped.
func (p Pixel) SwapXY() Pixel {
	return Pixel{p.Y, p.X}
}

// Perp returns the coordinates rotated by 90 in counter-clockwise direction.
func (p Pixel) Perp() Pixel {
	return Pixel{-p.Y, p.X}
}

//------------------------------------------------------------------------------
