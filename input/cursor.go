// Copyright (c) 2018-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package input

import (
	"errors"

	"github.com/cozely/cozely/coord"
	"github.com/cozely/cozely/internal"
)

////////////////////////////////////////////////////////////////////////////////

// CursorID identifes an absolute two-dimensional analog input, i.e. any action
// that is best represented by a pair of X and Y coordinates, and whose most
// important characteristic is the position in the game window.
type CursorID uint32

const noCursor = CursorID(maxID)

var cursors struct {
	// For each cursor
	name []string
}

type cursor struct {
	active   bool
	value    coord.XY
	previous coord.XY
}

// Cursor declares a new cursor action, and returns its ID.
func Cursor(name string) CursorID {
	if internal.Running {
		setErr(errors.New("input cursor declaration: declarations must happen before starting the framework"))
		return noCursor
	}

	_, ok := actions.name[name]
	if ok {
		setErr(errors.New("input cursor declaration: name already taken by another action"))
		return noCursor
	}

	a := len(cursors.name)
	if a >= maxID {
		setErr(errors.New("input cursor declaration: too many cursor actions"))
		return noCursor
	}

	actions.name[name] = CursorID(a)
	actions.list = append(actions.list, CursorID(a))
	cursors.name = append(cursors.name, name)

	return CursorID(a)
}

////////////////////////////////////////////////////////////////////////////////

// Name of the action.
func (a CursorID) Name() string {
	return cursors.name[a]
}

// Active returns true if the action is currently active on a specific device
// (i.e. if it is listed in the context currently active on the device).
func (a CursorID) Active(d DeviceID) bool {
	return devices.cursors[d][a].active
}

// XY returns the current status of the action on a specific device. The
// cursorinates are the current absolute position; the values of X and Y are
// normalized between -1 and 1.
func (a CursorID) XY(d DeviceID) coord.XY {
	return devices.cursors[d][a].value
}

////////////////////////////////////////////////////////////////////////////////

func (a CursorID) activate(d DeviceID, b source) {
	devices.cursors[d][a].active = true
	devices.cursorbinds[d][a] = append(devices.cursorbinds[d][a], b)
}

func (a CursorID) newframe(d DeviceID) {
	devices.cursors[d][a].previous = devices.cursors[d][a].value
}

func (a CursorID) update(d DeviceID) {
	if d == kbmouse && mouse.moved {
		v := coord.XY{
			float32(internal.MousePositionX),
			float32(internal.MousePositionY),
		}
		devices.cursors[d][a].value = v
		devices.cursors[0][a].value = v //TODO
		return
	}
	var v coord.XY
	for _, b := range devices.cursorbinds[d][a] {
		v = v.Plus(b.asDelta())
	}
	if v.X != 0 || v.Y != 0 {
		mouse.moved = false
	}
	if !mouse.moved {
		s := coord.XY{float32(internal.Window.Width), float32(internal.Window.Height)}
		v = v.Times(s.Y / 128)

		devices.cursors[d][a].value = devices.cursors[d][a].value.Plus(v)
		if devices.cursors[d][a].value.X < 0 {
			devices.cursors[d][a].value.X = 0
		} else if devices.cursors[d][a].value.X > s.X-1 {
			devices.cursors[d][a].value.X = s.X - 1
		}
		if devices.cursors[d][a].value.Y < 0 {
			devices.cursors[d][a].value.Y = 0
		} else if devices.cursors[d][a].value.Y > s.Y-1 {
			devices.cursors[d][a].value.Y = s.Y - 1
		}

		devices.cursors[0][a].value = devices.cursors[0][a].value.Plus(v)
		if devices.cursors[0][a].value.X < 0 {
			devices.cursors[0][a].value.X = 0
		} else if devices.cursors[0][a].value.X >= s.X-1 {
			devices.cursors[0][a].value.X = s.X - 1
		}
		if devices.cursors[0][a].value.Y < 0 {
			devices.cursors[0][a].value.Y = 0
		} else if devices.cursors[0][a].value.Y >= s.Y-1 {
			devices.cursors[0][a].value.Y = s.Y - 1
		}
	}
}

func (a CursorID) deactivate(d DeviceID) {
	devices.cursorbinds[d][a] = devices.cursorbinds[d][a][:0]
	devices.cursors[d][a].active = false
}
