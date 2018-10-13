// Copyright (c) 2013-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package pixel

import "github.com/cozely/cozely/color"

////////////////////////////////////////////////////////////////////////////////

const (
	cmdPicture    = 1
	cmdTile       = 2
	cmdSubpicture = 3
	cmdTriangle   = 4
	cmdLine       = 5
	cmdPoint      = 6
	cmdBox        = 7
)

////////////////////////////////////////////////////////////////////////////////

// Paint queues a GPU command to put a picture on the canvas.
func (p PictureID) Paint(pos XY, z Layer) {
	renderer.command(cmdPicture, 0, int16(z), pos.X, pos.Y, 0, 0, int16(p), 0)
}

// Tile
func (p PictureID) Tile(pos XY, size XY, z Layer) {
	if size.X < 0 {
		pos.X += size.X
		size.X = -size.X
	}
	if size.Y < 0 {
		pos.Y += size.Y
		size.Y = -size.Y
	}
	renderer.command(cmdTile, 0, int16(z), pos.X, pos.Y, size.X, size.Y, int16(p), 0)
}

////////////////////////////////////////////////////////////////////////////////

// Point queues a GPU command to draw a point on the canvas.
func Point(pos XY, z Layer, c color.Index) {
	renderer.command(cmdPoint, int16(c), int16(z), pos.X, pos.Y, 0, 0, 0, 0)
	// Line(c, layer, pos, pos)
}

////////////////////////////////////////////////////////////////////////////////

// Line queues a GPU command to draw a single line on the canvas.
func Line(p1, p2 XY, z Layer, c color.Index) {
	renderer.command(
		cmdLine,
		int16(c), int16(z),
		p1.X, p1.Y,
		p2.X, p2.Y,
		0, 0,
	)
}

////////////////////////////////////////////////////////////////////////////////

// Triangle queues a GPU command to draw a single triangle on the canvas.
func Triangle(p1, p2, p3 XY, z Layer, co color.Index) {
	renderer.command(
		cmdTriangle,
		int16(co), int16(z),
		p1.X, p1.Y,
		p2.X, p2.Y,
		p3.X, p3.Y,
	)
}

////////////////////////////////////////////////////////////////////////////////

// Box queues a GPU command to draw a box on the canvas.
func Box(position, size XY, z Layer, corner int16, fg, bg color.Index) {
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
		int16(bg),
		int16(z),
		position.X, position.Y,
		size.X, size.Y,
		corner,
		int16(fg),
	)
}
