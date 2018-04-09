// Copyright (c) 2018-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package action

import "github.com/drakmaniso/glam/internal"

type Bool uint32

const noBool = Bool(maxID)

func NewBool(name string) Bool {
	_, ok := actions[name]
	if ok {
		//TODO: set error
		return noBool
	}

	l := len(internal.Bools.Name)
	if l >= maxID {
		//TODO: set error
		return noBool
	}

	actions[name] = Bool(l)
	internal.Bools.Name = append(internal.Bools.Name, name)
	internal.Bools.Active = append(internal.Bools.Active, false)
	internal.Bools.Just = append(internal.Bools.Just, false)
	internal.Bools.Pressed = append(internal.Bools.Pressed, false)

	return Bool(l)
}

func (b Bool) Name() string {
	return internal.Bools.Name[b]
}

func (b Bool) Active() bool {
	return internal.Bools.Active[b]
}

func (b Bool) Pressed() bool {
	return internal.Bools.Pressed[b]
}

func (b Bool) JustPressed() bool {
	return internal.Bools.Just[b] && internal.Bools.Pressed[b]
}

func (b Bool) Released() bool {
	return !internal.Bools.Pressed[b]
}

func (b Bool) JustReleased() bool {
	return internal.Bools.Just[b] && !internal.Bools.Pressed[b]
}
