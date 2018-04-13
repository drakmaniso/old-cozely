// Copyright (c) 2013-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package pixel

import (
	"errors"
	"unsafe"

	"github.com/cozely/cozely/color"
	"github.com/cozely/cozely/coord"
	"github.com/cozely/cozely/input"
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
	target        coord.CR
	autozoom      bool
	origin        coord.CR // Offset when there is a border around the screen
	size          coord.CR
	pixel         int16
	commands      []gl.DrawIndirectCommand
	parameters    []int16
	cursor        TextCursor
}

var canvases []canvas

// CanvasID is the ID to handle the GPU framebuffer used to display pictures,
// print text and for various other drawing primitives.
type CanvasID uint16

////////////////////////////////////////////////////////////////////////////////

// Canvas declares a new canvas and returns its ID.
func Canvas(o ...CanvasOption) CanvasID {
	if len(canvases) >= 0xFFFF {
		setErr("in NewCanvas", errors.New("too many canvases"))
	}

	a := CanvasID(len(canvases))
	canvases = append(canvases, canvas{})

	aa := &canvases[a]
	aa.target = coord.CR{640, 360}
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

	if aa.autozoom {
		// Find best fit for pixel size
		p := win.Slashcw(aa.target)
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

	// Compute offset
	sz := aa.size.Times(aa.pixel)
	aa.origin = win.Minus(sz).Slash(2)
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
		setErr("while creating screen textures", errors.New(st.String()))
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

	blitPipeline.Bind()
	gl.DefaultFramebuffer.Bind(gl.DrawFramebuffer)
	gl.Enable(gl.FramebufferSRGB)
	gl.Disable(gl.Blend)
	gl.Viewport(int32(aa.origin.C), int32(aa.origin.R),
		int32(aa.origin.C+sz.C), int32(aa.origin.R+sz.R))
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
	return canvases[a].size
}

// PixelSize returns the size of one canvas pixel, in *window* pixels.
func (a CanvasID) PixelSize() int16 {
	return canvases[a].pixel
}

////////////////////////////////////////////////////////////////////////////////

// Mouse returns the mouse position on the canvas (in *canvas* pixels).
func (a CanvasID) Mouse() coord.CR {
	m := input.Cursor.Position()
	return m.Minus(canvases[a].origin).Slash(canvases[a].pixel)
}

////////////////////////////////////////////////////////////////////////////////
