// Copyright (c) 2013-2017 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package poly

//------------------------------------------------------------------------------

import (
	"strings"

	"github.com/drakmaniso/glam/gfx"
)

//------------------------------------------------------------------------------

func PipelineSetup() gfx.PipelineConfig {
	return func(p *gfx.Pipeline) {
		gfx.VertexShader(strings.NewReader(vertshader))(p)
		gfx.Topology(gfx.Triangles)(p)
		gfx.CullFace(false, true)(p)
		gfx.DepthTest(true)(p)
	}
}

//------------------------------------------------------------------------------
