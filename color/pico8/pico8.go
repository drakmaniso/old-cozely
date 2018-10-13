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

// Color indices.
const (
	Transparent color.Index = iota
	DarkBlue
	DarkPurple
	DarkGreen
	Brown
	DarkGray
	LightGray
	White
	Red
	Orange
	Yellow
	Green
	Blue
	Indigo
	Pink
	Peach
	Black
)

// Palette is the PICO-8 palette.
var Palette = color.Palette{
	Transparent: color.LRGBA{},
	DarkBlue:    color.SRGB8{0x28, 0x22, 0x53},
	DarkPurple:  color.SRGB8{0x7E, 0x25, 0x53},
	DarkGreen:   color.SRGB8{0x00, 0x87, 0x51},
	Brown:       color.SRGB8{0xAB, 0x52, 0x36},
	DarkGray:    color.SRGB8{0x5F, 0x57, 0x4F},
	LightGray:   color.SRGB8{0xC2, 0xC3, 0xC7},
	White:       color.SRGB8{0xFF, 0xF1, 0xE8},
	Red:         color.SRGB8{0xFF, 0x00, 0x4D},
	Orange:      color.SRGB8{0xFF, 0xA3, 0x00},
	Yellow:      color.SRGB8{0xFF, 0xEC, 0x27},
	Green:       color.SRGB8{0x00, 0xE4, 0x36},
	Blue:        color.SRGB8{0x29, 0xAD, 0xFF},
	Indigo:      color.SRGB8{0x83, 0x76, 0x9C},
	Pink:        color.SRGB8{0xFF, 0x77, 0xA8},
	Peach:       color.SRGB8{0xFF, 0xCC, 0xAA},
	Black:       color.SRGB8{0x00, 0x00, 0x00},
}
