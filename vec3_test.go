// Copyright (c) 2013 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package glm

import (
	"fmt"
	"testing"
	"unsafe"
)

//-----------------------------------------------------------------------------

func TestCreation(t *testing.T) {
	var a Vec3
	if a.X != 0 || a.Y != 0 || a.Z != 0 {
		t.Errorf("Zero-initialization failed")
	}
	b := Vec3{1, 2.5, 3}
	if b.X != 1 || b.Y != 2.5 || b.Z != 3 {
		t.Errorf("Literal initialization failed")
	}
	c := [2]Vec3{{1, 2, 3}, {4, 5, 6}}
	if unsafe.Pointer(&c) != unsafe.Pointer(&c[0].X) {
		t.Errorf("Padding before c[0].X")
	}
	if uintptr(unsafe.Pointer(&c[0].X))+4 != uintptr(unsafe.Pointer(&c[0].Y)) {
		t.Errorf("Padding between c[0].X an c[0].Y")
	}
	if uintptr(unsafe.Pointer(&c[0].Y))+4 != uintptr(unsafe.Pointer(&c[0].Z)) {
		t.Errorf("Padding between c[0].Y an c[0].Z")
	}
	if uintptr(unsafe.Pointer(&c[0].Z))+4 != uintptr(unsafe.Pointer(&c[1].X)) {
		t.Errorf("Padding between c[0].Z an c[1].X")
	}
}

func ExampleVec3() {
	var a Vec3
	fmt.Printf("a == %#v\n", a)
	b := Vec3{1, 2.5, 3}
	fmt.Printf("b == %#v\n", b)
	// Output: 
	// a == glm.Vec3{X:0, Y:0, Z:0}
	// b == glm.Vec3{X:1, Y:2.5, Z:3}
}

//-----------------------------------------------------------------------------

func TestHomogenized(t *testing.T) {
	a := Vec3{1.1, 2.2, 3.3}
	b := a.Homogenized()
	if b.X != 1.1 || b.Y != 2.2 || b.Z != 3.3 || b.W != 1.0 {
		t.Errorf("Wrong result: %#v", b)
	}
	if a.X != 1.1 || a.Y != 2.2 || a.Z != 3.3 {
		t.Errorf("First operand modified")
	}
}

//-----------------------------------------------------------------------------

func TestHomogenizedAsDirection(t *testing.T) {
	a := Vec3{1.1, 2.2, 3.3}
	b := a.HomogenizedAsDirection()
	if b.X != 1.1 || b.Y != 2.2 || b.Z != 3.3 || b.W != 0.0 {
		t.Errorf("Wrong result: %#v", b)
	}
	if a.X != 1.1 || a.Y != 2.2 || a.Z != 3.3 {
		t.Errorf("First operand modified")
	}
}

//-----------------------------------------------------------------------------

func TestDehomogenized(t *testing.T) {
	a := Vec3{1.1, 2.2, 3.3}
	b := a.Dehomogenized()
	if b.X != 0.33333334 || b.Y != 0.6666667 {
		t.Errorf("Wrong result: %#v", b)
	}
	if a.X != 1.1 || a.Y != 2.2 || a.Z != 3.3 {
		t.Errorf("First operand modified")
	}
}

//-----------------------------------------------------------------------------

func TestAdd(t *testing.T) {
	a := Vec3{1.1, 2.2, 3.3}
	b := Vec3{4.4, 5.5, 6.6}
	a.Add(b)
	if a.X != 5.5 || a.Y != 7.7 || a.Z != 9.9 {
		t.Errorf("Wrong result: %#v", a)
	}
}

func TestPlus(t *testing.T) {
	a := Vec3{1.1, 2.2, 3.3}
	b := Vec3{4.4, 5.5, 6.6}
	c := a.Plus(b)
	if c.X != 5.5 || c.Y != 7.7 || c.Z != 9.9 {
		t.Errorf("Wrong result: %#v", c)
	}
	if a.X != 1.1 || a.Y != 2.2 || a.Z != 3.3 {
		t.Errorf("First operand modified")
	}
}

//-----------------------------------------------------------------------------

func TestSubtract(t *testing.T) {
	a := Vec3{1.1, 2.2, 3.3}
	b := Vec3{4.4, 5.5, 6.6}
	a.Subtract(b)
	if a.X != -3.3000002 || a.Y != -3.3 || a.Z != -3.3 {
		t.Errorf("Wrong result: %#v", a)
	}
}

func TestMinus(t *testing.T) {
	a := Vec3{1.1, 2.2, 3.3}
	b := Vec3{4.4, 5.5, 6.6}
	c := a.Minus(b)
	if c.X != -3.3000002 || c.Y != -3.3 || c.Z != -3.3 {
		t.Errorf("Wrong result: %#v", c)
	}
	if a.X != 1.1 || a.Y != 2.2 || a.Z != 3.3 {
		t.Errorf("First operand modified")
	}
}

//-----------------------------------------------------------------------------

func TestInvert(t *testing.T) {
	a := Vec3{1.1, 2.2, 3.3}
	a.Invert()
	if a.X != -1.1 || a.Y != -2.2 || a.Z != -3.3 {
		t.Errorf("Wrong result: %#v", a)
	}
}

func TestInverse(t *testing.T) {
	a := Vec3{1.1, 2.2, 3.3}
	b := a.Inverse()
	if b.X != -1.1 || b.Y != -2.2 || b.Z != -3.3 {
		t.Errorf("Wrong result: %#v", b)
	}
	if a.X != 1.1 || a.Y != 2.2 || a.Z != 3.3 {
		t.Errorf("First operand modified")
	}
}

//-----------------------------------------------------------------------------

func TestMultiplyBy(t *testing.T) {
	a := Vec3{1.1, 2.2, 3.3}
	a.MultiplyBy(4.4)
	if a.X != 4.84 || a.Y != 9.68 || a.Z != 14.52 {
		t.Errorf("Wrong result: %#v", a)
	}
}

func TestTimes(t *testing.T) {
	a := Vec3{1.1, 2.2, 3.3}
	b := a.Times(4.4)
	if b.X != 4.84 || b.Y != 9.68 || b.Z != 14.52 {
		t.Errorf("Wrong result: %#v", b)
	}
	if a.X != 1.1 || a.Y != 2.2 || a.Z != 3.3 {
		t.Errorf("First operand modified")
	}
}

//-----------------------------------------------------------------------------

func TestDivideBy(t *testing.T) {
	a := Vec3{1.1, 2.2, 3.3}
	a.DivideBy(4.4)
	if a.X != 0.25 || a.Y != 0.5 || a.Z != 0.75 {
		t.Errorf("Wrong result: %#v", a)
	}
}

func TestSlash(t *testing.T) {
	a := Vec3{1.1, 2.2, 3.3}
	b := a.Slash(4.4)
	if b.X != 0.25 || b.Y != 0.5 || b.Z != 0.75 {
		t.Errorf("Wrong result: %#v", b)
	}
	if a.X != 1.1 || a.Y != 2.2 || a.Z != 3.3 {
		t.Errorf("First operand modified")
	}
}

//-----------------------------------------------------------------------------

func TestCross(t *testing.T) {
	a := Vec3{1.1, 2.2, 3.3}
	b := Vec3{4.4, 5.5, 6.6}
	c := a.Cross(b)
	if c.X != -3.6299992 || c.Y != 7.26 || c.Z != -3.63 {
		t.Errorf("Wrong result: %#v", c)
	}
	if a.X != 1.1 || a.Y != 2.2 || a.Z != 3.3 {
		t.Errorf("First operand modified")
	}
}

//-----------------------------------------------------------------------------

func TestDot(t *testing.T) {
	a := Vec3{1.1, 2.2, 3.3}
	b := Vec3{4.4, 5.5, 6.6}
	c := a.Dot(b)
	if c != 38.72 {
		t.Errorf("Wrong result: %#v", c)
	}
	if a.X != 1.1 || a.Y != 2.2 || a.Z != 3.3 {
		t.Errorf("First operand modified")
	}
}

//-----------------------------------------------------------------------------

func TestLength(t *testing.T) {
	a := Vec3{1.1, 2.2, 3.3}
	b := a.Length()
	if b != 4.115823 {
		t.Errorf("Wrong result: %#v", b)
	}
	if a.X != 1.1 || a.Y != 2.2 || a.Z != 3.3 {
		t.Errorf("First operand modified")
	}
}

func TestNormalize(t *testing.T) {
	a := Vec3{1.1, 2.2, 3.3}
	a.Normalize()
	if a.X != 0.26726127 || a.Y != 0.53452253 || a.Z != 0.8017838 {
		t.Errorf("Wrong result: %#v", a)
	}
}

//-----------------------------------------------------------------------------
