// Copyright (c) 2013-2017 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package colour

//------------------------------------------------------------------------------

type Index uint8

type Palette uint8

var palettes []paletteData

type paletteData struct {
	colours []RGBA
}

//------------------------------------------------------------------------------

func PaletteCount() uint8 {
	return uint8(len(palettes))
}

//------------------------------------------------------------------------------

func NewPalette(c []RGBA) Palette {
	cc := make([]RGBA, len(c), len(c))
	copy(cc, c)
	p := paletteData{colours: cc}
	palettes = append(palettes, p)
	//TODO: register for GPU
	//TODO: error handling?
	return Palette(len(palettes) - 1)
}

//------------------------------------------------------------------------------

func (p Palette) GetRGBA(i Index) (col RGBA, ok bool) {
	if int(p) < len(palettes) && int(i) < len(palettes[p].colours) {
		return palettes[p].colours[i], true
	}
	return RGBA{}, false
}

func FindIndex(colour RGBA) (i Index, ok bool) {
	return Index(0), false
}

//------------------------------------------------------------------------------
