// Copyright (c) 2013-2017 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package gfx

//------------------------------------------------------------------------------

import (
	"image/color"
)

//------------------------------------------------------------------------------

// RGBA is a color defined by its red, green and blue components, with alpha.
type RGBA struct {
	R float32
	G float32
	B float32
	A float32
}

//------------------------------------------------------------------------------

func MakeRGBA(c color.Color) RGBA {
	r, g, b, a := c.RGBA()
	println(r, g, b, a)
	return RGBA{
		R: float32(r) / float32(0xFFFF),
		G: float32(g) / float32(0xFFFF),
		B: float32(b) / float32(0xFFFF),
		A: float32(a) / float32(0xFFFF),
	}
}

//------------------------------------------------------------------------------

// RGBA implements the image.Color interface: it returns the four components
// scaled by 0xFFFF.
func (c RGBA) RGBA() (r, g, b, a uint32) {
	return uint32(c.R * 0xFFFF), uint32(c.G * 0xFFFF), uint32(c.B * 0xFFFF), uint32(c.A * 0xFFFF)
}

//------------------------------------------------------------------------------
