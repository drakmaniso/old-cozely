// Copyright (c) 2013-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package main

import (
	"github.com/cozely/cozely/examples/match3/ecs"
	"github.com/cozely/cozely/examples/match3/grid"
	"github.com/cozely/cozely/pixel"
)

////////////////////////////////////////////////////////////////////////////////

type colour int8

var colours [ecs.Size]colour

////////////////////////////////////////////////////////////////////////////////

func (loop) Render() {
	pixel.Clear(1)

	var pos pixel.XY

	for e := ecs.First; e < ecs.Last(); e++ {
		// Compute screen coords
		switch {
		case e.Has(ecs.GridPosition):
			pos = grid.PositionOf(e).Pixel()
		}
		// Draw
		switch {
		case e.Has(ecs.Color):
			c := colours[e]
			p := tilesPict[c].normal
			if grid.PositionOf(e) == current || e.Has(ecs.MatchFlag) {
				p = tilesPict[c].big
			}
			p.Paint(pos, 0, 0)
		}
	}
}

////////////////////////////////////////////////////////////////////////////////
