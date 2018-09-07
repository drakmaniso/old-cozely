// Copyright (c) 2013-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package pixel

import (
	"unsafe"

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

// Paint queues a GPU command to put a picture on the canvas.
func Paint(p PictureID, pos XY) {
	canvas.command(cmdPicture, 4, 1, int16(p), pos.X, pos.Y)
}

////////////////////////////////////////////////////////////////////////////////

// Point queues a GPU command to draw a point on the canvas.
func Point(c Color, pos XY) {
	canvas.command(cmdPoint, 3, 1, int16(c), pos.X, pos.Y)
}

////////////////////////////////////////////////////////////////////////////////

// Lines queues a GPU command to draw a line strip on the canvas. A line strip
// is a succession of connected lines; all lines share the same color.
func Lines(c Color, strip ...XY) {
	if len(strip) < 2 {
		return
	}
	prm := []int16{int16(c)} //TODO: remove alloc
	for _, p := range strip {
		prm = append(prm, p.X, p.Y)
	}
	canvas.command(cmdLines, 4, uint32(len(strip)-1), prm...)
}

////////////////////////////////////////////////////////////////////////////////

// Triangles queues a GPU command to draw a triangle strip on the canvas.
// Triangle strip have the same meaning than in OpenGL. All triangles share the
// same color.
func Triangles(c Color, strip ...XY) {
	if len(strip) < 3 {
		return
	}
	prm := []int16{int16(c)} //TODO: remove alloc
	for _, p := range strip {
		prm = append(prm, p.X, p.Y)
	}
	canvas.command(cmdTriangles, uint32(len(strip)), 1, prm...)
}

////////////////////////////////////////////////////////////////////////////////

// Box queues a GPU command to draw a box on the canvas.
func Box(fg, bg Color, corner int16, p1, p2 XY) {
	if p2.X < p1.X {
		p1.X, p2.X = p2.X, p1.X
	}
	if p2.Y < p1.Y {
		p1.Y, p2.Y = p2.Y, p1.Y
	}
	canvas.command(cmdBox, 4, 1,
		int16(uint32(fg)<<8|uint32(bg)),
		corner,
		p1.X, p1.Y,
		p2.X, p2.Y)
}

////////////////////////////////////////////////////////////////////////////////

func (a *cmdQueue) command(c uint32, v uint32, n uint32, params ...int16) {
	ccap, pcap := cap(a.commands), cap(a.parameters)

	l := len(a.commands)

	switch {

	case l > 0 && c == (a.commands[l-1].BaseInstance>>24) &&
		c != cmdLines && c != cmdTriangles:

		if c != cmdText {
			// Collapse with previous draw command
			a.commands[l-1].InstanceCount += n
			a.parameters = append(a.parameters, params...)
			break

		} else {
			// Check if same color and y
			pi := a.commands[l-1].BaseInstance & 0xFFFFFF
			if a.parameters[pi] == params[0] && a.parameters[pi+1] == params[1] {
				// Collapse with previous draw command
				a.commands[l-1].InstanceCount += n
				a.parameters = append(a.parameters, params[2:]...)
				break
			}
		}
		// cmdText but with different params
		fallthrough

	default:
		// Create new draw command
		a.commands = append(a.commands, gl.DrawIndirectCommand{
			VertexCount:   v,
			InstanceCount: n,
			FirstVertex:   0,
			BaseInstance:  uint32(c<<24 | uint32(len(a.parameters)&0xFFFFFF)),
		})
		a.parameters = append(a.parameters, params...)
	}

	if ccap < cap(a.commands) {
		a.commandsICBO.Delete()
		a.commandsICBO = gl.NewIndirectBuffer(
			uintptr(cap(a.commands))*unsafe.Sizeof(a.commands[0]),
			gl.DynamicStorage,
		)
	}

	if pcap < cap(a.parameters) {
		a.parametersTBO.Delete()
		a.parametersTBO = gl.NewBufferTexture(
			uintptr(cap(a.parameters))*unsafe.Sizeof(a.parameters[0]),
			gl.R16I,
			gl.DynamicStorage,
		)
	}
}

////////////////////////////////////////////////////////////////////////////////
