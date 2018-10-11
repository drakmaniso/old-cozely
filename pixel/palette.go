// Copyright (c) 2018-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package pixel

import (
	"github.com/cozely/cozely/color"
	"github.com/cozely/cozely/color/pico8"
)

////////////////////////////////////////////////////////////////////////////////

var palette struct {
	colors [256]color.LRGBA
	dirty  bool
}

func init() {
	SetPalette(pico8.Palette)
}

var debugColor = color.LRGBA{1, 0, 1, 1}

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
			switch c {
			case 254:
				palette.colors[c] = color.LRGBA{0, 0, 0, 1}
			case 255:
				palette.colors[c] = color.LRGBA{1, 1, 1, 1}
			default:
				palette.colors[c] = debugColor
			}
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
