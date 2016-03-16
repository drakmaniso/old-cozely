// Copyright (c) 2013-2016 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package math

import (
	"testing"
)

//------------------------------------------------------------------------------

func TestFastCos(t *testing.T) {
	var x float32
	maxDiff := float32(0)
	for _, tt := range cos_tests {
		x = FastCos(tt.in)
		if Abs(x-tt.out) > maxDiff {
			maxDiff = Abs(x - tt.out)
		}
		if !IsRoughlyEqual(x, tt.out, 1e-3) {
			t.Errorf("ULP error for Cos(%.100g): %.100g instead of %.100g\n", tt.in, x, tt.out)
		}
	}
	t.Logf("Max absolute error: %1.8e\n", maxDiff)
}

//------------------------------------------------------------------------------

func BenchmarkFastCos_go(b *testing.B) {
	a := float32(0.5)
	for i := 0; i < b.N; i++ {
		_ = fastCos(a)
	}
}

//------------------------------------------------------------------------------

func BenchmarkFastCos_glam(b *testing.B) {
	a := float32(0.5)
	for i := 0; i < b.N; i++ {
		_ = FastCos(a)
	}
}

//------------------------------------------------------------------------------
