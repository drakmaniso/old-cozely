// Copyright (c) 2013-2017 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package grid

import (
	"fmt"
)

//------------------------------------------------------------------------------

// A Position on the grid
type Position struct {
	x, y int8
}

var grid []uint32

var width, height int8

//------------------------------------------------------------------------------

// Setup prepare a new grid of width w and height h.
func Setup(w, h int8) {
	width, height = w, h
	grid = make([]uint32, w*h, w*h)
}

//------------------------------------------------------------------------------

func Fill(newTile func(p Position) uint32) {
	for y := int8(0); y < height; y++ {
		for x := int8(0); x < width; x++ {
			p := Position{x: x, y: y}
			e := newTile(p)
			put(e, p.x, p.y)
		}
	}
}

//------------------------------------------------------------------------------

// put entity e at grid coordinates x, y.
func put(e uint32, x, y int8) Position {
	grid[x+y*width] = e
	return Position{x: int8(x), y: int8(y)}
}

// get returns the entity at grid coordinates x, y.
func get(x, y int8) uint32 {
	return grid[x+y*width]
}

//------------------------------------------------------------------------------

// At returns the entity currently at the grid position.
func At(p Position) uint32 {
	return grid[p.x+p.y*width]
}

//------------------------------------------------------------------------------

func (p Position) String() string {
	return fmt.Sprintf("[%d,%d]", p.x, p.y)
}

//------------------------------------------------------------------------------

// Nowhere returns the nil value for grid positions.
func Nowhere() Position {
	return Position{x: -1, y: -1}
}

//------------------------------------------------------------------------------
