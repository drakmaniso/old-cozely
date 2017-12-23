// Copyright (c) 2013-2017 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package pixel

import (
	"strings"

	"github.com/drakmaniso/carol/core/gl"
	"github.com/drakmaniso/carol/internal"
)

//------------------------------------------------------------------------------

func init() {
	internal.PixelSetup = setupHook
}

func setupHook() error {
	var err error

	createScreen()

	stampPipeline = gl.NewPipeline(
		gl.VertexShader(strings.NewReader(vertexShader)),
		gl.FragmentShader(strings.NewReader(fragmentShader)),
		gl.Topology(gl.TriangleStrip),
	)

	screenUBO = gl.NewUniformBuffer(&screenUniforms, gl.DynamicStorage|gl.MapWrite)
	screenUBO.Bind(0)

	paletteSSBO = gl.NewStorageBuffer(uintptr(256*4*4), gl.DynamicStorage|gl.MapWrite)
	paletteSSBO.Bind(2)

	stampSSBO = gl.NewStorageBuffer(uintptr(1024), gl.DynamicStorage|gl.MapWrite)
	stampSSBO.Bind(0)

	err = gl.Err()
	if err != nil {
		return err
	}

	err = loadAllPictures()
	if err != nil {
		return err
	}

	return nil
}

//------------------------------------------------------------------------------
