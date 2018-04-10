// Copyright (c) 2018-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package input

import (
	"github.com/drakmaniso/glam/internal"
	"errors"
)

type binding interface {
	bind(c Context, a Action)
	device() Device
	action() Action
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
			act, ok := actions[an]
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
	println("NumJoysticks=", n)
	for j := 0; j < n; j++ {
		if internal.IsGameController(j) {
			println("joystick ", j, " is a gamepad")
			c := internal.GameControllerOpen(j)
			if c == nil {
				return errors.New("unable to open joystick as gamepad")
			}
		}
	}

	return err
}
