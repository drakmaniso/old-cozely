// Copyright (c) 2018-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package msx

import (
	"github.com/cozely/cozely/color"
	"github.com/cozely/cozely/pixel"
)

// A constant of each palette entry.
const (
	Transparent = pixel.Color(iota)
	Black
	MediumGreen
	LightGreen
	DarkBlue
	LightBlue
	DarkRed
	Cyan
	MediumRed
	LightRed
	DarkYellow
	LightYellow
	DarkGreen
	Magenta
	Gray
	White
)

var Colors = [256]color.Color{
	color.SRGB8{0x00, 0x00, 0x00},
	color.SRGB8{0x3E, 0xB8, 0x49},
	color.SRGB8{0x74, 0xd0, 0x7d},
	color.SRGB8{0x59, 0x55, 0xe0},
	color.SRGB8{0x80, 0x76, 0xf1},
	color.SRGB8{0xb9, 0x5e, 0x51},
	color.SRGB8{0x65, 0xdb, 0xef},
	color.SRGB8{0xdb, 0x65, 0x59},
	color.SRGB8{0xff, 0x89, 0x7d},
	color.SRGB8{0xcc, 0xc3, 0x5e},
	color.SRGB8{0xde, 0xd0, 0x87},
	color.SRGB8{0x3a, 0xa2, 0x41},
	color.SRGB8{0xb7, 0x66, 0xb5},
	color.SRGB8{0xcc, 0xcc, 0xcc},
	color.SRGB8{0xff, 0xff, 0xff},
}
