// Copyright (c) 2018-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package input

////////////////////////////////////////////////////////////////////////////////

// DeviceID identifies an input device, i.e. any kind of hardware that can be
// bound to a game action.
//
// Note: for convenience, the mouse and keyboard are considered to be the same
// device, and share the same ID.
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
	bools     [][]boolean
	unipolars [][]unipolar
	bipolars  [][]bipolar
	coords    [][]coordinates
	deltas    [][]delta

	// For each device/context combination, the list of bindings
	bindings [][][]source

	// For each device/action combination, the list of *current* bindings
	boolbinds     [][][]source
	unipolarbinds [][][]source
	bipolarbinds  [][][]source
	coordbinds    [][][]source
	deltabinds    [][][]source
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
	devices.boolbinds = append(devices.boolbinds, make([][]source, n))

	n = len(unipolars.name)
	devices.unipolars = append(devices.unipolars, make([]unipolar, n))
	devices.unipolarbinds = append(devices.unipolarbinds, make([][]source, n))

	n = len(bipolars.name)
	devices.bipolars = append(devices.bipolars, make([]bipolar, n))
	devices.bipolarbinds = append(devices.bipolarbinds, make([][]source, n))

	n = len(coords.name)
	devices.coords = append(devices.coords, make([]coordinates, n))
	devices.coordbinds = append(devices.coordbinds, make([][]source, n))

	n = len(deltas.name)
	devices.deltas = append(devices.deltas, make([]delta, n))
	devices.deltabinds = append(devices.deltabinds, make([][]source, n))

	n = len(contexts.name)
	devices.bindings = append(devices.bindings, make([][]source, n))

	return a
}

////////////////////////////////////////////////////////////////////////////////

func clearDevices() {
	devices.name = nil
	devices.context = nil
	devices.newcontext = nil
	devices.bools = nil
	devices.boolbinds = nil
	devices.unipolars = nil
	devices.unipolarbinds = nil
	devices.bipolars = nil
	devices.bipolarbinds = nil
	devices.coords = nil
	devices.coordbinds = nil
	devices.deltas = nil
	devices.deltabinds = nil
	devices.bindings = nil

	addDevice("Any")
	addDevice("KeyboardAndMouse")
}

////////////////////////////////////////////////////////////////////////////////

func (a DeviceID) isMouse() bool {
	return a == kbmouse
}
