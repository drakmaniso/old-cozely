// Copyright (c) 2018-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package input

import (
	"errors"

	"github.com/cozely/cozely/internal"
)

////////////////////////////////////////////////////////////////////////////////

// ContextID identifies an input context, i.e. a set of actions that can be
// active at the same time. The same action can be listed in several contexts,
// and have different (or identical) bindings in each.
type ContextID uint32

const noContext = ContextID(maxID)

var current, new = ContextID(0), ContextID(0)

var contexts struct {
	// For each context
	name []string

	// For each context, a list of actions
	actions [][]Action
}

////////////////////////////////////////////////////////////////////////////////

// Context declares a new input context. Its name should be unique.
func Context(name string, list ...Action) ContextID {
	if internal.Running {
		setErr(errors.New("input context declaration: declarations must happen before starting the framework"))
		return ContextID(maxID)
	}

	l := len(contexts.name)
	if l >= maxID {
		setErr(errors.New("input context declaration: too many contexts"))
		return ContextID(maxID)
	}

	c := ContextID(l)
	contexts.name = append(contexts.name, name)
	contexts.actions = append(contexts.actions, list)

	return c
}

////////////////////////////////////////////////////////////////////////////////

// Activate makes the context active on all devices.
func (a ContextID) Activate() {
	for d := range devices.name {
		devices.newcontext[d] = a
	}
}

// ActivateOn makes the context active on a specific device.
func (a ContextID) ActivateOn(d DeviceID) {
	devices.newcontext[d] = a
}

// Active returns true if the context is currently active for the current
// device.
func (a ContextID) Active() bool {
	return a.ActiveOn(devices.current)
}

// ActiveOn returns true if the context is currently active on a specific device.
func (a ContextID) ActiveOn(d DeviceID) bool {
	return a == devices.context[d]
}
