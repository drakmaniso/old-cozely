// Copyright (c) 2013-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package pixel

import (
	"github.com/cozely/cozely/color"
)

////////////////////////////////////////////////////////////////////////////////

const (
	cmdPicture    = 1
	cmdPictureExt = 2
	cmdText       = 3
	cmdPoint      = 4
	cmdLines      = 5
	cmdTriangles  = 6
	cmdBox        = 7
)

////////////////////////////////////////////////////////////////////////////////

// Paint queues a GPU command to put a picture on the canvas.
func (p PictureID) Paint(layer int16, pos XY) {
	renderer.command(cmdPicture, 4, 1, int16(p), layer, pos.X, pos.Y)
}

////////////////////////////////////////////////////////////////////////////////

// Point queues a GPU command to draw a point on the canvas.
func Point(c color.Index, layer int16, pos XY) {
	renderer.command(cmdPoint, 3, 1, int16(c), layer, pos.X, pos.Y)
}

////////////////////////////////////////////////////////////////////////////////

// Line queues a GPU command to draw a single line on the canvas.
func Line(c color.Index, layer int16, start, end XY) {
	renderer.command(
		cmdLines, 4,
		1,
		int16(c), layer,
		start.X, start.Y,
		end.X, end.Y,
	)
}

////////////////////////////////////////////////////////////////////////////////

// Triangle queues a GPU command to draw a single triangle on the canvas.
func Triangle(co color.Index, layer int16, a, b, c XY) {
	renderer.command(
		cmdTriangles, 3,
		1,
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
	renderer.command(cmdBox, 4, 1,
		int16(uint32(fg)<<8|uint32(bg)),
		layer,
		corner,
		position.X, position.Y,
		size.X, size.Y)
}
