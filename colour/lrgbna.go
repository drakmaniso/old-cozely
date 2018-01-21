// Copyright (c) 2013-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package colour

//------------------------------------------------------------------------------

// LRGBnA represents a color in *non* alpha-premultiplied linear color
// space. Each value ranges within [0, 1], and can be used directly by GPU
// shaders.
//
// Note: prefer RGBA for use in shaders.
type LRGBnA struct {
	R float32
	G float32
	B float32
	A float32
}

//------------------------------------------------------------------------------

// LRGBnAOf converts any color to non alpha-premultiplied linear color
// space.
func LRGBnAOf(c Colour) LRGBnA {
	r, g, b, a := c.Linear()
	return LRGBnA{r / a, g / a, b / a, a}
}

//------------------------------------------------------------------------------

// Linear implements the Colour interface.
func (c LRGBnA) Linear() (r, g, b, a float32) {
	return c.A * c.R, c.A * c.G, c.A * c.B, c.A
}

// Standard implements the Colour interface.
func (c LRGBnA) Standard() (r, g, b, a float32) {
	r = standardOf(c.A * c.R)
	g = standardOf(c.A * c.G)
	b = standardOf(c.A * c.B)
	return r, g, b, c.A
}

//------------------------------------------------------------------------------

// RGBA implements the image.Color interface: it returns the four components
// scaled by 0xFFFF.
func (c LRGBnA) RGBA() (r, g, b, a uint32) {
	return uint32(c.A * c.R * 0xFFFF), uint32(c.A * c.G * 0xFFFF), uint32(c.A * c.B * 0xFFFF), uint32(c.A * 0xFFFF)
}

//------------------------------------------------------------------------------
