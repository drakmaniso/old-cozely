// Copyright (c) 2018-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package input

import (
	"github.com/drakmaniso/cozely/plane"
)

type DeltaID uint32

const noDelta = DeltaID(maxID)

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

func Delta(name string) DeltaID {
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

	actions.names[name] = DeltaID(a)
	deltas.name = append(deltas.name, name)

	return DeltaID(a)
}

func (a DeltaID) Name() string {
	return bools.name[a]
}

func (a DeltaID) activate(d DeviceID, b binding) {
	devices.deltas[d][a].active = true
	devices.deltabinds[d][a] = append(devices.deltabinds[d][a], b)
}

func (a DeltaID) newframe(d DeviceID) {
}

func (a DeltaID) deactivate(d DeviceID) {
	devices.deltabinds[d][a] = devices.deltabinds[d][a][:0]
	devices.deltas[d][a].active = false
}

func (a DeltaID) Active(d DeviceID) bool {
	return devices.deltas[d][a].active
}

func (a DeltaID) Delta(d DeviceID) plane.Coord {
	return devices.deltas[d][a].value
}
