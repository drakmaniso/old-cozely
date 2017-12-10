// Copyright (c) 2013-2017 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package gfx

//------------------------------------------------------------------------------

import (
	"errors"
	"image/color"
)

//------------------------------------------------------------------------------

// A Color identifies a color inside the palette.
type Color uint8

//------------------------------------------------------------------------------

var (
	colours    [256]struct{ R, G, B, A float32 }
	colCount   int
	colNames   map[string]Color
	colChanged bool
)

func init() {
	colNames = make(map[string]Color, 256)
	ClearPalette()
}

//------------------------------------------------------------------------------

func ClearPalette() {
	for n := range colNames {
		delete(colNames, n)
	}
	for c := range colours {
		colours[c] = RGBA{1, 0, 1, 1}
	}
	colCount = 0
	NewColor("transparent", RGBA{0, 0, 0, 0})
	colours[255] = RGBA{1, 1, 1, 1}
}

//------------------------------------------------------------------------------

// ColorCount returns the number of colors in the palette.
func ColorCount() int {
	return colCount
}

//------------------------------------------------------------------------------

// NewColor adds a color to the  palette and returns its index. The name must be
// either unique or empty.
//
// Note: The palette contains a maximum of 256 colors.
func NewColor(name string, cc color.Color) Color {
	if colCount > 255 {
		setErr("in NewColor", errors.New("impossible to add color \""+name+"\": maximum color count reached."))
		return Color(0)
	}

	c := Color(colCount)
	colCount++

	v, ok := cc.(RGBA)
	if !ok {
		v = MakeRGBA(cc)
	}
	colours[c] = v

	colChanged = true

	if name != "" {
		if _, ok := colNames[name]; ok {
			setErr("in NewColor", errors.New(`impossible to add color: name "`+name+`" already taken.`))
			return Color(0)
		}
		colNames[name] = c
	}

	return c
}

//------------------------------------------------------------------------------

// GetColor returns the color associated with a name. If there isn't any, the first
// color is returned, and a sticky error is set.
func GetColor(name string) Color {
	c, ok := colNames[name]
	if !ok {
		setErr("in GetColor", errors.New("color \""+name+"\" not found"))
	}
	return c
}

// FindColor searches for a color by its RGBA values. If this exact color isn't in
// the palette, the first color is returned, and ok is set to false.
func FindColor(v RGBA) (c Color, ok bool) {
	for c, vv := range colours {
		if vv == v {
			return Color(c), true
		}
	}

	return Color(0), false
}

//------------------------------------------------------------------------------

// RGBA returns the RGBA values of a color.
func (c Color) RGBA() RGBA {
	return colours[c]
}

// SetRGBA changes the RGBA values of a color.
func (c Color) SetRGBA(cc color.Color) {
	v, ok := cc.(RGBA)
	if !ok {
		v = MakeRGBA(cc)
	}
	colours[c] = v
	colChanged = true
}

//------------------------------------------------------------------------------
