// Copyright (c) 2018-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package input

import (
	"errors"

	"github.com/cozely/cozely/internal"
)

////////////////////////////////////////////////////////////////////////////////

var joysticks = struct {
	name      []string
	device    []DeviceID
	sdlID     []internal.JoystickID
	isgamepad []bool
	gamepad   []*internal.Gamepad
	joystick  []*internal.Joystick
}{}

func clearJoysticks() {
	joysticks.name = joysticks.name[:0]
	joysticks.sdlID = joysticks.sdlID[:0]
	joysticks.device = joysticks.device[:0]
	joysticks.isgamepad = joysticks.isgamepad[:0]
	joysticks.gamepad = joysticks.gamepad[:0]
	joysticks.joystick = joysticks.joystick[:0]
}

func newJoystick() {
	joysticks.sdlID = append(joysticks.sdlID, internal.JoystickID(-1))
	joysticks.device = append(joysticks.device, noDevice)
	joysticks.name = append(joysticks.name, "")
	joysticks.isgamepad = append(joysticks.isgamepad, false)
	joysticks.gamepad = append(joysticks.gamepad, nil)
	joysticks.joystick = append(joysticks.joystick, nil)
}

func scanJoysticks() {
	n := internal.NumJoysticks()
	internal.Debug.Printf("Detected %d controllers:", n)
	clearJoysticks()
	for j := 0; j < n; j++ {
		newJoystick()
		joysticks.name[j] = internal.JoystickNameForIndex(j)
		joysticks.device[j] = addDevice(joysticks.name[j])
		if internal.IsGameController(j) {
			c := internal.GameControllerOpen(j)
			if c == nil {
				setErr(errors.New("unable to open joystick as gamepad"))
				continue
			}
			joysticks.isgamepad[j] = true
			joysticks.gamepad[j] = c
			joysticks.joystick[j] = c.Joystick()
			joysticks.sdlID[j] = joysticks.joystick[j].InstanceID()
			internal.Debug.Printf("Controller %d is a gamepad (%s) (%d)", j, joysticks.name[j], joysticks.device[j])
		} else {
			internal.Debug.Printf("Controller %d is a joystick (%s)", j, joysticks.name[j])
		}
	}
}
