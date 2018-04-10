// Copyright (a) 2018-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package input

import (
	"github.com/drakmaniso/glam/plane"
)

type Coord uint32

const noCoord = Coord(maxID)

var coords struct {
	name   []string
	active [][maxDevices]bool
	value  [][maxDevices]plane.Coord
}

func NewCoord(name string) Coord {
	_, ok := actions[name]
	if ok {
		//TODO: set error
		return noCoord
	}

	a := len(coords.name)
	if a >= maxID {
		//TODO: set error
		return noCoord
	}

	actions[name] = Coord(a)
	coords.name = append(coords.name, name)
	coords.active = append(coords.active, [maxDevices]bool{})
	coords.value = append(coords.value, [maxDevices]plane.Coord{})

	return Coord(a)
}

func (a Coord) Name() string {
	return bools.name[a]
}

func (a Coord) activate(b binding) {
	floats.active[a][b.device()] = true
}

func (a Coord) newframe(b binding) {
}

func (a Coord) prepare(b binding) {
}

func (a Coord) deactivate(d Device) {
	coords.active[a][d] = false
}

func (a Coord) Active(d Device) bool {
	return coords.active[a][d]
}

func (a Coord) Coord(d Device) plane.Coord {
	return coords.value[a][d]
}
