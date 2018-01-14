// Copyright (c) 2013-2017 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package pixel

import (
	"fmt"
	"strings"

	"github.com/drakmaniso/carol/x/gl"
	"github.com/drakmaniso/carol/internal"
)

//------------------------------------------------------------------------------

var screenUBO gl.UniformBuffer

var screenUniforms struct {
	PixelSize struct{ X, Y float32 }
}

var mappingsTBO gl.BufferTexture

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
	gl.Enable(gl.FramebufferSRGB)

	screenUBO = gl.NewUniformBuffer(&screenUniforms, gl.DynamicStorage|gl.MapWrite)
	screenUBO.Bind(0)

	paletteSSBO = gl.NewStorageBuffer(uintptr(256*4*4), gl.DynamicStorage|gl.MapWrite)
	paletteSSBO.Bind(0)

	stampSSBO = gl.NewStorageBuffer(uintptr(256*1024), gl.DynamicStorage|gl.MapWrite)
	stampSSBO.Bind(2)

	err = gl.Err()
	if err != nil {
		return err
	}

	err = loadAllPictures()
	if err != nil {
		return err
	}

	fmt.Printf("\n\n%v\n\n", mappings)
	mappingsTBO = gl.NewBufferTexture(mappings, gl.R16I, gl.StaticStorage)
	mappingsTBO.Bind(5)

	return gl.Err()
}

//------------------------------------------------------------------------------
