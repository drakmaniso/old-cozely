// Copyright (c) 2018-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package input

import (
	"github.com/drakmaniso/glam/plane"
)

type Delta uint32

const noDelta = Delta(maxID)

var deltas struct {
	name   []string
}

type delta  struct {
	active bool
	value plane.Coord
}

func NewDelta(name string) Delta {
	_, ok := actions[name]
	if ok {
		//TODO: set error
		return noDelta
	}

	a := len(deltas.name)
	if a >= maxID {
		//TODO: set error
		return noDelta
	}

	actions[name] = Delta(a)
	deltas.name = append(deltas.name, name)

	return Delta(a)
}

func (a Delta) Name() string {
	return bools.name[a]
}

func (a Delta) activate(b binding) {
	d := b.device()
	devices.deltas[d][a].active = true
}

func (a Delta) newframe(b binding) {
}

func (a Delta) prepare(b binding) {
}

func (a Delta) deactivate(d Device) {
	devices.deltas[d][a].active = false
}

func (a Delta) Active(d Device) bool {
	return devices.deltas[d][a].active
}

func (a Delta) Delta(d Device) plane.Coord {
	return devices.deltas[d][a].value
}
