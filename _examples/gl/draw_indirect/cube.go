// Copyright (c) 2013-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package main

//------------------------------------------------------------------------------

import "github.com/drakmaniso/cozely/space"

//------------------------------------------------------------------------------

func cube() mesh {
	return mesh{
		// Front Face
		{space.Coord{-0.5, -0.5, +0.5}},
		{space.Coord{+0.5, +0.5, +0.5}},
		{space.Coord{-0.5, +0.5, +0.5}},
		{space.Coord{-0.5, -0.5, +0.5}},
		{space.Coord{+0.5, -0.5, +0.5}},
		{space.Coord{+0.5, +0.5, +0.5}},
		// Back Face
		{space.Coord{-0.5, -0.5, -0.5}},
		{space.Coord{-0.5, +0.5, -0.5}},
		{space.Coord{+0.5, +0.5, -0.5}},
		{space.Coord{-0.5, -0.5, -0.5}},
		{space.Coord{+0.5, +0.5, -0.5}},
		{space.Coord{+0.5, -0.5, -0.5}},
		// Right Face
		{space.Coord{+0.5, -0.5, +0.5}},
		{space.Coord{+0.5, +0.5, -0.5}},
		{space.Coord{+0.5, +0.5, +0.5}},
		{space.Coord{+0.5, -0.5, +0.5}},
		{space.Coord{+0.5, -0.5, -0.5}},
		{space.Coord{+0.5, +0.5, -0.5}},
		// Left Face
		{space.Coord{-0.5, -0.5, +0.5}},
		{space.Coord{-0.5, +0.5, +0.5}},
		{space.Coord{-0.5, +0.5, -0.5}},
		{space.Coord{-0.5, -0.5, +0.5}},
		{space.Coord{-0.5, +0.5, -0.5}},
		{space.Coord{-0.5, -0.5, -0.5}},
		// Bottom Face
		{space.Coord{-0.5, -0.5, +0.5}},
		{space.Coord{-0.5, -0.5, -0.5}},
		{space.Coord{+0.5, -0.5, +0.5}},
		{space.Coord{-0.5, -0.5, -0.5}},
		{space.Coord{+0.5, -0.5, -0.5}},
		{space.Coord{+0.5, -0.5, +0.5}},
		// Top Face
		{space.Coord{-0.5, +0.5, +0.5}},
		{space.Coord{+0.5, +0.5, +0.5}},
		{space.Coord{-0.5, +0.5, -0.5}},
		{space.Coord{-0.5, +0.5, -0.5}},
		{space.Coord{+0.5, +0.5, +0.5}},
		{space.Coord{+0.5, +0.5, -0.5}},
	}
}

//------------------------------------------------------------------------------
