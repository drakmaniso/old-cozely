// Copyright (c) 2013-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package plane

////////////////////////////////////////////////////////////////////////////////

// Matrix represents a transformation matrix.
//
// Note: due to alignment rules, it's not possible to pass a Matrix directly
// to the GPU. You should convert it to a GPUMatrix first.
type Matrix [3][3]float32

////////////////////////////////////////////////////////////////////////////////

// Transpose of the matrix.
func (m Matrix) Transpose() Matrix {
	return Matrix{
		{m[0][0], m[1][0], m[2][0]},
		{m[0][1], m[1][1], m[2][1]},
		{m[0][2], m[1][2], m[2][2]},
	}
}

// Times returns the matrix product with another transformation matrix.
func (m Matrix) Times(o Matrix) Matrix {
	return Matrix{
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

////////////////////////////////////////////////////////////////////////////////

// Translation returns the translation vector of the transformation matrix.
func (m Matrix) Translation() Coord {
	return Coord{m[2][0], m[2][1]}
}

// WithoutTranslation returns the transformation matrix with the translation
// part removed.
func (m Matrix) WithoutTranslation() Matrix {
	m[2][0] = 0
	m[2][1] = 0
	return m
}

//TODO: implement matrix decomposition (see "Graphics Gem 2", p. 320)

////////////////////////////////////////////////////////////////////////////////

// GPUMatrix represents a transformation matrix with a memory layout compatible
// with the GPU.
type GPUMatrix [3][4]float32

// GPU returns a GPUMatrix version of the matrix.
func (m Matrix) GPU() GPUMatrix {
	return GPUMatrix{
		{m[0][0], m[1][0], m[2][0], 0},
		{m[0][1], m[1][1], m[2][1], 0},
		{m[0][2], m[1][2], m[2][2], 0},
	}
}

////////////////////////////////////////////////////////////////////////////////
