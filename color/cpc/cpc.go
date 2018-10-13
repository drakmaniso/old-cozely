// Copyright (c) 2018-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

// Package cpc provides the color palette of CPC microcomputers.
package cpc

import (
	"github.com/cozely/cozely/color"
)

////////////////////////////////////////////////////////////////////////////////

// Color indices.
const (
	Transparent color.Index = iota
	Black
	Blue
	BrightBlue
	Red
	Magenta
	Mauve
	BrightRed
	Purple
	BrightMagenta
	Green
	Cyan
	SkyBlue
	Yellow
	White
	PastelBlue
	Orange
	Pink
	PastelMagenta
	BrightGreen
	SeaGreen
	BrightCyan
	Lime
	PastelGreen
	PastelCyan
	BrightYellow
	PastelYellow
	BrightWhite
)

// Palette is the CPC palette.
var Palette = color.Palette{
	Transparent:   color.LRGBA{},
	Black:         color.SRGB8{0x00, 0x00, 0x00},
	Blue:          color.SRGB8{0x00, 0x00, 0x80},
	BrightBlue:    color.SRGB8{0x00, 0x00, 0xff},
	Red:           color.SRGB8{0x80, 0x00, 0x00},
	Magenta:       color.SRGB8{0x80, 0x00, 0x80},
	Mauve:         color.SRGB8{0x80, 0x00, 0xff},
	BrightRed:     color.SRGB8{0xff, 0x00, 0x00},
	Purple:        color.SRGB8{0xff, 0x00, 0x80},
	BrightMagenta: color.SRGB8{0xff, 0x00, 0xff},
	Green:         color.SRGB8{0x00, 0x80, 0x00},
	Cyan:          color.SRGB8{0x00, 0x80, 0x80},
	SkyBlue:       color.SRGB8{0x00, 0x80, 0xff},
	Yellow:        color.SRGB8{0x80, 0x80, 0x00},
	White:         color.SRGB8{0x80, 0x80, 0x80},
	PastelBlue:    color.SRGB8{0x80, 0x80, 0xff},
	Orange:        color.SRGB8{0xff, 0x80, 0x00},
	Pink:          color.SRGB8{0xff, 0x80, 0x80},
	PastelMagenta: color.SRGB8{0xff, 0x80, 0xff},
	BrightGreen:   color.SRGB8{0x00, 0xff, 0x00},
	SeaGreen:      color.SRGB8{0x00, 0xff, 0x80},
	BrightCyan:    color.SRGB8{0x00, 0xff, 0xff},
	Lime:          color.SRGB8{0x80, 0xff, 0x00},
	PastelGreen:   color.SRGB8{0x80, 0xff, 0x80},
	PastelCyan:    color.SRGB8{0x80, 0xff, 0xff},
	BrightYellow:  color.SRGB8{0xff, 0xff, 0x00},
	PastelYellow:  color.SRGB8{0xff, 0xff, 0x80},
	BrightWhite:   color.SRGB8{0xff, 0xff, 0xff},
}
