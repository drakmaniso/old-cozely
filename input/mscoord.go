// Copyright (c) 2018-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package input

import (
	"github.com/cozely/cozely/coord"
	"github.com/cozely/cozely/internal"
)

////////////////////////////////////////////////////////////////////////////////

type msCoord struct {
	target Action
}

////////////////////////////////////////////////////////////////////////////////

func (a *msCoord) bind(c ContextID, target Action) {
	aa := *a
	aa.target = target
	devices.bindings[kbmouse][c] =
		append(devices.bindings[kbmouse][c], &aa)
}

func (a *msCoord) activate(d DeviceID) {
	a.target.activate(d, a)
}

func (a *msCoord) asBool() (just bool, value bool) {
	return false, false
}

func (a *msCoord) asUnipolar() (just bool, value float32) {
	return false, 0
}

func (a *msCoord) asBipolar() (just bool, value float32) {
	return false, 0
}

func (a *msCoord) asCoord() (just bool, value coord.XY) {
	j := Mouse.delta.C != 0 || Mouse.delta.R != 0
	c := coord.XY{
		X: float32(internal.MousePositionX) / float32(internal.Window.Width),
		Y: float32(internal.MousePositionY) / float32(internal.Window.Height),
	}
	return j, c
}

func (a *msCoord) asDelta() coord.XY {
	return Mouse.delta.XY()
}
