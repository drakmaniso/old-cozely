// Copyright (c) 2013 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package math

import (
	"math"
	"testing"
)

//------------------------------------------------------------------------------

func TestTan(t *testing.T) {
	const epsilon = float32(1e-7)
	a := Tan(0)
	if a != 0 {
		t.Errorf("Wrong result for Tan(0): %v\n", a)
	}
	b := Tan(0.2)
	if !NearlyEqual(b, 0.202710035509, epsilon) {
		t.Errorf("Wrong result for Tan(0.2): %v\n", b)
	}
	c := Tan(0.5)
	if !NearlyEqual(c, 0.546302489844, epsilon) {
		t.Errorf("Wrong result for Tan(0.5): %v\n", c)
	}
	d := Tan(Pi / 4)
	if !NearlyEqual(d, 1, epsilon) {
		t.Errorf("Wrong result for Tan(Pi/4): %v\n", d)
	}
	e := Tan(Pi / 3)
	if !NearlyEqual(e, 1.73205080757, epsilon) {
		t.Errorf("Wrong result for Tan(Pi/3): %v\n", e)
	}
	f := Tan(Pi/2 - 0.1)
	if !NearlyEqual(f, 9.96664442326, epsilon) {
		t.Errorf("Wrong result for Tan(Pi/2 - 0.1): %v\n", f)
	}
	g := Tan(-0.2)
	if !NearlyEqual(g, -0.202710035509, epsilon) {
		t.Errorf("Wrong result for Tan(-0.2): %v\n", g)
	}
	h := Tan(-0.5)
	if !NearlyEqual(h, -0.546302489844, epsilon) {
		t.Errorf("Wrong result for Tan(-0.5): %v\n", h)
	}
	i := Tan(-Pi / 4)
	if !NearlyEqual(i, -1, epsilon) {
		t.Errorf("Wrong result for Tan(-Pi/4): %v\n", i)
	}
	j := Tan(-Pi / 3)
	if !NearlyEqual(j, -1.73205080757, epsilon) {
		t.Errorf("Wrong result for Tan(-Pi/3): %v\n", j)
	}
	k := Tan(-Pi/2 + 0.1)
	if !NearlyEqual(k, -9.96664442326, epsilon) {
		t.Errorf("Wrong result for Tan(-Pi/2 + 0.1): %v\n", k)
	}
}

//------------------------------------------------------------------------------

func BenchmarkTan_math64(b *testing.B) {
	a := float64(0.5)
	for i := 0; i < b.N; i++ {
		_ = math.Tan(a)
	}
}

//------------------------------------------------------------------------------

func BenchmarkTan_math32(b *testing.B) {
	a := float32(0.5)
	for i := 0; i < b.N; i++ {
		_ = float32(math.Tan(float64(a)))
	}
}

//------------------------------------------------------------------------------

func BenchmarkTan_glm(b *testing.B) {
	a := float32(0.5)
	for i := 0; i < b.N; i++ {
		_ = Tan(a)
	}
}

//------------------------------------------------------------------------------
