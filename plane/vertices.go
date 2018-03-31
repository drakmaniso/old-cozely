// Copyright (c) 2018-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package plane

//------------------------------------------------------------------------------

// Vertices represents a list of vertices.
type Vertices []Coord

// A VertexID refers to a point in a list of vertices.
type VertexID uint32

//------------------------------------------------------------------------------

// IsCCW returns true if a, b and c are in counter-clockwise order.
func (v Vertices) IsCCW(a, b, c VertexID) bool {
	// Compute the determinant of the following matrice:
	//   | xa  ya   1 |
	//   | xb  yb   1 |
	//   | xc  yc   1 |
	xa, ya := v[a].X, v[a].Y
	xb, yb := v[b].X, v[b].Y
	xc, yc := v[c].X, v[c].Y
	d := xa*yb + xb*yc + xc*ya - yb*xc - yc*xa - ya*xb
	return d > 0
}

// CCW returns a, b and c, but reordered so that they are in counter-clockwise
// order (by swapping b and c if necessary).
func (v Vertices) CCW(a, b, c VertexID) (VertexID, VertexID, VertexID) {
	if v.IsCCW(a, b, c) {
		return a, b, c
	}
	return a, c, b
}

//------------------------------------------------------------------------------

// PointInTriangle returns true if p is inside the triangle a b c.
func (v Vertices) PointInTriangle(p Coord, a, b, c VertexID) bool {
	xa, ya := v[a].X, v[a].Y
	xb, yb := v[b].X, v[b].Y
	xc, yc := v[c].X, v[c].Y
	s := ya*xc - xa*yc + (yc-ya)*p.X + (xa-xc)*p.Y
	d := xa*yb - ya*xb + (ya-yb)*p.X + (xb-xa)*p.Y

	if (s < 0) != (d < 0) {
		return false
	}

	r := -yb*xc + ya*(xc-xb) + xa*(yb-yc) + xb*yc

	if r < 0 {
		s = -s
		d = -d
		r = -r
	}
	return s > 0 && d > 0 && (s+d) <= r
}

// PointInTriangleCCW returns true if p is inside the triangle a b c (which must
// be in CCW order).
func (v Vertices) PointInTriangleCCW(p Coord, a, b, c VertexID) bool {
	xa, ya := v[a].X, v[a].Y
	xb, yb := v[b].X, v[b].Y
	xc, yc := v[c].X, v[c].Y

	// Translate to a as origin
	xbb, ybb := xb-xa, yb-ya
	xcc, ycc := xc-xa, yc-ya
	xpp, ypp := p.X-xa, p.Y-ya

	w := Coord{
		X: ycc*xpp - xcc*ypp,
		Y: -ybb*xpp + xbb*ypp,
	}
	if w.X <= 0 || w.Y <= 0 {
		return false
	}
	d := xbb*ycc - xcc*ybb
	return w.X+w.Y < d
}

//------------------------------------------------------------------------------
