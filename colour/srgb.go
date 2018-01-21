// Copyright (c) 2013-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package colour

//------------------------------------------------------------------------------

// SRGB represents a color in sRGB color space. Each value ranges within
// [0, 1], and can be used directly by GPU shaders.
type SRGB struct {
	R float32
	G float32
	B float32
}

//------------------------------------------------------------------------------

// SRGBOf converts any color to sRGB color space with no alpha.
func SRGBOf(c Colour) SRGB {
	r, g, b, _ := c.Standard()
	return SRGB{r, g, b}
}

//------------------------------------------------------------------------------

// Linear implements the Colour interface.
func (c SRGB) Linear() (r, g, b, a float32) {
	r = linearOf(c.R)
	g = linearOf(c.G)
	b = linearOf(c.B)
	return r, g, b, 1
}

// Standard implements the Colour interface.
func (c SRGB) Standard() (r, g, b, a float32) {
	return c.R, c.G, c.B, 1
}

//------------------------------------------------------------------------------

// RGBA implements the image.Color interface: it returns the four components
// scaled by 0xFFFF.
func (c SRGB) RGBA() (r, g, b, a uint32) {
	return uint32(c.R * 0xFFFF), uint32(c.G * 0xFFFF), uint32(c.B * 0xFFFF), uint32(0xFFFF)
}

//------------------------------------------------------------------------------
