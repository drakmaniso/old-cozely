// Copyright (c) 2013-2017 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package picture

//------------------------------------------------------------------------------

import (
	"github.com/drakmaniso/carol/internal/gfx"
)

//------------------------------------------------------------------------------

type Picture = gfx.Picture

//------------------------------------------------------------------------------

// Named returns the picture with name n and true, or an invalid Picture and false
// if no image correspond to n.
func Named(n string) (p Picture, ok bool) {
	p, ok = gfx.Pictures[n]
	return p, ok
}

//------------------------------------------------------------------------------
