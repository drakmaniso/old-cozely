// Copyright (c) 2018-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package input

import (
	"errors"

	"github.com/cozely/cozely/internal"
)

////////////////////////////////////////////////////////////////////////////////

// DigitalID identifies a digital Action, i.e. any action that can be
// represented with two status, "started" (aka "on", "pressed") and "stopped"
// (aka "off", "released").
type DigitalID uint32

const noDigital = DigitalID(maxID)

var digitals struct {
	// For each digital
	name []string
}

type digital struct {
	active   bool
	previous bool
	pressed  bool
}

////////////////////////////////////////////////////////////////////////////////

// Digital declares a new digital action, and returns its ID.
func Digital(name string) DigitalID {
	if internal.Running {
		setErr(errors.New("input digital declaration: declarations must happen before starting the framework"))
		return noDigital
	}

	_, ok := actions.name[name]
	if ok {
		setErr(errors.New("input digital declaration: name already taken by another action"))
		return noDigital
	}

	a := len(digitals.name)
	if a >= maxID {
		setErr(errors.New("input digital declaration: too many digital actions"))
		return noDigital
	}

	actions.name[name] = DigitalID(a)
	actions.list = append(actions.list, DigitalID(a))
	digitals.name = append(digitals.name, name)

	return DigitalID(a)
}

////////////////////////////////////////////////////////////////////////////////

// Name of the digital action.
func (a DigitalID) Name() string {
	return digitals.name[a]
}

// Active returns true if the action is currently active on a specific device
// (i.e. if it is listed in the context currently active on the device).
func (a DigitalID) Active(d DeviceID) bool {
	return devices.digitals[d][a].active
}

// Ongoing returns true if the action has been started and is currently ongoing
// (i.e. "on", "pressed") on a specific device.
func (a DigitalID) Ongoing(d DeviceID) bool {
	return devices.digitals[d][a].pressed
}

// JustChanged returns true if the action has just been pressed or released this
// very frame (i.e. just been started or stopped by the player).
//
// Note: this *must* be queried in the React method of the game loop, as this is
// the only method that is guaranteed to run at least once each frame.
func (a DigitalID) JustChanged(d DeviceID) bool {
	return devices.digitals[d][a].previous != devices.digitals[d][a].pressed
}

// Started returns true if the action has just been started this very frame
// (i.e. just been pressed by the player).
//
// Note: this *must* be queried in the React method of the game loop, as this is
// the only method that is guaranteed to run at least once each frame.
func (a DigitalID) Started(d DeviceID) bool {
	return devices.digitals[d][a].pressed && !devices.digitals[d][a].previous
}

// Stopped returns true if the action has just been stopped this very frame
// (i.e. just been released by the player).
//
// Note: this *must* be queried in the React method of the game loop, as this is
// the only method that is guaranteed to run at least once each frame.
func (a DigitalID) Stopped(d DeviceID) bool {
	return devices.digitals[d][a].previous && !devices.digitals[d][a].pressed
}

////////////////////////////////////////////////////////////////////////////////

func (a DigitalID) activate(d DeviceID, b source) {
	devices.digitals[d][a].active = true
	devices.digitalbinds[d][a] = append(devices.digitalbinds[d][a], b)
	_, v := b.asBool()
	if v {
		devices.digitals[d][a].pressed = true
		devices.digitals[d][a].previous = true
		devices.digitals[0][a].pressed = true
		devices.digitals[0][a].previous = true
	}
}

func (a DigitalID) newframe(d DeviceID) {
	devices.digitals[d][a].previous = devices.digitals[d][a].pressed
}

func (a DigitalID) update(d DeviceID) {
	for _, b := range devices.digitalbinds[d][a] {
		j, v := b.asBool()
		if j {
			devices.digitals[d][a].pressed = v
			devices.digitals[0][a].pressed = v
		}
	}
}

func (a DigitalID) deactivate(d DeviceID) {
	devices.digitalbinds[d][a] = devices.digitalbinds[d][a][:0]
	devices.digitals[d][a].active = false
	devices.digitals[d][a].pressed = false
	devices.digitals[0][a].active = false
	devices.digitals[0][a].pressed = false
}
