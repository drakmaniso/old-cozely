// Copyright (c) 2013-2017 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package colour

//------------------------------------------------------------------------------

type Colour interface {
	// Linear returns the alpha-premultiplied red, green, blue and alpha
	// values for the color, in linear color space. Each value ranges within
	// [0, 1] and can be used directly by GPU shaders.
	//
	// An alpha-premultiplied color component c has been scaled by alpha (a), so
	// has valid values 0 <= c <= a.
	//
	// Note that additive blending can also be achieved when alpha is set to 0
	// while the color components are non-null.
	Linear() (r, g, b, a float32)
}

//------------------------------------------------------------------------------
