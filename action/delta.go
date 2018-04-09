// Copyright (c) 2018-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package action

import (
	"github.com/drakmaniso/glam/internal"
	"github.com/drakmaniso/glam/plane"
)

type Delta uint32

func NewDelta(name string) Delta {
	l := len(internal.Deltas.Name)
	if l >= maxID {
		//TODO: set error
		return Delta(maxID)
	}

	internal.Deltas.Name = append(internal.Deltas.Name, name)
	internal.Deltas.Active = append(internal.Deltas.Active, false)
	internal.Deltas.X = append(internal.Deltas.X, 0)
	internal.Deltas.Y = append(internal.Deltas.Y, 0)

	return Delta(l)
}

func (c Delta) Name() string {
	return internal.Bools.Name[c]
}

func (c Delta) Active() bool {
	return internal.Deltas.Active[c]
}

func (c Delta) Delta() plane.Coord {
	return plane.Coord{internal.Deltas.X[c], internal.Deltas.Y[c]}
}
