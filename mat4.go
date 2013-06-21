// Copyright (c) 2013 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package glam

import "github.com/drakmaniso/glam/math"

//------------------------------------------------------------------------------

// `Mat4` is a single-precision matrix with 4 columns and 4 rows.
//
// Note: matrices are stored in column-major order, so when writing literals
// remember to use the transpose.
type Mat4 [4][4]float32

//------------------------------------------------------------------------------

// `NewMat4` allocates and returns a new matrix. The elements are stored in
// alphabetical order (column-major order).
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

// `MakeMat4` returns (by value) a matrix. The elements are stored in
// alphabetical order (column-major order).
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

// `SetTo` initializes `matrix`. The elements are stored in
// alphabetical order (column-major order).
func (matrix *Mat4) SetTo(
	a, e, i, m,
	b, f, j, n,
	c, g, k, o,
	d, h, l, p float32,
) {
	matrix[0][0] = a
	matrix[0][1] = b
	matrix[0][2] = c
	matrix[0][3] = d

	matrix[1][0] = e
	matrix[1][1] = f
	matrix[1][2] = g
	matrix[1][3] = h

	matrix[2][0] = i
	matrix[2][1] = j
	matrix[2][2] = k
	matrix[2][3] = l

	matrix[3][0] = m
	matrix[3][1] = n
	matrix[3][2] = o
	matrix[3][3] = p
}

//------------------------------------------------------------------------------

// `At` returns the element at '(column, row)`.
func (m Mat4) At(column, row int) float32 {
	return m[column][row]
}

// `Set` sets the element at `(column, row)` to `value`.
func (m *Mat4) Set(column, row int, value float32) {
	m[column][row] = value
}

//------------------------------------------------------------------------------

// `Perspective` returns (by value) a perspective projection matrix.
func Perspective(fieldOfView float32, aspectRatio float32, near float32, far float32) Mat4 {
	f := float32(1.0) / math.Tan(fieldOfView/float32(2.0))

	return Mat4{
		{f / aspectRatio, 0, 0, 0},
		{0, f, 0, 0},
		{0, 0, (far + near) / (near - far), -1},
		{0, 0, (2 * far * near) / (near - far), 0},
	}
}

// `SetToPerspective` sets `m` to a perspective projection matrix.
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
