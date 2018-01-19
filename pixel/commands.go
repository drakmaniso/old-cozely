// Copyright (c) 2013-2017 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package pixel

import "github.com/drakmaniso/glam/x/gl"

//------------------------------------------------------------------------------

const (
	cmdIndexed      = 1
	cmdIndexedExt   = 2
	cmdFullColor    = 3
	cmdFullColorExt = 4
	cmdPoint        = 5
	cmdLine         = 6
)

//------------------------------------------------------------------------------

func commandIndexed(m uint16, x, y int16) {
	commands = append(commands, gl.DrawIndirectCommand{
		VertexCount:   4,
		InstanceCount: 1,
		FirstVertex:   uint32(cmdIndexed << 2),
		BaseInstance:  uint32(len(parameters)),
	})
	parameters = append(parameters, int16(m), x, y)
}

func commandFullColor(m uint16, x, y int16) {
	commands = append(commands, gl.DrawIndirectCommand{
		VertexCount:   4,
		InstanceCount: 1,
		FirstVertex:   uint32(cmdFullColor << 2),
		BaseInstance:  uint32(len(parameters)),
	})
	parameters = append(parameters, int16(m), x, y)
}

//------------------------------------------------------------------------------

/*
type paramIndexedPicture struct {
	mapping int16
	x       int16
	y       int16
}

type paramFullIndexedPicture struct {
	mapping           int16
	x                 int16
	y                 int16
	w                 int16
	h                 int16
	transform, shift  uint8
	alpha, brightness uint8
}

type paramRGBAPicture struct {
	mapping int16
	x       int16
	y       int16
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

type paramPoint struct {
	rg uint16
	ba uint16
	x  int16
	y  int16
}

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

const (
	cmdIndexedPoint uint32 = 1 << 2
)

//------------------------------------------------------------------------------
