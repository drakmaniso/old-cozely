// Copyright (c) 2013-2017 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package picture

import (
	"github.com/drakmaniso/carol/internal/gpu"
)

//------------------------------------------------------------------------------

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

// Get returns the picture associated with a name. If there isn't any, an empty
// picture is returned and ok is set to false.
func Get(name string) (p Picture, ok bool) {
	p, ok = pictures[name]
	return p, ok
}

//------------------------------------------------------------------------------

func (p Picture) Paint(x, y int16) {
	gpu.Paint(p.address, p.width, p.height, x, y)
}

//------------------------------------------------------------------------------
