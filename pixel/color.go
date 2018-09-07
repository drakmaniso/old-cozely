// Copyright (c) 2018-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package pixel

import (
	"github.com/cozely/cozely/color"
)

////////////////////////////////////////////////////////////////////////////////

// A Color is an index in the palette currently in use.
type Color uint8

// Transparent is the only reserved color index. All palettes start with it.
const Transparent = Color(0)

////////////////////////////////////////////////////////////////////////////////

// LRGBA returns the color corresponding to a color index.
func (c Color) LRGBA() color.LRGBA {
	return palettes.stdcolors[palettes.current][c]
}

// Color returns the color corresponding to a color index.
func (c Color) Color() color.Color {
	return palettes.colors[palettes.current][c]
}
