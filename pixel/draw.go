// Copyright (c) 2013-2017 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package pixel

import (
	"github.com/drakmaniso/glam/internal"
	"github.com/drakmaniso/glam/x/gl"
)

//------------------------------------------------------------------------------

func init() {
	internal.PixelDraw = drawHook
}

func drawHook() error {
	internal.PaletteUpload()

	screenUniforms.PixelSize.X = 1.0 / float32(screen.size.X)
	screenUniforms.PixelSize.Y = 1.0 / float32(screen.size.Y)
	screenUBO.SubData(&screenUniforms, 0)

	screen.buffer.Bind(gl.DrawReadFramebuffer)
	gl.Viewport(0, 0, int32(screen.size.X), int32(screen.size.Y))
	pipeline.Bind()
	gl.ClearColorBuffer(screen.background)
	gl.Blending(gl.SrcAlpha, gl.OneMinusSrcAlpha)
	gl.Enable(gl.Blend)
	gl.Enable(gl.FramebufferSRGB)

	screenUBO.Bind(layoutScreen)
	commandsICBO.Bind()
	parametersTBO.Bind(layoutParameters)
	mappingsTBO.Bind(layoutMappings)
	indexedTextures.Bind(layoutIndexedTextures)
	fullColorTextures.Bind(layoutFullColorTextures)

	if true {
		if len(commands) > 0 {
			commandsICBO.SubData(commands, 0)
			parametersTBO.SubData(parameters, 0)
			gl.DrawIndirect(0, int32(len(commands)))
			commands = commands[:0]
			parameters = parameters[:0]
		}
	}

	blitScreen()

	return nil
}

//------------------------------------------------------------------------------
