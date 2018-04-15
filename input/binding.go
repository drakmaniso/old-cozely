// Copyright (c) 2018-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package input

import (
	"errors"

	"github.com/cozely/cozely/internal"
)

////////////////////////////////////////////////////////////////////////////////

type binding interface {
	bind(c ContextID, a Action)
	activate(d DeviceID)
	asBool() (just bool, value bool)
}

// Bindings is a list of bindings for each context/action combination. The first
// map level associates each context name to a sub-map; this sub map then
// associates each action name to a slice of strings (the actual bindings).
type Bindings map[string]map[string][]string

////////////////////////////////////////////////////////////////////////////////

// Load associates each context/action combination found in the bindings map to
// the requested bindings.
func (a Bindings) Load() {
	if ! internal.Running {
		setErr(errors.New("bindings must be loaded while the framework is running"))
	}

	// Forget devices (and previous bindings)
	clearDevices()

	lcn := "Loaded input bindings (contexts:"

	for cn, cb := range a {
		// Find context by name
		ctx := noContext
		for i, n := range contexts.name {
			if n == cn {
				ctx = ContextID(i)
				break
			}
		}
		if ctx == noContext {
				setErr(errors.New("unknown context: " + cn))
			continue
		}
		lcn = lcn + " " + cn

		for an, ab := range cb {
			// Find action by name
			act, ok := actions.names[an]
			if !ok {
				setErr(errors.New("unknown action: " + an))
				continue
			}

			for _, n := range ab {
				bnd, ok := binders[n]
				if !ok {
					setErr(errors.New("unknown binding: " + n))
					continue
				}
				bnd.bind(ctx, act)
			}
		}
		internal.Debug.Printf(lcn + ")")
	}

	// Add gamepad devices

	n := internal.NumJoysticks()
	internal.Debug.Printf("Detected %d controllers:", n)
	for j := 0; j < n; j++ {
		if internal.IsGameController(j) {
			c := internal.GameControllerOpen(j)
			if c == nil {
				setErr(errors.New("unable to open joystick as gamepad"))
				continue
			}
			// nm := c.Name()
			nm := internal.JoystickNameForIndex(j)
			internal.Debug.Printf("Controller %d is a gamepad (%s)", j, nm)
		} else {
			nm := internal.JoystickNameForIndex(j)
			internal.Debug.Printf("Controller %d is a joystick (%s)", j, nm)
		}
	}
}
