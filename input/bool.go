// Copyright (c) 2018-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package input

type BoolID uint32

const noBool = BoolID(maxID)

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

func Bool(name string) BoolID {
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

	actions.names[name] = BoolID(a)
	bools.name = append(bools.name, name)

	return BoolID(a)
}

func (a BoolID) Name() string {
	return bools.name[a]
}

func (a BoolID) activate(d DeviceID, b binding) {
	devices.bools[d][a].active = true
	devices.boolbinds[d][a] = append(devices.boolbinds[d][a], b)
	_, v := b.asBool()
	if v {
		devices.bools[d][a].pressed = true
	}
}

func (a BoolID) newframe(d DeviceID) {
	devices.bools[d][a].just = false
	for _, b := range devices.boolbinds[d][a] {
		j, v := b.asBool()
		if j {
			devices.bools[d][a].just = (v != devices.bools[d][a].pressed)
			devices.bools[d][a].pressed = v
		}
	}
}

func (a BoolID) deactivate(d DeviceID) {
	devices.boolbinds[d][a] = devices.boolbinds[d][a][:0]
	devices.bools[d][a].active = false
	devices.bools[d][a].just = false
	devices.bools[d][a].pressed = false
}

func (a BoolID) Active(d DeviceID) bool {
	return devices.bools[d][a].active
}

func (a BoolID) Pressed(d DeviceID) bool {
	return devices.bools[d][a].pressed
}

func (a BoolID) JustPressed(d DeviceID) bool {
	return devices.bools[d][a].just && devices.bools[d][a].pressed
}

func (a BoolID) Released(d DeviceID) bool {
	return !devices.bools[d][a].pressed
}

func (a BoolID) JustReleased(d DeviceID) bool {
	return devices.bools[d][a].just && !devices.bools[d][a].pressed
}
