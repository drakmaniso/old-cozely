// Copyright (c) 2013-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package main

import (
	"github.com/cozely/cozely/coord"
)


////////////////////////////////////////////////////////////////////////////////

func cube() mesh {
	return mesh{
		// Front Face
		{coord.XYZ{-0.5, -0.5, +0.5}},
		{coord.XYZ{+0.5, +0.5, +0.5}},
		{coord.XYZ{-0.5, +0.5, +0.5}},
		{coord.XYZ{-0.5, -0.5, +0.5}},
		{coord.XYZ{+0.5, -0.5, +0.5}},
		{coord.XYZ{+0.5, +0.5, +0.5}},
		// Back Face
		{coord.XYZ{-0.5, -0.5, -0.5}},
		{coord.XYZ{-0.5, +0.5, -0.5}},
		{coord.XYZ{+0.5, +0.5, -0.5}},
		{coord.XYZ{-0.5, -0.5, -0.5}},
		{coord.XYZ{+0.5, +0.5, -0.5}},
		{coord.XYZ{+0.5, -0.5, -0.5}},
		// Right Face
		{coord.XYZ{+0.5, -0.5, +0.5}},
		{coord.XYZ{+0.5, +0.5, -0.5}},
		{coord.XYZ{+0.5, +0.5, +0.5}},
		{coord.XYZ{+0.5, -0.5, +0.5}},
		{coord.XYZ{+0.5, -0.5, -0.5}},
		{coord.XYZ{+0.5, +0.5, -0.5}},
		// Left Face
		{coord.XYZ{-0.5, -0.5, +0.5}},
		{coord.XYZ{-0.5, +0.5, +0.5}},
		{coord.XYZ{-0.5, +0.5, -0.5}},
		{coord.XYZ{-0.5, -0.5, +0.5}},
		{coord.XYZ{-0.5, +0.5, -0.5}},
		{coord.XYZ{-0.5, -0.5, -0.5}},
		// Bottom Face
		{coord.XYZ{-0.5, -0.5, +0.5}},
		{coord.XYZ{-0.5, -0.5, -0.5}},
		{coord.XYZ{+0.5, -0.5, +0.5}},
		{coord.XYZ{-0.5, -0.5, -0.5}},
		{coord.XYZ{+0.5, -0.5, -0.5}},
		{coord.XYZ{+0.5, -0.5, +0.5}},
		// Top Face
		{coord.XYZ{-0.5, +0.5, +0.5}},
		{coord.XYZ{+0.5, +0.5, +0.5}},
		{coord.XYZ{-0.5, +0.5, -0.5}},
		{coord.XYZ{-0.5, +0.5, -0.5}},
		{coord.XYZ{+0.5, +0.5, +0.5}},
		{coord.XYZ{+0.5, +0.5, -0.5}},
	}
}

////////////////////////////////////////////////////////////////////////////////
