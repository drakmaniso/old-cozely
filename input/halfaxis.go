// Copyright (c) 2018-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package input

import (
	"errors"

	"github.com/cozely/cozely/internal"
)

////////////////////////////////////////////////////////////////////////////////

// HalfAxisID identifies an absolute, one-dimensional analog action with a
// resting position and one signle direction, i.e. any action that is best
// represented by one floating-point value between 0 and 1.
type HalfAxisID uint32

const noHalfAxis = HalfAxisID(maxID)

var halfaxes struct {
	// For each float
	name []string
}

type halfaxis struct {
	value    float32
	previous float32
}

////////////////////////////////////////////////////////////////////////////////

// HalfAxis declares a new unipolar analog action, and returns its ID.
func HalfAxis(name string) HalfAxisID {
	if internal.Running {
		setErr(errors.New("input half-axis declaration: declarations must happen before starting the framework"))
		return noHalfAxis
	}

	_, ok := actions.name[name]
	if ok {
		setErr(errors.New("input half-axis declaration: name already taken by another action"))
		return noHalfAxis
	}

	a := len(halfaxes.name)
	if a >= maxID {
		setErr(errors.New("input half-axis declaration: too many half-axis actions"))
		return noHalfAxis
	}

	actions.name[name] = HalfAxisID(a)
	actions.list = append(actions.list, HalfAxisID(a))
	halfaxes.name = append(halfaxes.name, name)

	return HalfAxisID(a)
}

////////////////////////////////////////////////////////////////////////////////

// Name of the action.
func (a HalfAxisID) Name() string {
	return halfaxes.name[a]
}

// Value returns the current value of the action on the current device. This
// value is normalized between 0 and 1.
func (a HalfAxisID) Value() float32 {
	return a.ValueOn(Any)
}

// ValueOn returns the current value of the action on a specific device. This
// value is normalized between 0 and 1.
func (a HalfAxisID) ValueOn(d DeviceID) float32 {
	return devices.halfaxes[d][a].value
}

////////////////////////////////////////////////////////////////////////////////

func (a HalfAxisID) activate(d DeviceID, b source) {
	devices.halfaxesbinds[d][a] = append(devices.halfaxesbinds[d][a], b)
}

func (a HalfAxisID) newframe(d DeviceID) {
	devices.halfaxes[d][a].previous = devices.halfaxes[d][a].value
}

func (a HalfAxisID) update(d DeviceID) {
	for _, b := range devices.halfaxesbinds[d][a] {
		j, v := b.asUnipolar()
		if j {
			devices.halfaxes[d][a].value = v
			devices.halfaxes[0][a].value = v
		}
	}
}

func (a HalfAxisID) deactivate(d DeviceID) {
	devices.halfaxesbinds[d][a] = devices.halfaxesbinds[d][a][:0]
}
