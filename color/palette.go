// Copyright (c) 2013-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package color

import (
	"image"
	stdcolor "image/color"
	"os"

	"github.com/cozely/cozely/internal"
)

////////////////////////////////////////////////////////////////////////////////

type Palette struct {
	Names  map[string]Index
	Colors []LRGBA
}

type Index uint8

////////////////////////////////////////////////////////////////////////////////

func PaletteFrom(name string) Palette {
	var pal = Palette{
		Names: map[string]Index{},
	}

	f, err := os.Open(internal.Path + name + ".png")
	if err != nil {
		//TODO: errors.New("unable to open file for palette " + name)
		return pal
	}
	defer f.Close() //TODO: error handling
	cf, _, err := image.DecodeConfig(f)
	if err != nil {
		//TODO: errors.New("unable to decode file for palette " + name)
		return pal
	}

	p, ok := cf.ColorModel.(stdcolor.Palette)
	if !ok {
		//TODO: errors.New("image file not paletted for palette " + name)
		return pal
	}

	//TODO: clear the palette?

	for i := range p {
		r, g, b, al := p[i].RGBA()
		if i == 0 {
			//TODO: check if first entry is transparent
			continue
		}
		if i > 255 {
			//TODO:errors.New("too many colors for palette " + name)
			return pal
		}
		c := SRGBA{
			R: float32(r) / float32(0xFFFF),
			G: float32(g) / float32(0xFFFF),
			B: float32(b) / float32(0xFFFF),
			A: float32(al) / float32(0xFFFF),
		}
		//TODO: append name
		pal.Colors = append(pal.Colors, LRGBAof(c))
	}

	internal.Debug.Printf("Loaded color palette (%d entries) from %s", len(p), name)

	return pal
}
