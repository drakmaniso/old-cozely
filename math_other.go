// +build !386
// +build !amd64

package glm

import (
	"math"
)

//------------------------------------------------------------------------------

func Sqrt(x float32) float32 {
	return float32(math.Sqrt(float64(x)))
}

//------------------------------------------------------------------------------

func Floor(x float32) float32 {
	return float32(math.Floor(float64(x)))
}

//------------------------------------------------------------------------------
// Copyright (c) 2013 - Laurent Moussault <moussault.laurent@gmail.com>
