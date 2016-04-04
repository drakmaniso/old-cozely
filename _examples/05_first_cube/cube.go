// Copyright (c) 2013-2016 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package main

//------------------------------------------------------------------------------

import (
	"github.com/drakmaniso/glam/color"
	. "github.com/drakmaniso/glam/geom"
)

//------------------------------------------------------------------------------

var (
	purple = color.RGB{R: 0.2, G: 0, B: 0.6}
	orange = color.RGB{R: 0.8, G: 0.3, B: 0}
	green  = color.RGB{R: 0, G: 0.3, B: 0.1}
)

//------------------------------------------------------------------------------

func cube() []perVertex {
	return []perVertex{
		// Front Face
		{Vec3{-0.5, -0.5, +0.5}, purple},
		{Vec3{+0.5, +0.5, +0.5}, purple},
		{Vec3{-0.5, +0.5, +0.5}, purple},
		{Vec3{-0.5, -0.5, +0.5}, purple},
		{Vec3{+0.5, -0.5, +0.5}, purple},
		{Vec3{+0.5, +0.5, +0.5}, purple},
		// Back Face
		{Vec3{-0.5, -0.5, -0.5}, purple},
		{Vec3{-0.5, +0.5, -0.5}, purple},
		{Vec3{+0.5, +0.5, -0.5}, purple},
		{Vec3{-0.5, -0.5, -0.5}, purple},
		{Vec3{+0.5, +0.5, -0.5}, purple},
		{Vec3{+0.5, -0.5, -0.5}, purple},
		// Right Face
		{Vec3{+0.5, -0.5, +0.5}, green},
		{Vec3{+0.5, +0.5, -0.5}, green},
		{Vec3{+0.5, +0.5, +0.5}, green},
		{Vec3{+0.5, -0.5, +0.5}, green},
		{Vec3{+0.5, -0.5, -0.5}, green},
		{Vec3{+0.5, +0.5, -0.5}, green},
		// Left Face
		{Vec3{-0.5, -0.5, +0.5}, green},
		{Vec3{-0.5, +0.5, +0.5}, green},
		{Vec3{-0.5, +0.5, -0.5}, green},
		{Vec3{-0.5, -0.5, +0.5}, green},
		{Vec3{-0.5, +0.5, -0.5}, green},
		{Vec3{-0.5, -0.5, -0.5}, green},
		// Bottom Face
		{Vec3{-0.5, -0.5, +0.5}, orange},
		{Vec3{-0.5, -0.5, -0.5}, orange},
		{Vec3{+0.5, -0.5, +0.5}, orange},
		{Vec3{-0.5, -0.5, -0.5}, orange},
		{Vec3{+0.5, -0.5, -0.5}, orange},
		{Vec3{+0.5, -0.5, +0.5}, orange},
		// Top Face
		{Vec3{-0.5, +0.5, +0.5}, orange},
		{Vec3{+0.5, +0.5, +0.5}, orange},
		{Vec3{-0.5, +0.5, -0.5}, orange},
		{Vec3{-0.5, +0.5, -0.5}, orange},
		{Vec3{+0.5, +0.5, +0.5}, orange},
		{Vec3{+0.5, +0.5, -0.5}, orange},
	}
}

//------------------------------------------------------------------------------
