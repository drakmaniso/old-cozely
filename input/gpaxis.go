// Copyright (c) 2018-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package input

import (
	"github.com/cozely/cozely/internal"
)

type gpAxis struct {
	gamepad *internal.Gamepad
	axis    internal.GamepadAxis
	target  Action
	value   float32
}

func (a *gpAxis) bind(c ContextID, target Action) {}
func (a *gpAxis) activate(d DeviceID)             {}
func (a *gpAxis) asBool() (just bool, value bool) { return false, false }
func (a *gpAxis) asFloat() (just bool, value float32) { return false, 0 }
