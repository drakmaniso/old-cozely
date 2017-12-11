// Copyright (c) 2013-2017 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package gfx

//------------------------------------------------------------------------------

import (
	"github.com/drakmaniso/carol/colour"
	"github.com/drakmaniso/carol/core/gl"
	"github.com/drakmaniso/carol/internal"
	"github.com/drakmaniso/carol/pixel"
)

//------------------------------------------------------------------------------

var screen struct {
	buffer  gl.Framebuffer
	texture gl.Texture2D
	depth   gl.Texture2D
	size    pixel.Coord
	pixel   pixel.Coord
}

//------------------------------------------------------------------------------

func ScreenSize() pixel.Coord {
	return screen.size
}

//------------------------------------------------------------------------------

func createScreen() {
	if internal.Config.ScreenMode == "direct" {
		gl.Viewport(0, 0, int32(internal.Window.Size.X), int32(internal.Window.Size.Y))
		return
	}

	screen.buffer = gl.NewFramebuffer()
	screen.size = internal.Config.ScreenSize
	screen.pixel = internal.Config.PixelSize

	createScreenTexture()

	screen.buffer.Bind(gl.DrawReadFramebuffer)
}

//------------------------------------------------------------------------------

func createScreenTexture() {
	//TODO: delete previous texture
	screen.texture = gl.NewTexture2D(1, int32(screen.size.X), int32(screen.size.Y), gl.RGB8)
	screen.buffer.Texture(gl.ColorAttachment0, screen.texture, 0)

	screen.depth = gl.NewTexture2D(1, int32(screen.size.X), int32(screen.size.Y), gl.Depth24)
	screen.buffer.Texture(gl.DepthAttachment, screen.depth, 0)

	screen.buffer.DrawBuffer(gl.ColorAttachment0)
}

//------------------------------------------------------------------------------

func init() {
	internal.ResizeScreen = func() {
		switch internal.Config.ScreenMode {
		case "Extend":
			screen.size = internal.Window.Size.SlashCW(screen.pixel)
			createScreenTexture()

		case "Zoom":
			r1 := float64(screen.pixel.X) / float64(screen.pixel.Y)
			screen.pixel = internal.Window.Size.SlashCW(screen.size)
			if screen.pixel.X < 1 {
				screen.pixel.X = 1
			}
			if screen.pixel.Y < 1 {
				screen.pixel.Y = 1
			}
			r2 := float64(screen.pixel.X) / float64(screen.pixel.Y)
			if r1 < r2 {
				screen.pixel.X = int16(float64(screen.pixel.Y) * r1)
				if screen.pixel.X < 1 {
					screen.pixel.X = 1
					screen.pixel.Y = int16(float64(screen.pixel.X) / r1)
				}
			} else if r1 > r2 {
				screen.pixel.Y = int16(float64(screen.pixel.X) / r1)
				if screen.pixel.Y < 1 {
					screen.pixel.Y = 1
					screen.pixel.X = int16(float64(screen.pixel.Y) * r1)
				}
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

	w := screen.size.X * screen.pixel.X
	h := screen.size.Y * screen.pixel.Y
	ox := (internal.Window.Size.X - w) / 2
	oy := (internal.Window.Size.Y - h) / 2
	screen.buffer.Blit(
		gl.DefaultFramebuffer,
		0, 0, int32(screen.size.X), int32(screen.size.Y),
		int32(ox), int32(oy), int32(ox+w), int32(oy+h),
		gl.ColorBufferBit,
		gl.Nearest,
	)

}

//------------------------------------------------------------------------------
