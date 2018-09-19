// Copyright (c) 2013-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package math32

import (
	"math"
)

////////////////////////////////////////////////////////////////////////////////

// Floor returns the nearest integer less than or equal to x.
func Floor(x float32) float32

func floor(x float32) float32 {
	return float32(math.Floor(float64(x)))
}
