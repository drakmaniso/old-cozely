// Copyright (c) 2018-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package pixel

import (
	"errors"
	"image"
	_ "image/png" // Activate PNG support
	"os"
	"path/filepath"

	"github.com/drakmaniso/glam/internal"
	"github.com/drakmaniso/glam/x/atlas"
)

//------------------------------------------------------------------------------

var (
	fntAtlas  *atlas.Atlas
	fntImages []*image.Paletted
)

//------------------------------------------------------------------------------

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
			return internal.Error("in path while loading font", err)
		}

		fl, err := os.Open(path)
		if err != nil {
			return internal.Error(`while opening font file "`+path+`"`, err)
		}
		defer fl.Close() //TODO: error handling

		img, _, err := image.Decode(fl)
		switch err {
		case nil:
		case image.ErrFormat:
			return nil
		default:
			return internal.Error("decoding font file", err)
		}

		var ok bool
		p, ok = img.(*image.Paletted)
		if !ok {
			return errors.New("impossible to load font " + path + " (color model not supported)")
		}
	}

	h := p.Bounds().Dy() - 1
	fonts[f].height = int16(h)
	g := uint16(len(glyphMap))
	fonts[f].first = g
	maxw := 0

	for y := 0; y < p.Bounds().Dy(); y++ {
		if p.Pix[0+y*p.Stride] != 0 {
			fonts[f].baseline = int16(y)
			break
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
		glyphMap = append(glyphMap, mapping{w: int16(w), h: int16(h)})
		*frects = append(*frects, uint32(g))
		fntImages = append(fntImages, mm)
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

//------------------------------------------------------------------------------

func fntSize(rect uint32) (width, height int16) {
	w, h := fntImages[rect].Bounds().Dx(), fntImages[rect].Bounds().Dy()
	return int16(w), int16(h)
}

func fntPut(rect uint32, bin int16, x, y int16) {
	glyphMap[rect].bin = bin
	glyphMap[rect].x = x
	glyphMap[rect].y = y
}

func fntPaint(rect uint32, dest interface{}) error {
	fx, fy := glyphMap[rect].x, glyphMap[rect].y
	fw, fh := glyphMap[rect].w, glyphMap[rect].h

	dm, ok := dest.(*image.Paletted)
	if !ok {
		return errors.New("unexpected dest argument to fntrune paint method")
	}
	for y := 0; y < int(fh); y++ {
		for x := 0; x < int(fw); x++ {
			w := dm.Bounds().Dx()
			ci := fntImages[rect].Pix[x+y*fntImages[rect].Stride]
			dm.Pix[int(fx)+x+w*(int(fy)+y)] = uint8(ci)
		}
	}

	return nil
}

//------------------------------------------------------------------------------
