// Copyright (c) 2018-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package input

import (
	"github.com/cozely/cozely/internal"
)

////////////////////////////////////////////////////////////////////////////////

func init() {
	internal.InputSetup = setup
}

func setup() error {
	internal.Log.Printf("Input declarations: %d contexts and %d actions",
		len(contexts.name), len(actions.name))

	if len(contexts.name) == 0 {
		aa := []Action{}
		for _, a := range actions.name {
			aa = append(aa, a)
		}
		c := Context("Default", aa...)
		c.ActivateOn(Any)
		internal.Log.Printf("(added default context with all actions)")
	}

	load()

	return nil
}
