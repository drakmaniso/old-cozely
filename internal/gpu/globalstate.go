// Copyright (c) 2013-2017 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package gpu

//------------------------------------------------------------------------------

/*
#include "glad.h"
*/
import "C"

//------------------------------------------------------------------------------

import (
	"github.com/drakmaniso/carol/pixel"
)

//------------------------------------------------------------------------------

// Framebuffer contains the state of the framebuffer
var Framebuffer struct {
	fbo       C.GLuint
	Size      pixel.Coord
	PixelSize int
}

//------------------------------------------------------------------------------
