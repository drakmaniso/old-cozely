// Copyright (c) 2013-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package colour

import (
	"math"
)

////////////////////////////////////////////////////////////////////////////////

// Colour can convert itself to alpha-premultipled RGBA as 32 bit floats, in
// both linear and standard (sRGB) color spaces.
type Colour interface {
	// Linear returns the red, green, blue and alpha values in linear color space.
	// The red, gren and blue values have been alpha-premultiplied in linear
	// space. Each value ranges within [0, 1] and can be used directly by GPU
	// shaders.
	Linear() (r, g, b, a float32)

	// Standard returns the red, green, blue and alpha values in standard (sRGB)
	// color space. The red, gren and blue values have been alpha-premultiplied in
	// linear space. Each value ranges within [0, 1].
	Standard() (r, g, b, a float32)
}

////////////////////////////////////////////////////////////////////////////////

func linearOf(c float32) float32 {
	if c <= 0.04045 {
		return c / 12.92
	}
	return float32(math.Pow(float64(c+0.055)/(1+0.055), 2.4))
}

func standardOf(c float32) float32 {
	if c <= 0.0031308 {
		return 12.92 * c
	}
	return (1+0.055)*float32(math.Pow(float64(c), 1/2.4)) - 0.055
}

////////////////////////////////////////////////////////////////////////////////
