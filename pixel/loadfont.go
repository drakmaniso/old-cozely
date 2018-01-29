// Copyright (c) 2018-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package pixel

import (
	"errors"
	"image"
	_ "image/png" // Activate PNG support
	"os"
	"path/filepath"
	"strings"

	"github.com/drakmaniso/glam/internal"
	"github.com/drakmaniso/glam/x/atlas"
)

//------------------------------------------------------------------------------

var (
	fntFiles []atlas.Image
	fntAtlas *atlas.Atlas
)

//------------------------------------------------------------------------------

func LoadFont(path string) error {
	if internal.Running {
		return errors.New("loading fonts while running not implemented")
	}

	f, err := os.Open(path)
	if err != nil {
		return internal.Error(`while opening font file "`+path+`"`, err)
	}
	defer f.Close() //TODO: error handling

	img, _, err := image.Decode(f)
	switch err {
	case nil:
	case image.ErrFormat:
		return nil
	default:
		return internal.Error("decoding font file", err)
	}

	n := filepath.Base(path)
	if err != nil {
		return err
	}
	n = strings.TrimSuffix(n, filepath.Ext(n))
	n = filepath.ToSlash(n)

	p, ok := img.(*image.Paletted)
	if !ok {
		internal.Debug.Println(`ignoring image: "` + path + `" (color model not supported)`)
		return nil
	}
	gw, gh := p.Bounds().Dx()/16, p.Bounds().Dy()/8

	h := gh - 1
	fnt := newFont(n, int16(h), 0)

	// Create images for each rune

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
			fontMap[fnt].chars[gx+16*gy].w = int16(w)
			fntFiles = append(
				fntFiles,
				fntrune{
					font: fnt,
					char: rune(gx + 16*gy),
					img:  mm,
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
	font Font
	char rune
	img  *image.Paletted
}

func (fr fntrune) Size() (width, height int16) {
	w, h := fr.img.Bounds().Dx(), fr.img.Bounds().Dy()
	return int16(w), int16(h)
}

func (fr fntrune) Put(bin int16, x, y int16) {
	fontMap[fr.font].chars[fr.char].bin = bin
	fontMap[fr.font].chars[fr.char].x = x
	fontMap[fr.font].chars[fr.char].y = y
}

func (fr fntrune) Paint(dest interface{}) error {
	_, fx, fy, fw, fh := fr.font.getMap(fr.char)

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
