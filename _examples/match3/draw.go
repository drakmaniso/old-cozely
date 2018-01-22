// Copyright (c) 2013-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package main

import (
	"github.com/drakmaniso/glam/_examples/match3/ecs"
	"github.com/drakmaniso/glam/_examples/match3/grid"
	"github.com/drakmaniso/glam/pixel"
	"github.com/drakmaniso/glam/x/gl"
)

//------------------------------------------------------------------------------

type color int8

var colors [ecs.Size]color

//------------------------------------------------------------------------------

func (loop) Draw() error {
	s := pixel.Screen()
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
			s.Picture(p, x, y)
		}
	}

	s.Blit()

	return gl.Err()
}

//------------------------------------------------------------------------------
