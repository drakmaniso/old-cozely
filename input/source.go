// Copyright (c) 2018-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package input

import (
	"github.com/cozely/cozely/coord"
)

////////////////////////////////////////////////////////////////////////////////

type source interface {
	bind(c ContextID, a Action)
	activate(d DeviceID)
	asBool() (just bool, value bool)
	asUnipolar() (just bool, value float32)
	asBipolar() (just bool, value float32)
	asCoord() (just bool, value coord.XY)
	asDelta() coord.XY
}
