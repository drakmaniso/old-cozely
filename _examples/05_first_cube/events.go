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
	"github.com/drakmaniso/glam/window"
)

//------------------------------------------------------------------------------

type handler struct {
	basic.WindowHandler
	basic.MouseHandler
}

//------------------------------------------------------------------------------

func (h handler) WindowResized(is pixel.Coord, _ uint32) {
	w := plane.CoordOf(is)
	r := w.X / w.Y
	projection = space.Perspective(math32.Pi/4, r, 0.001, 1000.0)
	printState()
}

//------------------------------------------------------------------------------

func (h handler) MouseWheel(motion pixel.Coord, _ uint32) {
	distance -= float32(motion.Y) / 4
	updateView()
	printState()
}

func (h handler) MouseButtonDown(b mouse.Button, _ int, _ uint32) {
	mouse.SetRelativeMode(true)
}

func (h handler) MouseButtonUp(b mouse.Button, _ int, _ uint32) {
	mouse.SetRelativeMode(false)
}

func (h handler) MouseMotion(motion pixel.Coord, _ pixel.Coord, _ uint32) {
	m := plane.CoordOf(motion)
	s := plane.CoordOf(window.Size())

	switch {
	case mouse.IsPressed(mouse.Left):
		d := m.Times(2).SlashCW(s)
		position.X += d.X
		position.Y -= d.Y
		updateModel()
		printState()

	case mouse.IsPressed(mouse.Right):
		yaw += 4 * m.X / s.X
		pitch += 4 * m.Y / s.Y
		switch {
		case pitch < -math32.Pi/2:
			pitch = -math32.Pi / 2
		case pitch > +math32.Pi/2:
			pitch = +math32.Pi / 2
		}
		updateModel()
		printState()
	}
}

//------------------------------------------------------------------------------

func printState() {
	mtx.Print(1, 0, "position=%+6.2f,%+6.2f\v", position.X, position.Y)
	mtx.Print(1, 1, "     yaw=%+6.2f\v", yaw)
	mtx.Print(1, 2, "   pitch=%+6.2f\v", pitch)
	mtx.Print(1, 3, "distance=% 6.2f\v", distance)
}

//------------------------------------------------------------------------------
