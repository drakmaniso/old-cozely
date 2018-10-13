// Copyright (c) 2018-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

// Package c64 provides the color palette of C64 microcomputers.
package c64

import (
	"github.com/cozely/cozely/color"
)

// Color indices.
const (
	Transparent color.Index = iota
	Black
	White
	Red
	Cyan
	Violet
	Green
	Blue
	Yellow
	Orange
	Brown
	LightRed
	DarkGrey
	Grey
	LightGreen
	LightBlue
	LightGrey
)

// Palette is the C64 palette.
var Palette = color.Palette{
	Transparent: color.LRGBA{},
	Black:       color.SRGB8{0x00, 0x00, 0x00},
	White:       color.SRGB8{0xff, 0xff, 0xff},
	Red:         color.SRGB8{0x68, 0x37, 0x2b},
	Cyan:        color.SRGB8{0x70, 0xa4, 0xb2},
	Violet:      color.SRGB8{0x6f, 0x3d, 0x86},
	Green:       color.SRGB8{0x58, 0x8d, 0x43},
	Blue:        color.SRGB8{0x35, 0x28, 0x79},
	Yellow:      color.SRGB8{0xb8, 0xc7, 0x6f},
	Orange:      color.SRGB8{0x6f, 0x4f, 0x25},
	Brown:       color.SRGB8{0x43, 0x39, 0x00},
	LightRed:    color.SRGB8{0x9a, 0x67, 0x59},
	DarkGrey:    color.SRGB8{0x44, 0x44, 0x44},
	Grey:        color.SRGB8{0x6c, 0x6c, 0x6c},
	LightGreen:  color.SRGB8{0x9a, 0xd2, 0x84},
	LightBlue:   color.SRGB8{0x6c, 0x5e, 0xb5},
	LightGrey:   color.SRGB8{0x95, 0x95, 0x95},
}
