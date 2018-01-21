// Copyright (c) 2013-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package math32

// FastFloor returns int32(x) if x>0, int32(x-1) otherwise.
func FastFloor(x float32) int32 {
	if x > 0 {
		return int32(x)
	}
	return int32(x - 1)
}

//------------------------------------------------------------------------------
