// Copyright (c) 2018-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package pixel

import (
	"github.com/cozely/cozely/color"
	"github.com/cozely/cozely/x/gl"
)

////////////////////////////////////////////////////////////////////////////////

// DefaultPalette is the palette initially loaded when the framework starts.
var DefaultPalette = color.Palette{
	Names: map[string]color.Index{
		"Black":       1,
		"Dark Blue":   2,
		"Dark Purple": 3,
		"Dark Green":  4,
		"Brown":       5,
		"Dark Gray":   6,
		"Light Gray":  7,
		"White":       8,
		"Red":         9,
		"Orange":      10,
		"Yellow":      11,
		"Green":       12,
		"Blue":        13,
		"Indigo":      14,
		"Pink":        15,
		"Peach":       16,
	},
	Colors: []color.LRGBA{
		color.LRGBAof(color.SRGB8{0x00, 0x00, 0x00}),
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
	},
}

var (
	paletteSSBO gl.StorageBuffer
	palette     [256]color.LRGBA
	dirtyPal    bool
)

func init() {
	Palette(DefaultPalette)
}

var debugColor = color.LRGBA{0, 0, 0, 1}

////////////////////////////////////////////////////////////////////////////////

// Palette asks the GPU to change the color palette.
//
// Note that the palette will be used for every drawing command of the current
// frame, even those issued before the call to Use. In other words, you cannot
// change the palette in the middle of a frame.
func Palette(p color.Palette) {
	for c := range palette {
		switch {
		case c == 0:
			palette[c] = color.LRGBA{0, 0, 0, 0}
		case c-1 < len(p.Colors):
			palette[c] = p.Colors[c-1]
		default:
			palette[c] = debugColor
		}
	}
	dirtyPal = true
}

////////////////////////////////////////////////////////////////////////////////

// SetColor changes the color associated with an index.
func SetColor(i color.Index, c color.Color) color.Index {
	if c == nil {
		palette[i] = color.LRGBA{1, 0, .5, 1}
	} else {
		palette[i] = color.LRGBAof(c)
	}
	dirtyPal = true //TODO: finer-grained palette upload
	return color.Index(i)
}

////////////////////////////////////////////////////////////////////////////////

// MatchColor searches for a color by its color.LRGBA values. If this exact color
// isn't in the palette, index 0 is returned.
func MatchColor(v color.Color) color.Index {
	lv := color.LRGBAof(v)
	for c, pv := range palette {
		if pv == lv {
			return color.Index(c)
		}
	}

	return color.Index(0)
}

//TODO: search by color proximity
