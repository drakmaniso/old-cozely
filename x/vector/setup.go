// Copyright (c) 2018-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package vector

import (
	"strings"
	"unsafe"

	"github.com/cozely/cozely/internal"
	"github.com/cozely/cozely/x/gl"
)

////////////////////////////////////////////////////////////////////////////////

func init() {
	internal.VectorSetup = setup
}

func setup() error {
	pipeline = gl.NewPipeline(
		gl.VertexShader(strings.NewReader(vertexShader)),
		gl.FragmentShader(strings.NewReader(fragmentShader)),
		gl.Topology(gl.TriangleStrip),
	)

	screenUBO = gl.NewUniformBuffer(&screenUniforms, gl.DynamicStorage|gl.MapWrite)

	commands = make([]gl.DrawIndirectCommand, 0, maxCommandCount)
	commandsICBO = gl.NewIndirectBuffer(
		uintptr(maxCommandCount)*unsafe.Sizeof(gl.DrawIndirectCommand{}),
		gl.DynamicStorage,
	)

	parameters = make([]int16, 0, maxCommandCount*maxParamCount)
	parametersTBO = gl.NewBufferTexture(
		uintptr(maxCommandCount*maxParamCount),
		gl.R16I,
		gl.DynamicStorage,
	)

	return gl.Err()
}

func cleanup() error {
	pipeline.Delete()
	screenUBO.Delete()
	commandsICBO.Delete()
	parametersTBO.Delete()
	return gl.Err()
}

////////////////////////////////////////////////////////////////////////////////
