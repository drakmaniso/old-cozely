// Copyright (c) 2013-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package math32

import (
	"math"
	"testing"
)

////////////////////////////////////////////////////////////////////////////////

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

////////////////////////////////////////////////////////////////////////////////

func BenchmarkFloor_math(b *testing.B) {
	x := float64(3.3)
	y := float64(-3.3)
	for i := 0; i < b.N; i++ {
		result64 = math.Floor(x)
		result64 = math.Floor(y)
	}
}

////////////////////////////////////////////////////////////////////////////////

func BenchmarkFloor_float32math(b *testing.B) {
	x := float32(3.3)
	y := float32(-3.3)
	for i := 0; i < b.N; i++ {
		result = float32(math.Floor(float64(x)))
		result = float32(math.Floor(float64(y)))
	}
}

////////////////////////////////////////////////////////////////////////////////

func BenchmarkFloor_glam(b *testing.B) {
	x := float32(3.3)
	y := float32(-3.3)
	for i := 0; i < b.N; i++ {
		result = Floor(x)
		result = Floor(y)
	}
}

////////////////////////////////////////////////////////////////////////////////
