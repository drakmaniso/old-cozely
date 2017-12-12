// Copyright (c) 2013-2017 Laurent Moussault. All rights reserved.
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
	r, g, b, _ := c.Linear()
	r = linearToSrgb(r)
	g = linearToSrgb(g)
	b = linearToSrgb(b)
	return SRGB{r, g, b}
}

//------------------------------------------------------------------------------

// RGBA implements the Colour interface.
func (c SRGB) RGBA() (r, g, b, a float32) {
	r = srgbToLinear(c.R)
	g = srgbToLinear(c.G)
	b = srgbToLinear(c.B)
	return r, g, b, 1
}

//------------------------------------------------------------------------------

// Linear implements the image.Color interface: it returns the four components
// scaled by 0xFFFF.
func (c SRGB) Linear() (r, g, b, a uint32) {
	return uint32(c.R * 0xFFFF), uint32(c.G * 0xFFFF), uint32(c.B * 0xFFFF), uint32(0xFFFF)
}

//------------------------------------------------------------------------------
