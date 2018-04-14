// Copyright (c) 2018-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package plane

import "github.com/cozely/cozely/coord"

////////////////////////////////////////////////////////////////////////////////

// Orientation returns a psitive value if the triangle a b c is in
// counter-clockwise order, a negative value if it is in clockwise order, and a
// null value if a b and c are colinear.
func Orientation(a, b, c coord.XY) float32 {
	A, B, C := a.XY64(), b.XY64(), c.XY64()
	// Compute the determinant of the following matrice:
	//   | a.X  a.Y   1 |
	//   | b.X  b.Y   1 |
	//   | c.X  c.Y   1 |
	d := A.X*B.Y + B.X*C.Y + C.X*A.Y - B.Y*C.X - C.Y*A.X - A.Y*B.X
	return float32(d)
}

// IsCCW returns true if a, b and c are in counter-clockwise order.
func IsCCW(a, b, c coord.XY) bool {
	A, B, C := a.XY64(), b.XY64(), c.XY64()
	// Compute the determinant of the following matrice:
	//   | a.X  a.Y   1 |
	//   | b.X  b.Y   1 |
	//   | c.X  c.Y   1 |
	d := A.X*B.Y + B.X*C.Y + C.X*A.Y - B.Y*C.X - C.Y*A.X - A.Y*B.X
	return d > 0
}

////////////////////////////////////////////////////////////////////////////////

// InTriangle returns true if p is inside the triangle a b c.
func InTriangle(a, b, c coord.XY, p coord.XY) bool {
	A, B, C, P := a.XY64(), b.XY64(), c.XY64(), p.XY64()

	s := A.Y*C.X - A.X*C.Y + (C.Y-A.Y)*P.X + (A.X-C.X)*P.Y
	d := A.X*B.Y - A.Y*B.X + (A.Y-B.Y)*P.X + (B.X-A.X)*P.Y

	if (s < 0) != (d < 0) {
		return false
	}

	r := -B.Y*C.X + A.Y*(C.X-B.X) + A.X*(B.Y-C.Y) + B.X*C.Y

	if r < 0 {
		s = -s
		d = -d
		r = -r
	}
	return s > 0 && d > 0 && (s+d) <= r
}

// InTriangleCCW returns true if p is inside the triangle a b c (which must
// be in counter-clockwise order).
func InTriangleCCW(a, b, c coord.XY, p coord.XY) bool {
	A, B, C, P := a.XY64(), b.XY64(), c.XY64(), p.XY64()

	// Translate to a as origin
	bb := B.Minus(A)
	cc := C.Minus(A)
	pp := P.Minus(A)

	w := coord.XY64{
		X: cc.Y*pp.X - cc.X*pp.Y,
		Y: -bb.Y*pp.X + bb.X*pp.Y,
	}
	if w.X <= 0 || w.Y <= 0 {
		return false
	}
	d := bb.X*cc.Y - cc.X*bb.Y
	return w.X+w.Y < d
}

////////////////////////////////////////////////////////////////////////////////

// InCircumcircle returns true if p is inside the circumcircle of triangle a b c
// (which must be in counter-clockwise order)
func InCircumcircle(a, b, c coord.XY, p coord.XY) bool {
	A, B, C, P := a.XY64(), b.XY64(), c.XY64(), p.XY64()

	return ((P.Y-A.Y)*(B.X-C.X)+(P.X-A.X)*(B.Y-C.Y))*
		((P.X-C.X)*(B.X-A.X)-(P.Y-C.Y)*(B.Y-A.Y)) >
		((P.Y-C.Y)*(B.X-A.X)+(P.X-C.X)*(B.Y-A.Y))*
			((P.X-A.X)*(B.X-C.X)-(P.Y-A.Y)*(B.Y-C.Y))
}

////////////////////////////////////////////////////////////////////////////////

// Circumcenter returns the coordinates of the circumcenter of triangle a b c.
func Circumcenter(a, b, c coord.XY) coord.XY {
	A, B, C := a.XY64(), b.XY64(), c.XY64()

	// Translate to a as origin
	ba := B.Minus(A)
	ca := C.Minus(A)

	lba := ba.Length2()
	lca := ca.Length2()

	d := 0.5 / (ba.X*ca.Y - ba.Y*ca.X) //TODO: handle div by zero case

	o := coord.XY64{
		X: (ca.Y*lba - ba.Y*lca) * d,
		Y: (ba.X*lca - ca.X*lba) * d,
	}

	return o.Plus(A).XY()
}

////////////////////////////////////////////////////////////////////////////////
