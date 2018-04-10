// Copyright (c) 2018-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package input

type Device uint32

const noDevice = Device(maxID)

const (
	Any              Device = 0
	KeyboardAndMouse Device = iota
)

const maxDevices = 16

var devices struct {
	// For each device
	name    []string
	context []Context
	new     []Context

	// For each device/action combination
	bools  [][]boolean
	floats [][]float
	coords [][]coord
	deltas [][]delta

	// For each device/context combination, a list of bindings
	bindings [][][]binding
}

func NewDevice(name string) Device {
	l := len(devices.name)
	if l >= maxID {
		//TODO: set error
		return Device(maxID)
	}

	a := Device(l)
	devices.name = append(devices.name, name)
	devices.context = append(devices.context, 0)
	devices.new = append(devices.new, 0)

	n := len(bools.name)
	devices.bools = append(devices.bools, make([]boolean, n))

	n = len(contexts.name)
	devices.bindings = append(devices.bindings, make([][]binding, n))

	return a
}

func clearDevices() {
	devices.name = nil
	devices.context = nil
	devices.new = nil
	devices.bools = nil
	devices.bindings = nil

	NewDevice("Any")
	NewDevice("KeyboardAndMouse")
}
