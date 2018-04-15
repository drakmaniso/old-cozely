// Copyright (c) 2018-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package input

import (
	"errors"

	"github.com/cozely/cozely/internal"
)

////////////////////////////////////////////////////////////////////////////////

// BoolID identifies a boolean Action, i.e. any action that can be represented
// with two status, "pressed" (aka "on", "started") and "released" (aka "off",
// "stopped").
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

////////////////////////////////////////////////////////////////////////////////

// Bool declares a new bool action, and returns its ID.
func Bool(name string) BoolID {
	if internal.Running {
		setErr(errors.New("input bool declaration: declarations must happen before starting the framework"))
		return noBool
	}

	_, ok := actions.names[name]
	if ok {
		setErr(errors.New("input bool declaration: name already taken by another action"))
		return noBool
	}

	a := len(bools.name)
	if a >= maxID {
		setErr(errors.New("input bool declaration: too many bool actions"))
		return noBool
	}

	actions.names[name] = BoolID(a)
	bools.name = append(bools.name, name)

	return BoolID(a)
}

////////////////////////////////////////////////////////////////////////////////

// Name of the boolean action.
func (a BoolID) Name() string {
	return bools.name[a]
}

// Active returns true if the action is currently active on a specific device
// (i.e. if it is listed in the context currently active on the device).
func (a BoolID) Active(d DeviceID) bool {
	return devices.bools[d][a].active
}

// Pressed returns true if the action is currently pressed (i.e. "on",
// "started") on a specific device.
func (a BoolID) Pressed(d DeviceID) bool {
	return devices.bools[d][a].pressed
}

// JustChanged returns true if the action has just been pressed or released this
// very frame (i.e. just been started or stopped by the player).
//
// Note: this must be queried in the React method of the game loop, as this is
// the only method that is guaranteed to run at least once each frame.
func (a BoolID) JustChanged(d DeviceID) bool {
	return devices.bools[d][a].just
}

// JustPressed returns true if the action has just been pressed this very frame
// (i.e. just been started by the player).
//
// Note: this must be queried in the React method of the game loop, as this is
// the only method that is guaranteed to run at least once each frame.
func (a BoolID) JustPressed(d DeviceID) bool {
	return devices.bools[d][a].just && devices.bools[d][a].pressed
}

// JustReleased returns true if the action has just been released this very
// frame (i.e. just been stopped by the player).
//
// Note: this must be queried in the React method of the game loop, as this is
// the only method that is guaranteed to run at least once each frame.
func (a BoolID) JustReleased(d DeviceID) bool {
	return devices.bools[d][a].just && !devices.bools[d][a].pressed
}

////////////////////////////////////////////////////////////////////////////////

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
