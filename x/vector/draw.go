// Copyright (c) 2018-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package vector

import (
	"github.com/drakmaniso/cozely/colour"
	"github.com/drakmaniso/cozely/internal"
	"github.com/drakmaniso/cozely/x/gl"
)

////////////////////////////////////////////////////////////////////////////////

func init() {
	internal.VectorDraw = drawHook
}

func drawHook() error {
	internal.PaletteUpload()

	screenUniforms.PixelSize.X = 1.0 / float32(internal.Window.Width)
	screenUniforms.PixelSize.Y = 1.0 / float32(internal.Window.Height)
	screenUBO.SubData(&screenUniforms, 0)

	gl.DefaultFramebuffer.Bind(gl.DrawReadFramebuffer)
	gl.Viewport(0, 0, int32(internal.Window.Width), int32(internal.Window.Height))
	pipeline.Bind()
	gl.ClearColorBuffer(colour.LRGBA{0, 0, 0, 0})
	gl.Blending(gl.SrcAlpha, gl.OneMinusSrcAlpha)
	gl.Enable(gl.Blend)
	gl.Enable(gl.FramebufferSRGB)

	screenUBO.Bind(layoutScreen)
	commandsICBO.Bind()
	parametersTBO.Bind(layoutParameters)

	if len(commands) > 0 {
		commandsICBO.SubData(commands, 0)
		parametersTBO.SubData(parameters, 0)
		gl.DrawIndirect(0, int32(len(commands)))
		commands = commands[:0]
		parameters = parameters[:0]
	}

	return nil
}

////////////////////////////////////////////////////////////////////////////////
