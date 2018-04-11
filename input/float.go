// Copyright (c) 2018-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package input

type Float uint32

const noFloat = Float(maxID)

var floats struct {
	// For each float
	name []string

	// For each device, a list of floats
	byDevice [][]float
}

type float struct {
	active bool
	value  float32
}

func NewFloat(name string) Float {
	_, ok := actions.names[name]
	if ok {
		//TODO: set error
		return noFloat
	}

	a := len(floats.name)
	if a >= maxID {
		//TODO: set error
		return noFloat
	}

	actions.names[name] = Float(a)
	floats.name = append(floats.name, name)

	return Float(a)
}

func (a Float) Name() string {
	return bools.name[a]
}

func (a Float) activate(d Device, b binding) {
	devices.floats[d][a].active = true
	devices.floatbinds[d][a] = append(devices.floatbinds[d][a], b)
}

func (a Float) newframe(d Device) {
}

func (a Float) deactivate(d Device) {
	devices.floatbinds[d][a] = devices.floatbinds[d][a][:0]
	devices.floats[d][a].active = false
}

func (a Float) Active(d Device) bool {
	return devices.floats[d][a].active
}

func (a Float) Value(d Device) float32 {
	return devices.floats[d][a].value
}
