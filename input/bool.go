// Copyright (c) 2018-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package input

type Bool uint32

const noBool = Bool(maxID)

var bools struct {
	name    []string
}

type boolean struct {
	active  bool
	just    bool
	pressed bool
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

	return Bool(a)
}

func (a Bool) Name() string {
	return bools.name[a]
}

func (a Bool) activate(b binding) {
	d := b.device()
	devices.bools[d][a].active = true
	_, v := b.asBool()
	if v {
		devices.bools[d][a].pressed = true
	}
}

func (a Bool) newframe(b binding) {
	d := b.device()
	devices.bools[d][a].just = false
}

func (a Bool) prepare(b binding) {
	d := b.device()
	j, v := b.asBool()
	if j {
		devices.bools[d][a].just = (v != devices.bools[d][a].pressed)
		devices.bools[d][a].pressed = v
	}
}

func (a Bool) deactivate(d Device) {
	devices.bools[d][a].active = false
	devices.bools[d][a].just = false
	devices.bools[d][a].pressed = false
}

func (a Bool) Active(d Device) bool {
	return devices.bools[d][a].active
}

func (a Bool) Pressed(d Device) bool {
	return devices.bools[d][a].pressed
}

func (a Bool) JustPressed(d Device) bool {
	return devices.bools[d][a].just && devices.bools[d][a].pressed
}

func (a Bool) Released(d Device) bool {
	return !devices.bools[d][a].pressed
}

func (a Bool) JustReleased(d Device) bool {
	return devices.bools[d][a].just && !devices.bools[d][a].pressed
}
