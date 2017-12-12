// Copyright (c) 2013-2017 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package colour

//------------------------------------------------------------------------------

// RGBA represents a color in alpha-premultiplied linear color space. Each
// value ranges within [0, 1], and can be used directly by GPU shaders.
//
// An alpha-premultiplied color component c has been scaled by alpha (a), so has
// valid values 0 <= c <= a.
//
// Note that additive blending can also be achieved when alpha is set to 0 while
// the color components are non-null.
type RGBA struct {
	R float32
	G float32
	B float32
	A float32
}

//------------------------------------------------------------------------------

// RGBAOf converts any color to alpha-premultiplied, linear color space.
func RGBAOf(c Colour) RGBA {
	r, g, b, a := c.Linear()
	return RGBA{r, g, b, a}
}

//------------------------------------------------------------------------------

// Linear implements the Colour interface.
func (c RGBA) Linear() (r, g, b, a float32) {
	return c.R, c.G, c.B, c.A
}

//------------------------------------------------------------------------------

// RGBA implements the image.Color interface: it returns the four components
// scaled by 0xFFFF.
func (c RGBA) RGBA() (r, g, b, a uint32) {
	r = uint32(c.R * 0xFFFF)
	g = uint32(c.G * 0xFFFF)
	b = uint32(c.B * 0xFFFF)
	a = uint32(c.A * 0xFFFF)
	return r, g, b, a
}

//------------------------------------------------------------------------------
