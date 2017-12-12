// Copyright (c) 2013-2017 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package mouse

//------------------------------------------------------------------------------

import (
	"github.com/drakmaniso/carol/internal"
)

//------------------------------------------------------------------------------

// Position returns the current mouse position, relative to the game window.
// Updated at the start of each game loop iteration.
func Position() (x, y int32) {
	return internal.MousePositionX, internal.MousePositionY
}

// Delta returns the mouse position relative to the last call of Delta.
func Delta() (dx, dy int32) {
	dx, dy = internal.MouseDeltaX, internal.MouseDeltaY
	internal.MouseDeltaX, internal.MouseDeltaY = 0, 0
	return dx, dy
}

// SetRelativeMode enables or disables the relative mode, where the mouse is
// hidden and mouse motions are continuously reported.
func SetRelativeMode(enabled bool) error {
	return internal.MouseSetRelativeMode(enabled)
}

// GetRelativeMode returns true if the relative mode is enabled.
func GetRelativeMode() bool {
	return internal.MouseGetRelativeMode()
}

//------------------------------------------------------------------------------
