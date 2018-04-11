// Copyright (c) 2013-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package internal

////////////////////////////////////////////////////////////////////////////////

/*
#include "sdl.h"
*/
import "C"

////////////////////////////////////////////////////////////////////////////////

// GetSeconds returns the number of seconds elapsed since program start.
//
// Note: This functions use the performance counter.
func GetSeconds() float64 {
	return float64(C.SDL_GetPerformanceCounter()) * perfUnit
}

func init() {
	perfUnit = 1.0 / float64(C.SDL_GetPerformanceFrequency())
	// ms := C.SDL_GetTicks()
	// s := C.SDL_GetPerformanceCounter()
	// perfOffset = float64(ms)/1000.0 - float64(s)*perfUnit
}

var perfUnit float64

////////////////////////////////////////////////////////////////////////////////
