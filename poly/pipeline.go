// Copyright (c) 2013-2017 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package poly

//------------------------------------------------------------------------------

import (
	"strings"

	"github.com/drakmaniso/glam/gfx"
)

//------------------------------------------------------------------------------

func SetupPipeline(o ...gfx.PipelineConfig) {
	o = append(o, gfx.VertexShader(strings.NewReader(vertshader)))
	o = append(o, gfx.Topology(gfx.Triangles))
	o = append(o, gfx.CullFace(false, true))
	o = append(o, gfx.DepthTest(true))
	pipeline = gfx.NewPipeline(o...)
}

//------------------------------------------------------------------------------

func BindPipeline() {
	pipeline.Bind()
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

var pipeline *gfx.Pipeline

//------------------------------------------------------------------------------
