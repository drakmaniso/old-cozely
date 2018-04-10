// Copyright (c) 2018-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package input

type Float uint32

const noFloat = Float(maxID)

var floats struct {
	name   []string
	active []bool
	value  []float32
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
	floats.active = append(floats.active, false)
	floats.value = append(floats.value, 0)

	return Float(a)
}

func (a Float) Name() string {
	return bools.name[a]
}

func (a Float) activate() {
	floats.active[a] = true
}

func (a Float) deactivate() {
	floats.active[a] = false
}

func (a Float) prepareKey(k KeyCode) {
}

func (a Float) Active() bool {
	return floats.active[a]
}

func (a Float) Value() float32 {
	return floats.value[a]
}
