// Copyright (c) 2018-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package plane

//------------------------------------------------------------------------------

// A Triangle is defined by three indices. To find the coordinate of the
// points, the first argument of each method is a slice of Coord.
type Triangle [3]uint32

//------------------------------------------------------------------------------

// IsClockwise returns true if the three points are in clockwise order.
func (t Triangle) IsClockwise(points []Coord) bool {
	// Compute the determinant of the following matrice:
	//   | X0  Y0   1 |
	//   | X1  Y1   1 |
	//   | X2  Y2   1 |
	x0, y0 := points[t[0]].X, points[t[0]].Y
	x1, y1 := points[t[1]].X, points[t[1]].Y
	x2, y2 := points[t[2]].X, points[t[2]].Y
	d := x0*y1 + x1*y2 + x2*y0 - y1*x2 - y2*x0 - y0*x1
	return d > 0
}

//------------------------------------------------------------------------------

// CounterClockwise returns the triangle in CCW order.
func (t Triangle) CounterClockwise(points []Coord) Triangle {
	if t.IsClockwise(points) {
		t[1], t[2] = t[2], t[1]
	}
	return t
}

//------------------------------------------------------------------------------
