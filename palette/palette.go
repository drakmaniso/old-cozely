// Copyright (c) 2013-2017 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package palette

import (
	"github.com/drakmaniso/glam/colour"
)

//------------------------------------------------------------------------------

// A Index identifies a color inside the palette.
type Index uint8

// Transparent is the only reserved index of every palettes.
const Transparent Index = 0

//------------------------------------------------------------------------------

var (
	// colours needs to be outside of palette, otherwise cgo will freak out
	colours [256]struct{ R, G, B, A float32 }
	count   int
	names   map[string]Index
	changed bool
)

func init() {
	names = make(map[string]Index, 256)
	Clear()
}

//------------------------------------------------------------------------------

// Clear removes all colors and names from the palette, initialize index 0 with
// a fully transparent color.
//
// Note: for debugging purpose, all unused indexes are initialized with pure
// magenta.
func Clear() {
	for n := range names {
		delete(names, n)
	}
	for c := range colours {
		colours[c] = colour.RGBA{1, 0, 1, 1}
	}
	count = 0
	New("transparent", colour.RGBA{0, 0, 0, 0})
	colours[255] = colour.RGBA{1, 1, 1, 1}
}

//------------------------------------------------------------------------------

// Count returns the number of colors in the palette.
func Count() int {
	return count
}

//------------------------------------------------------------------------------

// New adds a color to the  palette and returns its index. The name must be
// either unique or empty. If the palette is full, index 0 is returnd.
func New(name string, v colour.Colour) Index {
	if count > 255 {
		return Index(0)
	}

	c := Index(count)
	count++

	colours[c] = colour.RGBAOf(v)

	changed = true

	if name != "" {
		if _, ok := names[name]; ok {
			return Index(0)
		}
		names[name] = c
	}

	return c
}

//------------------------------------------------------------------------------

// Request tries to find an existing index for the specified colour, and returns
// it. If it is not present in the palette, it is added and the new index
// returned. If the palette is full, index 0 is returned.
func Request(v colour.Colour) Index {
	if count > 255 {
		return Index(0)
	}

	rgba := colour.RGBAOf(v)

	// Search the color in the existing palette

	for i := Index(0); i < Index(count); i++ {
		if colours[i] == rgba {
			return i
		}
	}

	// Index not found, create a new entry

	c := Index(count)
	count++

	colours[c] = rgba

	changed = true

	return c
}

//------------------------------------------------------------------------------

// Get returns the index associated with a name. If there isn't any, index 0 is
// returned.
func Get(name string) Index {
	c, _ := names[name]
	return c
}

// Find searches for a color by its colour.RGBA values. If this exact color
// isn't in the palette, index 0 is returned.
func Find(v colour.Colour) Index {
	lv := colour.RGBAOf(v)
	for c, vv := range colours {
		if vv == lv {
			return Index(c)
		}
	}

	return Index(0)
}

//------------------------------------------------------------------------------

// Colour returns the color corresponding to a palette index.
func (c Index) Colour() colour.RGBA {
	return colours[c]
}

// Set changes the colour.RGBA values of a color.
func (c Index) Set(v colour.Colour) {
	colours[c] = colour.RGBAOf(v)
	changed = true
}

//------------------------------------------------------------------------------