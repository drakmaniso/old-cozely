// Copyright (c) 2013-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package pixel

import (
	"github.com/drakmaniso/glam/internal"
	"github.com/drakmaniso/glam/x/gl"
)

//------------------------------------------------------------------------------

func (s *ScreenCanvas) Blit() error {
	internal.PaletteUpload()

	for c := range cursors {
		Cursor(c).Flush()
	}

	screenUniforms.PixelSize.X = 1.0 / float32(s.size.X)
	screenUniforms.PixelSize.Y = 1.0 / float32(s.size.Y)
	screenUBO.SubData(&screenUniforms, 0)

	gl.DefaultFramebuffer.Bind(gl.ReadFramebuffer) //TODO: Useless?
	s.buffer.Bind(gl.DrawFramebuffer)
	gl.Viewport(0, 0, int32(s.size.X), int32(s.size.Y))
	pipeline.Bind()
	gl.Disable(gl.Blend)
	gl.DepthMask(false)
	s.buffer.ClearColorUint(uint32(s.background), 0, 0, 0)

	screenUBO.Bind(layoutScreen)
	commandsICBO.Bind()
	parametersTBO.Bind(layoutParameters)
	pictureMapTBO.Bind(layoutPictureMap)
	glyphMapTBO.Bind(layoutGlyphMap)
	picturesTA.Bind(layoutPictures)
	glyphsTA.Bind(layoutGlyphs)

	if len(s.commands) > 0 {
		commandsICBO.SubData(s.commands, 0)
		parametersTBO.SubData(s.parameters, 0)
		gl.DrawIndirect(0, int32(len(s.commands)))
		s.commands = s.commands[:0]
		s.parameters = s.parameters[:0]
	}

	blitScreen()

	return nil
}

//------------------------------------------------------------------------------
