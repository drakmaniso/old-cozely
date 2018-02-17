// Copyright (c) 2013-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package pixel

import (
	"github.com/drakmaniso/glam/x/gl"
)

//------------------------------------------------------------------------------

const (
	cmdPicture    = 1
	cmdPictureExt = 2
	cmdText       = 3
	cmdPoint      = 4
	cmdPointList  = 5
	cmdLine       = 6
)

//------------------------------------------------------------------------------

func (cv Canvas) appendCommand(c uint32, v uint32, n uint32, params ...int16) {
	s := &canvases[cv]
	l := len(s.commands)
	if l > 0 &&
		c != cmdText &&
		c != cmdPointList &&
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
