// Copyright (c) 2013-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package grid

import (
	"github.com/cozely/cozely/pixel"
)

////////////////////////////////////////////////////////////////////////////////

var (
	origin pixel.XY
)

const cellSize = 20

////////////////////////////////////////////////////////////////////////////////

// ScreenResized repositions the grid on the screen
func ScreenResized(w, h int16) {
	origin.X = (w - (int16(width) * cellSize)) / 2
	origin.Y = (h - (int16(height) * cellSize)) / 2
}

////////////////////////////////////////////////////////////////////////////////

// Pixel returns the screen coordinates of the grid position, given a cell
// size of s.
func (p Position) Pixel() pixel.XY {
	return pixel.XY{
		origin.X + int16(p.x)*cellSize,
		origin.Y + int16(height-1-p.y)*cellSize,
	}
}

////////////////////////////////////////////////////////////////////////////////

// PositionAt returns the grid position containing the screen coordinates c.
func PositionAt(c pixel.XY) Position {
	if c.X < origin.X || c.Y < origin.Y {
		return Nowhere()
	}

	x := int8((int16(c.X) - origin.X) / cellSize)
	y := int8((int16(c.Y) - origin.Y) / cellSize)

	if x >= width || y >= height {
		return Nowhere()
	}

	return Position{x: x, y: height - 1 - y}
}

////////////////////////////////////////////////////////////////////////////////
