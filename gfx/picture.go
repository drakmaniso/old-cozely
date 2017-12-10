// Copyright (c) 2013-2017 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package gfx

//------------------------------------------------------------------------------

import (
	"errors"
)

//------------------------------------------------------------------------------

type Picture struct {
	address uint32
	width   int16
	height  int16
}

//------------------------------------------------------------------------------

var pictures map[string]Picture

func init() {
	pictures = make(map[string]Picture, 128)
}

//------------------------------------------------------------------------------

// GetPicture returns the picture associated with a name. If there isn't any, an
// empty picture is returned, and a sticky error is set.
func GetPicture(name string) Picture {
	p, ok := pictures[name]
	if !ok {
		setErr("in GetPicture", errors.New("picture \""+name+"\" not found"))
	}
	return p
}

//------------------------------------------------------------------------------

func (p Picture) Paint(x, y int16) {
	// gpu.Paint(p.address, p.width, p.height, x, y)
}

//------------------------------------------------------------------------------
