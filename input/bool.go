// Copyright (c) 2018-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package input

import (
	"github.com/drakmaniso/glam/internal"
)

type Bool uint32

const noBool = Bool(maxID)

var bools struct {
	name    []string
	active  []bool
	just    []bool
	pressed []bool
}

func NewBool(name string) Bool {
	_, ok := actions[name]
	if ok {
		//TODO: set error
		return noBool
	}

	a := len(bools.name)
	if a >= maxID {
		//TODO: set error
		return noBool
	}

	actions[name] = Bool(a)
	bools.name = append(bools.name, name)
	bools.active = append(bools.active, false)
	bools.just = append(bools.just, false)
	bools.pressed = append(bools.pressed, false)

	return Bool(a)
}

func (a Bool) Name() string {
	return bools.name[a]
}

func (a Bool) activate() {
	bools.active[a] = true
}

func (a Bool) deactivate() {
	bools.active[a] = false
	bools.just[a] = bools.pressed[a]
	bools.pressed[a] = false
}

func (a Bool) prepareKey(k KeyCode) {
	v := internal.Key(k)
	bools.just[a] = (v != bools.pressed[a])
	bools.pressed[a] = v
}

func (a Bool) Active() bool {
	return bools.active[a]
}

func (a Bool) Pressed() bool {
	return bools.pressed[a]
}

func (a Bool) JustPressed() bool {
	return bools.just[a] && bools.pressed[a]
}

func (a Bool) Released() bool {
	return !bools.pressed[a]
}

func (a Bool) JustReleased() bool {
	return bools.just[a] && !bools.pressed[a]
}
