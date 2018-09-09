// Copyright (c) 2018-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package input

import (
	"errors"

	"github.com/cozely/cozely/internal"
)

////////////////////////////////////////////////////////////////////////////////

// ButtonID identifies a digital Action, i.e. an action that can be either "on"
// or "off".
//
// In this package, the state of being "on" is called "pressed". The transition
// from "off" to "on" is called "pushed"; and the transition back to "off" is
// called "released".
type ButtonID uint32

const noButton = ButtonID(maxID)

var buttons = struct {
	// For each button
	name []string
}{
	name: []string{
		"Menu Select",
		"Menu Back",
		"Menu Up",
		"Menu Down",
		"Menu Left",
		"Menu Right",
		"Menu Click",
	},
}

type button struct {
	// active   bool
	previous bool
	pressed  bool
}

////////////////////////////////////////////////////////////////////////////////

// Button declares a new digital action, and returns its ID.
func Button(name string) ButtonID {
	if internal.Running {
		setErr(errors.New("input button declaration: declarations must happen before starting the framework"))
		return noButton
	}

	_, ok := actions.name[name]
	if ok {
		setErr(errors.New("input button declaration: name already taken by another action"))
		return noButton
	}

	a := len(buttons.name)
	if a >= maxID {
		setErr(errors.New("input button declaration: too many button actions"))
		return noButton
	}

	actions.name[name] = ButtonID(a)
	actions.list = append(actions.list, ButtonID(a))
	buttons.name = append(buttons.name, name)

	return ButtonID(a)
}

////////////////////////////////////////////////////////////////////////////////

// Name of the button action.
func (a ButtonID) Name() string {
	return buttons.name[a]
}

// Pressed returns true if the action has been started and is currently ongoing
// (i.e. "on", "pressed") on the current device.
func (a ButtonID) Pressed() bool {
	return a.PressedOn(devices.current)
}

// PressedOn returns true if the action has been started and is currently ongoing
// (i.e. "on", "pressed") on a specific device.
func (a ButtonID) PressedOn(d DeviceID) bool {
	return devices.buttons[d][a].pressed
}

// Changed returns true if the action has just been pressed or released this
// very frame (i.e. just been started or stopped by the player).
//
// Note: this *must* be queried in the React method of the game loop, as this is
// the only method that is guaranteed to run at least once each frame.
func (a ButtonID) Changed() bool {
	return a.ChangedOn(devices.current)
}

// ChangedOn returns true if the action has just been pressed or released this
// very frame (i.e. just been started or stopped by the player).
//
// Note: this *must* be queried in the React method of the game loop, as this is
// the only method that is guaranteed to run at least once each frame.
func (a ButtonID) ChangedOn(d DeviceID) bool {
	return devices.buttons[d][a].previous != devices.buttons[d][a].pressed
}

// Pushed returns true if the action has just been started this very frame (i.e.
// just been pressed by the player).
//
// Note: this *must* be queried in the React method of the game loop, as this is
// the only method that is guaranteed to run at least once each frame.
func (a ButtonID) Pushed() bool {
	return a.PushedOn(devices.current)
}

// PushedOn returns true if the action has just been started this very frame
// (i.e. just been pressed by the player).
//
// Note: this *must* be queried in the React method of the game loop, as this is
// the only method that is guaranteed to run at least once each frame.
func (a ButtonID) PushedOn(d DeviceID) bool {
	return devices.buttons[d][a].pressed && !devices.buttons[d][a].previous
}

// Released returns true if the action has just been stopped this very frame
// (i.e. just been released by the player).
//
// Note: this *must* be queried in the React method of the game loop, as this is
// the only method that is guaranteed to run at least once each frame.
func (a ButtonID) Released() bool {
	return a.ReleasedOn(devices.current)
}

// ReleasedOn returns true if the action has just been stopped this very frame
// (i.e. just been released by the player).
//
// Note: this *must* be queried in the React method of the game loop, as this is
// the only method that is guaranteed to run at least once each frame.
func (a ButtonID) ReleasedOn(d DeviceID) bool {
	return devices.buttons[d][a].previous && !devices.buttons[d][a].pressed
}

////////////////////////////////////////////////////////////////////////////////

func (a ButtonID) activate(d DeviceID, b source) {
	devices.buttonsbinds[d][a] = append(devices.buttonsbinds[d][a], b)
	_, v := b.asButton()
	if v {
		devices.buttons[d][a].pressed = true
		devices.buttons[d][a].previous = true
	}
}

func (a ButtonID) newframe(d DeviceID) {
	devices.buttons[d][a].previous = devices.buttons[d][a].pressed
}

func (a ButtonID) update(d DeviceID) {
	for _, s := range devices.buttonsbinds[d][a] {
		j, v := s.asButton()
		if j {
			devices.buttons[d][a].pressed = v
			devices.current = d
		}
	}
}

func (a ButtonID) deactivate(d DeviceID) {
	devices.buttonsbinds[d][a] = devices.buttonsbinds[d][a][:0]
	devices.buttons[d][a].pressed = false
}
