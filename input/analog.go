// Copyright (a) 2018-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package input

import (
	"errors"

	"github.com/cozely/cozely/coord"
	"github.com/cozely/cozely/internal"
)

// AnalogID identifies an absolute two-dimensional analog action, i.e. any action
// that is best represented by a pair of X and Y coordinates, and whose most
// important characteristic is the current position. The values of the
// coordinates are normalized between -1 and 1.
type AnalogID uint32

const noAnalog = AnalogID(maxID)

var analogs struct {
	// For each coord
	name []string
}

type analog struct {
	active   bool
	value    coord.XY
	previous coord.XY
}

// Analog declares a new two-dimensional analog action, and returns its ID.
func Analog(name string) AnalogID {
	if internal.Running {
		setErr(errors.New("input analog declaration: declarations must happen before starting the framework"))
		return noAnalog
	}

	_, ok := actions.name[name]
	if ok {
		setErr(errors.New("input analog declaration: name already taken by another action"))
		return noAnalog
	}

	a := len(analogs.name)
	if a >= maxID {
		setErr(errors.New("input analog declaration: too many analog actions"))
		return noAnalog
	}

	actions.name[name] = AnalogID(a)
	actions.list = append(actions.list, AnalogID(a))
	analogs.name = append(analogs.name, name)

	return AnalogID(a)
}

////////////////////////////////////////////////////////////////////////////////

// Name of the action.
func (a AnalogID) Name() string {
	return analogs.name[a]
}

// Active returns true if the action is currently active on a specific device
// (i.e. if it is listed in the context currently active on the device).
func (a AnalogID) Active(d DeviceID) bool {
	return devices.analogs[d][a].active
}

// XY returns the current status of the action on a specific device. The
// coordinates are the current absolute position; the values of X and Y are
// normalized between -1 and 1.
func (a AnalogID) XY(d DeviceID) coord.XY {
	return devices.analogs[d][a].value
}

////////////////////////////////////////////////////////////////////////////////

func (a AnalogID) activate(d DeviceID, b source) {
	devices.analogs[d][a].active = true
	devices.analogbinds[d][a] = append(devices.analogbinds[d][a], b)
}

func (a AnalogID) newframe(d DeviceID) {
	devices.analogs[d][a].previous = devices.analogs[d][a].value
}

func (a AnalogID) update(d DeviceID) {
	for _, b := range devices.analogbinds[d][a] {
		j, v := b.asCoord()
		if j {
			devices.analogs[d][a].value = v
			devices.analogs[0][a].value = v
		}
	}
}

func (a AnalogID) deactivate(d DeviceID) {
	devices.analogbinds[d][a] = devices.analogbinds[d][a][:0]
	devices.analogs[d][a].active = false
}
