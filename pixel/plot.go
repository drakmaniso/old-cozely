// Copyright (c) 2013-2017 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package pixel

import (
	"github.com/drakmaniso/glam/colour"
)

//------------------------------------------------------------------------------

func Point(x, y int16, c colour.Colour) {
	commandPoint(x, y, c)
}

//------------------------------------------------------------------------------
