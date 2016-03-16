// Copyright (c) 2013-2016 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package math

import (
	"math"
	"testing"
	"unsafe"
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

func bitsAbs(x float32) float32 {
	return Float32frombits(Float32bits(x) & 0x7FFFFFFF)
}

func BenchmarkAbs_bits(b *testing.B) {
	x := float32(3.3)
	y := float32(-3.3)
	for i := 0; i < b.N; i++ {
		_ = bitsAbs(x)
		_ = bitsAbs(y)
	}
}

//------------------------------------------------------------------------------

func unsafeAbs(x float32) float32 {
	ux := *(*uint32)(unsafe.Pointer(&x)) & 0x7FFFFFFF
	return *(*float32)(unsafe.Pointer(&ux))
}

func BenchmarkAbs_unsafe(b *testing.B) {
	x := float32(3.3)
	y := float32(-3.3)
	for i := 0; i < b.N; i++ {
		_ = unsafeAbs(x)
		_ = unsafeAbs(y)
	}
}

//------------------------------------------------------------------------------

func abs_asm(x float32) float32

func BenchmarkAbs_asm(b *testing.B) {
	x := float32(3.3)
	y := float32(-3.3)
	for i := 0; i < b.N; i++ {
		_ = abs_asm(x)
		_ = abs_asm(y)
	}
}

//------------------------------------------------------------------------------

func BenchmarkAbs_glam(b *testing.B) {
	x := float32(3.3)
	y := float32(-3.3)
	for i := 0; i < b.N; i++ {
		_ = Abs(x)
		_ = Abs(y)
	}
}

//------------------------------------------------------------------------------
