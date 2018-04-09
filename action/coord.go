// Copyright (c) 2018-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package action

import (
	"github.com/drakmaniso/glam/internal"
	"github.com/drakmaniso/glam/plane"
)

type Coord uint32

const noCoord = Coord(maxID)

func NewCoord(name string) Coord {
	_, ok := actions[name]
	if ok {
		//TODO: set error
		return noCoord
	}

	l := len(internal.Coords.Name)
	if l >= maxID {
		//TODO: set error
		return noCoord
	}

	actions[name] = Coord(l)
	internal.Coords.Name = append(internal.Coords.Name, name)
	internal.Coords.Active = append(internal.Coords.Active, false)
	internal.Coords.X = append(internal.Coords.X, 0)
	internal.Coords.Y = append(internal.Coords.Y, 0)

	return Coord(l)
}

func (c Coord) Name() string {
	return internal.Bools.Name[c]
}

func (c Coord) Active() bool {
	return internal.Coords.Active[c]
}

func (c Coord) Coord() plane.Coord {
	return plane.Coord{internal.Coords.X[c], internal.Coords.Y[c]}
}
