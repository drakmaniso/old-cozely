// Copyright (c) 2018-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package color

import (
	"errors"
	"image"
	stdcolor "image/color"
	_ "image/png" // Activate PNG support
	"os"

	"github.com/cozely/cozely/internal"
)

////////////////////////////////////////////////////////////////////////////////

func (a PaletteID) load(name string) error {
	f, err := os.Open(internal.Path + name + ".png")
	if err != nil {
		return errors.New("unable to open file for palette " + name)
	}
	cf, _, err := image.DecodeConfig(f)
	if err != nil {
		return errors.New("unable to decode file for palette " + name)
	}

	p, ok := cf.ColorModel.(stdcolor.Palette)
	if !ok {
		return errors.New("image file not paletted for palette " + name)
	}

	palettes.colours[a][0] = LRGBA{0, 0, 0, 0}
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
		c := SRGBA{
			R: float32(r) / float32(0xFFFF),
			G: float32(g) / float32(0xFFFF),
			B: float32(b) / float32(0xFFFF),
			A: float32(al) / float32(0xFFFF),
		}
		palettes.colours[a][j] = c
		j++
	}
	palettes.changed[a] = true

	internal.Debug.Printf("Loaded color palette (%d entries) from %s", len(p), name)

	return nil
}

////////////////////////////////////////////////////////////////////////////////
