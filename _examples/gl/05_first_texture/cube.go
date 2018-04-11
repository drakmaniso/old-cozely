// Copyright (c) 2013-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package main

////////////////////////////////////////////////////////////////////////////////

import (
	"github.com/cozely/cozely/plane"
	"github.com/cozely/cozely/space"
)

////////////////////////////////////////////////////////////////////////////////

func cube() mesh {
	return mesh{
		// Front Face
		{space.Coord{-0.5, -0.5, +0.5}, plane.Coord{0, 1}},
		{space.Coord{+0.5, -0.5, +0.5}, plane.Coord{1, 1}},
		{space.Coord{+0.5, +0.5, +0.5}, plane.Coord{1, 0}},
		{space.Coord{-0.5, -0.5, +0.5}, plane.Coord{0, 1}},
		{space.Coord{+0.5, +0.5, +0.5}, plane.Coord{1, 0}},
		{space.Coord{-0.5, +0.5, +0.5}, plane.Coord{0, 0}},
		// Back Face
		{space.Coord{+0.5, -0.5, -0.5}, plane.Coord{0, 1}},
		{space.Coord{-0.5, -0.5, -0.5}, plane.Coord{1, 1}},
		{space.Coord{-0.5, +0.5, -0.5}, plane.Coord{1, 0}},
		{space.Coord{+0.5, -0.5, -0.5}, plane.Coord{0, 1}},
		{space.Coord{-0.5, +0.5, -0.5}, plane.Coord{1, 0}},
		{space.Coord{+0.5, +0.5, -0.5}, plane.Coord{0, 0}},
		// Right Face
		{space.Coord{+0.5, -0.5, +0.5}, plane.Coord{0, 1}},
		{space.Coord{+0.5, -0.5, -0.5}, plane.Coord{1, 1}},
		{space.Coord{+0.5, +0.5, -0.5}, plane.Coord{1, 0}},
		{space.Coord{+0.5, -0.5, +0.5}, plane.Coord{0, 1}},
		{space.Coord{+0.5, +0.5, -0.5}, plane.Coord{1, 0}},
		{space.Coord{+0.5, +0.5, +0.5}, plane.Coord{0, 0}},
		// Left Face
		{space.Coord{-0.5, -0.5, -0.5}, plane.Coord{0, 1}},
		{space.Coord{-0.5, -0.5, +0.5}, plane.Coord{1, 1}},
		{space.Coord{-0.5, +0.5, +0.5}, plane.Coord{1, 0}},
		{space.Coord{-0.5, -0.5, -0.5}, plane.Coord{0, 1}},
		{space.Coord{-0.5, +0.5, +0.5}, plane.Coord{1, 0}},
		{space.Coord{-0.5, +0.5, -0.5}, plane.Coord{0, 0}},
		// Bottom Face
		{space.Coord{-0.5, -0.5, -0.5}, plane.Coord{0, 1}},
		{space.Coord{+0.5, -0.5, -0.5}, plane.Coord{1, 1}},
		{space.Coord{+0.5, -0.5, +0.5}, plane.Coord{1, 0}},
		{space.Coord{-0.5, -0.5, -0.5}, plane.Coord{0, 1}},
		{space.Coord{+0.5, -0.5, +0.5}, plane.Coord{1, 0}},
		{space.Coord{-0.5, -0.5, +0.5}, plane.Coord{0, 0}},
		// Top Face
		{space.Coord{-0.5, +0.5, +0.5}, plane.Coord{0, 1}},
		{space.Coord{+0.5, +0.5, +0.5}, plane.Coord{1, 1}},
		{space.Coord{+0.5, +0.5, -0.5}, plane.Coord{1, 0}},
		{space.Coord{-0.5, +0.5, +0.5}, plane.Coord{0, 1}},
		{space.Coord{+0.5, +0.5, -0.5}, plane.Coord{1, 0}},
		{space.Coord{-0.5, +0.5, -0.5}, plane.Coord{0, 0}},
	}
}

////////////////////////////////////////////////////////////////////////////////
