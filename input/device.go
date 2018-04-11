// Copyright (c) 2018-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package input

////////////////////////////////////////////////////////////////////////////////

type DeviceID uint32

const noDevice = DeviceID(maxID)

const (
	anydev  DeviceID = 0
	kbmouse DeviceID = iota
)

const maxdevices = 16

var devices struct {
	// For each device
	name       []string
	context    []ContextID
	newcontext []ContextID

	// For each device/action combination
	bools  [][]boolean
	floats [][]float
	coords [][]coord
	deltas [][]delta

	// For each device/context combination, the list of bindings
	bindings [][][]binding

	// For each device/action combination, the list of *current* bindings
	boolbinds  [][][]binding
	floatbinds [][][]binding
	coordbinds [][][]binding
	deltabinds [][][]binding
}

////////////////////////////////////////////////////////////////////////////////

func addDevice(name string) DeviceID {
	l := len(devices.name)
	if l >= maxID {
		//TODO: set error
		return DeviceID(maxID)
	}

	a := DeviceID(l)
	devices.name = append(devices.name, name)
	devices.context = append(devices.context, noContext)
	devices.newcontext = append(devices.newcontext, 0)

	n := len(bools.name)
	devices.bools = append(devices.bools, make([]boolean, n))
	devices.boolbinds = append(devices.boolbinds, make([][]binding, n))

	n = len(floats.name)
	devices.floats = append(devices.floats, make([]float, n))
	devices.floatbinds = append(devices.floatbinds, make([][]binding, n))

	n = len(coords.name)
	devices.coords = append(devices.coords, make([]coord, n))
	devices.coordbinds = append(devices.coordbinds, make([][]binding, n))

	n = len(deltas.name)
	devices.deltas = append(devices.deltas, make([]delta, n))
	devices.deltabinds = append(devices.deltabinds, make([][]binding, n))

	n = len(contexts.name)
	devices.bindings = append(devices.bindings, make([][]binding, n))

	return a
}

////////////////////////////////////////////////////////////////////////////////

func clearDevices() {
	devices.name = nil
	devices.context = nil
	devices.newcontext = nil
	devices.bools = nil
	devices.boolbinds = nil
	devices.floats = nil
	devices.floatbinds = nil
	devices.coords = nil
	devices.coordbinds = nil
	devices.deltas = nil
	devices.deltabinds = nil
	devices.bindings = nil

	addDevice("Any")
	addDevice("KeyboardAndMouse")
}

////////////////////////////////////////////////////////////////////////////////
