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

func (a *msAxis) bind(c ContextID, target Action) {
	aa := *a
	aa.target = target
	devices.bindings[kbmouse][c] =
		append(devices.bindings[kbmouse][c], &aa)
}

func (a *msAxis) activate(d DeviceID) {
	a.target.activate(d, a)
}

func (a *msAxis) asBool() (just bool, value bool) {
	return false, false
}

func (a *msAxis) asUnipolar() (just bool, value float32) {
	return false, 0
}

func (a *msAxis) asBipolar() (just bool, value float32) {
	return false, 0
}

func (a *msAxis) asCoord() (just bool, value coord.XY) {
	return false, coord.XY{0, 0}
}

func (a *msAxis) asDelta() coord.XY {
	return coord.XY{0, 0}
}
