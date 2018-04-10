// Copyright (c) 2018-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package input

type Float uint32

const noFloat = Float(maxID)

var floats struct {
	name   []string
	active [][maxDevices]bool
	value  [][maxDevices]float32
}

func NewFloat(name string) Float {
	_, ok := actions[name]
	if ok {
		//TODO: set error
		return noFloat
	}

	a := len(floats.name)
	if a >= maxID {
		//TODO: set error
		return noFloat
	}

	actions[name] = Float(a)
	floats.name = append(floats.name, name)
	floats.active = append(floats.active, [maxDevices]bool{})
	floats.value = append(floats.value, [maxDevices]float32{})

	return Float(a)
}

func (a Float) Name() string {
	return bools.name[a]
}

func (a Float) deactivate(d Device) {
	floats.active[a][d] = false
}

func (a Float) activateKey(k KeyCode) {
	floats.active[a][Keyboard] = true
}

func (a Float) prepareKey(k KeyCode) {
}

func (a Float) Active(d Device) bool {
	return floats.active[a][d]
}

func (a Float) Value(d Device) float32 {
	return floats.value[a][d]
}
