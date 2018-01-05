// Copyright (c) 2013-2017 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package main

import (
	"math/rand"
	"time"

	"github.com/drakmaniso/carol/_examples/match3/grid"
)

//------------------------------------------------------------------------------

var entities struct {
	mask    []componentMask
	gridPos []grid.Position
	color   []color
}

type componentMask int32

const (
	hasGridPos componentMask = 1 << iota
	hasColor
)

type color int8

//------------------------------------------------------------------------------

const (
	gridWidth  = 8
	gridHeight = 8
)

//------------------------------------------------------------------------------

func init() {
	rand.Seed(int64(time.Now().Unix()))
}

func fillGrid() {
	for y := int8(0); y < gridHeight; y++ {
		for x := int8(0); x < gridWidth; x++ {
			e := int32(len(entities.mask))
			entities.mask = append(entities.mask, hasGridPos|hasColor)
			p := grid.Put(e, x, y)
			entities.gridPos = append(entities.gridPos, p)
			c := color(rand.Int31n(7))
			if rand.Int31n(16) == 0 {
				c = 7
			}
			entities.color = append(entities.color, c)
		}
	}
}

//------------------------------------------------------------------------------

func sysDraw() {
	var x, y int16 // screen coords

	for e, m := range entities.mask {
		// Compute screen coords
		switch {
		case m&hasGridPos != 0:
			x, y = entities.gridPos[e].ScreenXY()
		}
		// Draw
		switch {
		case m&hasColor != 0:
			c := entities.color[e]
			p := tilesPict[c].normal
			if entities.gridPos[e] == current {
				p = tilesPict[c].big
			}
			p.Paint(x, y)
		}
	}
}

//------------------------------------------------------------------------------
