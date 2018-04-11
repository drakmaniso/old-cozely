// Copyright (c) 2013-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package main

import (
	"github.com/drakmaniso/cozely/_examples/match3/ecs"
	"github.com/drakmaniso/cozely/_examples/match3/grid"
	"github.com/drakmaniso/cozely/plane"
	"github.com/drakmaniso/cozely/x/gl"
)

//------------------------------------------------------------------------------

type color int8

var colors [ecs.Size]color

//------------------------------------------------------------------------------

func (loop) Draw() error {
	screen.Clear(1)

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
			c := colors[e]
			p := tilesPict[c].normal
			if grid.PositionOf(e) == current || e.Has(ecs.MatchFlag) {
				p = tilesPict[c].big
			}
			screen.Picture(p, 0, plane.Pixel{x, y})
		}
	}

	screen.Display()

	return gl.Err()
}

//------------------------------------------------------------------------------
