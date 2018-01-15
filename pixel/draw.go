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
	stampPipeline.Bind()
	gl.ClearColorBuffer(screen.background)
	gl.Blending(gl.SrcAlpha, gl.OneMinusSrcAlpha)
	gl.Enable(gl.Blend)

	if true {
		if len(stamps) > 0 {
			stampSSBO.SubData(stamps, 0)
			gl.DrawInstanced(0, 4, int32(len(stamps)))
			stamps = stamps[:0]
		}
	}

	blitScreen()

	return nil
}

//------------------------------------------------------------------------------
