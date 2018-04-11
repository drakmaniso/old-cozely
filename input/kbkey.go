// Copyright (c) 2018-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package input

import (
	"github.com/drakmaniso/glam/internal"
)

type kbKey struct {
	keycode       keyCode
	target        Action
	just, pressed bool
}

func (a *kbKey) bind(c Context, target Action) {
	aa := *a
	aa.target = target
	devices.bindings[kbmouse][c] =
		append(devices.bindings[kbmouse][c], &aa)
}

func (a *kbKey) device() Device {
	return kbmouse
}

func (a *kbKey) action() Action {
	return a.target
}

func (a *kbKey) activate(d Device) {
	a.target.activate(d, a)
}

func (a *kbKey) asBool() (just bool, value bool) {
	v := internal.Key(a.keycode)
	a.just = (v != a.pressed) //TODO: no need to store?
	a.pressed = v
	return a.just, a.pressed
}
