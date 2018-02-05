// Copyright (c) 2013-2018 Laurent Moussault. All rights reserved.
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

type palette [256]struct{ R, G, B, A float32 }

var (
	// colours needs to be outside of palette, otherwise cgo will freak out
	colours palette
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
		colours[c] = colour.LRGBA{1, 0, 1, 1}
	}
	Entry(0, "transparent", colour.LRGBA{0, 0, 0, 0})
}

//------------------------------------------------------------------------------

// Entry configures an entry in the palette and returns its index. The name
// should be either unique or empty.
func Entry(index uint8, name string, v colour.Colour) Index {
	c := Index(index)

	colours[c] = colour.LRGBAOf(v)

	changed = true

	if name != "" {
		_, ok := names[name]
		if ok {
			//TODO: error?
		}
		names[name] = c
	}

	return c
}

//------------------------------------------------------------------------------

// Get returns the index associated with a name. If there isn't any, index 0 is
// returned.
func Get(name string) Index {
	c, _ := names[name]
	return c
}

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

// Set changes the colour.LRGBA values of a color.
func (c Index) Set(v colour.Colour) {
	colours[c] = colour.LRGBAOf(v)
	changed = true
}

// Name returns the current name of the index, or the empty string if it is
// unnamed.
func (c Index) Name() string {
	for n, nc := range names {
		if c == nc {
			return n
		}
	}
	return ""
}

// Rename changes the name of an index. If the empty string is used, the index
// becomes unnamed.
func (c Index) Rename(n string) {
	if n != "" {
		names[n] = c
	} else {
		delete(names, n)
	}
}

//------------------------------------------------------------------------------
