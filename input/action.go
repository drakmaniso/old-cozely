// Copyright (c) 2018-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package input

////////////////////////////////////////////////////////////////////////////////

// An Action represents something the player can do in the game. Actions can be
// bound to hardware input (by the player). During the game loop, actions can be
// queried and reacted upon.
type Action interface {
	newframe(d DeviceID)
	update(d DeviceID)
	activate(d DeviceID, b source)
	deactivate(d DeviceID)
}

var actions = struct {
	name map[string]Action
	// For fast iteration, the same list in a slice:
	list []Action
}{
	name: map[string]Action{
		// "Select":  Select,
		// "Back":    Close,
		// "Up":      Up,
		// "Down":    Down,
		// "Left":    Left,
		// "Right":   Right,
		// "Pointer": Pointer,
		// "Click":   Click,
	},
	list: []Action{
		Select,
		Close,
		Up,
		Down,
		Left,
		Right,
		Pointer,
		Click,
	},
}

// Default actions with automatic bindings. If a context contain one of these,
// but no bindings is found, default bindings will be added. If there is no
// declared context, a default context will be created and include all these
// actions.
const (
	Select = ButtonID(iota)
	Close
	Up
	Down
	Left
	Right
	Click

	Pointer = CursorID(0)
)

const maxID = 0xFFFFFFFF
