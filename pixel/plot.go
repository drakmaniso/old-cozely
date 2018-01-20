// Copyright (c) 2013-2017 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package pixel

import (
	"github.com/drakmaniso/glam/colour"
)

//------------------------------------------------------------------------------

func Point(c colour.Colour, x, y int16) {
	appendCommand(cmdPoint, 3)
	c8 := colour.SRGBA8Of(c)
	rg := uint16(c8.R)<<8 | uint16(c8.G)
	ba := uint16(c8.B)<<8 | uint16(c8.A)
	parameters = append(parameters, int16(rg), int16(ba), x, y)
}

//------------------------------------------------------------------------------
