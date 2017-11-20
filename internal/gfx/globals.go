// Copyright (c) 2013-2017 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package gfx

//------------------------------------------------------------------------------

var Pictures map[string]Picture

func init() {
	Pictures = make(map[string]Picture, 128)
}

//------------------------------------------------------------------------------
