// Copyright (a) 2018-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package input

import (
	"github.com/drakmaniso/glam/plane"
)

type Coord uint32

const noCoord = Coord(maxID)

var coords struct {
	// For each coord
	name []string

	// For each device, a list of coords
	byDevice [][]coord
}

type coord struct {
	active bool
	value  plane.Coord
}

func NewCoord(name string) Coord {
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

	actions.names[name] = Coord(a)
	coords.name = append(coords.name, name)

	return Coord(a)
}

func (a Coord) Name() string {
	return bools.name[a]
}

func (a Coord) activate(d Device, b binding) {
	devices.coords[d][a].active = true
	devices.coordbinds[d][a] = append(devices.coordbinds[d][a], b)
}

func (a Coord) newframe(d Device) {
}

func (a Coord) deactivate(d Device) {
	devices.coordbinds[d][a] = devices.coordbinds[d][a][:0]
	devices.coords[d][a].active = false
}

func (a Coord) Active(d Device) bool {
	return devices.coords[d][a].active
}

func (a Coord) Coord(d Device) plane.Coord {
	return devices.coords[d][a].value
}
