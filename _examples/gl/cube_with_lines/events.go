// Copyright (c) 2013-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package main

////////////////////////////////////////////////////////////////////////////////

import (
	"github.com/drakmaniso/cozely"
	"github.com/drakmaniso/cozely/x/gl"
	"github.com/drakmaniso/cozely/x/math32"
	"github.com/drakmaniso/cozely/mouse"
	"github.com/drakmaniso/cozely/plane"
	"github.com/drakmaniso/cozely/space"
)

////////////////////////////////////////////////////////////////////////////////

func (loop) WindowResized(w, h int32) {
	gl.Viewport(0, 0, w, h)
	r := float32(w) / float32(h)
	screenFromView = space.Perspective(math32.Pi/4, r, 0.001, 1000.0)
}

////////////////////////////////////////////////////////////////////////////////

func (loop) MouseButtonDown(b mouse.Button, _ int) {
	mouse.SetRelativeMode(true)
}

func (loop) MouseButtonUp(b mouse.Button, _ int) {
	mouse.SetRelativeMode(false)
}

func (loop) MouseMotion(dx, dy int32, _, _ int32) {
	m := plane.Coord{float32(dx), float32(dy)}
	w, h := cozely.WindowSize()
	s := plane.Coord{float32(w), float32(h)}

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
		d := m.Times(2).Slashcw(s)
		position.X += d.X
		position.Z += d.Y
		computeWorldFromObject()

	case mouse.IsPressed(mouse.Right):
		d := m.Times(2).Slashcw(s)
		position.X += d.X
		position.Y -= d.Y
		computeWorldFromObject()
	}
}

////////////////////////////////////////////////////////////////////////////////
