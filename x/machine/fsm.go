// Copyright (c) 2013-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package machine

////////////////////////////////////////////////////////////////////////////////

// State is used to implement a lightweight state machine. Each state is
// represented by a closure, that, when run, returns the next state (or nil to
// stop the machine).
type State func() State

////////////////////////////////////////////////////////////////////////////////

// Update runs s (i.e. the current state), and replaces it with the result of
// that call. It's safe to call Update on a nil target.
func (s *State) Update() {
	if *s != nil {
		*s = (*s)()
	}
}

////////////////////////////////////////////////////////////////////////////////
