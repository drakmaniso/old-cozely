// Copyright (c) 2013-2016 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package main

//------------------------------------------------------------------------------

import (
	"github.com/drakmaniso/glam/basic"
	"github.com/drakmaniso/glam/key"
	"github.com/drakmaniso/glam/math32"
	"github.com/drakmaniso/glam/mouse"
	"github.com/drakmaniso/glam/pixel"
	"github.com/drakmaniso/glam/window"
)

//------------------------------------------------------------------------------

type handler struct {
	basic.WindowHandler
	basic.MouseHandler
	basic.KeyHandler
}

//------------------------------------------------------------------------------

func (h handler) WindowResized(s pixel.Coord, _ uint32) {
	camera.WindowResized()
}

//------------------------------------------------------------------------------

func (h handler) MouseWheel(motion pixel.Coord, _ uint32) {
	if motion.Y >= 0 {
		object.scale *= 2.0 * float32(motion.Y)
	} else {
		object.scale /= 2.0 * float32(-motion.Y)
	}
	updateModel()
}

func (h handler) MouseButtonDown(b mouse.Button, _ int, _ uint32) {
	if b == mouse.Right {
		firstPerson = !firstPerson
		mouse.SetRelativeMode(firstPerson)
		_ = mouse.Delta()
	} else {
		mouse.SetRelativeMode(true)
	}
}

func (h handler) MouseButtonUp(b mouse.Button, _ int, _ uint32) {
	if b != mouse.Right {
		mouse.SetRelativeMode(false)
		firstPerson = false
	}
}

func (h handler) MouseMotion(motion pixel.Coord, _ pixel.Coord, _ uint32) {
	mx, my := motion.Cartesian()
	sx, sy := window.Size().Cartesian()

	switch {
	case mouse.IsPressed(mouse.Middle):
		object.position.X += 2 * mx / sx
		object.position.Y -= 2 * my / sy
		updateModel()

	case mouse.IsPressed(mouse.Extra1):
		object.roll -= 4 * mx / sx
		switch {
		case object.roll > 2*math32.Pi:
			object.roll = 2 * math32.Pi
		case object.roll < -2*math32.Pi:
			object.roll = -2 * math32.Pi
		}
		updateModel()
	case mouse.IsPressed(mouse.Left):
		object.yaw += 4 * mx / sx
		object.pitch += 4 * my / sy
		switch {
		case object.pitch < -2*math32.Pi:
			object.pitch = -2 * math32.Pi
		case object.pitch > +2*math32.Pi:
			object.pitch = +2 * math32.Pi
		}
		updateModel()
	case mouse.IsPressed(mouse.Extra2):
		object.yaw += 4 * mx / sx
		updateModel()
	}
}

var firstPerson bool

//------------------------------------------------------------------------------

func (h handler) KeyDown(l key.Label, p key.Position, t uint32) {
	const s = 2.0
	switch p {
	case key.PositionW:
		forward = -s
	case key.PositionS:
		forward = s
	case key.PositionA:
		lateral = -s
	case key.PositionD:
		lateral = s
	case key.PositionSpace:
		vertical = s
	case key.PositionLShift:
		vertical = -s
	case key.PositionQ:
		rolling = -1.0
	case key.PositionE:
		rolling = 1.0
	default:
		h.KeyHandler.KeyDown(l, p, t)
	}
}

func (h handler) KeyUp(_ key.Label, p key.Position, _ uint32) {
	const s = 5.0
	switch p {
	case key.PositionW, key.PositionS:
		forward = 0.0
	case key.PositionA, key.PositionD:
		lateral = 0.0
	case key.PositionSpace, key.PositionLShift:
		vertical = 0.0
	case key.PositionQ, key.PositionE:
		rolling = 0.0
	}
}

//------------------------------------------------------------------------------
