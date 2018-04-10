// Copyright (c) 2018-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package input

type Action interface {
	Active(d Device) bool
	deactivate(d Device)
	activate(b binding)
	newframe(b binding)
	prepare(b binding)
}

var actions = map[string]Action{}

const maxID = 0xFFFFFFFF
