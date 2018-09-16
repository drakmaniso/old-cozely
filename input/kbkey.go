// Copyright (c) 2018-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package input

import (
	"github.com/cozely/cozely/coord"
	"github.com/cozely/cozely/internal"
)

////////////////////////////////////////////////////////////////////////////////

type kbKey struct {
	target  Action
	keycode keyCode
	pressed bool
}

////////////////////////////////////////////////////////////////////////////////

func (a *kbKey) bind(c ContextID, target Action) {
	aa := *a
	aa.target = target
	devices.bindings[KeyboardAndMouse][c] =
		append(devices.bindings[KeyboardAndMouse][c], &aa)
}

func (a *kbKey) device() DeviceID {
	return KeyboardAndMouse
}

func (a *kbKey) action() Action {
	return a.target
}

func (a *kbKey) activate(d DeviceID) {
	a.target.activate(d, a)
}

func (a *kbKey) asButton() (just bool, value bool) {
	v := internal.Key(a.keycode)
	j := (v != a.pressed)
	a.pressed = v
	return j, a.pressed
}

func (a *kbKey) asHalfAxis() (just bool, value float32) {
	v := internal.Key(a.keycode)
	j := (v != a.pressed)
	a.pressed = v
	if v {
		return j, 1
	}
	return j, 0
}

func (a *kbKey) asAxis() (just bool, value float32) {
	v := internal.Key(a.keycode)
	j := (v != a.pressed)
	a.pressed = v
	if v {
		return j, +1
	}
	return j, 0
}

func (a *kbKey) asDualAxis() (just bool, value coord.XY) {
	v := internal.Key(a.keycode)
	j := (v != a.pressed)
	a.pressed = v
	if v {
		return j, coord.XY{1, 0}
	}
	return j, coord.XY{0, 0}
}

func (a *kbKey) asDelta() (just bool, value coord.XY) {
	v := internal.Key(a.keycode)
	j := (v != a.pressed)
	a.pressed = v
	if v {
		return j, coord.XY{1, 0}
	}
	return j, coord.XY{0, 0}
}
