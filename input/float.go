// Copyright (c) 2018-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package input

import (
	"errors"

	"github.com/cozely/cozely/internal"
)

////////////////////////////////////////////////////////////////////////////////

// FloatID identifies an absolute one-dimensional analog action, i.e. any action
// that can be represented by one floating-point value. This value is normalized
// between -1 and 1.
type FloatID uint32

const noFloat = FloatID(maxID)

var floats struct {
	// For each float
	name []string

	// For each device, a list of floats
	byDevice [][]float
}

type float struct {
	active bool
	value  float32
}

////////////////////////////////////////////////////////////////////////////////

// Float declares a new float action, and returns its ID.
func Float(name string) FloatID {
	if internal.Running {
		setErr(errors.New("input float declaration: declarations must happen before starting the framework"))
		return noFloat
	}

	_, ok := actions.names[name]
	if ok {
		setErr(errors.New("input float declaration: name already taken by another action"))
		return noFloat
	}

	a := len(floats.name)
	if a >= maxID {
		setErr(errors.New("input float declaration: too many float actions"))
		return noFloat
	}

	actions.names[name] = FloatID(a)
	floats.name = append(floats.name, name)

	return FloatID(a)
}

////////////////////////////////////////////////////////////////////////////////

// Name of the action.
func (a FloatID) Name() string {
	return bools.name[a]
}

// Active returns true if the action is currently active on a specific device
// (i.e. if it is listed in the context currently active on the device).
func (a FloatID) Active(d DeviceID) bool {
	return devices.floats[d][a].active
}

// Float returns the current value of the action on a specific device. This
// value is normalized between -1 and 1.
func (a FloatID) Float(d DeviceID) float32 {
	return devices.floats[d][a].value
}

////////////////////////////////////////////////////////////////////////////////

func (a FloatID) activate(d DeviceID, b binding) {
	devices.floats[d][a].active = true
	devices.floatbinds[d][a] = append(devices.floatbinds[d][a], b)
}

func (a FloatID) newframe(d DeviceID) {
}

func (a FloatID) deactivate(d DeviceID) {
	devices.floatbinds[d][a] = devices.floatbinds[d][a][:0]
	devices.floats[d][a].active = false
}
