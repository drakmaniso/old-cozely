// Copyright (c) 2013-2016 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package math

import (
	"testing"
)

//------------------------------------------------------------------------------

func TestFastSin(t *testing.T) {
	var x float32
	maxDiff := float32(0)
	for _, tt := range sin_tests {
		x = FastSin(tt.in)
		if Abs(x-tt.out) > maxDiff {
			maxDiff = Abs(x - tt.out)
		}
		if !IsRoughlyEqual(x, tt.out, 1e-3) {
			t.Errorf("ULP error for Sin(%.100g): %.100g instead of %.100g\n", tt.in, x, tt.out)
		}
	}
	t.Logf("Max absolute error: %1.8e\n", maxDiff)
}

//------------------------------------------------------------------------------

func BenchmarkFastSin_go(b *testing.B) {
	a := float32(0.5)
	for i := 0; i < b.N; i++ {
		_ = fastSin(a)
	}
}

//------------------------------------------------------------------------------

func BenchmarkFastSin_glam(b *testing.B) {
	a := float32(0.5)
	for i := 0; i < b.N; i++ {
		_ = FastSin(a)
	}
}

//------------------------------------------------------------------------------
