// Copyright (c) 2013-2016 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package main

//------------------------------------------------------------------------------

import (
	"github.com/drakmaniso/glam/basic"
	"github.com/drakmaniso/glam/math"
	"github.com/drakmaniso/glam/mouse"
	"github.com/drakmaniso/glam/pixel"
	"github.com/drakmaniso/glam/space"
	"github.com/drakmaniso/glam/window"
)

//------------------------------------------------------------------------------

type handler struct {
	basic.WindowHandler
	basic.MouseHandler
}

//------------------------------------------------------------------------------

func (h handler) WindowResized(s pixel.Coord, _ uint32) {
	sx, sy := window.Size().Cartesian()
	r := sx / sy
	projection = space.Perspective(math.Pi/4, r, 0.001, 1000.0)
}

//------------------------------------------------------------------------------

func (h handler) MouseWheel(motion pixel.Coord, _ uint32) {
	distance -= float32(motion.Y) / 4
	updateView()
}

func (h handler) MouseButtonDown(b mouse.Button, _ int, _ uint32) {
	mouse.SetRelativeMode(true)
}

func (h handler) MouseButtonUp(b mouse.Button, _ int, _ uint32) {
	mouse.SetRelativeMode(false)
}

func (h handler) MouseMotion(motion pixel.Coord, _ pixel.Coord, _ uint32) {
	mx, my := motion.Cartesian()
	sx, sy := window.Size().Cartesian()

	switch {
	case mouse.IsPressed(mouse.Middle):
		fallthrough
	case mouse.IsPressed(mouse.Left):
		position.X += 2 * mx / sx
		position.Y -= 2 * my / sy
		updateModel()

	case mouse.IsPressed(mouse.Extra1):
		fallthrough
	case mouse.IsPressed(mouse.Right):
		yaw += 4 * mx / sx
		pitch += 4 * my / sy
		switch {
		case pitch < -math.Pi/2:
			pitch = -math.Pi / 2
		case pitch > +math.Pi/2:
			pitch = +math.Pi / 2
		}
		updateModel()
	}
}

//------------------------------------------------------------------------------
