// Copyright (c) 2013-2017 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package noise

//------------------------------------------------------------------------------

import (
	"github.com/drakmaniso/carol/x/math32"
	"github.com/drakmaniso/carol/plane"
	"github.com/drakmaniso/carol/space"
)

//------------------------------------------------------------------------------

var Gradient3D = []space.Coord{
	space.Coord{X: +1, Y: +1, Z: 0}, space.Coord{X: -1, Y: +1, Z: 0}, space.Coord{X: +1, Y: -1, Z: 0}, space.Coord{X: -1, Y: -1, Z: 0},
	space.Coord{X: +1, Y: 0, Z: +1}, space.Coord{X: -1, Y: 0, Z: +1}, space.Coord{X: +1, Y: 0, Z: -1}, space.Coord{X: -1, Y: 0, Z: -1},
	space.Coord{X: 0, Y: +1, Z: +1}, space.Coord{X: 0, Y: -1, Z: +1}, space.Coord{X: 0, Y: +1, Z: -1}, space.Coord{X: 0, Y: -1, Z: -1},
}

//------------------------------------------------------------------------------

const cos15 = 0.9659258262890682867497431997289
const sin15 = 0.25881904510252076234889883762405
const cos30 = 0.5 * sqrt3
const sin30 = 0.5
const cos45 = 1.414213562373095 / 2.0
const sin45 = 1.414213562373095 / 2.0
const cos60 = 0.5
const sin60 = 0.5 * sqrt3
const cos75 = 0.2588190451025207623488988376240
const sin75 = 0.9659258262890682867497431997289

var Gradient4 = []plane.Coord{
	plane.Coord{1.0, 0.0}, plane.Coord{0.0, 1.0}, plane.Coord{-1.0, 0.0}, plane.Coord{0.0, -1.0},
}

var Gradient8 = []plane.Coord{
	plane.Coord{1.0, 0.0}, plane.Coord{0.0, 1.0}, plane.Coord{-1.0, 0.0}, plane.Coord{0.0, -1.0},
	plane.Coord{cos45, cos45}, plane.Coord{-cos45, cos45}, plane.Coord{-cos45, -cos45}, plane.Coord{cos45, -cos45},
}

var Gradient6 = []plane.Coord{
	plane.Coord{0.0, 1.0}, plane.Coord{sin60, cos60}, plane.Coord{-sin60, cos60},
	plane.Coord{0.0, -1.0}, plane.Coord{-sin60, -cos60}, plane.Coord{sin60, -cos60},
}

var Gradient12 = []plane.Coord{
	plane.Coord{0.0, 1.0}, plane.Coord{sin60, cos60}, plane.Coord{-sin60, cos60},
	plane.Coord{1.0, 0.0}, plane.Coord{sin30, cos30}, plane.Coord{-sin30, cos30},
	plane.Coord{0.0, -1.0}, plane.Coord{-sin60, -cos60}, plane.Coord{sin60, -cos60},
	plane.Coord{-1.0, 0.0}, plane.Coord{-sin30, -cos30}, plane.Coord{sin30, -cos30},
}

var Gradient24 = []plane.Coord{
	plane.Coord{0.0, 1.0}, plane.Coord{sin60, cos60}, plane.Coord{-sin60, cos60},
	plane.Coord{1.0, 0.0}, plane.Coord{sin30, cos30}, plane.Coord{-sin30, cos30},
	plane.Coord{0.0, -1.0}, plane.Coord{-sin60, -cos60}, plane.Coord{sin60, -cos60},
	plane.Coord{-1.0, 0.0}, plane.Coord{-sin30, -cos30}, plane.Coord{sin30, -cos30},
	plane.Coord{sin15, cos15}, plane.Coord{sin45, cos45}, plane.Coord{sin75, cos75},
	plane.Coord{-sin15, cos15}, plane.Coord{-sin45, cos45}, plane.Coord{-sin75, cos75},
	plane.Coord{sin15, -cos15}, plane.Coord{sin45, -cos45}, plane.Coord{sin75, -cos75},
	plane.Coord{-sin15, -cos15}, plane.Coord{-sin45, -cos45}, plane.Coord{-sin75, -cos75},
}

//------------------------------------------------------------------------------

func Gradient(a0 float32, n uint) []plane.Coord {
	var g = make([]plane.Coord, n, n)
	var a = 2.0 * math32.Pi / float32(n)
	for i := 0; i < int(n); i++ {
		g[i].X = math32.Cos(a0 + float32(i)*a)
		g[i].Y = math32.Sin(a0 + float32(i)*a)
	}
	return g
}

func GradientFrom(a []float32) []plane.Coord {
	var n = len(a)
	var g = make([]plane.Coord, n, n)
	for i := 0; i < int(n); i++ {
		g[i].X = math32.Cos(a[i])
		g[i].Y = math32.Sin(a[i])
	}
	return g
}

//------------------------------------------------------------------------------
