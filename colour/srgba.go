// Copyright (c) 2013-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package colour

//------------------------------------------------------------------------------

// SRGBA represents a color in alpha-premultiplied sRGB color space. Each value
// ranges within [0, 1].
//
// An alpha-premultiplied color component c has been scaled by alpha (a), so has
// valid values 0 <= c <= a.
//
// Note that additive blending can also be achieved when alpha is set to 0 while
// the color components are non-null.
type SRGBA struct {
	R float32
	G float32
	B float32
	A float32
}

//------------------------------------------------------------------------------

// SRGBAOf converts any color to sRGB alpha-premultiplied color space.
func SRGBAOf(c Colour) SRGBA {
	r, g, b, a := c.Standard()
	return SRGBA{r, g, b, a}
}

//------------------------------------------------------------------------------

// Linear implements the Colour interface: it returns the color converted to
// linear color space.
func (c SRGBA) Linear() (r, g, b, a float32) {
	r = linearOf(c.R)
	g = linearOf(c.G)
	b = linearOf(c.B)
	return r, g, b, c.A
}

// Standard implements the Colour interface.
func (c SRGBA) Standard() (r, g, b, a float32) {
	return c.R, c.G, c.B, c.A
}

//------------------------------------------------------------------------------

// RGBA implements the image.Color interface: it returns the four components
// scaled by 0xFFFF.
func (c SRGBA) RGBA() (r, g, b, a uint32) {
	return uint32(c.R * 0xFFFF), uint32(c.G * 0xFFFF), uint32(c.B * 0xFFFF), uint32(c.A * 0xFFFF)
}

//------------------------------------------------------------------------------
