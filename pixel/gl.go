// Copyright (c) 2013-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package pixel

import (
	"errors"
	"image"
	stdcolor "image/color"
	"strings"
	"unsafe"

	"github.com/cozely/cozely/internal"
	"github.com/cozely/cozely/palette"
	"github.com/cozely/cozely/x/atlas"
	"github.com/cozely/cozely/x/gl"
)

////////////////////////////////////////////////////////////////////////////////

// Uniform Buffers Bindings
const (
	layoutScreen = 0
)

// Textures Bindings
const (
	layoutParameters = 0
	layoutPictureMap = 1
	layoutPictures   = 3
)

const (
	maxParamCount = 8 * 1024
)

////////////////////////////////////////////////////////////////////////////////

var renderer = glRenderer{}

type glRenderer struct {
	// Canvas and filter drawing pipeline
	drawPipeline  *gl.Pipeline
	pictureMapTBO gl.BufferTexture
	picturesTA    gl.TextureArray2D
	drawUBO       gl.UniformBuffer

	// Blitting pipeline
	blitPipeline *gl.Pipeline
	blitUBO      gl.UniformBuffer

	// Framebuffers and their textures
	canvasBuf gl.Framebuffer
	canvasTex gl.Texture2D
	filterBuf gl.Framebuffer
	filterTex gl.Texture2D
	depthTex  gl.Renderbuffer

	// Command queue
	clearQueued   bool
	clearColor    palette.Index
	parametersTBO gl.BufferTexture
	parameters    []int16
}

// Note: The uniform structs need to be at top level to pass cgo's pointer
// check.

var drawUniforms struct {
	PixelSize    struct{ X, Y float32 }
	CanvasMargin struct{ X, Y int32 }
}

var blitUniforms struct {
	WindowSize struct{ X, Y float32 }
	ScreenSize struct{ X, Y float32 }
	ScreenZoom struct{ X, Y float32 }
}

////////////////////////////////////////////////////////////////////////////////

func init() {
	internal.PixelSetup = renderer.setup
	internal.PixelCleanup = renderer.cleanup
	internal.PixelRender = renderer.render
}

func (r *glRenderer) setup() error {
	// Prepare the canvas

	r.canvasBuf = gl.NewFramebuffer()
	r.filterBuf = gl.NewFramebuffer()

	r.parameters = make([]int16, 0, maxParamCount)
	r.parametersTBO = gl.NewBufferTexture(
		uintptr(cap(r.parameters))*unsafe.Sizeof(r.parameters[0]),
		gl.R16I,
		gl.DynamicStorage,
	)

	//TODO: create textures if not autoresize

	// Create the paint pipeline

	r.drawPipeline = gl.NewPipeline(
		gl.VertexShader(strings.NewReader(drawVertexShader)),
		gl.FragmentShader(strings.NewReader(drawFragmentShader)),
		gl.CullFace(false, false),
		gl.Topology(gl.Triangles),
		gl.DepthTest(true),
		gl.DepthWrite(true),
		gl.DepthComparison(gl.GreaterOrEqual),
	)

	r.drawUBO = gl.NewUniformBuffer(&drawUniforms, gl.DynamicStorage|gl.MapWrite)

	r.clearQueued = true

	// Create the display pipeline

	r.blitPipeline = gl.NewPipeline(
		gl.VertexShader(strings.NewReader(blitVertexShader)),
		gl.FragmentShader(strings.NewReader(blitFragmentShader)),
		gl.Topology(gl.TriangleStrip),
		gl.DepthTest(false),
		gl.DepthWrite(false),
	)

	r.blitUBO = gl.NewUniformBuffer(&blitUniforms, gl.DynamicStorage|gl.MapWrite)

	// Create texture atlas for pictures (and fonts glyphs)

	pictures.atlas = atlas.New(1024, 1024)

	err := loadAssets()
	if err != nil {
		return err
	}

	// Mappings Buffer
	r.pictureMapTBO = gl.NewBufferTexture(pictures.mapping, gl.R16I, gl.StaticStorage)

	// Create the pictures texture array
	w, h := pictures.atlas.BinSize()
	if pictures.atlas.BinCount() > 0 {
		r.picturesTA = gl.NewTextureArray2D(1, gl.R8UI, int32(w), int32(h), int32(pictures.atlas.BinCount()))
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

		r.picturesTA.SubImage(0, 0, 0, int32(i), m)
	}

	pictures.path = pictures.path[:2]
	pictures.image = pictures.image[:2]
	pictures.lut = pictures.lut[:2]

	return gl.Err()
}

////////////////////////////////////////////////////////////////////////////////

func (r *glRenderer) cleanup() error {
	// Canvases
	r.depthTex.Delete()
	r.canvasTex.Delete()
	r.canvasBuf.Delete()
	r.filterTex.Delete()
	r.filterBuf.Delete()

	// Display pipeline
	r.drawPipeline.Delete()
	r.drawPipeline = nil
	r.drawUBO.Delete()

	// Pictures
	pictures.atlas = nil
	pictures.mapping = pictures.mapping[:2]
	r.pictureMapTBO.Delete()
	r.picturesTA.Delete()

	// Fonts
	fonts = fonts[:1]
	fontPaths = fontPaths[:1]

	return gl.Err()
}

////////////////////////////////////////////////////////////////////////////////

func (r *glRenderer) adjustScreenTextures() {
	r.depthTex.Delete()
	r.depthTex = gl.NewRenderbuffer(gl.Depth32F, int32(screen.size.X), int32(screen.size.Y))

	r.canvasTex.Delete()
	r.canvasTex = gl.NewTexture2D(1, gl.R8UI, int32(screen.size.X), int32(screen.size.Y))
	r.canvasBuf.Texture(gl.ColorAttachment0, r.canvasTex, 0)

	r.canvasBuf.Renderbuffer(gl.DepthAttachment, r.depthTex)
	r.canvasBuf.DrawBuffer(gl.ColorAttachment0)
	r.canvasBuf.ReadBuffer(gl.NoAttachment)

	st := r.canvasBuf.CheckStatus(gl.DrawReadFramebuffer)
	if st != gl.FramebufferComplete {
		setErr(errors.New("pixel canvas texture creation: " + st.String()))
	}

	r.filterTex.Delete()
	r.filterTex = gl.NewTexture2D(1, gl.R8UI, int32(screen.size.X), int32(screen.size.Y))
	r.filterBuf.Texture(gl.ColorAttachment0, r.filterTex, 0)

	r.filterBuf.Renderbuffer(gl.DepthAttachment, r.depthTex)
	r.filterBuf.DrawBuffer(gl.ColorAttachment0)
	r.filterBuf.ReadBuffer(gl.NoAttachment)

	st = r.filterBuf.CheckStatus(gl.DrawReadFramebuffer)
	if st != gl.FramebufferComplete {
		setErr(errors.New("pixel canvas texture creation: " + st.String()))
	}
}

////////////////////////////////////////////////////////////////////////////////

// Clear sets the color of all pixels on the canvas; it also resets the filter
// of all pixels.
func (r *glRenderer) clear(c palette.Index) {
	r.clearQueued = true
	r.clearColor = c
}

////////////////////////////////////////////////////////////////////////////////

func (r *glRenderer) command(c int16, color, layer, x, y, p4, p5, p6, p7 int16) {
	pcap := cap(r.parameters)

	r.parameters = append(
		r.parameters,
		c<<12|(color&0xFF),
		layer,
		x, y,
		p4, p5,
		p6, p7,
	)

	if pcap < cap(r.parameters) {
		r.parametersTBO.Delete()
		r.parametersTBO = gl.NewBufferTexture(
			uintptr(cap(r.parameters))*unsafe.Sizeof(r.parameters[0]),
			gl.R16I,
			gl.DynamicStorage,
		)
	}
}

////////////////////////////////////////////////////////////////////////////////

// render executes all pending commands on the canvas. It is automatically called
// by Display; the only reason to call it manually is to be able to read from it
// before display.
func (r *glRenderer) render() error {
	// Execute all pending commands

	r.canvasBuf.Bind(gl.DrawFramebuffer)
	gl.Viewport(0, 0, int32(screen.size.X), int32(screen.size.Y))
	r.drawPipeline.Bind()
	gl.Disable(gl.Blend)

	if r.clearQueued {
		r.clearQueued = false
		r.canvasBuf.ClearColorUint(uint32(r.clearColor), 0, 0, 1)
		r.canvasBuf.ClearDepth(-1.0)
	}

	if len(r.parameters) == 0 {
		goto display
	}

	drawUniforms.PixelSize.X = 1.0 / float32(screen.size.X)
	drawUniforms.PixelSize.Y = 1.0 / float32(screen.size.Y)
	drawUniforms.CanvasMargin.X = int32(screen.margin.X)
	drawUniforms.CanvasMargin.Y = int32(screen.margin.Y)
	r.drawUBO.SubData(&drawUniforms, 0)

	r.drawUBO.Bind(layoutScreen)
	r.parametersTBO.Bind(layoutParameters)
	r.pictureMapTBO.Bind(layoutPictureMap)
	r.picturesTA.Bind(layoutPictures)

	r.parametersTBO.SubData(r.parameters, 0)
	gl.Draw(0, 6*int32(len(r.parameters)/8))
	r.parameters = r.parameters[:0]

	// Display the canvas on the game window.

display:
	blitUniforms.WindowSize.X = float32(internal.Window.Width)
	blitUniforms.WindowSize.Y = float32(internal.Window.Height)
	blitUniforms.ScreenSize.X = float32(screen.size.X)
	blitUniforms.ScreenSize.Y = float32(screen.size.Y)
	blitUniforms.ScreenZoom.X = float32(screen.zoom)
	blitUniforms.ScreenZoom.Y = float32(screen.zoom)
	r.blitUBO.SubData(&blitUniforms, 0)

	r.blitPipeline.Bind()
	gl.DefaultFramebuffer.Bind(gl.DrawFramebuffer)
	gl.Enable(gl.FramebufferSRGB)
	gl.Disable(gl.Blend)
	gl.Viewport(0, 0, int32(internal.Window.Width), int32(internal.Window.Height))
	r.blitUBO.Bind(0)
	r.canvasTex.Bind(0)
	gl.Draw(0, 4)

	return gl.Err()
}
