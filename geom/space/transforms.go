// Copyright (c) 2013-2016 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package space

import (
	"github.com/drakmaniso/glam/geom"
	"github.com/drakmaniso/glam/math"
)

//------------------------------------------------------------------------------

// Transform a vector by a matrix.
func Transform(m geom.Mat4, v geom.Vec4) geom.Vec4 {
	return geom.Vec4{
		X: m[0][0]*v.X + m[1][0]*v.Y + m[2][0]*v.Z + m[3][0]*v.W,
		Y: m[0][1]*v.X + m[1][1]*v.Y + m[2][1]*v.Z + m[3][1]*v.W,
		Z: m[0][2]*v.X + m[1][2]*v.Y + m[2][2]*v.Z + m[3][2]*v.W,
		W: m[0][3]*v.X + m[1][3]*v.Y + m[2][3]*v.Z + m[3][3]*v.W,
	}
}

//------------------------------------------------------------------------------

// Translation returns a translation matrix.
func Translation(t geom.Vec3) geom.Mat4 {
	return geom.Mat4{
		{1, 0, 0, 0},
		{0, 1, 0, 0},
		{0, 0, 1, 0},
		{t.X, t.Y, t.Z, 1},
	}
}

//------------------------------------------------------------------------------

// Rotation returns a rotation matrix.
func Rotation(angle float32, axis geom.Vec3) geom.Mat4 {
	c := math.Cos(angle)
	s := math.Sin(angle)

	return geom.Mat4{
		{c + axis.X*axis.X*(1-c), -axis.Z*s + axis.X*axis.Y*(1-c), axis.Y*s + axis.X*axis.Z*(1-c), 0},
		{axis.Z*s + axis.Y*axis.X*(1-c), c + axis.Y*axis.Y*(1-c), -axis.X*s + axis.Y*axis.Z*(1-c), 0},
		{-axis.Y*s + axis.Z*axis.X*(1-c), axis.X*s + axis.Z*axis.Y*(1-c), c + axis.Z*axis.Z*(1-c), 0},
		{0, 0, 0, 1},
	}
}

//------------------------------------------------------------------------------

// Scaling returns a scaling matrix.
func Scaling(s geom.Vec3) geom.Mat4 {
	return geom.Mat4{
		{s.X, 0, 0, 0},
		{0, s.Y, 0, 0},
		{0, 0, s.Z, 0},
		{0, 0, 0, 1},
	}
}

//------------------------------------------------------------------------------

// Identity returns an Identity matrix.
func Identity() geom.Mat4 {
	return geom.Mat4{
		{1, 0, 0, 0},
		{0, 1, 0, 0},
		{0, 0, 1, 0},
		{0, 0, 0, 1},
	}
}

//------------------------------------------------------------------------------

// LookAt returns a transform from world space into the specific eye space
// that the projective matrix functions (Perspective, OrthographicFrustum, ...)
// are designed to expect.
//
// See also Perspective and OrthographicFrustum.
func LookAt(eye, center, up geom.Vec3) geom.Mat4 {
	f := center.Minus(eye).Normalized()
	u := up.Normalized()
	s := f.Cross(u).Normalized()
	u = s.Cross(f)

	return geom.Mat4{
		{s.X, u.X, -f.X, 0},
		{s.Y, u.Y, -f.Y, 0},
		{s.Z, u.Z, -f.Z, 0},
		{-s.Dot(eye), -u.Dot(eye), f.Dot(eye), 1},
	}
}

//------------------------------------------------------------------------------

// Perspective returns a perspective projection matrix.
//
// See also PerspectiveFrustum.
func Perspective(fieldOfView, aspectRatio, near, far float32) geom.Mat4 {
	f := float32(1.0) / math.Tan(fieldOfView/float32(2.0))

	return geom.Mat4{
		{f / aspectRatio, 0, 0, 0},
		{0, f, 0, 0},
		{0, 0, (far + near) / (near - far), -1},
		{0, 0, (2 * far * near) / (near - far), 0},
	}
}

// PerspectiveFrustum returns a perspective projection matrix.
//
// See also Perspective.
func PerspectiveFrustum(left, right, bottom, top, near, far float32) geom.Mat4 {
	return geom.Mat4{
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
func Orthographic(zoom, aspectRatio, near, far float32) geom.Mat4 {
	top := zoom / 2
	right := top * aspectRatio
	return geom.Mat4{
		{1 / right, 0, 0, 0},
		{0, 1 / top, 0, 0},
		{0, 0, -2 / (far - near), 0},
		{0, 0, -(far + near) / (far - near), 1},
	}
}

// OrthographicFrustum returns an orthographic (parallel) projection matrix.
//
// See also Orthographic.
func OrthographicFrustum(left, right, bottom, top, near, far float32) geom.Mat4 {
	return geom.Mat4{
		{2 / (right - left), 0, 0, 0},
		{0, 2 / (top - bottom), 0, 0},
		{0, 0, -2 / (far - near), 0},
		{-(right + left) / (right - left), -(top + bottom) / (top - bottom), -(far + near) / (far - near), 1},
	}
}

//------------------------------------------------------------------------------
