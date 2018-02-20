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
	cmdLine       = 5
)

//------------------------------------------------------------------------------

func (cv Canvas) Picture(p Picture, x, y, z int16) {
	cv.appendCommand(cmdPicture, 4, 1, int16(p), x, y, z)
}

//------------------------------------------------------------------------------

func (cv Canvas) Point(c palette.Index, x, y, z int16) {
	cv.appendCommand(cmdPoint, 3, 1, int16(c), x, y, z)
}

//------------------------------------------------------------------------------

func (cv Canvas) Line(c palette.Index, x1, y1, z1, x2, y2, z2 int16) {
	cv.appendCommand(cmdLine, 4, 1, int16(c), x1, y1, z1, x2, y2, z2)
}

//------------------------------------------------------------------------------

func (cv Canvas) appendCommand(c uint32, v uint32, n uint32, params ...int16) {
	s := &canvases[cv]
	l := len(s.commands)
	if l > 0 &&
		c != cmdText &&
		(s.commands[l-1].FirstVertex>>2) == c {
		// Collapse with previous draw
		s.commands[l-1].InstanceCount += n
	} else {
		// Create new draw
		s.commands = append(s.commands, gl.DrawIndirectCommand{
			VertexCount:   v,
			InstanceCount: n,
			FirstVertex:   uint32(c << 2),
			BaseInstance:  uint32(len(s.parameters)),
		})
	}

	s.parameters = append(s.parameters, params...)
}

//------------------------------------------------------------------------------
