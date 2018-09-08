// Copyright (c) 2013-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package pixel

import (
	"image"
	stdcolor "image/color"
	"strings"
	"unsafe"

	"github.com/cozely/cozely/internal"
	"github.com/cozely/cozely/x/atlas"
	"github.com/cozely/cozely/x/gl"
)

////////////////////////////////////////////////////////////////////////////////

func init() {
	internal.PixelSetup = setup
	internal.PixelCleanup = cleanup
}

func setup() error {
	// Prepare the palettes

	palettes.ssbo = gl.NewStorageBuffer(uintptr(256*4*4), gl.DynamicStorage|gl.MapWrite)

	for id, pp := range palettes.path {
		if pp != "" {
			PaletteID(id).load(pp)
		}
	}

	//TODO: Add default palette
	palettes.current = PaletteID(0)

	// Prepare the canvas

	canvas.commands = make([]gl.DrawIndirectCommand, 0, maxCommandCount)
	canvas.parameters = make([]int16, 0, maxParamCount)

	canvas.buffer = gl.NewFramebuffer()

	canvas.commandsICBO = gl.NewIndirectBuffer(
		uintptr(cap(canvas.commands))*unsafe.Sizeof(canvas.commands[0]),
		gl.DynamicStorage,
	)
	canvas.parametersTBO = gl.NewBufferTexture(
		uintptr(cap(canvas.parameters))*unsafe.Sizeof(canvas.parameters[0]),
		gl.R16I,
		gl.DynamicStorage,
	)

	//TODO: create textures if not autoresize

	// Create the paint pipeline

	pipeline = gl.NewPipeline(
		gl.VertexShader(strings.NewReader(vertexShader)),
		gl.FragmentShader(strings.NewReader(fragmentShader)),
		gl.CullFace(false, false),
		gl.Topology(gl.TriangleStrip),
		gl.DepthTest(false),
		gl.DepthWrite(false),
	)

	screenUBO = gl.NewUniformBuffer(&screenUniforms, gl.DynamicStorage|gl.MapWrite)

	// Create the display pipeline

	blitPipeline = gl.NewPipeline(
		gl.VertexShader(strings.NewReader(blitVertexShader)),
		gl.FragmentShader(strings.NewReader(blitFragmentShader)),
		gl.Topology(gl.TriangleStrip),
		gl.DepthTest(false),
		gl.DepthWrite(false),
	)

	blitUBO = gl.NewUniformBuffer(&blitUniforms, gl.DynamicStorage|gl.MapWrite)

	// Create texture atlas for pictures (and fonts glyphs)

	pictures.atlas = atlas.New(1024, 1024)

	err := loadAssets()
	if err != nil {
		return err
	}
	//TODO: handle the case when there is no pictures

	// Mappings Buffer
	pictureMapTBO = gl.NewBufferTexture(pictures.mapping, gl.R16I, gl.StaticStorage)

	// Create the pictures texture array
	w, h := pictures.atlas.BinSize()
	if pictures.atlas.BinCount() > 0 {
		picturesTA = gl.NewTextureArray2D(1, gl.R8UI, int32(w), int32(h), int32(pictures.atlas.BinCount()))
	}
	for i := int16(0); i < pictures.atlas.BinCount(); i++ {
		m := image.NewPaletted(image.Rectangle{
			Min: image.Point{0, 0},
			Max: image.Point{int(w), int(h)},
		},
			stdcolor.Palette{},
		)

		err := pictures.atlas.Paint(i, m, pictPaint)
		if err != nil {
			return err
		}

		// of, err := os.Create("testdata/atlas" + string('0'+i) + ".png")
		// if err != nil {
		// 	panic(err)
		// }
		// m.Palette = stdcolor.Palette{
		// 	stdcolor.RGBA{0, 0, 0, 255},
		// 	stdcolor.RGBA{255, 255, 255, 255},
		// 	stdcolor.RGBA{255, 0, 255, 255},
		// }
		// err = png.Encode(of, m)
		// if err != nil {
		// 	panic(err)
		// }
		// of.Close()

		picturesTA.SubImage(0, 0, 0, int32(i), m)
	}

	pictures.path = pictures.path[:2]
	pictures.image = pictures.image[:2]

	return gl.Err()
}

func cleanup() error {
	// Palette
	palettes.ssbo.Delete()
	palettes.path = palettes.path[:1]
	palettes.changed = palettes.changed[:1]
	palettes.stdcolors = palettes.stdcolors[:1]
	palettes.current = 0
	palettes.changed[0] = true

	// Canvases
	canvas.texture.Delete()
	canvas.buffer.Delete()

	// Display pipeline
	pipeline.Delete()
	pipeline = nil
	screenUBO.Delete()

	// Pictures
	pictures.atlas = nil
	pictures.mapping = pictures.mapping[:2]
	pictureMapTBO.Delete()
	picturesTA.Delete()

	// Fonts
	fonts = fonts[:1]
	fontPaths = fontPaths[:1]

	return gl.Err()
}

////////////////////////////////////////////////////////////////////////////////
