// Copyright (c) 2018-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package input

import (
	"github.com/cozely/cozely/coord"
	"github.com/cozely/cozely/internal"
)

////////////////////////////////////////////////////////////////////////////////

type msButton struct {
	target  Action
	button  mouseButton
	pressed bool
}

// mouseButton identifies a mouse button
type mouseButton uint32

// mouseButton constants
const (
	mouseLeft mouseButton = 1 << iota
	mouseMiddle
	mouseRight
	mouseBack
	mouseForward
	mouse6
	mouse7
	mouse8
)

////////////////////////////////////////////////////////////////////////////////

func (a *msButton) bind(c ContextID, target Action) {
	aa := *a
	aa.target = target
	devices.bindings[KeyboardAndMouse][c] =
		append(devices.bindings[KeyboardAndMouse][c], &aa)
}

func (a *msButton) activate(d DeviceID) {
	a.target.activate(d, a)
}

func (a *msButton) asButton() (just bool, value bool) {
	v := (mouseButton(internal.MouseButtons) & a.button) != 0
	j := (v != a.pressed)
	a.pressed = v
	return j, a.pressed
}

func (a *msButton) asHalfAxis() (just bool, value float32) {
	v := (mouseButton(internal.MouseButtons) & a.button) != 0
	j := (v != a.pressed)
	a.pressed = v
	if v {
		return j, 1
	}
	return j, 0
}

func (a *msButton) asAxis() (just bool, value float32) {
	v := (mouseButton(internal.MouseButtons) & a.button) != 0
	j := (v != a.pressed)
	a.pressed = v
	if v {
		return j, 1
	}
	return j, 0
}

func (a *msButton) asDualAxis() (just bool, value coord.XY) {
	v := (mouseButton(internal.MouseButtons) & a.button) != 0
	j := (v != a.pressed)
	a.pressed = v
	if v {
		return j, coord.XY{1, 0}
	}
	return j, coord.XY{0, 0}
}

func (a *msButton) asDelta() (just bool, value coord.XY) {
	v := (mouseButton(internal.MouseButtons) & a.button) != 0
	j := (v != a.pressed)
	a.pressed = v
	if v {
		return j, coord.XY{1, 0}
	}
	return j, coord.XY{0, 0}
}
