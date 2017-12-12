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
	buffer         gl.Framebuffer
	texture        gl.Texture2D
	depth          gl.Texture2D
	size           pixel.Coord
	pixelW, pixelH int32
}

//------------------------------------------------------------------------------

func ScreenSize() pixel.Coord {
	return screen.size
}

//------------------------------------------------------------------------------

func createScreen() {
	if internal.Config.ScreenMode == "direct" {
		gl.Viewport(0, 0, int32(internal.Window.Width), int32(internal.Window.Height))
		return
	}

	screen.buffer = gl.NewFramebuffer()
	screen.size = pixel.Coord{
		int16(internal.Config.ScreenSize[0]),
		int16(internal.Config.ScreenSize[1]),
	}
	screen.pixelW = internal.Config.PixelSize[0]
	screen.pixelH = internal.Config.PixelSize[1]

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
			screen.size = pixel.Coord{
				int16(internal.Window.Width / screen.pixelW),
				int16(internal.Window.Height / screen.pixelH),
			}
			createScreenTexture()

		case "Zoom":
			r1 := float64(screen.pixelW) / float64(screen.pixelH)
			screen.size = pixel.Coord{
				int16(internal.Window.Width / screen.pixelW),
				int16(internal.Window.Height / screen.pixelH),
			}
			if screen.pixelW < 1 {
				screen.pixelW = 1
			}
			if screen.pixelH < 1 {
				screen.pixelH = 1
			}
			r2 := float64(screen.pixelW) / float64(screen.pixelH)
			if r1 < r2 {
				screen.pixelW = int32(float64(screen.pixelH) * r1)
				if screen.pixelW < 1 {
					screen.pixelW = 1
					screen.pixelH = int32(float64(screen.pixelW) / r1)
				}
			} else if r1 > r2 {
				screen.pixelH = int32(float64(screen.pixelW) / r1)
				if screen.pixelH < 1 {
					screen.pixelH = 1
					screen.pixelW = int32(float64(screen.pixelH) * r1)
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

	w := int32(screen.size.X) * screen.pixelW
	h := int32(screen.size.Y) * screen.pixelH
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
