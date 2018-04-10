// Copyright (c) 2018-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package input

import (
	"github.com/drakmaniso/glam/internal"
)

var keybmouse struct {
	context, new Context
	keys         [][]*kbKey
	buttons      [][]*msButton
}

type kbKey struct {
	keycode       KeyCode
	action        Action
	just, pressed bool
}

func (a *kbKey) bind(c Context, target Action) {
	aa := *a
	aa.action = target
	keybmouse.keys[c] = append(keybmouse.keys[c], &aa)
}

func (a *kbKey) device() Device {
	return KeyboardAndMouse
}

func (a *kbKey) asBool() (just bool, value bool) {
	v := internal.Key(a.keycode)
	a.just = (v != a.pressed) //TODO: no need to store?
	a.pressed = v
	return a.just, a.pressed
}

type msPosition struct {
}

func (a msPosition) bind(c Context, target Action) {}
func (a msPosition) device() Device {return noDevice}
func (a msPosition) asBool() (just bool, value bool) {
	return false, false
}

type msButton struct {
	button MouseButton
	action Action
	just, pressed bool
}

// MouseButton identifies a mouse button
type MouseButton uint32

// MouseButton constants
const (
	MouseLeft MouseButton = 1 << iota
	MouseMiddle
	MouseRight
	MouseBack
	MouseForward
	Mouse6
	Mouse7
	Mouse8
)

func (a *msButton) bind(c Context, target Action) {
	aa := *a
	aa.action = target
	keybmouse.buttons[c] = append(keybmouse.buttons[c], &aa)
}

func (a *msButton) device() Device {
	return KeyboardAndMouse
}

func (a *msButton) asBool() (just bool, value bool) {
	v := (MouseButton(internal.MouseButtons) & a.button) != 0
	a.just = (v != a.pressed) //TODO: no need to store?
	a.pressed = v
	return a.just, a.pressed
}
