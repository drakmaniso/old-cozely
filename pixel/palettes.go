// Copyright (c) 2013-2017 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package pixel

//------------------------------------------------------------------------------

import (
	"github.com/drakmaniso/glam/colour"
)

//------------------------------------------------------------------------------

func PaletteMSX2() {
	ClearPalette()
	for i := 1; i < 256; i++ {
		NewColor("", colour.RGBA{
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
	NewColor("msx: black", colour.SRGB8{0x00, 0x00, 0x00})
	NewColor("msx: medium green", colour.SRGB8{0x3E, 0xB8, 0x49})
	NewColor("msx: light green", colour.SRGB8{0x74, 0xd0, 0x7d})
	NewColor("msx: dark blue", colour.SRGB8{0x59, 0x55, 0xe0})
	NewColor("msx: light blue", colour.SRGB8{0x80, 0x76, 0xf1})
	NewColor("msx: dark red", colour.SRGB8{0xb9, 0x5e, 0x51})
	NewColor("msx: cyan", colour.SRGB8{0x65, 0xdb, 0xef})
	NewColor("msx: medium red", colour.SRGB8{0xdb, 0x65, 0x59})
	NewColor("msx: light red", colour.SRGB8{0xff, 0x89, 0x7d})
	NewColor("msx: dark yellow", colour.SRGB8{0xcc, 0xc3, 0x5e})
	NewColor("msx: light yellow", colour.SRGB8{0xde, 0xd0, 0x87})
	NewColor("msx: dark green", colour.SRGB8{0x3a, 0xa2, 0x41})
	NewColor("msx: magenta", colour.SRGB8{0xb7, 0x66, 0xb5})
	NewColor("msx: gray", colour.SRGB8{0xcc, 0xcc, 0xcc})
	NewColor("msx: white", colour.SRGB8{0xff, 0xff, 0xff})
}

//------------------------------------------------------------------------------

func PaletteCPC() {
	ClearPalette()
	NewColor("cpc: black", colour.SRGB8{0x00, 0x00, 0x00})
	NewColor("cpc: blue", colour.SRGB8{0x00, 0x00, 0x80})
	NewColor("cpc: bright blue", colour.SRGB8{0x00, 0x00, 0xff})
	NewColor("cpc: red", colour.SRGB8{0x80, 0x00, 0x00})
	NewColor("cpc: magenta", colour.SRGB8{0x80, 0x00, 0x80})
	NewColor("cpc: mauve", colour.SRGB8{0x80, 0x00, 0xff})
	NewColor("cpc: bright red", colour.SRGB8{0xff, 0x00, 0x00})
	NewColor("cpc: purple", colour.SRGB8{0xff, 0x00, 0x80})
	NewColor("cpc: bright magenta", colour.SRGB8{0xff, 0x00, 0xff})
	NewColor("cpc: green", colour.SRGB8{0x00, 0x80, 0x00})
	NewColor("cpc: cyan", colour.SRGB8{0x00, 0x80, 0x80})
	NewColor("cpc: sky blue", colour.SRGB8{0x00, 0x80, 0xff})
	NewColor("cpc: yellow", colour.SRGB8{0x80, 0x80, 0x00})
	NewColor("cpc: white", colour.SRGB8{0x80, 0x80, 0x80})
	NewColor("cpc: pastel blue", colour.SRGB8{0x80, 0x80, 0xff})
	NewColor("cpc: orange", colour.SRGB8{0xff, 0x80, 0x00})
	NewColor("cpc: pink", colour.SRGB8{0xff, 0x80, 0x80})
	NewColor("cpc: pastel magenta", colour.SRGB8{0xff, 0x80, 0xff})
	NewColor("cpc: bright green", colour.SRGB8{0x00, 0xff, 0x00})
	NewColor("cpc: sea green", colour.SRGB8{0x00, 0xff, 0x80})
	NewColor("cpc: bright cyan", colour.SRGB8{0x00, 0xff, 0xff})
	NewColor("cpc: lime", colour.SRGB8{0x80, 0xff, 0x00})
	NewColor("cpc: pastel green", colour.SRGB8{0x80, 0xff, 0x80})
	NewColor("cpc: pastel cyan", colour.SRGB8{0x80, 0xff, 0xff})
	NewColor("cpc: bright yellow", colour.SRGB8{0xff, 0xff, 0x00})
	NewColor("cpc: pastel yellow", colour.SRGB8{0xff, 0xff, 0x80})
	NewColor("cpc: bright white", colour.SRGB8{0xff, 0xff, 0xff})
}

//------------------------------------------------------------------------------

// http://unusedino.de/ec64/technical/misc/vic656x/colors/
//
// http://www.pepto.de/projects/colorvic/
func PaletteC64() {
	ClearPalette()
	NewColor("c64: black", colour.SRGB8{0x00, 0x00, 0x00})
	NewColor("c64: white", colour.SRGB8{0xff, 0xff, 0xff})
	NewColor("c64: red", colour.SRGB8{0x68, 0x37, 0x2b})
	NewColor("c64: cyan", colour.SRGB8{0x70, 0xa4, 0xb2})
	NewColor("c64: violet", colour.SRGB8{0x6f, 0x3d, 0x86})
	NewColor("c64: green", colour.SRGB8{0x58, 0x8d, 0x43})
	NewColor("c64: blue", colour.SRGB8{0x35, 0x28, 0x79})
	NewColor("c64: yellow", colour.SRGB8{0xb8, 0xc7, 0x6f})
	NewColor("c64: orange", colour.SRGB8{0x6f, 0x4f, 0x25})
	NewColor("c64: brown", colour.SRGB8{0x43, 0x39, 0x00})
	NewColor("c64: light red", colour.SRGB8{0x9a, 0x67, 0x59})
	NewColor("c64: dark grey", colour.SRGB8{0x44, 0x44, 0x44})
	NewColor("c64: grey", colour.SRGB8{0x6c, 0x6c, 0x6c})
	NewColor("c64: light green", colour.SRGB8{0x9a, 0xd2, 0x84})
	NewColor("c64: light blue", colour.SRGB8{0x6c, 0x5e, 0xb5})
	NewColor("c64: light grey", colour.SRGB8{0x95, 0x95, 0x95})
}

//------------------------------------------------------------------------------
