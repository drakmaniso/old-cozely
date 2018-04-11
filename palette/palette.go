// Copyright (c) 2013-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package palette

import (
	"github.com/drakmaniso/cozely/colour"
)

//------------------------------------------------------------------------------

// An Index identifies a color inside the palette.
type Index uint8

// Transparent is the only reserved index of every palettes.
const Transparent Index = 0

var (
	colours [256]struct{ R, G, B, A float32 }
	names   map[string]Index
	changed bool
)

func init() {
	names = make(map[string]Index, 256)
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
	Index(0).SetColour(colour.LRGBA{0, 0, 0, 0})
	Index(0).Rename("transparent")
}

//------------------------------------------------------------------------------

// Find returns the index associated with a name. If there isn't any, index 0 is
// returned.
func Find(name string) Index {
  c, _ := names[name]
  return c
}

// Match searches for a color by its colour.LRGBA values. If this exact color
// isn't in the palette, index 0 is returned.
func Match(v colour.Colour) Index {
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
