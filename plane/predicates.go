// Copyright (c) 2018-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package plane

//------------------------------------------------------------------------------

// IsCCW returns true if a, b and c are in counter-clockwise order.
func IsCCW(a, b, c Coord) bool {
	// Compute the determinant of the following matrice:
	//   | a.X  a.Y   1 |
	//   | b.X  b.Y   1 |
	//   | c.X  c.Y   1 |
	d := a.X*b.Y + b.X*c.Y + c.X*a.Y - b.Y*c.X - c.Y*a.X - a.Y*b.X
	return d > 0
}

//------------------------------------------------------------------------------

// InTriangle returns true if p is inside the triangle a b c.
func InTriangle(a, b, c Coord, p Coord) bool {
	s := a.Y*c.X - a.X*c.Y + (c.Y-a.Y)*p.X + (a.X-c.X)*p.Y
	d := a.X*b.Y - a.Y*b.X + (a.Y-b.Y)*p.X + (b.X-a.X)*p.Y

	if (s < 0) != (d < 0) {
		return false
	}

	r := -b.Y*c.X + a.Y*(c.X-b.X) + a.X*(b.Y-c.Y) + b.X*c.Y

	if r < 0 {
		s = -s
		d = -d
		r = -r
	}
	return s > 0 && d > 0 && (s+d) <= r
}

// InTriangleCCW returns true if p is inside the triangle a b c (which must
// be in counter-clockwise order).
func InTriangleCCW(a, b, c Coord, p Coord) bool {
	// Translate to a as origin
	bb := b.Minus(a)
	cc := c.Minus(a)
	pp := p.Minus(a)

	w := Coord{
		X: cc.Y*pp.X - cc.X*pp.Y,
		Y: -bb.Y*pp.X + bb.X*pp.Y,
	}
	if w.X <= 0 || w.Y <= 0 {
		return false
	}
	d := bb.X*cc.Y - cc.X*bb.Y
	return w.X+w.Y < d
}

//------------------------------------------------------------------------------

// InCircumcircle returns true if p is inside the circumcircle of triangle a b c
// (which must be in counter-clockwise order)
func InCircumcircle(a, b, c Coord, p Coord) bool {
	return ((p.Y-a.Y)*(b.X-c.X)+(p.X-a.X)*(b.Y-c.Y))*
		((p.X-c.X)*(b.X-a.X)-(p.Y-c.Y)*(b.Y-a.Y)) >
		((p.Y-c.Y)*(b.X-a.X)+(p.X-c.X)*(b.Y-a.Y))*
			((p.X-a.X)*(b.X-c.X)-(p.Y-a.Y)*(b.Y-c.Y))
}

//------------------------------------------------------------------------------
