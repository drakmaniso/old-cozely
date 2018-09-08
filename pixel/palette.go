// Copyright (c) 2018-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package pixel

import (
	"errors"
	"image"
	stdcolor "image/color"
	_ "image/png" // Activate PNG support
	"os"

	"github.com/cozely/cozely/color"
	"github.com/cozely/cozely/internal"
	"github.com/cozely/cozely/x/gl"
)

////////////////////////////////////////////////////////////////////////////////

// A PaletteID identifies a declared palette
type PaletteID uint8

// DefaultPalette is a palette that is always loaded.
const DefaultPalette = PaletteID(0)

var palettes struct {
	current PaletteID
	ssbo    gl.StorageBuffer

	// For each palette
	changed []bool
	path    []string

	// For each palette/color combination
	stdcolors [][256]color.LRGBA
}

func init() {
	palettes.changed = append(palettes.changed, true)
	palettes.path = append(palettes.path, "")
	pal := [256]color.Color{
		color.SRGBA{0, 0, 0, 0},
		color.SRGB8{0x00, 0x00, 0x00},
		color.SRGB8{0x28, 0x22, 0x53},
		color.SRGB8{0x7E, 0x25, 0x53},
		color.SRGB8{0x00, 0x87, 0x51},
		color.SRGB8{0xAB, 0x52, 0x36},
		color.SRGB8{0x5F, 0x57, 0x4F},
		color.SRGB8{0xC2, 0xC3, 0xC7},
		color.SRGB8{0xFF, 0xF1, 0xE8},
		color.SRGB8{0xFF, 0x00, 0x4D},
		color.SRGB8{0xFF, 0xA3, 0x00},
		color.SRGB8{0xFF, 0xEC, 0x27},
		color.SRGB8{0x00, 0xE4, 0x36},
		color.SRGB8{0x29, 0xAD, 0xFF},
		color.SRGB8{0x83, 0x76, 0x9C},
		color.SRGB8{0xFF, 0x77, 0xA8},
		color.SRGB8{0xFF, 0xCC, 0xAA},
	}
	palettes.stdcolors = append(palettes.stdcolors, [256]color.LRGBA{})

	for c := range pal {
		cc := pal[c]
		if cc == nil {
			palettes.stdcolors[0][c] = debugColor
			continue
		}
		palettes.stdcolors[0][c] = color.LRGBAof(cc)
	}
}

var debugColor = color.LRGBA{0, 0, 0, 1}

////////////////////////////////////////////////////////////////////////////////

func PaletteColors(pal [256]color.Color) PaletteID {
	//TODO: avoid copy?
	a := len(palettes.stdcolors)

	if a >= 0xFF {
		//TODO: set error
	}

	palettes.changed = append(palettes.changed, true)
	palettes.path = append(palettes.path, "")
	palettes.stdcolors = append(palettes.stdcolors, [256]color.LRGBA{})
	for c := range pal {
		cc := pal[c]
		if cc != nil {
			palettes.stdcolors[a][c] = color.LRGBAof(cc)
		} else {
			palettes.stdcolors[a][c] = debugColor
		}
	}

	return PaletteID(a)
}

func Palette(path string) PaletteID {
	a := len(palettes.stdcolors)

	if a >= 0xFF {
		//TODO: set error
	}

	palettes.changed = append(palettes.changed, false)
	palettes.path = append(palettes.path, path)
	palettes.stdcolors = append(palettes.stdcolors, [256]color.LRGBA{})

	return PaletteID(a)
}

func (a PaletteID) load(name string) error {
	f, err := os.Open(internal.Path + name + ".png")
	if err != nil {
		return errors.New("unable to open file for palette " + name)
	}
	defer f.Close() //TODO: error handling
	cf, _, err := image.DecodeConfig(f)
	if err != nil {
		return errors.New("unable to decode file for palette " + name)
	}

	p, ok := cf.ColorModel.(stdcolor.Palette)
	if !ok {
		return errors.New("image file not paletted for palette " + name)
	}

	//TODO: clear the palette?

	palettes.stdcolors[a][0] = color.LRGBA{0, 0, 0, 0}
	j := 1
	for i := range p {
		r, g, b, al := p[i].RGBA()
		if i == 0 {
			//TODO: check that first entry is transparent
			continue
		}
		if j > 255 {
			return errors.New("too many colors for palette " + name)
		}
		c := color.SRGBA{
			R: float32(r) / float32(0xFFFF),
			G: float32(g) / float32(0xFFFF),
			B: float32(b) / float32(0xFFFF),
			A: float32(al) / float32(0xFFFF),
		}
		palettes.stdcolors[a][j] = color.LRGBAof(c)
		j++
	}
	palettes.changed[a] = true

	internal.Debug.Printf("Loaded color palette (%d entries) from %s", len(p), name)

	return nil
}

////////////////////////////////////////////////////////////////////////////////

func uploadPalette() error {

	return gl.Err()
}

////////////////////////////////////////////////////////////////////////////////

// Use asks the GPU to use this palette.
//
// Note that the palette will be used for every drawing command of the current
// frame, even those issued before the call to Use. In other words, you cannot
// change the palette in the middle of a frame.
func (a PaletteID) Use() {
	palettes.current = a
	palettes.changed[palettes.current] = true
}

////////////////////////////////////////////////////////////////////////////////

// Clear removes all colors from the palette.
func (a PaletteID) Clear() {
	for c := range palettes.stdcolors[a] {
		palettes.stdcolors[a][c] = debugColor
	}
	palettes.stdcolors[a][0] = color.LRGBA{0, 0, 0, 0}
	palettes.changed[a] = true
}

// Set changes the color associated with an index.
func (a PaletteID) Set(i Color, c color.Color) Color {
	palettes.changed[a] = true
	if c == nil {
		palettes.stdcolors[a][i] = color.LRGBA{1, 0, .5, 1}
	} else {
		palettes.stdcolors[a][i] = color.LRGBAof(c)
	}
	return Color(i)
}

////////////////////////////////////////////////////////////////////////////////

// Match searches for a color by its color.LRGBA values. If this exact color
// isn't in the palette, index 0 is returned.
func (a PaletteID) Match(v color.Color) Color {
	lv := color.LRGBAof(v)
	for c, pv := range palettes.stdcolors[a] {
		if pv == lv {
			return Color(c)
		}
	}

	return Color(0)
}

//TODO: search by color proximity
