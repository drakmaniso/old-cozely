// Copyright (c) 2013-2017 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package colour

//------------------------------------------------------------------------------

// RGB represents a color in linear color space. Each value ranges within
// [0, 1], and can be used directly by GPU shaders.
type RGB struct {
	R float32
	G float32
	B float32
}

//------------------------------------------------------------------------------

// RGBOf converts any color to linear color space with no alpha.
func RGBOf(c Colour) RGB {
	r, g, b, _ := c.Linear()
	return RGB{r, g, b}
}

//------------------------------------------------------------------------------

// Linear implements the Colour interface.
func (c RGB) Linear() (r, g, b, a float32) {
	return c.R, c.G, c.B, 1
}

//------------------------------------------------------------------------------

// RGBA implements the image.Color interface: it returns the four components
// scaled by 0xFFFF.
func (c RGB) RGBA() (r, g, b, a uint32) {
	r = uint32(c.R * 0xFFFF)
	g = uint32(c.G * 0xFFFF)
	b = uint32(c.B * 0xFFFF)
	a = uint32(0xFFFF)
	return r, g, b, a
}

//------------------------------------------------------------------------------
