// Copyright (c) 2018-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package pixel

import (
	"errors"
	"image"
	_ "image/png" // Activate PNG support
	"os"
	"path/filepath"

	"github.com/cozely/cozely/color"
	"github.com/cozely/cozely/internal"
)

////////////////////////////////////////////////////////////////////////////////

// FontID is the ID to handle font assets.
type FontID uint8

const (
	// Monozela10 is the default font (10 pixel high, monospace). This is the only
	// font that is always loaded and doesn't need declaration.
	Monozela10 = FontID(0)

	noFont = FontID(maxFontID)
)

const maxFontID = 0xFF

var fonts = struct {
	path      []string
	height    []int16
	baseline  []int16
	basecolor []color.Index
	first     []uint16 // index of the first glyph
	lut       []color.LUT
}{
	path:      []string{"builtin monozela 10"},
	height:    []int16{0},
	baseline:  []int16{0},
	basecolor: []color.Index{0},
	first:     []uint16{0},
	lut:       []color.LUT{initLUT},
}

////////////////////////////////////////////////////////////////////////////////

// Font declares a new font and returns its ID.
func Font(path string) FontID {
	if internal.Running {
		setErr(errors.New("pixel font declaration: declarations must happen before starting the framework"))
		return noFont
	}

	if len(fonts.path) >= maxFontID {
		setErr(errors.New("pixel font declaration: too many fonts"))
		return noFont
	}

	fonts.path = append(fonts.path, path)
	fonts.height = append(fonts.height, 0)
	fonts.baseline = append(fonts.baseline, 0)
	fonts.basecolor = append(fonts.basecolor, 0)
	fonts.first = append(fonts.first, 0)
	fonts.lut = append(fonts.lut, initLUT)
	return FontID(len(fonts.path) - 1)
}

////////////////////////////////////////////////////////////////////////////////

func (f FontID) glyph(r rune) uint16 {
	//TODO: add support for non-ascii runes
	switch {
	case r < ' ':
		r = 0x7F - ' '
	case r <= 0x7F:
		r = r - ' '
	default:
		r = 0x7F - ' '
	}
	return fonts.first[f] + uint16(r)
}

////////////////////////////////////////////////////////////////////////////////

// Height returns the height of the font, i.e. the height of the images used to
// store the glyphs.
func (f FontID) Height() int16 {
	return fonts.height[f]
}

// Interline returns the default interline of the font, i.e. the vertical
// distance between the baselines of two consecutive lines.
func (f FontID) Interline() int16 {
	return int16(float32(fonts.height[f]) * 1.25)
}

////////////////////////////////////////////////////////////////////////////////

func (f FontID) load(frects *[]uint32) error {
	var err error
	//TODO: support other image formats?
	var p *image.Paletted

	if f == 0 {
		p = &monozela10
	} else {
		n := fonts.path[f]
		path := filepath.FromSlash(internal.Path + n + ".png")
		path, err := filepath.EvalSymlinks(path)
		if err != nil {
			return internal.Wrap("in path while loading font", err)
		}

		fl, err := os.Open(path)
		if err != nil {
			return internal.Wrap(`while opening font file "`+path+`"`, err)
		}
		defer fl.Close() //TODO: error handling

		img, _, err := image.Decode(fl)
		switch err {
		case nil:
		case image.ErrFormat:
			return nil
		default:
			return internal.Wrap("decoding font file", err)
		}

		var ok bool
		p, ok = img.(*image.Paletted)
		if !ok {
			return errors.New("impossible to load font " + path + " (color model not supported)")
		}
	}

	fonts.lut[f], err = color.ToMaster(p)
	if err != nil {
		return errors.New("impossible to load font: " + err.Error())
	}

	h := p.Bounds().Dy() - 1
	fonts.height[f] = int16(h)
	g := uint16(len(pictures.mapping))
	fonts.first[f] = g
	maxw := 0

	for y := 0; y < p.Bounds().Dy(); y++ {
		if p.Pix[0+y*p.Stride] != 0 {
			fonts.baseline[f] = int16(y)
			break
		}
	}

	fonts.basecolor[f] = 255
	// for _, c := range p.Pix {
	// 	lc := fonts.lut[f][c]
	// 	if c != 0 &&  lc < fonts.basecolor[f] {
	// 		fonts.basecolor[f] = lc
	// 	}
	// }

	// Create images and reserve mapping for each rune

	for x := 1; x < p.Bounds().Dx(); g++ {
		w := 0
		for x+w < p.Bounds().Dx() && p.Pix[x+w+h*p.Stride] != 0 {
			w++
		}
		if w > maxw {
			maxw = w
		}
		m := p.SubImage(image.Rect(x, 0, x+w, h))
		mm, ok := m.(*image.Paletted)
		if !ok {
			return errors.New("unexpected subimage in Loadfont")
		}
		gg := picture(mm)
		if gg != PictureID(g) {
			//TODO:
		}
		pictures.lut[gg] = fonts.lut[f]
		x += w
		for x < p.Bounds().Dx() && p.Pix[x+h*p.Stride] == 0 {
			x++
		}
	}

	internal.Debug.Printf(
		"Loaded font %s (%d glyphs, %dx%d)",
		fonts.path[f],
		g-fonts.first[f],
		maxw,
		fonts.height[f],
	)

	return nil
}
