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
		"Menu Select":  MenuSelect,
		"Menu Back":    MenuBack,
		"Menu Up":      MenuUp,
		"Menu Down":    MenuDown,
		"Menu Left":    MenuLeft,
		"Menu Right":   MenuRight,
		"Menu Pointer": MenuPointer,
		"Menu Click":   MenuClick,
	},
	list: []Action{
		MenuSelect,
		MenuBack,
		MenuUp,
		MenuDown,
		MenuLeft,
		MenuRight,
		MenuPointer,
		MenuClick,
	},
}

// Default actions with automatic bindings. If a context contain one of these,
// but no bindings is found, default bindings will be added. If there is no
// declared context, a default context will be created and include all these
// actions.
const (
	MenuSelect = ButtonID(iota)
	MenuBack
	MenuUp
	MenuDown
	MenuLeft
	MenuRight
	MenuClick

	MenuPointer = CursorID(0)
)

const maxID = 0xFFFFFFFF
