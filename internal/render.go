// Copyright (c) 2013-2016 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package internal

//------------------------------------------------------------------------------

/*
#include "sdl.h"

static inline void SwapWindow(SDL_Window* w) {SDL_GL_SwapWindow(w);}
*/
import "C"

//------------------------------------------------------------------------------

// SwapWindow swaps the double-buffer.
func SwapWindow() {
	C.SwapWindow(Window.window)
}

//------------------------------------------------------------------------------
