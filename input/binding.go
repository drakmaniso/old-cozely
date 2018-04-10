// Copyright (c) 2018-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package input

import "errors"

type binding interface {
	bind(c Context, a Action)
	device() Device
	action() Action
	asBool() (just bool, value bool)
}

type gpStick struct{}

func (a gpStick) bind(c Context, target Action)   {}
func (a gpStick) device() Device                  { return noDevice }
func (a gpStick) action() Action                  { return nil }
func (a gpStick) asBool() (just bool, value bool) { return false, false }

type gpTrigger struct{}

func (a gpTrigger) bind(c Context, target Action)   {}
func (a gpTrigger) device() Device                  { return noDevice }
func (a gpTrigger) action() Action                  { return nil }
func (a gpTrigger) asBool() (just bool, value bool) { return false, false }

type gpButton struct{}

func (a gpButton) bind(c Context, target Action)   {}
func (a gpButton) device() Device                  { return noDevice }
func (a gpButton) action() Action                  { return nil }
func (a gpButton) asBool() (just bool, value bool) { return false, false }

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

	return err
}
