// Copyright (c) 2013-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package pixel

import (
	"image"
	stdcolor "image/color"
	"strings"

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
	// Create the canvases

	for i := range canvases {
		CanvasID(i).createBuffer()
	}

	// Create the paint pipeline

	pipeline = gl.NewPipeline(
		gl.VertexShader(strings.NewReader(vertexShader)),
		gl.FragmentShader(strings.NewReader(fragmentShader)),
		gl.CullFace(false, false),
		gl.Topology(gl.TriangleStrip),
		gl.DepthTest(true),
		gl.DepthWrite(true),
		gl.DepthComparison(gl.GreaterOrEqual),
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

	// Create texture atlases for pictures and fonts

	pictAtlas = atlas.New(1024, 1024)
	fntAtlas = atlas.New(128, 128)

	err := loadAssets()
	if err != nil {
		return err
	}
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
			stdcolor.Palette{},
		)

		err := pictAtlas.Paint(i, m, pictPaint)
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
			stdcolor.Palette{},
		)

		err := fntAtlas.Paint(i, m, fntPaint)
		if err != nil {
			return err
		}

		glyphsTA.SubImage(0, 0, 0, int32(i), m)

		// of, err := os.Create("testdata/fnt" + string('0'+i) + ".png")
		// if err != nil {
		// 	panic(err)
		// }
		// m.Palette = color.Palette{
		// 	color.RGBA{0, 0, 0, 255},
		// 	color.RGBA{255, 255, 255, 255},
		// 	color.RGBA{255, 0, 255, 255},
		// }
		// err = png.Encode(of, m)
		// if err != nil {
		// 	panic(err)
		// }
		// of.Close()
	}

	fntImages = fntImages[:0]

	return gl.Err()
}

func cleanup() error {
	// Canvases
	for i := range canvases {
		s := &canvases[i]
		s.texture.Delete()
		s.depth.Delete()
		s.buffer.Delete()
	}

	// Display pipeline
	pipeline.Delete()
	pipeline = nil
	screenUBO.Delete()

	// Pictures
	pictAtlas = nil
	pictureMapTBO.Delete()
	picturesTA.Delete()

	// Fonts
	fntAtlas = nil
	glyphMap = glyphMap[:0]
	glyphMapTBO.Delete()
	glyphsTA.Delete()

	return gl.Err()
}

////////////////////////////////////////////////////////////////////////////////

func init() {
	internal.PixelResize = func() {
		for i := range canvases {
			CanvasID(i).autoresize()
		}
	}
}

////////////////////////////////////////////////////////////////////////////////
