// Copyright (c) 2013-2017 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package gfx

//------------------------------------------------------------------------------

import (
	"github.com/drakmaniso/carol/core/gl"
	pixel "github.com/drakmaniso/carol/pixel"
)

//------------------------------------------------------------------------------

var screen struct {
	framebuffer gl.Framebuffer
	size        pixel.Coord
	pixel       pixel.Coord
}

//------------------------------------------------------------------------------

func createScreen(size pixel.Coord, pixel pixel.Coord) {
	screen.framebuffer = gl.NewFramebuffer()
	screen.size = size
	screen.pixel = pixel

	ct := gl.NewTexture2D(1, int32(size.X), int32(size.Y), gl.RGB8)
	//TODO: parameter NEAREST

	screen.framebuffer.Texture(gl.ColorAttachment0, ct, 0)

	screen.framebuffer.DrawBuffer(gl.ColorAttachment0)

	gl.Viewport(0, 0, int32(size.X), int32(size.Y))
	screen.framebuffer.Bind(gl.DrawReadFramebuffer)
}

//------------------------------------------------------------------------------

func blitScreen(winSize pixel.Coord) {
	// screen.framebuffer.Bind(gl.ReadFramebuffer)
	gl.DefaultFramebuffer.Bind(gl.DrawFramebuffer)
	gl.ClearColorBuffer(RGBA{0.2, 0.2, 0.2, 1})

	w := screen.size.X * screen.pixel.X
	h := screen.size.Y * screen.pixel.Y
	ox := (winSize.X - w) / 2
	oy := (winSize.Y - h) / 2
	screen.framebuffer.Blit(
		gl.DefaultFramebuffer,
		0, 0, int32(screen.size.X), int32(screen.size.Y),
		int32(ox), int32(oy), int32(ox+w), int32(oy+h),
		gl.ColorBufferBit,
		gl.Nearest,
	)

	screen.framebuffer.Bind(gl.DrawFramebuffer)
	gl.ClearColorBuffer(RGBA{0, 0, 0, 1})
}

//------------------------------------------------------------------------------
