// Copyright (c) 2018-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package action

type binding interface {
	BindTo(a Action)
	UnbindFrom(a Action)
}

type gamepadStick struct {
}

func (g gamepadStick) BindTo(a Action)     {}
func (g gamepadStick) UnbindFrom(a Action) {}

type gamepadTrigger struct {
}

func (g gamepadTrigger) BindTo(a Action)     {}
func (g gamepadTrigger) UnbindFrom(a Action) {}

type gamepadButton struct {
}

func (g gamepadButton) BindTo(a Action)     {}
func (g gamepadButton) UnbindFrom(a Action) {}

type mouse struct {
}

func (g mouse) BindTo(a Action)     {}
func (g mouse) UnbindFrom(a Action) {}

type mouseFloat struct {
}

func (g mouseFloat) BindTo(a Action)     {}
func (g mouseFloat) UnbindFrom(a Action) {}

type mouseButton struct {
}

func (g mouseButton) BindTo(a Action)     {}
func (g mouseButton) UnbindFrom(a Action) {}

type keyboard struct {
}

func (g keyboard) BindTo(a Action)     {}
func (g keyboard) UnbindFrom(a Action) {}
