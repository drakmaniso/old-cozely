// Copyright (c) 2018-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package input

import (
	"github.com/cozely/cozely/internal"
)

////////////////////////////////////////////////////////////////////////////////

type gpTrigger struct {
	gamepad *internal.Gamepad
	axis    internal.GamepadAxis
	target  Action
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
	return false, false
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
	return j, 2*float32(v) / float32(0x7FFF) - 1
}
