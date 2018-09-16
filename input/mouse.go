// Copyright (c) 2018-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package input

import (
	"github.com/cozely/cozely/coord"
	"github.com/cozely/cozely/internal"
)

////////////////////////////////////////////////////////////////////////////////

var mouse struct {
	grabbed bool
	delta   coord.CR
	moved   bool
	wheel   coord.CR
}

func updateMouse() {
	mouse.delta = coord.CR{
		C: internal.MouseDeltaX,
		R: internal.MouseDeltaY,
	}
	mouse.moved = mouse.delta.C != 0 || mouse.delta.R != 0
	internal.MouseDeltaX = 0
	internal.MouseDeltaY = 0

	if mouse.wheel.C > 0 {
		mouse.wheel.C--
	} else if mouse.wheel.C < 0 {
		mouse.wheel.C++
	}
	if mouse.wheel.R > 0 {
		mouse.wheel.R--
	} else if mouse.wheel.R < 0 {
		mouse.wheel.R++
	}
	// Wheel delta multiplied by 2 to generate on/off events
	mw := coord.CR{internal.MouseWheelX, internal.MouseWheelY}
	mouse.wheel = mouse.wheel.Plus(mw.Times(2))
	internal.MouseWheelX = 0
	internal.MouseWheelY = 0
}

// GrabMouse puts the mouse in relative mode: the cursor is hidden and cannot leave
// the game window, but the mouse movements (delta) are continuously reported,
// without constraints.
func GrabMouse(grab bool) {
	mouse.grabbed = grab
	_ = internal.MouseSetRelative(grab)
}

// MouseGrabbed returns true if the relative mode is enabled.
func MouseGrabbed() bool {
	return mouse.grabbed
}

// ShowMouse shows or hides the (system) mouse cursor.
func ShowMouse(show bool) {
	internal.MouseShow(show)
}
