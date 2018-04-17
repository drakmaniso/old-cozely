// Copyright (a) 2018-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package input

import (
	"errors"

	"github.com/cozely/cozely/coord"
	"github.com/cozely/cozely/internal"
)

// CoordID identifies an absolute two-dimensional analog action, i.e. any action
// that is best represented by a pair of X and Y coordinates, and whose most
// important characteristic is the current position. The values of the
// coordinates are normalized between -1 and 1.
type CoordID uint32

const noCoord = CoordID(maxID)

var coords struct {
	// For each coord
	name []string
}

type coordinates struct {
	active   bool
	value    coord.XY
	previous coord.XY
}

// Coord declares a new coord action, and returns its ID.
func Coord(name string) CoordID {
	if internal.Running {
		setErr(errors.New("input coord declaration: declarations must happen before starting the framework"))
		return noCoord
	}

	_, ok := actions.name[name]
	if ok {
		setErr(errors.New("input coord declaration: name already taken by another action"))
		return noCoord
	}

	a := len(coords.name)
	if a >= maxID {
		setErr(errors.New("input coord declaration: too many coord actions"))
		return noCoord
	}

	actions.name[name] = CoordID(a)
	actions.list = append(actions.list, CoordID(a))
	coords.name = append(coords.name, name)

	return CoordID(a)
}

////////////////////////////////////////////////////////////////////////////////

// Name of the action.
func (a CoordID) Name() string {
	return bools.name[a]
}

// Active returns true if the action is currently active on a specific device
// (i.e. if it is listed in the context currently active on the device).
func (a CoordID) Active(d DeviceID) bool {
	return devices.coords[d][a].active
}

// Coord returns the current status of the action on a specific device. The
// coordinates are the current absolute position; the values of X and Y are
// normalized between -1 and 1.
func (a CoordID) Coord(d DeviceID) coord.XY {
	return devices.coords[d][a].value
}

////////////////////////////////////////////////////////////////////////////////

func (a CoordID) activate(d DeviceID, b source) {
	devices.coords[d][a].active = true
	devices.coordbinds[d][a] = append(devices.coordbinds[d][a], b)
}

func (a CoordID) newframe(d DeviceID) {
	devices.coords[d][a].previous = devices.coords[d][a].value
}

func (a CoordID) update(d DeviceID) {
	for _, b := range devices.coordbinds[d][a] {
		j, v := b.asCoord()
		if j {
			devices.coords[d][a].value = v
			devices.coords[0][a].value = v
		}
	}
}

func (a CoordID) deactivate(d DeviceID) {
	devices.coordbinds[d][a] = devices.coordbinds[d][a][:0]
	devices.coords[d][a].active = false
}
