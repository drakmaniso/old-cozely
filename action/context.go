// Copyright (c) 2018-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package action

type Context uint32

const noContext = Context(maxID)

var contexts struct {
	Name   []string
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

	contexts.Name = append(contexts.Name, name)
	contexts.Bools = append(contexts.Bools, make([]Bool, 0, 8))
	contexts.Floats = append(contexts.Floats, make([]Float, 0, 8))
	contexts.Coords = append(contexts.Coords, make([]Coord, 0, 8))
	contexts.Deltas = append(contexts.Deltas, make([]Delta, 0, 8))

	return Context(l)
}

func (c Context) Activate() {

}
