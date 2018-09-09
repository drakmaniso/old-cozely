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
	// Any is a pseudo-device that aggregates the status of all other devices.
	Any     DeviceID = 0
	kbmouse DeviceID = iota
)

const maxdevices = 16

var devices struct {
	// For each device
	name       []string
	context    []ContextID
	newcontext []ContextID

	// For each device/action combination
	buttons  [][]button
	halfaxes [][]halfaxis
	axes     [][]axis
	dualaxes [][]dualaxis
	cursors  [][]cursor
	deltas   [][]delta

	// For each device/context combination, the list of bindings
	bindings [][][]source

	// For each device/action combination, the list of *current* bindings
	buttonsbinds  [][][]source
	halfaxesbinds [][][]source
	axesbinds     [][][]source
	dualaxesbinds [][][]source
	cursorsbinds  [][][]source
	deltasbinds   [][][]source
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

	n := len(buttons.name)
	devices.buttons = append(devices.buttons, make([]button, n))
	devices.buttonsbinds = append(devices.buttonsbinds, make([][]source, n))

	n = len(halfaxes.name)
	devices.halfaxes = append(devices.halfaxes, make([]halfaxis, n))
	devices.halfaxesbinds = append(devices.halfaxesbinds, make([][]source, n))

	n = len(axes.name)
	devices.axes = append(devices.axes, make([]axis, n))
	devices.axesbinds = append(devices.axesbinds, make([][]source, n))

	n = len(dualaxes.name)
	devices.dualaxes = append(devices.dualaxes, make([]dualaxis, n))
	devices.dualaxesbinds = append(devices.dualaxesbinds, make([][]source, n))

	n = len(cursors.name)
	devices.cursors = append(devices.cursors, make([]cursor, n))
	devices.cursorsbinds = append(devices.cursorsbinds, make([][]source, n))

	n = len(deltas.name)
	devices.deltas = append(devices.deltas, make([]delta, n))
	devices.deltasbinds = append(devices.deltasbinds, make([][]source, n))

	n = len(contexts.name)
	devices.bindings = append(devices.bindings, make([][]source, n))

	return a
}

////////////////////////////////////////////////////////////////////////////////

func clearDevices() {
	devices.name = nil
	devices.context = nil
	devices.newcontext = nil
	devices.buttons = nil
	devices.buttonsbinds = nil
	devices.halfaxes = nil
	devices.halfaxesbinds = nil
	devices.axes = nil
	devices.axesbinds = nil
	devices.dualaxes = nil
	devices.dualaxesbinds = nil
	devices.cursors = nil
	devices.cursorsbinds = nil
	devices.deltas = nil
	devices.deltasbinds = nil
	devices.bindings = nil

	addDevice("Any")
	addDevice("KeyboardAndMouse")
}

////////////////////////////////////////////////////////////////////////////////

func (a DeviceID) isMouse() bool {
	return a == kbmouse
}
