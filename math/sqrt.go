// Copyright (c) 2013 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package math

import (
	smath "math"
)

//------------------------------------------------------------------------------

//go:noescape

// `Sqrt` returns the square root of `x`.
func Sqrt(x float32) float32

func sqrt(x float32) float32 {
	return float32(smath.Sqrt(float64(x)))
}

//------------------------------------------------------------------------------
