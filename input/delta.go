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
	active [][maxDevices]bool
	value  [][maxDevices]plane.Coord
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
	deltas.active = append(deltas.active, [maxDevices]bool{})
	deltas.value = append(deltas.value, [maxDevices]plane.Coord{})

	return Delta(a)
}

func (a Delta) Name() string {
	return bools.name[a]
}

func (a Delta) activate(b binding) {
	floats.active[a][b.device()] = true
}

func (a Delta) newframe(b binding) {
}

func (a Delta) prepare(b binding) {
}

func (a Delta) deactivate(d Device) {
	deltas.active[a][d] = false
}

func (a Delta) Active(d Device) bool {
	return deltas.active[a][d]
}

func (a Delta) Delta(d Device) plane.Coord {
	return deltas.value[a][d]
}
