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
//
// See also `MakeMat4` and `SetTo`.
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

// `MakeMat4` returns a matrix. The elements are stored in
// alphabetical order (column-major order).
//
// See also `NewMat4` and `SetTo`.
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
//
// See also `NewMat4` and `SetTo`.
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

// `At` returns the element at '(row, column)`.
func (m Mat4) At(row, column int) float32 {
	return m[column][row]
}

// `Set` sets the element at `(row, column)` to `value`.
func (m *Mat4) Set(row, column int, value float32) {
	m[column][row] = value
}

//------------------------------------------------------------------------------

// `Perspective` returns a perspective projection matrix.
//
// See also `SetToPerspective`, `PerspectiveFrustum` and `SetToPerspectiveFrustum`.
func Perspective(fieldOfView, aspectRatio, near, far float32) Mat4 {
	f := float32(1.0) / math.Tan(fieldOfView/float32(2.0))

	return Mat4{
		{f / aspectRatio, 0, 0, 0},
		{0, f, 0, 0},
		{0, 0, (far + near) / (near - far), -1},
		{0, 0, (2 * far * near) / (near - far), 0},
	}
}

// `SetToPerspective` sets `m` to a perspective projection matrix.
//
// See also `Perspective`, `PerspectiveFrustum` and `SetToPerspectiveFrustum`.
func (m *Mat4) SetToPerspective(fieldOfView, aspectRatio, near, far float32) {
	f := float32(1.0) / math.Tan(fieldOfView/float32(2.0))

	m[0][0] = f / aspectRatio
	m[0][1] = 0
	m[0][2] = 0
	m[0][3] = 0

	m[1][0] = 0
	m[1][1] = f
	m[1][2] = 0
	m[1][3] = 0

	m[2][0] = 0
	m[2][1] = 0
	m[2][2] = (far + near) / (near - far)
	m[2][3] = -1

	m[3][0] = 0
	m[3][1] = 0
	m[3][2] = (2 * far * near) / (near - far)
	m[3][3] = 0
}

//------------------------------------------------------------------------------

// `PerspectiveFrustum` returns a perspective projection matrix.
//
// See also `SetToPerspectiveFrustum`, `Perspective` and `SetToPerspective`.
func PerspectiveFrustum(left, right, bottom, top, near, far float32) Mat4 {
	return Mat4{
		{(2 * near) / (right - left), 0, 0, 0},
		{0, (2 * near) / (top - bottom), 0, 0},
		{(right + left) / (right - left), (top + bottom) / (top - bottom), -(far + near) / (far - near), -1},
		{0, 0, -(2 * far * near) / (far - near), 0},
	}
}

// `SetToPerspectiveFrustum` sets `m` to a perspective projection matrix.
//
// See also `PerspectiveFrustum`, `Perspective` and `SetToPerspective`.
func (m *Mat4) SetToPerspectiveFrustum(left, right, bottom, top, near, far float32) {
	m[0][0] = (2 * near) / (right - left)
	m[0][1] = 0
	m[0][2] = 0
	m[0][3] = 0

	m[1][0] = 0
	m[1][1] = (2 * near) / (top - bottom)
	m[1][2] = 0
	m[1][3] = 0

	m[2][0] = (right + left) / (right - left)
	m[2][1] = (top + bottom) / (top - bottom)
	m[2][2] = -(far + near) / (far - near)
	m[2][3] = -1

	m[3][0] = 0
	m[3][1] = 0
	m[3][2] = -(2 * far * near) / (far - near)
	m[3][3] = 0
}

//------------------------------------------------------------------------------

// `Orthographic` returns an orthographic (parallel) projection matrix.
// `zoom` is the height of the projection plane.
//
// See also `SetToOrthographic`, `OrthographicFrustum` and `SetToOrthographicFrustum`.
func Orthographic(zoom, aspectRatio, near, far float32) Mat4 {
	top := zoom / 2
	right := top * aspectRatio
	return Mat4{
		{1 / right, 0, 0, 0},
		{0, 1 / top, 0, 0},
		{0, 0, -2 / (far - near), 0},
		{0, 0, -(far + near) / (far - near), 1},
	}
}

// `SetToOrthographic` sets `m` to an orthographic (parallel) projection matrix.
// `zoom` is the height of the projection plane.
//
// See also `Orthographic`, `OrthographicFrustum` and `SetToOrthographicFrustum`.
func (m *Mat4) SetToOrthographic(zoom, aspectRatio, near, far float32) {
	top := zoom / 2
	right := top * aspectRatio

	m[0][0] = 1 / right
	m[0][1] = 0
	m[0][2] = 0
	m[0][3] = 0

	m[1][0] = 0
	m[1][1] = 1 / top
	m[1][2] = 0
	m[1][3] = 0

	m[2][0] = 0
	m[2][1] = 0
	m[2][2] = -2 / (far - near)
	m[2][3] = 0

	m[3][0] = 0
	m[3][1] = 0
	m[3][2] = -(far + near) / (far - near)
	m[3][3] = 1
}

//------------------------------------------------------------------------------

// `OrthographicFrustum` returns an orthographic (parallel) projection matrix.
//
// See also `SetToOrthographicFrustum`, `Orthographic` and `SetToOrthographic`.
func OrthographicFrustum(left, right, bottom, top, near, far float32) Mat4 {
	return Mat4{
		{2 / (right - left), 0, 0, 0},
		{0, 2 / (top - bottom), 0, 0},
		{0, 0, -2 / (far - near), 0},
		{-(right + left) / (right - left), -(top + bottom) / (top - bottom), -(far + near) / (far - near), 1},
	}
}

// `SetToOrthographicFrustum` returns an orthographic (parallel) projection matrix.
//
// See also `OrthographicFrustum`, `Orthographic` and `SetToOrthographic`.
func (m *Mat4) SetToOrthographicFrustum(left, right, bottom, top, near, far float32) {
	m[0][0] = 2 / (right - left)
	m[0][1] = 0
	m[0][2] = 0
	m[0][3] = 0

	m[1][0] = 0
	m[1][1] = 2 / (top - bottom)
	m[1][2] = 0
	m[1][3] = 0

	m[2][0] = 0
	m[2][1] = 0
	m[2][2] = -2 / (far - near)
	m[2][3] = 0

	m[3][0] = -(right + left) / (right - left)
	m[3][1] = -(top + bottom) / (top - bottom)
	m[3][2] = -(far + near) / (far - near)
	m[3][3] = 1
}

//------------------------------------------------------------------------------

// `Translation` returns a translation matrix.
//
// See also `SetToTranslation`.
func Translation(t Vec3) Mat4 {
	return Mat4{
		{1, 0, 0, 0},
		{0, 1, 0, 0},
		{0, 0, 1, 0},
		{t.X, t.Y, t.Z, 1},
	}
}

// `SetToTranslation` sets `m` to a translation matrix.
//
// See also `Translation`.
func (m *Mat4) SetToTranslation(t Vec3) {
	m[0][0] = 1
	m[0][1] = 0
	m[0][2] = 0
	m[0][3] = 0

	m[1][0] = 0
	m[1][1] = 1
	m[1][2] = 0
	m[1][3] = 0

	m[2][0] = 0
	m[2][1] = 0
	m[2][2] = 1
	m[2][3] = 0

	m[3][0] = t.X
	m[3][1] = t.Y
	m[3][2] = t.Z
	m[3][3] = 1
}

//------------------------------------------------------------------------------

// `Rotation` returns a rotation matrix.
//
// See also `SetToRotation`.
func Rotation(angle float32, axis Vec3) Mat4 {
	c := math.Cos(angle)
	s := math.Sin(angle)

	return Mat4{
		{c + axis.X*axis.X*(1-c), -axis.Z*s + axis.X*axis.Y*(1-c), axis.Y*s + axis.X*axis.Z*(1-c), 0},
		{axis.Z*s + axis.Y*axis.X*(1-c), c + axis.Y*axis.Y*(1-c), -axis.X*s + axis.Y*axis.Z*(1-c), 0},
		{-axis.Y*s + axis.Z*axis.X*(1-c), axis.X*s + axis.Z*axis.Y*(1-c), c + axis.Z*axis.Z*(1-c), 0},
		{0, 0, 0, 1},
	}
}

// `SetToRotation` sets `m` to a rotation matrix.
//
// See also `Rotation`.
func (m *Mat4) Rotation(angle float32, axis Vec3) {
	c := math.Cos(angle)
	s := math.Sin(angle)

	m[0][0] = c + axis.X*axis.X*(1-c)
	m[0][1] = -axis.Z*s + axis.X*axis.Y*(1-c)
	m[0][2] = axis.Y*s + axis.X*axis.Z*(1-c)
	m[0][3] = 0

	m[1][0] = axis.Z*s + axis.Y*axis.X*(1-c)
	m[1][1] = c + axis.Y*axis.Y*(1-c)
	m[1][2] = -axis.X*s + axis.Y*axis.Z*(1-c)
	m[1][3] = 0

	m[2][0] = -axis.Y*s + axis.Z*axis.X*(1-c)
	m[2][1] = axis.X*s + axis.Z*axis.Y*(1-c)
	m[2][2] = c + axis.Z*axis.Z*(1-c)
	m[2][3] = 0

	m[3][0] = 0
	m[3][1] = 0
	m[3][2] = 0
	m[3][3] = 1
}

//------------------------------------------------------------------------------

// `Times` returns the matrix product of `m` and `o`.
//
// See also `Multiply`, `TimesVec` and `MultiplyVec`.
func (m *Mat4) Times(o *Mat4) Mat4 {
	return Mat4{
		{
			m[0][0]*o[0][0] + m[0][1]*o[1][0] + m[0][2]*o[2][0] + m[0][3]*o[3][0],
			m[0][0]*o[0][1] + m[0][1]*o[1][1] + m[0][2]*o[2][1] + m[0][3]*o[3][1],
			m[0][0]*o[0][2] + m[0][1]*o[1][2] + m[0][2]*o[2][2] + m[0][3]*o[3][2],
			m[0][0]*o[0][3] + m[0][1]*o[1][3] + m[0][2]*o[2][3] + m[0][3]*o[3][3],
		},
		{
			m[1][0]*o[0][0] + m[1][1]*o[1][0] + m[1][2]*o[2][0] + m[1][3]*o[3][0],
			m[1][0]*o[0][1] + m[1][1]*o[1][1] + m[1][2]*o[2][1] + m[1][3]*o[3][1],
			m[1][0]*o[0][2] + m[1][1]*o[1][2] + m[1][2]*o[2][2] + m[1][3]*o[3][2],
			m[1][0]*o[0][3] + m[1][1]*o[1][3] + m[1][2]*o[2][3] + m[1][3]*o[3][3],
		},
		{
			m[2][0]*o[0][0] + m[2][1]*o[1][0] + m[2][2]*o[2][0] + m[2][3]*o[3][0],
			m[2][0]*o[0][1] + m[2][1]*o[1][1] + m[2][2]*o[2][1] + m[2][3]*o[3][1],
			m[2][0]*o[0][2] + m[2][1]*o[1][2] + m[2][2]*o[2][2] + m[2][3]*o[3][2],
			m[2][0]*o[0][3] + m[2][1]*o[1][3] + m[2][2]*o[2][3] + m[2][3]*o[3][3],
		},
		{
			m[3][0]*o[0][0] + m[3][1]*o[1][0] + m[3][2]*o[2][0] + m[3][3]*o[3][0],
			m[3][0]*o[0][1] + m[3][1]*o[1][1] + m[3][2]*o[2][1] + m[3][3]*o[3][1],
			m[3][0]*o[0][2] + m[3][1]*o[1][2] + m[3][2]*o[2][2] + m[3][3]*o[3][2],
			m[3][0]*o[0][3] + m[3][1]*o[1][3] + m[3][2]*o[2][3] + m[3][3]*o[3][3],
		},
	}
}

// `Multiply` sets `r` to the matrix product of `m` and `o`.
//
// `r` must not be `m` or `o`.
//
// See also `Multiply`, `TimesVec` and `MultiplyVec`.
func (r *Mat4) Multiply(m, o *Mat4) Mat4 {
	r[0][0] = m[0][0]*o[0][0] + m[0][1]*o[1][0] + m[0][2]*o[2][0] + m[0][3]*o[3][0]
	r[0][1] = m[0][0]*o[0][1] + m[0][1]*o[1][1] + m[0][2]*o[2][1] + m[0][3]*o[3][1]
	r[0][2] = m[0][0]*o[0][2] + m[0][1]*o[1][2] + m[0][2]*o[2][2] + m[0][3]*o[3][2]
	r[0][3] = m[0][0]*o[0][3] + m[0][1]*o[1][3] + m[0][2]*o[2][3] + m[0][3]*o[3][3]

	r[1][0] = m[1][0]*o[0][0] + m[1][1]*o[1][0] + m[1][2]*o[2][0] + m[1][3]*o[3][0]
	r[1][1] = m[1][0]*o[0][1] + m[1][1]*o[1][1] + m[1][2]*o[2][1] + m[1][3]*o[3][1]
	r[1][2] = m[1][0]*o[0][2] + m[1][1]*o[1][2] + m[1][2]*o[2][2] + m[1][3]*o[3][2]
	r[1][3] = m[1][0]*o[0][3] + m[1][1]*o[1][3] + m[1][2]*o[2][3] + m[1][3]*o[3][3]

	r[2][0] = m[2][0]*o[0][0] + m[2][1]*o[1][0] + m[2][2]*o[2][0] + m[2][3]*o[3][0]
	r[2][1] = m[2][0]*o[0][1] + m[2][1]*o[1][1] + m[2][2]*o[2][1] + m[2][3]*o[3][1]
	r[2][2] = m[2][0]*o[0][2] + m[2][1]*o[1][2] + m[2][2]*o[2][2] + m[2][3]*o[3][2]
	r[2][3] = m[2][0]*o[0][3] + m[2][1]*o[1][3] + m[2][2]*o[2][3] + m[2][3]*o[3][3]

	r[3][0] = m[3][0]*o[0][0] + m[3][1]*o[1][0] + m[3][2]*o[2][0] + m[3][3]*o[3][0]
	r[3][1] = m[3][0]*o[0][1] + m[3][1]*o[1][1] + m[3][2]*o[2][1] + m[3][3]*o[3][1]
	r[3][2] = m[3][0]*o[0][2] + m[3][1]*o[1][2] + m[3][2]*o[2][2] + m[3][3]*o[3][2]
	r[3][3] = m[3][0]*o[0][3] + m[3][1]*o[1][3] + m[3][2]*o[2][3] + m[3][3]*o[3][3]
}

//------------------------------------------------------------------------------
