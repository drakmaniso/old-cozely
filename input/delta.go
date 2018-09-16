// Copyright (c) 2018-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package input

import (
	"errors"

	"github.com/cozely/cozely/coord"
	"github.com/cozely/cozely/internal"
)

////////////////////////////////////////////////////////////////////////////////

// DeltaID identifes a relative two-dimensional analog input, i.e. any action
// that is best represented by a pair of X and Y coordinates, and whose most
// important characteristic is the change in position.
type DeltaID uint32

const noDelta = DeltaID(maxID)

var deltas struct {
	// For each delta
	name []string
}

type delta struct {
	value    coord.XY
	previous coord.XY
}

////////////////////////////////////////////////////////////////////////////////

// Delta declares a new delta action, and returns its ID.
func Delta(name string) DeltaID {
	if internal.Running {
		setErr(errors.New("input delta declaration: declarations must happen before starting the framework"))
		return noDelta
	}

	_, ok := actions.name[name]
	if ok {
		setErr(errors.New("input delta declaration: name already taken by another action"))
		return noDelta
	}

	a := len(deltas.name)
	if a >= maxID {
		setErr(errors.New("input delta declaration: too many delta actions"))
		return noDelta
	}

	actions.name[name] = DeltaID(a)
	actions.list = append(actions.list, DeltaID(a))
	deltas.name = append(deltas.name, name)

	return DeltaID(a)
}

////////////////////////////////////////////////////////////////////////////////

// Name of the action.
func (a DeltaID) Name() string {
	return dualaxes.name[a]
}

// XY returns the current status of the action on the current device. The
// coordinates correspond to the change in position since the last frame.
func (a DeltaID) XY() coord.XY {
	return a.XYon(devices.current)
}

// XYon returns the current status of the action on a specific device. The
// coordinates correspond to the change in position since the last frame.
func (a DeltaID) XYon(d DeviceID) coord.XY {
	return devices.deltas[d][a].value
}

////////////////////////////////////////////////////////////////////////////////

func (a DeltaID) activate(d DeviceID, s source) {
	devices.deltasbinds[d][a] = append(devices.deltasbinds[d][a], s)
}

func (a DeltaID) newframe(d DeviceID) {
	devices.deltas[d][a].previous = devices.deltas[d][a].value
	devices.deltas[d][a].value = coord.XY{}
}

func (a DeltaID) update(d DeviceID) {
	j := false
	v := coord.XY{}
	for _, s := range devices.deltasbinds[d][a] {
		j, v = s.asDelta()
		devices.deltas[d][a].value = devices.deltas[d][a].value.Plus(v)
		if j {
			devices.current = d
		}
	}
}

func (a DeltaID) deactivate(d DeviceID) {
	devices.deltasbinds[d][a] = devices.deltasbinds[d][a][:0]
}
