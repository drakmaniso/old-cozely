// Copyright (c) 2013-2016 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package engine

//------------------------------------------------------------------------------

// #cgo windows LDFLAGS: -lSDL2
// #cgo linux freebsd darwin pkg-config: sdl2
//
// #include "engine.h"
import "C"

import "errors"

//------------------------------------------------------------------------------

func getError() error {
	if s := C.SDL_GetError(); s != nil {
		return errors.New(C.GoString(s))
	}
	return nil
}

//------------------------------------------------------------------------------
