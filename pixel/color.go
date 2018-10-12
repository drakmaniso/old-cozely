// Copyright (c) 2018-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package pixel

import (
	"github.com/cozely/cozely/color"
)

////////////////////////////////////////////////////////////////////////////////

// Reserved color indices.
const (
	Transparent = color.Index(0)
	Black = color.Index(251)
	DarkGray = color.Index(252)
	MidGray = color.Index(253)
	LightGray = color.Index(254)
	White = color.Index(255)
)

////////////////////////////////////////////////////////////////////////////////

// LRGBAof returns the color corresponding to a color index in the current
// palette.
func LRGBAof(c color.Index) color.LRGBA {
	return palette.colors[c]
}
