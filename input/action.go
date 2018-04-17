// Copyright (c) 2018-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package input

////////////////////////////////////////////////////////////////////////////////

// An Action represents something the player can do in the game. Actions can be
// bound to hardware input (by the player). During the game loop, actions can be
// queried and reacted upon.
type Action interface {
	Active(d DeviceID) bool
	deactivate(d DeviceID)
	activate(d DeviceID, b binding)
	newframe(d DeviceID)
	update(d DeviceID)
}

var actions = struct {
	names map[string]Action
}{
	names: map[string]Action{},
}

const maxID = 0xFFFFFFFF
