// Copyright (c) 2013-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package pixel

import (
	"errors"

	"github.com/drakmaniso/cozely/plane"
)

//------------------------------------------------------------------------------

// A PictureID identifies an image than can be displayed on a Canvas.
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

//------------------------------------------------------------------------------

// Picture reserves an ID for a picture, that will be loaded from path by
// cozely.Run.
func Picture(path string) PictureID {
	if len(pictureMap) >= 0xFFFF {
		setErr("in NewPitcture", errors.New("too many pictures"))
		return PictureID(0)
	}

	picturePaths = append(picturePaths, path)
	pictureMap = append(pictureMap, mapping{})
	return PictureID(len(picturePaths) - 1)
}

//------------------------------------------------------------------------------

// Size returns the width and height of the picture.
func (p PictureID) Size() plane.Pixel {
	return plane.Pixel{pictureMap[p].w, pictureMap[p].h}
}

//------------------------------------------------------------------------------
