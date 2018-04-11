// Copyright (c) 2018-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package input

import (
	"github.com/drakmaniso/cozely/internal"
)

//------------------------------------------------------------------------------

func init() {
	internal.ActionNewFrame = newframe
}

func newframe() error {
	Cursor.newframe(1)

	for d := range devices.name {
		// Activate this device context if necessary
		c := devices.context[d]
		if c != devices.newcontext[d] {
			if c != noContext {
				for _, t := range contexts.actions[c] {
					t.deactivate(DeviceID(d))
				}
			}
			devices.context[d] = devices.newcontext[d]
			c = devices.context[d]
			for _, b := range devices.bindings[d][c] {
				b.activate(DeviceID(d))
			}
		}

		for _, t := range contexts.actions[c] {
			t.newframe(DeviceID(d))
		}
	}

	return nil
}

//------------------------------------------------------------------------------
