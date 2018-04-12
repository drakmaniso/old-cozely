// Copyright (c) 2013-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package main

////////////////////////////////////////////////////////////////////////////////

import (
	"github.com/cozely/cozely/colour"
	"github.com/cozely/cozely/coord"
)

////////////////////////////////////////////////////////////////////////////////

var (
	purple = colour.LRGB{0.2, 0, 0.6}
	orange = colour.LRGB{0.8, 0.3, 0}
	green  = colour.LRGB{0, 0.3, 0.1}
)

////////////////////////////////////////////////////////////////////////////////

func cube() mesh {
	return mesh{
		// Front Face
		{coord.XYZ{-0.5, -0.5, +0.5}, purple},
		{coord.XYZ{+0.5, +0.5, +0.5}, purple},
		{coord.XYZ{-0.5, +0.5, +0.5}, purple},
		{coord.XYZ{-0.5, -0.5, +0.5}, purple},
		{coord.XYZ{+0.5, -0.5, +0.5}, purple},
		{coord.XYZ{+0.5, +0.5, +0.5}, purple},
		// Back Face
		{coord.XYZ{-0.5, -0.5, -0.5}, purple},
		{coord.XYZ{-0.5, +0.5, -0.5}, purple},
		{coord.XYZ{+0.5, +0.5, -0.5}, purple},
		{coord.XYZ{-0.5, -0.5, -0.5}, purple},
		{coord.XYZ{+0.5, +0.5, -0.5}, purple},
		{coord.XYZ{+0.5, -0.5, -0.5}, purple},
		// Right Face
		{coord.XYZ{+0.5, -0.5, +0.5}, green},
		{coord.XYZ{+0.5, +0.5, -0.5}, green},
		{coord.XYZ{+0.5, +0.5, +0.5}, green},
		{coord.XYZ{+0.5, -0.5, +0.5}, green},
		{coord.XYZ{+0.5, -0.5, -0.5}, green},
		{coord.XYZ{+0.5, +0.5, -0.5}, green},
		// Left Face
		{coord.XYZ{-0.5, -0.5, +0.5}, green},
		{coord.XYZ{-0.5, +0.5, +0.5}, green},
		{coord.XYZ{-0.5, +0.5, -0.5}, green},
		{coord.XYZ{-0.5, -0.5, +0.5}, green},
		{coord.XYZ{-0.5, +0.5, -0.5}, green},
		{coord.XYZ{-0.5, -0.5, -0.5}, green},
		// Bottom Face
		{coord.XYZ{-0.5, -0.5, +0.5}, orange},
		{coord.XYZ{-0.5, -0.5, -0.5}, orange},
		{coord.XYZ{+0.5, -0.5, +0.5}, orange},
		{coord.XYZ{-0.5, -0.5, -0.5}, orange},
		{coord.XYZ{+0.5, -0.5, -0.5}, orange},
		{coord.XYZ{+0.5, -0.5, +0.5}, orange},
		// Top Face
		{coord.XYZ{-0.5, +0.5, +0.5}, orange},
		{coord.XYZ{+0.5, +0.5, +0.5}, orange},
		{coord.XYZ{-0.5, +0.5, -0.5}, orange},
		{coord.XYZ{-0.5, +0.5, -0.5}, orange},
		{coord.XYZ{+0.5, +0.5, +0.5}, orange},
		{coord.XYZ{+0.5, +0.5, -0.5}, orange},
	}
}

////////////////////////////////////////////////////////////////////////////////
