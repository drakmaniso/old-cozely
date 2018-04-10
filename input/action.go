// Copyright (c) 2018-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package input

import (
	"errors"
)

type Action interface {
	Active() bool
	activate()
	deactivate()
	prepareKey(k KeyCode)
}

var actions = map[string]Action{}

const (
	flagActive byte = 1 << iota
)

const maxID = 0xFFFFFFFF

func LoadBindings(b map[string]map[string][]string) error {
	var err error

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
				bnd, ok := bindings[n]
				if !ok {
					if err == nil {
						err = errors.New("unkown binding: " + n)
					}
					continue
				}
				bnd.BindTo(ctx, act)
				print(", ")
			}
			println("")

		}

	}

	return err
}
