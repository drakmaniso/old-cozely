// Copyright (c) 2018-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package input

type gpStick struct{}

func (a gpStick) bind(c Context, target Action)   {}
func (a gpStick) device() Device                  { return noDevice }
func (a gpStick) action() Action                  { return nil }
func (a gpStick) asBool() (just bool, value bool) { return false, false }
