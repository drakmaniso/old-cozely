package pixel

import (
	"errors"
	"image"
	_ "image/png" // Activate PNG support
	"os"

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

var fonts struct {
	path      []string
	height    []int16
	baseline  []int16
	basecolor []color.Index
	first     []uint16 // index of the first glyph
	image     []*image.Paletted
	lut       []*color.LUT
}

////////////////////////////////////////////////////////////////////////////////

// Font declares a new font and returns its ID.
func Font(path string) FontID {
	return font(path, nil, nil)
}

func font(path string, m *image.Paletted, l *color.LUT) FontID {
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
	fonts.image = append(fonts.image, m)
	fonts.lut = append(fonts.lut, l)
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

	m := fonts.image[f]

	if m == nil {
		fl, err := os.Open(internal.Path + fonts.path[f] + ".png")
		if err != nil {
			return internal.Wrap(`while opening font file "`+fonts.path[f]+`"`, err)
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
		m, ok = img.(*image.Paletted)
		if !ok {
			return errors.New("impossible to load font " + fonts.path[f] + " (color model not supported)")
		}
	}

	if fonts.lut[f] == nil {
		// Construct the font LUT
		sm, ok := m.SubImage(image.Rect(1, 1, m.Bounds().Dx()-2, m.Bounds().Dy()-2)).(*image.Paletted)
		if !ok {
			return errors.New("unexpected subimage in Loadfont")
		}
		var a int
		fonts.lut[f], a, err = color.ToMaster(sm)
		if a != 0 {
			internal.Debug.Printf("Warning: %d new colors in font "+fonts.path[f], a)
		}
		if err != nil {
			return errors.New("impossible to load font: " + err.Error())
		}
	}

	h := m.Bounds().Dy() - 1
	fonts.height[f] = int16(h)
	g := uint16(len(pictures.mapping))
	fonts.first[f] = g
	maxw := 0

	// Find the baseline
	for y := 0; y < m.Bounds().Dy(); y++ {
		if m.Pix[0+y*m.Stride] != 0 {
			fonts.baseline[f] = int16(y)
			break
		}
	}

	// Find the basecolor
	fonts.basecolor[f] = 255
	for y := 0; y < m.Bounds().Dy()-1; y++ {
		for x := 1; x < m.Bounds().Dx(); x++ {
			c := color.Index(m.Pix[x+y*m.Stride])
			lc := fonts.lut[f][c]
			if lc != 0 && c < fonts.basecolor[f] {
				fonts.basecolor[f] = lc
			}
		}
	}

	// Create images and reserve mapping for each rune
	for x := 1; x < m.Bounds().Dx(); g++ {
		w := 0
		for x+w < m.Bounds().Dx() && m.Pix[x+w+h*m.Stride] != 0 {
			w++
		}
		if w > maxw {
			maxw = w
		}
		sm, ok := m.SubImage(image.Rect(x, 0, x+w, h)).(*image.Paletted)
		if !ok {
			return errors.New("unexpected subimage in Loadfont")
		}
		gg := picture("(font)", sm, fonts.lut[f])
		if gg != PictureID(g) {
			//TODO:
		}
		x += w
		for x < m.Bounds().Dx() && m.Pix[x+h*m.Stride] == 0 {
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

//// Copyright (c) 2018-2018 Laurent Moussault. All rights reserved.
//// Licensed under a simplified BSD license (see LICENSE file).
