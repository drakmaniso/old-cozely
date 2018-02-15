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

type ScreenCanvas struct {
	buffer     gl.Framebuffer
	texture    gl.Texture2D
	depth      gl.RenderBuffer
	target     Coord
	autozoom   bool
	size       Coord
	pixel      int32
	ox, oy     int32 // Offset when there is a border around the screen
	background palette.Index
	commands   []gl.DrawIndirectCommand
	parameters []int16
}

var screen ScreenCanvas

func init() {
	screen.target = Coord{X: 640, Y: 360}
	screen.pixel = 2
}

//------------------------------------------------------------------------------

func Screen() *ScreenCanvas {
	return &screen
}

//------------------------------------------------------------------------------

func (s *ScreenCanvas) Size() Coord {
	return s.size
}

func (s *ScreenCanvas) Pixel() int32 {
	return s.pixel
}

//------------------------------------------------------------------------------

func SetBackground(c palette.Index) {
	screen.background = c
}

//------------------------------------------------------------------------------

func createScreen() {
	screen.commands = make([]gl.DrawIndirectCommand, 0, maxCommandCount)
	screen.parameters = make([]int16, 0, maxCommandCount*maxParamCount)
	screen.buffer = gl.NewFramebuffer()
	internal.ResizeScreen()
	screen.buffer.Bind(gl.DrawReadFramebuffer)
}

//------------------------------------------------------------------------------

func createScreenTexture() {
	screen.texture.Delete()
	screen.texture = gl.NewTexture2D(1, gl.R8UI, int32(screen.size.X), int32(screen.size.Y))
	screen.buffer.Texture(gl.ColorAttachment0, screen.texture, 0)

	screen.depth.Delete()
	screen.depth = gl.NewRenderBuffer(gl.Depth32F, int32(screen.size.X), int32(screen.size.Y))
	screen.buffer.RenderBuffer(gl.DepthAttachment, screen.depth)

	screen.buffer.DrawBuffer(gl.ColorAttachment0)
	screen.buffer.ReadBuffer(gl.NoAttachment)

	s := screen.buffer.CheckStatus(gl.DrawReadFramebuffer)
	if s != gl.FramebufferComplete {
		setErr("while creating screen textures", errors.New(s.String()))
	}
}

//------------------------------------------------------------------------------

func init() {
	internal.ResizeScreen = func() {
		if screen.autozoom {
			// Find best fit for pixel size
			p1 := internal.Window.Width / int32(screen.target.X)
			p2 := internal.Window.Height / int32(screen.target.Y)
			if p1 < p2 {
				screen.pixel = p1
			} else {
				screen.pixel = p2
			}
			if screen.pixel < 1 {
				screen.pixel = 1
			}
		}

		// Extend the screen to cover the window
		screen.size = Coord{
			int16(internal.Window.Width / screen.pixel),
			int16(internal.Window.Height / screen.pixel),
		}
		createScreenTexture()

		// Compute offset
		w := int32(screen.size.X) * screen.pixel
		h := int32(screen.size.Y) * screen.pixel
		screen.ox = (internal.Window.Width - w) / 2
		screen.oy = (internal.Window.Height - h) / 2

		internal.Loop.ScreenResized(screen.size.X, screen.size.Y, screen.pixel)
	}
}

//------------------------------------------------------------------------------

func blitScreen() {
	w := int32(screen.size.X) * screen.pixel
	h := int32(screen.size.Y) * screen.pixel

	blitUniforms.ScreenSize.X = float32(screen.size.X)
	blitUniforms.ScreenSize.Y = float32(screen.size.Y)
	blitUBO.SubData(&blitUniforms, 0)

	blitPipeline.Bind()
	// screen.buffer.Bind(gl.ReadFramebuffer) //TODO: Useless?
	gl.DefaultFramebuffer.Bind(gl.ReadFramebuffer) //TODO: Useless?
	gl.DefaultFramebuffer.Bind(gl.DrawFramebuffer)
	gl.Enable(gl.FramebufferSRGB)
	gl.Disable(gl.Blend)
	gl.DepthMask(false)
	gl.Viewport(screen.ox, screen.oy, screen.ox+w, screen.ox+h)
	blitUBO.Bind(0)
	screen.texture.Bind(0)
	gl.Draw(0, 4)
}

//------------------------------------------------------------------------------

// Mouse returns the mouse position on the virtual screen.
func (s *ScreenCanvas) Mouse() Coord {
	mx, my := mouse.Position()
	return Coord{
		X: int16((mx - s.ox) / s.pixel),
		Y: int16((my - s.oy) / s.pixel),
	}
}

//------------------------------------------------------------------------------
