// Copyright (c) 2018-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package palette

import (
	"errors"
	"image"
	"image/color"
	_ "image/png" // Activate PNG support
	"os"
	"strconv"

	"github.com/drakmaniso/glam/colour"
	"github.com/drakmaniso/glam/internal"
)

//------------------------------------------------------------------------------

// Load clears the palette and load a new one. The only format currently
// supported is indexed PNG files. Note that once loaded, the palette is cached,
// so only the first call to Load will read the file. Several names are
// predefined: "MSX", "MSX2", "CPC", "C64".
func Load(name string) error {
	p, ok := palettes[name]
	if ok {
		Clear()
		for i := range p {
			Entry(uint8(i+1), p[i].colour)
		}
		return nil
	}

	return loadFile(name)
}

//------------------------------------------------------------------------------

func loadFile(name string) error {
	f, err := os.Open(internal.Path + name + ".png")
	if err != nil {
		return errors.New("unable to open file for palette " + name)
	}
	cf, _, err := image.DecodeConfig(f)
	if err != nil {
		return errors.New("unable to decode file for palette " + name)
	}

	p, ok := cf.ColorModel.(color.Palette)
	if !ok {
		return errors.New("image file not paletted for palette " + name)
	}

	Clear()
	pal := make([]struct {
		name   string
		colour colour.Colour
	},
		len(p)-1,
		len(p)-1)
	j := 1
	for i := range p {
		r, g, b, _ := p[i].RGBA()
		if i == 0 {
			//TODO: check that first entry is transparent
			continue
		}
		if j > 255 {
			return errors.New("too many colors for palette " + name)
		}
		c := colour.SRGBA{
			R: float32(r) / float32(0xFFFF),
			G: float32(g) / float32(0xFFFF),
			B: float32(b) / float32(0xFFFF),
			A: 1,
		}
		n := "png" + strconv.Itoa(j)
		pal[j-1].name = n
		pal[j-1].colour = c
		Entry(uint8(j), c)
		j++
	}
	palettes[name] = pal

	internal.Debug.Printf("Loaded color palette (%d entries) from %s", len(p), name)

	return nil
}

//------------------------------------------------------------------------------
