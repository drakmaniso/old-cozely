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
	maxFontID = 0xFF
	noFont    = FontID(maxFontID)
)

// Monozela10 is the ID of the default font (a 10 pixel high monospace). This is
// the only font that is always loaded and doesn't need declaration.
const Monozela10 = FontID(0)

var fontPaths = []string{"builtin monozela 10"}

var fonts = []font{{}}

type font struct {
	height    int16
	baseline  int16
	basecolor color.Index
	first     uint16 // index of the first glyph
}

////////////////////////////////////////////////////////////////////////////////

// Font declares a new font and returns its ID.
func Font(path string) FontID {
	if internal.Running {
		setErr(errors.New("pixel font declaration: declarations must happen before starting the framework"))
		return noFont
	}

	if len(fonts) >= maxFontID {
		setErr(errors.New("pixel font declaration: too many fonts"))
		return noFont
	}

	fonts = append(fonts, font{})
	fontPaths = append(fontPaths, path)
	return FontID(len(fonts) - 1)
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
	return fonts[f].first + uint16(r)
}

////////////////////////////////////////////////////////////////////////////////

// Height returns the height of the font, i.e. the height of the images used to
// store the glyphs.
func (f FontID) Height() int16 {
	return fonts[f].height
}

////////////////////////////////////////////////////////////////////////////////

////////////////////////////////////////////////////////////////////////////////

func (f FontID) load(frects *[]uint32) error {
	//TODO: support other image formats?
	var p *image.Paletted

	if f == 0 {
		p = &monozela10
	} else {
		n := fontPaths[f]
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

	h := p.Bounds().Dy() - 1
	fonts[f].height = int16(h)
	g := uint16(len(pictures.mapping))
	fonts[f].first = g
	maxw := 0

	for y := 0; y < p.Bounds().Dy(); y++ {
		if p.Pix[0+y*p.Stride] != 0 {
			fonts[f].baseline = int16(y)
			break
		}
	}

	fonts[f].basecolor = 255
	for _, c := range p.Pix {
		if c != 0 && color.Index(c) < fonts[f].basecolor {
			fonts[f].basecolor = color.Index(c)
		}
	}

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
		x += w
		for x < p.Bounds().Dx() && p.Pix[x+h*p.Stride] == 0 {
			x++
		}
	}

	internal.Debug.Printf(
		"Loaded font %s (%d glyphs, %dx%d)",
		fontPaths[f],
		g-fonts[f].first,
		maxw,
		fonts[f].height,
	)

	return nil
}
