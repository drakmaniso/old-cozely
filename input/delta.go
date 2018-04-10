// Copyright (c) 2018-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package input

import (
	"github.com/drakmaniso/glam/plane"
)

type Delta uint32

const noDelta = Delta(maxID)

var deltas struct {
	name   []string
	active []bool
	value  []plane.Coord
}

func NewDelta(name string) Delta {
	_, ok := actions[name]
	if ok {
		//TODO: set error
		return noDelta
	}

	a := len(deltas.name)
	if a >= maxID {
		//TODO: set error
		return noDelta
	}

	actions[name] = Delta(a)
	deltas.name = append(deltas.name, name)
	deltas.active = append(deltas.active, false)
	deltas.value = append(deltas.value, plane.Coord{})

	return Delta(a)
}

func (c Delta) Name() string {
	return bools.name[c]
}

func (a Delta) activate() {
	deltas.active[a] = true
}

func (a Delta) deactivate() {
	deltas.active[a] = false
}

func (a Delta) prepareKey(k KeyCode) {
}

func (c Delta) Active() bool {
	return deltas.active[c]
}

func (c Delta) Delta() plane.Coord {
	return deltas.value[c]
}
