// Copyright (c) 2018-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package input

import (
	"github.com/cozely/cozely/coord"
	"github.com/cozely/cozely/internal"
)

////////////////////////////////////////////////////////////////////////////////

type gpStick struct {
	target       Action
	gamepad      *internal.Gamepad
	xaxis, yaxis internal.GamepadAxis
	x, y         int16
}

////////////////////////////////////////////////////////////////////////////////

func (a *gpStick) bind(c ContextID, target Action) {
	for j := range joysticks.name {
		if joysticks.isgamepad[j] {
			aa := *a
			aa.target = target
			d := joysticks.device[j]
			aa.gamepad = joysticks.gamepad[j]
			devices.bindings[d][c] =
				append(devices.bindings[d][c], &aa)
		}
	}
}

func (a *gpStick) activate(d DeviceID) {
	a.target.activate(d, a)
}

func (a *gpStick) asBool() (just bool, value bool) {
	return false, false
}

func (a *gpStick) asUnipolar() (just bool, value float32) {
	return false, 0
}

func (a *gpStick) asBipolar() (just bool, value float32) {
	return false, 0
}

func (a *gpStick) asCoord() (just bool, value coord.XY) {
	vx, vy := a.gamepad.Axis(a.xaxis), a.gamepad.Axis(a.yaxis)
	j := (vx != a.x) || (vy != a.y)
	a.x, a.y = vx, vy
	var c coord.XY
	if vx < 0 {
		c.X = float32(vx) / float32(0x8000)
	} else {
		c.X = float32(vx) / float32(0x7FFF)
	}
	if vy < 0 {
		c.Y = float32(vy) / float32(0x8000)
	} else {
		c.Y = float32(vy) / float32(0x7FFF)
	}
	return j, c
}

func (a *gpStick) asDelta() coord.XY {
	vx, vy := a.gamepad.Axis(a.xaxis), a.gamepad.Axis(a.yaxis)
	a.x, a.y = vx, vy
	var c coord.XY
	if vx < 0 {
		c.X = float32(vx) / float32(0x8000)
	} else {
		c.X = float32(vx) / float32(0x7FFF)
	}
	if vy < 0 {
		c.Y = float32(vy) / float32(0x8000)
	} else {
		c.Y = float32(vy) / float32(0x7FFF)
	}
	return c
}
