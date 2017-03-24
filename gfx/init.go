// Copyright (c) 2013-2016 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package gfx

//------------------------------------------------------------------------------

/*
#cgo linux LDFLAGS: -ldl
#include "glad.h"

int InitOpenGL(int debug);
*/
import "C"

import (
	"fmt"

	"github.com/drakmaniso/glam/internal"
)

//------------------------------------------------------------------------------

func Setup() {
	var d C.int
	if internal.Debug {
		d = 1
	}
	if C.InitOpenGL(d) != 0 {
		internal.InitError = fmt.Errorf("impossible to initialize OpenGL")
	}
}

//------------------------------------------------------------------------------
