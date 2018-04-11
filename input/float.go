// Copyright (c) 2018-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package input

type FloatID uint32

const noFloat = FloatID(maxID)

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

func NewFloat(name string) FloatID {
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

	actions.names[name] = FloatID(a)
	floats.name = append(floats.name, name)

	return FloatID(a)
}

func (a FloatID) Name() string {
	return bools.name[a]
}

func (a FloatID) activate(d DeviceID, b binding) {
	devices.floats[d][a].active = true
	devices.floatbinds[d][a] = append(devices.floatbinds[d][a], b)
}

func (a FloatID) newframe(d DeviceID) {
}

func (a FloatID) deactivate(d DeviceID) {
	devices.floatbinds[d][a] = devices.floatbinds[d][a][:0]
	devices.floats[d][a].active = false
}

func (a FloatID) Active(d DeviceID) bool {
	return devices.floats[d][a].active
}

func (a FloatID) Value(d DeviceID) float32 {
	return devices.floats[d][a].value
}
