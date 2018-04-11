// Copyright (c) 2018-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package input

type gpStick struct{}

func (a gpStick) bind(c Context, target Action)   {}
func (a gpStick) activate(d Device)              {}
func (a gpStick) asBool() (just bool, value bool) { return false, false }
