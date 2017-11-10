// Copyright (c) 2013-2016 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package internal

//------------------------------------------------------------------------------

/*
#include "sdl.h"
*/
import "C"

import (
	"errors"
)

//------------------------------------------------------------------------------

// SDLQuit is called when the game loop stops.
func SDLQuit() {
	C.SDL_Quit()
}

//------------------------------------------------------------------------------

// GetMilliseconds returns the number of milliseconds elapsed since program
// start.
func GetMilliseconds() uint32 {
	return uint32(C.SDL_GetTicks())
}

//------------------------------------------------------------------------------

// GetSeconds returns the number of seconds elapsed since program start. This
// functions use the performance counter, so is more precise than
// GetMilliseconds.
func GetSeconds() float64 {
	return perfOffset + float64(C.SDL_GetPerformanceCounter())*perfUnit
}

func init() {
	perfUnit = 1.0 / float64(C.SDL_GetPerformanceFrequency())
	ms := C.SDL_GetTicks()
	s := C.SDL_GetPerformanceCounter()
	perfOffset = float64(ms)/1000.0 - float64(s)*perfUnit
}

var perfUnit, perfOffset float64

//------------------------------------------------------------------------------

// GetSDLError returns nil or the current SDL Error wrapped in a Go error.
func GetSDLError() error {
	if s := C.SDL_GetError(); s != nil {
		return errors.New(C.GoString(s))
	}
	return nil
}

//------------------------------------------------------------------------------
