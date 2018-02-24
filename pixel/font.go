// Copyright (c) 2018-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package pixel

import (
	"errors"
)

//------------------------------------------------------------------------------

// A Font identifies  apixel font that can be used by Cursor to display text.
type Font uint8

var fontPaths []string

type font struct {
	height   int16
	baseline int16
	first    uint16 // index of the first glyph
}

var fonts []font

var glyphMap []mapping

//------------------------------------------------------------------------------

// NewFont reserves an ID for a new font, that will be loaded from path by
// glam.Run.
func NewFont(path string) Font {
	if len(fonts) >= 0xFF {
		setErr("in NewFont", errors.New("too many fonts"))
		return Font(0)
	}

	fonts = append(fonts, font{})
	fontPaths = append(fontPaths, path)
	return Font(len(fonts) - 1)
}

//------------------------------------------------------------------------------

func (f Font) glyph(r rune) uint16 {
	//TODO: add support for non-ascii runes
	if r < ' ' || r > 0x7F {
		r = 0x7F
	}
	return fonts[f].first + uint16(r)
}

//------------------------------------------------------------------------------

// Height returns the height of the font, i.e. the height of the images used to
// store the glyphs.
func (f Font) Height() int16 {
	return fonts[f].height
}

//------------------------------------------------------------------------------
