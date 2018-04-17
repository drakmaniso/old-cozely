// Copyright (c) 2018-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package input

import (
	coordi "github.com/cozely/cozely/coord"
	"github.com/cozely/cozely/internal"
)

type cursor struct{}

// Cursor is a special action automatically created by the framework, and linked
// to the position of the system mouse cursor.
var Cursor cursor

var (
	cursorhidden bool
	cursordelta  coordi.CR
)

func (cursor) Active(DeviceID) bool {
	return true
}

func (cursor) deactivate(DeviceID) {}

func (cursor) activate(DeviceID, source) {}

func (cursor) newframe(DeviceID) {
	cursordelta = coordi.CR{
		C: internal.MouseDeltaX,
		R: internal.MouseDeltaY,
	}
	internal.MouseDeltaX = 0
	internal.MouseDeltaY = 0
}

func (cursor) Position() coordi.CR {
	return coordi.CR{
		C: internal.MousePositionX,
		R: internal.MousePositionY,
	}
}

func (cursor) Delta() coordi.CR {
	return cursordelta
}

// SetRelative enables or disables the relative mode, where the mouse is hidden
// and mouse motions are continuously reported.
func (cursor) SetRelative(enabled bool) error {
	return internal.MouseSetRelative(enabled)
}

// Hide puts the mouse in relative mode: the cursor is hidden and the mouse
// cannot leave the game window.
func (cursor) Hide() {
	cursorhidden = true
	_ = internal.MouseSetRelative(true)
}

// Show gets the mouse out of relative mode: the cursor is shown and the mouse
// is free to leave the game window.
func (cursor) Show() {
	cursorhidden = false
	_ = internal.MouseSetRelative(false)
}

// Hidden returns true if the relative mode is enabled.
func (cursor) Hidden() bool {
	return internal.MouseRelative()
}
