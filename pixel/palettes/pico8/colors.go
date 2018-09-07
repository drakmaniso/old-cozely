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
	"github.com/cozely/cozely/pixel"
)

// Color names, in palette order.
const (
	Transparent = pixel.Color(iota)
	Black       // 1
	DarkBlue    // 2
	DarkPurple  // 3
	DarkGreen   // 4
	Brown       // 5
	DarkGrey    // 6
	LightGray   // 7
	White       // 8
	Red         // 9
	Orange      // 10
	Yellow      // 11
	Green       // 12
	Blue        // 13
	Indigo      // 14
	Pink        // 15
	Peach       // 16
)

// Colors is the PICO-8 palette.
var Colors = [256]color.Color{
	color.SRGBA{0, 0, 0, 0},
	color.SRGB8{0x00, 0x00, 0x00},
	color.SRGB8{0x28, 0x22, 0x53},
	color.SRGB8{0x7E, 0x25, 0x53},
	color.SRGB8{0x00, 0x87, 0x51},
	color.SRGB8{0xAB, 0x52, 0x36},
	color.SRGB8{0x5F, 0x57, 0x4F},
	color.SRGB8{0xC2, 0xC3, 0xC7},
	color.SRGB8{0xFF, 0xF1, 0xE8},
	color.SRGB8{0xFF, 0x00, 0x4D},
	color.SRGB8{0xFF, 0xA3, 0x00},
	color.SRGB8{0xFF, 0xEC, 0x27},
	color.SRGB8{0x00, 0xE4, 0x36},
	color.SRGB8{0x29, 0xAD, 0xFF},
	color.SRGB8{0x83, 0x76, 0x9C},
	color.SRGB8{0xFF, 0x77, 0xA8},
	color.SRGB8{0xFF, 0xCC, 0xAA},
}
