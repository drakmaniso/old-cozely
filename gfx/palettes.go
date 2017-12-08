// Copyright (c) 2013-2017 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package gfx

import "image/color"

//------------------------------------------------------------------------------

func PaletteMSX2() {
	ClearPalette()
	for i := 1; i < 256; i++ {
		NewColor("", RGBA{
			float32(i>>5) / 7.0,
			float32((i&0x1C)>>2) / 7.0,
			float32(i&0x3) / 3.0,
			1.0,
		})
	}
}

//------------------------------------------------------------------------------

func PaletteMSX() {
	ClearPalette()
	NewColor("msx: black", color.NRGBA{0x00, 0x00, 0x00, 0xFF})
	NewColor("msx: medium green", color.NRGBA{0x3E, 0xB8, 0x49, 0xFF})
	NewColor("msx: light green", color.NRGBA{0x74, 0xd0, 0x7d, 0xFF})
	NewColor("msx: dark blue", color.NRGBA{0x59, 0x55, 0xe0, 0xFF})
	NewColor("msx: light blue", color.NRGBA{0x80, 0x76, 0xf1, 0xFF})
	NewColor("msx: dark red", color.NRGBA{0xb9, 0x5e, 0x51, 0xFF})
	NewColor("msx: cyan", color.NRGBA{0x65, 0xdb, 0xef, 0xFF})
	NewColor("msx: medium red", color.NRGBA{0xdb, 0x65, 0x59, 0xFF})
	NewColor("msx: light red", color.NRGBA{0xff, 0x89, 0x7d, 0xFF})
	NewColor("msx: dark yellow", color.NRGBA{0xcc, 0xc3, 0x5e, 0xFF})
	NewColor("msx: light yellow", color.NRGBA{0xde, 0xd0, 0x87, 0xFF})
	NewColor("msx: dark green", color.NRGBA{0x3a, 0xa2, 0x41, 0xFF})
	NewColor("msx: magenta", color.NRGBA{0xb7, 0x66, 0xb5, 0xFF})
	NewColor("msx: gray", color.NRGBA{0xcc, 0xcc, 0xcc, 0xFF})
	NewColor("msx: white", color.NRGBA{0xff, 0xff, 0xff, 0xFF})
}

//------------------------------------------------------------------------------

func PaletteCPC() {
	ClearPalette()
	NewColor("cpc: black", color.NRGBA{0x00, 0x00, 0x00, 0xFF})
	NewColor("cpc: blue", color.NRGBA{0x00, 0x00, 0x80, 0xFF})
	NewColor("cpc: bright blue", color.NRGBA{0x00, 0x00, 0xff, 0xFF})
	NewColor("cpc: red", color.NRGBA{0x80, 0x00, 0x00, 0xFF})
	NewColor("cpc: magenta", color.NRGBA{0x80, 0x00, 0x80, 0xFF})
	NewColor("cpc: mauve", color.NRGBA{0x80, 0x00, 0xff, 0xFF})
	NewColor("cpc: bright red", color.NRGBA{0xff, 0x00, 0x00, 0xFF})
	NewColor("cpc: purple", color.NRGBA{0xff, 0x00, 0x80, 0xFF})
	NewColor("cpc: bright magenta", color.NRGBA{0xff, 0x00, 0xff, 0xFF})
	NewColor("cpc: green", color.NRGBA{0x00, 0x80, 0x00, 0xFF})
	NewColor("cpc: cyan", color.NRGBA{0x00, 0x80, 0x80, 0xFF})
	NewColor("cpc: sky blue", color.NRGBA{0x00, 0x80, 0xff, 0xFF})
	NewColor("cpc: yellow", color.NRGBA{0x80, 0x80, 0x00, 0xFF})
	NewColor("cpc: white", color.NRGBA{0x80, 0x80, 0x80, 0xFF})
	NewColor("cpc: pastel blue", color.NRGBA{0x80, 0x80, 0xff, 0xFF})
	NewColor("cpc: orange", color.NRGBA{0xff, 0x80, 0x00, 0xFF})
	NewColor("cpc: pink", color.NRGBA{0xff, 0x80, 0x80, 0xFF})
	NewColor("cpc: pastel magenta", color.NRGBA{0xff, 0x80, 0xff, 0xFF})
	NewColor("cpc: bright green", color.NRGBA{0x00, 0xff, 0x00, 0xFF})
	NewColor("cpc: sea green", color.NRGBA{0x00, 0xff, 0x80, 0xFF})
	NewColor("cpc: bright cyan", color.NRGBA{0x00, 0xff, 0xff, 0xFF})
	NewColor("cpc: lime", color.NRGBA{0x80, 0xff, 0x00, 0xFF})
	NewColor("cpc: pastel green", color.NRGBA{0x80, 0xff, 0x80, 0xFF})
	NewColor("cpc: pastel cyan", color.NRGBA{0x80, 0xff, 0xff, 0xFF})
	NewColor("cpc: bright yellow", color.NRGBA{0xff, 0xff, 0x00, 0xFF})
	NewColor("cpc: pastel yellow", color.NRGBA{0xff, 0xff, 0x80, 0xFF})
	NewColor("cpc: bright white", color.NRGBA{0xff, 0xff, 0xff, 0xFF})
}

//------------------------------------------------------------------------------

// http://unusedino.de/ec64/technical/misc/vic656x/colors/
//
// http://www.pepto.de/projects/colorvic/
func PaletteC64() {
	ClearPalette()
	NewColor("c64: black", color.NRGBA{0x00, 0x00, 0x00, 0xFF})
	NewColor("c64: white", color.NRGBA{0xff, 0xff, 0xff, 0xFF})
	NewColor("c64: red", color.NRGBA{0x68, 0x37, 0x2b, 0xFF})
	NewColor("c64: cyan", color.NRGBA{0x70, 0xa4, 0xb2, 0xFF})
	NewColor("c64: violet", color.NRGBA{0x6f, 0x3d, 0x86, 0xFF})
	NewColor("c64: green", color.NRGBA{0x58, 0x8d, 0x43, 0xFF})
	NewColor("c64: blue", color.NRGBA{0x35, 0x28, 0x79, 0xFF})
	NewColor("c64: yellow", color.NRGBA{0xb8, 0xc7, 0x6f, 0xFF})
	NewColor("c64: orange", color.NRGBA{0x6f, 0x4f, 0x25, 0xFF})
	NewColor("c64: brown", color.NRGBA{0x43, 0x39, 0x00, 0xFF})
	NewColor("c64: light red", color.NRGBA{0x9a, 0x67, 0x59, 0xFF})
	NewColor("c64: dark grey", color.NRGBA{0x44, 0x44, 0x44, 0xFF})
	NewColor("c64: grey", color.NRGBA{0x6c, 0x6c, 0x6c, 0xFF})
	NewColor("c64: light green", color.NRGBA{0x9a, 0xd2, 0x84, 0xFF})
	NewColor("c64: light blue", color.NRGBA{0x6c, 0x5e, 0xb5, 0xFF})
	NewColor("c64: light grey", color.NRGBA{0x95, 0x95, 0x95, 0xFF})
}

//------------------------------------------------------------------------------
