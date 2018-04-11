// Copyright (c) 2018-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package input

import (
	"github.com/drakmaniso/glam/internal"
	"github.com/drakmaniso/glam/plane"
)

type cursor struct {}

// Cursor is special action linked to the position of the mouse cursor.
var Cursor cursor

var (
	cursorhidden bool
	cursordelta plane.Pixel
)

func (cursor) Active(Device) bool {
	return true
}

func (cursor) deactivate(Device) {}

func (cursor) activate(Device, binding) {}

func (cursor) newframe(Device) {
	cursordelta = plane.Pixel{
		X: internal.MouseDeltaX,
		Y: internal.MouseDeltaY,
	}
	internal.MouseDeltaX = 0
	internal.MouseDeltaY = 0
}

func (cursor) Position() plane.Pixel {
	return plane.Pixel{
		X: internal.MousePositionX,
		Y: internal.MousePositionY,
	}
}

func (cursor) Delta() plane.Pixel {
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
