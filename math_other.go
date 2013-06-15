// Copyright (c) 2013 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

// +build !386
// +build !amd64

package glm

import (
	"math"
)

//------------------------------------------------------------------------------

// Sqrt returns the square root of x.
func Sqrt(x float32) float32 {
	return float32(math.Sqrt(float64(x)))
}

//------------------------------------------------------------------------------

// Mix returns the linear interpolation between x and y using a to weight between them.
// The value is computed as follows: x*(1-a) + y*a
func Mix(x, y float32, a float32) float32 {
	return x*(1-a) + y*a
}

//------------------------------------------------------------------------------

// Floor returns the nearest integer less than or equal to x.
func Floor(x float32) float32 {
	return float32(math.Floor(float64(x)))
}

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
