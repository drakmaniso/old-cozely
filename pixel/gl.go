// Copyright (c) 2013-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package pixel

import (
	"errors"
	"image"
	stdcolor "image/color"
	"strings"
	"unsafe"

	"github.com/cozely/cozely/color"
	"github.com/cozely/cozely/internal"
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
	maxCommandCount = 1024
	maxParamCount   = 4 * 1024
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

	// Palette
	paletteSSBO gl.StorageBuffer

	// Command queue
	clearQueued   bool
	clearColor    color.Index
	commandsICBO  gl.IndirectBuffer
	commands      []gl.DrawIndirectCommand
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
	internal.PixelSetup = setup
	internal.PixelCleanup = cleanup
	internal.PixelRender = render
}

func setup() error {
	// Prepare the palette

	renderer.paletteSSBO = gl.NewStorageBuffer(uintptr(256*4*4), gl.DynamicStorage|gl.MapWrite)

	// Prepare the canvas

	renderer.commands = make([]gl.DrawIndirectCommand, 0, maxCommandCount)
	renderer.parameters = make([]int16, 0, maxParamCount)

	renderer.canvasBuf = gl.NewFramebuffer()
	renderer.filterBuf = gl.NewFramebuffer()

	renderer.commandsICBO = gl.NewIndirectBuffer(
		uintptr(cap(renderer.commands))*unsafe.Sizeof(renderer.commands[0]),
		gl.DynamicStorage,
	)
	renderer.parametersTBO = gl.NewBufferTexture(
		uintptr(cap(renderer.parameters))*unsafe.Sizeof(renderer.parameters[0]),
		gl.R16I,
		gl.DynamicStorage,
	)

	//TODO: create textures if not autoresize

	// Create the paint pipeline

	renderer.drawPipeline = gl.NewPipeline(
		gl.VertexShader(strings.NewReader(drawVertexShader)),
		gl.FragmentShader(strings.NewReader(drawFragmentShader)),
		gl.CullFace(false, false),
		gl.Topology(gl.TriangleStrip),
		gl.DepthTest(true),
		gl.DepthWrite(true),
		gl.DepthComparison(gl.GreaterOrEqual),
	)

	renderer.drawUBO = gl.NewUniformBuffer(&drawUniforms, gl.DynamicStorage|gl.MapWrite)

	renderer.clearQueued = true

	// Create the display pipeline

	renderer.blitPipeline = gl.NewPipeline(
		gl.VertexShader(strings.NewReader(blitVertexShader)),
		gl.FragmentShader(strings.NewReader(blitFragmentShader)),
		gl.Topology(gl.TriangleStrip),
		gl.DepthTest(false),
		gl.DepthWrite(false),
	)

	renderer.blitUBO = gl.NewUniformBuffer(&blitUniforms, gl.DynamicStorage|gl.MapWrite)

	// Create texture atlas for pictures (and fonts glyphs)

	pictures.atlas = atlas.New(1024, 1024)

	err := loadAssets()
	if err != nil {
		return err
	}

	// Mappings Buffer
	renderer.pictureMapTBO = gl.NewBufferTexture(pictures.mapping, gl.R16I, gl.StaticStorage)

	// Create the pictures texture array
	w, h := pictures.atlas.BinSize()
	if pictures.atlas.BinCount() > 0 {
		renderer.picturesTA = gl.NewTextureArray2D(1, gl.R8UI, int32(w), int32(h), int32(pictures.atlas.BinCount()))
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

		renderer.picturesTA.SubImage(0, 0, 0, int32(i), m)
	}

	pictures.path = pictures.path[:2]
	pictures.image = pictures.image[:2]

	return gl.Err()
}

////////////////////////////////////////////////////////////////////////////////

func cleanup() error {
	// Palette
	SetPalette(DefaultPalette)
	palette.dirty = true

	// Canvases
	renderer.depthTex.Delete()
	renderer.canvasTex.Delete()
	renderer.canvasBuf.Delete()
	renderer.filterTex.Delete()
	renderer.filterBuf.Delete()

	// Display pipeline
	renderer.drawPipeline.Delete()
	renderer.drawPipeline = nil
	renderer.drawUBO.Delete()

	// Pictures
	pictures.atlas = nil
	pictures.mapping = pictures.mapping[:2]
	renderer.pictureMapTBO.Delete()
	renderer.picturesTA.Delete()

	// Fonts
	fonts = fonts[:1]
	fontPaths = fontPaths[:1]

	return gl.Err()
}

////////////////////////////////////////////////////////////////////////////////

func adjustScreenTextures() {
	renderer.depthTex.Delete()
	renderer.depthTex = gl.NewRenderbuffer(gl.Depth32F, int32(screen.size.X), int32(screen.size.Y))

	renderer.canvasTex.Delete()
	renderer.canvasTex = gl.NewTexture2D(1, gl.R8UI, int32(screen.size.X), int32(screen.size.Y))
	renderer.canvasBuf.Texture(gl.ColorAttachment0, renderer.canvasTex, 0)

	renderer.canvasBuf.Renderbuffer(gl.DepthAttachment, renderer.depthTex)
	renderer.canvasBuf.DrawBuffer(gl.ColorAttachment0)
	renderer.canvasBuf.ReadBuffer(gl.NoAttachment)

	st := renderer.canvasBuf.CheckStatus(gl.DrawReadFramebuffer)
	if st != gl.FramebufferComplete {
		setErr(errors.New("pixel canvas texture creation: " + st.String()))
	}

	renderer.filterTex.Delete()
	renderer.filterTex = gl.NewTexture2D(1, gl.R8UI, int32(screen.size.X), int32(screen.size.Y))
	renderer.filterBuf.Texture(gl.ColorAttachment0, renderer.filterTex, 0)

	renderer.filterBuf.Renderbuffer(gl.DepthAttachment, renderer.depthTex)
	renderer.filterBuf.DrawBuffer(gl.ColorAttachment0)
	renderer.filterBuf.ReadBuffer(gl.NoAttachment)

	st = renderer.filterBuf.CheckStatus(gl.DrawReadFramebuffer)
	if st != gl.FramebufferComplete {
		setErr(errors.New("pixel canvas texture creation: " + st.String()))
	}
}

////////////////////////////////////////////////////////////////////////////////

// Clear sets the color of all pixels on the canvas; it also resets the filter
// of all pixels.
func (a *glRenderer) clear(c color.Index) {
	renderer.clearQueued = true
	renderer.clearColor = c
}

////////////////////////////////////////////////////////////////////////////////

func (a *glRenderer) command(c uint32, params ...int16) {
	ccap, pcap := cap(a.commands), cap(a.parameters)

	a.commands = append(a.commands, gl.DrawIndirectCommand{
		VertexCount:   4,
		InstanceCount: 1,
		FirstVertex:   0,
		BaseInstance:  uint32(c<<24 | uint32(len(a.parameters)&0xFFFFFF)),
	})
	a.parameters = append(a.parameters, params...)

	if ccap < cap(a.commands) {
		a.commandsICBO.Delete()
		a.commandsICBO = gl.NewIndirectBuffer(
			uintptr(cap(a.commands))*unsafe.Sizeof(a.commands[0]),
			gl.DynamicStorage,
		)
	}

	if pcap < cap(a.parameters) {
		a.parametersTBO.Delete()
		a.parametersTBO = gl.NewBufferTexture(
			uintptr(cap(a.parameters))*unsafe.Sizeof(a.parameters[0]),
			gl.R16I,
			gl.DynamicStorage,
		)
	}
}

////////////////////////////////////////////////////////////////////////////////

// render executes all pending commands on the canvas. It is automatically called
// by Display; the only reason to call it manually is to be able to read from it
// before display.
func render() error {
	// Upload the current palette

	if palette.dirty {
		renderer.paletteSSBO.SubData(palette.colors[:], 0)
		palette.dirty = false
	}
	renderer.paletteSSBO.Bind(0)

	// Execute all pending commands

	renderer.canvasBuf.Bind(gl.DrawFramebuffer)
	gl.Viewport(0, 0, int32(screen.size.X), int32(screen.size.Y))
	renderer.drawPipeline.Bind()
	gl.Disable(gl.Blend)

	if renderer.clearQueued {
		renderer.clearQueued = false
		renderer.canvasBuf.ClearColorUint(uint32(renderer.clearColor), 0, 0, 1)
		renderer.canvasBuf.ClearDepth(-1.0)
	}

	if len(renderer.commands) == 0 {
		goto display
	}

	drawUniforms.PixelSize.X = 1.0 / float32(screen.size.X)
	drawUniforms.PixelSize.Y = 1.0 / float32(screen.size.Y)
	drawUniforms.CanvasMargin.X = int32(screen.margin.X)
	drawUniforms.CanvasMargin.Y = int32(screen.margin.Y)
	renderer.drawUBO.SubData(&drawUniforms, 0)

	renderer.drawUBO.Bind(layoutScreen)
	renderer.commandsICBO.Bind()
	renderer.parametersTBO.Bind(layoutParameters)
	renderer.pictureMapTBO.Bind(layoutPictureMap)
	renderer.picturesTA.Bind(layoutPictures)

	renderer.commandsICBO.SubData(renderer.commands, 0)
	renderer.parametersTBO.SubData(renderer.parameters, 0)
	gl.DrawIndirect(0, int32(len(renderer.commands)))
	renderer.commands = renderer.commands[:0]
	renderer.parameters = renderer.parameters[:0]

	// Display the canvas on the game window.

display:
	blitUniforms.WindowSize.X = float32(internal.Window.Width)
	blitUniforms.WindowSize.Y = float32(internal.Window.Height)
	blitUniforms.ScreenSize.X = float32(screen.size.X)
	blitUniforms.ScreenSize.Y = float32(screen.size.Y)
	blitUniforms.ScreenZoom.X = float32(screen.zoom)
	blitUniforms.ScreenZoom.Y = float32(screen.zoom)
	renderer.blitUBO.SubData(&blitUniforms, 0)

	renderer.blitPipeline.Bind()
	gl.DefaultFramebuffer.Bind(gl.DrawFramebuffer)
	gl.Enable(gl.FramebufferSRGB)
	gl.Disable(gl.Blend)
	// gl.Viewport(int32(screen.border.X), int32(screen.border.Y),
	// 	int32(screen.border.X+sz.X), int32(screen.border.Y+sz.Y))
	gl.Viewport(0, 0, int32(internal.Window.Width), int32(internal.Window.Height))
	renderer.blitUBO.Bind(0)
	renderer.canvasTex.Bind(0)
	gl.Draw(0, 4)

	return gl.Err()
}
