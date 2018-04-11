// Copyright (c) 2013-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package colour

////////////////////////////////////////////////////////////////////////////////

// SRGBnA8 represents a 32-bit color in *non* alpha-premultiplied sRGB color
// space. There is 8 bits for each compnent.
type SRGBnA8 struct {
	R uint8
	G uint8
	B uint8
	A uint8
}

////////////////////////////////////////////////////////////////////////////////

// SRGBnA8Of converts any color to non alpha-premultiplied sRGB color space.
func SRGBnA8Of(c Colour) SRGBnA8 {
	cc, ok := c.(SRGBnA8)
	if ok {
		return cc
	}
	r, g, b, a := c.Linear()
	r = standardOf(r / a)
	g = standardOf(g / a)
	b = standardOf(b / a)
	return SRGBnA8{uint8(r * 0xFF), uint8(g * 0xFF), uint8(b * 0xFF), uint8(a * 0xFF)}
}

////////////////////////////////////////////////////////////////////////////////

// Linear implements the Colour interface.
func (c SRGBnA8) Linear() (r, g, b, a float32) {
	a = float32(c.A) / float32(0xFF)
	r = a * linearOf(float32(c.R)/float32(0xFF))
	g = a * linearOf(float32(c.G)/float32(0xFF))
	b = a * linearOf(float32(c.B)/float32(0xFF))
	return r, g, b, a
}

// Standard implements the Colour interface.
func (c SRGBnA8) Standard() (r, g, b, a float32) {
	a = float32(c.A) / float32(0xFF)
	r = standardOf(a * linearOf(float32(c.R)/float32(0xFF)))
	g = standardOf(a * linearOf(float32(c.R)/float32(0xFF)))
	b = standardOf(a * linearOf(float32(c.R)/float32(0xFF)))
	return r, g, b, a
}

////////////////////////////////////////////////////////////////////////////////

// RGBA implements the image.Color interface: it returns the four components
// scaled by 0xFFFF.
func (c SRGBnA8) RGBA() (r, g, b, a uint32) {
	alpha := float64(c.A) / float64(0xFF)
	r = uint32(alpha * float64(c.A) * (float64(c.R) / float64(0xFF)) * float64(0xFFFF))
	g = uint32(alpha * float64(c.A) * (float64(c.G) / float64(0xFF)) * float64(0xFFFF))
	b = uint32(alpha * float64(c.A) * (float64(c.B) / float64(0xFF)) * float64(0xFFFF))
	a = uint32(alpha * float64(0xFFFF))
	return r, g, b, a
}

////////////////////////////////////////////////////////////////////////////////
