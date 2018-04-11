// Copyright (c) 2018-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package input

//------------------------------------------------------------------------------

type Context uint32

const noContext = Context(maxID)

var current, new = Context(0), Context(0)

var contexts struct {
	// For each context
	name []string

	// For each context, a list of actions
	actions [][]Action
}

//------------------------------------------------------------------------------

func NewContext(name string, la ...Action) Context {
	l := len(contexts.name)
	if l >= maxID {
		//TODO: set error
		return Context(maxID)
	}

	c := Context(l)
	contexts.name = append(contexts.name, name)
	contexts.actions = append(contexts.actions, la)

	return c
}

//------------------------------------------------------------------------------

func (a Context) Activate(d Device) {
	devices.newcontext[d] = a
}

func (a Context) Active(d Device) bool {
	return a == devices.context[d]
}

//------------------------------------------------------------------------------
