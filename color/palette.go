// Copyright (c) 2013-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package color

import (
	"errors"
	"image"
	stdcolor "image/color"
	"os"

	"github.com/cozely/cozely/internal"
)

////////////////////////////////////////////////////////////////////////////////

// A Palette is an ordered list of colors.
type Palette []Color

////////////////////////////////////////////////////////////////////////////////

// PaletteFrom returns a new Palette created from the file at the specified
// path.
func PaletteFrom(path string) (Palette, error) {
	var pal = Palette{}

	f, err := os.Open(internal.Path + path + ".png")
	if err != nil {
		return pal, errors.New("unable to open file for palette " + path)
	}
	defer f.Close() //TODO: error handling
	cf, _, err := image.DecodeConfig(f)
	if err != nil {
		return pal, errors.New("unable to decode file for palette " + path)
	}

	p, ok := cf.ColorModel.(stdcolor.Palette)
	if !ok {
		return pal, errors.New("image file not paletted for palette " + path)
	}

	for i := range p {
		r, g, b, al := p[i].RGBA()
		if i > 255 {
			return pal, errors.New("too many colors for palette " + path)
		}
		c := SRGBA{
			R: float32(r) / float32(0xFFFF),
			G: float32(g) / float32(0xFFFF),
			B: float32(b) / float32(0xFFFF),
			A: float32(al) / float32(0xFFFF),
		}
		//TODO: append name
		pal = append(pal, c)
	}

	internal.Debug.Printf("Loaded color palette (%d entries) from %s", len(p), path)

	return pal, nil
}
