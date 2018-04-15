// Copyright (c) 2013-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package coord

import (
	"testing"

	"github.com/cozely/cozely/x/math32"
)

//-----------------------------------------------------------------------------

func add(a, b XYZ) XYZ {
	return XYZ{
		a.X + b.X,
		a.Y + b.Y,
		a.Z + b.Z,
	}
}

func (v *XYZ) add(a, b XYZ) {
	v.X = a.X + b.X
	v.Y = a.Y + b.Y
	v.Z = a.Z + b.Z
}

func (v *XYZ) subtract(a, b XYZ) {
	v.X = a.X - b.X
	v.Y = a.Y - b.Y
	v.Z = a.Z - b.Z
}

func (v *XYZ) invert() {
	v.X = -v.X
	v.Y = -v.Y
	v.Z = -v.Z
}

func (v *XYZ) multiply(o XYZ, s float32) {
	v.X = o.X * s
	v.Y = o.Y * s
	v.Z = o.Z * s
}

func (v *XYZ) divide(o XYZ, s float32) {
	v.X = o.X / s
	v.Y = o.Y / s
	v.Z = o.Z / s
}

func (v *XYZ) normalize() {
	length := math32.Sqrt(v.X*v.X + v.Y*v.Y + v.Z*v.Z)
	v.X /= length
	v.Y /= length
	v.Z /= length
}

type arrayCoord [3]float32

func (v arrayCoord) plus(o arrayCoord) arrayCoord {
	return arrayCoord{v[0] + o[0], v[1] + o[1], v[2] + o[2]}
}

func (v *arrayCoord) add(a, b arrayCoord) {
	v[0] = a[0] + b[0]
	v[1] = a[1] + b[1]
	v[2] = a[2] + b[2]
}

//-----------------------------------------------------------------------------

func BenchmarkCoord_Plus(b *testing.B) {
	m := XYZ{1.1, 2.2, 3.3}
	n := XYZ{5.5, 6.6, 7.7}
	var o XYZ
	for i := 0; i < b.N; i++ {
		o = m.Plus(n)
	}
	_ = o
}

func BenchmarkCoord_Plus_Add(b *testing.B) {
	m := XYZ{1.1, 2.2, 3.3}
	n := XYZ{5.5, 6.6, 7.7}
	var o XYZ
	for i := 0; i < b.N; i++ {
		o = add(m, n)
	}
	_ = o
}

func BenchmarkCoord_Plus_Array(b *testing.B) {
	m := arrayCoord{1.1, 2.2, 3.3}
	n := arrayCoord{5.5, 6.6, 7.7}
	var o arrayCoord
	for i := 0; i < b.N; i++ {
		o = m.plus(n)
	}
	_ = o
}

func BenchmarkCoord_Plus_Self(b *testing.B) {
	m := XYZ{1.1, 2.2, 3.3}
	n := XYZ{5.5, 6.6, 7.7}
	for i := 0; i < b.N; i++ {
		m = m.Plus(n)
	}
	_ = m
}

func BenchmarkCoord_Plus_ArraySelf(b *testing.B) {
	m := arrayCoord{1.1, 2.2, 3.3}
	n := arrayCoord{5.5, 6.6, 7.7}
	for i := 0; i < b.N; i++ {
		m = m.plus(n)
	}
	_ = m
}

func BenchmarkCoord_Plus_ByRef(b *testing.B) {
	m := XYZ{1.1, 2.2, 3.3}
	n := XYZ{5.5, 6.6, 7.7}
	var o XYZ
	for i := 0; i < b.N; i++ {
		o.add(m, n)
	}
	_ = o
}

func BenchmarkCoord_Plus_ArrayByRef(b *testing.B) {
	m := arrayCoord{1.1, 2.2, 3.3}
	n := arrayCoord{5.5, 6.6, 7.7}
	var o arrayCoord
	for i := 0; i < b.N; i++ {
		o.add(m, n)
	}
	_ = o
}

func BenchmarkCoord_Plus_ByRefSelf(b *testing.B) {
	m := XYZ{1.1, 2.2, 3.3}
	n := XYZ{5.5, 6.6, 7.7}
	for i := 0; i < b.N; i++ {
		m.add(m, n)
	}
	_ = m
}

func BenchmarkCoord_Plus_ArrayByRefSelf(b *testing.B) {
	m := arrayCoord{1.1, 2.2, 3.3}
	n := arrayCoord{5.5, 6.6, 7.7}
	for i := 0; i < b.N; i++ {
		m.add(m, n)
	}
	_ = m
}

//-----------------------------------------------------------------------------

func BenchmarkCoord_Times(b *testing.B) {
	m := XYZ{1.1, 2.2, 3.3}
	n := float32(5.5)
	var o XYZ
	for i := 0; i < b.N; i++ {
		o = m.Times(n)
	}
	_ = o
}

func BenchmarkCoord_Times_Self(b *testing.B) {
	m := XYZ{1.1, 2.2, 3.3}
	n := float32(5.5)
	for i := 0; i < b.N; i++ {
		m = m.Times(n)
	}
	_ = m
}

func BenchmarkCoord_Times_ByRef(b *testing.B) {
	m := XYZ{1.1, 2.2, 3.3}
	n := float32(5.5)
	var o XYZ
	for i := 0; i < b.N; i++ {
		o.multiply(m, n)
	}
	_ = o
}

func BenchmarkCoord_Times_ByRefSelf(b *testing.B) {
	m := XYZ{1.1, 2.2, 3.3}
	n := float32(5.5)
	for i := 0; i < b.N; i++ {
		m.multiply(m, n)
	}
	_ = m
}

//-----------------------------------------------------------------------------

func BenchmarkCoord_Slash(b *testing.B) {
	m := XYZ{1.1, 2.2, 3.3}
	n := float32(5.5)
	var o XYZ
	for i := 0; i < b.N; i++ {
		o = m.Slash(n)
	}
	_ = o
}

func BenchmarkCoord_Slash_Self(b *testing.B) {
	m := XYZ{1.1, 2.2, 3.3}
	n := float32(5.5)
	for i := 0; i < b.N; i++ {
		m = m.Slash(n)
	}
	_ = m
}

func BenchmarkCoord_Slash_ByRef(b *testing.B) {
	m := XYZ{1.1, 2.2, 3.3}
	n := float32(5.5)
	var o XYZ
	for i := 0; i < b.N; i++ {
		o.divide(m, n)
	}
	_ = o
}

func BenchmarkCoord_Slash_ByRefSelf(b *testing.B) {
	m := XYZ{1.1, 2.2, 3.3}
	n := float32(5.5)
	for i := 0; i < b.N; i++ {
		m.divide(m, n)
	}
	_ = m
}

//-----------------------------------------------------------------------------

func BenchmarkCoord_Normalized(b *testing.B) {
	m := XYZ{1.1, 2.2, 3.3}
	var o XYZ
	for i := 0; i < b.N; i++ {
		o = m.Normalized()
	}
	_ = o
}

func BenchmarkCoord_Normalized_Self(b *testing.B) {
	m := XYZ{1.1, 2.2, 3.3}
	for i := 0; i < b.N; i++ {
		m = m.Normalized()
	}
	_ = m
}

func BenchmarkCoord_Normalize_ByRef(b *testing.B) {
	m := XYZ{1.1, 2.2, 3.3}
	for i := 0; i < b.N; i++ {
		m.normalize()
	}
	_ = m
}

////////////////////////////////////////////////////////////////////////////////
