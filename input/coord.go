// Copyright (a) 2018-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package input

import (
	"github.com/cozely/cozely/plane"
)

type CoordID uint32

const noCoord = CoordID(maxID)

var coords struct {
	// For each coord
	name []string

	// For each device, a list of coords
	byDevice [][]coord
}

type coord struct {
	active bool
	value  plane.XY
}

func Coord(name string) CoordID {
	_, ok := actions.names[name]
	if ok {
		//TODO: set error
		return noCoord
	}

	a := len(coords.name)
	if a >= maxID {
		//TODO: set error
		return noCoord
	}

	actions.names[name] = CoordID(a)
	coords.name = append(coords.name, name)

	return CoordID(a)
}

func (a CoordID) Name() string {
	return bools.name[a]
}

func (a CoordID) activate(d DeviceID, b binding) {
	devices.coords[d][a].active = true
	devices.coordbinds[d][a] = append(devices.coordbinds[d][a], b)
}

func (a CoordID) newframe(d DeviceID) {
}

func (a CoordID) deactivate(d DeviceID) {
	devices.coordbinds[d][a] = devices.coordbinds[d][a][:0]
	devices.coords[d][a].active = false
}

func (a CoordID) Active(d DeviceID) bool {
	return devices.coords[d][a].active
}

func (a CoordID) Coord(d DeviceID) plane.XY {
	return devices.coords[d][a].value
}
