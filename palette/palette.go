// Copyright (c) 2013-2017 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package palette

//------------------------------------------------------------------------------

import (
	"errors"
)

//------------------------------------------------------------------------------

// An Color identifies a color inside a palette.
type Color uint8

// A Palette identifies one of the 256 available palettes.
type Palette uint8

//------------------------------------------------------------------------------

var (
	palettes [256]struct {
		colours [256]RGBA
		count   int
		names   map[string]Color
	}
	palCount int
	palNames map[string]Palette
)

func init() {
	palNames = make(map[string]Palette, 256)
	for p := range palettes {
		palettes[p].names = make(map[string]Color, 256)
		for c := range palettes[p].colours {
			palettes[p].colours[c] = RGBA{1, 0, 1, 1}
		}
	}
}

//------------------------------------------------------------------------------

// Count returns the number of palettes created.
func Count() int {
	return palCount
}

//------------------------------------------------------------------------------

// New creates and name a new palette, and returns its identifier. It returns
// the default palette and an error if the maximum number of palette is reached,
// or if the name is already taken.
func New(name string) (Palette, error) {
	if palCount > 255 {
		return Palette(0), errors.New("impossible to create palette \"" + name + "\": maximum palette count reached.")
	}

	if _, ok := palNames[name]; ok {
		return Palette(0), errors.New(`impossible to create palette: name "` + name + `" already taken.`)
	}

	p := Palette(palCount)
	palNames[name] = p
	palCount++
	//TODO: register for GPU

	return p, nil
}

// Get returns the palette associated with a name. If there isn't any, the
// default palette is returned, and ok is set to false.
func Get(name string) (p Palette, ok bool) {
	p, ok = palNames[name]
	return p, ok
}

//------------------------------------------------------------------------------

// New adds a named color to the palette, and returns its index. It returns
// color 0 and an error if the maximum number of colours is reached, or the name
// already taken.
func (p Palette) New(name string, v RGBA) (Color, error) {
	if palettes[p].count > 255 {
		return Color(0), errors.New("impossible to add color \"" + name + "\": maximum color count reached.")
	}

	if _, ok := palettes[p].names[name]; ok {
		return Color(0), errors.New(`impossible to add color: name "` + name + `" already taken.`)
	}

	c := Color(palettes[p].count)
	palettes[p].names[name] = c
	palettes[p].colours[c] = v
	palettes[p].count++

	return c, nil
}

//------------------------------------------------------------------------------

// Get returns the color associated with a name. If there isn't any, the color 0
// is returned, and ok is set to false.
func (p Palette) Get(name string) (c Color, ok bool) {
	c, ok = palettes[p].names[name]
	return c, ok
}

// Find searches for a color by its RGBA values. It returns its index and true
// if found, or color 0 and false otherwise.
func (p Palette) Find(v RGBA) (Color, bool) {
	for c, vv := range palettes[p].colours {
		if vv == v {
			return Color(c), true
		}
	}

	return Color(0), false
}

//------------------------------------------------------------------------------

// RGBA returns the RGBA values of a color.
func (p Palette) RGBA(i Color) RGBA {
	return palettes[p].colours[i]
}

// SetRGBA changes the RGBA values of a color.
func (p Palette) SetRGBA(c Color, v RGBA) {
	palettes[p].colours[c] = v
}

//------------------------------------------------------------------------------
