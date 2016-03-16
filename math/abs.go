// Copyright (c) 2013-2016 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package math

import "unsafe"

//------------------------------------------------------------------------------

// Abs returns the absolute value of x.
func Abs(x float32) float32 {
	ux := *(*uint32)(unsafe.Pointer(&x)) & 0x7FFFFFFF
	return *(*float32)(unsafe.Pointer(&ux))
}

//------------------------------------------------------------------------------
