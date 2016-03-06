// Copyright (c) 2013-2016 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package geom

import (
	"testing"

	"github.com/drakmaniso/glam/math"
)

//-----------------------------------------------------------------------------
func makeMat4(
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

func (m *Mat4) setTo(
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

func (v *Vec4) add(a, b Vec4) {
	v.X = a.X + b.X
	v.Y = a.Y + b.Y
	v.Z = a.Z + b.Z
	v.W = a.W + b.W
}

func (v *Vec4) subtract(a, b Vec4) {
	v.X = a.X - b.X
	v.Y = a.Y - b.Y
	v.Z = a.Z - b.Z
	v.W = a.W - b.W
}

func (v *Vec4) invert() {
	v.X = -v.X
	v.Y = -v.Y
	v.Z = -v.Z
	v.W = -v.W
}

func (v *Vec4) multiply(o Vec4, s float32) {
	v.X = o.X * s
	v.Y = o.Y * s
	v.Z = o.Z * s
	v.W = o.W * s
}

func (v *Vec4) divide(o Vec4, s float32) {
	v.X = o.X / s
	v.Y = o.Y / s
	v.Z = o.Z / s
	v.W = o.W / s
}

func (v *Vec4) normalize() {
	length := math.Sqrt(v.X*v.X + v.Y*v.Y + v.Z*v.Z + v.W*v.W)
	v.X /= length
	v.Y /= length
	v.Z /= length
	v.W /= length
}

type arrayVec4 [4]float32

func (v arrayVec4) plus(o arrayVec4) arrayVec4 {
	return arrayVec4{v[0] + o[0], v[1] + o[1], v[2] + o[2], v[3] + o[3]}
}

func (v *arrayVec4) add(a, b arrayVec4) {
	v[0] = a[0] + b[0]
	v[1] = a[1] + b[1]
	v[2] = a[2] + b[2]
	v[3] = a[3] + b[3]
}

//-----------------------------------------------------------------------------

func BenchmarkVec4_Plus(b *testing.B) {
	m := Vec4{1.1, 2.2, 3.3, 4.4}
	n := Vec4{5.5, 6.6, 7.7, 8.8}
	var o Vec4
	for i := 0; i < b.N; i++ {
		o = m.Plus(n)
	}
	_ = o
}

func BenchmarkVec4_Plus_Array(b *testing.B) {
	m := arrayVec4{1.1, 2.2, 3.3, 4.4}
	n := arrayVec4{5.5, 6.6, 7.7, 8.8}
	var o arrayVec4
	for i := 0; i < b.N; i++ {
		o = m.plus(n)
	}
	_ = o
}

func BenchmarkVec4_Plus_Self(b *testing.B) {
	m := Vec4{1.1, 2.2, 3.3, 4.4}
	n := Vec4{5.5, 6.6, 7.7, 8.8}
	for i := 0; i < b.N; i++ {
		m = m.Plus(n)
	}
	_ = m
}

func BenchmarkVec4_Plus_ArraySelf(b *testing.B) {
	m := arrayVec4{1.1, 2.2, 3.3, 4.4}
	n := arrayVec4{5.5, 6.6, 7.7, 8.8}
	for i := 0; i < b.N; i++ {
		m = m.plus(n)
	}
	_ = m
}

func BenchmarkVec4_Plus_ByRef(b *testing.B) {
	m := Vec4{1.1, 2.2, 3.3, 4.4}
	n := Vec4{5.5, 6.6, 7.7, 8.8}
	var o Vec4
	for i := 0; i < b.N; i++ {
		o.add(m, n)
	}
	_ = o
}

func BenchmarkVec4_Plus_ArrayByRef(b *testing.B) {
	m := arrayVec4{1.1, 2.2, 3.3, 4.4}
	n := arrayVec4{5.5, 6.6, 7.7, 8.8}
	var o arrayVec4
	for i := 0; i < b.N; i++ {
		o.add(m, n)
	}
	_ = o
}

func BenchmarkVec4_Plus_ByRefSelf(b *testing.B) {
	m := Vec4{1.1, 2.2, 3.3, 4.4}
	n := Vec4{5.5, 6.6, 7.7, 8.8}
	for i := 0; i < b.N; i++ {
		m.add(m, n)
	}
	_ = m
}

func BenchmarkVec4_Plus_ArrayByRefSelf(b *testing.B) {
	m := arrayVec4{1.1, 2.2, 3.3, 4.4}
	n := arrayVec4{5.5, 6.6, 7.7, 8.8}
	for i := 0; i < b.N; i++ {
		m.add(m, n)
	}
	_ = m
}

//-----------------------------------------------------------------------------

func BenchmarkVec4_Times(b *testing.B) {
	m := Vec4{1.1, 2.2, 3.3, 4.4}
	n := float32(5.5)
	var o Vec4
	for i := 0; i < b.N; i++ {
		o = m.Times(n)
	}
	_ = o
}

func BenchmarkVec4_Times_Self(b *testing.B) {
	m := Vec4{1.1, 2.2, 3.3, 4.4}
	n := float32(5.5)
	for i := 0; i < b.N; i++ {
		m = m.Times(n)
	}
	_ = m
}

func BenchmarkVec4_Times_ByRef(b *testing.B) {
	m := Vec4{1.1, 2.2, 3.3, 4.4}
	n := float32(5.5)
	var o Vec4
	for i := 0; i < b.N; i++ {
		o.multiply(m, n)
	}
	_ = o
}

func BenchmarkVec4_Times_ByRefSelf(b *testing.B) {
	m := Vec4{1.1, 2.2, 3.3, 4.4}
	n := float32(5.5)
	for i := 0; i < b.N; i++ {
		m.multiply(m, n)
	}
	_ = m
}

//-----------------------------------------------------------------------------

func BenchmarkVec4_Slash(b *testing.B) {
	m := Vec4{1.1, 2.2, 3.3, 4.4}
	n := float32(5.5)
	var o Vec4
	for i := 0; i < b.N; i++ {
		o = m.Slash(n)
	}
	_ = o
}

func BenchmarkVec4_Slash_Self(b *testing.B) {
	m := Vec4{1.1, 2.2, 3.3, 4.4}
	n := float32(5.5)
	for i := 0; i < b.N; i++ {
		m = m.Slash(n)
	}
	_ = m
}

func BenchmarkVec4_Slash_ByRef(b *testing.B) {
	m := Vec4{1.1, 2.2, 3.3, 4.4}
	n := float32(5.5)
	var o Vec4
	for i := 0; i < b.N; i++ {
		o.divide(m, n)
	}
	_ = o
}

func BenchmarkVec4_Slash_ByRefSelf(b *testing.B) {
	m := Vec4{1.1, 2.2, 3.3, 4.4}
	n := float32(5.5)
	for i := 0; i < b.N; i++ {
		m.divide(m, n)
	}
	_ = m
}

//-----------------------------------------------------------------------------

func BenchmarkVec4_Normalized(b *testing.B) {
	m := Vec4{1.1, 2.2, 3.3, 4.4}
	var o Vec4
	for i := 0; i < b.N; i++ {
		o = m.Normalized()
	}
	_ = o
}

func BenchmarkVec4_Normalized_Self(b *testing.B) {
	m := Vec4{1.1, 2.2, 3.3, 4.4}
	for i := 0; i < b.N; i++ {
		m = m.Normalized()
	}
	_ = m
}

func BenchmarkVec4_Normalize_ByRef(b *testing.B) {
	m := Vec4{1.1, 2.2, 3.3, 4.4}
	for i := 0; i < b.N; i++ {
		m.normalize()
	}
	_ = m
}

//-----------------------------------------------------------------------------

func (m *Mat4) multiply(a, b *Mat4) {
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

func (m Mat4) timesOneRef(o *Mat4) Mat4 {
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

func (m Mat4) timesTwoValues(o Mat4) Mat4 {
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

//------------------------------------------------------------------------------

func BenchmarkMat4_literal(b *testing.B) {
	var m Mat4
	for i := 0; i < b.N; i++ {
		m = Mat4{
			{1.1, 2.1, 3.1, 4.1},
			{1.2, 2.2, 3.2, 4.2},
			{1.3, 2.3, 3.3, 4.3},
			{1.4, 2.4, 3.4, 4.4},
		}
	}
	_ = m
}

func BenchmarkMat4_MakeMat4(b *testing.B) {
	var a Mat4
	for i := 0; i < b.N; i++ {
		a = makeMat4(
			1.1, 11.1, 111.1, 1111.1,
			2.2, 22.2, 222.2, 2222.2,
			3.3, 33.3, 333.3, 3333.3,
			4.4, 44.4, 444.4, 4444.4,
		)
	}
	_ = a
}

func BenchmarkMat4_NewMat4(b *testing.B) {
	var a *Mat4
	for i := 0; i < b.N; i++ {
		a = &Mat4{
			{1.1, 2.1, 3.1, 4.1},
			{1.2, 2.2, 3.2, 4.2},
			{1.3, 2.3, 3.3, 4.3},
			{1.4, 2.4, 3.4, 4.4},
		}
	}
	_ = a
}

func BenchmarkMat4_SetTo(b *testing.B) {
	var a Mat4
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

func BenchmarkMat4_Times(b *testing.B) {
	m := &Mat4{
		{1.1, 2.1, 3.1, 4.1},
		{1.2, 2.2, 3.2, 4.2},
		{1.3, 2.3, 3.3, 4.3},
		{1.4, 2.4, 3.4, 4.4},
	}
	n := &Mat4{
		{10.1, 20.1, 30.1, 40.1},
		{10.2, 20.2, 30.2, 40.2},
		{10.3, 20.3, 30.3, 40.3},
		{10.4, 20.4, 30.4, 40.4},
	}
	var o Mat4
	for i := 0; i < b.N; i++ {
		o = m.Times(n)
	}
	_ = o
}

func BenchmarkMat4_Times_ThreeRefs(b *testing.B) {
	m := &Mat4{
		{1.1, 2.1, 3.1, 4.1},
		{1.2, 2.2, 3.2, 4.2},
		{1.3, 2.3, 3.3, 4.3},
		{1.4, 2.4, 3.4, 4.4},
	}
	n := &Mat4{
		{10.1, 20.1, 30.1, 40.1},
		{10.2, 20.2, 30.2, 40.2},
		{10.3, 20.3, 30.3, 40.3},
		{10.4, 20.4, 30.4, 40.4},
	}
	o := &Mat4{}
	for i := 0; i < b.N; i++ {
		o.multiply(m, n)
	}
	_ = o
}

func BenchmarkMat4_Times_OneRef(b *testing.B) {
	m := Mat4{
		{1.1, 2.1, 3.1, 4.1},
		{1.2, 2.2, 3.2, 4.2},
		{1.3, 2.3, 3.3, 4.3},
		{1.4, 2.4, 3.4, 4.4},
	}
	n := &Mat4{
		{10.1, 20.1, 30.1, 40.1},
		{10.2, 20.2, 30.2, 40.2},
		{10.3, 20.3, 30.3, 40.3},
		{10.4, 20.4, 30.4, 40.4},
	}
	var o Mat4
	for i := 0; i < b.N; i++ {
		o = m.timesOneRef(n)
	}
	_ = o
}

func BenchmarkMat4_Times_TwoValues(b *testing.B) {
	m := Mat4{
		{1.1, 2.1, 3.1, 4.1},
		{1.2, 2.2, 3.2, 4.2},
		{1.3, 2.3, 3.3, 4.3},
		{1.4, 2.4, 3.4, 4.4},
	}
	n := Mat4{
		{10.1, 20.1, 30.1, 40.1},
		{10.2, 20.2, 30.2, 40.2},
		{10.3, 20.3, 30.3, 40.3},
		{10.4, 20.4, 30.4, 40.4},
	}
	var o Mat4
	for i := 0; i < b.N; i++ {
		o = m.timesTwoValues(n)
	}
	_ = o
}

//------------------------------------------------------------------------------
