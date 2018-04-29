// Copyright (c) 2013-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package internal

////////////////////////////////////////////////////////////////////////////////

/*
#include "sdl.h"
*/
import "C"

////////////////////////////////////////////////////////////////////////////////

// MouseDelta holds the delta from last mouse position.
var MouseDeltaX, MouseDeltaY int16

// MousePosition holds the current mouse position.
var MousePositionX, MousePositionY int16

// MouseButtons holds the state of the mouse buttons.
var MouseButtons uint32

var MouseWheelX, MouseWheelY int16

////////////////////////////////////////////////////////////////////////////////

// MouseSetRelative enables or disables the relative mode, where the mouse is
// hidden and mouse motions are continuously reported.
func MouseSetRelative(enabled bool) error {
	var m C.SDL_bool
	if enabled {
		m = 1
	}
	if C.SDL_SetRelativeMouseMode(m) != 0 {
		return Wrap("setting relative mouse mode", GetSDLError())
	}
	return nil
}

// MouseRelative returns true if the relative mode is enabled.
func MouseRelative() bool {
	return C.SDL_GetRelativeMouseMode() == C.SDL_TRUE
}

// MouseShow shows or hides the (system) mouse cursor
func MouseShow(show bool) {
	if show {
		C.SDL_ShowCursor(C.SDL_ENABLE)
	} else {
		C.SDL_ShowCursor(C.SDL_DISABLE)
	}
}

// MouseWarp moves the (system) mouse cursor in the window
func MouseWarp(x, y int16) {
	C.SDL_WarpMouseInWindow(Window.window, C.int(x), C.int(y))
}

////////////////////////////////////////////////////////////////////////////////
