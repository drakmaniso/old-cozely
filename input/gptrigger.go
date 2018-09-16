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

func (a *gpTrigger) asButton() (just bool, value bool) {
	//TODO: implement hair-trigger instead
	v := a.gamepad.Axis(a.axis)
	vv := v > (0x7FFF / 2)
	j := vv != (a.value > (0x7FFF / 2))
	a.value = v
	return j, vv
}

func (a *gpTrigger) asHalfAxis() (just bool, value float32) {
	v := a.gamepad.Axis(a.axis)
	j := v != a.value
	a.value = v
	return j, float32(v) / float32(0x7FFF)
}

func (a *gpTrigger) asAxis() (just bool, value float32) {
	v := a.gamepad.Axis(a.axis)
	j := v != a.value
	a.value = v
	return j, float32(v) / float32(0x7FFF)
	// return j, 2*float32(v) / float32(0x7FFF) - 1
}

func (a *gpTrigger) asDualAxis() (just bool, value coord.XY) {
	v := a.gamepad.Axis(a.axis)
	j := v != a.value
	a.value = v
	return j, coord.XY{float32(v) / float32(0x7FFF), 0}
}

func (a *gpTrigger) asDelta() (just bool, value coord.XY) {
	v := a.gamepad.Axis(a.axis)
	j := v != a.value
	a.value = v
	return j, coord.XY{float32(v) / float32(0x7FFF), 0}
}
