// Copyright (c) 2013 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package math

//------------------------------------------------------------------------------

// `Mix` returns the linear interpolation between `x` and `y` using `a` to weight between them.
// The value is computed as follows: `x*(1-a) + y*a`
func Mix(x, y float32, a float32) float32 {
	return x*(1-a) + y*a
}

//------------------------------------------------------------------------------
