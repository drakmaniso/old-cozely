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

type Bindings map[string]map[string][]string

////////////////////////////////////////////////////////////////////////////////

func Load(b Bindings) error {
	var err error

	// Forget devices (and previous bindings)
	clearDevices()

	lcn := "Loaded input bindings (contexts:"

	for cn, cb := range b {
		// Find context by name
		ctx := noContext
		for i, n := range contexts.name {
			if n == cn {
				ctx = ContextID(i)
				break
			}
		}
		if ctx == noContext {
			if err == nil {
				err = errors.New("unknown context: " + cn)
			}
			continue
		}
		lcn = lcn + " " + cn

		for an, ab := range cb {
			// Find action by name
			act, ok := actions.names[an]
			if !ok {
				if err == nil {
					err = errors.New("unknown action: " + an)
				}
				continue
			}

			for _, n := range ab {
				bnd, ok := binders[n]
				if !ok {
					if err == nil {
						err = errors.New("unknown binding: " + n)
					}
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
				return errors.New("unable to open joystick as gamepad")
			}
			// nm := c.Name()
			nm := internal.JoystickNameForIndex(j)
			internal.Debug.Printf("Controller %d is a gamepad (%s)", j, nm)
		} else {
			nm := internal.JoystickNameForIndex(j)
			internal.Debug.Printf("Controller %d is a joystick (%s)", j, nm)
		}
	}

	return err
}

////////////////////////////////////////////////////////////////////////////////
