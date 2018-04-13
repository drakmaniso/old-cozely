// Copyright (c) 2013-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package pixel

import (
	"errors"

	"github.com/cozely/cozely/coord"
)

////////////////////////////////////////////////////////////////////////////////

// PictureID is the ID to handle static image assets.
type PictureID uint16

var picturePaths = []string{""}

var pictureMap = []mapping{
	mapping{},
}

type mapping struct {
	bin  int16
	x, y int16
	w, h int16
}

////////////////////////////////////////////////////////////////////////////////

// Picture declares a new picture and returns its ID.
func Picture(path string) PictureID {
	if len(pictureMap) >= 0xFFFF {
		setErr("in NewPitcture", errors.New("too many pictures"))
		return PictureID(0)
	}

	picturePaths = append(picturePaths, path)
	pictureMap = append(pictureMap, mapping{})
	return PictureID(len(picturePaths) - 1)
}

////////////////////////////////////////////////////////////////////////////////

// Size returns the width and height of the picture.
func (p PictureID) Size() coord.CR {
	return coord.CR{pictureMap[p].w, pictureMap[p].h}
}

////////////////////////////////////////////////////////////////////////////////
