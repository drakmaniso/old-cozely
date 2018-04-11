// Copyright (c) 2013-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package pixel

import (
	"github.com/cozely/cozely/x/gl"
)

////////////////////////////////////////////////////////////////////////////////

const (
	maxCommandCount = 1024
	maxParamCount   = 4 * 1024
)

////////////////////////////////////////////////////////////////////////////////

var screenUniforms struct {
	PixelSize struct{ X, Y float32 }
}

var blitUniforms struct {
	ScreenSize struct{ X, Y float32 }
}

////////////////////////////////////////////////////////////////////////////////

var (
	pipeline      *gl.Pipeline
	screenUBO     gl.UniformBuffer
	pictureMapTBO gl.BufferTexture
	glyphMapTBO   gl.BufferTexture
	picturesTA    gl.TextureArray2D
	glyphsTA      gl.TextureArray2D
)

var (
	blitPipeline *gl.Pipeline
	blitUBO      gl.UniformBuffer
	blitTexture  gl.Sampler
)

////////////////////////////////////////////////////////////////////////////////
