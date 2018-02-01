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

	screenUniforms.PixelSize.X = 1.0 / float32(s.size.X)
	screenUniforms.PixelSize.Y = 1.0 / float32(s.size.Y)
	screenUBO.SubData(&screenUniforms, 0)

	s.buffer.Bind(gl.DrawReadFramebuffer)
	gl.Viewport(0, 0, int32(s.size.X), int32(s.size.Y))
	pipeline.Bind()
	s.buffer.ClearColorUint(uint32(s.background), 0, 0, 0)
	gl.Disable(gl.Blend)

	screenUBO.Bind(layoutScreen)
	commandsICBO.Bind()
	parametersTBO.Bind(layoutParameters)
	picturesMapTBO.Bind(layoutMappings)
	glyphsMapTBO.Bind(layoutFontMap)
	picturesTA.Bind(layoutPictures)
	glyphsTA.Bind(layoutFonts)

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
