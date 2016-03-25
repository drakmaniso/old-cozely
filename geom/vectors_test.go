// Copyright (c) 2013-2016 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package geom_test

import (
	"testing"
	"unsafe"

	"github.com/drakmaniso/glam/geom"
)

//-----------------------------------------------------------------------------

func TestVec4_Creation(t *testing.T) {
	var a geom.Vec4
	if a.X != 0 || a.Y != 0 || a.Z != 0 || a.W != 0 {
		t.Errorf("Zero-initialization failed")
	}
	b := geom.Vec4{X: 1.1, Y: 2.2, Z: 3.3, W: 4.4}
	if b.X != 1.1 || b.Y != 2.2 || b.Z != 3.3 || b.W != 4.4 {
		t.Errorf("Literal initialization failed")
	}
	c := [2]geom.Vec4{{1, 2, 3, 4}, {5, 6, 7, 8}}
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
		t.Errorf("Padding between c[0].Z an c[0].W")
	}
	if uintptr(unsafe.Pointer(&c[0].W))+4 != uintptr(unsafe.Pointer(&c[1].X)) {
		t.Errorf("Padding between c[0].W an c[1].X")
	}
}

func TestVec3_creation(t *testing.T) {
	var a geom.Vec3
	if a.X != 0 || a.Y != 0 || a.Z != 0 {
		t.Errorf("Zero-initialization failed")
	}
	b := geom.Vec3{1.1, 2.2, 3.3}
	if b.X != 1.1 || b.Y != 2.2 || b.Z != 3.3 {
		t.Errorf("Literal initialization failed")
	}
	c := [2]geom.Vec3{{1, 2, 3}, {4, 5, 6}}
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

//-----------------------------------------------------------------------------

func TestVec4_Dehomogenized(t *testing.T) {
	a := geom.Vec4{X: 1.1, Y: 2.2, Z: 3.3, W: 4.4}
	b := a.Dehomogenized()
	if b.X != 0.25 || b.Y != 0.5 || b.Z != 0.75 {
		t.Errorf("Wrong result: %#v", b)
	}
	if a.X != 1.1 || a.Y != 2.2 || a.Z != 3.3 || a.W != 4.4 {
		t.Errorf("First operand modified")
	}
}

func TestVec3_Homogenized(t *testing.T) {
	a := geom.Vec3{1.1, 2.2, 3.3}
	b := a.Homogenized()
	if b.X != 1.1 || b.Y != 2.2 || b.Z != 3.3 || b.W != 1.0 {
		t.Errorf("Wrong result: %#v", b)
	}
	if a.X != 1.1 || a.Y != 2.2 || a.Z != 3.3 {
		t.Errorf("First operand modified")
	}
}

func TestVec3_Dehomogenized(t *testing.T) {
	a := geom.Vec3{1.1, 2.2, 3.3}
	b := a.Dehomogenized()
	if b.X != 0.33333334 || b.Y != 0.6666667 {
		t.Errorf("Wrong result: %#v", b)
	}
	if a.X != 1.1 || a.Y != 2.2 || a.Z != 3.3 {
		t.Errorf("First operand modified")
	}
}

//-----------------------------------------------------------------------------

func TestVec4_Plus(t *testing.T) {
	a := geom.Vec4{X: 1.1, Y: 2.2, Z: 3.3, W: 4.4}
	b := geom.Vec4{X: 5.5, Y: 6.6, Z: 7.7, W: 8.8}
	c := a.Plus(b)
	if c.X != 6.6 || c.Y != 8.8 || c.Z != 11 || c.W != 13.200001 {
		t.Errorf("Wrong result: %#v", c)
	}
	if a.X != 1.1 || a.Y != 2.2 || a.Z != 3.3 || a.W != 4.4 {
		t.Errorf("First operand modified")
	}
}

func TestVec3_Plus(t *testing.T) {
	a := geom.Vec3{1.1, 2.2, 3.3}
	b := geom.Vec3{4.4, 5.5, 6.6}
	c := a.Plus(b)
	if c.X != 5.5 || c.Y != 7.7 || c.Z != 9.9 {
		t.Errorf("Wrong result: %#v", c)
	}
	if a.X != 1.1 || a.Y != 2.2 || a.Z != 3.3 {
		t.Errorf("First operand modified")
	}
}

//-----------------------------------------------------------------------------

func TestVec4_Minus(t *testing.T) {
	a := geom.Vec4{X: 1.1, Y: 2.2, Z: 3.3, W: 4.4}
	b := geom.Vec4{X: 5.5, Y: 6.6, Z: 7.7, W: 8.8}
	c := a.Minus(b)
	if c.X != -4.4 || c.Y != -4.3999996 || c.Z != -4.3999996 || c.W != -4.4 {
		t.Errorf("Wrong result: %#v", c)
	}
	if a.X != 1.1 || a.Y != 2.2 || a.Z != 3.3 || a.W != 4.4 {
		t.Errorf("First operand modified")
	}
}

func TestVec3_Minus(t *testing.T) {
	a := geom.Vec3{1.1, 2.2, 3.3}
	b := geom.Vec3{4.4, 5.5, 6.6}
	c := a.Minus(b)
	if c.X != -3.3000002 || c.Y != -3.3 || c.Z != -3.3 {
		t.Errorf("Wrong result: %#v", c)
	}
	if a.X != 1.1 || a.Y != 2.2 || a.Z != 3.3 {
		t.Errorf("First operand modified")
	}
}

//-----------------------------------------------------------------------------

func TestVec4_Inverse(t *testing.T) {
	a := geom.Vec4{X: 1.1, Y: 2.2, Z: 3.3, W: 4.4}
	b := a.Inverse()
	if b.X != -1.1 || b.Y != -2.2 || b.Z != -3.3 || b.W != -4.4 {
		t.Errorf("Wrong result: %#v", b)
	}
	if a.X != 1.1 || a.Y != 2.2 || a.Z != 3.3 || a.W != 4.4 {
		t.Errorf("First operand modified")
	}
}

func TestVec3_Inverse(t *testing.T) {
	a := geom.Vec3{1.1, 2.2, 3.3}
	b := a.Inverse()
	if b.X != -1.1 || b.Y != -2.2 || b.Z != -3.3 {
		t.Errorf("Wrong result: %#v", b)
	}
	if a.X != 1.1 || a.Y != 2.2 || a.Z != 3.3 {
		t.Errorf("First operand modified")
	}
}

//-----------------------------------------------------------------------------

func TestVec4_Times(t *testing.T) {
	a := geom.Vec4{X: 1.1, Y: 2.2, Z: 3.3, W: 4.4}
	b := a.Times(5.5)
	if b.X != 6.05 || b.Y != 12.1 || b.Z != 18.15 || b.W != 24.2 {
		t.Errorf("Wrong result: %#v", b)
	}
	if a.X != 1.1 || a.Y != 2.2 || a.Z != 3.3 || a.W != 4.4 {
		t.Errorf("First operand modified")
	}
}

func TestVec3_Times(t *testing.T) {
	a := geom.Vec3{1.1, 2.2, 3.3}
	b := a.Times(4.4)
	if b.X != 4.84 || b.Y != 9.68 || b.Z != 14.52 {
		t.Errorf("Wrong result: %#v", b)
	}
	if a.X != 1.1 || a.Y != 2.2 || a.Z != 3.3 {
		t.Errorf("First operand modified")
	}
}

//-----------------------------------------------------------------------------

func TestVec4_Slash(t *testing.T) {
	a := geom.Vec4{X: 1.1, Y: 2.2, Z: 3.3, W: 4.4}
	b := a.Slash(5.5)
	if b.X != 0.2 || b.Y != 0.4 || b.Z != 0.59999996 || b.W != 0.8 {
		t.Errorf("Wrong result: %#v", b)
	}
	if a.X != 1.1 || a.Y != 2.2 || a.Z != 3.3 || a.W != 4.4 {
		t.Errorf("First operand modified")
	}
}

func TestVec3_Slash(t *testing.T) {
	a := geom.Vec3{1.1, 2.2, 3.3}
	b := a.Slash(4.4)
	if b.X != 0.25 || b.Y != 0.5 || b.Z != 0.75 {
		t.Errorf("Wrong result: %#v", b)
	}
	if a.X != 1.1 || a.Y != 2.2 || a.Z != 3.3 {
		t.Errorf("First operand modified")
	}
}

//-----------------------------------------------------------------------------

func TestVec4_Dot(t *testing.T) {
	a := geom.Vec4{X: 1.1, Y: 2.2, Z: 3.3, W: 4.4}
	b := geom.Vec4{X: 5.5, Y: 6.6, Z: 7.7, W: 8.8}
	c := a.Dot(b)
	if c != 84.7 {
		t.Errorf("Wrong result: %#v", c)
	}
	if a.X != 1.1 || a.Y != 2.2 || a.Z != 3.3 || a.W != 4.4 {
		t.Errorf("First operand modified")
	}
}

func TestVec3_Dot(t *testing.T) {
	a := geom.Vec3{1.1, 2.2, 3.3}
	b := geom.Vec3{4.4, 5.5, 6.6}
	c := a.Dot(b)
	if c != 38.72 {
		t.Errorf("Wrong result: %#v", c)
	}
	if a.X != 1.1 || a.Y != 2.2 || a.Z != 3.3 {
		t.Errorf("First operand modified")
	}
}

//-----------------------------------------------------------------------------

func TestVec3_Cross(t *testing.T) {
	a := geom.Vec3{1.1, 2.2, 3.3}
	b := geom.Vec3{4.4, 5.5, 6.6}
	c := a.Cross(b)
	if c.X != -3.6299992 || c.Y != 7.26 || c.Z != -3.63 {
		t.Errorf("Wrong result: %#v", c)
	}
	if a.X != 1.1 || a.Y != 2.2 || a.Z != 3.3 {
		t.Errorf("First operand modified")
	}
}

//-----------------------------------------------------------------------------

func TestVec4_Length(t *testing.T) {
	a := geom.Vec4{X: 1.1, Y: 2.2, Z: 3.3, W: 4.4}
	b := a.Length()
	if b != 6.024948 {
		t.Errorf("Wrong result: %#v", b)
	}
	if a.X != 1.1 || a.Y != 2.2 || a.Z != 3.3 || a.W != 4.4 {
		t.Errorf("First operand modified")
	}
}

func TestVec3_Length(t *testing.T) {
	a := geom.Vec3{1.1, 2.2, 3.3}
	b := a.Length()
	if b != 4.115823 {
		t.Errorf("Wrong result: %#v", b)
	}
	if a.X != 1.1 || a.Y != 2.2 || a.Z != 3.3 {
		t.Errorf("First operand modified")
	}
}

//-----------------------------------------------------------------------------

func TestVec4_Normalized(t *testing.T) {
	a := geom.Vec4{X: 1.1, Y: 2.2, Z: 3.3, W: 4.4}
	b := a.Normalized()
	if b.X != 0.18257418 || b.Y != 0.36514837 || b.Z != 0.5477226 || b.W != 0.73029673 {
		t.Errorf("Wrong result: %#v", b)
	}
}

func TestVec3_Normalized(t *testing.T) {
	a := geom.Vec3{1.1, 2.2, 3.3}
	b := a.Normalized()
	if b.X != 0.26726127 || b.Y != 0.53452253 || b.Z != 0.8017838 {
		t.Errorf("Wrong result: %#v", b)
	}
}

//-----------------------------------------------------------------------------
