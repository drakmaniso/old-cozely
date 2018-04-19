// Copyright (c) 2018-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package pixel

import (
	"errors"

	"github.com/cozely/cozely/color"
	"github.com/cozely/cozely/internal"
)

////////////////////////////////////////////////////////////////////////////////

// FontID is the ID to handle font assets.
type FontID uint8

const (
	maxFontID = 0xFF
	noFont    = FontID(maxFontID)
)

// Monozela10 is the ID of the default font (a 10 pixel high monospace). This is
// the only font that is always loaded and doesn't need declaration.
const Monozela10 = FontID(0)

var fontPaths = []string{"builtin monozela 10"}

var fonts = []font{{}}

type font struct {
	height    int16
	baseline  int16
	basecolor color.Index
	first     uint16 // index of the first glyph
}

var glyphMap []mapping

////////////////////////////////////////////////////////////////////////////////

// Font declares a new font and returns its ID.
func Font(path string) FontID {
	if internal.Running {
		setErr(errors.New("pixel font declaration: declarations must happen before starting the framework"))
		return noFont
	}

	if len(fonts) >= maxFontID {
		setErr(errors.New("pixel font declaration: too many fonts"))
		return noFont
	}

	fonts = append(fonts, font{})
	fontPaths = append(fontPaths, path)
	return FontID(len(fonts) - 1)
}

////////////////////////////////////////////////////////////////////////////////

func (f FontID) glyph(r rune) uint16 {
	//TODO: add support for non-ascii runes
	switch {
	case r < ' ':
		r = 0x7F - ' '
	case r <= 0x7F:
		r = r - ' '
	default:
		r = 0x7F - ' '
	}
	return fonts[f].first + uint16(r)
}

////////////////////////////////////////////////////////////////////////////////

// Height returns the height of the font, i.e. the height of the images used to
// store the glyphs.
func (f FontID) Height() int16 {
	return fonts[f].height
}

////////////////////////////////////////////////////////////////////////////////
