// Copyright (c) 2018-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package input

type msPosition struct {
}

func (a msPosition) bind(c ContextID, target Action) {}
func (a msPosition) activate(d DeviceID)             {}
func (a msPosition) asBool() (just bool, value bool) { return false, false }
func (a msPosition) asFloat() (just bool, value float32) { return false, 0}
