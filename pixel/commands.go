// Copyright (c) 2013-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package pixel

import (
	"github.com/drakmaniso/glam/palette"
	"github.com/drakmaniso/glam/x/gl"
)

//------------------------------------------------------------------------------

const (
	cmdPicture    = 1
	cmdPictureExt = 2
	cmdText       = 3
	cmdPoint      = 4
	cmdLines      = 5
	cmdTriangles  = 6
)

//------------------------------------------------------------------------------

// Picture adds a command to show a picture on the canvas.
func (cv Canvas) Picture(p Picture, x, y, z int16) {
	cv.appendCommand(cmdPicture, 4, 1, int16(p), x, y, z)
}

//------------------------------------------------------------------------------

// Point adds a command to draw a point on the canvas.
func (cv Canvas) Point(c palette.Index, x, y, z int16) {
	cv.appendCommand(cmdPoint, 3, 1, int16(c), x, y, z)
}

//------------------------------------------------------------------------------

// Lines adds a command to draw a line strip on the canvas. A line strip is a
// succesion of points connected by lines; all points and lines share the same
// depth and color.
func (cv Canvas) Lines(c palette.Index, z int16, strip ...Coord) {
	if len(strip) < 2 {
		return
	}
	prm := []int16{int16(c), z} //TODO: remove alloc
	for _, p := range strip {
		prm = append(prm, p.X, p.Y)
	}
	cv.appendCommand(cmdLines, 4, uint32(len(strip)-1), prm...)
	// cv.appendCommand(cmdLine, 4, 1, int16(c), z, x1, y1, x2, y2)
}

//------------------------------------------------------------------------------

// Triangles adds a command to draw a triangle strip on the canvas. Triangle
// strip have the same meaning than in OpenGL. All points and triangles share
// the same depth and color.
func (cv Canvas) Triangles(c palette.Index, z int16, strip ...Coord) {
	if len(strip) < 3 {
		return
	}
	prm := []int16{int16(c), z} //TODO: remove alloc
	for _, p := range strip {
		prm = append(prm, p.X, p.Y)
	}
	cv.appendCommand(cmdTriangles, uint32(len(strip)), 1, prm...)
}

//------------------------------------------------------------------------------

func (cv Canvas) appendCommand(c uint32, v uint32, n uint32, params ...int16) {
	s := &canvases[cv]
	l := len(s.commands)
	if l > 0 &&
		c != cmdText &&
		c != cmdLines &&
		c != cmdTriangles &&
		(s.commands[l-1].BaseInstance>>24) == c {
		// Collapse with previous draw
		s.commands[l-1].InstanceCount += n
	} else {
		// Create new draw
		s.commands = append(s.commands, gl.DrawIndirectCommand{
			VertexCount:   v,
			InstanceCount: n,
			FirstVertex:   0,
			BaseInstance:  uint32(c<<24 | uint32(len(s.parameters)&0xFFFFFF)),
		})
	}

	s.parameters = append(s.parameters, params...)
}

//------------------------------------------------------------------------------
