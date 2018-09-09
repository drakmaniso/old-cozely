// Copyright (c) 2018-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package pixel

import (
	"github.com/cozely/cozely/color"
)

////////////////////////////////////////////////////////////////////////////////

// Transparent is the only reserved color index. All palettes start with it.
const Transparent = color.Index(0)

////////////////////////////////////////////////////////////////////////////////

// LRGBA returns the color corresponding to a color index.
func LRGBA(c color.Index) color.LRGBA {
	return palette.colors[c]
}
