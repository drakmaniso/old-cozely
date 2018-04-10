// Copyright (c) 2018-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package input

type Float uint32

const noFloat = Float(maxID)

var floats struct {
	name   []string
}
type float struct {
	active bool
	value float32
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

	return Float(a)
}

func (a Float) Name() string {
	return bools.name[a]
}

func (a Float) activate(b binding) {
	d := b.device()
	devices.floats[d][a].active = true
}

func (a Float) newframe(b binding) {
}

func (a Float) prepare(b binding) {
}

func (a Float) deactivate(d Device) {
	devices.floats[d][a].active = false
}

func (a Float) Active(d Device) bool {
	return devices.floats[d][a].active
}

func (a Float) Value(d Device) float32 {
	return devices.floats[d][a].value
}
