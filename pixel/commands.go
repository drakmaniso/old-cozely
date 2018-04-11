// Copyright (c) 2013-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package pixel

import (
	"unsafe"

	"github.com/cozely/cozely/palette"
	"github.com/cozely/cozely/plane"
	"github.com/cozely/cozely/x/gl"
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

// Picture adds a command to show a picture on the canvas.
func (cv CanvasID) Picture(p PictureID, depth int16, pos plane.Pixel) {
	cv.appendCommand(cmdPicture, 4, 1, int16(p), depth, pos.X, pos.Y)
}

////////////////////////////////////////////////////////////////////////////////

// Point adds a command to draw a point on the canvas.
func (cv CanvasID) Point(color palette.Index, depth int16, pos plane.Pixel) {
	cv.appendCommand(cmdPoint, 3, 1, int16(color), depth, pos.X, pos.Y)
}

////////////////////////////////////////////////////////////////////////////////

// Lines adds a command to draw a line strip on the canvas. A line strip is a
// succession of points connected by lines; all points and lines share the same
// depth and color.
func (cv CanvasID) Lines(c palette.Index, depth int16, strip ...plane.Pixel) {
	if len(strip) < 2 {
		return
	}
	prm := []int16{int16(c), depth} //TODO: remove alloc
	for _, p := range strip {
		prm = append(prm, p.X, p.Y)
	}
	cv.appendCommand(cmdLines, 4, uint32(len(strip)-1), prm...)
}

////////////////////////////////////////////////////////////////////////////////

// Triangles adds a command to draw a triangle strip on the canvas. Triangle
// strip have the same meaning than in OpenGL. All points and triangles share
// the same depth and color.
func (cv CanvasID) Triangles(c palette.Index, depth int16, strip ...plane.Pixel) {
	if len(strip) < 3 {
		return
	}
	prm := []int16{int16(c), depth} //TODO: remove alloc
	for _, p := range strip {
		prm = append(prm, p.X, p.Y)
	}
	cv.appendCommand(cmdTriangles, uint32(len(strip)), 1, prm...)
}

////////////////////////////////////////////////////////////////////////////////

// Box adds a command to draw a box on the canvas.
func (cv CanvasID) Box(fg, bg palette.Index, corner int16, depth int16, a, b plane.Pixel) {
	if b.X < a.X {
		a.X, b.X = b.X, a.X
	}
	if b.Y < a.Y {
		a.Y, b.Y = b.Y, a.Y
	}
	cv.appendCommand(cmdBox, 4, 1,
		int16(uint16(fg)<<8|uint16(bg)),
		corner,
		depth,
		a.X, a.Y,
		b.X, b.Y)
}

////////////////////////////////////////////////////////////////////////////////

func (cv CanvasID) appendCommand(c uint32, v uint32, n uint32, params ...int16) {
	s := &canvases[cv]
	ccap, pcap := cap(s.commands), cap(s.parameters)

	l := len(s.commands)

	switch {

	case l > 0 && c == (s.commands[l-1].BaseInstance>>24) &&
		c != cmdLines && c != cmdTriangles:

		if c != cmdText {
			// Collapse with previous draw command
			s.commands[l-1].InstanceCount += n
			s.parameters = append(s.parameters, params...)
			break

		} else {
			// Check if same color, depth and y
			pi := s.commands[l-1].BaseInstance & 0xFFFFFF
			if s.parameters[pi] == params[0] && s.parameters[pi+1] == params[1] &&
				s.parameters[pi+2] == params[2] {
				// Collapse with previous draw command
				s.commands[l-1].InstanceCount += n
				s.parameters = append(s.parameters, params[3:]...)
				break
			}
		}
		// cmdText but with different params
		fallthrough

	default:
		// Create new draw command
		s.commands = append(s.commands, gl.DrawIndirectCommand{
			VertexCount:   v,
			InstanceCount: n,
			FirstVertex:   0,
			BaseInstance:  uint32(c<<24 | uint32(len(s.parameters)&0xFFFFFF)),
		})
		s.parameters = append(s.parameters, params...)
	}

	if ccap < cap(s.commands) {
		s.commandsICBO.Delete()
		s.commandsICBO = gl.NewIndirectBuffer(
			uintptr(cap(s.commands))*unsafe.Sizeof(s.commands[0]),
			gl.DynamicStorage,
		)
	}

	if pcap < cap(s.parameters) {
		s.parametersTBO.Delete()
		s.parametersTBO = gl.NewBufferTexture(
			uintptr(cap(s.parameters))*unsafe.Sizeof(s.parameters[0]),
			gl.R16I,
			gl.DynamicStorage,
		)
	}
}

////////////////////////////////////////////////////////////////////////////////
