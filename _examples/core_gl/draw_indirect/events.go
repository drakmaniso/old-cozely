// Copyright (c) 2013-2016 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package main

//------------------------------------------------------------------------------

import (
	"github.com/drakmaniso/carol"
	"github.com/drakmaniso/carol/core/math32"
	"github.com/drakmaniso/carol/mouse"
	"github.com/drakmaniso/carol/pixel"
	"github.com/drakmaniso/carol/plane"
	"github.com/drakmaniso/carol/space"
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
	s := plane.CoordOf(carol.WindowSize())

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
