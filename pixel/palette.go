// Copyright (c) 2018-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package pixel

import (
	"github.com/cozely/cozely/color"
)

////////////////////////////////////////////////////////////////////////////////

var palette struct {
	colors [256]color.LRGBA
	used   [256]bool
	dirty  bool
}

func init() {
	clearPalette()
}

var debugColor = color.LRGBA{1, 0, 1, 1}

////////////////////////////////////////////////////////////////////////////////

func clearPalette() {
	for j := range palette.colors {
		i := color.Index(j)
		switch i {
		case Transparent:
			SetColor(i, color.LRGBA{0, 0, 0, 0})
		case Black:
			SetColor(i, color.SRGBA{0, 0, 0, 1})
		case DarkGray:
			SetColor(i, color.SRGBA{0.25, 0.25, 0.25, 1})
		case MidGray:
			SetColor(i, color.SRGBA{0.5, 0.5, 0.5, 1})
		case LightGray:
			SetColor(i, color.SRGBA{0.75, 0.75, 0.75, 1})
		case White:
			SetColor(i, color.SRGBA{1, 1, 1, 1})
		default:
			SetColor(i, nil)
		}
		palette.dirty = true
	}
}

////////////////////////////////////////////////////////////////////////////////

// SetColor changes the color associated with an index.
//
// Note that the modified palette will be used for every drawing command of the
// current frame, even those issued before the call to this function. In other
// words, you cannot modify the palette in the middle of a frame.
func SetColor(i color.Index, c color.Color) {
	switch {
	case i == 0:
		return

	case c == nil:
		palette.colors[i] = debugColor
		palette.used[i] = (i == 0)

	default:
		palette.colors[i] = color.LRGBAof(c)
		palette.used[i] = true
	}
	palette.dirty = true //TODO: finer-grained palette upload?
}

// AddColor finds an unsed index in the palette and adds a new color. It returns
// the found index, or 0 if the palette is full.
func AddColor(c color.Color) color.Index {
	i := color.Index(1)
	for ; i < Black; i++ {
		if !palette.used[i] {
			SetColor(i, c)
			return i
		}
	}
	return color.Index(0)
}

////////////////////////////////////////////////////////////////////////////////

// FindColor returns the first color index associated with specific LRGBA
// values. If there isn't any color with these values in the palette, index 0 is
// returned.
func FindColor(c color.Color) color.Index {
	lc := color.LRGBAof(c)
	for i, pc := range palette.colors {
		if i == 0 || !palette.used[i] {
			continue
		}
		if pc == lc {
			return color.Index(i)
		}
	}

	return color.Index(0)
}

//TODO: search by color proximity
