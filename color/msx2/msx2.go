// Copyright (c) 2018-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

// Package msx2 provides the color palette of MSX2 microcomputers.
package msx2

import (
	"github.com/cozely/cozely/color"
)

////////////////////////////////////////////////////////////////////////////////

// Palette is the MSX2 palette.
var Palette = color.Palette{
	1: color.LRGBA{1,0,0,1},
}

func init() {
	for i := 1; i < 256; i++ {
		g, r, b := i>>5, (i&0x1C)>>2, i&0x3
		Palette[i] = color.LRGBA{
			float32(r) / 7.0,
			float32(g) / 7.0,
			float32(b) / 3.0,
			1.0,
		}
	}
}
