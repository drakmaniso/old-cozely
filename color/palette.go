// Copyright (c) 2018-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package color

type PaletteID uint8

////////////////////////////////////////////////////////////////////////////////

var palettes struct {
	// For each palette
	path    [][]string
	changed []bool

	// For each palette/color combination
	colours [][256]Color
}

// Active palette
var (
	active    PaletteID
	activated bool
	colours   [256]struct{ R, G, B, A float32 }
)

////////////////////////////////////////////////////////////////////////////////

func Palette(path ...string) PaletteID {
	a := len(palettes.path)

	if a >= 0xFF {
		//TODO: set error
	}

	palettes.path = append(palettes.path, path)
	palettes.changed = append(palettes.changed, false)
	palettes.colours = append(palettes.colours, [256]Color{})

	return PaletteID(a)
}

func (a PaletteID) Activate() {
	//TODO
	for i := range palettes.colours[a] {
		Index(i).Set(palettes.colours[a][i])
	}

	active = a
	activated = true
}

////////////////////////////////////////////////////////////////////////////////

// Clear removes all colors from the palette.
func (a PaletteID) Clear() {
	//TODO:
	for c := range palettes.colours[a] {
		palettes.colours[a][c] = nil
	}
	for c := range colours {
		if active == a {
			colours[c] = LRGBA{0, 0, 0, 0}
		}
	}
}

// Entry associates a color to an unused entry of the palette, and returns its
// index. If every entries are used, it returns the transparent index.
//
// It starts looking from the end of the palette (index 255). This way, the same
// palette can also load colors from a file (starting at index 0).
func (a PaletteID) Entry(c Color) Index {
	for i := 255; i > 0; i-- {
		cc := palettes.colours[a][i]
		if cc == nil {
			palettes.colours[a][i] = c
			if active == a {
				colours[i] = LRGBAof(c)
			}
			return Index(i)
		}
	}
	return Transparent
}

////////////////////////////////////////////////////////////////////////////////

// Match searches for a color by its color.LRGBA values. If this exact color
// isn't in the palette, index 0 is returned.
//TODO: search by color proximity
func (a PaletteID) Match(v Color) Index {
	lv := LRGBAof(v)
	for c, pv := range palettes.colours[a] {
		lpv := LRGBAof(pv)
		if lpv == lv {
			return Index(c)
		}
	}

	return Index(0)
}

////////////////////////////////////////////////////////////////////////////////
