// Copyright (c) 2013-2017 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package colour

//------------------------------------------------------------------------------

// SRGB8 represents a 24-bit color in sRGB color space. There is 8 bits for each
// components.
type SRGB8 struct {
	R uint8
	G uint8
	B uint8
}

//------------------------------------------------------------------------------

// SRGB8Of converts any color to sRGB color space.
func SRGB8Of(c Colour) SRGB8 {
	cc, ok := c.(SRGB8)
	if ok {
		return cc
	}
	ccc, ok := c.(SRGBA8)
	if ok {
		return SRGB8{ccc.R, ccc.G, ccc.B}
	}
	r, g, b, _ := c.Standard()
	return SRGB8{uint8(r * 0xFF), uint8(g * 0xFF), uint8(b * 0xFF)}
}

//------------------------------------------------------------------------------

// Linear implements the Colour interface.
func (c SRGB8) Linear() (r, g, b, a float32) {
	r = linearOf(float32(c.R) / float32(0xFF))
	g = linearOf(float32(c.G) / float32(0xFF))
	b = linearOf(float32(c.B) / float32(0xFF))
	a = 1
	return r, g, b, a
}

// Standard implements the Colour interface.
func (c SRGB8) Standard() (r, g, b, a float32) {
	r = float32(c.R) / float32(0xFF)
	g = float32(c.R) / float32(0xFF)
	b = float32(c.R) / float32(0xFF)
	return r, g, b, 1
}

//------------------------------------------------------------------------------

// RGBA implements the image.Color interface.
func (c SRGB8) RGBA() (r, g, b, a uint32) {
	r = uint32((float64(c.R) / float64(0xFF)) * float64(0xFFFF))
	g = uint32((float64(c.G) / float64(0xFF)) * float64(0xFFFF))
	b = uint32((float64(c.B) / float64(0xFF)) * float64(0xFFFF))
	a = uint32(0xFFFF)
	return r, g, b, a
}

//------------------------------------------------------------------------------
