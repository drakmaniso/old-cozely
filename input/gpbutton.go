// Copyright (c) 2018-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package input

import (
	"github.com/cozely/cozely/coord"
	"github.com/cozely/cozely/internal"
)

////////////////////////////////////////////////////////////////////////////////

type gpButton struct {
	target  Action
	gamepad *internal.Gamepad
	button  internal.GamepadButton
	pressed bool
}

////////////////////////////////////////////////////////////////////////////////

func (a *gpButton) bind(c ContextID, target Action) {
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

func (a *gpButton) activate(d DeviceID) {
	a.target.activate(d, a)
}

func (a *gpButton) asBool() (just bool, value bool) {
	v := a.gamepad.Button(a.button)
	j := (v != a.pressed)
	a.pressed = v
	return j, a.pressed
}

func (a *gpButton) asUnipolar() (just bool, value float32) {
	v := a.gamepad.Button(a.button)
	j := (v != a.pressed)
	a.pressed = v
	if v {
		return j, 1
	}
	return j, 0
}

func (a *gpButton) asBipolar() (just bool, value float32) {
	v := a.gamepad.Button(a.button)
	j := (v != a.pressed)
	a.pressed = v
	if v {
		return j, +1
	}
	return j, 0
}

func (a *gpButton) asCoord() (just bool, value coord.XY) {
	v := a.gamepad.Button(a.button)
	j := (v != a.pressed)
	a.pressed = v
	if v {
		return j, coord.XY{1, 0}
	}
	return j, coord.XY{0, 0}
}

func (a *gpButton) asDelta() coord.XY {
	v := a.gamepad.Button(a.button)
	a.pressed = v
	if v {
		return coord.XY{1, 0}
	}
	return coord.XY{0, 0}
}
