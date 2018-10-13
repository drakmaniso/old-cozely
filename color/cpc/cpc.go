// Copyright (c) 2018-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

// Package cpc provides the color palette of CPC microcomputers.
package cpc

import (
	"github.com/cozely/cozely/color"
)

////////////////////////////////////////////////////////////////////////////////

// Color names.
const (
	Black         = 1
	Blue          = 2
	BrightBlue    = 3
	Red           = 4
	Magenta       = 5
	Mauve         = 6
	BrightRed     = 7
	Purple        = 8
	BrightMagenta = 9
	Green         = 10
	Cyan          = 11
	SkyBlue       = 12
	Yellow        = 13
	White         = 14
	PastelBlue    = 15
	Orange        = 16
	Pink          = 17
	PastelMagenta = 18
	BrightGreen   = 19
	SeaGreen      = 20
	BrightCyan    = 21
	Lime          = 22
	PastelGreen   = 23
	PastelCyan    = 24
	BrightYellow  = 25
	PastelYellow  = 26
	BrightWhite   = 27
)

// Palette is the CPC palette.
var Palette = []color.Color{
	color.SRGB8{0x00, 0x00, 0x00},
	color.SRGB8{0x00, 0x00, 0x80},
	color.SRGB8{0x00, 0x00, 0xff},
	color.SRGB8{0x80, 0x00, 0x00},
	color.SRGB8{0x80, 0x00, 0x80},
	color.SRGB8{0x80, 0x00, 0xff},
	color.SRGB8{0xff, 0x00, 0x00},
	color.SRGB8{0xff, 0x00, 0x80},
	color.SRGB8{0xff, 0x00, 0xff},
	color.SRGB8{0x00, 0x80, 0x00},
	color.SRGB8{0x00, 0x80, 0x80},
	color.SRGB8{0x00, 0x80, 0xff},
	color.SRGB8{0x80, 0x80, 0x00},
	color.SRGB8{0x80, 0x80, 0x80},
	color.SRGB8{0x80, 0x80, 0xff},
	color.SRGB8{0xff, 0x80, 0x00},
	color.SRGB8{0xff, 0x80, 0x80},
	color.SRGB8{0xff, 0x80, 0xff},
	color.SRGB8{0x00, 0xff, 0x00},
	color.SRGB8{0x00, 0xff, 0x80},
	color.SRGB8{0x00, 0xff, 0xff},
	color.SRGB8{0x80, 0xff, 0x00},
	color.SRGB8{0x80, 0xff, 0x80},
	color.SRGB8{0x80, 0xff, 0xff},
	color.SRGB8{0xff, 0xff, 0x00},
	color.SRGB8{0xff, 0xff, 0x80},
	color.SRGB8{0xff, 0xff, 0xff},
}
