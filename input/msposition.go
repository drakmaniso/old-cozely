// Copyright (c) 2018-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package input

import "github.com/cozely/cozely/coord"

type msPosition struct {
}

func (a msPosition) bind(c ContextID, target Action)        {}
func (a msPosition) activate(d DeviceID)                    {}
func (a msPosition) asBool() (just bool, value bool)        { return false, false }
func (a msPosition) asUnipolar() (just bool, value float32) { return false, 0 }
func (a msPosition) asBipolar() (just bool, value float32)  { return false, 0 }
func (a msPosition) asCoord() (just bool, value coord.XY)   { return false, coord.XY{0, 0} }
