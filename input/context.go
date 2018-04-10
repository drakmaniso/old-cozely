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
	Name []string

	Bools  [][]Bool
	Floats [][]Float
	Coords [][]Coord
	Deltas [][]Delta
}

func NewContext(name string, actions ...Action) Context {
	l := len(contexts.Name)
	if l >= maxID {
		//TODO: set error
		return Context(maxID)
	}

	c := Context(l)
	contexts.Name = append(contexts.Name, name)

	contexts.Bools = append(contexts.Bools, make([]Bool, 0, 8))
	contexts.Floats = append(contexts.Floats, make([]Float, 0, 8))
	contexts.Coords = append(contexts.Coords, make([]Coord, 0, 8))
	contexts.Deltas = append(contexts.Deltas, make([]Delta, 0, 8))

	for _, a := range actions {
		switch a := a.(type) {
		case Bool:
			contexts.Bools[c] = append(contexts.Bools[c], a)
		case Float:
			contexts.Floats[c] = append(contexts.Floats[c], a)
		case Coord:
			contexts.Coords[c] = append(contexts.Coords[c], a)
		case Delta:
			contexts.Deltas[c] = append(contexts.Deltas[c], a)
		}
	}

	return c
}

func (c Context) Activate(d Device) {
	switch d {
	case KeyboardAndMouse:
		keybmouse.new = c
	}
}

func (c Context) Active(d Device) bool {
	switch d {
	case KeyboardAndMouse:
		return c == keybmouse.context
	}
	return false
}

func init() {
	internal.ActionPrepare = prepare
}

func prepare() error {
	// Keyboard
	if keybmouse.context != keybmouse.new {
		for _, b := range keybmouse.keys[keybmouse.context] {
			b.action.deactivate(KeyboardAndMouse)
		}
		keybmouse.context = keybmouse.new
		for _, b := range keybmouse.keys[keybmouse.context] {
			b.action.activateKey(b.keycode)
		}
	}
	for _, b := range keybmouse.keys[keybmouse.context] {
		b.action.prepareKey(b.keycode)
	}

	// Mouse

	return nil
}
