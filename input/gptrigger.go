// Copyright (c) 2018-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package input

import (
	"github.com/cozely/cozely/coord"
	"github.com/cozely/cozely/internal"
)

////////////////////////////////////////////////////////////////////////////////

type gpTrigger struct {
	target  Action
	gamepad *internal.Gamepad
	axis    internal.GamepadAxis
	value   int16
}

////////////////////////////////////////////////////////////////////////////////

func (a *gpTrigger) bind(c ContextID, target Action) {
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

func (a *gpTrigger) activate(d DeviceID) {
	a.target.activate(d, a)
}

func (a *gpTrigger) asBool() (just bool, value bool) {
	return false, false //TODO:
}

func (a *gpTrigger) asUnipolar() (just bool, value float32) {
	v := a.gamepad.Axis(a.axis)
	j := v != a.value
	a.value = v
	return j, float32(v) / float32(0x7FFF)
}

func (a *gpTrigger) asBipolar() (just bool, value float32) {
	v := a.gamepad.Axis(a.axis)
	j := v != a.value
	a.value = v
	return j, float32(v) / float32(0x7FFF)
	// return j, 2*float32(v) / float32(0x7FFF) - 1
}

func (a *gpTrigger) asCoord() (just bool, value coord.XY) {
	v := a.gamepad.Axis(a.axis)
	j := v != a.value
	a.value = v
	return j, coord.XY{float32(v) / float32(0x7FFF), 0}
}
