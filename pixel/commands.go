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
func (a SceneID) Picture(p PictureID, pos coord.CR) {
	a.command(cmdPicture, 4, 1, int16(p), pos.C, pos.R)
}

////////////////////////////////////////////////////////////////////////////////

// Point asks the GPU to draw a point on the canvas.
func (a SceneID) Point(color color.Index, pos coord.CR) {
	a.command(cmdPoint, 3, 1, int16(color), pos.C, pos.R)
}

////////////////////////////////////////////////////////////////////////////////

// Lines asks the GPU to draw a line strip on the canvas. A line strip is a
// succession of connected lines; all lines share the same color.
func (a SceneID) Lines(c color.Index, strip ...coord.CR) {
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
func (a SceneID) Triangles(c color.Index, strip ...coord.CR) {
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
func (a SceneID) Box(fg, bg color.Index, corner int16, p1, p2 coord.CR) {
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

func (a SceneID) command(c uint32, v uint32, n uint32, params ...int16) {
	ccap, pcap := cap(scenes.commands[a]), cap(scenes.parameters[a])

	l := len(scenes.commands[a])

	switch {

	case l > 0 && c == (scenes.commands[a][l-1].BaseInstance>>24) &&
		c != cmdLines && c != cmdTriangles:

		if c != cmdText {
			// Collapse with previous draw command
			scenes.commands[a][l-1].InstanceCount += n
			scenes.parameters[a] = append(scenes.parameters[a], params...)
			break

		} else {
			// Check if same color and y
			pi := scenes.commands[a][l-1].BaseInstance & 0xFFFFFF
			if scenes.parameters[a][pi] == params[0] && scenes.parameters[a][pi+1] == params[1] {
				// Collapse with previous draw command
				scenes.commands[a][l-1].InstanceCount += n
				scenes.parameters[a] = append(scenes.parameters[a], params[2:]...)
				break
			}
		}
		// cmdText but with different params
		fallthrough

	default:
		// Create new draw command
		scenes.commands[a] = append(scenes.commands[a], gl.DrawIndirectCommand{
			VertexCount:   v,
			InstanceCount: n,
			FirstVertex:   0,
			BaseInstance:  uint32(c<<24 | uint32(len(scenes.parameters[a])&0xFFFFFF)),
		})
		scenes.parameters[a] = append(scenes.parameters[a], params...)
	}

	scenes.changed[a] = true

	if ccap < cap(scenes.commands[a]) {
		scenes.commandsICBO[a].Delete()
		scenes.commandsICBO[a] = gl.NewIndirectBuffer(
			uintptr(cap(scenes.commands[a]))*unsafe.Sizeof(scenes.commands[a][0]),
			gl.DynamicStorage,
		)
	}

	if pcap < cap(scenes.parameters[a]) {
		scenes.parametersTBO[a].Delete()
		scenes.parametersTBO[a] = gl.NewBufferTexture(
			uintptr(cap(scenes.parameters[a]))*unsafe.Sizeof(scenes.parameters[a][0]),
			gl.R16I,
			gl.DynamicStorage,
		)
	}
}

////////////////////////////////////////////////////////////////////////////////
