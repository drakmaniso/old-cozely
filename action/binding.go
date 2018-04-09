// Copyright (c) 2018-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package action

type binding interface {
	BindTo(c Context, a Action)
	Unbind()
}

type gamepadStick struct {
}

func (g gamepadStick) BindTo(c Context, a Action)     {}
func (g gamepadStick) Unbind() {}

type gamepadTrigger struct {
}

func (g gamepadTrigger) BindTo(c Context, a Action)     {}
func (g gamepadTrigger) Unbind() {}

type gamepadButton struct {
}

func (g gamepadButton) BindTo(c Context, a Action)     {}
func (g gamepadButton) Unbind() {}

type mouse struct {
}

func (m mouse) BindTo(c Context, a Action)     {}
func (m mouse) Unbind() {}

type mouseButton struct {
}

func (m mouseButton) BindTo(c Context, a Action)     {}
func (m mouseButton) Unbind() {}

type keyboard struct {
}

func (k keyboard) BindTo(c Context, a Action) {
	switch a := a.(type) {
	case Bool:
		print("keyboard", "->bool", a)
	case Float:
		print("keyboard", "->float", a)
	case Coord:
		print("keyboard", "->coord", a)
	case Delta:
		print("keyboard", "->delta", a)
	}
}
func (k keyboard) Unbind() {}
