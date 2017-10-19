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
	"github.com/drakmaniso/carol/plane"
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

// SetSmoothing sets the smoothing factor for SmoothDelta.
func SetSmoothing(s float32) {
	smoothing = s
}

// SmoothDelta returns relative to the last call of SmoothDelta (or Delta), but
// smoothed to avoid jitter. The is best used with a fixed timestep (see
// carol.LoopStable).
func SmoothDelta() plane.Coord {
	d := plane.CoordOf(Delta())
	smoothed = smoothed.Plus(d.Minus(smoothed).Times(smoothing))
	return smoothed
}

var smoothed plane.Coord
var smoothing = float32(0.4)

//------------------------------------------------------------------------------
