// Copyright (c) 2013-2016 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package main

//------------------------------------------------------------------------------

import (
	"github.com/drakmaniso/carol/color"
	"github.com/drakmaniso/carol/space"
)

//------------------------------------------------------------------------------

var (
	purple = color.RGB{R: 0.08, G: 0.15, B: 0.15}
	orange = color.RGB{R: 0.15, G: 0.21, B: 0.21}
	green  = color.RGB{R: 0.10, G: 0.17, B: 0.17}
)

//------------------------------------------------------------------------------

func cube() mesh {
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

//------------------------------------------------------------------------------
