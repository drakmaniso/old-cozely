// Copyright (c) 2013-2016 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package geom

//------------------------------------------------------------------------------

// Mat4 is a 32-bit float matrix with 4 columns and 4 rows.
type Mat4 [4][4]float32

// Mat3 is a 32-bit float matrix with 3 columns and 3 rows.
//
// Note: due to alignment rules, it's not possible to pass a Mat3 directly to
// the GPU. You should use Mat3x4 instead.
type Mat3 [3][3]float32

// Mat2 is a 32-bit float matrix with 2 columns and 2 rows.
type Mat2 [2][2]float32

// Mat3x4 is a 32-bit float matrix with 3 columns and 4 rows.
type Mat3x4 [3][4]float32

//------------------------------------------------------------------------------

// Mat3 returns the upper-left part of the matrix.
func (m *Mat4) Mat3() Mat3 {
	return Mat3{
		{m[0][0], m[1][0], m[2][0]},
		{m[0][1], m[1][1], m[2][1]},
		{m[0][2], m[1][2], m[2][2]},
	}
}

// Mat3x4 returns the left part of the matrix.
func (m *Mat4) Mat3x4() Mat3x4 {
	return Mat3x4{
		{m[0][0], m[1][0], m[2][0], m[3][0]},
		{m[0][1], m[1][1], m[2][1], m[3][1]},
		{m[0][2], m[1][2], m[2][2], m[3][2]},
	}
}

// Mat2 returns the upper-left part of the matrix.
func (m *Mat4) Mat2() Mat2 {
	return Mat2{
		{m[0][0], m[1][0]},
		{m[0][1], m[1][1]},
	}
}

// Mat2 returns the upper-left part of the matrix.
func (m *Mat3) Mat2() Mat2 {
	return Mat2{
		{m[0][0], m[1][0]},
		{m[0][1], m[1][1]},
	}
}

// Mat3x4 returns the left part of the matrix.
func (m *Mat3) Mat3x4() Mat3x4 {
	return Mat3x4{
		{m[0][0], m[1][0], m[2][0], 0},
		{m[0][1], m[1][1], m[2][1], 0},
		{m[0][2], m[1][2], m[2][2], 0},
	}
}

//------------------------------------------------------------------------------

// Transpose of the matrix.
func (m *Mat4) Transpose() Mat4 {
	return Mat4{
		{m[0][0], m[1][0], m[2][0], m[3][0]},
		{m[0][1], m[1][1], m[2][1], m[3][1]},
		{m[0][2], m[1][2], m[2][2], m[3][2]},
		{m[0][3], m[1][3], m[2][3], m[3][3]},
	}
}

// Transpose of the matrix.
func (m *Mat3) Transpose() Mat3 {
	return Mat3{
		{m[0][0], m[1][0], m[2][0]},
		{m[0][1], m[1][1], m[2][1]},
		{m[0][2], m[1][2], m[2][2]},
	}
}

// Transpose of the matrix.
func (m *Mat2) Transpose() Mat2 {
	return Mat2{
		{m[0][0], m[1][0]},
		{m[0][1], m[1][1]},
	}
}

//------------------------------------------------------------------------------

// Times returns the matrix product with another matrix.
func (m *Mat4) Times(o Mat4) Mat4 {
	return Mat4{
		{
			m[0][0]*o[0][0] + m[1][0]*o[0][1] + m[2][0]*o[0][2] + m[3][0]*o[0][3],
			m[0][1]*o[0][0] + m[1][1]*o[0][1] + m[2][1]*o[0][2] + m[3][1]*o[0][3],
			m[0][2]*o[0][0] + m[1][2]*o[0][1] + m[2][2]*o[0][2] + m[3][2]*o[0][3],
			m[0][3]*o[0][0] + m[1][3]*o[0][1] + m[2][3]*o[0][2] + m[3][3]*o[0][3],
		},
		{
			m[0][0]*o[1][0] + m[1][0]*o[1][1] + m[2][0]*o[1][2] + m[3][0]*o[1][3],
			m[0][1]*o[1][0] + m[1][1]*o[1][1] + m[2][1]*o[1][2] + m[3][1]*o[1][3],
			m[0][2]*o[1][0] + m[1][2]*o[1][1] + m[2][2]*o[1][2] + m[3][2]*o[1][3],
			m[0][3]*o[1][0] + m[1][3]*o[1][1] + m[2][3]*o[1][2] + m[3][3]*o[1][3],
		},
		{
			m[0][0]*o[2][0] + m[1][0]*o[2][1] + m[2][0]*o[2][2] + m[3][0]*o[2][3],
			m[0][1]*o[2][0] + m[1][1]*o[2][1] + m[2][1]*o[2][2] + m[3][1]*o[2][3],
			m[0][2]*o[2][0] + m[1][2]*o[2][1] + m[2][2]*o[2][2] + m[3][2]*o[2][3],
			m[0][3]*o[2][0] + m[1][3]*o[2][1] + m[2][3]*o[2][2] + m[3][3]*o[2][3],
		},
		{
			m[0][0]*o[3][0] + m[1][0]*o[3][1] + m[2][0]*o[3][2] + m[3][0]*o[3][3],
			m[0][1]*o[3][0] + m[1][1]*o[3][1] + m[2][1]*o[3][2] + m[3][1]*o[3][3],
			m[0][2]*o[3][0] + m[1][2]*o[3][1] + m[2][2]*o[3][2] + m[3][2]*o[3][3],
			m[0][3]*o[3][0] + m[1][3]*o[3][1] + m[2][3]*o[3][2] + m[3][3]*o[3][3],
		},
	}
}

// Times returns the matrix product with another matrix.
func (m *Mat3) Times(o Mat3) Mat3 {
	return Mat3{
		{
			m[0][0]*o[0][0] + m[1][0]*o[0][1] + m[2][0]*o[0][2],
			m[0][1]*o[0][0] + m[1][1]*o[0][1] + m[2][1]*o[0][2],
			m[0][2]*o[0][0] + m[1][2]*o[0][1] + m[2][2]*o[0][2],
		},
		{
			m[0][0]*o[1][0] + m[1][0]*o[1][1] + m[2][0]*o[1][2],
			m[0][1]*o[1][0] + m[1][1]*o[1][1] + m[2][1]*o[1][2],
			m[0][2]*o[1][0] + m[1][2]*o[1][1] + m[2][2]*o[1][2],
		},
		{
			m[0][0]*o[2][0] + m[1][0]*o[2][1] + m[2][0]*o[2][2],
			m[0][1]*o[2][0] + m[1][1]*o[2][1] + m[2][1]*o[2][2],
			m[0][2]*o[2][0] + m[1][2]*o[2][1] + m[2][2]*o[2][2],
		},
	}
}

// Times returns the matrix product with another matrix.
func (m *Mat2) Times(o Mat2) Mat2 {
	return Mat2{
		{
			m[0][0]*o[0][0] + m[1][0]*o[0][1],
			m[0][1]*o[0][0] + m[1][1]*o[0][1],
		},
		{
			m[0][0]*o[1][0] + m[1][0]*o[1][1],
			m[0][1]*o[1][0] + m[1][1]*o[1][1],
		},
	}
}

//------------------------------------------------------------------------------
