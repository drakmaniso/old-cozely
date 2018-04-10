// Copyright (c) 2018-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package input

import "errors"

type binding interface {
	bind(c Context, a Action)
	device() Device
	asBool() (just bool, value bool)
}

type gpStick struct{}

func (g gpStick) bind(c Context, a Action) {}
func (g gpStick) device() Device {return noDevice}
func (g gpStick) asBool() (just bool, value bool) {return false, false}

type gpTrigger struct{}

func (g gpTrigger) bind(c Context, a Action) {}
func (g gpTrigger) device() Device {return noDevice}
func (g gpTrigger) asBool() (just bool, value bool) {return false, false}

type gpButton struct{}

func (g gpButton) bind(c Context, a Action) {}
func (g gpButton) device() Device {return noDevice}
func (g gpButton) asBool() (just bool, value bool) {return false, false}

func LoadBindings(b map[string]map[string][]string) error {
	var err error

	// Forget previous bindings
	var nbctx = len(contexts.Name)
	keybmouse.keys = make([][]*kbKey, nbctx)
	keybmouse.buttons = make([][]*msButton, nbctx)

	for cn, cb := range b {
		// Find context by name
		ctx := noContext
		for i, n := range contexts.Name {
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

	return err
}
