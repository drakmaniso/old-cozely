// Copyright (c) 2018-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package input

////////////////////////////////////////////////////////////////////////////////

// An Action represents something the player can do in the game. Actions can be
// bound to hardware input (by the player). During the game loop, actions can be
// queried and reacted upon.
type Action interface {
	ActiveOn(d DeviceID) bool
	deactivate(d DeviceID)
	activate(d DeviceID, b source)
	newframe(d DeviceID)
	update(d DeviceID)
}

var actions = struct {
	name map[string]Action
	// For fast iteration, the same list in a slice:
	list []Action
}{
	name: map[string]Action{
		"Menu Select": Select,
		"Menu Back":   Back,
		"Menu Up":     Up,
		"Menu Down":   Down,
		"Menu Left":   Left,
		"Menu Right":  Right,
		"Menu Pointer": Pointer,
		"Menu Click":  Click,
	},
	list: []Action{
		Select,
		Back,
		Up,
		Down,
		Left,
		Right,
		Pointer,
		Click,
	},
}

const maxID = 0xFFFFFFFF
