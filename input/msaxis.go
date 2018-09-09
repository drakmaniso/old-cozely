// Copyright (c) 2018-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package input

import (
	"github.com/cozely/cozely/coord"
)

////////////////////////////////////////////////////////////////////////////////

type msAxis struct {
	target Action
}

////////////////////////////////////////////////////////////////////////////////

//TODO: implement!

func (a *msAxis) bind(c ContextID, target Action) {
	aa := *a
	aa.target = target
	devices.bindings[KeyboardAndMouse][c] =
		append(devices.bindings[KeyboardAndMouse][c], &aa)
}

func (a *msAxis) activate(d DeviceID) {
	a.target.activate(d, a)
}

func (a *msAxis) asButton() (just bool, value bool) {
	return false, false
}

func (a *msAxis) asHalfAxis() (just bool, value float32) {
	return false, 0
}

func (a *msAxis) asAxis() (just bool, value float32) {
	return false, 0
}

func (a *msAxis) asDualAxis() (just bool, value coord.XY) {
	return false, coord.XY{0, 0}
}

func (a *msAxis) asDelta() (just bool, value coord.XY) {
	return false, coord.XY{0, 0}
}
