// Copyright (c) 2018-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package input

import (
	"errors"

	"github.com/cozely/cozely/internal"
)

////////////////////////////////////////////////////////////////////////////////

// UnipolarID identifies an absolute, one-dimensional analog action with a
// resting position and one signle direction, i.e. any action that is best
// represented by one floating-point value between 0 and 1.
type UnipolarID uint32

const noUnipolar = UnipolarID(maxID)

var unipolars struct {
	// For each float
	name []string
}

type unipolar struct {
	active   bool
	value    float32
	previous float32
}

////////////////////////////////////////////////////////////////////////////////

// Unipolar declares a new unipolar analog action, and returns its ID.
func Unipolar(name string) UnipolarID {
	if internal.Running {
		setErr(errors.New("input unipolar declaration: declarations must happen before starting the framework"))
		return noUnipolar
	}

	_, ok := actions.name[name]
	if ok {
		setErr(errors.New("input unipolar declaration: name already taken by another action"))
		return noUnipolar
	}

	a := len(unipolars.name)
	if a >= maxID {
		setErr(errors.New("input unipolar declaration: too many unipolar actions"))
		return noUnipolar
	}

	actions.name[name] = UnipolarID(a)
	actions.list = append(actions.list, UnipolarID(a))
	unipolars.name = append(unipolars.name, name)

	return UnipolarID(a)
}

////////////////////////////////////////////////////////////////////////////////

// Name of the action.
func (a UnipolarID) Name() string {
	return unipolars.name[a]
}

// Active returns true if the action is currently active on a specific device
// (i.e. if it is listed in the context currently active on the device).
func (a UnipolarID) Active(d DeviceID) bool {
	return devices.unipolars[d][a].active
}

// Value returns the current value of the action on a specific device. This
// value is normalized between 0 and 1.
func (a UnipolarID) Value(d DeviceID) float32 {
	return devices.unipolars[d][a].value
}

////////////////////////////////////////////////////////////////////////////////

func (a UnipolarID) activate(d DeviceID, b source) {
	devices.unipolars[d][a].active = true
	devices.unipolarbinds[d][a] = append(devices.unipolarbinds[d][a], b)
}

func (a UnipolarID) newframe(d DeviceID) {
	devices.unipolars[d][a].previous = devices.unipolars[d][a].value
}

func (a UnipolarID) update(d DeviceID) {
	for _, b := range devices.unipolarbinds[d][a] {
		j, v := b.asUnipolar()
		if j {
			devices.unipolars[d][a].value = v
			devices.unipolars[0][a].value = v
		}
	}
}

func (a UnipolarID) deactivate(d DeviceID) {
	devices.unipolarbinds[d][a] = devices.unipolarbinds[d][a][:0]
	devices.unipolars[d][a].active = false
}
