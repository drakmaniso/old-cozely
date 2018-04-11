// Copyright (c) 2018-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package vector

import (
	"github.com/drakmaniso/cozely/x/gl"
)

//------------------------------------------------------------------------------

const (
	cmdIndexed      = 1
	cmdIndexedExt   = 2
	cmdFullColor    = 3
	cmdFullColorExt = 4
	cmdPoint        = 5
	cmdPointList    = 6
	cmdLine         = 7
	cmdLineAA       = 8
)

//------------------------------------------------------------------------------

func appendCommand(c uint32, v uint32, n uint32) {
	l := len(commands)
	if l > 0 &&
		c != cmdPointList &&
		(commands[l-1].FirstVertex>>2) == c {
		commands[l-1].InstanceCount += n
		return
	}

	commands = append(commands, gl.DrawIndirectCommand{
		VertexCount:   v,
		InstanceCount: n,
		FirstVertex:   uint32(c << 2),
		BaseInstance:  uint32(len(parameters)),
	})
}

//------------------------------------------------------------------------------
