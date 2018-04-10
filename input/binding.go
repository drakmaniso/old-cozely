// Copyright (c) 2018-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package input

type binding interface {
	bind(c Context, a Action)
}

type gpStick struct {
}

func (g gpStick) bind(c Context, a Action) {
}

type gpTrigger struct {
}

func (g gpTrigger) bind(c Context, a Action) {
}

type gpButton struct {
}

func (g gpButton) bind(c Context, a Action) {
}
