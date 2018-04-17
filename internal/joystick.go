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
	GamepadButtonInvalid = GamepadButton(C.SDL_CONTROLLER_BUTTON_INVALID)
	GamepadA             = GamepadButton(C.SDL_CONTROLLER_BUTTON_A)
	GamepadB             = GamepadButton(C.SDL_CONTROLLER_BUTTON_B)
	GamepadX             = GamepadButton(C.SDL_CONTROLLER_BUTTON_X)
	GamepadY             = GamepadButton(C.SDL_CONTROLLER_BUTTON_Y)
	GamepadBack          = GamepadButton(C.SDL_CONTROLLER_BUTTON_BACK)
	GamepadGuide         = GamepadButton(C.SDL_CONTROLLER_BUTTON_GUIDE)
	GamepadStart         = GamepadButton(C.SDL_CONTROLLER_BUTTON_START)
	GamepadLeftClick     = GamepadButton(C.SDL_CONTROLLER_BUTTON_LEFTSTICK)
	GamepadRightClick    = GamepadButton(C.SDL_CONTROLLER_BUTTON_RIGHTSTICK)
	GamepadLeftBumper    = GamepadButton(C.SDL_CONTROLLER_BUTTON_LEFTSHOULDER)
	GamepadRightBumper   = GamepadButton(C.SDL_CONTROLLER_BUTTON_RIGHTSHOULDER)
	GamepadDpadUp        = GamepadButton(C.SDL_CONTROLLER_BUTTON_DPAD_UP)
	GamepadDpadDown      = GamepadButton(C.SDL_CONTROLLER_BUTTON_DPAD_DOWN)
	GamepadDpadLeft      = GamepadButton(C.SDL_CONTROLLER_BUTTON_DPAD_LEFT)
	GamepadDpadRight     = GamepadButton(C.SDL_CONTROLLER_BUTTON_DPAD_RIGHT)
	GamepadButtonMax     = GamepadButton(C.SDL_CONTROLLER_BUTTON_MAX)
)

type GamepadAxis = C.SDL_GameControllerAxis

const (
	GamepadInvalidAxis  = GamepadAxis(C.SDL_CONTROLLER_AXIS_INVALID)
	GamepadLeftX        = GamepadAxis(C.SDL_CONTROLLER_AXIS_LEFTX)
	GamepadLeftY        = GamepadAxis(C.SDL_CONTROLLER_AXIS_LEFTY)
	GamepadRightX       = GamepadAxis(C.SDL_CONTROLLER_AXIS_RIGHTX)
	GamepadRightY       = GamepadAxis(C.SDL_CONTROLLER_AXIS_RIGHTY)
	GamepadLeftTrigger  = GamepadAxis(C.SDL_CONTROLLER_AXIS_TRIGGERLEFT)
	GamepadRightTrigger = GamepadAxis(C.SDL_CONTROLLER_AXIS_TRIGGERRIGHT)
	GamepadMaxAxis      = GamepadAxis(C.SDL_CONTROLLER_AXIS_MAX)
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

func (a *Gamepad) Axis(x GamepadAxis) int16 {
	v := C.SDL_GameControllerGetAxis((*C.SDL_GameController)(a),
		C.SDL_GameControllerAxis(x))
	return int16(v)
}

////////////////////////////////////////////////////////////////////////////////
