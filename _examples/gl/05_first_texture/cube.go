// Copyright (c) 2013-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package main

////////////////////////////////////////////////////////////////////////////////

import (
	"github.com/cozely/cozely/coord"
)

////////////////////////////////////////////////////////////////////////////////

func cube() mesh {
	return mesh{
		// Front Face
		{coord.XYZ{-0.5, -0.5, +0.5}, coord.XY{0, 1}},
		{coord.XYZ{+0.5, -0.5, +0.5}, coord.XY{1, 1}},
		{coord.XYZ{+0.5, +0.5, +0.5}, coord.XY{1, 0}},
		{coord.XYZ{-0.5, -0.5, +0.5}, coord.XY{0, 1}},
		{coord.XYZ{+0.5, +0.5, +0.5}, coord.XY{1, 0}},
		{coord.XYZ{-0.5, +0.5, +0.5}, coord.XY{0, 0}},
		// Back Face
		{coord.XYZ{+0.5, -0.5, -0.5}, coord.XY{0, 1}},
		{coord.XYZ{-0.5, -0.5, -0.5}, coord.XY{1, 1}},
		{coord.XYZ{-0.5, +0.5, -0.5}, coord.XY{1, 0}},
		{coord.XYZ{+0.5, -0.5, -0.5}, coord.XY{0, 1}},
		{coord.XYZ{-0.5, +0.5, -0.5}, coord.XY{1, 0}},
		{coord.XYZ{+0.5, +0.5, -0.5}, coord.XY{0, 0}},
		// Right Face
		{coord.XYZ{+0.5, -0.5, +0.5}, coord.XY{0, 1}},
		{coord.XYZ{+0.5, -0.5, -0.5}, coord.XY{1, 1}},
		{coord.XYZ{+0.5, +0.5, -0.5}, coord.XY{1, 0}},
		{coord.XYZ{+0.5, -0.5, +0.5}, coord.XY{0, 1}},
		{coord.XYZ{+0.5, +0.5, -0.5}, coord.XY{1, 0}},
		{coord.XYZ{+0.5, +0.5, +0.5}, coord.XY{0, 0}},
		// Left Face
		{coord.XYZ{-0.5, -0.5, -0.5}, coord.XY{0, 1}},
		{coord.XYZ{-0.5, -0.5, +0.5}, coord.XY{1, 1}},
		{coord.XYZ{-0.5, +0.5, +0.5}, coord.XY{1, 0}},
		{coord.XYZ{-0.5, -0.5, -0.5}, coord.XY{0, 1}},
		{coord.XYZ{-0.5, +0.5, +0.5}, coord.XY{1, 0}},
		{coord.XYZ{-0.5, +0.5, -0.5}, coord.XY{0, 0}},
		// Bottom Face
		{coord.XYZ{-0.5, -0.5, -0.5}, coord.XY{0, 1}},
		{coord.XYZ{+0.5, -0.5, -0.5}, coord.XY{1, 1}},
		{coord.XYZ{+0.5, -0.5, +0.5}, coord.XY{1, 0}},
		{coord.XYZ{-0.5, -0.5, -0.5}, coord.XY{0, 1}},
		{coord.XYZ{+0.5, -0.5, +0.5}, coord.XY{1, 0}},
		{coord.XYZ{-0.5, -0.5, +0.5}, coord.XY{0, 0}},
		// Top Face
		{coord.XYZ{-0.5, +0.5, +0.5}, coord.XY{0, 1}},
		{coord.XYZ{+0.5, +0.5, +0.5}, coord.XY{1, 1}},
		{coord.XYZ{+0.5, +0.5, -0.5}, coord.XY{1, 0}},
		{coord.XYZ{-0.5, +0.5, +0.5}, coord.XY{0, 1}},
		{coord.XYZ{+0.5, +0.5, -0.5}, coord.XY{1, 0}},
		{coord.XYZ{-0.5, +0.5, -0.5}, coord.XY{0, 0}},
	}
}

////////////////////////////////////////////////////////////////////////////////
