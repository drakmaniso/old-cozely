// Copyright (c) 2013-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package pixel

import (
	"unsafe"

	"github.com/cozely/cozely/color"
	"github.com/cozely/cozely/coord"
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

// Picture asks the GPU to show a picture on the canvas.
func (a CanvasID) Picture(p PictureID, pos coord.CR) {
	a.command(cmdPicture, 4, 1, int16(p), pos.C, pos.R)
}

////////////////////////////////////////////////////////////////////////////////

// Point asks the GPU to draw a point on the canvas.
func (a CanvasID) Point(color color.Index, pos coord.CR) {
	a.command(cmdPoint, 3, 1, int16(color), pos.C, pos.R)
}

////////////////////////////////////////////////////////////////////////////////

// Lines asks the GPU to draw a line strip on the canvas. A line strip is a
// succession of connected lines; all lines share the same color.
func (a CanvasID) Lines(c color.Index, strip ...coord.CR) {
	if len(strip) < 2 {
		return
	}
	prm := []int16{int16(c)} //TODO: remove alloc
	for _, p := range strip {
		prm = append(prm, p.C, p.R)
	}
	a.command(cmdLines, 4, uint32(len(strip)-1), prm...)
}

////////////////////////////////////////////////////////////////////////////////

// Triangles asks the GPU to draw a triangle strip on the canvas. Triangle strip
// have the same meaning than in OpenGL. All triangles share the same color.
func (a CanvasID) Triangles(c color.Index, strip ...coord.CR) {
	if len(strip) < 3 {
		return
	}
	prm := []int16{int16(c)} //TODO: remove alloc
	for _, p := range strip {
		prm = append(prm, p.C, p.R)
	}
	a.command(cmdTriangles, uint32(len(strip)), 1, prm...)
}

////////////////////////////////////////////////////////////////////////////////

// Box asks the GPU to draw a box on the canvas.
func (a CanvasID) Box(fg, bg color.Index, corner int16, p1, p2 coord.CR) {
	if p2.C < p1.C {
		p1.C, p2.C = p2.C, p1.C
	}
	if p2.R < p1.R {
		p1.R, p2.R = p2.R, p1.R
	}
	a.command(cmdBox, 4, 1,
		int16(uint32(fg)<<8|uint32(bg)),
		corner,
		p1.C, p1.R,
		p2.C, p2.R)
}

////////////////////////////////////////////////////////////////////////////////

func (a CanvasID) command(c uint32, v uint32, n uint32, params ...int16) {
	s := &canvases[a]
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
			// Check if same color and y
			pi := s.commands[l-1].BaseInstance & 0xFFFFFF
			if s.parameters[pi] == params[0] && s.parameters[pi+1] == params[1] {
				// Collapse with previous draw command
				s.commands[l-1].InstanceCount += n
				s.parameters = append(s.parameters, params[2:]...)
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
