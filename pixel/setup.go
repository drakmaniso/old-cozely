// Copyright (c) 2013-2017 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package pixel

import (
	"image"
	"image/color"
	"strings"

	"github.com/drakmaniso/glam/internal"
	"github.com/drakmaniso/glam/x/atlas"
	"github.com/drakmaniso/glam/x/gl"
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

	stampSSBO = gl.NewStorageBuffer(uintptr(256*1024), gl.DynamicStorage|gl.MapWrite)
	stampSSBO.Bind(2)

	err = gl.Err()
	if err != nil {
		return err
	}

	// Prepare picture loading

	indexedAtlas = atlas.New(1024, 1024)
	rgbaAtlas = atlas.New(1024, 1024)

	return gl.Err()
}

//------------------------------------------------------------------------------

func init() {
	internal.PixelPostSetup = postSetupHook
}

func postSetupHook() error {
	//TODO: handle the case when there is no pictures

	// Create the indexed texture atlas
	w, h := indexedAtlas.BinSize()
	indexedTexture = gl.NewTextureArray2D(1, gl.R8UI, int32(w), int32(h), int32(indexedAtlas.BinCount()))
	for i := int16(0); i < indexedAtlas.BinCount(); i++ {
		m := image.NewPaletted(image.Rectangle{
			Min: image.Point{0, 0},
			Max: image.Point{int(w), int(h)},
		},
			color.Palette{},
		)

		err := indexedAtlas.Paint(i, m)
		if err != nil {
			return err
		}

		indexedTexture.SubImage(0, 0, 0, int32(i), m)
	}
	indexedTexture.Bind(1)

	// Create the RGBA texture atlas
	w, h = rgbaAtlas.BinSize()
	rgbaTexture = gl.NewTextureArray2D(1, gl.SRGBA8, int32(w), int32(h), int32(rgbaAtlas.BinCount()))
	for i := int16(0); i < rgbaAtlas.BinCount(); i++ {
		m := image.NewNRGBA(image.Rectangle{
			Min: image.Point{0, 0},
			Max: image.Point{int(w), int(h)},
		})

		err := rgbaAtlas.Paint(i, m)
		if err != nil {
			return err
		}

		rgbaTexture.SubImage(0, 0, 0, int32(i), m)
	}
	rgbaTexture.Bind(2)

	internal.Debug.Printf("Loaded %d pictures.", len(pictures))

	mappingsTBO = gl.NewBufferTexture(mappings, gl.R16I, gl.StaticStorage)
	mappingsTBO.Bind(5)

	return gl.Err()
}

//------------------------------------------------------------------------------
