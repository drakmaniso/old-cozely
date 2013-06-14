package glm

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

func BenchmarkMathSqrtFloat64(b *testing.B) {
	a := float64(3.3)
	for i := 0; i < b.N; i++ {
		_ = math.Sqrt(a)
	}
}

func BenchmarkMathSqrtFloat32(b *testing.B) {
	a := float32(3.3)
	for i := 0; i < b.N; i++ {
		_ = float32(math.Sqrt(float64(a)))
	}
}

func BenchmarkGlmSqrt(b *testing.B) {
	a := float32(3.3)
	for i := 0; i < b.N; i++ {
		_ = Sqrt(a)
	}
}

//------------------------------------------------------------------------------
// Copyright (c) 2013 - Laurent Moussault <moussault.laurent@gmail.com>
