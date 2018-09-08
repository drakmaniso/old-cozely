// Copyright (c) 2018-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

// Package c64 provides the color palette of C64 microcomputers.
package c64

import (
	"github.com/cozely/cozely/color"
)

// Palette is the C64 palette.
var Palette = color.Palette{
	Names: map[string]color.Index{
		"Black":       1,
		"White":       2,
		"Red":         3,
		"Cyan":        4,
		"Violet":      5,
		"Green":       6,
		"Blue":        7,
		"Yellow":      8,
		"Orange":      9,
		"Brown":       10,
		"Light Red":   11,
		"Dark Grey":   12,
		"Grey":        13,
		"Light Green": 14,
		"Light Blue":  15,
		"Light Grey":  16,
	},

	Colors: []color.LRGBA{
		color.LRGBAof(color.SRGB8{0x00, 0x00, 0x00}),
		color.LRGBAof(color.SRGB8{0xff, 0xff, 0xff}),
		color.LRGBAof(color.SRGB8{0x68, 0x37, 0x2b}),
		color.LRGBAof(color.SRGB8{0x70, 0xa4, 0xb2}),
		color.LRGBAof(color.SRGB8{0x6f, 0x3d, 0x86}),
		color.LRGBAof(color.SRGB8{0x58, 0x8d, 0x43}),
		color.LRGBAof(color.SRGB8{0x35, 0x28, 0x79}),
		color.LRGBAof(color.SRGB8{0xb8, 0xc7, 0x6f}),
		color.LRGBAof(color.SRGB8{0x6f, 0x4f, 0x25}),
		color.LRGBAof(color.SRGB8{0x43, 0x39, 0x00}),
		color.LRGBAof(color.SRGB8{0x9a, 0x67, 0x59}),
		color.LRGBAof(color.SRGB8{0x44, 0x44, 0x44}),
		color.LRGBAof(color.SRGB8{0x6c, 0x6c, 0x6c}),
		color.LRGBAof(color.SRGB8{0x9a, 0xd2, 0x84}),
		color.LRGBAof(color.SRGB8{0x6c, 0x5e, 0xb5}),
		color.LRGBAof(color.SRGB8{0x95, 0x95, 0x95}),
	},
}
