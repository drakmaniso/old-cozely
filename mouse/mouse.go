// Copyright (c) 2013-2016 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package mouse

//------------------------------------------------------------------------------

// #cgo windows LDFLAGS: -lSDL2
// #cgo linux freebsd darwin pkg-config: sdl2
// #include "../sdl.h"
import "C"

import (
	"github.com/drakmaniso/carol/internal"
	"github.com/drakmaniso/carol/pixel"
)

//------------------------------------------------------------------------------

// Position returns the current mouse position, relative to the game window.
// Updated at the start of each game loop iteration.
func Position() pixel.Coord {
	return internal.MousePosition
}

// Delta returns the mouse position relative to the last call of Delta.
func Delta() pixel.Coord {
	result := internal.MouseDelta
	internal.MouseDelta.X, internal.MouseDelta.Y = 0, 0
	return result
}

// SetRelativeMode enables or disables the relative mode, where the mouse is
// hidden and mouse motions are continuously reported.
func SetRelativeMode(enabled bool) error {
	var m C.SDL_bool
	if enabled {
		m = 1
		C.SDL_ShowCursor(C.SDL_DISABLE)
	}
	if C.SDL_SetRelativeMouseMode(m) != 0 {
		C.SDL_ShowCursor(C.SDL_ENABLE)
		return internal.Error("setting relative mouse mode", internal.GetSDLError())
	}
	C.SDL_ShowCursor(C.SDL_ENABLE)
	return nil
}

// GetRelativeMode returns true if the relative mode is enabled.
func GetRelativeMode() bool {
	return C.SDL_GetRelativeMouseMode() == C.SDL_TRUE
}

//------------------------------------------------------------------------------
