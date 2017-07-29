// Copyright (c) 2013-2016 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package main

//------------------------------------------------------------------------------

import (
	"github.com/drakmaniso/glam/math32"
	"github.com/drakmaniso/glam/mouse"
	"github.com/drakmaniso/glam/pixel"
	"github.com/drakmaniso/glam/plane"
	"github.com/drakmaniso/glam/space"
	"github.com/drakmaniso/glam/window"
)

//------------------------------------------------------------------------------

func (loop) WindowResized(is pixel.Coord) {
	w := plane.CoordOf(is)
	r := w.X / w.Y
	screenFromView = space.Perspective(math32.Pi/4, r, 0.001, 1000.0)
}

//------------------------------------------------------------------------------

func (loop) MouseButtonDown(b mouse.Button, _ int) {
	mouse.SetRelativeMode(true)
}

func (loop) MouseButtonUp(b mouse.Button, _ int) {
	mouse.SetRelativeMode(false)
}

func (loop) MouseMotion(motion pixel.Coord, _ pixel.Coord) {
	m := plane.CoordOf(motion)
	s := plane.CoordOf(window.Size())

	switch {
	case mouse.IsPressed(mouse.Left):
		yaw += 4 * m.X / s.X
		pitch += 4 * m.Y / s.Y
		switch {
		case pitch < -math32.Pi/2:
			pitch = -math32.Pi / 2
		case pitch > +math32.Pi/2:
			pitch = +math32.Pi / 2
		}
		computeWorldFromObject()

	case mouse.IsPressed(mouse.Middle):
		d := m.Times(2).SlashCW(s)
		position.X += d.X
		position.Z += d.Y
		computeWorldFromObject()

	case mouse.IsPressed(mouse.Right):
		d := m.Times(2).SlashCW(s)
		position.X += d.X
		position.Y -= d.Y
		computeWorldFromObject()
	}
}

//------------------------------------------------------------------------------
