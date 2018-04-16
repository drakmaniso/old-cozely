// Copyright (c) 2013-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package main

import (
	"github.com/cozely/cozely/coord"

	"github.com/cozely/cozely/_examples/match3/ecs"
	"github.com/cozely/cozely/_examples/match3/grid"
)

////////////////////////////////////////////////////////////////////////////////

type colour int8

var colours [ecs.Size]colour

////////////////////////////////////////////////////////////////////////////////

func (loop) Render() {
	canvas.Clear(1)

	var x, y int16 // screen coords

	for e := ecs.First; e < ecs.Last(); e++ {
		// Compute screen coords
		switch {
		case e.Has(ecs.GridPosition):
			x, y = grid.PositionOf(e).ScreenXY()
		}
		// Draw
		switch {
		case e.Has(ecs.Color):
			c := colours[e]
			p := tilesPict[c].normal
			if grid.PositionOf(e) == current || e.Has(ecs.MatchFlag) {
				p = tilesPict[c].big
			}
			canvas.Picture(p, 0, coord.CR{x, y})
		}
	}

	canvas.Display()
}

////////////////////////////////////////////////////////////////////////////////
