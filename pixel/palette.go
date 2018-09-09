// Copyright (c) 2018-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package pixel

import (
	"github.com/cozely/cozely/color"
)

////////////////////////////////////////////////////////////////////////////////

// DefaultPalette is the palette initially loaded when the framework starts.
var DefaultPalette = color.Palette{
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

var palette struct {
	colors [256]color.LRGBA
	dirty  bool
}

func init() {
	SetPalette(DefaultPalette)
}

var debugColor = color.LRGBA{0, 0, 0, 1}

////////////////////////////////////////////////////////////////////////////////

// SetPalette asks the GPU to change the color palette.
//
// Note that the palette will be used for every drawing command of the current
// frame, even those issued before the call to Use. In other words, you cannot
// change the palette in the middle of a frame.
func SetPalette(p color.Palette) {
	for c := range palette.colors {
		switch {
		case c == 0:
			palette.colors[c] = color.LRGBA{0, 0, 0, 0}
		case c-1 < len(p.Colors):
			palette.colors[c] = p.Colors[c-1]
		default:
			palette.colors[c] = debugColor
		}
	}
	palette.dirty = true
}

////////////////////////////////////////////////////////////////////////////////

// SetColor changes the color associated with an index.
func SetColor(i color.Index, c color.Color) color.Index {
	if c == nil {
		palette.colors[i] = color.LRGBA{1, 0, .5, 1}
	} else {
		palette.colors[i] = color.LRGBAof(c)
	}
	palette.dirty = true //TODO: finer-grained palette upload
	return color.Index(i)
}

////////////////////////////////////////////////////////////////////////////////

// FindColor returns the first color index associated with specific LRGBA
// values. If there isn't any color with these values in the palette, index 0 is
// returned.
func FindColor(v color.Color) color.Index {
	lv := color.LRGBAof(v)
	for c, pv := range palette.colors {
		if pv == lv {
			return color.Index(c)
		}
	}

	return color.Index(0)
}

//TODO: search by color proximity
