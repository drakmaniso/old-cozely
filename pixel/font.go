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

type fontMapping struct {
	h        int16
	baseline int16
	chars    [128]struct {
		bin  int16
		w    int16
		x, y int16
	}
}

var fontMap []fontMapping

//------------------------------------------------------------------------------

func newFont(name string, h int16, baseline int16) Font {
	m := fontMapping{
		h:        h,
		baseline: baseline,
	}
	fontMap = append(fontMap, m)
	f := Font(len(fontMap) - 1)
	fonts[name] = f
	return f
}

//------------------------------------------------------------------------------

func (f Font) getMap(c rune) (bin int16, x, y, w, h int16) {
	return fontMap[f].chars[c].bin, fontMap[f].chars[c].x, fontMap[f].chars[c].y,
		fontMap[f].chars[c].w, fontMap[f].h
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
