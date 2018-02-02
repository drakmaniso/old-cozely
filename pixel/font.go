// Copyright (c) 2018-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package pixel

import (
	"errors"
)

//------------------------------------------------------------------------------

type Font uint16

var fonts map[string]Font

func init() {
	fonts = make(map[string]Font, 8)
}

//------------------------------------------------------------------------------

var glyphsMap []mapping

type fontDesc struct {
	height   int16
	baseline int16
	ascii    int16 // Glyph index of ASCII 0x00
}

var fontsDesc []fontDesc

//------------------------------------------------------------------------------

func newFont(name string, h int16, baseline int16) Font {
	d := fontDesc{
		height:   h,
		baseline: baseline,
		ascii:    int16(len(glyphsMap)),
	}
	fontsDesc = append(fontsDesc, d)
	f := Font(len(fontsDesc) - 1)
	fonts[name] = f
	return f
}

//------------------------------------------------------------------------------

func (f Font) getGlyph(r rune) int16 {
	if r < ' ' || r > 0x7F {
		r = 0x7F
	}
	return fontsDesc[f].ascii + int16(r)
}

func (f Font) getMap(c rune) (bin int16, x, y, w, h int16) {
	g := fontsDesc[f].ascii + int16(c) //TODO:
	return glyphsMap[g].binFlip, glyphsMap[g].x, glyphsMap[g].y,
		glyphsMap[g].w, glyphsMap[g].h
}

//------------------------------------------------------------------------------

func (f Font) Height() int16 {
	return fontsDesc[f].height
}

//------------------------------------------------------------------------------

// GetFont returns the font associated with a name. If there isn't any, an
// empty font is returned, and a sticky error is set.
func GetFont(name string) Font {
	f, ok := fonts[name]
	if !ok {
		setErr("in GetFont", errors.New("font \""+name+"\" not found"))
	}
	return f
}

//------------------------------------------------------------------------------
