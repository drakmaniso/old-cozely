// Copyright (c) 2018-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package input

import (
	"github.com/drakmaniso/glam/internal"
)

type Context uint32

const noContext = Context(maxID)

var current, new = Context(0), Context(0)

var contexts struct {
	name []string
}

func NewContext(name string, actions ...Action) Context {
	l := len(contexts.name)
	if l >= maxID {
		//TODO: set error
		return Context(maxID)
	}

	c := Context(l)
	contexts.name = append(contexts.name, name)

	return c
}

func (c Context) Activate(d Device) {
	devices.new[d] = c
}

func (c Context) Active(d Device) bool {
	return c == devices.context[d]
}

func init() {
	internal.ActionPrepare = prepare
}

func prepare() error {
	for d := range devices.name {
		if devices.context[d] != devices.new[d] {
			for _, b := range devices.bindings[d][devices.context[d]] {
				b.action().deactivate(KeyboardAndMouse)
			}

			devices.context[d] = devices.new[d]

			for _, b := range devices.bindings[d][devices.context[d]] {
				b.action().activate(b)
			}
		}

		for _, b := range devices.bindings[d][devices.context[d]] {
			b.action().newframe(b)
		}

		for _, b := range devices.bindings[d][devices.context[d]] {
			b.action().prepare(b)
		}

	}

	return nil
}
