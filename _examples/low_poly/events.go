// Copyright (c) 2013-2016 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package main

//------------------------------------------------------------------------------

import (
	"github.com/drakmaniso/glam/basic"
	"github.com/drakmaniso/glam/key"
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
	basic.KeyHandler
}

//------------------------------------------------------------------------------

func (h handler) WindowResized(s pixel.Coord, _ uint32) {
	sx, sy := window.Size().Cartesian()
	r := sx / sy
	projection = space.Perspective(math.Pi/4, r, 0.001, 1000.0)
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
		fallthrough
	case mouse.IsPressed(mouse.Left):
		object.position.X += 2 * mx / sx
		object.position.Y -= 2 * my / sy
		updateModel()

	case mouse.IsPressed(mouse.Extra1):
		fallthrough
	case mouse.IsPressed(mouse.Left):
		object.yaw += 4 * mx / sx
		object.pitch += 4 * my / sy
		switch {
		case object.pitch < -math.Pi/2:
			object.pitch = -math.Pi / 2
		case object.pitch > +math.Pi/2:
			object.pitch = +math.Pi / 2
		}
		updateModel()

	case firstPerson:
		// camera.yaw += 2 * mx / sx
		// camera.pitch += 2 * my / sy
		// switch {
		// case camera.pitch < -math.Pi/2:
		// 	camera.pitch = -math.Pi / 2
		// case camera.pitch > +math.Pi/2:
		// 	camera.pitch = +math.Pi / 2
		// }
		// updateView()
	}
}

var firstPerson bool

//------------------------------------------------------------------------------

func (h handler) KeyDown(l key.Label, p key.Position, t uint32) {
	const s = 2.0
	switch p {
	case key.PositionW:
		camera.velocity.Z = -s
	case key.PositionS:
		camera.velocity.Z = s
	case key.PositionA:
		camera.velocity.X = -s
	case key.PositionD:
		camera.velocity.X = s
	default:
		h.KeyHandler.KeyDown(l, p, t)
	}
}

func (h handler) KeyUp(_ key.Label, p key.Position, _ uint32) {
	const s = 5.0
	switch p {
	case key.PositionW, key.PositionS:
		camera.velocity.Z = 0.0
	case key.PositionA, key.PositionD:
		camera.velocity.X = 0.0
	}
}

//------------------------------------------------------------------------------
