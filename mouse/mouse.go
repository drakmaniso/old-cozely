// Copyright (c) 2013-2016 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package mouse

//------------------------------------------------------------------------------

// #cgo windows LDFLAGS: -lSDL2
// #cgo linux freebsd darwin pkg-config: sdl2
// #include "../internal/sdl.h"
import "C"

import (
	"fmt"

	"github.com/drakmaniso/glam/internal"
	"github.com/drakmaniso/glam/pixel"
)

//------------------------------------------------------------------------------

// Handler receives the mouse events.
type Handler interface {
	MouseMotion(motion pixel.Coord, position pixel.Coord, timestamp uint32)
	MouseButtonDown(b Button, clicks int, timestamp uint32)
	MouseButtonUp(b Button, clicks int, timestamp uint32)
	MouseWheel(motion pixel.Coord, timestamp uint32)
}

// Handle is the current handlers for mouse events
//
// It can be changed while the loop is running, but must never be nil.
var Handle Handler

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
		return fmt.Errorf("impossible to set relative mouse mode: %s", internal.GetSDLError())
	}
	C.SDL_ShowCursor(C.SDL_ENABLE)
	return nil
}

// GetRelativeMode returns true if the relative mode is enabled.
func GetRelativeMode() bool {
	return C.SDL_GetRelativeMouseMode() == C.SDL_TRUE
}

//------------------------------------------------------------------------------
