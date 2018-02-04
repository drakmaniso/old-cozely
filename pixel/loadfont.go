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
	fntFiles []atlas.Image
	fntAtlas *atlas.Atlas
)

//------------------------------------------------------------------------------

func loadFont(n string, f Font) error {
	//TODO: support other image formats?
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

	p, ok := img.(*image.Paletted)
	if !ok {
		internal.Debug.Println(`ignoring image: "` + path + `" (color model not supported)`)
		return nil
	}
	gw, gh := p.Bounds().Dx()/16, p.Bounds().Dy()/8

	h := gh - 1
	fonts[f].height = int16(h)
	gly := uint16(len(glyphMap))
	fonts[f].first = gly

	// Create images and reserve mapping for each rune

	for gy := 0; gy < 8; gy++ {
		for gx := 0; gx < 16; gx++ {
			w := 0
			for w < gw && p.Pix[gx*gw+w+(gy*gh+h)*p.Stride] != 0 {
				w++
			}
			m := p.SubImage(image.Rect(gx*gw, gy*gh, gx*gw+w, gy*gh+h))
			mm, ok := m.(*image.Paletted)
			if !ok {
				return errors.New("unexpected subimage in Loadfont")
			}
			glyphMap = append(glyphMap, mapping{w: int16(w), h: int16(h)})
			fntFiles = append(
				fntFiles,
				fntrune{
					glyph: gly + uint16(gx+16*gy),
					img:   mm,
				},
			)

		}
	}

	// Pack them into the atlas

	fntAtlas.Pack(fntFiles)

	return nil
}

//------------------------------------------------------------------------------

type fntrune struct {
	glyph uint16
	img   *image.Paletted
}

func (fr fntrune) Size() (width, height int16) {
	w, h := fr.img.Bounds().Dx(), fr.img.Bounds().Dy()
	return int16(w), int16(h)
}

func (fr fntrune) Put(bin int16, x, y int16) {
	glyphMap[fr.glyph].bin = bin
	glyphMap[fr.glyph].x = x
	glyphMap[fr.glyph].y = y
}

func (fr fntrune) Paint(dest interface{}) error {
	fx, fy := glyphMap[fr.glyph].x, glyphMap[fr.glyph].y
	fw, fh := glyphMap[fr.glyph].w, glyphMap[fr.glyph].h

	dm, ok := dest.(*image.Paletted)
	if !ok {
		return errors.New("unexpected dest argument to fntrune paint method")
	}
	for y := 0; y < int(fh); y++ {
		for x := 0; x < int(fw); x++ {
			w := dm.Bounds().Dx()
			ci := fr.img.Pix[x+y*fr.img.Stride]
			dm.Pix[int(fx)+x+w*(int(fy)+y)] = uint8(ci)
		}
	}

	return nil
}

//------------------------------------------------------------------------------
