// Copyright (c) 2013-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package pixel

import (
	"errors"
)

//------------------------------------------------------------------------------

type Picture uint16

var pictureNames map[string]Picture

func init() {
	pictureNames = make(map[string]Picture, 128)
}

//------------------------------------------------------------------------------

type mapping struct {
	bin  int16
	x, y int16
	w, h int16
}

var pictureMap = []mapping{
	mapping{},
}

//------------------------------------------------------------------------------

func newPicture(name string, w, h int16) Picture {
	pictureMap = append(pictureMap, mapping{w: w, h: h})
	p := Picture(len(pictureMap) - 1)
	pictureNames[name] = p
	return p
}

//------------------------------------------------------------------------------

// Size returns the width and height of the picture.
func (p Picture) Size() Coord {
	return Coord{pictureMap[p].w, pictureMap[p].h}
}

//------------------------------------------------------------------------------

// GetPicture returns the picture associated with a name. If there isn't any, an
// empty picture is returned, and a sticky error is set.
func GetPicture(name string) Picture {
	p, ok := pictureNames[name]
	if !ok {
		setErr("in GetPicture", errors.New("picture \""+name+"\" not found"))
	}
	return p
}

//------------------------------------------------------------------------------
