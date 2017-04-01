// Copyright (c) 2013-2016 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package math32

//------------------------------------------------------------------------------

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

var result64 float64
var result float32
var resultInt32 int32

//------------------------------------------------------------------------------

func BenchmarkAbs_math(b *testing.B) {
	x := float64(3.3)
	y := float64(-3.3)
	for i := 0; i < b.N; i++ {
		result64 = math.Abs(x)
		result64 = math.Abs(y)
	}
}

//------------------------------------------------------------------------------

func BenchmarkAbs_float32math(b *testing.B) {
	x := float32(3.3)
	y := float32(-3.3)
	for i := 0; i < b.N; i++ {
		result = float32(math.Abs(float64(x)))
		result = float32(math.Abs(float64(y)))
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
		result = switchAbs(x)
		result = switchAbs(y)
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
		result = bitsAbs(x)
		result = bitsAbs(y)
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
		result = unsafeAbs(x)
		result = unsafeAbs(y)
	}
}

//------------------------------------------------------------------------------

func absAsm(x float32) float32

func BenchmarkAbs_asm(b *testing.B) {
	x := float32(3.3)
	y := float32(-3.3)
	for i := 0; i < b.N; i++ {
		result = absAsm(x)
		result = absAsm(y)
	}
}

//------------------------------------------------------------------------------

func ifAbs(x float32) float32 {
	if x < 0 {
		return -x
	}
	if x == 0 {
		return 0 // return correctly abs(-0)
	}
	return x
}

func BenchmarkAbs_if(b *testing.B) {
	x := float32(3.3)
	y := float32(-3.3)
	for i := 0; i < b.N; i++ {
		result = ifAbs(x)
		result = ifAbs(y)
	}
}

//------------------------------------------------------------------------------

func BenchmarkAbs_glam(b *testing.B) {
	x := float32(3.3)
	y := float32(-3.3)
	for i := 0; i < b.N; i++ {
		result = Abs(x)
		result = Abs(y)
	}
}

//------------------------------------------------------------------------------
