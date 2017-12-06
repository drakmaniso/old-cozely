// Copyright (c) 2013-2017 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package gfx

//------------------------------------------------------------------------------

import (
	"errors"

	"github.com/drakmaniso/carol/internal/core"
	"github.com/drakmaniso/carol/internal/gpu"
)

//------------------------------------------------------------------------------

// An Color identifies a color inside a palette.
type Color uint8

// A Palette identifies one of the 256 available palettes.
type Palette uint8

//------------------------------------------------------------------------------

var (
	palettes [256]struct {
		colours    [256]struct{ R, G, B, A float32 }
		count      int
		names      map[string]Color
		hasChanged bool
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

// PaletteCount returns the number of palettes created.
func PaletteCount() int {
	return palCount
}

//------------------------------------------------------------------------------

// NewPalette creates a new palette and returns its identifier. The name must be
// either unique or empty.
//
// Note: there is a maximum of 256 palettes.
func NewPalette(name string) (Palette, error) {
	if palCount > 255 {
		return Palette(0), errors.New("impossible to create palette \"" + name + "\": maximum palette count reached.")
	}

	p := Palette(palCount)
	palCount++

	if name != "" {
		if _, ok := palNames[name]; ok {
			return Palette(0), errors.New(`impossible to create palette: name "` + name + `" already taken.`)
		}
		palNames[name] = p
	}

	return p, nil
}

// GetPalette returns the palette associated with a name. If there isn't any, the
// default palette is returned, and a sticky error is set.
func GetPalette(name string) Palette {
	p, ok := palNames[name]
	if !ok {
		setErr("in GetPalette", errors.New("palette \"" + name + "\" not found"))
	}
	return p
}

//------------------------------------------------------------------------------

// New adds a color to the palette and returns its index. The name must be
// either unique or empty.
//
// Note: A palette contains a maximum of 256 colors.
func (p Palette) New(name string, v RGBA) (Color, error) {
	if palettes[p].count > 255 {
		return Color(0), errors.New("impossible to add color \"" + name + "\": maximum color count reached.")
	}

	c := Color(palettes[p].count)
	palettes[p].count++
	palettes[p].colours[c] = v
	palettes[p].hasChanged = true

	if name != "" {
		if _, ok := palettes[p].names[name]; ok {
			return Color(0), errors.New(`impossible to add color: name "` + name + `" already taken.`)
		}
		palettes[p].names[name] = c
	}

	return c, nil
}

//------------------------------------------------------------------------------

// Get returns the color associated with a name. If there isn't any, the first
// color is returned, and a sticky error is set.
func (p Palette) Get(name string) Color {
	c, ok := palettes[p].names[name]
	if !ok {
		setErr("in palette Get", errors.New("color \"" + name + "\" not found"))
	}
	return c
}

// Find searches for a color by its RGBA values. If this exact color isn't in
// the palette, the first color is returned, and ok is set to false.
func (p Palette) Find(v RGBA) (c Color, ok bool) {
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
	palettes[p].hasChanged = true
}

//------------------------------------------------------------------------------

func init() {
	c := core.Hook{
		Callback: postDrawHook,
		Context:  "in screen package post-Draw hook",
	}
	core.PostDrawHooks = append(core.PostDrawHooks, c)
}

func postDrawHook() error {
	for p := range palettes {
		if palettes[p].hasChanged {
			gpu.UpdatePaletteBuffer(uint8(p), palettes[p].colours[:])
			palettes[p].hasChanged = false
		}
	}
	return nil
}

//------------------------------------------------------------------------------
