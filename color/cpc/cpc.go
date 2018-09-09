// Copyright (c) 2018-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

// Package cpc provides the color palette of CPC microcomputers.
package cpc

import (
	"github.com/cozely/cozely/color"
)

////////////////////////////////////////////////////////////////////////////////

// Palette is the CPC palette.
var Palette = color.Palette{
	ByName: map[string]color.Index{
		"Black":          1,
		"Blue":           2,
		"Bright Blue":    3,
		"Red":            4,
		"Magenta":        5,
		"Mauve":          6,
		"Bright Red":     7,
		"Purple":         8,
		"Bright Magenta": 9,
		"Green":          10,
		"Cyan":           11,
		"SkyBlue":        12,
		"Yellow":         13,
		"White":          14,
		"Pastel Blue":    15,
		"Orange":         16,
		"Pink":           17,
		"Pastel Magenta": 18,
		"Bright Green":   19,
		"Sea Green":      20,
		"Bright Cyan":    21,
		"Lime":           22,
		"Pastel Green":   23,
		"Pastel Cyan":    24,
		"Bright Yellow":  25,
		"Pastel Yellow":  26,
		"Bright White":   27,
	},
	Colors: []color.LRGBA{
		color.LRGBAof(color.SRGB8{0x00, 0x00, 0x00}),
		color.LRGBAof(color.SRGB8{0x00, 0x00, 0x80}),
		color.LRGBAof(color.SRGB8{0x00, 0x00, 0xff}),
		color.LRGBAof(color.SRGB8{0x80, 0x00, 0x00}),
		color.LRGBAof(color.SRGB8{0x80, 0x00, 0x80}),
		color.LRGBAof(color.SRGB8{0x80, 0x00, 0xff}),
		color.LRGBAof(color.SRGB8{0xff, 0x00, 0x00}),
		color.LRGBAof(color.SRGB8{0xff, 0x00, 0x80}),
		color.LRGBAof(color.SRGB8{0xff, 0x00, 0xff}),
		color.LRGBAof(color.SRGB8{0x00, 0x80, 0x00}),
		color.LRGBAof(color.SRGB8{0x00, 0x80, 0x80}),
		color.LRGBAof(color.SRGB8{0x00, 0x80, 0xff}),
		color.LRGBAof(color.SRGB8{0x80, 0x80, 0x00}),
		color.LRGBAof(color.SRGB8{0x80, 0x80, 0x80}),
		color.LRGBAof(color.SRGB8{0x80, 0x80, 0xff}),
		color.LRGBAof(color.SRGB8{0xff, 0x80, 0x00}),
		color.LRGBAof(color.SRGB8{0xff, 0x80, 0x80}),
		color.LRGBAof(color.SRGB8{0xff, 0x80, 0xff}),
		color.LRGBAof(color.SRGB8{0x00, 0xff, 0x00}),
		color.LRGBAof(color.SRGB8{0x00, 0xff, 0x80}),
		color.LRGBAof(color.SRGB8{0x00, 0xff, 0xff}),
		color.LRGBAof(color.SRGB8{0x80, 0xff, 0x00}),
		color.LRGBAof(color.SRGB8{0x80, 0xff, 0x80}),
		color.LRGBAof(color.SRGB8{0x80, 0xff, 0xff}),
		color.LRGBAof(color.SRGB8{0xff, 0xff, 0x00}),
		color.LRGBAof(color.SRGB8{0xff, 0xff, 0x80}),
		color.LRGBAof(color.SRGB8{0xff, 0xff, 0xff}),
	},
}
