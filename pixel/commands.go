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
	cmdPrint      = 3
	cmdPoint      = 4
	cmdPointList  = 5
	cmdLine       = 6
)

//------------------------------------------------------------------------------

func (s *ScreenCanvas) appendCommand(c uint32, v uint32, n uint32) {
	l := len(s.commands)
	if l > 0 &&
		c != cmdPrint &&
		c != cmdPointList &&
		(s.commands[l-1].FirstVertex>>2) == c {
		s.commands[l-1].InstanceCount += n
		return
	}

	s.commands = append(s.commands, gl.DrawIndirectCommand{
		VertexCount:   v,
		InstanceCount: n,
		FirstVertex:   uint32(c << 2),
		BaseInstance:  uint32(len(s.parameters)),
	})
}

//------------------------------------------------------------------------------
