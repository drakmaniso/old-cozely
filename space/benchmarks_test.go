// Copyright (c) 2013-2016 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package space

//-----------------------------------------------------------------------------

import (
	"testing"

	"github.com/drakmaniso/glam/math"
)

//-----------------------------------------------------------------------------

func makeMatrix(
	a, e, i, m,
	b, f, j, n,
	c, g, k, o,
	d, h, l, p float32,
) Matrix {
	return Matrix{
		{a, b, c, d},
		{e, f, g, h},
		{i, j, k, l},
		{m, n, o, p},
	}
}

func (m *Matrix) setTo(
	a, e, i, mm,
	b, f, j, n,
	c, g, k, o,
	d, h, l, p float32,
) {
	m[0][0] = a
	m[0][1] = b
	m[0][2] = c
	m[0][3] = d

	m[1][0] = e
	m[1][1] = f
	m[1][2] = g
	m[1][3] = h

	m[2][0] = i
	m[2][1] = j
	m[2][2] = k
	m[2][3] = l

	m[3][0] = mm
	m[3][1] = n
	m[3][2] = o
	m[3][3] = p
}

func add(a, b Homogen) Homogen {
	return Homogen{
		a.X + b.X,
		a.Y + b.Y,
		a.Z + b.Z,
		a.W + b.W,
	}
}

func (v *Homogen) add(a, b Homogen) {
	v.X = a.X + b.X
	v.Y = a.Y + b.Y
	v.Z = a.Z + b.Z
	v.W = a.W + b.W
}

func (v *Homogen) subtract(a, b Homogen) {
	v.X = a.X - b.X
	v.Y = a.Y - b.Y
	v.Z = a.Z - b.Z
	v.W = a.W - b.W
}

func (v *Homogen) invert() {
	v.X = -v.X
	v.Y = -v.Y
	v.Z = -v.Z
	v.W = -v.W
}

func (v *Homogen) multiply(o Homogen, s Homogen) {
	v.X = o.X * s.X
	v.Y = o.Y * s.Y
	v.Z = o.Z * s.Z
	v.W = o.W * s.W
}

func (v *Homogen) divide(o Homogen, s Homogen) {
	v.X = o.X / s.X
	v.Y = o.Y / s.Y
	v.Z = o.Z / s.Z
	v.W = o.W / s.W
}

func (v *Homogen) normalize() {
	length := math.Sqrt(v.X*v.X + v.Y*v.Y + v.Z*v.Z + v.W*v.W)
	v.X /= length
	v.Y /= length
	v.Z /= length
	v.W /= length
}

type arrayHomogen [4]float32

func (v arrayHomogen) plus(o arrayHomogen) arrayHomogen {
	return arrayHomogen{v[0] + o[0], v[1] + o[1], v[2] + o[2], v[3] + o[3]}
}

func (v *arrayHomogen) add(a, b arrayHomogen) {
	v[0] = a[0] + b[0]
	v[1] = a[1] + b[1]
	v[2] = a[2] + b[2]
	v[3] = a[3] + b[3]
}

//-----------------------------------------------------------------------------

func BenchmarkHomogen_Plus(b *testing.B) {
	m := Homogen{1.1, 2.2, 3.3, 4.4}
	n := Homogen{5.5, 6.6, 7.7, 8.8}
	var o Homogen
	for i := 0; i < b.N; i++ {
		o = m.Plus(n)
	}
	_ = o
}

func BenchmarkHomogen_Plus_Add(b *testing.B) {
	m := Homogen{1.1, 2.2, 3.3, 4.4}
	n := Homogen{5.5, 6.6, 7.7, 8.8}
	var o Homogen
	for i := 0; i < b.N; i++ {
		o = add(m, n)
	}
	_ = o
}

func BenchmarkHomogen_Plus_Array(b *testing.B) {
	m := arrayHomogen{1.1, 2.2, 3.3, 4.4}
	n := arrayHomogen{5.5, 6.6, 7.7, 8.8}
	var o arrayHomogen
	for i := 0; i < b.N; i++ {
		o = m.plus(n)
	}
	_ = o
}

func BenchmarkHomogen_Plus_Self(b *testing.B) {
	m := Homogen{1.1, 2.2, 3.3, 4.4}
	n := Homogen{5.5, 6.6, 7.7, 8.8}
	for i := 0; i < b.N; i++ {
		m = m.Plus(n)
	}
	_ = m
}

func BenchmarkHomogen_Plus_ArraySelf(b *testing.B) {
	m := arrayHomogen{1.1, 2.2, 3.3, 4.4}
	n := arrayHomogen{5.5, 6.6, 7.7, 8.8}
	for i := 0; i < b.N; i++ {
		m = m.plus(n)
	}
	_ = m
}

func BenchmarkHomogen_Plus_ByRef(b *testing.B) {
	m := Homogen{1.1, 2.2, 3.3, 4.4}
	n := Homogen{5.5, 6.6, 7.7, 8.8}
	var o Homogen
	for i := 0; i < b.N; i++ {
		o.add(m, n)
	}
	_ = o
}

func BenchmarkHomogen_Plus_ArrayByRef(b *testing.B) {
	m := arrayHomogen{1.1, 2.2, 3.3, 4.4}
	n := arrayHomogen{5.5, 6.6, 7.7, 8.8}
	var o arrayHomogen
	for i := 0; i < b.N; i++ {
		o.add(m, n)
	}
	_ = o
}

func BenchmarkHomogen_Plus_ByRefSelf(b *testing.B) {
	m := Homogen{1.1, 2.2, 3.3, 4.4}
	n := Homogen{5.5, 6.6, 7.7, 8.8}
	for i := 0; i < b.N; i++ {
		m.add(m, n)
	}
	_ = m
}

func BenchmarkHomogen_Plus_ArrayByRefSelf(b *testing.B) {
	m := arrayHomogen{1.1, 2.2, 3.3, 4.4}
	n := arrayHomogen{5.5, 6.6, 7.7, 8.8}
	for i := 0; i < b.N; i++ {
		m.add(m, n)
	}
	_ = m
}

//-----------------------------------------------------------------------------

func BenchmarkHomogen_Times(b *testing.B) {
	m := Homogen{1.1, 2.2, 3.3, 4.4}
	n := float32(5.5)
	var o Homogen
	for i := 0; i < b.N; i++ {
		o = m.Times(Homogen{n, n, n, n})
	}
	_ = o
}

func BenchmarkHomogen_Times_Self(b *testing.B) {
	m := Homogen{1.1, 2.2, 3.3, 4.4}
	n := float32(5.5)
	for i := 0; i < b.N; i++ {
		m = m.Times(Homogen{n, n, n, n})
	}
	_ = m
}

func BenchmarkHomogen_Times_ByRef(b *testing.B) {
	m := Homogen{1.1, 2.2, 3.3, 4.4}
	n := float32(5.5)
	var o Homogen
	for i := 0; i < b.N; i++ {
		o.multiply(m, Homogen{n, n, n, n})
	}
	_ = o
}

func BenchmarkHomogen_Times_ByRefSelf(b *testing.B) {
	m := Homogen{1.1, 2.2, 3.3, 4.4}
	n := float32(5.5)
	for i := 0; i < b.N; i++ {
		m.multiply(m, Homogen{n, n, n, n})
	}
	_ = m
}

//-----------------------------------------------------------------------------

func BenchmarkHomogen_Slash(b *testing.B) {
	m := Homogen{1.1, 2.2, 3.3, 4.4}
	n := float32(5.5)
	var o Homogen
	for i := 0; i < b.N; i++ {
		o = m.Slash(Homogen{n, n, n, n})
	}
	_ = o
}

func BenchmarkHomogen_Slash_Self(b *testing.B) {
	m := Homogen{1.1, 2.2, 3.3, 4.4}
	n := float32(5.5)
	for i := 0; i < b.N; i++ {
		m = m.Slash(Homogen{n, n, n, n})
	}
	_ = m
}

func BenchmarkHomogen_Slash_ByRef(b *testing.B) {
	m := Homogen{1.1, 2.2, 3.3, 4.4}
	n := float32(5.5)
	var o Homogen
	for i := 0; i < b.N; i++ {
		o.divide(m, Homogen{n, n, n, n})
	}
	_ = o
}

func BenchmarkHomogen_Slash_ByRefSelf(b *testing.B) {
	m := Homogen{1.1, 2.2, 3.3, 4.4}
	n := float32(5.5)
	for i := 0; i < b.N; i++ {
		m.divide(m, Homogen{n, n, n, n})
	}
	_ = m
}

//-----------------------------------------------------------------------------

func BenchmarkHomogen_Normalized(b *testing.B) {
	m := Homogen{1.1, 2.2, 3.3, 4.4}
	var o Homogen
	for i := 0; i < b.N; i++ {
		o = m.Normalized()
	}
	_ = o
}

func BenchmarkHomogen_Normalized_Self(b *testing.B) {
	m := Homogen{1.1, 2.2, 3.3, 4.4}
	for i := 0; i < b.N; i++ {
		m = m.Normalized()
	}
	_ = m
}

func BenchmarkHomogen_Normalize_ByRef(b *testing.B) {
	m := Homogen{1.1, 2.2, 3.3, 4.4}
	for i := 0; i < b.N; i++ {
		m.normalize()
	}
	_ = m
}

//-----------------------------------------------------------------------------

func (m *Matrix) timesThreeRefs(a, b *Matrix) {
	m[0][0] = a[0][0]*b[0][0] + a[0][1]*b[1][0] + a[0][2]*b[2][0] + a[0][3]*b[3][0]
	m[0][1] = a[0][0]*b[0][1] + a[0][1]*b[1][1] + a[0][2]*b[2][1] + a[0][3]*b[3][1]
	m[0][2] = a[0][0]*b[0][2] + a[0][1]*b[1][2] + a[0][2]*b[2][2] + a[0][3]*b[3][2]
	m[0][3] = a[0][0]*b[0][3] + a[0][1]*b[1][3] + a[0][2]*b[2][3] + a[0][3]*b[3][3]

	m[1][0] = a[1][0]*b[0][0] + a[1][1]*b[1][0] + a[1][2]*b[2][0] + a[1][3]*b[3][0]
	m[1][1] = a[1][0]*b[0][1] + a[1][1]*b[1][1] + a[1][2]*b[2][1] + a[1][3]*b[3][1]
	m[1][2] = a[1][0]*b[0][2] + a[1][1]*b[1][2] + a[1][2]*b[2][2] + a[1][3]*b[3][2]
	m[1][3] = a[1][0]*b[0][3] + a[1][1]*b[1][3] + a[1][2]*b[2][3] + a[1][3]*b[3][3]

	m[2][0] = a[2][0]*b[0][0] + a[2][1]*b[1][0] + a[2][2]*b[2][0] + a[2][3]*b[3][0]
	m[2][1] = a[2][0]*b[0][1] + a[2][1]*b[1][1] + a[2][2]*b[2][1] + a[2][3]*b[3][1]
	m[2][2] = a[2][0]*b[0][2] + a[2][1]*b[1][2] + a[2][2]*b[2][2] + a[2][3]*b[3][2]
	m[2][3] = a[2][0]*b[0][3] + a[2][1]*b[1][3] + a[2][2]*b[2][3] + a[2][3]*b[3][3]

	m[3][0] = a[3][0]*b[0][0] + a[3][1]*b[1][0] + a[3][2]*b[2][0] + a[3][3]*b[3][0]
	m[3][1] = a[3][0]*b[0][1] + a[3][1]*b[1][1] + a[3][2]*b[2][1] + a[3][3]*b[3][1]
	m[3][2] = a[3][0]*b[0][2] + a[3][1]*b[1][2] + a[3][2]*b[2][2] + a[3][3]*b[3][2]
	m[3][3] = a[3][0]*b[0][3] + a[3][1]*b[1][3] + a[3][2]*b[2][3] + a[3][3]*b[3][3]
}

func (m Matrix) timesRecvValArgRef(o *Matrix) Matrix {
	return Matrix{
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

func (m *Matrix) timesRecvRefArgVal(o Matrix) Matrix {
	return Matrix{
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

func (m Matrix) timesRecvValArgVal(o Matrix) Matrix {
	return Matrix{
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

//------------------------------------------------------------------------------

func BenchmarkMatrix_literal(b *testing.B) {
	var m Matrix
	for i := 0; i < b.N; i++ {
		m = Matrix{
			{1.1, 2.1, 3.1, 4.1},
			{1.2, 2.2, 3.2, 4.2},
			{1.3, 2.3, 3.3, 4.3},
			{1.4, 2.4, 3.4, 4.4},
		}
	}
	_ = m
}

func BenchmarkMatrix_MakeMatrix(b *testing.B) {
	var a Matrix
	for i := 0; i < b.N; i++ {
		a = makeMatrix(
			1.1, 11.1, 111.1, 1111.1,
			2.2, 22.2, 222.2, 2222.2,
			3.3, 33.3, 333.3, 3333.3,
			4.4, 44.4, 444.4, 4444.4,
		)
	}
	_ = a
}

func BenchmarkMatrix_NewMatrix(b *testing.B) {
	var a *Matrix
	for i := 0; i < b.N; i++ {
		a = &Matrix{
			{1.1, 2.1, 3.1, 4.1},
			{1.2, 2.2, 3.2, 4.2},
			{1.3, 2.3, 3.3, 4.3},
			{1.4, 2.4, 3.4, 4.4},
		}
	}
	_ = a
}

func BenchmarkMatrix_SetTo(b *testing.B) {
	var a Matrix
	for i := 0; i < b.N; i++ {
		a.setTo(
			1.1, 2.1, 3.1, 4.1,
			1.2, 2.2, 3.2, 4.2,
			1.3, 2.3, 3.3, 4.3,
			1.4, 2.4, 3.4, 4.4,
		)
	}
	_ = a
}

//------------------------------------------------------------------------------

func BenchmarkMatrix_Times(b *testing.B) {
	m := &Matrix{
		{1.1, 2.1, 3.1, 4.1},
		{1.2, 2.2, 3.2, 4.2},
		{1.3, 2.3, 3.3, 4.3},
		{1.4, 2.4, 3.4, 4.4},
	}
	n := Matrix{
		{10.1, 20.1, 30.1, 40.1},
		{10.2, 20.2, 30.2, 40.2},
		{10.3, 20.3, 30.3, 40.3},
		{10.4, 20.4, 30.4, 40.4},
	}
	var o Matrix
	for i := 0; i < b.N; i++ {
		o = m.Times(n)
	}
	_ = o
}

func BenchmarkMatrix_Times_RecvMutated_TwoArgsByRef(b *testing.B) {
	m := &Matrix{
		{1.1, 2.1, 3.1, 4.1},
		{1.2, 2.2, 3.2, 4.2},
		{1.3, 2.3, 3.3, 4.3},
		{1.4, 2.4, 3.4, 4.4},
	}
	n := &Matrix{
		{10.1, 20.1, 30.1, 40.1},
		{10.2, 20.2, 30.2, 40.2},
		{10.3, 20.3, 30.3, 40.3},
		{10.4, 20.4, 30.4, 40.4},
	}
	o := &Matrix{}
	for i := 0; i < b.N; i++ {
		o.timesThreeRefs(m, n)
	}
	_ = o
}

func BenchmarkMatrix_Times_RecvByRef_ArgByVal(b *testing.B) {
	m := Matrix{
		{1.1, 2.1, 3.1, 4.1},
		{1.2, 2.2, 3.2, 4.2},
		{1.3, 2.3, 3.3, 4.3},
		{1.4, 2.4, 3.4, 4.4},
	}
	n := Matrix{
		{10.1, 20.1, 30.1, 40.1},
		{10.2, 20.2, 30.2, 40.2},
		{10.3, 20.3, 30.3, 40.3},
		{10.4, 20.4, 30.4, 40.4},
	}
	var o Matrix
	for i := 0; i < b.N; i++ {
		o = m.timesRecvRefArgVal(n)
	}
	_ = o
}

func BenchmarkMatrix_Times_RecvByVal_ArgByRef(b *testing.B) {
	m := Matrix{
		{1.1, 2.1, 3.1, 4.1},
		{1.2, 2.2, 3.2, 4.2},
		{1.3, 2.3, 3.3, 4.3},
		{1.4, 2.4, 3.4, 4.4},
	}
	n := &Matrix{
		{10.1, 20.1, 30.1, 40.1},
		{10.2, 20.2, 30.2, 40.2},
		{10.3, 20.3, 30.3, 40.3},
		{10.4, 20.4, 30.4, 40.4},
	}
	var o Matrix
	for i := 0; i < b.N; i++ {
		o = m.timesRecvValArgRef(n)
	}
	_ = o
}

func BenchmarkMatrix_Times_RecvByVal_ArgByVal(b *testing.B) {
	m := Matrix{
		{1.1, 2.1, 3.1, 4.1},
		{1.2, 2.2, 3.2, 4.2},
		{1.3, 2.3, 3.3, 4.3},
		{1.4, 2.4, 3.4, 4.4},
	}
	n := Matrix{
		{10.1, 20.1, 30.1, 40.1},
		{10.2, 20.2, 30.2, 40.2},
		{10.3, 20.3, 30.3, 40.3},
		{10.4, 20.4, 30.4, 40.4},
	}
	var o Matrix
	for i := 0; i < b.N; i++ {
		o = m.timesRecvValArgVal(n)
	}
	_ = o
}

func BenchmarkMatrix_Times_RecvByValAddr_ArgByVal(b *testing.B) {
	m := &Matrix{
		{1.1, 2.1, 3.1, 4.1},
		{1.2, 2.2, 3.2, 4.2},
		{1.3, 2.3, 3.3, 4.3},
		{1.4, 2.4, 3.4, 4.4},
	}
	n := Matrix{
		{10.1, 20.1, 30.1, 40.1},
		{10.2, 20.2, 30.2, 40.2},
		{10.3, 20.3, 30.3, 40.3},
		{10.4, 20.4, 30.4, 40.4},
	}
	var o Matrix
	for i := 0; i < b.N; i++ {
		o = m.timesRecvValArgVal(n)
	}
	_ = o
}

//------------------------------------------------------------------------------

func (m *Matrix) rotationSetAndReturn(angle float32, axis Coord) Matrix {
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

	return *m
}

func (m *Matrix) rotationSetOnly(angle float32, axis Coord) {
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

// func BenchmarkMatrix_Rotation(b *testing.B) {
// 	axis := Coord{1, 2, 3}
// 	var o Matrix
// 	for i := 0; i < b.N; i++ {
// 		o = space.Rotation(3.14, axis)
// 	}
// 	_ = o
// }

func BenchmarkMatrix_Rotation_SetAndReturn(b *testing.B) {
	axis := Coord{1, 2, 3}
	var m Matrix
	var o Matrix
	for i := 0; i < b.N; i++ {
		o = m.rotationSetAndReturn(3.14, axis)
	}
	_ = o
}

func BenchmarkMatrix_Rotation_SetAndDiscardReturn(b *testing.B) {
	axis := Coord{1, 2, 3}
	var o Matrix
	for i := 0; i < b.N; i++ {
		o.rotationSetAndReturn(3.14, axis)
	}
	_ = o
}

func BenchmarkMatrix_Rotation_SetOnly(b *testing.B) {
	axis := Coord{1, 2, 3}
	var o Matrix
	for i := 0; i < b.N; i++ {
		o.rotationSetOnly(3.14, axis)
	}
	_ = o
}

//------------------------------------------------------------------------------
