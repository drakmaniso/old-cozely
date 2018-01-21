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
	cmdPoint      = 3
	cmdPointList  = 4
	cmdLine       = 5
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

/*
type paramFullIndexedPicture struct {
	mapping           int16
	x                 int16
	y                 int16
	w                 int16
	h                 int16
	transform, shift  uint8
	alpha, brightness uint8
}

type paramFulRGBAPicture struct {
	mapping          int16
	x                int16
	y                int16
	w                int16
	h                int16
	transform, alpha uint8
}

//------------------------------------------------------------------------------

type paramLine struct {
	rg               uint16
	ba               uint16
	x1               int16
	y1               int16
	x2               int16
	y2               int16
	width, antialias uint8
}

type paramBezier struct {
	rg               uint16
	ba               uint16
	x1               int16
	y1               int16
	x2               int16
	y2               int16
	x3               int16
	y3               int16
	x4               int16
	y4               int16
	width, antialias uint8
}

*/

//------------------------------------------------------------------------------
