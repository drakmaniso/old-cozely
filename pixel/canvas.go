// Copyright (c) 2013-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package pixel

import (
	"errors"
	"unsafe"

	"github.com/cozely/cozely/color"
	"github.com/cozely/cozely/coord"
	"github.com/cozely/cozely/internal"
	"github.com/cozely/cozely/x/gl"
)

////////////////////////////////////////////////////////////////////////////////

type canvas struct {
	buffer        gl.Framebuffer
	texture       gl.Texture2D
	depth         gl.Renderbuffer
	commandsICBO  gl.IndirectBuffer
	parametersTBO gl.BufferTexture
	resolution    coord.CR
	fixedres      bool
	size          coord.CR // in canvas pixels
	pixel         int16    // in window pixels
	margin        coord.CR // in canvas pixels
	border        coord.CR // in window pixels (leftover from division by pixel size)
	commands      []gl.DrawIndirectCommand
	parameters    []int16
	cursor        TextCursor
}

var canvases []canvas

// CanvasID is the ID to handle the GPU framebuffer used to display pictures,
// print text and for various other drawing primitives.
type CanvasID uint16

const (
	maxCanvasID = 0xFFFF
	noCanvas = CanvasID(maxCanvasID)
)

////////////////////////////////////////////////////////////////////////////////

// Canvas declares a new canvas and returns its ID.
func Canvas(o ...CanvasOption) CanvasID {
	if internal.Running {
		setErr(errors.New("pixel canvas declaration: declarations must happen before starting the framework"))
		return noCanvas
	}

	if len(canvases) >= maxCanvasID {
		setErr(errors.New("pixel canvas declaration: too many canvases"))
		return noCanvas
	}

	a := CanvasID(len(canvases))
	canvases = append(canvases, canvas{})

	aa := &canvases[a]
	aa.resolution = coord.CR{640, 360}
	aa.pixel = 2
	aa.commands = make([]gl.DrawIndirectCommand, 0, maxCommandCount)
	aa.parameters = make([]int16, 0, maxParamCount)

	for i := range o {
		o[i](a)
	}

	//TODO: create textures if not autoresize

	return a
}

////////////////////////////////////////////////////////////////////////////////

func (a CanvasID) createBuffer() {
	aa := &canvases[a]
	aa.buffer = gl.NewFramebuffer()

	aa.commandsICBO = gl.NewIndirectBuffer(
		uintptr(cap(aa.commands))*unsafe.Sizeof(aa.commands[0]),
		gl.DynamicStorage,
	)
	aa.parametersTBO = gl.NewBufferTexture(
		uintptr(cap(aa.parameters))*unsafe.Sizeof(aa.parameters[0]),
		gl.R16I,
		gl.DynamicStorage,
	)
}

////////////////////////////////////////////////////////////////////////////////

func (a CanvasID) autoresize() {
	aa := &canvases[a]
	win := coord.CR{internal.Window.Width, internal.Window.Height}

	if aa.fixedres {
		// Find best fit for pixel size
		p := win.Slashcw(aa.resolution)
		if p.C < p.R {
			aa.pixel = p.C
		} else {
			aa.pixel = p.R
		}
		if aa.pixel < 1 {
			aa.pixel = 1
		}
	}

	// Extend the screen to cover the window
	aa.size = win.Slash(aa.pixel)
	a.createTextures()

	// For fixed resolution, compute the margin and fix the size
	if aa.fixedres {
		aa.margin = aa.size.Minus(aa.resolution).Slash(2)
	}

	// Compute outside border
	sz := aa.size.Times(aa.pixel)
	aa.border = win.Minus(sz).Slash(2)
}

////////////////////////////////////////////////////////////////////////////////

func (a CanvasID) createTextures() {
	aa := &canvases[a]

	aa.texture.Delete()
	aa.texture = gl.NewTexture2D(1, gl.R8UI, int32(aa.size.C), int32(aa.size.R))
	aa.buffer.Texture(gl.ColorAttachment0, aa.texture, 0)

	aa.depth.Delete()
	aa.depth = gl.NewRenderbuffer(gl.Depth32F, int32(aa.size.C), int32(aa.size.R))
	aa.buffer.Renderbuffer(gl.DepthAttachment, aa.depth)

	aa.buffer.DrawBuffer(gl.ColorAttachment0)
	aa.buffer.ReadBuffer(gl.NoAttachment)

	st := aa.buffer.CheckStatus(gl.DrawReadFramebuffer)
	if st != gl.FramebufferComplete {
		setErr(errors.New("pixel canvas texture creation: " + st.String()))
	}
}

////////////////////////////////////////////////////////////////////////////////

// paint executes all pending commands on the canvas. It is automatically called
// by Display; the only reason to call it manually is to be able to read from it
// before display.
func (a CanvasID) paint() {
	aa := &canvases[a]

	if len(aa.commands) == 0 {
		return
	}

	internal.PaletteUpload()

	screenUniforms.PixelSize.X = 1.0 / float32(aa.size.C)
	screenUniforms.PixelSize.Y = 1.0 / float32(aa.size.R)
	screenUniforms.CanvasMargin.X = int32(aa.margin.C)
	screenUniforms.CanvasMargin.Y = int32(aa.margin.R)
	screenUBO.SubData(&screenUniforms, 0)

	aa.buffer.Bind(gl.DrawFramebuffer)
	gl.Viewport(0, 0, int32(aa.size.C), int32(aa.size.R))
	pipeline.Bind()
	gl.Disable(gl.Blend)

	screenUBO.Bind(layoutScreen)
	aa.commandsICBO.Bind()
	aa.parametersTBO.Bind(layoutParameters)
	pictureMapTBO.Bind(layoutPictureMap)
	glyphMapTBO.Bind(layoutGlyphMap)
	picturesTA.Bind(layoutPictures)
	glyphsTA.Bind(layoutGlyphs)

	aa.commandsICBO.SubData(aa.commands, 0)
	aa.parametersTBO.SubData(aa.parameters, 0)
	gl.DrawIndirect(0, int32(len(aa.commands)))
	aa.commands = aa.commands[:0]
	aa.parameters = aa.parameters[:0]
}

////////////////////////////////////////////////////////////////////////////////

// Display tells the GPU to execute all pending drawing commands on the canvas
// (if any), and then display it on the game window.
func (a CanvasID) Display() {
	a.paint()

	aa := &canvases[a]

	sz := aa.size.Times(aa.pixel)

	blitUniforms.ScreenSize.X = float32(aa.size.C)
	blitUniforms.ScreenSize.Y = float32(aa.size.R)
	blitUBO.SubData(&blitUniforms, 0)

	internal.PaletteUpload()

	blitPipeline.Bind()
	gl.DefaultFramebuffer.Bind(gl.DrawFramebuffer)
	gl.Enable(gl.FramebufferSRGB)
	gl.Disable(gl.Blend)
	gl.Viewport(int32(aa.border.C), int32(aa.border.R),
		int32(aa.border.C+sz.C), int32(aa.border.R+sz.R))
	blitUBO.Bind(0)
	aa.texture.Bind(0)
	gl.Draw(0, 4)
}

////////////////////////////////////////////////////////////////////////////////

// Clear sets the color of all pixels on the canvas; it also resets depth
// information, used to implement layers.
func (a CanvasID) Clear(color color.Index) {
	aa := &canvases[a]
	pipeline.Bind() //TODO: find another way to enable depthWrite
	aa.buffer.ClearColorUint(uint32(color), 0, 0, 0)
	aa.buffer.ClearDepth(-1.0)
}

// ClearDepth resets depth information, used to implement layers. Call this
// method if you don't need to clear the color of the canvas, but still want to
// discard all layer information from the previous frame.
func (a CanvasID) ClearDepth() {
	aa := &canvases[a]
	pipeline.Bind() //TODO: find another way to enable depthWrite
	aa.buffer.ClearDepth(-1.0)
}

////////////////////////////////////////////////////////////////////////////////

// Size returns the current dimension of the canvas (in *canvas* pixels).
func (a CanvasID) Size() coord.CR {
	if canvases[a].fixedres {
		return canvases[a].resolution
	}
	return canvases[a].size
}

// PixelSize returns the size of one canvas pixel, in *window* pixels.
func (a CanvasID) PixelSize() int16 {
	return canvases[a].pixel
}

// FromWindow takes a coordinates in window space and returns it in canvas
// space.
func (a CanvasID) FromWindow(p coord.CR) coord.CR {
	aa := &canvases[a]
	if aa.fixedres {
		return p.Minus(aa.border).Slash(aa.pixel).Minus(aa.margin)
	}
	return p.Minus(aa.border).Slash(aa.pixel)
}

// ToWindow takes a coordinates in canvas space and returns it in window
// space.
func (a CanvasID) ToWindow(p coord.CR) coord.CR {
	aa := &canvases[a]
	if aa.fixedres {
		return p.Times(aa.pixel).Plus(aa.border)
	}
	return p.Plus(aa.margin).Times(aa.pixel).Plus(aa.border)
}

////////////////////////////////////////////////////////////////////////////////
