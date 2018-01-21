// Copyright (c) 2013-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package pixel

import (
	"image"
	"image/color"
	"strings"
	"unsafe"

	"github.com/drakmaniso/glam/internal"
	"github.com/drakmaniso/glam/x/atlas"
	"github.com/drakmaniso/glam/x/gl"
)

//------------------------------------------------------------------------------

func init() {
	internal.PixelSetup = setupHook
}

func setupHook() error {
	createScreen()

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

	// Prepare picture loading

	pictAtlas = atlas.New(1024, 1024)

	return gl.Err()
}

//------------------------------------------------------------------------------

func init() {
	internal.PixelPostSetup = postSetupHook
}

func postSetupHook() error {
	//TODO: handle the case when there is no pictures

	// Mappings Buffer
	mappingsTBO = gl.NewBufferTexture(mappings, gl.R16I, gl.StaticStorage)

	// Create the texture atlas
	w, h := pictAtlas.BinSize()
	picturesTA = gl.NewTextureArray2D(1, gl.R8UI, int32(w), int32(h), int32(pictAtlas.BinCount()))
	for i := int16(0); i < pictAtlas.BinCount(); i++ {
		m := image.NewPaletted(image.Rectangle{
			Min: image.Point{0, 0},
			Max: image.Point{int(w), int(h)},
		},
			color.Palette{},
		)

		err := pictAtlas.Paint(i, m)
		if err != nil {
			return err
		}

		picturesTA.SubImage(0, 0, 0, int32(i), m)
	}

	return gl.Err()
}

//------------------------------------------------------------------------------
