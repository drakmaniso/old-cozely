// Copyright (c) 2013-2016 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package geom

import "github.com/drakmaniso/glam/math"

//------------------------------------------------------------------------------

/*
Mat4 is a single-precision matrix with 4 columns and 4 rows.
*/
type Mat4 [4][4]float32

//------------------------------------------------------------------------------

// Mat4Identity returns a 4x4 Identity matrix.
func Mat4Identity() Mat4 {
	return Mat4{
		{1, 0, 0, 0},
		{0, 1, 0, 0},
		{0, 0, 1, 0},
		{0, 0, 0, 1},
	}
}

//------------------------------------------------------------------------------

// Perspective returns a perspective projection matrix.
//
// See also PerspectiveFrustum.
func Perspective(fieldOfView, aspectRatio, near, far float32) Mat4 {
	f := float32(1.0) / math.Tan(fieldOfView/float32(2.0))

	return Mat4{
		{f / aspectRatio, 0, 0, 0},
		{0, f, 0, 0},
		{0, 0, (far + near) / (near - far), -1},
		{0, 0, (2 * far * near) / (near - far), 0},
	}
}

// PerspectiveFrustum returns a perspective projection matrix.
//
// See also Perspective.
func PerspectiveFrustum(left, right, bottom, top, near, far float32) Mat4 {
	return Mat4{
		{(2 * near) / (right - left), 0, 0, 0},
		{0, (2 * near) / (top - bottom), 0, 0},
		{(right + left) / (right - left), (top + bottom) / (top - bottom), -(far + near) / (far - near), -1},
		{0, 0, -(2 * far * near) / (far - near), 0},
	}
}

// Orthographic returns an orthographic (parallel) projection matrix.
// (zoom is the height of the projection plane).
//
// See also OrthographicFrustum.
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

// OrthographicFrustum returns an orthographic (parallel) projection matrix.
//
// See also Orthographic.
func OrthographicFrustum(left, right, bottom, top, near, far float32) Mat4 {
	return Mat4{
		{2 / (right - left), 0, 0, 0},
		{0, 2 / (top - bottom), 0, 0},
		{0, 0, -2 / (far - near), 0},
		{-(right + left) / (right - left), -(top + bottom) / (top - bottom), -(far + near) / (far - near), 1},
	}
}

//------------------------------------------------------------------------------

// Translation returns a translation matrix.
func Translation(t Vec3) Mat4 {
	return Mat4{
		{1, 0, 0, 0},
		{0, 1, 0, 0},
		{0, 0, 1, 0},
		{t.X, t.Y, t.Z, 1},
	}
}

//------------------------------------------------------------------------------

// Rotation returns a rotation matrix.
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

//------------------------------------------------------------------------------

// Times returns the matrix product with another matrix.
//
// See also TimesVector.
func (m *Mat4) Times(o *Mat4) Mat4 {
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

//------------------------------------------------------------------------------

// LookAt returns a transform from world space into the specific eye space
// that the projective matrix functions (Perspective, OrthographicFrustum, ...)
// are designed to expect.
//
// See also Perspective and OrthographicFrustum.
func LookAt(eye, center, up Vec3) Mat4 {
	center = center.Minus(eye)
	f := center.Normalized()
	u := up.Normalized()
	s := f.Cross(u).Normalized()
	u = s.Cross(f)

	return Mat4{
		{s.X, u.X, -f.X, 0},
		{s.Y, u.Y, -f.Y, 0},
		{s.Z, u.Z, -f.Z, 0},
		{-s.Dot(eye), -u.Dot(eye), f.Dot(eye), 1},
	}
}

//------------------------------------------------------------------------------
