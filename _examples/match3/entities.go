// Copyright (c) 2013-2017 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package main

import (
	"github.com/drakmaniso/carol/_examples/match3/grid"
)

//------------------------------------------------------------------------------

var entities struct {
	mask    []componentMask
	gridPos []grid.Position
	color   []color
}

type componentMask uint32

const (
	hasGridPos componentMask = 1 << iota
	hasColor
)

type color int8

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

func sysTestAndMark() {
	for e, m := range entities.mask {
		// Compute screen coords
		switch {
		case m&hasGridPos != 0 && m&hasColor != 0:
			entities.gridPos[e].TestAndMark(testMatch, markMatch)
		}
	}
}

func testMatch(e1, e2 uint32) bool {
	m1 := entities.mask[e1]
	m2 := entities.mask[e2]
	if (m1&hasColor == 0) || (m2&hasColor == 0) {
		return false
	}
	c1 := entities.color[e1]
	c2 := entities.color[e2]
	return c1 == c2
}

func markMatch(e uint32) {
	println(entities.gridPos[e].String())
}

//------------------------------------------------------------------------------
