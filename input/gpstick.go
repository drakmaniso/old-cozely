// Copyright (c) 2018-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package input

type gpStick struct{}

func (a gpStick) bind(c ContextID, target Action) {}
func (a gpStick) activate(d DeviceID)             {}
func (a gpStick) asBool() (just bool, value bool) { return false, false }
func (a gpStick) asFloat() (just bool, value float32) { return false, 0}
