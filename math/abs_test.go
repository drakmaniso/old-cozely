// Copyright (c) 2013 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package math

import (
	"math"
	"testing"
)

//------------------------------------------------------------------------------

func TestAbs(t *testing.T) {
	a := Abs(1.1)
	if a != 1.1 {
		t.Errorf("Wrong result for Abs(1.1): %#v\n", a)
	}
	b := Abs(-1.1)
	if b != 1.1 {
		t.Errorf("Wrong result for Abs(-1.1): %#v\n", b)
	}
	c := Abs(0.0)
	if c != 0.0 {
		t.Errorf("Wrong result for Abs(0.0): %#v\n", c)
	}
	d := Abs(-0.0)
	if d != 0.0 {
		t.Errorf("Wrong result for Abs(-0.0): %#v\n", d)
	}
}

//------------------------------------------------------------------------------

func BenchmarkAbs_math64(b *testing.B) {
	x := float64(3.3)
	y := float64(-3.3)
	for i := 0; i < b.N; i++ {
		_ = math.Abs(x)
		_ = math.Abs(y)
	}	
}

//------------------------------------------------------------------------------

func BenchmarkAbs_math32(b *testing.B) {
	x := float32(3.3)
	y := float32(-3.3)
	for i := 0; i < b.N; i++ {
		_ = float64(math.Abs(float64(x)))
		_ = float64(math.Abs(float64(y)))
	}	
}

//------------------------------------------------------------------------------

func switchAbs(x float32) float32 {
	switch {
	case x < 0:
		return -x
	case x == 0:
		return 0 // return correctly abs(-0)
	}
	return x
}

func BenchmarkAbs_switch(b *testing.B) {
	x := float32(3.3)
	y := float32(-3.3)
	for i := 0; i < b.N; i++ {
		_ = switchAbs(x)
		_ = switchAbs(y)
	}	
}

//------------------------------------------------------------------------------

func BenchmarkAbs_glm(b *testing.B) {
	x := float32(3.3)
	y := float32(-3.3)
	for i := 0; i < b.N; i++ {
		_ = Abs(x)
		_ = Abs(y)
	}	
}

//------------------------------------------------------------------------------
