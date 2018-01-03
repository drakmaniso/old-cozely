// Copyright (c) 2013-2017 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package pixel

//------------------------------------------------------------------------------

import (
	"errors"

	"github.com/drakmaniso/carol/colour"
	"github.com/drakmaniso/carol/core/gl"
)

//------------------------------------------------------------------------------

// A Color identifies a color inside the palette.
type Color uint8

//------------------------------------------------------------------------------

var (
	// colours needs to be outside of palette, otherwise cgo will freak out
	colours [256]struct{ R, G, B, A float32 }
	palette struct {
		count   int
		names   map[string]Color
		changed bool
	}
)

func init() {
	palette.names = make(map[string]Color, 256)
	ClearPalette()
}

//------------------------------------------------------------------------------

var paletteSSBO gl.StorageBuffer

//------------------------------------------------------------------------------

func ClearPalette() {
	for n := range palette.names {
		delete(palette.names, n)
	}
	for c := range colours {
		colours[c] = colour.RGBA{1, 0, 1, 1}
	}
	palette.count = 0
	NewColor("transparent", colour.RGBA{0, 0, 0, 0})
	colours[255] = colour.RGBA{1, 1, 1, 1}
}

//------------------------------------------------------------------------------

// ColorCount returns the number of colors in the palette.
func ColorCount() int {
	return palette.count
}

//------------------------------------------------------------------------------

// NewColor adds a color to the  palette and returns its index. The name must be
// either unique or empty.
//
// Note: The palette contains a maximum of 256 colors.
func NewColor(name string, v colour.Colour) Color {
	if palette.count > 255 {
		setErr("in NewColor", errors.New("impossible to add color \""+name+"\": maximum color count reached."))
		return Color(0)
	}

	c := Color(palette.count)
	palette.count++

	colours[c] = colour.RGBAOf(v)

	palette.changed = true

	if name != "" {
		if _, ok := palette.names[name]; ok {
			setErr("in NewColor", errors.New(`impossible to add color: name "`+name+`" already taken.`))
			return Color(0)
		}
		palette.names[name] = c
	}

	return c
}

//------------------------------------------------------------------------------

func requestColor(v colour.Colour) Color {
	if palette.count > 255 {
		setErr("in requestColor", errors.New("impossible to automatically add color: maximum color count reached."))
		return Color(0)
	}

	rgba := colour.RGBAOf(v)

	// Search the color in the existing palette

	for i := Color(0); i < Color(palette.count); i++ {
		if colours[i] == rgba {
			return i
		}
	}

	// Color not found, create a new entry

	c := Color(palette.count)
	palette.count++

	colours[c] = rgba

	palette.changed = true

	return c
}

//------------------------------------------------------------------------------

// GetColor returns the color associated with a name. If there isn't any, the first
// color is returned, and a sticky error is set.
func GetColor(name string) Color {
	c, ok := palette.names[name]
	if !ok {
		setErr("in GetColor", errors.New("color \""+name+"\" not found"))
	}
	return c
}

// FindColor searches for a color by its colour.RGBA values. If this exact color isn't in
// the palette, the first color is returned, and ok is set to false.
func FindColor(v colour.Colour) (c Color, ok bool) {
	lv := colour.RGBAOf(v)
	for c, vv := range colours {
		if vv == lv {
			return Color(c), true
		}
	}

	return Color(0), false
}

//------------------------------------------------------------------------------

// RGBA returns the color corresponding to a palette index.
func (c Color) RGBA() colour.RGBA {
	return colours[c]
}

// SetRGBA changes the colour.RGBA values of a color.
func (c Color) SetRGBA(v colour.Colour) {
	colours[c] = colour.RGBAOf(v)
	palette.changed = true
}

//------------------------------------------------------------------------------
