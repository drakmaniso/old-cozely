// Copyright (c) 2013-2017 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package grid

import (
	"github.com/drakmaniso/carol/pixel"
)

//------------------------------------------------------------------------------

// A Position on the grid
type Position struct {
	x, y int8
}

var grid []int32

var width, height int8

var (
	origin pixel.Coord
)

const cellSize = 20

//------------------------------------------------------------------------------

// Setup prepare a new grid of width w and height h.
func Setup(w, h int8) {
	width, height = w, h
	grid = make([]int32, w*h, w*h)
}

//------------------------------------------------------------------------------

// ScreenResized repositions the grid on the screen
func ScreenResized(w, h int16) {
	origin.X = (w - (8 * cellSize)) / 2
	origin.Y = (h - (8 * cellSize)) / 2
}

//------------------------------------------------------------------------------

// Put entity e at grid coordinates x, y.
func Put(e int32, x, y int8) Position {
	grid[x+y*width] = e
	return Position{x: int8(x), y: int8(y)}
}

//------------------------------------------------------------------------------

// ScreenXY returns the screen coordinates of the grid position, given a cell
// size of s.
func (p Position) ScreenXY() (x, y int16) {
	return origin.X + int16(p.x)*cellSize, origin.Y + int16(p.y)*cellSize
}

//------------------------------------------------------------------------------

// PositionAt returns the grid position containing the screen coordinates c.
func PositionAt(c pixel.Coord) Position {
	if c.X < origin.X || c.Y < origin.Y {
		return Nowhere()
	}

	x := int8((int16(c.X) - origin.X) / cellSize)
	y := int8((int16(c.Y) - origin.Y) / cellSize)

	if x >= width || y >= height {
		return Nowhere()
	}

	return Position{x: x, y: y}
}

//------------------------------------------------------------------------------

// Nowhere returns the nil value for grid positions.
func Nowhere() Position {
	return Position{x: -1, y: -1}
}

//------------------------------------------------------------------------------
