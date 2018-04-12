// Copyright (c) 2013-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package space

import "github.com/cozely/cozely/coord"

////////////////////////////////////////////////////////////////////////////////

// Matrix represents a transformation matrix.
type Matrix [4][4]float32

////////////////////////////////////////////////////////////////////////////////

// Transpose of the matrix.
func (m Matrix) Transpose() Matrix {
	return Matrix{
		{m[0][0], m[1][0], m[2][0], m[3][0]},
		{m[0][1], m[1][1], m[2][1], m[3][1]},
		{m[0][2], m[1][2], m[2][2], m[3][2]},
		{m[0][3], m[1][3], m[2][3], m[3][3]},
	}
}

// Times returns the matrix product with another matrix.
func (m Matrix) Times(o Matrix) Matrix {
	return Matrix{
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

////////////////////////////////////////////////////////////////////////////////

// Translation returns the translation vector of the transformation matrix.
func (m Matrix) Translation() coord.XYZ {
	return coord.XYZ{m[3][0], m[3][1], m[3][2]}
}

// WithoutTranslation returns the transformation matrix with the translation
// part removed.
func (m Matrix) WithoutTranslation() Matrix {
	m[3][0] = 0
	m[3][1] = 0
	m[3][2] = 0
	return m
}

//TODO: implement matrix decomposition (see "Graphics Gem 2", p. 320)

////////////////////////////////////////////////////////////////////////////////
