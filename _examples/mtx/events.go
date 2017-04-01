// Copyright (c) 2013-2016 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package main

//------------------------------------------------------------------------------

import (
	"github.com/drakmaniso/glam/basic"
	"github.com/drakmaniso/glam/math32"
	"github.com/drakmaniso/glam/mouse"
	"github.com/drakmaniso/glam/mtx"
	"github.com/drakmaniso/glam/pixel"
	"github.com/drakmaniso/glam/plane"
	"github.com/drakmaniso/glam/space"
)

//------------------------------------------------------------------------------

type handler struct {
	basic.WindowHandler
	basic.MouseHandler
}

//------------------------------------------------------------------------------

func (h handler) WindowResized(is pixel.Coord, _ uint32) {
	s := plane.CoordOf(is)
	r := s.X / s.Y
	projection = space.Perspective(math32.Pi/4, r, 0.001, 1000.0)

	// MTX

	for y := 0; y < 16; y++ {
		for x := 0; x < 16; x++ {
			mtx.Poke(-16+x, -16+y, byte(x+16*y))
		}
	}
}

//------------------------------------------------------------------------------

func (h handler) MouseWheel(motion pixel.Coord, _ uint32) {
}

func (h handler) MouseButtonDown(b mouse.Button, _ int, _ uint32) {
}

func (h handler) MouseButtonUp(b mouse.Button, _ int, _ uint32) {
}

func (h handler) MouseMotion(motion pixel.Coord, _ pixel.Coord, _ uint32) {
}

//------------------------------------------------------------------------------
