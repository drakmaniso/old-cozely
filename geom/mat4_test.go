// Copyright (c) 2013-2016 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package geom_test

import (
	"testing"
	"unsafe"

	"github.com/drakmaniso/glam/geom"
)

//------------------------------------------------------------------------------

func TestMat4_Creation(t *testing.T) {
	var a geom.Mat4
	if a[0][0] != 0 || a[1][0] != 0 || a[2][0] != 0 || a[3][0] != 0 ||
		a[0][1] != 0 || a[1][1] != 0 || a[2][1] != 0 || a[3][1] != 0 ||
		a[0][2] != 0 || a[1][2] != 0 || a[2][2] != 0 || a[3][2] != 0 ||
		a[0][3] != 0 || a[1][3] != 0 || a[2][3] != 0 || a[3][3] != 0 {
		t.Errorf("Not zeroed: %#v", a)
	}
	b := geom.MakeMat4(
		1.1, 11.1, 111.1, 1111.1,
		2.2, 22.2, 222.2, 2222.2,
		3.3, 33.3, 333.3, 3333.3,
		4.4, 44.4, 444.4, 4444.4,
	)
	if b[0][0] != 1.1 || b[1][0] != 11.1 || b[2][0] != 111.1 || b[3][0] != 1111.1 ||
		b[0][1] != 2.2 || b[1][1] != 22.2 || b[2][1] != 222.2 || b[3][1] != 2222.2 ||
		b[0][2] != 3.3 || b[1][2] != 33.3 || b[2][2] != 333.3 || b[3][2] != 3333.3 ||
		b[0][3] != 4.4 || b[1][3] != 44.4 || b[2][3] != 444.4 || b[3][3] != 4444.4 {
		t.Errorf("Not maked: %#v", b)
	}
	c := geom.NewMat4(
		1.1, 11.1, 111.1, 1111.1,
		2.2, 22.2, 222.2, 2222.2,
		3.3, 33.3, 333.3, 3333.3,
		4.4, 44.4, 444.4, 4444.4,
	)
	if c[0][0] != 1.1 || c[1][0] != 11.1 || c[2][0] != 111.1 || c[3][0] != 1111.1 ||
		c[0][1] != 2.2 || c[1][1] != 22.2 || c[2][1] != 222.2 || c[3][1] != 2222.2 ||
		c[0][2] != 3.3 || c[1][2] != 33.3 || c[2][2] != 333.3 || c[3][2] != 3333.3 ||
		c[0][3] != 4.4 || c[1][3] != 44.4 || c[2][3] != 444.4 || c[3][3] != 4444.4 {
		t.Errorf("Not newed: %#v", c)
	}
	var d geom.Mat4
	d.SetTo(
		1.1, 11.1, 111.1, 1111.1,
		2.2, 22.2, 222.2, 2222.2,
		3.3, 33.3, 333.3, 3333.3,
		4.4, 44.4, 444.4, 4444.4,
	)
	if d[0][0] != 1.1 || d[1][0] != 11.1 || d[2][0] != 111.1 || d[3][0] != 1111.1 ||
		d[0][1] != 2.2 || d[1][1] != 22.2 || d[2][1] != 222.2 || d[3][1] != 2222.2 ||
		d[0][2] != 3.3 || d[1][2] != 33.3 || d[2][2] != 333.3 || d[3][2] != 3333.3 ||
		d[0][3] != 4.4 || d[1][3] != 44.4 || d[2][3] != 444.4 || d[3][3] != 4444.4 {
		t.Errorf("Not set: %#v", d)
	}
	e := geom.Mat4{
		{1.1, 2.2, 3.3, 4.4},
		{11.1, 22.2, 33.3, 44.4},
		{111.1, 222.2, 333.3, 444.4},
		{1111.1, 2222.2, 3333.3, 4444.4},
	}
	if e[0][0] != 1.1 || e[1][0] != 11.1 || e[2][0] != 111.1 || e[3][0] != 1111.1 ||
		e[0][1] != 2.2 || e[1][1] != 22.2 || e[2][1] != 222.2 || e[3][1] != 2222.2 ||
		e[0][2] != 3.3 || e[1][2] != 33.3 || e[2][2] != 333.3 || e[3][2] != 3333.3 ||
		e[0][3] != 4.4 || e[1][3] != 44.4 || e[2][3] != 444.4 || e[3][3] != 4444.4 {
		t.Errorf("Not set: %#v", e)
	}
	arrmat := [2]geom.Mat4{
		{
			{1.1, 2.2, 3.3, 4.4},
			{11.1, 22.2, 33.3, 44.4},
			{111.1, 222.2, 333.3, 444.4},
			{1111.1, 2222.2, 3333.3, 4444.4},
		},
		{
			{5.5, 6.6, 7.7, 8.8},
			{55.5, 66.6, 77.7, 88.8},
			{555.5, 666.6, 777.7, 888.8},
			{5555.5, 6666.6, 7777.7, 8888.8},
		},
	}

	if unsafe.Pointer(&arrmat) != unsafe.Pointer(&arrmat[0][0][0]) {
		t.Errorf("Padding before arrmat[0][0][0]\n")
	}
	if uintptr(unsafe.Pointer(&arrmat[0][0][0]))+4 != uintptr(unsafe.Pointer(&arrmat[0][0][1])) {
		t.Errorf("Padding before arrmat[0][0][1]\n")
	}
	if uintptr(unsafe.Pointer(&arrmat[0][0][1]))+4 != uintptr(unsafe.Pointer(&arrmat[0][0][2])) {
		t.Errorf("Padding before arrmat[0][0][2]\n")
	}
	if uintptr(unsafe.Pointer(&arrmat[0][0][2]))+4 != uintptr(unsafe.Pointer(&arrmat[0][0][3])) {
		t.Errorf("Padding before arrmat[0][0][3]\n")
	}
	if uintptr(unsafe.Pointer(&arrmat[0][0][3]))+4 != uintptr(unsafe.Pointer(&arrmat[0][1][0])) {
		t.Errorf("Padding before arrmat[0][1][0]\n")
	}

	if uintptr(unsafe.Pointer(&arrmat[0][1][0]))+4 != uintptr(unsafe.Pointer(&arrmat[0][1][1])) {
		t.Errorf("Padding before arrmat[0][1][1]\n")
	}
	if uintptr(unsafe.Pointer(&arrmat[0][1][1]))+4 != uintptr(unsafe.Pointer(&arrmat[0][1][2])) {
		t.Errorf("Padding before arrmat[0][1][2]\n")
	}
	if uintptr(unsafe.Pointer(&arrmat[0][1][2]))+4 != uintptr(unsafe.Pointer(&arrmat[0][1][3])) {
		t.Errorf("Padding before arrmat[0][1][3]\n")
	}
	if uintptr(unsafe.Pointer(&arrmat[0][1][3]))+4 != uintptr(unsafe.Pointer(&arrmat[0][2][0])) {
		t.Errorf("Padding before arrmat[0][2][0]\n")
	}

	if uintptr(unsafe.Pointer(&arrmat[0][2][0]))+4 != uintptr(unsafe.Pointer(&arrmat[0][2][1])) {
		t.Errorf("Padding before arrmat[0][2][1]\n")
	}
	if uintptr(unsafe.Pointer(&arrmat[0][2][1]))+4 != uintptr(unsafe.Pointer(&arrmat[0][2][2])) {
		t.Errorf("Padding before arrmat[0][2][2]\n")
	}
	if uintptr(unsafe.Pointer(&arrmat[0][2][2]))+4 != uintptr(unsafe.Pointer(&arrmat[0][2][3])) {
		t.Errorf("Padding before arrmat[0][2][3]\n")
	}
	if uintptr(unsafe.Pointer(&arrmat[0][2][3]))+4 != uintptr(unsafe.Pointer(&arrmat[0][3][0])) {
		t.Errorf("Padding before arrmat[0][3][0]\n")
	}

	if uintptr(unsafe.Pointer(&arrmat[0][3][0]))+4 != uintptr(unsafe.Pointer(&arrmat[0][3][1])) {
		t.Errorf("Padding before arrmat[0][3][1]\n")
	}
	if uintptr(unsafe.Pointer(&arrmat[0][3][1]))+4 != uintptr(unsafe.Pointer(&arrmat[0][3][2])) {
		t.Errorf("Padding before arrmat[0][3][2]\n")
	}
	if uintptr(unsafe.Pointer(&arrmat[0][3][2]))+4 != uintptr(unsafe.Pointer(&arrmat[0][3][3])) {
		t.Errorf("Padding before arrmat[0][3][3]\n")
	}
	if uintptr(unsafe.Pointer(&arrmat[0][3][3]))+4 != uintptr(unsafe.Pointer(&arrmat[1][0][0])) {
		t.Errorf("Padding before arrmat[1][0][0]\n")
	}
}

//------------------------------------------------------------------------------

func TestMat4_StorageInColumnMajorOrder(t *testing.T) {
	a := geom.MakeMat4(
		1.1, 1.2, 1.3, 1.4,
		2.1, 2.2, 2.3, 2.4,
		3.1, 3.2, 3.3, 3.4,
		4.1, 4.2, 4.3, 4.4,
	)
	b := geom.Mat4{
		{1.1, 2.1, 3.1, 4.1},
		{1.2, 2.2, 3.2, 4.2},
		{1.3, 2.3, 3.3, 4.3},
		{1.4, 2.4, 3.4, 4.4},
	}
	if a != b {
		t.Errorf("Storage not in column major order: %#v", a)
	}
}

//------------------------------------------------------------------------------

func TestMat4_Times(t *testing.T) {
	a := geom.MakeMat4(
		1.1, 1.2, 1.3, 1.4,
		2.1, 2.2, 2.3, 2.4,
		3.1, 3.2, 3.3, 3.4,
		4.1, 4.2, 4.3, 4.4,
	)
	b := geom.MakeMat4(
		10.1, 10.2, 10.3, 10.4,
		20.1, 20.2, 20.3, 20.4,
		30.1, 30.2, 30.3, 30.4,
		40.1, 40.2, 40.3, 40.4,
	)
	c := a.Times(&b)
	d := geom.MakeMat4(
		130.5, 131, 131.5, 132,
		230.9, 231.80002, 232.7, 233.6,
		331.3, 332.60004, 333.90002, 335.2,
		431.7, 433.40002, 435.1, 436.8,
	)
	if c != d {
		t.Errorf("Wrong result: %v", c)
	}
}

//------------------------------------------------------------------------------
