// Copyright (c) 2018-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package internal

////////////////////////////////////////////////////////////////////////////////

/*
#include "sdl.h"
*/
import "C"

////////////////////////////////////////////////////////////////////////////////

// GameController is a pointer to a SDL object.
type GameController C.SDL_GameController

////////////////////////////////////////////////////////////////////////////////

// NumJoysticks returns the number of attached joysticks on success or a
// negative error code on failure; call SDL_GetError() for more information.
func NumJoysticks() int {
	return int(C.SDL_NumJoysticks())
}

// IsGameController returns true if the given joystick is supported by the game
// controller interface, false if it isn't or it's an invalid index.
func IsGameController(j int) bool {
	return C.SDL_IsGameController(C.int(j)) == C.SDL_TRUE
}

// GameControllerOpen Returns a gamecontroller pointer or nil if an error
// occurred; call SDL_GetError() for more information.
func GameControllerOpen(j int) *GameController {
	c := C.SDL_GameControllerOpen(C.int(j))
	return (*GameController)(c)
}

func (a *GameController) Name() string {
	n := C.SDL_GameControllerName((*C.SDL_GameController)(a))
	return C.GoString(n)
}

func JoystickNameForIndex(j int) string {
	n := C.SDL_JoystickNameForIndex(C.int(j))
	return C.GoString(n)
}

////////////////////////////////////////////////////////////////////////////////
