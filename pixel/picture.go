// Copyright (c) 2013-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package pixel

import (
	"errors"

	"github.com/cozely/cozely/coord"
	"github.com/cozely/cozely/internal"
)

////////////////////////////////////////////////////////////////////////////////

// PictureID is the ID to handle static image assets.
type PictureID uint16

const (
	maxPictureID = 0xFFFF
	noPicture = PictureID(0)
	MouseCursor = PictureID(1)
)

var picturePaths = []string{"", ""}

var pictureMap = []mapping{
	mapping{},
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
	if internal.Running {
		setErr(errors.New("pixel picture declaration: declarations must happen before starting the framework"))
		return noPicture
	}

	if len(pictureMap) >= maxPictureID {
		setErr(errors.New("pixel picture declaration: too many pictures"))
		return noPicture
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
