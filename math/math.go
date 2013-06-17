// Copyright (c) 2013 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package math

//------------------------------------------------------------------------------

// FastFloor returns int32(x) if x>0, int32(x-1) otherwise.
func FastFloor(x float32) int32 {
	if x > 0 {
		return int32(x)
	} else {
		return int32(x - 1)
	}
}

//------------------------------------------------------------------------------

// Round returns the nearest integer to x.
func Round(x float32) int32 {
	if x > 0 {
		return int32(x + 0.5)
	} else {
		return int32(x - 0.5)
	}
}

//------------------------------------------------------------------------------
