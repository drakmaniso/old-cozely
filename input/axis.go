// Copyright (c) 2018-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package input

import (
	"errors"

	"github.com/cozely/cozely/internal"
)

////////////////////////////////////////////////////////////////////////////////

// AxisID identifies an absolute, one-dimensional analog action with a
// resting position and two directions, i.e. any action that is best represented
// by one floating-point value between -1 and +1.
type AxisID uint32

const noAxis = AxisID(maxID)

var axes struct {
	// For each axis
	name []string
}

type axis struct {
	value    float32
	previous float32
}

////////////////////////////////////////////////////////////////////////////////

// Axis declares a new bipolar action, and returns its ID.
func Axis(name string) AxisID {
	if internal.Running {
		setErr(errors.New("input axis declaration: declarations must happen before starting the framework"))
		return noAxis
	}

	_, ok := actions.name[name]
	if ok {
		setErr(errors.New("input axis declaration: name already taken by another action"))
		return noAxis
	}

	a := AxisID(len(axes.name))
	if a >= maxID {
		setErr(errors.New("input axis declaration: too many axis actions"))
		return noAxis
	}

	actions.name[name] = AxisID(a)
	actions.list = append(actions.list, AxisID(a))
	axes.name = append(axes.name, name)

	return AxisID(a)
}

////////////////////////////////////////////////////////////////////////////////

// Name of the action.
func (a AxisID) Name() string {
	return axes.name[a]
}

// Value returns the current value of the action on the current device. This
// value is normalized between -1 and +1.
func (a AxisID) Value() float32 {
	return a.ValueOn(devices.current)
}

// ValueOn returns the current value of the action on a specific device. This
// value is normalized between -1 and +1.
func (a AxisID) ValueOn(d DeviceID) float32 {
	return devices.axes[d][a].value
}

////////////////////////////////////////////////////////////////////////////////

func (a AxisID) activate(d DeviceID, b source) {
	devices.axesbinds[d][a] = append(devices.axesbinds[d][a], b)
}

func (a AxisID) newframe(d DeviceID) {
	devices.axes[d][a].previous = devices.axes[d][a].value
}

func (a AxisID) update(d DeviceID) {
	for _, s := range devices.axesbinds[d][a] {
		j, v := s.asAxis()
		if j {
			devices.axes[d][a].value = v
			devices.current = d
		}
	}
}

func (a AxisID) deactivate(d DeviceID) {
	devices.axesbinds[d][a] = devices.axesbinds[d][a][:0]
}
