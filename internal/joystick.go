// Copyright (c) 2018-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package internal

////////////////////////////////////////////////////////////////////////////////

/*
#include "sdl.h"
*/
import "C"

////////////////////////////////////////////////////////////////////////////////

// Joystick is a pointer to a SDL object.
type Joystick C.SDL_Joystick

// JoystickID is the unique identifier given by SDL to a joystick.
type JoystickID C.SDL_JoystickID

// Gamepad is a pointer to a SDL object.
type Gamepad C.SDL_GameController

type GamepadButton C.SDL_GameControllerButton

const (
	GamepadButtonInvalid     = GamepadButton(C.SDL_CONTROLLER_BUTTON_INVALID)
	GamepadButtonA           = GamepadButton(C.SDL_CONTROLLER_BUTTON_A)
	GamepadButtonB           = GamepadButton(C.SDL_CONTROLLER_BUTTON_B)
	GamepadButtonX           = GamepadButton(C.SDL_CONTROLLER_BUTTON_X)
	GamepadButtonY           = GamepadButton(C.SDL_CONTROLLER_BUTTON_Y)
	GamepadButtonBack        = GamepadButton(C.SDL_CONTROLLER_BUTTON_BACK)
	GamepadButtonGuide       = GamepadButton(C.SDL_CONTROLLER_BUTTON_GUIDE)
	GamepadButtonStart       = GamepadButton(C.SDL_CONTROLLER_BUTTON_START)
	GamepadButtonLeftStick   = GamepadButton(C.SDL_CONTROLLER_BUTTON_LEFTSTICK)
	GamepadButtonRightStick  = GamepadButton(C.SDL_CONTROLLER_BUTTON_RIGHTSTICK)
	GamepadButtonLeftBumper  = GamepadButton(C.SDL_CONTROLLER_BUTTON_LEFTSHOULDER)
	GamepadButtonRightBumper = GamepadButton(C.SDL_CONTROLLER_BUTTON_RIGHTSHOULDER)
	GamepadButtonDpadUp      = GamepadButton(C.SDL_CONTROLLER_BUTTON_DPAD_UP)
	GamepadButtonDpadDown    = GamepadButton(C.SDL_CONTROLLER_BUTTON_DPAD_DOWN)
	GamepadButtonDpadLeft    = GamepadButton(C.SDL_CONTROLLER_BUTTON_DPAD_LEFT)
	GamepadButtonDpadRight   = GamepadButton(C.SDL_CONTROLLER_BUTTON_DPAD_RIGHT)
	GamepadButtonMax         = GamepadButton(C.SDL_CONTROLLER_BUTTON_MAX)
)

////////////////////////////////////////////////////////////////////////////////

// NumJoysticks returns the number of attached joysticks on success or a
// negative error code on failure; call SDL_GetError() for more information.
func NumJoysticks() int {
	return int(C.SDL_NumJoysticks())
}

func JoystickNameForIndex(j int) string {
	n := C.SDL_JoystickNameForIndex(C.int(j))
	return C.GoString(n)
}

func (a *Joystick) InstanceID() JoystickID {
	id := C.SDL_JoystickInstanceID((*C.SDL_Joystick)(a))
	return JoystickID(id)
}

////////////////////////////////////////////////////////////////////////////////

// IsGameController returns true if the given joystick is supported by the game
// controller interface, false if it isn't or it's an invalid index.
func IsGameController(j int) bool {
	return C.SDL_IsGameController(C.int(j)) == C.SDL_TRUE
}

// GameControllerOpen Returns a gamecontroller pointer or nil if an error
// occurred; call SDL_GetError() for more information.
func GameControllerOpen(j int) *Gamepad {
	c := C.SDL_GameControllerOpen(C.int(j))
	return (*Gamepad)(c)
}

func (a *Gamepad) Name() string {
	n := C.SDL_GameControllerName((*C.SDL_GameController)(a))
	return C.GoString(n)
}

func (a *Gamepad) Joystick() *Joystick {
	j := C.SDL_GameControllerGetJoystick((*C.SDL_GameController)(a))
	return (*Joystick)(j)
}

func (a *Gamepad) Button(b GamepadButton) bool {
	v := C.SDL_GameControllerGetButton((*C.SDL_GameController)(a),
		C.SDL_GameControllerButton(b))
	return v == 1
}

////////////////////////////////////////////////////////////////////////////////
