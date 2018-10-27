package coord_test

import (
	"testing"
	"unsafe"

	"github.com/cozely/cozely/coord"
)

////////////////////////////////////////////////////////////////////////////////

func TestCoord_creation(t *testing.T) {
	var a coord.XYZ
	if a.X != 0 || a.Y != 0 || a.Z != 0 {
		t.Errorf("Zero-initialization failed")
	}
	b := coord.XYZ{1.1, 2.2, 3.3}
	if b.X != 1.1 || b.Y != 2.2 || b.Z != 3.3 {
		t.Errorf("Literal initialization failed")
	}
	c := [2]coord.XYZ{{1, 2, 3}, {4, 5, 6}}
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

////////////////////////////////////////////////////////////////////////////////

func TestHomogen_Dehomogenized(t *testing.T) {
	a := coord.XYZW{1.1, 2.2, 3.3, 4.4}
	b := a.XYZ()
	if b.X != 0.25 || b.Y != 0.5 || b.Z != 0.75 {
		t.Errorf("Wrong result: %#v", b)
	}
	if a.X != 1.1 || a.Y != 2.2 || a.Z != 3.3 || a.W != 4.4 {
		t.Errorf("First operand modified")
	}
}

func TestCoord_Homogenized(t *testing.T) {
	a := coord.XYZ{1.1, 2.2, 3.3}
	b := a.XYZW(1)
	if b.X != 1.1 || b.Y != 2.2 || b.Z != 3.3 || b.W != 1.0 {
		t.Errorf("Wrong result: %#v", b)
	}
	if a.X != 1.1 || a.Y != 2.2 || a.Z != 3.3 {
		t.Errorf("First operand modified")
	}
}

////////////////////////////////////////////////////////////////////////////////

func TestCoord_Plus(t *testing.T) {
	a := coord.XYZ{1.1, 2.2, 3.3}
	b := coord.XYZ{4.4, 5.5, 6.6}
	c := a.Plus(b)
	if c.X != 5.5 || c.Y != 7.7 || c.Z != 9.9 {
		t.Errorf("Wrong result: %#v", c)
	}
	if a.X != 1.1 || a.Y != 2.2 || a.Z != 3.3 {
		t.Errorf("First operand modified")
	}
}

////////////////////////////////////////////////////////////////////////////////

func TestCoord_Minus(t *testing.T) {
	a := coord.XYZ{1.1, 2.2, 3.3}
	b := coord.XYZ{4.4, 5.5, 6.6}
	c := a.Minus(b)
	if c.X != -3.3000002 || c.Y != -3.3 || c.Z != -3.3 {
		t.Errorf("Wrong result: %#v", c)
	}
	if a.X != 1.1 || a.Y != 2.2 || a.Z != 3.3 {
		t.Errorf("First operand modified")
	}
}

////////////////////////////////////////////////////////////////////////////////

func TestCoord_Inverse(t *testing.T) {
	a := coord.XYZ{1.1, 2.2, 3.3}
	b := a.Opposite()
	if b.X != -1.1 || b.Y != -2.2 || b.Z != -3.3 {
		t.Errorf("Wrong result: %#v", b)
	}
	if a.X != 1.1 || a.Y != 2.2 || a.Z != 3.3 {
		t.Errorf("First operand modified")
	}
}

////////////////////////////////////////////////////////////////////////////////

func TestCoord_Times(t *testing.T) {
	a := coord.XYZ{1.1, 2.2, 3.3}
	b := a.Timess(4.4)
	if b.X != 4.84 || b.Y != 9.68 || b.Z != 14.52 {
		t.Errorf("Wrong result: %#v", b)
	}
	if a.X != 1.1 || a.Y != 2.2 || a.Z != 3.3 {
		t.Errorf("First operand modified")
	}
}

////////////////////////////////////////////////////////////////////////////////

func TestCoord_Slash(t *testing.T) {
	a := coord.XYZ{1.1, 2.2, 3.3}
	b := a.Slashs(4.4)
	if b.X != 0.25 || b.Y != 0.5 || b.Z != 0.75 {
		t.Errorf("Wrong result: %#v", b)
	}
	if a.X != 1.1 || a.Y != 2.2 || a.Z != 3.3 {
		t.Errorf("First operand modified")
	}
}

////////////////////////////////////////////////////////////////////////////////

func TestCoord_Dot(t *testing.T) {
	a := coord.XYZ{1.1, 2.2, 3.3}
	b := coord.XYZ{4.4, 5.5, 6.6}
	c := a.Dot(b)
	if c != 38.72 {
		t.Errorf("Wrong result: %#v", c)
	}
	if a.X != 1.1 || a.Y != 2.2 || a.Z != 3.3 {
		t.Errorf("First operand modified")
	}
}

////////////////////////////////////////////////////////////////////////////////

func TestCoord_Cross(t *testing.T) {
	a := coord.XYZ{1.1, 2.2, 3.3}
	b := coord.XYZ{4.4, 5.5, 6.6}
	c := a.Cross(b)
	if c.X != -3.6299992 || c.Y != 7.26 || c.Z != -3.63 {
		t.Errorf("Wrong result: %#v", c)
	}
	if a.X != 1.1 || a.Y != 2.2 || a.Z != 3.3 {
		t.Errorf("First operand modified")
	}
}

////////////////////////////////////////////////////////////////////////////////

func TestCoord_Length(t *testing.T) {
	a := coord.XYZ{1.1, 2.2, 3.3}
	b := a.Length()
	if b != 4.115823 {
		t.Errorf("Wrong result: %#v", b)
	}
	if a.X != 1.1 || a.Y != 2.2 || a.Z != 3.3 {
		t.Errorf("First operand modified")
	}
}

////////////////////////////////////////////////////////////////////////////////

func TestCoord_Normalized(t *testing.T) {
	a := coord.XYZ{1.1, 2.2, 3.3}
	b := a.Normalized()
	if b.X != 0.26726127 || b.Y != 0.53452253 || b.Z != 0.8017838 {
		t.Errorf("Wrong result: %#v", b)
	}
}

// Copyright (c) 2013-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).
