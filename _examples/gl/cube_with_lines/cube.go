// Copyright (c) 2013-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package main

////////////////////////////////////////////////////////////////////////////////

import (
	"github.com/drakmaniso/cozely/colour"
	"github.com/drakmaniso/cozely/space"
)

////////////////////////////////////////////////////////////////////////////////

var (
	purple = colour.LRGB{R: 0.2, G: 0, B: 0.6}
	orange = colour.LRGB{R: 0.8, G: 0.3, B: 0}
	// purple = colour.LRGB{R: 0.8, G: 0.3, B: 0}
	// green  = colour.LRGB{R: 0.8, G: 0.3, B: 0}
	green = colour.LRGB{R: 0, G: 0.3, B: 0.1}

	black = colour.LRGB{R: 0, G: 0, B: 0}
)

////////////////////////////////////////////////////////////////////////////////

func cubeFaces() mesh {
	return mesh{
		// Front Face
		{space.Coord{-0.5, -0.5, +0.5}, purple},
		{space.Coord{+0.5, +0.5, +0.5}, purple},
		{space.Coord{-0.5, +0.5, +0.5}, purple},
		{space.Coord{-0.5, -0.5, +0.5}, purple},
		{space.Coord{+0.5, -0.5, +0.5}, purple},
		{space.Coord{+0.5, +0.5, +0.5}, purple},
		// Back Face
		{space.Coord{-0.5, -0.5, -0.5}, purple},
		{space.Coord{-0.5, +0.5, -0.5}, purple},
		{space.Coord{+0.5, +0.5, -0.5}, purple},
		{space.Coord{-0.5, -0.5, -0.5}, purple},
		{space.Coord{+0.5, +0.5, -0.5}, purple},
		{space.Coord{+0.5, -0.5, -0.5}, purple},
		// Right Face
		{space.Coord{+0.5, -0.5, +0.5}, green},
		{space.Coord{+0.5, +0.5, -0.5}, green},
		{space.Coord{+0.5, +0.5, +0.5}, green},
		{space.Coord{+0.5, -0.5, +0.5}, green},
		{space.Coord{+0.5, -0.5, -0.5}, green},
		{space.Coord{+0.5, +0.5, -0.5}, green},
		// Left Face
		{space.Coord{-0.5, -0.5, +0.5}, green},
		{space.Coord{-0.5, +0.5, +0.5}, green},
		{space.Coord{-0.5, +0.5, -0.5}, green},
		{space.Coord{-0.5, -0.5, +0.5}, green},
		{space.Coord{-0.5, +0.5, -0.5}, green},
		{space.Coord{-0.5, -0.5, -0.5}, green},
		// Bottom Face
		{space.Coord{-0.5, -0.5, +0.5}, orange},
		{space.Coord{-0.5, -0.5, -0.5}, orange},
		{space.Coord{+0.5, -0.5, +0.5}, orange},
		{space.Coord{-0.5, -0.5, -0.5}, orange},
		{space.Coord{+0.5, -0.5, -0.5}, orange},
		{space.Coord{+0.5, -0.5, +0.5}, orange},
		// Top Face
		{space.Coord{-0.5, +0.5, +0.5}, orange},
		{space.Coord{+0.5, +0.5, +0.5}, orange},
		{space.Coord{-0.5, +0.5, -0.5}, orange},
		{space.Coord{-0.5, +0.5, -0.5}, orange},
		{space.Coord{+0.5, +0.5, +0.5}, orange},
		{space.Coord{+0.5, +0.5, -0.5}, orange},
	}
}

func cubeEdges() mesh {
	return mesh{
		// Front Edges
		{space.Coord{-0.5, -0.5, +0.5}, black},
		{space.Coord{-0.5, +0.5, +0.5}, black},

		{space.Coord{-0.5, +0.5, +0.5}, black},
		{space.Coord{+0.5, +0.5, +0.5}, black},

		{space.Coord{+0.5, +0.5, +0.5}, black},
		{space.Coord{+0.5, -0.5, +0.5}, black},

		{space.Coord{+0.5, -0.5, +0.5}, black},
		{space.Coord{-0.5, -0.5, +0.5}, black},

		// Back Edges
		{space.Coord{-0.5, -0.5, -0.5}, black},
		{space.Coord{-0.5, +0.5, -0.5}, black},

		{space.Coord{-0.5, +0.5, -0.5}, black},
		{space.Coord{+0.5, +0.5, -0.5}, black},

		{space.Coord{+0.5, +0.5, -0.5}, black},
		{space.Coord{+0.5, -0.5, -0.5}, black},

		{space.Coord{+0.5, -0.5, -0.5}, black},
		{space.Coord{-0.5, -0.5, -0.5}, black},

		// Side Edges
		{space.Coord{-0.5, -0.5, +0.5}, black},
		{space.Coord{-0.5, -0.5, -0.5}, black},

		{space.Coord{-0.5, +0.5, +0.5}, black},
		{space.Coord{-0.5, +0.5, -0.5}, black},

		{space.Coord{+0.5, +0.5, +0.5}, black},
		{space.Coord{+0.5, +0.5, -0.5}, black},

		{space.Coord{+0.5, -0.5, +0.5}, black},
		{space.Coord{+0.5, -0.5, -0.5}, black},
	}
}

////////////////////////////////////////////////////////////////////////////////
