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
	active []bool
	value  []plane.Coord
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
	coords.active = append(coords.active, false)
	coords.value = append(coords.value, plane.Coord{})

	return Coord(a)
}

func (a Coord) Name() string {
	return bools.name[a]
}

func (a Coord) activate() {
	coords.active[a] = true
}

func (a Coord) deactivate() {
	coords.active[a] = false
}

func (a Coord) prepareKey(k KeyCode) {
}

func (a Coord) Active() bool {
	return coords.active[a]
}

func (a Coord) Coord() plane.Coord {
	return coords.value[a]
}
