// Copyright (c) 2018-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

// Package pico8 provides the color palette of the PICO-8 fantasy console.
//
// The palette was created by Joseph White, and is licensed under CC-0.
//
// Source:
//   https://twitter.com/lexaloffle/status/732649035165667329
//   https://www.lexaloffle.com/pico-8.php?page=faq
//   http://www.lexaloffle.com/gfx/pico8_palette.png
package pico8

import (
	"github.com/cozely/cozely/color"
)

////////////////////////////////////////////////////////////////////////////////

// Palette is the PICO-8 palette.
var Palette = color.Palette{
	ByName: map[string]color.Index{
		"Dark Blue":   1,
		"Dark Purple": 2,
		"Dark Green":  3,
		"Brown":       4,
		"Dark Gray":   5,
		"Light Gray":  6,
		"White":       7,
		"Red":         8,
		"Orange":      9,
		"Yellow":      10,
		"Green":       11,
		"Blue":        12,
		"Indigo":      13,
		"Pink":        14,
		"Peach":       15,
		"Black":       16,
	},
	Colors: []color.LRGBA{
		color.LRGBAof(color.SRGB8{0x28, 0x22, 0x53}),
		color.LRGBAof(color.SRGB8{0x7E, 0x25, 0x53}),
		color.LRGBAof(color.SRGB8{0x00, 0x87, 0x51}),
		color.LRGBAof(color.SRGB8{0xAB, 0x52, 0x36}),
		color.LRGBAof(color.SRGB8{0x5F, 0x57, 0x4F}),
		color.LRGBAof(color.SRGB8{0xC2, 0xC3, 0xC7}),
		color.LRGBAof(color.SRGB8{0xFF, 0xF1, 0xE8}),
		color.LRGBAof(color.SRGB8{0xFF, 0x00, 0x4D}),
		color.LRGBAof(color.SRGB8{0xFF, 0xA3, 0x00}),
		color.LRGBAof(color.SRGB8{0xFF, 0xEC, 0x27}),
		color.LRGBAof(color.SRGB8{0x00, 0xE4, 0x36}),
		color.LRGBAof(color.SRGB8{0x29, 0xAD, 0xFF}),
		color.LRGBAof(color.SRGB8{0x83, 0x76, 0x9C}),
		color.LRGBAof(color.SRGB8{0xFF, 0x77, 0xA8}),
		color.LRGBAof(color.SRGB8{0xFF, 0xCC, 0xAA}),
		color.LRGBAof(color.SRGB8{0x00, 0x00, 0x00}),
	},
}

// Color indices constants
const (
	DarkBlue   color.Index = 1
	DarkPurple color.Index = 2
	DarkGreen  color.Index = 3
	Brown       color.Index = 4
	DarkGray   color.Index = 5
	LightGray  color.Index = 6
	White       color.Index = 7
	Red         color.Index = 8
	Orange      color.Index = 9
	Yellow      color.Index = 10
	Green       color.Index = 11
	Blue        color.Index = 12
	Indigo      color.Index = 13
	Pink        color.Index = 14
	Peach       color.Index = 15
	Black       color.Index = 16
)
