// Copyright (c) 2013-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package coord

////////////////////////////////////////////////////////////////////////////////

// CR represents 2D integer coordinates. C stands for "column" and corresponds
// to the x-coordinate, while R stands for "row" and corresponds to the
// y-coordinate.
type CR struct {
	C, R int16
}

// CRof returns an integer vector corresponding to the first two coordinates of
// v.
func CRof(v Coordinates) CR {
	x, y, _ := v.Cartesian()
	return CR{int16(x), int16(y)}
}

// Cartesian returns the coordinates of the vector in 3D space. C and R are
// casted to float32, and the third coordinate is always 0.
func (a CR) Cartesian() (x, y, z float32) {
	return float32(a.C), float32(a.R), 0
}

// CRD returns the integer coordinates of the vector in 3D space, with d as
// third coordinate.
func (a CR) CRD(d int16) CRD {
	return CRD{a.C, a.R, d}
}

// XY returns the floating point coordinates of the vector.
func (a CR) XY() XY {
	return XY{float32(a.C), float32(a.R)}
}

// XYZ returns the floating-point coordinates in 3D space, with the third
// dimension set to z.
func (a CR) XYZ(z float32) XYZ {
	return XYZ{float32(a.C), float32(a.R), z}
}

// XYZW returns the floating-point coordinates in 3D projective space, with the
// third and fourth dimensions set to z and w.
func (a CR) XYZW(z, w float32) XYZW {
	return XYZW{float32(a.C), float32(a.R), z, w}
}

////////////////////////////////////////////////////////////////////////////////

// Plus returns the sum with another vector.
func (a CR) Plus(b CR) CR {
	return CR{a.C + b.C, a.R + b.R}
}

// Pluss returns the sum with the vector (s, s).
func (a CR) Pluss(s int16) CR {
	return CR{a.C + s, a.R + s}
}

// Minus returns the difference with another vector.
func (a CR) Minus(b CR) CR {
	return CR{a.C - b.C, a.R - b.R}
}

// Minuss returns the difference with the vector (s, s).
func (a CR) Minuss(s int16) CR {
	return CR{a.C - s, a.R - s}
}

// Opposite returns the opposite of the vector.
func (a CR) Opposite() CR {
	return CR{-a.C, -a.R}
}

// Timess returns the product with a scalar.
func (a CR) Timess(s int16) CR {
	return CR{a.C * s, a.R * s}
}

// Times returns the component-wise product with another vector.
func (a CR) Times(b CR) CR {
	return CR{a.C * b.C, a.R * b.R}
}

// Slashs returns the integer quotient of the division by a scalar (which must be
// non-zero).
func (a CR) Slashs(s int16) CR {
	return CR{a.C / s, a.R / s}
}

// Slash returns the integer quotients of the component-wise division by
// another vector (of which both C and R must be non-zero).
func (a CR) Slash(b CR) CR {
	return CR{a.C / b.C, a.R / b.R}
}

// Mods returns the remainder (modulus) of the division by a scalar (which must
// be non-zero).
func (a CR) Mods(s int16) CR {
	return CR{a.C % s, a.R % s}
}

// Mod returns the remainder (modulus) of the component-wise division by
// another vector (of which both C and R must be non-zero).
func (a CR) Mod(b CR) CR {
	return CR{a.C % b.C, a.R % b.R}
}

////////////////////////////////////////////////////////////////////////////////

// FlipC returns the vector with the sign of C flipped.
func (a CR) FlipC() CR {
	return CR{-a.C, a.R}
}

// FlipR returns the vector with the sign of R flipped.
func (a CR) FlipR() CR {
	return CR{a.C, -a.R}
}

// Col returns the vector projected on the C axis (i.e. with R nulled).
func (a CR) Col() CR {
	return CR{a.C, 0}
}

// Row returns the vector projected on the R axis (i.e. with C nulled).
func (a CR) Row() CR {
	return CR{0, a.R}
}

// RC returns the vector with coordinates C and R swapped.
func (a CR) RC() CR {
	return CR{a.R, a.C}
}

// Perp returns the vector rotated by 90 in counter-clockwise direction.
func (a CR) Perp() CR {
	return CR{-a.R, a.C}
}

////////////////////////////////////////////////////////////////////////////////

// Null returns true if both coordinates are null.
func (a CR) Null() bool {
	return a.C == 0 && a.R == 0
}
