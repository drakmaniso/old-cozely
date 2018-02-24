// Copyright (c) 2013-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package palette

import (
	"github.com/drakmaniso/glam/colour"
)

//------------------------------------------------------------------------------

// An Index identifies a color inside the palette.
type Index uint8

// Transparent is the only reserved index of every palettes.
const Transparent Index = 0

var (
	colours [256]struct{ R, G, B, A float32 }
	changed bool
)

//------------------------------------------------------------------------------

// Clear removes all colors and names from the palette, initialize index 0 with
// a fully transparent color.
//
// Note: for debugging purpose, all unused indexes are initialized with pure
// magenta.
func Clear() {
	for c := range colours {
		colours[c] = colour.LRGBA{1, 0, 1, 1}
	}
	Index(0).SetColour(colour.LRGBA{0, 0, 0, 0})
}

//------------------------------------------------------------------------------

// Find searches for a color by its colour.LRGBA values. If this exact color
// isn't in the palette, index 0 is returned.
func Find(v colour.Colour) Index {
	lv := colour.LRGBAOf(v)
	for c, vv := range colours {
		if vv == lv {
			return Index(c)
		}
	}

	return Index(0)
}

//------------------------------------------------------------------------------

// Colour returns the color corresponding to a palette index.
func (c Index) Colour() colour.LRGBA {
	return colours[c]
}

// SetColour changes the color corresponding to a palette index.
func (c Index) SetColour(v colour.Colour) {
	colours[c] = colour.LRGBAOf(v)
	changed = true
}

//------------------------------------------------------------------------------
