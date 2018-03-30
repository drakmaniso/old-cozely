// Copyright (c) 2018-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package plane

//------------------------------------------------------------------------------

// A Triangle is defined by three indices. To find the coordinate of the
// points, the first argument of each method is a slice of Coord.
type Triangle [3]uint32

//------------------------------------------------------------------------------

// IsClockwise returns true if the three points are in clockwise order.
func (t Triangle) IsClockwise(coords []Coord) bool {
	// Compute the determinant of the following matrice:
	//   | X0  Y0   1 |
	//   | X1  Y1   1 |
	//   | X2  Y2   1 |
	x0, y0 := coords[t[0]].X, coords[t[0]].Y
	x1, y1 := coords[t[1]].X, coords[t[1]].Y
	x2, y2 := coords[t[2]].X, coords[t[2]].Y
	d := x0*y1 + x1*y2 + x2*y0 - y1*x2 - y2*x0 - y0*x1
	return d < 0
}

//------------------------------------------------------------------------------

// CounterClockwise returns the triangle in CCW order.
func (t Triangle) CounterClockwise(coords []Coord) Triangle {
	if t.IsClockwise(coords) {
		t[1], t[2] = t[2], t[1]
	}
	return t
}

//------------------------------------------------------------------------------

// Contains returns true if point is inside the triangle.
func (t Triangle) Contains(coords []Coord, point Coord) bool {
	x0, y0 := coords[t[0]].X, coords[t[0]].Y
	x1, y1 := coords[t[1]].X, coords[t[1]].Y
	x2, y2 := coords[t[2]].X, coords[t[2]].Y
	s := y0*x2 - x0*y2 + (y2-y0)*point.X + (x0-x2)*point.Y
	d := x0*y1 - y0*x1 + (y0-y1)*point.X + (x1-x0)*point.Y

	if (s < 0) != (d < 0) {
		return false
	}

	a := -y1*x2 + y0*(x2-x1) + x0*(y1-y2) + x1*y2

	if a < 0 {
		s = -s
		d = -d
		a = -a
	}
	return s > 0 && d > 0 && (s+d) <= a
}

//------------------------------------------------------------------------------

// ContainsCCW returns true if point is inside the triangle (wich must be CCW
// ordered).
func (t Triangle) ContainsCCW(coords []Coord, point Coord) bool {
	x0, y0 := coords[t[0]].X, coords[t[0]].Y
	x1, y1 := coords[t[1]].X, coords[t[1]].Y
	x2, y2 := coords[t[2]].X, coords[t[2]].Y
	xB, yB := x1-x0, y1-y0
	xC, yC := x2-x0, y2-y0
	xP, yP := point.X-x0, point.Y-y0
	v := Coord{
		X: yC*xP - xC*yP,
		Y: -yB*xP + xB*yP,
	}
	if v.X <= 0 || v.Y <= 0 {
		return false
	}
	d := xB*yC - xC*yB
	return v.X+v.Y < d
}

//------------------------------------------------------------------------------
