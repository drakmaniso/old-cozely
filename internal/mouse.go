// Copyright (c) 2013-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package internal

//------------------------------------------------------------------------------

/*
#include "sdl.h"
*/
import "C"

//------------------------------------------------------------------------------

// A MouseButton on the mouse
type MouseButton uint8 //TODO: remove

//------------------------------------------------------------------------------

// MouseDelta holds the delta from last mouse position.
var MouseDeltaX, MouseDeltaY int32

// MousePosition holds the current mouse position.
var MousePositionX, MousePositionY int32

// MouseButtons holds the state of the mouse buttons.
var MouseButtons uint32

//------------------------------------------------------------------------------

// MouseSetRelativeMode enables or disables the relative mode, where the mouse is
// hidden and mouse motions are continuously reported.
func MouseSetRelativeMode(enabled bool) error {
	var m C.SDL_bool
	if enabled {
		m = 1
		C.SDL_ShowCursor(C.SDL_DISABLE)
	}
	if C.SDL_SetRelativeMouseMode(m) != 0 {
		C.SDL_ShowCursor(C.SDL_ENABLE)
		return Error("setting relative mouse mode", GetSDLError())
	}
	C.SDL_ShowCursor(C.SDL_ENABLE)
	return nil
}

// MouseGetRelativeMode returns true if the relative mode is enabled.
func MouseGetRelativeMode() bool {
	return C.SDL_GetRelativeMouseMode() == C.SDL_TRUE
}

//------------------------------------------------------------------------------
