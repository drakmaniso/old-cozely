// Copyright (c) 2013-2017 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package colour

//------------------------------------------------------------------------------

// RGBnA represents a color in *non* alpha-premultiplied linear color
// space. Each value ranges within [0, 1], and can be used directly by GPU
// shaders.
//
// Note: prefer RGBA for use in shaders.
type RGBnA struct {
	R float32
	G float32
	B float32
	A float32
}

//------------------------------------------------------------------------------

// RGBnAOf converts any color to non alpha-premultiplied linear color
// space.
func RGBnAOf(c Colour) RGBnA {
	r, g, b, a := c.Linear()
	return RGBnA{r / a, g / a, b / a, a}
}

//------------------------------------------------------------------------------

// Linear implements the Colour interface.
func (c RGBnA) Linear() (r, g, b, a float32) {
	return c.A * c.R, c.A * c.G, c.A * c.B, c.A
}

//------------------------------------------------------------------------------

// RGBA implements the image.Color interface: it returns the four components
// scaled by 0xFFFF.
func (c RGBnA) RGBA() (r, g, b, a uint32) {
	return uint32(c.A * c.R * 0xFFFF), uint32(c.A * c.G * 0xFFFF), uint32(c.A * c.B * 0xFFFF), uint32(c.A * 0xFFFF)
}

//------------------------------------------------------------------------------
