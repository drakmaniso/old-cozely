// Copyright (c) 2013-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package palette

import (
	"errors"

	"github.com/drakmaniso/glam/colour"
)

//------------------------------------------------------------------------------

// Change clears the palette and replaces it by the palette corresponding to
// name.
//
// There is a few predefinde names: "MSX", "MSX2", "CPC", and "C64"
func Change(name string) error {
	switch name {

	case "MSX2":
		Clear()
		for i := 1; i < 256; i++ {
			New("", colour.LRGBA{
				float32(i>>5) / 7.0,
				float32((i&0x1C)>>2) / 7.0,
				float32(i&0x3) / 3.0,
				1.0,
			})
		}

	case "MSX":
		Clear()
		New("msx: black", colour.SRGB8{0x00, 0x00, 0x00})
		New("msx: medium green", colour.SRGB8{0x3E, 0xB8, 0x49})
		New("msx: light green", colour.SRGB8{0x74, 0xd0, 0x7d})
		New("msx: dark blue", colour.SRGB8{0x59, 0x55, 0xe0})
		New("msx: light blue", colour.SRGB8{0x80, 0x76, 0xf1})
		New("msx: dark red", colour.SRGB8{0xb9, 0x5e, 0x51})
		New("msx: cyan", colour.SRGB8{0x65, 0xdb, 0xef})
		New("msx: medium red", colour.SRGB8{0xdb, 0x65, 0x59})
		New("msx: light red", colour.SRGB8{0xff, 0x89, 0x7d})
		New("msx: dark yellow", colour.SRGB8{0xcc, 0xc3, 0x5e})
		New("msx: light yellow", colour.SRGB8{0xde, 0xd0, 0x87})
		New("msx: dark green", colour.SRGB8{0x3a, 0xa2, 0x41})
		New("msx: magenta", colour.SRGB8{0xb7, 0x66, 0xb5})
		New("msx: gray", colour.SRGB8{0xcc, 0xcc, 0xcc})
		New("msx: white", colour.SRGB8{0xff, 0xff, 0xff})

	case "CPC":
		Clear()
		New("cpc: black", colour.SRGB8{0x00, 0x00, 0x00})
		New("cpc: blue", colour.SRGB8{0x00, 0x00, 0x80})
		New("cpc: bright blue", colour.SRGB8{0x00, 0x00, 0xff})
		New("cpc: red", colour.SRGB8{0x80, 0x00, 0x00})
		New("cpc: magenta", colour.SRGB8{0x80, 0x00, 0x80})
		New("cpc: mauve", colour.SRGB8{0x80, 0x00, 0xff})
		New("cpc: bright red", colour.SRGB8{0xff, 0x00, 0x00})
		New("cpc: purple", colour.SRGB8{0xff, 0x00, 0x80})
		New("cpc: bright magenta", colour.SRGB8{0xff, 0x00, 0xff})
		New("cpc: green", colour.SRGB8{0x00, 0x80, 0x00})
		New("cpc: cyan", colour.SRGB8{0x00, 0x80, 0x80})
		New("cpc: sky blue", colour.SRGB8{0x00, 0x80, 0xff})
		New("cpc: yellow", colour.SRGB8{0x80, 0x80, 0x00})
		New("cpc: white", colour.SRGB8{0x80, 0x80, 0x80})
		New("cpc: pastel blue", colour.SRGB8{0x80, 0x80, 0xff})
		New("cpc: orange", colour.SRGB8{0xff, 0x80, 0x00})
		New("cpc: pink", colour.SRGB8{0xff, 0x80, 0x80})
		New("cpc: pastel magenta", colour.SRGB8{0xff, 0x80, 0xff})
		New("cpc: bright green", colour.SRGB8{0x00, 0xff, 0x00})
		New("cpc: sea green", colour.SRGB8{0x00, 0xff, 0x80})
		New("cpc: bright cyan", colour.SRGB8{0x00, 0xff, 0xff})
		New("cpc: lime", colour.SRGB8{0x80, 0xff, 0x00})
		New("cpc: pastel green", colour.SRGB8{0x80, 0xff, 0x80})
		New("cpc: pastel cyan", colour.SRGB8{0x80, 0xff, 0xff})
		New("cpc: bright yellow", colour.SRGB8{0xff, 0xff, 0x00})
		New("cpc: pastel yellow", colour.SRGB8{0xff, 0xff, 0x80})
		New("cpc: bright white", colour.SRGB8{0xff, 0xff, 0xff})

	case "C64":
		// http://unusedino.de/ec64/technical/misc/vic656x/colors/
		//
		// http://www.pepto.de/projects/colorvic/
		Clear()
		New("c64: black", colour.SRGB8{0x00, 0x00, 0x00})
		New("c64: white", colour.SRGB8{0xff, 0xff, 0xff})
		New("c64: red", colour.SRGB8{0x68, 0x37, 0x2b})
		New("c64: cyan", colour.SRGB8{0x70, 0xa4, 0xb2})
		New("c64: violet", colour.SRGB8{0x6f, 0x3d, 0x86})
		New("c64: green", colour.SRGB8{0x58, 0x8d, 0x43})
		New("c64: blue", colour.SRGB8{0x35, 0x28, 0x79})
		New("c64: yellow", colour.SRGB8{0xb8, 0xc7, 0x6f})
		New("c64: orange", colour.SRGB8{0x6f, 0x4f, 0x25})
		New("c64: brown", colour.SRGB8{0x43, 0x39, 0x00})
		New("c64: light red", colour.SRGB8{0x9a, 0x67, 0x59})
		New("c64: dark grey", colour.SRGB8{0x44, 0x44, 0x44})
		New("c64: grey", colour.SRGB8{0x6c, 0x6c, 0x6c})
		New("c64: light green", colour.SRGB8{0x9a, 0xd2, 0x84})
		New("c64: light blue", colour.SRGB8{0x6c, 0x5e, 0xb5})
		New("c64: light grey", colour.SRGB8{0x95, 0x95, 0x95})

	default:
		return errors.New("unknown palette")
	}

	return nil
}

//------------------------------------------------------------------------------
