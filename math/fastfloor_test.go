// Copyright (c) 2013 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package math

import (
	"testing"
)

//------------------------------------------------------------------------------

func TestFastFloor(t *testing.T) {
	a := FastFloor(float32(3.3))
	if a != 3 {
		t.Errorf("Wrong result for FastFloor(3.3): %v %T", a, a)
	}
	b := FastFloor(float32(-3.3))
	if b != -4 {
		t.Errorf("Wrong result for FastFloor(-3.3): %v %T", b, b)
	}
	c := FastFloor(float32(3))
	if c != 3 {
		t.Errorf("Wrong result for FastFloor(3): %v %T", c, c)
	}
	d := FastFloor(float32(-3))
	if d != -4 {
		t.Errorf("Wrong result for FastFloor(-3): %v %T", d, d)
	}
}

//------------------------------------------------------------------------------

func castFastFloor(x float32) int32 {
	if x > 0 {
		return int32(x)
	} else {
		return int32(x - 1)
	}
}

func BenchmarkFastFloor_cast(b *testing.B) {
	x := float32(3.3)
	y := float32(-3.3)
	for i := 0; i < b.N; i++ {
		_ = castFastFloor(x)
		_ = castFastFloor(y)
	}
}

//------------------------------------------------------------------------------

func asmFastFloor(s float32) int32

func BenchmarkFastFloor_asm(b *testing.B) {
	x := float32(3.3)
	y := float32(-3.3)
	for i := 0; i < b.N; i++ {
		_ = asmFastFloor(x)
		_ = asmFastFloor(y)
	}
}

//------------------------------------------------------------------------------

func BenchmarkFastFloor_glm(b *testing.B) {
	x := float32(3.3)
	y := float32(-3.3)
	for i := 0; i < b.N; i++ {
		_ = FastFloor(x)
		_ = FastFloor(y)
	}
}

//------------------------------------------------------------------------------
