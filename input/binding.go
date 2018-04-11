// Copyright (c) 2018-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package input

import (
	"errors"
	"github.com/drakmaniso/glam/internal"
)

type binding interface {
	bind(c Context, a Action)
	activate(d Device)
	asBool() (just bool, value bool)
}

func LoadBindings(b map[string]map[string][]string) error {
	var err error

	// Forget devices (and previous bindings)
	clearDevices()

	for cn, cb := range b {
		// Find context by name
		ctx := noContext
		for i, n := range contexts.name {
			if n == cn {
				ctx = Context(i)
				break
			}
		}
		if ctx == noContext {
			if err == nil {
				err = errors.New("unkown context: " + cn)
			}
			continue
		}
		println(cn)

		for an, ab := range cb {
			// Find action by name
			act, ok := actions.names[an]
			if !ok {
				if err == nil {
					err = errors.New("unkown action: " + an)
				}
				continue
			}
			print("    ", an, " = ")

			for _, n := range ab {
				bnd, ok := binders[n]
				if !ok {
					if err == nil {
						err = errors.New("unkown binding: " + n)
					}
					continue
				}
				bnd.bind(ctx, act)
				print(", ")
			}
			println("")

		}

	}

	// Add gamepad devices

	n := internal.NumJoysticks()
	internal.Debug.Printf("Number of controllers detected: %d", n)
	for j := 0; j < n; j++ {
		if internal.IsGameController(j) {
			c := internal.GameControllerOpen(j)
			if c == nil {
				return errors.New("unable to open joystick as gamepad")
			}
			// nm := c.Name()
			nm := internal.JoystickNameForIndex(j)
			internal.Debug.Printf("#%d: %s is a gamepad", j, nm)
		} else {
			nm := internal.JoystickNameForIndex(j)
			internal.Debug.Printf("#%d: %s is a joystick", j, nm)
		}
	}

	return err
}
