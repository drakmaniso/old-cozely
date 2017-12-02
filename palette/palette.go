// Copyright (c) 2013-2017 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package palette

//------------------------------------------------------------------------------

import (
	"errors"
)

//------------------------------------------------------------------------------

// An Colour identifies a colour inside a palette.
type Colour uint8

// A Palette identifies one of the 256 available palettes.
type Palette uint8

//------------------------------------------------------------------------------

var (
	palettes [256]struct {
		colours [256]RGBA
		count   int
		names   map[string]Colour
	}
	palCount int
	palNames map[string]Palette
)

func init() {
	palNames = make(map[string]Palette, 256)
	for p := range palettes {
		palettes[p].names = make(map[string]Colour, 256)
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

// Get searches for a palette by name. It returns its identifier and true if
// found, or the default palette and false otherwise.
func Get(name string) (Palette, bool) {
	p, ok := palNames[name]
	return p, ok
}

//------------------------------------------------------------------------------

// New adds a named colour to the palette, and returns its index. It returns
// color 0 and an error if the maximum number of colours is reached, or the name
// already taken.
func (p Palette) New(name string, v RGBA) (Colour, error) {
	if palettes[p].count > 255 {
		return Colour(0), errors.New("impossible to add colour \"" + name + "\": maximum colour count reached.")
	}

	if _, ok := palettes[p].names[name]; ok {
		return Colour(0), errors.New(`impossible to add colour: name "` + name + `" already taken.`)
	}

	c := Colour(palettes[p].count)
	palettes[p].names[name] = c
	palettes[p].colours[c] = v
	palettes[p].count++

	return c, nil
}

//------------------------------------------------------------------------------

// Get searches for a colour by name. It returns its index and true if found, or
// colour 0 and false otherwise.
func (p Palette) Get(name string) (Colour, bool) {
	c, ok := palettes[p].names[name]
	return c, ok
}

// Find searches for a colour by its RGBA values. It returns its index and true
// if found, or colour 0 and false otherwise.
func (p Palette) Find(v RGBA) (Colour, bool) {
	for c, vv := range palettes[p].colours {
		if vv == v {
			return Colour(c), true
		}
	}

	return Colour(0), false
}

//------------------------------------------------------------------------------

// RGBA returns the RGBA values of a colour.
func (p Palette) RGBA(i Colour) RGBA {
	return palettes[p].colours[i]
}

// SetRGBA changes the RGBA values of a colour.
func (p Palette) SetRGBA(c Colour, v RGBA) {
	palettes[p].colours[c] = v
}

//------------------------------------------------------------------------------
