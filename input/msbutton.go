// Copyright (c) 2018-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package input

import (
	"github.com/cozely/cozely/internal"
)

////////////////////////////////////////////////////////////////////////////////

type msButton struct {
	button        mouseButton
	target        Action
	just, pressed bool
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
	devices.bindings[kbmouse][c] =
		append(devices.bindings[kbmouse][c], &aa)
}

func (a *msButton) activate(d DeviceID) {
	a.target.activate(d, a)
}

func (a *msButton) asBool() (just bool, value bool) {
	v := (mouseButton(internal.MouseButtons) & a.button) != 0
	a.just = (v != a.pressed) //TODO: no need to store?
	a.pressed = v
	return a.just, a.pressed
}
