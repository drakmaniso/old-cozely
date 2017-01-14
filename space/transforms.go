// Copyright (c) 2013-2017 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package space

//------------------------------------------------------------------------------

import "github.com/drakmaniso/glam/math"

//------------------------------------------------------------------------------

// Identity matrix.
func Identity() Matrix {
	return Matrix{
		{1, 0, 0, 0},
		{0, 1, 0, 0},
		{0, 0, 1, 0},
		{0, 0, 0, 1},
	}
}

//------------------------------------------------------------------------------

// Apply a tranformation matrix to a vector (i.e. returns matrix multiplied by
// column vector).
func Apply(m Matrix, v Homogen) Homogen {
	return Homogen{
		X: m[0][0]*v.X + m[1][0]*v.Y + m[2][0]*v.Z + m[3][0]*v.W,
		Y: m[0][1]*v.X + m[1][1]*v.Y + m[2][1]*v.Z + m[3][1]*v.W,
		Z: m[0][2]*v.X + m[1][2]*v.Y + m[2][2]*v.Z + m[3][2]*v.W,
		W: m[0][3]*v.X + m[1][3]*v.Y + m[2][3]*v.Z + m[3][3]*v.W,
	}
}

//------------------------------------------------------------------------------

// Translation by a vector.
func Translation(t Vector) Matrix {
	x, y, z := t.Cartesian()

	return Matrix{
		{1, 0, 0, 0},
		{0, 1, 0, 0},
		{0, 0, 1, 0},
		{x, y, z, 1},
	}
}

//------------------------------------------------------------------------------

// Rotation around an axis.
func Rotation(angle float32, axis Vector) Matrix {
	c := math.Cos(angle)
	s := math.Sin(angle)

	x, y, z := axis.Cartesian()

	return Matrix{
		{c + x*x*(1-c), -z*s + x*y*(1-c), y*s + x*z*(1-c), 0},
		{z*s + y*x*(1-c), c + y*y*(1-c), -x*s + y*z*(1-c), 0},
		{-y*s + z*x*(1-c), x*s + z*y*(1-c), c + z*z*(1-c), 0},
		{0, 0, 0, 1},
	}
}

//------------------------------------------------------------------------------

// EulerYXZ returns a transform matrix corresponding to a sequence of three
// intrinsic rotations around the (local) axis Y, X, Z, in that order.
func EulerYXZ(angleX, angleY, angleZ float32) Matrix {
	c1 := math.Cos(angleY)
	s1 := math.Sin(angleY)
	c2 := math.Cos(angleX)
	s2 := math.Sin(angleX)
	c3 := math.Cos(angleZ)
	s3 := math.Sin(angleZ)

	return Matrix{
		{c1*c3 + s1*s2*s3, c2 * s3, c1*s2*s3 - c3*s1},
		{c3*s1*s2 - c1*s3, c2 * c3, c1*c3*s2 + s1*s3},
		{c2 * s1, -s2, c1 * c2},
		{0, 0, 0, 1},
	}
}

// EulerXYZ returns a transform matrix corresponding to a sequence of three
// intrinsic rotations around the (local) axis X, Y, Z, in that order.
func EulerXYZ(angleX, angleY, angleZ float32) Matrix {
	c1 := math.Cos(angleX)
	s1 := math.Sin(angleX)
	c2 := math.Cos(angleY)
	s2 := math.Sin(angleY)
	c3 := math.Cos(angleZ)
	s3 := math.Sin(angleZ)

	return Matrix{
		{c2 * c3, c1*s3 + c3*s1*s2, s1*s3 - c1*c3*s2},
		{-c2 * s3, c1*c3 - s1*s2*s3, c3*s1 + c1*s2*s3},
		{s2, -c2 * s1, c1 * c2},
		{0, 0, 0, 1},
	}
}

// EulerZYX returns a transform matrix corresponding to a sequence of three
// intrinsic rotations around the (local) axis Z, Y, X, in that order.
func EulerZYX(angleX, angleY, angleZ float32) Matrix {
	c1 := math.Cos(angleZ)
	s1 := math.Sin(angleZ)
	c2 := math.Cos(angleY)
	s2 := math.Sin(angleY)
	c3 := math.Cos(angleX)
	s3 := math.Sin(angleX)

	return Matrix{
		{c1 * c2, c2 * s1, -s2},
		{c1*s2*s3 - c3*s1, c1*c3 + s1*s2*s3, c2 * s3},
		{s1*s3 + c1*c3*s2, c3*s1*s2 - c1*s3, c2 * c3},
		{0, 0, 0, 1},
	}
}

// EulerXZY returns a transform matrix corresponding to a sequence of three
// intrinsic rotations around the (local) axis X, Z, Y, in that order.
func EulerXZY(angleX, angleY, angleZ float32) Matrix {
	c1 := math.Cos(angleX)
	s1 := math.Sin(angleX)
	c2 := math.Cos(angleZ)
	s2 := math.Sin(angleZ)
	c3 := math.Cos(angleY)
	s3 := math.Sin(angleY)

	return Matrix{
		{c2 * c3, s1*s3 + c1*c3*s2, c3*s1*s2 - c1*s3},
		{-s2, c1 * c2, c2 * s1},
		{c2 * s3, c1*s2*s3 - c3*s1, c1*c3 + s1*s2*s3},
		{0, 0, 0, 1},
	}
}

// EulerYZX returns a transform matrix corresponding to a sequence of three
// intrinsic rotations around the (local) axis Y, Z, X, in that order.
func EulerYZX(angleX, angleY, angleZ float32) Matrix {
	c1 := math.Cos(angleY)
	s1 := math.Sin(angleY)
	c2 := math.Cos(angleZ)
	s2 := math.Sin(angleZ)
	c3 := math.Cos(angleX)
	s3 := math.Sin(angleX)

	return Matrix{
		{c1 * c2, s2, -c2 * s1},
		{s1*s3 - c1*c3*s2, c2 * c3, c1*s3 + c3*s1*s2},
		{c3*s1 + c1*s2*s3, -c2 * s3, c1*c3 - s1*s2*s3},
		{0, 0, 0, 1},
	}
}

// EulerZXY returns a transform matrix corresponding to a sequence of three
// intrinsic rotations around the (local) axis Z, X, Y, in that order.
func EulerZXY(angleX, angleY, angleZ float32) Matrix {
	c1 := math.Cos(angleZ)
	s1 := math.Sin(angleZ)
	c2 := math.Cos(angleX)
	s2 := math.Sin(angleX)
	c3 := math.Cos(angleY)
	s3 := math.Sin(angleY)

	return Matrix{
		{c1*c3 - s1*s2*s3, c3*s1 + c1*s2*s3, -c2 * s3},
		{-c2 * s1, c1 * c2, s2},
		{c1*s3 + c3*s1*s2, s1*s3 - c1*c3*s2, c2 * c3},
		{0, 0, 0, 1},
	}
}

//------------------------------------------------------------------------------

// Scaling along the 3 axis.
func Scaling(s Vector) Matrix {
	x, y, z := s.Cartesian()

	return Matrix{
		{x, 0, 0, 0},
		{0, y, 0, 0},
		{0, 0, z, 0},
		{0, 0, 0, 1},
	}
}

//------------------------------------------------------------------------------

// LookAt returns a transform which put eye at origin and target along
// negative Z. In other words, if a projection matrix is applied to the result,
// target will be in the center of the viewport.
func LookAt(eye, target, up Vector) Matrix {
	e := NewCoord(eye)
	f := NewCoord(target).Minus(e).Normalized()
	u := NewCoord(up).Normalized()
	s := f.Cross(u).Normalized()
	u = s.Cross(f)

	return Matrix{
		{s.X, u.X, -f.X, 0},
		{s.Y, u.Y, -f.Y, 0},
		{s.Z, u.Z, -f.Z, 0},
		{-s.Dot(e), -u.Dot(e), f.Dot(e), 1},
	}
}

//------------------------------------------------------------------------------

// Perspective returns a perspective projection matrix.
func Perspective(fieldOfView, aspectRatio, near, far float32) Matrix {
	f := float32(1.0) / math.Tan(fieldOfView/float32(2.0))

	return Matrix{
		{f / aspectRatio, 0, 0, 0},
		{0, f, 0, 0},
		{0, 0, (far + near) / (near - far), -1},
		{0, 0, (2 * far * near) / (near - far), 0},
	}
}

// PerspectiveFrustum returns a perspective projection matrix.
func PerspectiveFrustum(left, right, bottom, top, near, far float32) Matrix {
	return Matrix{
		{(2 * near) / (right - left), 0, 0, 0},
		{0, (2 * near) / (top - bottom), 0, 0},
		{(right + left) / (right - left), (top + bottom) / (top - bottom), -(far + near) / (far - near), -1},
		{0, 0, -(2 * far * near) / (far - near), 0},
	}
}

// Orthographic returns an orthographic (parallel) projection matrix.
// (zoom is the height of the projection plane).
func Orthographic(zoom, aspectRatio, near, far float32) Matrix {
	top := zoom / 2
	right := top * aspectRatio
	return Matrix{
		{1 / right, 0, 0, 0},
		{0, 1 / top, 0, 0},
		{0, 0, -2 / (far - near), 0},
		{0, 0, -(far + near) / (far - near), 1},
	}
}

// OrthographicFrustum returns an orthographic (parallel) projection matrix.
func OrthographicFrustum(left, right, bottom, top, near, far float32) Matrix {
	return Matrix{
		{2 / (right - left), 0, 0, 0},
		{0, 2 / (top - bottom), 0, 0},
		{0, 0, -2 / (far - near), 0},
		{-(right + left) / (right - left), -(top + bottom) / (top - bottom), -(far + near) / (far - near), 1},
	}
}

//------------------------------------------------------------------------------
