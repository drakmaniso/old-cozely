// Copyright (c) 2013 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package glm

import (
	"fmt"
	"testing"
	"unsafe"
)

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
	d.Add(Vec4{5.5, 6.6, 7.7, 8.8})
	e := b.Slash(2.2)
	f := e.Dehomogenized()
	g := b
	g.Normalize()

	fmt.Printf("a == %#v\n", a)
	fmt.Printf("b == %#v\n", b)
	fmt.Printf("c == %#v\n", c)
	fmt.Printf("d == %#v\n", d)
	fmt.Printf("e == %#v\n", e)
	fmt.Printf("f == %#v\n", f)
	fmt.Printf("g == %#v\n", g)
	// Output:
	// a == glm.Vec4{X:0, Y:0, Z:0, W:0}
	// b == glm.Vec4{X:1.1, Y:2.2, Z:3.3, W:4.4}
	// c == glm.Vec4{X:6.6, Y:8.8, Z:11, W:13.200001}
	// d == glm.Vec4{X:6.6, Y:8.8, Z:11, W:13.200001}
	// e == glm.Vec4{X:0.5, Y:1, Z:1.5, W:2}
	// f == glm.Vec3{X:0.25, Y:0.5, Z:0.75}
	// g == glm.Vec4{X:0.18257418, Y:0.36514837, Z:0.5477226, W:0.73029673}
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

func TestVec4_Add(t *testing.T) {
	a := Vec4{1.1, 2.2, 3.3, 4.4}
	b := Vec4{5.5, 6.6, 7.7, 8.8}
	a.Add(b)
	if a.X != 6.6 || a.Y != 8.8 || a.Z != 11 || a.W != 13.200001 {
		t.Errorf("Wrong result: %#v", a)
	}
}

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

//-----------------------------------------------------------------------------

func TestVec4_Subtract(t *testing.T) {
	a := Vec4{1.1, 2.2, 3.3, 4.4}
	b := Vec4{5.5, 6.6, 7.7, 8.8}
	a.Subtract(b)
	if a.X != -4.4 || a.Y != -4.3999996 || a.Z != -4.3999996 || a.W != -4.4 {
		t.Errorf("Wrong result: %#v", a)
	}
}

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

func TestVec4_Invert(t *testing.T) {
	a := Vec4{1.1, 2.2, 3.3, 4.4}
	a.Invert()
	if a.X != -1.1 || a.Y != -2.2 || a.Z != -3.3 || a.W != -4.4 {
		t.Errorf("Wrong result: %#v", a)
	}
}

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

func TestVec4_Multiply(t *testing.T) {
	a := Vec4{1.1, 2.2, 3.3, 4.4}
	a.Multiply(5.5)
	if a.X != 6.05 || a.Y != 12.1 || a.Z != 18.15 || a.W != 24.2 {
		t.Errorf("Wrong result: %#v", a)
	}
}

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

//-----------------------------------------------------------------------------

func TestVec4_Divide(t *testing.T) {
	a := Vec4{1.1, 2.2, 3.3, 4.4}
	a.Divide(5.5)
	if a.X != 0.2 || a.Y != 0.4 || a.Z != 0.59999996 || a.W != 0.8 {
		t.Errorf("Wrong result: %#v", a)
	}
}

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

func TestVec4_Normalize(t *testing.T) {
	a := Vec4{1.1, 2.2, 3.3, 4.4}
	a.Normalize()
	if a.X != 0.18257418 || a.Y != 0.36514837 || a.Z != 0.5477226 || a.W != 0.73029673 {
		t.Errorf("Wrong result: %#v", a)
	}
}

//-----------------------------------------------------------------------------
