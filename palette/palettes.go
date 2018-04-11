// Copyright (c) 2013-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package palette

import (
	"strconv"

	"github.com/cozely/cozely/colour"
)

////////////////////////////////////////////////////////////////////////////////

var palettes = map[string][]struct {
	name   string
	colour colour.Colour
}{
	"MSX": {
		{"black", colour.SRGB8{0x00, 0x00, 0x00}},
		{"medium green", colour.SRGB8{0x3E, 0xB8, 0x49}},
		{"light green", colour.SRGB8{0x74, 0xd0, 0x7d}},
		{"dark blue", colour.SRGB8{0x59, 0x55, 0xe0}},
		{"light blue", colour.SRGB8{0x80, 0x76, 0xf1}},
		{"dark red", colour.SRGB8{0xb9, 0x5e, 0x51}},
		{"cyan", colour.SRGB8{0x65, 0xdb, 0xef}},
		{"medium red", colour.SRGB8{0xdb, 0x65, 0x59}},
		{"light red", colour.SRGB8{0xff, 0x89, 0x7d}},
		{"dark yellow", colour.SRGB8{0xcc, 0xc3, 0x5e}},
		{"light yellow", colour.SRGB8{0xde, 0xd0, 0x87}},
		{"dark green", colour.SRGB8{0x3a, 0xa2, 0x41}},
		{"magenta", colour.SRGB8{0xb7, 0x66, 0xb5}},
		{"gray", colour.SRGB8{0xcc, 0xcc, 0xcc}},
		{"white", colour.SRGB8{0xff, 0xff, 0xff}},
	},
	"CPC": {
		{"black", colour.SRGB8{0x00, 0x00, 0x00}},
		{"blue", colour.SRGB8{0x00, 0x00, 0x80}},
		{"bright blue", colour.SRGB8{0x00, 0x00, 0xff}},
		{"red", colour.SRGB8{0x80, 0x00, 0x00}},
		{"magenta", colour.SRGB8{0x80, 0x00, 0x80}},
		{"mauve", colour.SRGB8{0x80, 0x00, 0xff}},
		{"bright red", colour.SRGB8{0xff, 0x00, 0x00}},
		{"purple", colour.SRGB8{0xff, 0x00, 0x80}},
		{"bright magenta", colour.SRGB8{0xff, 0x00, 0xff}},
		{"green", colour.SRGB8{0x00, 0x80, 0x00}},
		{"cyan", colour.SRGB8{0x00, 0x80, 0x80}},
		{"sky blue", colour.SRGB8{0x00, 0x80, 0xff}},
		{"yellow", colour.SRGB8{0x80, 0x80, 0x00}},
		{"white", colour.SRGB8{0x80, 0x80, 0x80}},
		{"pastel blue", colour.SRGB8{0x80, 0x80, 0xff}},
		{"orange", colour.SRGB8{0xff, 0x80, 0x00}},
		{"pink", colour.SRGB8{0xff, 0x80, 0x80}},
		{"pastel magenta", colour.SRGB8{0xff, 0x80, 0xff}},
		{"bright green", colour.SRGB8{0x00, 0xff, 0x00}},
		{"sea green", colour.SRGB8{0x00, 0xff, 0x80}},
		{"bright cyan", colour.SRGB8{0x00, 0xff, 0xff}},
		{"lime", colour.SRGB8{0x80, 0xff, 0x00}},
		{"pastel green", colour.SRGB8{0x80, 0xff, 0x80}},
		{"pastel cyan", colour.SRGB8{0x80, 0xff, 0xff}},
		{"bright yellow", colour.SRGB8{0xff, 0xff, 0x00}},
		{"pastel yellow", colour.SRGB8{0xff, 0xff, 0x80}},
		{"bright white", colour.SRGB8{0xff, 0xff, 0xff}},
	},
	"C64": {
		{"black", colour.SRGB8{0x00, 0x00, 0x00}},
		{"white", colour.SRGB8{0xff, 0xff, 0xff}},
		{"red", colour.SRGB8{0x68, 0x37, 0x2b}},
		{"cyan", colour.SRGB8{0x70, 0xa4, 0xb2}},
		{"violet", colour.SRGB8{0x6f, 0x3d, 0x86}},
		{"green", colour.SRGB8{0x58, 0x8d, 0x43}},
		{"blue", colour.SRGB8{0x35, 0x28, 0x79}},
		{"yellow", colour.SRGB8{0xb8, 0xc7, 0x6f}},
		{"orange", colour.SRGB8{0x6f, 0x4f, 0x25}},
		{"brown", colour.SRGB8{0x43, 0x39, 0x00}},
		{"light red", colour.SRGB8{0x9a, 0x67, 0x59}},
		{"dark grey", colour.SRGB8{0x44, 0x44, 0x44}},
		{"grey", colour.SRGB8{0x6c, 0x6c, 0x6c}},
		{"light green", colour.SRGB8{0x9a, 0xd2, 0x84}},
		{"light blue", colour.SRGB8{0x6c, 0x5e, 0xb5}},
		{"light grey", colour.SRGB8{0x95, 0x95, 0x95}},
	},
}

func init() {
	pal := make([]struct {
		name   string
		colour colour.Colour
	},
		255,
		255)
	for i := 1; i < 256; i++ {
		g, r, b := i>>5, (i&0x1C)>>2, i&0x3
		pal[i-1].colour = colour.LRGBA{
			float32(r) / 7.0,
			float32(g) / 7.0,
			float32(b) / 3.0,
			1.0,
		}
		pal[i-1].name = strconv.Itoa(g) + strconv.Itoa(r) + strconv.Itoa(b)
	}

	palettes["MSX2"] = pal
}

////////////////////////////////////////////////////////////////////////////////
