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
	// Setup pipeline for rendering on virtual screen
	createScreen()

	pipeline = gl.NewPipeline(
		gl.VertexShader(strings.NewReader(vertexShader)),
		gl.FragmentShader(strings.NewReader(fragmentShader)),
		gl.Topology(gl.TriangleStrip),
	)

	screenUBO = gl.NewUniformBuffer(&screenUniforms, gl.DynamicStorage|gl.MapWrite)

	commandsICBO = gl.NewIndirectBuffer(
		uintptr(maxCommandCount)*unsafe.Sizeof(gl.DrawIndirectCommand{}),
		gl.DynamicStorage,
	)

	parametersTBO = gl.NewBufferTexture(
		uintptr(maxCommandCount*maxParamCount),
		gl.R16I,
		gl.DynamicStorage,
	)

	// Setup pipeline for blitting screen on window

	blitPipeline = gl.NewPipeline(
		gl.VertexShader(strings.NewReader(blitVertexShader)),
		gl.FragmentShader(strings.NewReader(blitFragmentShader)),
		gl.Topology(gl.TriangleStrip),
	)

	blitUBO = gl.NewUniformBuffer(&blitUniforms, gl.DynamicStorage|gl.MapWrite)

	// Prepare picture loading

	pictAtlas = atlas.New(1024, 1024)
	fntAtlas = atlas.New(256, 256)

	return gl.Err()
}

//------------------------------------------------------------------------------

func init() {
	internal.PixelPostSetup = postSetupHook
}

func postSetupHook() error {
	//TODO: handle the case when there is no pictures

	// Mappings Buffer
	pictureMapTBO = gl.NewBufferTexture(pictureMap, gl.R16I, gl.StaticStorage)
	if len(glyphMap) > 0 {
		glyphMapTBO = gl.NewBufferTexture(glyphMap, gl.R16I, gl.StaticStorage)
	}

	// Create the pictures texture array
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

	// Create the font texture array
	w, h = fntAtlas.BinSize()
	glyphsTA = gl.NewTextureArray2D(1, gl.R8UI, int32(w), int32(h), int32(fntAtlas.BinCount()))
	for i := int16(0); i < fntAtlas.BinCount(); i++ {
		m := image.NewPaletted(image.Rectangle{
			Min: image.Point{0, 0},
			Max: image.Point{int(w), int(h)},
		},
			color.Palette{},
		)

		err := fntAtlas.Paint(i, m)
		if err != nil {
			return err
		}

		glyphsTA.SubImage(0, 0, 0, int32(i), m)

		// if i == 0 {
		// 	f, err := os.Create("FOO.png")
		// 	if err != nil {
		// 		panic(err)
		// 	}
		// 	defer f.Close()
		// 	m.Palette = []color.Color{
		// 		color.RGBA{0, 0, 0, 0xFF},
		// 		color.RGBA{0xFF, 0xFF, 0xFF, 0xFF},
		// 		color.RGBA{0, 0xFF, 0, 0xFF},
		// 		color.RGBA{0xFF, 0, 0, 0xFF},
		// 	}
		// 	err = png.Encode(f, m)
		// 	if err != nil {
		// 		panic(err)
		// 	}
		// }
	}

	return gl.Err()
}

//------------------------------------------------------------------------------
