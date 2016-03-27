// Copyright (c) 2013-2016 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package plane

import (
	"github.com/drakmaniso/glam/geom"
	"github.com/drakmaniso/glam/math"
)

//------------------------------------------------------------------------------

// Identity matrix.
func Identity() geom.Mat3 {
	return geom.Mat3{
		{1, 0, 0},
		{0, 1, 0},
		{0, 0, 1},
	}
}

//------------------------------------------------------------------------------

// Apply a tranformation matrix to a vector (i.e. returns matrix multiplied by
// column vector).
func Apply(m geom.Mat3, v geom.Vec3) geom.Vec3 {
	return geom.Vec3{
		X: m[0][0]*v.X + m[1][0]*v.Y + m[2][0]*v.Z,
		Y: m[0][1]*v.X + m[1][1]*v.Y + m[2][1]*v.Z,
		Z: m[0][2]*v.X + m[1][2]*v.Y + m[2][2]*v.Z,
	}
}

//------------------------------------------------------------------------------

// Translation by a vector.
func Translation(t geom.Vec2) geom.Mat3 {
	return geom.Mat3{
		{1, 0, 0},
		{0, 1, 0},
		{t.X, t.Y, 1},
	}
}

//------------------------------------------------------------------------------

// Rotation.
func Rotation(angle float32) geom.Mat3 {
	c := math.Cos(angle)
	s := math.Sin(angle)

	return geom.Mat3{
		{c, -s, 0},
		{s, c, 0},
		{0, 0, 1},
	}
}

//------------------------------------------------------------------------------

// RotationAround a point.
func RotationAround(angle float32, center geom.Vec2) geom.Mat3 {
	c := math.Cos(angle)
	s := math.Sin(angle)

	return geom.Mat3{
		{c, -s, 0},
		{s, c, 0},
		{center.X - c*center.X - s*center.Y, center.Y + s*center.X - c*center.Y, 1},
	}
}

//------------------------------------------------------------------------------

// Scaling and/or mirror along both axis.
func Scaling(s geom.Vec2) geom.Mat3 {
	return geom.Mat3{
		{s.X, 0, 0},
		{0, s.Y, 0},
		{0, 0, 1},
	}
}

//------------------------------------------------------------------------------

// ScalingAround a point (and/or mirror).
func ScalingAround(s geom.Vec2, center geom.Vec2) geom.Mat3 {
	return geom.Mat3{
		{s.X, 0, 0},
		{0, s.Y, 0},
		{center.X - center.X*s.X, center.Y - center.Y*s.Y, 1},
	}
}

//------------------------------------------------------------------------------

// Shearing along both axis.
func Shearing(s geom.Vec2) geom.Mat3 {
	return geom.Mat3{
		{1, s.Y, 0},
		{s.X, 1, 0},
		{0, 0, 1},
	}
}

//------------------------------------------------------------------------------

// Viewport returns a matrix that scale to an aspect ratio and zoom.
func Viewport(zoom, aspectRatio float32) geom.Mat3 {
	height := zoom / 2
	width := height * aspectRatio
	return Scaling(geom.Vec2{X: width, Y: height})
}

//------------------------------------------------------------------------------
