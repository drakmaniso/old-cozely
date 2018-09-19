// Copyright (c) 2013-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package internal

import (
	"unsafe"
)

////////////////////////////////////////////////////////////////////////////////

/*
#include <stdlib.h>
#include "sdl.h"

static inline void SwapWindow(SDL_Window* w) {
	SDL_GL_SwapWindow(w);
}
*/
import "C"

////////////////////////////////////////////////////////////////////////////////

// SwapWindow swaps the double-buffer.
func SwapWindow() {
	C.SwapWindow(Window.window)
}

////////////////////////////////////////////////////////////////////////////////

func SetFullscreen(f bool) {
	var fs C.Uint32
	if f {
		if Config.FullscreenMode == "Desktop" {
			fs = C.SDL_WINDOW_FULLSCREEN_DESKTOP
		} else {
			fs = C.SDL_WINDOW_FULLSCREEN
		}
	}
	C.SDL_SetWindowFullscreen(Window.window, fs)
}

func GetFullscreen() bool {
	fs := C.SDL_GetWindowFlags(Window.window)
	fs &= (C.SDL_WINDOW_FULLSCREEN_DESKTOP | C.SDL_WINDOW_FULLSCREEN)
	return fs != 0
}

func ToggleFullscreen() {
	fs := !GetFullscreen()
	SetFullscreen(fs)
}

////////////////////////////////////////////////////////////////////////////////

func SetWindowTitle(title string) {
	t := C.CString(title)
	defer C.free(unsafe.Pointer(t))
	C.SDL_SetWindowTitle(Window.window, t)
}

////////////////////////////////////////////////////////////////////////////////
