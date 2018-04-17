// Copyright (c) 2018-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package input

import (
	"errors"

	"github.com/cozely/cozely/internal"
)

////////////////////////////////////////////////////////////////////////////////

// BipolarID identifies an absolute, one-dimensional analog action with a
// resting position and two directions, i.e. any action that is best represented
// by one floating-point value between -1 and +1.
type BipolarID uint32

const noBipolar = BipolarID(maxID)

var bipolars struct {
	// For each bipolar
	name []string
}

type bipolar struct {
	active   bool
	value    float32
	previous float32
}

////////////////////////////////////////////////////////////////////////////////

// Bipolar declares a new bipolar action, and returns its ID.
func Bipolar(name string) BipolarID {
	if internal.Running {
		setErr(errors.New("input bipolar declaration: declarations must happen before starting the framework"))
		return noBipolar
	}

	_, ok := actions.name[name]
	if ok {
		setErr(errors.New("input bipolar declaration: name already taken by another action"))
		return noBipolar
	}

	a := len(bipolars.name)
	if a >= maxID {
		setErr(errors.New("input bipolar declaration: too many bipolar actions"))
		return noBipolar
	}

	actions.name[name] = BipolarID(a)
	actions.list = append(actions.list, BipolarID(a))
	bipolars.name = append(bipolars.name, name)

	return BipolarID(a)
}

////////////////////////////////////////////////////////////////////////////////

// Name of the action.
func (a BipolarID) Name() string {
	return bipolars.name[a]
}

// Active returns true if the action is currently active on a specific device
// (i.e. if it is listed in the context currently active on the device).
func (a BipolarID) Active(d DeviceID) bool {
	return devices.bipolars[d][a].active
}

// Value returns the current value of the action on a specific device. This
// value is normalized between -1 and +1.
func (a BipolarID) Value(d DeviceID) float32 {
	return devices.bipolars[d][a].value
}

////////////////////////////////////////////////////////////////////////////////

func (a BipolarID) activate(d DeviceID, b binding) {
	devices.bipolars[d][a].active = true
	devices.bipolarbinds[d][a] = append(devices.bipolarbinds[d][a], b)
}

func (a BipolarID) newframe(d DeviceID) {
	devices.bipolars[d][a].previous = devices.bipolars[d][a].value
}

func (a BipolarID) update(d DeviceID) {
	for _, b := range devices.bipolarbinds[d][a] {
		j, v := b.asBipolar()
		if j {
			devices.bipolars[d][a].value = v
			devices.bipolars[0][a].value = v
		}
	}
}

func (a BipolarID) deactivate(d DeviceID) {
	devices.bipolarbinds[d][a] = devices.bipolarbinds[d][a][:0]
	devices.bipolars[d][a].active = false
}
