// Copyright (c) 2013-2016 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package geom

import (
	"fmt"
	"testing"
	"unsafe"

	"github.com/drakmaniso/glam/math"
)

//-----------------------------------------------------------------------------

func (v *Vec4) add(a, b Vec4) {
	v.X = a.X + b.X
	v.Y = a.Y + b.Y
	v.Z = a.Z + b.Z
	v.W = a.W + b.W
}

func (v *Vec4) subtract(a, b Vec4) {
	v.X = a.X - b.X
	v.Y = a.Y - b.Y
	v.Z = a.Z - b.Z
	v.W = a.W - b.W
}

func (v *Vec4) invert() {
	v.X = -v.X
	v.Y = -v.Y
	v.Z = -v.Z
	v.W = -v.W
}

func (v *Vec4) multiply(o Vec4, s float32) {
	v.X = o.X * s
	v.Y = o.Y * s
	v.Z = o.Z * s
	v.W = o.W * s
}

func (v *Vec4) divide(o Vec4, s float32) {
	v.X = o.X / s
	v.Y = o.Y / s
	v.Z = o.Z / s
	v.W = o.W / s
}

func (v *Vec4) normalize() {
	length := math.Sqrt(v.X*v.X + v.Y*v.Y + v.Z*v.Z + v.W*v.W)
	v.X /= length
	v.Y /= length
	v.Z /= length
	v.W /= length
}

//-----------------------------------------------------------------------------

type arrayVec4 [4]float32

func (v arrayVec4) plus(o arrayVec4) arrayVec4 {
	return arrayVec4{v[0] + o[0], v[1] + o[1], v[2] + o[2], v[3] + o[3]}
}

func (v *arrayVec4) add(a, b arrayVec4) {
	v[0] = a[0] + b[0]
	v[1] = a[1] + b[1]
	v[2] = a[2] + b[2]
	v[3] = a[3] + b[3]
}

//-----------------------------------------------------------------------------

func TestVec4_Creation(t *testing.T) {
	var a Vec4
	if a.X != 0 || a.Y != 0 || a.Z != 0 || a.W != 0 {
		t.Errorf("Zero-initialization failed")
	}
	b := Vec4{1.1, 2.2, 3.3, 4.4}
	if b.X != 1.1 || b.Y != 2.2 || b.Z != 3.3 || b.W != 4.4 {
		t.Errorf("Literal initialization failed")
	}
	c := [2]Vec4{{1, 2, 3, 4}, {5, 6, 7, 8}}
	if unsafe.Pointer(&c) != unsafe.Pointer(&c[0].X) {
		t.Errorf("Padding before c[0].X")
	}
	if uintptr(unsafe.Pointer(&c[0].X))+4 != uintptr(unsafe.Pointer(&c[0].Y)) {
		t.Errorf("Padding between c[0].X an c[0].Y")
	}
	if uintptr(unsafe.Pointer(&c[0].Y))+4 != uintptr(unsafe.Pointer(&c[0].Z)) {
		t.Errorf("Padding between c[0].Y an c[0].Z")
	}
	if uintptr(unsafe.Pointer(&c[0].Z))+4 != uintptr(unsafe.Pointer(&c[0].W)) {
		t.Errorf("Padding between c[0].Z an c[1].X")
	}
	if uintptr(unsafe.Pointer(&c[0].W))+4 != uintptr(unsafe.Pointer(&c[1].X)) {
		t.Errorf("Padding between c[0].Z an c[1].X")
	}
}

func ExampleVec4() {
	var a Vec4
	b := Vec4{1.1, 2.2, 3.3, 4.4}
	c := b.Plus(Vec4{5.5, 6.6, 7.7, 8.8})
	d := b
	d = d.Plus(Vec4{5.5, 6.6, 7.7, 8.8})
	e := b.Slash(2.2)
	f := e.Dehomogenized()
	g := b
	g = g.Normalized()

	fmt.Printf("a == %#v\n", a)
	fmt.Printf("b == %#v\n", b)
	fmt.Printf("c == %#v\n", c)
	fmt.Printf("d == %#v\n", d)
	fmt.Printf("e == %#v\n", e)
	fmt.Printf("f == %#v\n", f)
	fmt.Printf("g == %#v\n", g)
	// Output:
	// a == geom.Vec4{X:0, Y:0, Z:0, W:0}
	// b == geom.Vec4{X:1.1, Y:2.2, Z:3.3, W:4.4}
	// c == geom.Vec4{X:6.6, Y:8.8, Z:11, W:13.200001}
	// d == geom.Vec4{X:6.6, Y:8.8, Z:11, W:13.200001}
	// e == geom.Vec4{X:0.5, Y:1, Z:1.5, W:2}
	// f == geom.Vec3{X:0.25, Y:0.5, Z:0.75}
	// g == geom.Vec4{X:0.18257418, Y:0.36514837, Z:0.5477226, W:0.73029673}
}

//-----------------------------------------------------------------------------

func TestVec4_Dehomogenized(t *testing.T) {
	a := Vec4{1.1, 2.2, 3.3, 4.4}
	b := a.Dehomogenized()
	if b.X != 0.25 || b.Y != 0.5 || b.Z != 0.75 {
		t.Errorf("Wrong result: %#v", b)
	}
	if a.X != 1.1 || a.Y != 2.2 || a.Z != 3.3 || a.W != 4.4 {
		t.Errorf("First operand modified")
	}
}

//-----------------------------------------------------------------------------

func TestVec4_Plus(t *testing.T) {
	a := Vec4{1.1, 2.2, 3.3, 4.4}
	b := Vec4{5.5, 6.6, 7.7, 8.8}
	c := a.Plus(b)
	if c.X != 6.6 || c.Y != 8.8 || c.Z != 11 || c.W != 13.200001 {
		t.Errorf("Wrong result: %#v", c)
	}
	if a.X != 1.1 || a.Y != 2.2 || a.Z != 3.3 || a.W != 4.4 {
		t.Errorf("First operand modified")
	}
}

func BenchmarkVec4_Plus(b *testing.B) {
	m := Vec4{1.1, 2.2, 3.3, 4.4}
	n := Vec4{5.5, 6.6, 7.7, 8.8}
	var o Vec4
	for i := 0; i < b.N; i++ {
		o = m.Plus(n)
	}
	_ = o
}

func BenchmarkVec4_Plus_Array(b *testing.B) {
	m := arrayVec4{1.1, 2.2, 3.3, 4.4}
	n := arrayVec4{5.5, 6.6, 7.7, 8.8}
	var o arrayVec4
	for i := 0; i < b.N; i++ {
		o = m.plus(n)
	}
	_ = o
}

func BenchmarkVec4_Plus_Self(b *testing.B) {
	m := Vec4{1.1, 2.2, 3.3, 4.4}
	n := Vec4{5.5, 6.6, 7.7, 8.8}
	for i := 0; i < b.N; i++ {
		m = m.Plus(n)
	}
	_ = m
}

func BenchmarkVec4_Plus_ArraySelf(b *testing.B) {
	m := arrayVec4{1.1, 2.2, 3.3, 4.4}
	n := arrayVec4{5.5, 6.6, 7.7, 8.8}
	for i := 0; i < b.N; i++ {
		m = m.plus(n)
	}
	_ = m
}

func BenchmarkVec4_Plus_ByRef(b *testing.B) {
	m := Vec4{1.1, 2.2, 3.3, 4.4}
	n := Vec4{5.5, 6.6, 7.7, 8.8}
	var o Vec4
	for i := 0; i < b.N; i++ {
		o.add(m, n)
	}
	_ = o
}

func BenchmarkVec4_Plus_ArrayByRef(b *testing.B) {
	m := arrayVec4{1.1, 2.2, 3.3, 4.4}
	n := arrayVec4{5.5, 6.6, 7.7, 8.8}
	var o arrayVec4
	for i := 0; i < b.N; i++ {
		o.add(m, n)
	}
	_ = o
}

func BenchmarkVec4_Plus_ByRefSelf(b *testing.B) {
	m := Vec4{1.1, 2.2, 3.3, 4.4}
	n := Vec4{5.5, 6.6, 7.7, 8.8}
	for i := 0; i < b.N; i++ {
		m.add(m, n)
	}
	_ = m
}

func BenchmarkVec4_Plus_ArrayByRefSelf(b *testing.B) {
	m := arrayVec4{1.1, 2.2, 3.3, 4.4}
	n := arrayVec4{5.5, 6.6, 7.7, 8.8}
	for i := 0; i < b.N; i++ {
		m.add(m, n)
	}
	_ = m
}

//-----------------------------------------------------------------------------

func TestVec4_Minus(t *testing.T) {
	a := Vec4{1.1, 2.2, 3.3, 4.4}
	b := Vec4{5.5, 6.6, 7.7, 8.8}
	c := a.Minus(b)
	if c.X != -4.4 || c.Y != -4.3999996 || c.Z != -4.3999996 || c.W != -4.4 {
		t.Errorf("Wrong result: %#v", c)
	}
	if a.X != 1.1 || a.Y != 2.2 || a.Z != 3.3 || a.W != 4.4 {
		t.Errorf("First operand modified")
	}
}

//-----------------------------------------------------------------------------

func TestVec4_Inverse(t *testing.T) {
	a := Vec4{1.1, 2.2, 3.3, 4.4}
	b := a.Inverse()
	if b.X != -1.1 || b.Y != -2.2 || b.Z != -3.3 || b.W != -4.4 {
		t.Errorf("Wrong result: %#v", b)
	}
	if a.X != 1.1 || a.Y != 2.2 || a.Z != 3.3 || a.W != 4.4 {
		t.Errorf("First operand modified")
	}
}

//-----------------------------------------------------------------------------

func TestVec4_Times(t *testing.T) {
	a := Vec4{1.1, 2.2, 3.3, 4.4}
	b := a.Times(5.5)
	if b.X != 6.05 || b.Y != 12.1 || b.Z != 18.15 || b.W != 24.2 {
		t.Errorf("Wrong result: %#v", b)
	}
	if a.X != 1.1 || a.Y != 2.2 || a.Z != 3.3 || a.W != 4.4 {
		t.Errorf("First operand modified")
	}
}

func BenchmarkVec4_Times(b *testing.B) {
	m := Vec4{1.1, 2.2, 3.3, 4.4}
	n := float32(5.5)
	var o Vec4
	for i := 0; i < b.N; i++ {
		o = m.Times(n)
	}
	_ = o
}

func BenchmarkVec4_Times_Self(b *testing.B) {
	m := Vec4{1.1, 2.2, 3.3, 4.4}
	n := float32(5.5)
	for i := 0; i < b.N; i++ {
		m = m.Times(n)
	}
	_ = m
}

func BenchmarkVec4_Times_ByRef(b *testing.B) {
	m := Vec4{1.1, 2.2, 3.3, 4.4}
	n := float32(5.5)
	var o Vec4
	for i := 0; i < b.N; i++ {
		o.multiply(m, n)
	}
	_ = o
}

func BenchmarkVec4_Times_ByRefSelf(b *testing.B) {
	m := Vec4{1.1, 2.2, 3.3, 4.4}
	n := float32(5.5)
	for i := 0; i < b.N; i++ {
		m.multiply(m, n)
	}
	_ = m
}

//-----------------------------------------------------------------------------

func TestVec4_Slash(t *testing.T) {
	a := Vec4{1.1, 2.2, 3.3, 4.4}
	b := a.Slash(5.5)
	if b.X != 0.2 || b.Y != 0.4 || b.Z != 0.59999996 || b.W != 0.8 {
		t.Errorf("Wrong result: %#v", b)
	}
	if a.X != 1.1 || a.Y != 2.2 || a.Z != 3.3 || a.W != 4.4 {
		t.Errorf("First operand modified")
	}
}

func BenchmarkVec4_Slash(b *testing.B) {
	m := Vec4{1.1, 2.2, 3.3, 4.4}
	n := float32(5.5)
	var o Vec4
	for i := 0; i < b.N; i++ {
		o = m.Slash(n)
	}
	_ = o
}

func BenchmarkVec4_Slash_Self(b *testing.B) {
	m := Vec4{1.1, 2.2, 3.3, 4.4}
	n := float32(5.5)
	for i := 0; i < b.N; i++ {
		m = m.Slash(n)
	}
	_ = m
}

func BenchmarkVec4_Slash_ByRef(b *testing.B) {
	m := Vec4{1.1, 2.2, 3.3, 4.4}
	n := float32(5.5)
	var o Vec4
	for i := 0; i < b.N; i++ {
		o.divide(m, n)
	}
	_ = o
}

func BenchmarkVec4_Slash_ByRefSelf(b *testing.B) {
	m := Vec4{1.1, 2.2, 3.3, 4.4}
	n := float32(5.5)
	for i := 0; i < b.N; i++ {
		m.divide(m, n)
	}
	_ = m
}

//-----------------------------------------------------------------------------

func TestVec4_Dot(t *testing.T) {
	a := Vec4{1.1, 2.2, 3.3, 4.4}
	b := Vec4{5.5, 6.6, 7.7, 8.8}
	c := a.Dot(b)
	if c != 84.7 {
		t.Errorf("Wrong result: %#v", c)
	}
	if a.X != 1.1 || a.Y != 2.2 || a.Z != 3.3 || a.W != 4.4 {
		t.Errorf("First operand modified")
	}
}

//-----------------------------------------------------------------------------

func TestVec4_Length(t *testing.T) {
	a := Vec4{1.1, 2.2, 3.3, 4.4}
	b := a.Length()
	if b != 6.024948 {
		t.Errorf("Wrong result: %#v", b)
	}
	if a.X != 1.1 || a.Y != 2.2 || a.Z != 3.3 || a.W != 4.4 {
		t.Errorf("First operand modified")
	}
}

//-----------------------------------------------------------------------------

func TestVec4_Normalized(t *testing.T) {
	a := Vec4{1.1, 2.2, 3.3, 4.4}
	b := a.Normalized()
	if b.X != 0.18257418 || b.Y != 0.36514837 || b.Z != 0.5477226 || b.W != 0.73029673 {
		t.Errorf("Wrong result: %#v", b)
	}
}

func BenchmarkVec4_Normalized(b *testing.B) {
	m := Vec4{1.1, 2.2, 3.3, 4.4}
	var o Vec4
	for i := 0; i < b.N; i++ {
		o = m.Normalized()
	}
	_ = o
}

func BenchmarkVec4_Normalized_Self(b *testing.B) {
	m := Vec4{1.1, 2.2, 3.3, 4.4}
	for i := 0; i < b.N; i++ {
		m = m.Normalized()
	}
	_ = m
}

func BenchmarkVec4_Normalize_ByRef(b *testing.B) {
	m := Vec4{1.1, 2.2, 3.3, 4.4}
	for i := 0; i < b.N; i++ {
		m.normalize()
	}
	_ = m
}

//-----------------------------------------------------------------------------
