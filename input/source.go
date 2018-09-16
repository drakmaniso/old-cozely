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
	asButton() (just bool, value bool)
	asHalfAxis() (just bool, value float32)
	asAxis() (just bool, value float32)
	asDualAxis() (just bool, value coord.XY)
	asDelta() (just bool, value coord.XY)
}
