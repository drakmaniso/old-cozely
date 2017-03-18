// Copyright (c) 2013-2017 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package poly

//------------------------------------------------------------------------------

import (
	"strings"

	"github.com/drakmaniso/glam/gfx"
)

//------------------------------------------------------------------------------

func SetupPipeline(o ...gfx.PipelineOption) {
	o = append(o, gfx.VertexShader(strings.NewReader(vertshader)))
	pipeline = gfx.NewPipeline(o...)
}

//------------------------------------------------------------------------------

func BindPipeline() {
	pipeline.Bind()
	gfx.Enable(gfx.DepthTest)
	// gfx.CullFace(false, true)
	gfx.ClearDepthBuffer(1.0)
}

func UnbindPipeline() {
	pipeline.Unbind()
}

//------------------------------------------------------------------------------

func ClosePipeline() {
	pipeline.Close()
}

//------------------------------------------------------------------------------

var pipeline gfx.Pipeline

//------------------------------------------------------------------------------
