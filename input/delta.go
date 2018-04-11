// Copyright (c) 2018-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package input

import (
	"github.com/drakmaniso/glam/plane"
)

type Delta uint32

const noDelta = Delta(maxID)

var deltas struct {
	// For each delta
	name []string

	// For each device, a list of deltas
	byDevice [][]delta
}

type delta struct {
	active bool
	value  plane.Coord
}

func NewDelta(name string) Delta {
	_, ok := actions.names[name]
	if ok {
		//TODO: set error
		return noDelta
	}

	a := len(deltas.name)
	if a >= maxID {
		//TODO: set error
		return noDelta
	}

	actions.names[name] = Delta(a)
	deltas.name = append(deltas.name, name)

	return Delta(a)
}

func (a Delta) Name() string {
	return bools.name[a]
}

func (a Delta) activate(d Device, b binding) {
	devices.deltas[d][a].active = true
	devices.deltabinds[d][a] = append(devices.deltabinds[d][a], b)
}

func (a Delta) newframe(d Device) {
}

func (a Delta) deactivate(d Device) {
	devices.deltabinds[d][a] = devices.deltabinds[d][a][:0]
	devices.deltas[d][a].active = false
}

func (a Delta) Active(d Device) bool {
	return devices.deltas[d][a].active
}

func (a Delta) Delta(d Device) plane.Coord {
	return devices.deltas[d][a].value
}
