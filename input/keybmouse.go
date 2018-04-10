// Copyright (c) 2018-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package input

var keybmouse struct {
	context, new Context
	keys         [][]kbKey
	buttons      [][]msButton
}

type kbKey struct {
	keycode KeyCode
	action  Action
}

func (k kbKey) bind(c Context, a Action) {
	for len(keybmouse.keys) < int(c+1) {
		keybmouse.keys = append(keybmouse.keys, []kbKey{})
	}
	k.action = a
	keybmouse.keys[c] = append(keybmouse.keys[c], k)
	// switch a := a.(type) {
	// case Bool:
	// 	print("keyboard", "->bool", a)
	// case Float:
	// 	print("keyboard", "->float", a)
	// case Coord:
	// 	print("keyboard", "->coord", a)
	// case Delta:
	// 	print("keyboard", "->delta", a)
	// }
}

type msPosition struct {
}

func (m msPosition) bind(c Context, a Action) {

}

type msButton struct {
	button MouseButton
	action Action
}

func (m msButton) bind(c Context, a Action) {

}
