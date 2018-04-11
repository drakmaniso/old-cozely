// Copyright (c) 2013-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package main

////////////////////////////////////////////////////////////////////////////////

import (
	"github.com/cozely/cozely"
	"github.com/cozely/cozely/input"
	"github.com/cozely/cozely/space"
	"github.com/cozely/cozely/x/math32"
)

////////////////////////////////////////////////////////////////////////////////

func (loop) React() error {
	if quit.JustPressed(1) {
		cozely.Stop()
	}

	m := input.Cursor.Delta().Cartesian()
	s := input.Cursor.Position().Cartesian()

	if rotate.JustPressed(1) || move.JustPressed(1) || zoom.JustPressed(1) {
		input.Cursor.Hide()
	}
	if rotate.JustReleased(1) || move.JustReleased(1) || zoom.JustReleased(1) {
		input.Cursor.Show()
	}

	if rotate.Pressed(1) {
		yaw += 4 * m.X / s.X
		pitch += 4 * m.Y / s.Y
		switch {
		case pitch < -math32.Pi/2:
			pitch = -math32.Pi / 2
		case pitch > +math32.Pi/2:
			pitch = +math32.Pi / 2
		}
		computeWorldFromObject()
	}

	if move.Pressed(1) {
		d := m.Times(2).Slashcw(s)
		position.X += d.X
		position.Y -= d.Y
		computeWorldFromObject()
	}

	if zoom.Pressed(1) {
		d := m.Times(2).Slashcw(s)
		position.X += d.X
		position.Z += d.Y
		computeWorldFromObject()
	}

	return nil
}

////////////////////////////////////////////////////////////////////////////////

func computeWorldFromObject() {
	rot := space.EulerZXY(pitch, yaw, 0)
	worldFromObject = space.Translation(position).Times(rot)
}

func computeViewFromWorld() {
	viewFromWorld = space.LookAt(
		space.Coord{0, 0, 3},
		space.Coord{0, 0, 0},
		space.Coord{0, 1, 0},
	)
}

////////////////////////////////////////////////////////////////////////////////
