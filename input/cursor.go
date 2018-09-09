// Copyright (c) 2018-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package input

import (
	"errors"

	"github.com/cozely/cozely/internal"
	"github.com/cozely/cozely/window"
)

////////////////////////////////////////////////////////////////////////////////

//TODO: it seems the mouse is always bound to all cursors! :(

// CursorID identifes an absolute two-dimensional analog input, i.e. any action
// that is best represented by a pair of X and Y coordinates, and whose most
// important characteristic is the position in the game window.
type CursorID uint32

const noCursor = CursorID(maxID)

// Pointer is a default action.

var cursors = struct {
	// For each cursor
	name []string
}{
	name: []string{
		"Menu Pointer",
	},
}

type cursor struct {
	value    window.XY
	previous window.XY
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

// XY returns the current status of the action on the current device. The
// cursorinates are the current absolute position; the values of X and Y are
// normalized between -1 and 1.
func (a CursorID) XY() window.XY {
	return a.XYon(Any)
}

// XYon returns the current status of the action on a specific device. The
// cursorinates are the current absolute position; the values of X and Y are
// normalized between -1 and 1.
func (a CursorID) XYon(d DeviceID) window.XY {
	return devices.cursors[d][a].value
}

////////////////////////////////////////////////////////////////////////////////

func (a CursorID) activate(d DeviceID, b source) {
	devices.cursorsbinds[d][a] = append(devices.cursorsbinds[d][a], b)
}

func (a CursorID) newframe(d DeviceID) {
	devices.cursors[d][a].previous = devices.cursors[d][a].value
}

func (a CursorID) update(d DeviceID) {
	if d == kbmouse && mouse.moved {
		v := window.XY{
			internal.MousePositionX,
			internal.MousePositionY,
		}
		devices.cursors[d][a].value = v
		devices.cursors[0][a].value = v //TODO
		return
	}
	var v window.XY
	for _, b := range devices.cursorsbinds[d][a] {
		v = v.Plus(window.XYof(b.asDelta()))
	}
	if v.X != 0 || v.Y != 0 {
		mouse.moved = false
	}
	if !mouse.moved {
		s := window.XY{internal.Window.Width, internal.Window.Height}
		v = v.Times(int16(float32(s.Y) / 128)) //TODO: handle stick->cursor

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
	devices.cursorsbinds[d][a] = devices.cursorsbinds[d][a][:0]
}
