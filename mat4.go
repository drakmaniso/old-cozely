// Copyright (c) 2013 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package glam

import "github.com/drakmaniso/glam/math"

//------------------------------------------------------------------------------

type Mat4 [4][4]float32

//------------------------------------------------------------------------------

func NewMat4(
	a, e, i, m,
	b, f, j, n,
	c, g, k, o,
	d, h, l, p float32,
) *Mat4 {
	return &Mat4{
		{a, b, c, d},
		{e, f, g, h},
		{i, j, k, l},
		{m, n, o, p},
	}
}

func MakeMat4(
	a, e, i, m,
	b, f, j, n,
	c, g, k, o,
	d, h, l, p float32,
) Mat4 {
	return Mat4{
		{a, b, c, d},
		{e, f, g, h},
		{i, j, k, l},
		{m, n, o, p},
	}
}

func (self *Mat4) SetTo(
	a, e, i, m,
	b, f, j, n,
	c, g, k, o,
	d, h, l, p float32,
) {
	self[0][0] = a
	self[0][1] = b
	self[0][2] = c
	self[0][3] = d

	self[1][0] = e
	self[1][1] = f
	self[1][2] = g
	self[1][3] = h

	self[2][0] = i
	self[2][1] = j
	self[2][2] = k
	self[2][3] = l

	self[3][0] = m
	self[3][1] = n
	self[3][2] = o
	self[3][3] = p
}

//------------------------------------------------------------------------------

func (self Mat4) At(column, row int) float32 {
	return self[column][row]
}

func (self *Mat4) Set(column, row int, value float32) {
	self[column][row] = value
}

//------------------------------------------------------------------------------

func Perspective(fieldOfView float32, aspectRatio float32, near float32, far float32) Mat4 {
	f := float32(1.0) / math.Tan(fieldOfView/float32(2.0))

	return Mat4{
		{f / aspectRatio, 0, 0, 0},
		{0, f, 0, 0},
		{0, 0, (far + near) / (near - far), -1},
		{0, 0, (2 * far * near) / (near - far), 0},
	}
}

func (m *Mat4) SetToPerspective(fieldOfView float32, aspectRatio float32, near float32, far float32) {
	f := float32(1.0) / math.Tan(fieldOfView/float32(2.0))

	m[0][0] = f / aspectRatio
	m[0][1] = 0
	m[0][2] = 0
	m[0][3] = 0

	m[0][0] = 0
	m[0][1] = f
	m[0][2] = 0
	m[0][3] = 0

	m[0][0] = 0
	m[0][1] = 0
	m[0][2] = (far + near) / (near - far)
	m[0][3] = -1

	m[0][0] = 0
	m[0][1] = 0
	m[0][2] = (2 * far * near) / (near - far)
	m[0][3] = 0
}

//------------------------------------------------------------------------------
