// Copyright (c) 2018-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package pixel

import (
	"errors"
)

////////////////////////////////////////////////////////////////////////////////

// A FontID identifies a pixel font that can be used to display text on the
// canvas.
type FontID uint8

// Monozela10 is the default font. It's a 10 pixel high monospace font that is
// automatically loaded when the framework starts.
const Monozela10 = FontID(0)

var fontPaths = []string{"builtin monozela 10"}

var fonts = []font{{}}

type font struct {
	height   int16
	baseline int16
	first    uint16 // index of the first glyph
}

var glyphMap []mapping

////////////////////////////////////////////////////////////////////////////////

// Font reserves an ID for a new font, that will be loaded from path when the
// framework starts.
func Font(path string) FontID {
	if len(fonts) >= 0xFF {
		setErr("in NewFont", errors.New("too many fonts"))
		return FontID(0)
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
