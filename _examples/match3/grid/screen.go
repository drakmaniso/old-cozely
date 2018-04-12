// Copyright (c) 2013-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package grid

import "github.com/cozely/cozely/coord"

////////////////////////////////////////////////////////////////////////////////

var (
	origin coord.CR
)

const cellSize = 20

////////////////////////////////////////////////////////////////////////////////

// ScreenResized repositions the grid on the screen
func ScreenResized(w, h int16) {
	origin.C = (w - (int16(width) * cellSize)) / 2
	origin.R = (h - (int16(height) * cellSize)) / 2
}

////////////////////////////////////////////////////////////////////////////////

// ScreenXY returns the screen coordinates of the grid position, given a cell
// size of s.
func (p Position) ScreenXY() (x, y int16) {
	x = origin.C + int16(p.x)*cellSize
	y = origin.R + int16(height-1-p.y)*cellSize
	return x, y
}

////////////////////////////////////////////////////////////////////////////////

// PositionAt returns the grid position containing the screen coordinates c.
func PositionAt(c coord.CR) Position {
	if c.C < origin.C || c.R < origin.R {
		return Nowhere()
	}

	x := int8((int16(c.C) - origin.C) / cellSize)
	y := int8((int16(c.R) - origin.R) / cellSize)

	if x >= width || y >= height {
		return Nowhere()
	}

	return Position{x: x, y: height - 1 - y}
}

////////////////////////////////////////////////////////////////////////////////
