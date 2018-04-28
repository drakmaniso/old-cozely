// Copyright (c) 2018-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package input

import (
	"github.com/cozely/cozely/coord"
	"github.com/cozely/cozely/internal"
)

////////////////////////////////////////////////////////////////////////////////

// Mouse is a special, unique action automatically created by the framework,
// and linked to the position of the mouse cursor.
var Mouse mouse

type mouse struct {
	active   bool
	hidden   bool
	position coord.CR
	delta    coord.CR
	wheel    coord.CR
}

func (mouse) Active(DeviceID) bool {
	return Mouse.active
}

func (mouse) activate(_ DeviceID, s source) {
	Mouse.active = true
}

func (mouse) deactivate(DeviceID) {
	Mouse.active = false
}

// The cursor is unique action shared by several devices, it needs a special
// update.
func (mouse) updateMouse() {
	Mouse.delta = coord.CR{
		C: internal.MouseDeltaX,
		R: internal.MouseDeltaY,
	}
	internal.MouseDeltaX = 0
	internal.MouseDeltaY = 0

	if Mouse.wheel.C > 0 {
		Mouse.wheel.C--
	} else if Mouse.wheel.C < 0 {
		Mouse.wheel.C++
	}
	if Mouse.wheel.R > 0 {
		Mouse.wheel.R--
	} else if Mouse.wheel.R < 0 {
		Mouse.wheel.R++
	}
	// Wheel delta multiplied by 2 to generate on/off events
	Mouse.wheel = Mouse.wheel.Plus(
		coord.CR{internal.MouseWheelX, internal.MouseWheelY}.
			Times(2))
	internal.MouseWheelX = 0
	internal.MouseWheelY = 0
}

func (mouse) newframe(DeviceID) {
}

func (mouse) update(DeviceID) {
	Mouse.position = coord.CR{
		C: internal.MousePositionX,
		R: internal.MousePositionY,
	}
}

func (mouse) CR() coord.CR {
	if Mouse.active {
		return Mouse.position
	}
	return coord.CR{}
}

// Hide puts the mouse in relative mode: the cursor is hidden and cannot leave
// the game window, but the mouse movements (delta) are continuously reported,
// without constraints.
func (mouse) Hide() {
	Mouse.hidden = true
	_ = internal.MouseSetRelative(true)
}

// Show gets the mouse out of relative mode: the cursor is shown and  free to
// leave the game window, but mouse movements are only reported when the cursor
// is inside the window.
func (mouse) Show() {
	Mouse.hidden = false
	_ = internal.MouseSetRelative(false)
}

// Visible returns true if the relative mode is enabled.
func (mouse) Visible() bool {
	return !Mouse.hidden
}
