// Copyright (c) 2013-2016 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package internal

//------------------------------------------------------------------------------

/*
#include "sdl.h"

#define PEEP_SIZE 128

SDL_Event Events[PEEP_SIZE];

int PeepEvents();
*/
import "C"
import "unsafe"

//------------------------------------------------------------------------------

// PeepSize is the size of the event buffer.
const PeepSize = C.PEEP_SIZE

//------------------------------------------------------------------------------

// PeepEvents fill the event buffer and returns the number of events fetched.
func PeepEvents() int {
	return int(C.PeepEvents())
}

// EventAt returns a pointer to an event in the event buffer.
func EventAt(i int) unsafe.Pointer {
	return unsafe.Pointer(&C.Events[i])
}

//------------------------------------------------------------------------------
