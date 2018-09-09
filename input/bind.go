// Copyright (c) 2018-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package input

import (
	"errors"

	"github.com/cozely/cozely/internal"
)

////////////////////////////////////////////////////////////////////////////////

// Bindings is a list of bindings for each context/action combination. The first
// map level associates each context name to a sub-map; this sub map then
// associates each action name to a slice of strings (the actual bindings).
type Bindings map[string]map[string][]string

var bindings = Bindings{}

////////////////////////////////////////////////////////////////////////////////

//TODO: auto-load bindings

// Load associates each context/action combination found in the bindings map to
// the requested bindings.
func Load(b Bindings) {
	bindings = b
	if internal.Running {
		load()
	}
}

func load() {
	// Forget devices (and previous bindings)
	clearDevices()
	// Add gamepad devices
	scanJoysticks()

	lcn := "Loaded input bindings (contexts:"

	for cn, cb := range bindings {
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
			act, ok := actions.name[an]
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
	}
	internal.Debug.Printf(lcn + ")")
}
