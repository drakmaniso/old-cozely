// Copyright (c) 2018-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package input

type binding interface {
	BindTo(c Context, a Action)
	Unbind()
}

type gpStick struct {
}

func (g gpStick) BindTo(c Context, a Action) {}
func (g gpStick) Unbind()                    {}

type gpTrigger struct {
}

func (g gpTrigger) BindTo(c Context, a Action) {}
func (g gpTrigger) Unbind()                    {}

type gpButton struct {
}

func (g gpButton) BindTo(c Context, a Action) {}
func (g gpButton) Unbind()                    {}

type msPosition struct {
}

func (m msPosition) BindTo(c Context, a Action) {}
func (m msPosition) Unbind()                    {}

type msButton struct {
}

func (m msButton) BindTo(c Context, a Action) {}
func (m msButton) Unbind()                    {}

type kbKey struct {
	pos KeyCode
}

func (k kbKey) BindTo(c Context, a Action) {
	switch a := a.(type) {
	case Bool:
		for len(keyboard.actions) < int(c+1) {
			keyboard.actions = append(keyboard.actions, []keyAction{})
		}
		keyboard.actions[c] = append(keyboard.actions[c],
			keyAction{
				position: k.pos,
				action:   a,
			})
		print("keyboard", "->bool", a)
	case Float:
		print("keyboard", "->float", a)
	case Coord:
		print("keyboard", "->coord", a)
	case Delta:
		print("keyboard", "->delta", a)
	}
}
func (k kbKey) Unbind() {}
