// Copyright (c) 2013-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package color

////////////////////////////////////////////////////////////////////////////////

// An Index identifies a color inside the palette.
type Index uint8

// Transparent is a reserved index in every palettes.
const Transparent Index = 0

////////////////////////////////////////////////////////////////////////////////

// Color returns the color corresponding to a color index.
func (c Index) Color() LRGBA {
	return colours[c]
}

// Set changes a color in the active palette.
func (c Index) Set(v Color) {
	if v == nil {
		colours[c] = LRGBA{1, 0, .5, 1}
	} else {
		colours[c] = LRGBAof(v)
	}
	// palcolours[active][c] = LRGBAOf(v)
	palettes.changed[active] = true
}

////////////////////////////////////////////////////////////////////////////////
