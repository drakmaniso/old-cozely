// Copyright (c) 2013-2016 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package gfx

import (
	"io"

	"github.com/drakmaniso/glam/internal"
)

//------------------------------------------------------------------------------

// A Shader is a compiled program run by the GPU.
type Shader struct {
	internal internal.Shader
}

//------------------------------------------------------------------------------

func (s *Shader) Create(st shaderStages, r io.Reader) {
	s.internal.Create(uint32(st), r)
}

//------------------------------------------------------------------------------

type shaderStages uint32

const (
	VertexShader         shaderStages = 0x8B31
	FragmentShader       shaderStages = 0x8B30
	GeometryShader       shaderStages = 0x8DD9
	TessControlShader    shaderStages = 0x8E88
	TessEvaluationShader shaderStages = 0x8E87
	ComputeShader        shaderStages = 0x91B9
)

//------------------------------------------------------------------------------
