// Copyright (c) 2013-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package pixel

import (
	"errors"

	"github.com/cozely/cozely/internal"
	"github.com/cozely/cozely/window"
	"github.com/cozely/cozely/x/gl"
)

////////////////////////////////////////////////////////////////////////////////

var canvas = struct {
	resolution XY    // fixed resolution, or {0, 0} for fixed zoom
	zoom       int16 // in window pixels

	size   XY        // size of the canvas
	margin XY        // for fixed resolution only, = size - resolution
	border window.XY // leftover from division by pixel size

	buffer  gl.Framebuffer
	texture gl.Texture2D
	filter  gl.Texture2D

	cmdQueue
}{
	resolution: XY{640, 360},
	zoom:       2,
}

type cmdQueue struct {
	commands      []gl.DrawIndirectCommand
	parameters    []int16
	commandsICBO  gl.IndirectBuffer
	parametersTBO gl.BufferTexture
}

////////////////////////////////////////////////////////////////////////////////

// SetResolution defines a target resolution for the automatic resizing of
// the canvas.
//
// It guarantees that:
// - the canvas will never be smaller than the target resolution,
// - the target resolution will occupy as much screen as possible.
func SetResolution(w, h int16) {
	if internal.Running {
		setErr(errors.New("Resolution must be called before starting the framework"))
		return
	}
	canvas.resolution = XY{w, h}
	// if internal.Running {
	// 	CanvasID(0).autoresize()
	// }
}

// SetZoom sets the pixel size used to display the canvas.
func SetZoom(z int16) {
	if internal.Running {
		setErr(errors.New("Resolution must be called before starting the framework"))
		return
	}
	if z < 1 {
		z = 1
	}
	canvas.zoom = z
	canvas.resolution = XY{}
	canvas.margin = XY{}
	// if internal.Running {
	// 	CanvasID(0).autoresize()
	// }
}

////////////////////////////////////////////////////////////////////////////////

func init() {
	internal.PixelResize = resize
}

func resize() {
	//TODO: use window.XY
	win := window.XY{internal.Window.Width, internal.Window.Height}

	if !canvas.resolution.Null() {
		// Find best fit for pixel size
		p := win.SlashXY(window.XYof(canvas.resolution))
		if p.X < p.Y {
			canvas.zoom = p.X
		} else {
			canvas.zoom = p.Y
		}
		if canvas.zoom < 1 {
			canvas.zoom = 1
		}
	}

	// Extend the screen to cover the window
	canvas.size = XYof(win.Slash(canvas.zoom))
	createCanvasTextures()

	// For fixed resolution, compute the margin and fix the size
	if !canvas.resolution.Null() {
		canvas.margin = canvas.size.Minus(canvas.resolution).Slash(2)
	}

	// Compute outside border
	sz := window.XYof(canvas.size.Times(canvas.zoom))
	canvas.border = win.Minus(sz).Slash(2)
}

////////////////////////////////////////////////////////////////////////////////

func createCanvasTextures() {
	canvas.texture.Delete()
	canvas.texture = gl.NewTexture2D(1, gl.R8, int32(canvas.size.X), int32(canvas.size.Y))
	canvas.buffer.Texture(gl.ColorAttachment0, canvas.texture, 0)

	canvas.filter.Delete()
	canvas.filter = gl.NewTexture2D(1, gl.R8, int32(canvas.size.X), int32(canvas.size.Y))
	canvas.buffer.Texture(gl.ColorAttachment1, canvas.filter, 0)

	canvas.buffer.DrawBuffers([]gl.FramebufferAttachment{gl.ColorAttachment0, gl.ColorAttachment1})
	canvas.buffer.ReadBuffer(gl.NoAttachment)

	st := canvas.buffer.CheckStatus(gl.DrawReadFramebuffer)
	if st != gl.FramebufferComplete {
		setErr(errors.New("pixel canvas texture creation: " + st.String()))
	}
}

////////////////////////////////////////////////////////////////////////////////

func init() {
	internal.PixelRender = render
}

// render executes all pending commands on the canvas. It is automatically called
// by Display; the only reason to call it manually is to be able to read from it
// before display.
func render() error {
	// Upload the current palette

	p := palettes.current
	if palettes.changed[p] {
		palettes.ssbo.SubData(palettes.stdcolors[p][:], 0)
		palettes.changed[p] = false
	}
	palettes.ssbo.Bind(0)

	// Execute all pending commands

	if len(canvas.commands) == 0 {
		goto display
	}

	screenUniforms.PixelSize.X = 1.0 / float32(canvas.size.X)
	screenUniforms.PixelSize.Y = 1.0 / float32(canvas.size.Y)
	screenUniforms.CanvasMargin.X = int32(canvas.margin.X)
	screenUniforms.CanvasMargin.Y = int32(canvas.margin.Y)
	screenUBO.SubData(&screenUniforms, 0)

	canvas.buffer.Bind(gl.DrawFramebuffer)
	gl.Viewport(0, 0, int32(canvas.size.X), int32(canvas.size.Y))
	pipeline.Bind()
	gl.Enable(gl.Blend)
	gl.Blending(gl.SrcAlpha, gl.OneMinusSrcAlpha)

	screenUBO.Bind(layoutScreen)
	canvas.commandsICBO.Bind()
	canvas.parametersTBO.Bind(layoutParameters)
	pictureMapTBO.Bind(layoutPictureMap)
	picturesTA.Bind(layoutPictures)

	canvas.commandsICBO.SubData(canvas.commands, 0)
	canvas.parametersTBO.SubData(canvas.parameters, 0)
	gl.DrawIndirect(0, int32(len(canvas.commands)))
	canvas.commands = canvas.commands[:0]
	canvas.parameters = canvas.parameters[:0]

	// Display the canvas on the game window.

display:
	sz := canvas.size.Times(canvas.zoom)

	blitUniforms.ScreenSize.X = float32(canvas.size.X)
	blitUniforms.ScreenSize.Y = float32(canvas.size.Y)
	blitUBO.SubData(&blitUniforms, 0)

	blitPipeline.Bind()
	gl.DefaultFramebuffer.Bind(gl.DrawFramebuffer)
	gl.Enable(gl.FramebufferSRGB)
	gl.Disable(gl.Blend)
	gl.Viewport(int32(canvas.border.X), int32(canvas.border.Y),
		int32(canvas.border.X+sz.X), int32(canvas.border.Y+sz.Y))
	blitUBO.Bind(0)
	canvas.texture.Bind(0)
	canvas.filter.Bind(1)
	gl.Draw(0, 4)

	return gl.Err()
}

////////////////////////////////////////////////////////////////////////////////

// Clear sets the color of all pixels on the canvas; it also resets the filter
// of all pixels.
func Clear(color Color) {
	pipeline.Bind() //TODO: find another way to enable depthWrite
	canvas.buffer.ClearColor(0, float32(color)/255, 0, 0, 0)
	canvas.buffer.ClearColor(1, 0, 0, 0, 0)
}

////////////////////////////////////////////////////////////////////////////////

// Resolution returns the current dimension of the canvas (in *canvas* pixels).
func Resolution() XY {
	if !canvas.resolution.Null() {
		return canvas.resolution
	}
	return canvas.size
}

// Zoom returns the size of one canvas pixel, in *window* pixels.
func Zoom() int16 {
	return canvas.zoom
}

// ToCanvas takes coordinates in window space and returns them in canvas
// space.
func ToCanvas(p window.XY) XY {
	//TODO:
	if !canvas.resolution.Null() {
		return XYof(p.Minus(canvas.border).Slash(canvas.zoom)).Minus(canvas.margin)
	}
	return XYof(p.Minus(canvas.border).Slash(canvas.zoom))
}

// ToWindow takes coordinates in canvas space and returns them in window space.
func ToWindow(p XY) window.XY {
	//TODO:
	if !canvas.resolution.Null() {
		return window.XYof(p.Times(canvas.zoom)).Plus(canvas.border)
	}
	return window.XY(p.Plus(canvas.margin).Times(canvas.zoom)).Plus(canvas.border)
}

////////////////////////////////////////////////////////////////////////////////
