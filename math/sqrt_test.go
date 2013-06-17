// Copyright (c) 2013 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package math

import (
	"math"
	"testing"
)

//------------------------------------------------------------------------------

func TestSqrt(t *testing.T) {
	a := Sqrt(float32(4))
	if a != 2.0 {
		t.Errorf("Wrong result: %v %T", a, a)
	}
	b := Sqrt(float32(3.3))
	if b != 1.8165902 {
		t.Errorf("Wrong result: %v %T", b, b)
	}
}

//------------------------------------------------------------------------------

func BenchmarkSqrt_math64(b *testing.B) {
	a := float64(3.3)
	for i := 0; i < b.N; i++ {
		_ = math.Sqrt(a)
	}
}

//------------------------------------------------------------------------------

func BenchmarkSqrt_math32(b *testing.B) {
	a := float32(3.3)
	for i := 0; i < b.N; i++ {
		_ = float32(math.Sqrt(float64(a)))
	}
}

//------------------------------------------------------------------------------

func BenchmarkSqrt_glm(b *testing.B) {
	a := float32(3.3)
	for i := 0; i < b.N; i++ {
		_ = Sqrt(a)
	}
}

//------------------------------------------------------------------------------
