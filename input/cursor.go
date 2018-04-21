// Copyright (c) 2018-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package input

import (
	"github.com/cozely/cozely/coord"
	"github.com/cozely/cozely/internal"
)

////////////////////////////////////////////////////////////////////////////////

// Cursor is a special, unique action automatically created by the framework,
// and linked to the position of the system mouse cursor.
var Cursor cursor

type cursor struct {
	active bool
	hidden bool
	delta  coord.CR
	wheel  coord.CR
	binds  []source
}

func (cursor) Active(DeviceID) bool {
	return Cursor.active
}

func (cursor) activate(_ DeviceID, s source) {
	Cursor.active = true
	Cursor.binds = append(Cursor.binds, s)
}

func (cursor) deactivate(DeviceID) {
	Cursor.binds = Cursor.binds[:0]
	Cursor.active = false
}

// The cursor is unique action shared by several devices, it needs a special
// update.
func (cursor) specialnewframe() {
	Cursor.delta = coord.CR{
		C: internal.MouseDeltaX,
		R: internal.MouseDeltaY,
	}
	internal.MouseDeltaX = 0
	internal.MouseDeltaY = 0

	if Cursor.wheel.C > 0 {
		Cursor.wheel.C--
	} else if Cursor.wheel.C < 0 {
		Cursor.wheel.C++
	}
	if Cursor.wheel.R > 0 {
		Cursor.wheel.R--
	} else if Cursor.wheel.R < 0 {
		Cursor.wheel.R++
	}
	// Wheel delta multiplied by 2 to generate on/off events
	Cursor.wheel = Cursor.wheel.Plus(
		coord.CR{internal.MouseWheelX, internal.MouseWheelY}.
			Times(2))
	internal.MouseWheelX = 0
	internal.MouseWheelY = 0
}

func (cursor) newframe(DeviceID) {
}

func (cursor) update(DeviceID) {
	p := coord.CR{}
	for _, b := range Cursor.binds {
		v := b.asDelta().CR()
		p = p.Plus(v)
		p = p.Plus(v)
	}
	if p.C != 0 || p.R != 0 {
		//TODO:
	}
}

func (cursor) Position() coord.CR {
	return coord.CR{
		C: internal.MousePositionX,
		R: internal.MousePositionY,
	}
}

func (cursor) Delta() coord.CR {
	return Cursor.delta
}

// Hide puts the mouse in relative mode: the cursor is hidden and cannot leave
// the game window, but the mouse movements (delta) are continuously reported,
// without constraints.
func (cursor) Hide() {
	Cursor.hidden = true
	_ = internal.MouseSetRelative(true)
}

// Show gets the mouse out of relative mode: the cursor is shown and  free to
// leave the game window, but mouse movements are only reported when the cursor
// is inside the window.
func (cursor) Show() {
	Cursor.hidden = false
	_ = internal.MouseSetRelative(false)
}

// Hidden returns true if the relative mode is enabled.
func (cursor) Hidden() bool {
	return internal.MouseRelative()
}
