// Copyright (c) 2013-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package math32

//------------------------------------------------------------------------------

// Round returns the nearest integer to x.
func Round(x float32) int32 {
	if x > 0 {
		return int32(x + 0.5)
	}
	return int32(x - 0.5)
}

//------------------------------------------------------------------------------
