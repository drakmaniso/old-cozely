// Copyright (c) 2018-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package action

import "github.com/drakmaniso/glam/internal"

type Context uint32

const noContext = Context(maxID)

var current, new = Context(0), Context(0)

var contexts struct {
	Name []string

	Bools  [][]Bool
	Floats [][]Float
	Coords [][]Coord
	Deltas [][]Delta

	KeyboardHooks [][]func()
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

	contexts.KeyboardHooks = append(contexts.KeyboardHooks, make([]func(), 0, 8))

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

func (c Context) Activate() {
	new = c
}

func  activate() {
	for _, b := range contexts.Bools[current] {
		internal.Bools.Active[b] = false
		internal.Bools.Just[b] = internal.Bools.Pressed[b]
		internal.Bools.Pressed[b] = false
	}
	for _, f := range contexts.Floats[current] {
		internal.Floats.Active[f] = false
	}
	for _, c := range contexts.Coords[current] {
		internal.Coords.Active[c] = false
	}
	for _, d := range contexts.Deltas[current] {
		internal.Deltas.Active[d] = false
	}

	current = new
	for _, b := range contexts.Bools[new] {
		internal.Bools.Active[b] = true
	}
	for _, f := range contexts.Floats[new] {
		internal.Floats.Active[f] = true
	}
	for _, c := range contexts.Coords[new] {
		internal.Coords.Active[c] = true
	}
	for _, d := range contexts.Deltas[new] {
		internal.Deltas.Active[d] = true
	}
}

func (c Context) Active() bool {
	return c == current
}

func init() {
	internal.ActionPrepare = prepare
	internal.ActionAmend = amend
}

func prepare() error {
	if current != new {
		activate()
	}
	for _, h := range contexts.KeyboardHooks[current] {
		h()
	}
	return nil
}

func amend() error {
	for i := range internal.Bools.Just {
		internal.Bools.Just[i] = false
	}
	return nil
}
