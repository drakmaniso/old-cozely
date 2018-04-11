// Copyright (c) 2018-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package input

type Action interface {
	Active(d DeviceID) bool
	deactivate(d DeviceID)
	activate(d DeviceID, b binding)
	newframe(d DeviceID)
}

var actions = struct {
	names map[string]Action
}{
	names: map[string]Action{},
}

const maxID = 0xFFFFFFFF
