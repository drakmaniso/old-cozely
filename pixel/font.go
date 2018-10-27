package pixel

import (
	"errors"
	"image"
	_ "image/png" // Activate PNG support
	"io"

	"github.com/cozely/cozely/color"
	"github.com/cozely/cozely/internal"
)

////////////////////////////////////////////////////////////////////////////////

// FontID is the ID to handle font assets.
type FontID uint8

const noFont = FontID(0) //TODO

const maxFontID = 0xFF

var fonts = struct {
	dictionary map[string]FontID
	name       []string
	height     []int16
	baseline   []int16
	first      []uint16 // index of the first glyph
	image      []*image.Paletted
	lut        []*color.LUT
}{
	dictionary: map[string]FontID{},
}

////////////////////////////////////////////////////////////////////////////////

// Font returns the font ID corresponding to a name.
//
// Should only be called when the framework is running (since resources are
// loaded when the framework starts).
func Font(name string) FontID {
	return fonts.dictionary[name]
}

////////////////////////////////////////////////////////////////////////////////

// newFont creates a new font from an image.
//
// Must be called *before* running the framework.
func newFont(name string, m *image.Paletted, l *color.LUT) FontID {
	var err error

	if internal.Running {
		setErr(errors.New("pixel font declaration: declarations must happen before starting the framework"))
		return noFont
	}

	_, ok := fonts.dictionary[name]
	if ok && name != "" {
		setErr(errors.New(`new font: name "` + name + `" already taken`))
		return noFont
	}

	if len(fonts.name) >= maxFontID {
		setErr(errors.New("pixel font declaration: too many fonts"))
		return noFont
	}

	fonts.name = append(fonts.name, name)
	fonts.height = append(fonts.height, 0)
	fonts.baseline = append(fonts.baseline, 0)
	fonts.first = append(fonts.first, 0)
	fonts.image = append(fonts.image, m)
	fonts.lut = append(fonts.lut, l)
	f := FontID(len(fonts.name) - 1)

	if fonts.lut[f] == nil {
		// Construct the font LUT
		r := m.Bounds()
		r.Min.X++
		r.Min.Y++
		sm, ok := m.SubImage(r).(*image.Paletted)
		if !ok {
			setErr(errors.New("unexpected subimage in Loadfont"))
			return noFont
		}
		var a int
		fonts.lut[f], a, err = color.ToMaster(sm)
		if a != 0 {
			internal.Debug.Printf("WARNING: %d new colors in font "+fonts.name[f], a)
		}
		if err != nil {
			setErr(errors.New("impossible to load font: " + err.Error()))
			return noFont
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
			setErr(errors.New("unexpected subimage in Loadfont"))
			return noFont
		}
		gg := NewPicture("", sm, fonts.lut[f])
		if gg != PictureID(g) {
			//TODO:
			panic("load font: gg != g")
		}
		x += w
		for x < m.Bounds().Dx() && m.Pix[x+h*m.Stride] == 0 {
			x++
		}
	}

	internal.Debug.Printf(
		"Loaded font %s (%d glyphs, %dx%d)",
		fonts.name[f],
		g-fonts.first[f],
		maxw,
		fonts.height[f],
	)

	if name != "" {
		fonts.dictionary[name] = f
	}

	return f
}

////////////////////////////////////////////////////////////////////////////////

// loadFont is the resource handler for fonts.
func loadFont(name string, tags []string, ext string, r io.Reader) error {
	var err error

	if ext != "png" {
		return errors.New(`load font "` + name + `": format "` + ext + `" not supported`)
	}

	m, _, err := image.Decode(r)
	switch err {
	case nil:
	case image.ErrFormat:
		return nil
	default:
		return internal.Wrap("decoding font file", err)
	}

	mm, ok := m.(*image.Paletted)
	if !ok {
		return errors.New("impossible to load font " + name + " (color model not supported)")
	}

	// Check the optional tags
	for _, t := range tags {
		switch t {
		case "meta":
			// ignore, always on
		default:
			setErr(errors.New(`load font "` + name + `": invalid tag`))
		}
	}

	newFont(name, mm, nil)

	return nil
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

//// Copyright (c) 2018-2018 Laurent Moussault. All rights reserved.
//// Licensed under a simplified BSD license (see LICENSE file).
