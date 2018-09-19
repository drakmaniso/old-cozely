// Copyright (a) 2018-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package input

import (
	"errors"

	"github.com/cozely/cozely/coord"
	"github.com/cozely/cozely/internal"
)

// DualAxisID identifies an absolute two-dimensional analog action, i.e. any action
// that is best represented by a pair of X and Y coordinates, and whose most
// important characteristic is the current position. The values of the
// coordinates are normalized between -1 and 1.
type DualAxisID uint32

const noDualAxis = DualAxisID(maxID)

var dualaxes struct {
	// For each coord
	name []string
}

type dualaxis struct {
	value    coord.XY
	previous coord.XY
}

// DualAxis declares a new two-dimensional analog action, and returns its ID.
func DualAxis(name string) DualAxisID {
	if internal.Running {
		setErr(errors.New("input dual-axis declaration: declarations must happen before starting the framework"))
		return noDualAxis
	}

	_, ok := actions.name[name]
	if ok {
		setErr(errors.New("input dual-axis declaration: name already taken by another action"))
		return noDualAxis
	}

	a := DualAxisID(len(dualaxes.name))
	if a >= maxID {
		setErr(errors.New("input dual-axis declaration: too many dual-axis actions"))
		return noDualAxis
	}

	actions.name[name] = DualAxisID(a)
	actions.list = append(actions.list, DualAxisID(a))
	dualaxes.name = append(dualaxes.name, name)

	return DualAxisID(a)
}

////////////////////////////////////////////////////////////////////////////////

// Name of the action.
func (a DualAxisID) Name() string {
	return dualaxes.name[a]
}

// XY returns the current status of the action on the current device. The
// coordinates are the current absolute position; the values of X and Y are
// normalized between -1 and 1.
func (a DualAxisID) XY() coord.XY {
	return a.XYon(devices.current)
}

// XYon returns the current status of the action on a specific device. The
// coordinates are the current absolute position; the values of X and Y are
// normalized between -1 and 1.
func (a DualAxisID) XYon(d DeviceID) coord.XY {
	return devices.dualaxes[d][a].value
}

////////////////////////////////////////////////////////////////////////////////

func (a DualAxisID) activate(d DeviceID, s source) {
	devices.dualaxesbinds[d][a] = append(devices.dualaxesbinds[d][a], s)
}

func (a DualAxisID) newframe(d DeviceID) {
	devices.dualaxes[d][a].previous = devices.dualaxes[d][a].value
}

func (a DualAxisID) update(d DeviceID) {
	for _, s := range devices.dualaxesbinds[d][a] {
		j, v := s.asDualAxis()
		if j {
			devices.dualaxes[d][a].value = v
			devices.current = d
		}
	}
}

func (a DualAxisID) deactivate(d DeviceID) {
	devices.dualaxesbinds[d][a] = devices.dualaxesbinds[d][a][:0]
}
