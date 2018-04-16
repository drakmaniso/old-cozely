// Copyright (c) 2018-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package coord

////////////////////////////////////////////////////////////////////////////////

// CRD is a vector of 3D integer coordinates. C stands for "column" and
// corresponds to the x-coordinate, R stands for "row" and corresponds to the
// y-coordinate, and D stands for "depth" and corresponds to the z-coordinate.
type CRD struct {
	C, R, D int16
}

// CRDof returns an integer vector corresponding to v.
func CRDof(v Vector) CRD {
	x, y, z := v.Cartesian()
	return CRD{int16(x), int16(y), int16(z)}
}

// Cartesian returns the coordinates of the vector in 3D space. C,R and D are
// casted to float32.
//
// This method implements the Vector interface.
func (a CRD) Cartesian() (x, y, z float32) {
	return float32(a.C), float32(a.R), float32(a.D)
}

// CR returns the 2D integer vector corresponding to the first two dimensions of
// the vector.
func (a CRD) CR() CR {
	return CR{a.C, a.R}
}

// XY returns the 3D floating-point vector corresponding to the first two
// dimensions of the vector.
func (a CRD) XY() XY {
	return XY{float32(a.C), float32(a.R)}
}

// XYZ returns the floating-point version of the vector.
func (a CRD) XYZ() XYZ {
	return XYZ{float32(a.C), float32(a.R), float32(a.D)}
}

// XYZW returns the floating-point vector in 3D projective space, with the
// fourth dimensions set to w.
func (a CRD) XYZW(w float32) XYZW {
	return XYZW{float32(a.C), float32(a.R), float32(a.D), w}
}

////////////////////////////////////////////////////////////////////////////////

// Plus returns the sum with another vector.
func (a CRD) Plus(b CRD) CRD {
	return CRD{a.C + b.C, a.R + b.R, a.D + b.D}
}

// Pluss returns the sum with the vector (s, s, s).
func (a CRD) Pluss(s int16) CRD {
	return CRD{a.C + s, a.R + s, a.D + s}
}

// Minus returns the difference with another vector.
func (a CRD) Minus(b CRD) CRD {
	return CRD{a.C - b.C, a.R - b.R, a.D - b.D}
}

// Minuss returns the difference with the vector (s, s).
func (a CRD) Minuss(s int16) CRD {
	return CRD{a.C - s, a.R - s, a.D - s}
}

// Opposite returns the opposite of the vector.
func (a CRD) Opposite() CRD {
	return CRD{-a.C, -a.R, -a.D}
}

// Times returns the product with a scalar.
func (a CRD) Times(s int16) CRD {
	return CRD{a.C * s, a.R * s, a.D * s}
}

// Timescr returns the component-wise product with another vector.
func (a CRD) Timescr(b CRD) CRD {
	return CRD{a.C * b.C, a.R * b.R, a.D * b.D}
}

// Slash returns the integer quotient of the division by a scalar (which must be
// non-zero).
func (a CRD) Slash(s int16) CRD {
	return CRD{a.C / s, a.R / s, a.D / s}
}

// Slashcr returns the integer quotients of the component-wise division by
// another vector (of which both C and R must be non-zero).
func (a CRD) Slashcr(b CRD) CRD {
	return CRD{a.C / b.C, a.R / b.R, a.D / b.D}
}

// Mod returns the remainder (modulus) of the division by a scalar (which must
// be non-zero).
func (a CRD) Mod(s int16) CRD {
	return CRD{a.C % s, a.R % s, a.D % s}
}

// Modcr returns the remainder (modulus) of the component-wise division by
// another vector (of which both C and R must be non-zero).
func (a CRD) Modcr(b CRD) CRD {
	return CRD{a.C % b.C, a.R % b.R, a.D % b.D}
}

////////////////////////////////////////////////////////////////////////////////

// FlipC returns the vector with the sign of C flipped.
func (a CRD) FlipC() CRD {
	return CRD{-a.C, a.R, a.D}
}

// FlipR returns the vector with the sign of R flipped.
func (a CRD) FlipR() CRD {
	return CRD{a.C, -a.R, a.D}
}

// FlipD returns the vector with the sign of D flipped.
func (a CRD) FlipD() CRD {
	return CRD{a.C, a.R, -a.D}
}

// Col returns the vector projected on the C axis (i.e. with R and D nulled).
func (a CRD) Col() CRD {
	return CRD{a.C, 0, 0}
}

// Row returns the vector projected on the R axis (i.e. with C and D nulled).
func (a CRD) Row() CRD {
	return CRD{0, a.R, 0}
}

// Depth returns the vector projected on the D axis (i.e. with C and D nulled).
func (a CRD) Depth() CRD {
	return CRD{0, 0, a.D}
}
