// Copyright (c) 2013-2016 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package internal

//------------------------------------------------------------------------------

/*
#include "sdl.h"

void Render(SDL_Window* w);
*/
import "C"

//------------------------------------------------------------------------------

// Render send the draw commands to the GPU and swaps the double-buffer.
func Render() {
	C.Render(Window.window)
}

//------------------------------------------------------------------------------
