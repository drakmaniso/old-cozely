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
	active  [][maxDevices]bool
	just    [][maxDevices]bool
	pressed [][maxDevices]bool
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
	bools.active = append(bools.active, [maxDevices]bool{})
	bools.just = append(bools.just, [maxDevices]bool{})
	bools.pressed = append(bools.pressed, [maxDevices]bool{})

	return Bool(a)
}

func (a Bool) Name() string {
	return bools.name[a]
}

func (a Bool) deactivate(d Device) {
	bools.active[a][d] = false
	bools.just[a][d] = false
	bools.pressed[a][d] = false
}

func (a Bool) activateKey(k KeyCode) {
	bools.active[a][Keyboard] = true
	v := internal.Key(k)
	bools.pressed[a][Keyboard] = v
}

func (a Bool) prepareKey(k KeyCode) {
	v := internal.Key(k)
	bools.just[a][Keyboard] = (v != bools.pressed[a][Keyboard])
	bools.pressed[a][Keyboard] = v
}

func (a Bool) Active(d Device) bool {
	return bools.active[a][d]
}

func (a Bool) Pressed(d Device) bool {
	return bools.pressed[a][d]
}

func (a Bool) JustPressed(d Device) bool {
	return bools.just[a][d] && bools.pressed[a][d]
}

func (a Bool) Released(d Device) bool {
	return !bools.pressed[a][d]
}

func (a Bool) JustReleased(d Device) bool {
	return bools.just[a][d] && !bools.pressed[a][d]
}
