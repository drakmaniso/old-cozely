// Copyright (c) 2013-2017 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package gfx

//------------------------------------------------------------------------------

var PictureFormats []PictureFormat

var Pictures map[string]Picture

func init() {
	PictureFormats = make([]PictureFormat, 0, 8)
	Pictures = make(map[string]Picture, 128)
}

//------------------------------------------------------------------------------
