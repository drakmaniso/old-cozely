// Copyright (c) 2018-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package input

type gpTrigger struct{}

func (a gpTrigger) bind(c ContextID, target Action) {}
func (a gpTrigger) activate(d DeviceID)             {}
func (a gpTrigger) asBool() (just bool, value bool) { return false, false }
