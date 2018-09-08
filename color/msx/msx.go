// Copyright (c) 2018-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

// Package msx provides the color palette of MSX1 microcomputers.
//
// Sources:
//  https://github.com/openMSX/openMSX/files/275509/TMS9928_family_Palette_v2.xlsx
//  https://github.com/openMSX/openMSX/issues/1024
package msx

import (
	"github.com/cozely/cozely/color"
)

var names = map[string]color.Index{
	"Black":        1,
	"Medium Green": 2,
	"Light Green":  3,
	"Dark Blue":    4,
	"Light Blue":   5,
	"Dark Red":     6,
	"Cyan":         7,
	"Medium Red":   8,
	"Light Red":    9,
	"Dark Yellow":  10,
	"Light Yellow": 11,
	"Dark Green":   12,
	"Magenta":      13,
	"Gray":         14,
	"White":        15,
}

// PaletteIdealized is an idealized MSX palette (source: Wikipedia)
var PaletteIdealized = color.Palette{
	Names: names,
	// Wikipedia
	Colors: []color.LRGBA{
		color.LRGBAof(color.SRGB8{0x00, 0x00, 0x00}),
		color.LRGBAof(color.SRGB8{0x3E, 0xB8, 0x49}),
		color.LRGBAof(color.SRGB8{0x74, 0xd0, 0x7d}),
		color.LRGBAof(color.SRGB8{0x59, 0x55, 0xe0}),
		color.LRGBAof(color.SRGB8{0x80, 0x76, 0xf1}),
		color.LRGBAof(color.SRGB8{0xb9, 0x5e, 0x51}),
		color.LRGBAof(color.SRGB8{0x65, 0xdb, 0xef}),
		color.LRGBAof(color.SRGB8{0xdb, 0x65, 0x59}),
		color.LRGBAof(color.SRGB8{0xff, 0x89, 0x7d}),
		color.LRGBAof(color.SRGB8{0xcc, 0xc3, 0x5e}),
		color.LRGBAof(color.SRGB8{0xde, 0xd0, 0x87}),
		color.LRGBAof(color.SRGB8{0x3a, 0xa2, 0x41}),
		color.LRGBAof(color.SRGB8{0xb7, 0x66, 0xb5}),
		color.LRGBAof(color.SRGB8{0xcc, 0xcc, 0xcc}),
		color.LRGBAof(color.SRGB8{0xff, 0xff, 0xff}),
	},
}

// Palette is the theorical MSX palette (i.e. conversions done with ITU-R BT.601
// formula).
var Palette = color.Palette{
	Names: names,
	// ITU-R BT.601
	Colors: []color.LRGBA{
		color.LRGBAof(color.SRGB8{0, 6, 0}),
		color.LRGBAof(color.SRGB8{26, 207, 60}),
		color.LRGBAof(color.SRGB8{85, 224, 112}),
		color.LRGBAof(color.SRGB8{77, 91, 230}),
		color.LRGBAof(color.SRGB8{119, 124, 247}),
		color.LRGBAof(color.SRGB8{203, 85, 68}),
		color.LRGBAof(color.SRGB8{60, 243, 238}),
		color.LRGBAof(color.SRGB8{246, 91, 78}),
		color.LRGBAof(color.SRGB8{255, 125, 112}),
		color.LRGBAof(color.SRGB8{205, 200, 77}),
		color.LRGBAof(color.SRGB8{221, 211, 119}),
		color.LRGBAof(color.SRGB8{25, 180, 50}),
		color.LRGBAof(color.SRGB8{195, 98, 179}),
		color.LRGBAof(color.SRGB8{196, 209, 196}),
		color.LRGBAof(color.SRGB8{247, 255, 247}),
	},
}

// PaletteCVtoRGB is the MSX palette through a Component Video to RGB conversion.
var PaletteCVtoRGB = color.Palette{
	Names: names,
	// Component Video to RGB
	Colors: []color.LRGBA{
		color.LRGBAof(color.SRGB8{0, 4, 0}),
		color.LRGBAof(color.SRGB8{58, 187, 67}),
		color.LRGBAof(color.SRGB8{112, 211, 119}),
		color.LRGBAof(color.SRGB8{84, 89, 215}),
		color.LRGBAof(color.SRGB8{123, 123, 232}),
		color.LRGBAof(color.SRGB8{179, 99, 75}),
		color.LRGBAof(color.SRGB8{97, 223, 231}),
		color.LRGBAof(color.SRGB8{212, 106, 83}),
		color.LRGBAof(color.SRGB8{248, 142, 119}),
		color.LRGBAof(color.SRGB8{199, 199, 89}),
		color.LRGBAof(color.SRGB8{217, 212, 129}),
		color.LRGBAof(color.SRGB8{54, 165, 59}),
		color.LRGBAof(color.SRGB8{176, 107, 174}),
		color.LRGBAof(color.SRGB8{199, 208, 197}),
		color.LRGBAof(color.SRGB8{250, 255, 248}),
	},
}

// PaletteCheapRGB is the MSX palette through a typical analog conversion (as
// done in the 80's), featured in the majority of MSX1 models with a TMS VDP and
// RGB output.
var PaletteCheapRGB = color.Palette{
	Names: names,
	// El Cheapo RGB
	Colors: []color.LRGBA{
		color.LRGBAof(color.SRGB8{0, 5, 0}),
		color.LRGBAof(color.SRGB8{26, 205, 59}),
		color.LRGBAof(color.SRGB8{87, 225, 112}),
		color.LRGBAof(color.SRGB8{77, 91, 230}),
		color.LRGBAof(color.SRGB8{117, 124, 245}),
		color.LRGBAof(color.SRGB8{204, 86, 69}),
		color.LRGBAof(color.SRGB8{59, 242, 237}),
		color.LRGBAof(color.SRGB8{245, 90, 77}),
		color.LRGBAof(color.SRGB8{255, 126, 112}),
		color.LRGBAof(color.SRGB8{204, 197, 77}),
		color.LRGBAof(color.SRGB8{222, 211, 120}),
		color.LRGBAof(color.SRGB8{26, 181, 51}),
		color.LRGBAof(color.SRGB8{194, 97, 179}),
		color.LRGBAof(color.SRGB8{196, 209, 196}),
		color.LRGBAof(color.SRGB8{247, 255, 247}),
	},
}

// PaletteCheapRGBTrim is the same as ColorsCheapRGB, but with trimpots to
// normalize each channels (provides more reddish tones for the red colors).
var PaletteCheapRGBTrim = color.Palette{
	Names: names,
	// El Cheapo RGB with trimpots
	Colors: []color.LRGBA{
		color.LRGBAof(color.SRGB8{0, 0, 0}),
		color.LRGBAof(color.SRGB8{26, 208, 62}),
		color.LRGBAof(color.SRGB8{88, 229, 118}),
		color.LRGBAof(color.SRGB8{78, 90, 241}),
		color.LRGBAof(color.SRGB8{120, 123, 255}),
		color.LRGBAof(color.SRGB8{208, 84, 72}),
		color.LRGBAof(color.SRGB8{60, 246, 249}),
		color.LRGBAof(color.SRGB8{250, 88, 80}),
		color.LRGBAof(color.SRGB8{255, 126, 118}),
		color.LRGBAof(color.SRGB8{208, 200, 80}),
		color.LRGBAof(color.SRGB8{226, 214, 126}),
		color.LRGBAof(color.SRGB8{26, 183, 54}),
		color.LRGBAof(color.SRGB8{198, 96, 188}),
		color.LRGBAof(color.SRGB8{200, 213, 206}),
		color.LRGBAof(color.SRGB8{252, 255, 255}),
	},
}

// PaletteLazyRGB is the same as ColorsCheapRGB, but with a lower quality analog
// conversion.
var PaletteLazyRGB = color.Palette{
	Names: names,
	// Lazy El Cheapo RGB
	Colors: []color.LRGBA{
		color.LRGBAof(color.SRGB8{0, 8, 0}),
		color.LRGBAof(color.SRGB8{21, 202, 53}),
		color.LRGBAof(color.SRGB8{80, 221, 105}),
		color.LRGBAof(color.SRGB8{71, 91, 219}),
		color.LRGBAof(color.SRGB8{110, 122, 234}),
		color.LRGBAof(color.SRGB8{194, 87, 63}),
		color.LRGBAof(color.SRGB8{53, 237, 227}),
		color.LRGBAof(color.SRGB8{234, 90, 71}),
		color.LRGBAof(color.SRGB8{255, 125, 105}),
		color.LRGBAof(color.SRGB8{194, 195, 71}),
		color.LRGBAof(color.SRGB8{212, 207, 113}),
		color.LRGBAof(color.SRGB8{21, 178, 46}),
		color.LRGBAof(color.SRGB8{184, 97, 170}),
		color.LRGBAof(color.SRGB8{187, 206, 187}),
		color.LRGBAof(color.SRGB8{236, 255, 236}),
	},
}
