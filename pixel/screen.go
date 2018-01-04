// Copyright (c) 2013-2017 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package pixel

//------------------------------------------------------------------------------

import (
	"github.com/drakmaniso/carol/colour"
	"github.com/drakmaniso/carol/core/gl"
	"github.com/drakmaniso/carol/internal"
)

//------------------------------------------------------------------------------

var screen struct {
	buffer     gl.Framebuffer
	texture    gl.Texture2D
	depth      gl.Texture2D
	size       Coord
	pixel      int32
	background colour.RGBA
}

//------------------------------------------------------------------------------

func ScreenSize() Coord {
	return screen.size
}

func PixelSize() int32 {
	return screen.pixel
}

//------------------------------------------------------------------------------

func SetBackground(c colour.Colour) {
	screen.background = colour.RGBAOf(c)
}

//------------------------------------------------------------------------------

func createScreen() {
	if internal.Config.ScreenMode == "direct" {
		gl.Viewport(0, 0, int32(internal.Window.Width), int32(internal.Window.Height))
		return
	}

	screen.buffer = gl.NewFramebuffer()
	screen.size = Coord{
		int16(internal.Config.ScreenSize[0]),
		int16(internal.Config.ScreenSize[1]),
	}
	screen.pixel = internal.Config.PixelSize

	createScreenTexture()

	screen.buffer.Bind(gl.DrawReadFramebuffer)
}

//------------------------------------------------------------------------------

func createScreenTexture() {
	//TODO: delete previous texture
	screen.texture = gl.NewTexture2D(1, gl.SRGB8, int32(screen.size.X), int32(screen.size.Y))
	screen.buffer.Texture(gl.ColorAttachment0, screen.texture, 0)

	screen.buffer.DrawBuffer(gl.ColorAttachment0)
}

//------------------------------------------------------------------------------

func init() {
	internal.ResizeScreen = func() {
		switch internal.Config.ScreenMode {
		case "Extend":
			screen.size = Coord{
				int16(internal.Window.Width / screen.pixel),
				int16(internal.Window.Height / screen.pixel),
			}
			createScreenTexture()

		case "Zoom":
			p1 := internal.Window.Width / int32(screen.size.X)
			p2 := internal.Window.Height / int32(screen.size.Y)
			if p1 < p2 {
				screen.pixel = p1
			} else {
				screen.pixel = p2
			}
			if screen.pixel < 1 {
				screen.pixel = 1
			}

		default:
		}
	}
}

//------------------------------------------------------------------------------

func blitScreen() {
	if internal.Config.ScreenMode == "direct" {
		return
	}

	gl.DefaultFramebuffer.Bind(gl.DrawFramebuffer)
	gl.ClearColorBuffer(colour.RGBA{0.2, 0.2, 0.2, 1})

	w := int32(screen.size.X) * screen.pixel
	h := int32(screen.size.Y) * screen.pixel
	ox := (internal.Window.Width - w) / 2
	oy := (internal.Window.Height - h) / 2
	screen.buffer.Blit(
		gl.DefaultFramebuffer,
		0, 0, int32(screen.size.X), int32(screen.size.Y),
		ox, oy, ox+w, oy+h,
		gl.ColorBufferBit,
		gl.Nearest,
	)

}

//------------------------------------------------------------------------------
