// Copyright (c) 2013-2017 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package mtx

import (
	"strings"

	"github.com/drakmaniso/glam/color"
	"github.com/drakmaniso/glam/gfx"
	"github.com/drakmaniso/glam/internal"
)

//------------------------------------------------------------------------------

var (
	pipeline   gfx.Pipeline
	fontSSBO   gfx.StorageBuffer
	screenSSBO gfx.StorageBuffer
)

//------------------------------------------------------------------------------

func Setup() {
	pipeline = gfx.NewPipeline(
		gfx.VertexShader(strings.NewReader(vertexShader)),
		gfx.FragmentShader(strings.NewReader(fragmentShader)),
	)

	for i := range screen.chars {
		if i%120 == 0 {
			screen.chars[i] = 0
		} else if i%120 == 119 {
			screen.chars[i] = 1
		} else {
			screen.chars[i] = byte(i & 0x7F)
		}
	}
	fontSSBO = gfx.NewStorageBuffer(&Font, gfx.StaticStorage)
	screen.top = 0
	screen.left = 0 //160 * 2
	screenSSBO = gfx.NewStorageBuffer(&screen, gfx.DynamicStorage)

	internal.Log("len Font = %v\n", len(Font.Data))
}

//------------------------------------------------------------------------------

var screen struct {
	top      uint32
	left     uint32
	_        uint32
	_        uint32
	txtColor color.RGBA
	chars    [120 * 45]byte
}

//------------------------------------------------------------------------------

func SetColor(c color.RGB) {
	screen.txtColor.R = c.R
	screen.txtColor.G = c.G
	screen.txtColor.B = c.B
	//TODO: update screenSSBO
}

//------------------------------------------------------------------------------

func Draw() {
	pipeline.Bind()
	gfx.Disable(gfx.DepthTest)
	gfx.CullFace(false, false)
	fontSSBO.Bind(0)
	screenSSBO.Bind(1)
	gfx.Draw(gfx.TriangleStrip, 0, 4)
	pipeline.Unbind()
}

//------------------------------------------------------------------------------
