// Copyright (c) 2018-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package input

type gpButton struct{}

func (a gpButton) bind(c Context, target Action)   {}
func (a gpButton) device() Device                  { return noDevice }
func (a gpButton) action() Action                  { return nil }
func (a gpButton) asBool() (just bool, value bool) { return false, false }
