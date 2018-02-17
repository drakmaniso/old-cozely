// Copyright (c) 2013-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package pixel

import (
	"errors"

	"github.com/drakmaniso/glam/internal"
	"github.com/drakmaniso/glam/mouse"
	"github.com/drakmaniso/glam/palette"
	"github.com/drakmaniso/glam/x/gl"
)

//------------------------------------------------------------------------------

type canvas struct {
	buffer     gl.Framebuffer
	texture    gl.Texture2D
	depth      gl.Renderbuffer
	target     Coord
	autozoom   bool
	size       Coord
	pixel      int32
	ox, oy     int32 // Offset when there is a border around the screen
	background palette.Index
	commands   []gl.DrawIndirectCommand
	parameters []int16
}

var canvases []canvas

type Canvas uint16

const Screen = Canvas(0)

func init() {
	canvases = append(canvases, canvas{
		target: Coord{X: 640, Y: 360},
		pixel:  2,
	})
}

//------------------------------------------------------------------------------

func (cv Canvas) Size() Coord {
	return canvases[cv].size
}

func (cv Canvas) Pixel() int32 {
	return canvases[cv].pixel
}

//------------------------------------------------------------------------------

func (cv Canvas) SetBackground(c palette.Index) {
	canvases[cv].background = c
}

//------------------------------------------------------------------------------

func createScreen() {
	s := &canvases[Screen]

	s.commands = make([]gl.DrawIndirectCommand, 0, maxCommandCount)
	s.parameters = make([]int16, 0, maxCommandCount*maxParamCount)
	s.buffer = gl.NewFramebuffer()
	internal.ResizeScreen()
	s.buffer.Bind(gl.DrawReadFramebuffer)
}

//------------------------------------------------------------------------------

func createScreenTexture() {
	s := &canvases[Screen]

	s.texture.Delete()
	s.texture = gl.NewTexture2D(1, gl.R8UI, int32(s.size.X), int32(s.size.Y))
	s.buffer.Texture(gl.ColorAttachment0, s.texture, 0)

	s.depth.Delete()
	s.depth = gl.NewRenderbuffer(gl.Depth32F, int32(s.size.X), int32(s.size.Y))
	s.buffer.Renderbuffer(gl.DepthAttachment, s.depth)

	s.buffer.DrawBuffer(gl.ColorAttachment0)
	s.buffer.ReadBuffer(gl.NoAttachment)

	st := s.buffer.CheckStatus(gl.DrawReadFramebuffer)
	if st != gl.FramebufferComplete {
		setErr("while creating screen textures", errors.New(st.String()))
	}
}

//------------------------------------------------------------------------------

func init() {
	internal.ResizeScreen = func() {
		s := &canvases[Screen]

		if s.autozoom {
			// Find best fit for pixel size
			p1 := internal.Window.Width / int32(s.target.X)
			p2 := internal.Window.Height / int32(s.target.Y)
			if p1 < p2 {
				s.pixel = p1
			} else {
				s.pixel = p2
			}
			if s.pixel < 1 {
				s.pixel = 1
			}
		}

		// Extend the screen to cover the window
		s.size = Coord{
			int16(internal.Window.Width / s.pixel),
			int16(internal.Window.Height / s.pixel),
		}
		createScreenTexture()

		// Compute offset
		w := int32(s.size.X) * s.pixel
		h := int32(s.size.Y) * s.pixel
		s.ox = (internal.Window.Width - w) / 2
		s.oy = (internal.Window.Height - h) / 2

		internal.Loop.ScreenResized(s.size.X, s.size.Y, s.pixel)
	}
}

//------------------------------------------------------------------------------

func blitScreen() {
	s := &canvases[Screen]

	w := int32(s.size.X) * s.pixel
	h := int32(s.size.Y) * s.pixel

	blitUniforms.ScreenSize.X = float32(s.size.X)
	blitUniforms.ScreenSize.Y = float32(s.size.Y)
	blitUBO.SubData(&blitUniforms, 0)

	blitPipeline.Bind()
	// s.buffer.Bind(gl.ReadFramebuffer) //TODO: Useless?
	gl.DefaultFramebuffer.Bind(gl.ReadFramebuffer) //TODO: Useless?
	gl.DefaultFramebuffer.Bind(gl.DrawFramebuffer)
	gl.Enable(gl.FramebufferSRGB)
	gl.Disable(gl.Blend)
	gl.DepthMask(false)
	gl.Viewport(s.ox, s.oy, s.ox+w, s.ox+h)
	blitUBO.Bind(0)
	s.texture.Bind(0)
	gl.Draw(0, 4)
}

//------------------------------------------------------------------------------

// Mouse returns the mouse position on the virtual screen.
func (cv Canvas) Mouse() Coord {
	mx, my := mouse.Position()
	return Coord{
		X: int16((mx - canvases[cv].ox) / canvases[cv].pixel),
		Y: int16((my - canvases[cv].oy) / canvases[cv].pixel),
	}
}

//------------------------------------------------------------------------------
