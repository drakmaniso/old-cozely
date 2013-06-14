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

func TestFloor(t *testing.T) {
	a := Floor(float32(3.3))
	if a != 3 {
		t.Errorf("Wrong result for Floor(3.3): %v %T", a, a)
	}
	b := Floor(float32(-3.3))
	if b != -4 {
		t.Errorf("Wrong result for Floor(-3.3): %v %T", b, b)
	}
	c := Floor(float32(3))
	if c != 3 {
		t.Errorf("Wrong result for Floor(3): %v %T", c, c)
	}
	d := Floor(float32(-3))
	if d != -3 {
		t.Errorf("Wrong result for Floor(-3): %v %T", d, d)
	}
}

func BenchmarkMathFloorFloat64(b *testing.B) {
	x := float64(3.3)
	y := float64(-3.3)
	for i := 0; i < b.N; i++ {
		_ = math.Floor(x)
		_ = math.Floor(y)
	}
}

func BenchmarkMathFloorFloat32(b *testing.B) {
	x := float32(3.3)
	y := float32(-3.3)
	for i := 0; i < b.N; i++ {
		_ = float32(math.Floor(float64(x)))
		_ = float32(math.Floor(float64(y)))
	}
}

func BenchmarkGlmFloor(b *testing.B) {
	x := float32(3.3)
	y := float32(-3.3)
	for i := 0; i < b.N; i++ {
		_ = Floor(x)
		_ = Floor(y)
	}
}

//------------------------------------------------------------------------------
// Copyright (c) 2013 - Laurent Moussault <moussault.laurent@gmail.com>
