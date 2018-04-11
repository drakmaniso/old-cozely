// Copyright (c) 2018-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package input

type Bool uint32

const noBool = Bool(maxID)

var bools struct {
	// For each bool
	name []string

	// For each device, a list of bools
	byDevice [][]boolean
}

type boolean struct {
	active  bool
	just    bool
	pressed bool
}

func NewBool(name string) Bool {
	_, ok := actions.names[name]
	if ok {
		//TODO: set error
		return noBool
	}

	a := len(bools.name)
	if a >= maxID {
		//TODO: set error
		return noBool
	}

	actions.names[name] = Bool(a)
	bools.name = append(bools.name, name)

	return Bool(a)
}

func (a Bool) Name() string {
	return bools.name[a]
}

func (a Bool) activate(d Device, b binding) {
	devices.bools[d][a].active = true
	devices.boolbinds[d][a] = append(devices.boolbinds[d][a], b)
	_, v := b.asBool()
	if v {
		devices.bools[d][a].pressed = true
	}
}

func (a Bool) newframe(d Device) {
	devices.bools[d][a].just = false
	for _, b := range devices.boolbinds[d][a] {
		j, v := b.asBool()
		if j {
			devices.bools[d][a].just = (v != devices.bools[d][a].pressed)
			devices.bools[d][a].pressed = v
		}
	}
}

func (a Bool) deactivate(d Device) {
	devices.boolbinds[d][a] = devices.boolbinds[d][a][:0]
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
