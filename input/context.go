// Copyright (c) 2018-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package input

////////////////////////////////////////////////////////////////////////////////

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

func Context(name string, la ...Action) ContextID {
	l := len(contexts.name)
	if l >= maxID {
		//TODO: set error
		return ContextID(maxID)
	}

	c := ContextID(l)
	contexts.name = append(contexts.name, name)
	contexts.actions = append(contexts.actions, la)

	return c
}

////////////////////////////////////////////////////////////////////////////////

func (a ContextID) Activate(d DeviceID) {
	devices.newcontext[d] = a
}

func (a ContextID) Active(d DeviceID) bool {
	return a == devices.context[d]
}

////////////////////////////////////////////////////////////////////////////////
