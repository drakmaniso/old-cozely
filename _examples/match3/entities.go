// Copyright (c) 2013-2017 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package main

import (
	"github.com/drakmaniso/carol/pixel"
	"math/rand"
	"time"
)

//------------------------------------------------------------------------------

var entities struct {
	mask    []componentMask
	gridPos []gridPos
	color   []color
}

type componentMask int32

const (
	hasGridPos componentMask = 1 << iota
	hasColor
)

type gridPos struct {
	x, y int8
}

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
			entities.mask = append(entities.mask, hasGridPos|hasColor)
			entities.gridPos = append(entities.gridPos, gridPos{x, y})
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
	ox := (pixel.ScreenSize().X - (8 * tileSize)) / 2
	oy := (pixel.ScreenSize().Y - (8 * tileSize)) / 2

	var gx, gy int8 // grid coords
	var x, y int16  // screen coords

	for e, m := range entities.mask {
		// Compute screen coords
		switch {
		case m&hasGridPos != 0:
			gx, gy = entities.gridPos[e].x, entities.gridPos[e].y
			x, y = ox+int16(gx)*tileSize, oy+int16(gy)*tileSize
		}
		// Draw
		switch {
		case m&hasColor != 0:
			c := entities.color[e]
			p := tilesPict[c].normal
			if gx == currentX && gy == currentY {
				p = tilesPict[c].big
			}
			p.Paint(x, y)
		}
	}
}

//------------------------------------------------------------------------------
