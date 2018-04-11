// Copyright (c) 2018-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package input

type Action interface {
	Active(d Device) bool
	deactivate(d Device)
	activate(d Device, b binding)
	newframe(d Device)
}

var actions = struct {
	names map[string]Action
}{
	names: map[string]Action{},
}

const maxID = 0xFFFFFFFF
