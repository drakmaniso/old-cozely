// Copyright (c) 2018-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package input

import (
	"github.com/cozely/cozely/coord"
)

////////////////////////////////////////////////////////////////////////////////

type msWheel struct {
	target                Action
	direction             mouseWheel
	up, down, left, right bool
}

type mouseWheel byte

const (
	mouseScrollUp mouseWheel = iota
	mouseScrollDown
	mouseScrollLeft
	mouseScrollRight
)

////////////////////////////////////////////////////////////////////////////////

func (a *msWheel) bind(c ContextID, target Action) {
	aa := *a
	aa.target = target
	devices.bindings[kbmouse][c] =
		append(devices.bindings[kbmouse][c], &aa)
}

func (a *msWheel) activate(d DeviceID) {
	a.target.activate(d, a)
}

func (a *msWheel) asBool() (just bool, value bool) {
	var j, v bool
	switch a.direction {
	case mouseScrollUp:
		v = mouse.wheel.R > 0 && mouse.wheel.R%2 == 0
		j = (v != a.up)
		a.up = v
	case mouseScrollDown:
		v = mouse.wheel.R < 0 && mouse.wheel.R%2 == 0
		j = (v != a.down)
		a.down = v
	case mouseScrollLeft:
		v = mouse.wheel.C < 0 && mouse.wheel.C%2 == 0
		j = (v != a.left)
		a.left = v
	case mouseScrollRight:
		v = mouse.wheel.C > 0 && mouse.wheel.C%2 == 0
		j = (v != a.right)
		a.right = v
	}
	return j, v
}

func (a *msWheel) asUnipolar() (just bool, value float32) {
	var j, v bool
	switch a.direction {
	case mouseScrollUp:
		v = mouse.wheel.R > 0 && mouse.wheel.R%2 == 0
		j = (v != a.up)
		a.up = v
	case mouseScrollDown:
		v = mouse.wheel.R < 0 && mouse.wheel.R%2 == 0
		j = (v != a.down)
		a.down = v
	case mouseScrollLeft:
		v = mouse.wheel.C < 0 && mouse.wheel.C%2 == 0
		j = (v != a.left)
		a.left = v
	case mouseScrollRight:
		v = mouse.wheel.C > 0 && mouse.wheel.C%2 == 0
		j = (v != a.right)
		a.right = v
	}
	if v {
		return j, 1
	}
	return j, 0
}

func (a *msWheel) asBipolar() (just bool, value float32) {
	var j, v bool
	switch a.direction {
	case mouseScrollUp:
		v = mouse.wheel.R > 0 && mouse.wheel.R%2 == 0
		j = (v != a.up)
		a.up = v
	case mouseScrollDown:
		v = mouse.wheel.R < 0 && mouse.wheel.R%2 == 0
		j = (v != a.down)
		a.down = v
	case mouseScrollLeft:
		v = mouse.wheel.C < 0 && mouse.wheel.C%2 == 0
		j = (v != a.left)
		a.left = v
	case mouseScrollRight:
		v = mouse.wheel.C > 0 && mouse.wheel.C%2 == 0
		j = (v != a.right)
		a.right = v
	}
	if v {
		return j, 1
	}
	return j, 0
}

func (a *msWheel) asCoord() (just bool, value coord.XY) {
	var j, v bool
	switch a.direction {
	case mouseScrollUp:
		v = mouse.wheel.R > 0 && mouse.wheel.R%2 == 0
		j = (v != a.up)
		a.up = v
	case mouseScrollDown:
		v = mouse.wheel.R < 0 && mouse.wheel.R%2 == 0
		j = (v != a.down)
		a.down = v
	case mouseScrollLeft:
		v = mouse.wheel.C < 0 && mouse.wheel.C%2 == 0
		j = (v != a.left)
		a.left = v
	case mouseScrollRight:
		v = mouse.wheel.C > 0 && mouse.wheel.C%2 == 0
		j = (v != a.right)
		a.right = v
	}
	if v {
		return j, coord.XY{1, 0}
	}
	return j, coord.XY{0, 0}
}

func (a *msWheel) asDelta() coord.XY {
	var v bool
	switch a.direction {
	case mouseScrollUp:
		v = mouse.wheel.R > 0 && mouse.wheel.R%2 == 0
		a.up = v
	case mouseScrollDown:
		v = mouse.wheel.R < 0 && mouse.wheel.R%2 == 0
		a.down = v
	case mouseScrollLeft:
		v = mouse.wheel.C < 0 && mouse.wheel.C%2 == 0
		a.left = v
	case mouseScrollRight:
		v = mouse.wheel.C > 0 && mouse.wheel.C%2 == 0
		a.right = v
	}
	if v {
		return coord.XY{1, 0}
	}
	return coord.XY{0, 0}
}
