// Copyright (c) 2018-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package pixel

import (
	"errors"
)

//------------------------------------------------------------------------------

type Font uint16

var fontNames map[string]Font

func init() {
	fontNames = make(map[string]Font, 8)
}

var glyphMap []mapping

type font struct {
	height   int16
	baseline int16
	first    uint16 // index of the first glyph
}

var fonts []font

//------------------------------------------------------------------------------

func newFont(name string, h int16, baseline int16) Font {
	d := font{
		height:   h,
		baseline: baseline,
		first:    uint16(len(glyphMap)),
	}
	fonts = append(fonts, d)
	f := Font(len(fonts) - 1)
	fontNames[name] = f
	return f
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

func (f Font) Height() int16 {
	return fonts[f].height
}

//------------------------------------------------------------------------------

// GetFont returns the font associated with a name. If there isn't any, an
// empty font is returned, and a sticky error is set.
func GetFont(name string) Font {
	f, ok := fontNames[name]
	if !ok {
		setErr("in GetFont", errors.New("font \""+name+"\" not found"))
	}
	return f
}

//------------------------------------------------------------------------------
