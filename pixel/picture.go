// Copyright (c) 2013-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package pixel

import (
	"errors"
)

//------------------------------------------------------------------------------

// A Picture identifies an image than can be displayed on a Canvas.
type Picture uint16

var picturePaths = []string{""}

var pictureMap = []mapping{
	mapping{},
}

type mapping struct {
	bin  int16
	x, y int16
	w, h int16
}

//------------------------------------------------------------------------------

// NewPicture reserves an ID for a picture, that will be loaded from path by
// glam.Run.
func NewPicture(path string) Picture {
	if len(pictureMap) >= 0xFFFF {
		setErr("in NewPitcture", errors.New("too many pictures"))
		return Picture(0)
	}

	picturePaths = append(picturePaths, path)
	pictureMap = append(pictureMap, mapping{})
	return Picture(len(picturePaths) - 1)
}

//------------------------------------------------------------------------------

// Size returns the width and height of the picture.
func (p Picture) Size() Coord {
	return Coord{pictureMap[p].w, pictureMap[p].h}
}

//------------------------------------------------------------------------------
