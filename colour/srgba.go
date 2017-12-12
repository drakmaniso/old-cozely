// Copyright (c) 2013-2017 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package colour

import (
	"math"
)

//------------------------------------------------------------------------------

func srgbToLinear(c float32) float32 {
	if c <= 0.04045 {
		return c / 12.92
	}
	return float32(math.Pow(float64(c+0.055)/(1+0.055), 2.4))
}

func linearToSrgb(c float32) float32 {
	if c <= 0.0031308 {
		return 12.92 * c
	}
	return (1+0.055)*float32(math.Pow(float64(c), 1/2.4)) - 0.055
}

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
	r, g, b, a := c.Linear()
	r = linearToSrgb(r)
	g = linearToSrgb(g)
	b = linearToSrgb(b)
	return SRGBA{r, g, b, a}
}

//------------------------------------------------------------------------------

// Linear implements the Colour interface: it returns the color converted to
// linear color space.
func (c SRGBA) Linear() (r, g, b, a float32) {
	a = c.A
	r = srgbToLinear(c.R)
	g = srgbToLinear(c.G)
	b = srgbToLinear(c.B)
	return r, g, b, a
}

//------------------------------------------------------------------------------

// RGBA implements the image.Color interface: it returns the four components
// scaled by 0xFFFF.
func (c SRGBA) RGBA() (r, g, b, a uint32) {
	return uint32(c.R * 0xFFFF), uint32(c.G * 0xFFFF), uint32(c.B * 0xFFFF), uint32(c.A * 0xFFFF)
}

//------------------------------------------------------------------------------
