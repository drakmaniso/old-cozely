// Copyright (c) 2013 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package math

import (
	"testing"
)

//------------------------------------------------------------------------------

func TestRound(t *testing.T) {
	a := Round(float32(3.3))
	if a != 3 {
		t.Errorf("Wrong result for Round(3.3): %v %T", a, a)
	}
	b := Round(float32(-3.3))
	if b != -3 {
		t.Errorf("Wrong result for Round(-3.3): %v %T", b, b)
	}
	c := Round(float32(3))
	if c != 3 {
		t.Errorf("Wrong result for Round(3): %v %T", c, c)
	}
	d := Round(float32(-3))
	if d != -3 {
		t.Errorf("Wrong result for Round(-3): %v %T", d, d)
	}
	e := Round(float32(3.7))
	if e != 4 {
		t.Errorf("Wrong result for Round(3.7): %v %T", e, e)
	}
	f := Round(float32(-3.7))
	if f != -4 {
		t.Errorf("Wrong result for Round(-3.7): %v %T", f, f)
	}
}

//------------------------------------------------------------------------------

func round_cast(x float32) int32 {
	if x > 0 {
		return int32(x + 0.5)
	} else {
		return int32(x - 0.5)
	}
}

func BenchmarkRound_cast(b *testing.B) {
	x := float32(3.3)
	y := float32(-3.3)
	for i := 0; i < b.N; i++ {
		_ = round_cast(x)
		_ = round_cast(y)
	}
}

//------------------------------------------------------------------------------

func round_asm(s float32) float32

func BenchmarkRound_asm(b *testing.B) {
	x := float32(3.3)
	y := float32(-3.3)
	for i := 0; i < b.N; i++ {
		_ = round_asm(x)
		_ = round_asm(y)
	}
}

//------------------------------------------------------------------------------

func BenchmarkRound_glam(b *testing.B) {
	x := float32(3.3)
	y := float32(-3.3)
	for i := 0; i < b.N; i++ {
		_ = Round(x)
		_ = Round(y)
	}
}

//------------------------------------------------------------------------------
