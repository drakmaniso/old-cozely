// Copyright (c) 2018-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package input

import (
	"github.com/cozely/cozely/internal"
)

////////////////////////////////////////////////////////////////////////////////

func init() {
	internal.InputNewFrame = newframe
}

func newframe() error {
	updateMouse()

	for _, t := range actions.list {
		for d := range devices.name {
			t.newframe(DeviceID(d))
		}
	}

	for d := range devices.name {
		if d == 0 {
			// Pseudo-device 0 has no context and no binds
			continue
		}

		c := devices.context[d]
		// Activate this device context if necessary
		if c != devices.newcontext[d] {
			if c != noContext {
				for _, t := range contexts.actions[c] {
					t.deactivate(DeviceID(d))
				}
			}
			c = devices.newcontext[d]
			devices.context[d] = c
			for _, b := range devices.bindings[d][c] {
				b.activate(DeviceID(d))
			}
		}

		for _, t := range contexts.actions[c] {
			t.update(DeviceID(d))
		}
	}

	return nil
}

////////////////////////////////////////////////////////////////////////////////
