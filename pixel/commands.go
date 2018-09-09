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

// Lines queues a GPU command to draw a line strip on the canvas. A line strip
// is a succession of connected lines; all lines share the same color.
func Lines(c color.Index, layer int16, strip ...XY) {
	if len(strip) < 2 {
		return
	}
	prm := []int16{int16(c), layer} //TODO: remove alloc
	for _, p := range strip {
		prm = append(prm, p.X, p.Y)
	}
	renderer.command(cmdLines, 4, uint32(len(strip)-1), prm...)
}

////////////////////////////////////////////////////////////////////////////////

// Triangles queues a GPU command to draw a triangle strip on the canvas.
// Triangle strip have the same meaning than in OpenGL. All triangles share the
// same color.
func Triangles(c color.Index, layer int16, strip ...XY) {
	if len(strip) < 3 {
		return
	}
	prm := []int16{int16(c), layer} //TODO: remove alloc
	for _, p := range strip {
		prm = append(prm, p.X, p.Y)
	}
	renderer.command(cmdTriangles, uint32(len(strip)), 1, prm...)
}

////////////////////////////////////////////////////////////////////////////////

// Box queues a GPU command to draw a box on the canvas.
func Box(fg, bg color.Index, layer int16, corner int16, p1, p2 XY) {
	if p2.X < p1.X {
		p1.X, p2.X = p2.X, p1.X
	}
	if p2.Y < p1.Y {
		p1.Y, p2.Y = p2.Y, p1.Y
	}
	renderer.command(cmdBox, 4, 1,
		int16(uint32(fg)<<8|uint32(bg)),
		layer,
		corner,
		p1.X, p1.Y,
		p2.X, p2.Y)
}
