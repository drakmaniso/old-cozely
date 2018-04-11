// Copyright (c) 2013-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package space_test

import (
	"testing"
	"unsafe"

	"github.com/drakmaniso/cozely/x/space"
)

//------------------------------------------------------------------------------

func TestMatrix_ZeroValue(t *testing.T) {
	var m space.Matrix
	if m[0][0] != 0 || m[1][0] != 0 || m[2][0] != 0 || m[3][0] != 0 ||
		m[0][1] != 0 || m[1][1] != 0 || m[2][1] != 0 || m[3][1] != 0 ||
		m[0][2] != 0 || m[1][2] != 0 || m[2][2] != 0 || m[3][2] != 0 ||
		m[0][3] != 0 || m[1][3] != 0 || m[2][3] != 0 || m[3][3] != 0 {
		t.Errorf("Not zeroed: %v", m)
	}
}

func TestMatrix_Literal(t *testing.T) {
	m := space.Matrix{
		{1.1, 2.1, 3.1, 4.1},
		{1.2, 2.2, 3.2, 4.2},
		{1.3, 2.3, 3.3, 4.3},
		{1.4, 2.4, 3.4, 4.4},
	}
	if m[0][0] != 1.1 || m[1][0] != 1.2 || m[2][0] != 1.3 || m[3][0] != 1.4 ||
		m[0][1] != 2.1 || m[1][1] != 2.2 || m[2][1] != 2.3 || m[3][1] != 2.4 ||
		m[0][2] != 3.1 || m[1][2] != 3.2 || m[2][2] != 3.3 || m[3][2] != 3.4 ||
		m[0][3] != 4.1 || m[1][3] != 4.2 || m[2][3] != 4.3 || m[3][3] != 4.4 {
		t.Errorf("Not initialized: %v", m)
	}
}

func TestMatrix_Allocation(t *testing.T) {
	m := &space.Matrix{
		{1.1, 2.1, 3.1, 4.1},
		{1.2, 2.2, 3.2, 4.2},
		{1.3, 2.3, 3.3, 4.3},
		{1.4, 2.4, 3.4, 4.4},
	}
	if m[0][0] != 1.1 || m[1][0] != 1.2 || m[2][0] != 1.3 || m[3][0] != 1.4 ||
		m[0][1] != 2.1 || m[1][1] != 2.2 || m[2][1] != 2.3 || m[3][1] != 2.4 ||
		m[0][2] != 3.1 || m[1][2] != 3.2 || m[2][2] != 3.3 || m[3][2] != 3.4 ||
		m[0][3] != 4.1 || m[1][3] != 4.2 || m[2][3] != 4.3 || m[3][3] != 4.4 {
		t.Errorf("Not allocated correctly: %v", m)
	}
}

func TestMatrix_Assignment(t *testing.T) {
	var m space.Matrix
	_ = m
	m = space.Matrix{
		{1.1, 2.1, 3.1, 4.1},
		{1.2, 2.2, 3.2, 4.2},
		{1.3, 2.3, 3.3, 4.3},
		{1.4, 2.4, 3.4, 4.4},
	}
	if m[0][0] != 1.1 || m[1][0] != 1.2 || m[2][0] != 1.3 || m[3][0] != 1.4 ||
		m[0][1] != 2.1 || m[1][1] != 2.2 || m[2][1] != 2.3 || m[3][1] != 2.4 ||
		m[0][2] != 3.1 || m[1][2] != 3.2 || m[2][2] != 3.3 || m[3][2] != 3.4 ||
		m[0][3] != 4.1 || m[1][3] != 4.2 || m[2][3] != 4.3 || m[3][3] != 4.4 {
		t.Errorf("Not assigned correctly: %v", m)
	}
}

func TestMatrix_ArrayPadding(t *testing.T) {
	a := [2]space.Matrix{
		{
			{1.1, 2.1, 3.1, 4.1},
			{1.2, 2.2, 3.2, 4.2},
			{1.3, 2.3, 3.3, 4.3},
			{1.4, 2.4, 3.4, 4.4},
		},
		{
			{10.1, 20.1, 30.1, 40.1},
			{10.2, 20.2, 30.2, 40.2},
			{10.3, 20.3, 30.3, 40.3},
			{10.4, 20.4, 30.4, 40.4},
		},
	}

	if unsafe.Pointer(&a) != unsafe.Pointer(&a[0][0][0]) {
		t.Errorf("Padding before arrmat[0][0][0]\n")
	}
	if uintptr(unsafe.Pointer(&a[0][0][0]))+4 != uintptr(unsafe.Pointer(&a[0][0][1])) {
		t.Errorf("Padding before arrmat[0][0][1]\n")
	}
	if uintptr(unsafe.Pointer(&a[0][0][1]))+4 != uintptr(unsafe.Pointer(&a[0][0][2])) {
		t.Errorf("Padding before arrmat[0][0][2]\n")
	}
	if uintptr(unsafe.Pointer(&a[0][0][2]))+4 != uintptr(unsafe.Pointer(&a[0][0][3])) {
		t.Errorf("Padding before arrmat[0][0][3]\n")
	}
	if uintptr(unsafe.Pointer(&a[0][0][3]))+4 != uintptr(unsafe.Pointer(&a[0][1][0])) {
		t.Errorf("Padding before arrmat[0][1][0]\n")
	}

	if uintptr(unsafe.Pointer(&a[0][1][0]))+4 != uintptr(unsafe.Pointer(&a[0][1][1])) {
		t.Errorf("Padding before arrmat[0][1][1]\n")
	}
	if uintptr(unsafe.Pointer(&a[0][1][1]))+4 != uintptr(unsafe.Pointer(&a[0][1][2])) {
		t.Errorf("Padding before arrmat[0][1][2]\n")
	}
	if uintptr(unsafe.Pointer(&a[0][1][2]))+4 != uintptr(unsafe.Pointer(&a[0][1][3])) {
		t.Errorf("Padding before arrmat[0][1][3]\n")
	}
	if uintptr(unsafe.Pointer(&a[0][1][3]))+4 != uintptr(unsafe.Pointer(&a[0][2][0])) {
		t.Errorf("Padding before arrmat[0][2][0]\n")
	}

	if uintptr(unsafe.Pointer(&a[0][2][0]))+4 != uintptr(unsafe.Pointer(&a[0][2][1])) {
		t.Errorf("Padding before arrmat[0][2][1]\n")
	}
	if uintptr(unsafe.Pointer(&a[0][2][1]))+4 != uintptr(unsafe.Pointer(&a[0][2][2])) {
		t.Errorf("Padding before arrmat[0][2][2]\n")
	}
	if uintptr(unsafe.Pointer(&a[0][2][2]))+4 != uintptr(unsafe.Pointer(&a[0][2][3])) {
		t.Errorf("Padding before arrmat[0][2][3]\n")
	}
	if uintptr(unsafe.Pointer(&a[0][2][3]))+4 != uintptr(unsafe.Pointer(&a[0][3][0])) {
		t.Errorf("Padding before arrmat[0][3][0]\n")
	}

	if uintptr(unsafe.Pointer(&a[0][3][0]))+4 != uintptr(unsafe.Pointer(&a[0][3][1])) {
		t.Errorf("Padding before arrmat[0][3][1]\n")
	}
	if uintptr(unsafe.Pointer(&a[0][3][1]))+4 != uintptr(unsafe.Pointer(&a[0][3][2])) {
		t.Errorf("Padding before arrmat[0][3][2]\n")
	}
	if uintptr(unsafe.Pointer(&a[0][3][2]))+4 != uintptr(unsafe.Pointer(&a[0][3][3])) {
		t.Errorf("Padding before arrmat[0][3][3]\n")
	}
	if uintptr(unsafe.Pointer(&a[0][3][3]))+4 != uintptr(unsafe.Pointer(&a[1][0][0])) {
		t.Errorf("Padding before arrmat[1][0][0]\n")
	}
}

//------------------------------------------------------------------------------

func TestMatrix_Times(t *testing.T) {
	a := space.Matrix{
		{1.1, 2.1, 3.1, 4.1},
		{1.2, 2.2, 3.2, 4.2},
		{1.3, 2.3, 3.3, 4.3},
		{1.4, 2.4, 3.4, 4.4},
	}
	b := space.Matrix{
		{10.1, 20.1, 30.1, 40.1},
		{10.2, 20.2, 30.2, 40.2},
		{10.3, 20.3, 30.3, 40.3},
		{10.4, 20.4, 30.4, 40.4},
	}
	c := a.Times(b)
	d := space.Matrix{
		{130.5, 230.9, 331.3, 431.7},
		{131, 231.80002, 332.60004, 433.40002},
		{131.5, 232.7, 333.90002, 435.1},
		{132, 233.6, 335.2, 436.8},
	}
	if c != d {
		t.Errorf("Wrong result: %v", c)
	}
}

//------------------------------------------------------------------------------
