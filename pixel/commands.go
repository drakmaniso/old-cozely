// Copyright (c) 2013-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package pixel

import (
	"github.com/cozely/cozely/color"
)

////////////////////////////////////////////////////////////////////////////////

const (
	cmdPicture    = 1
	cmdTriangle   = 2
	cmdLine       = 3
	cmdBox        = 4
	cmdPoint      = 5
)

////////////////////////////////////////////////////////////////////////////////

// Paint queues a GPU command to put a picture on the canvas.
func (p PictureID) Paint(layer int16, pos XY) {
	renderer.command(cmdPicture, 0, layer, pos.X, pos.Y, 0, 0, int16(p), 0)
}

////////////////////////////////////////////////////////////////////////////////

// Point queues a GPU command to draw a point on the canvas.
func Point(c color.Index, layer int16, pos XY) {
	renderer.command(cmdPoint, int16(c), layer, pos.X, pos.Y, 0, 0, 0, 0)
	// Line(c, layer, pos, pos)
}

////////////////////////////////////////////////////////////////////////////////

// Line queues a GPU command to draw a single line on the canvas.
func Line(c color.Index, layer int16, start, end XY) {
	renderer.command(
		cmdLine,
		int16(c), layer,
		start.X, start.Y,
		end.X, end.Y,
		0, 0,
	)
}

////////////////////////////////////////////////////////////////////////////////

// Triangle queues a GPU command to draw a single triangle on the canvas.
func Triangle(co color.Index, layer int16, a, b, c XY) {
	renderer.command(
		cmdTriangle,
		int16(co), layer,
		a.X, a.Y,
		b.X, b.Y,
		c.X, c.Y,
	)
}

////////////////////////////////////////////////////////////////////////////////

// Box queues a GPU command to draw a box on the canvas.
func Box(fg, bg color.Index, layer int16, corner int16, position, size XY) {
	//TODO: maybe the shader can take this?
	if size.X < 0 {
		position.X = position.X + size.X
		size.X = -size.X
	}
	if size.Y < 0 {
		position.X = position.Y + size.Y
		size.Y = -size.Y
	}
	renderer.command(cmdBox,
		int16(uint32(fg)<<8|uint32(bg)),
		layer,
		position.X, position.Y,
		size.X, size.Y,
		corner,
		0,
	)
}
