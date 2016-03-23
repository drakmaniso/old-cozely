// Copyright (c) 2013-2016 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package gfx

import (
	"github.com/drakmaniso/glam/internal"
)

//------------------------------------------------------------------------------

type primitive uint32

// Drawing Primitives
const (
	Points               primitive = 0x0000
	Lines                primitive = 0x0001
	LineLoop             primitive = 0x0002
	LineStrip            primitive = 0x0003
	Triangles            primitive = 0x0004
	TriangleStrip        primitive = 0x0005
	TriangleFan          primitive = 0x0006
	LinesAdjency         primitive = 0x000A
	LineStripAdjency     primitive = 0x000B
	TrianglesAdjency     primitive = 0x000C
	TriangleStripAdjency primitive = 0x000D
	Patches              primitive = 0x000E
)

//------------------------------------------------------------------------------

func Draw(mode primitive, first int32, count int32) {
	internal.Draw(uint32(mode), first, count)
}

//------------------------------------------------------------------------------
