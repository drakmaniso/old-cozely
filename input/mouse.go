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
	hidden bool
	delta  coord.CR
	moved  bool
	wheel  coord.CR
}

// The cursor is unique action shared by several devices, it needs a special
// update.
func updateMouse() {
	Mouse.delta = coord.CR{
		C: internal.MouseDeltaX,
		R: internal.MouseDeltaY,
	}
	Mouse.moved = Mouse.moved || Mouse.delta.C != 0 || Mouse.delta.R != 0
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

// GrabMouse puts the mouse in relative mode: the cursor is hidden and cannot leave
// the game window, but the mouse movements (delta) are continuously reported,
// without constraints.
func GrabMouse(grab bool) {
	Mouse.hidden = grab
	_ = internal.MouseSetRelative(grab)
}

// MouseGrabbed returns true if the relative mode is enabled.
func MouseGrabbed() bool {
	return Mouse.hidden
}
