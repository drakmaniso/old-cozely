// Copyright (c) 2013-2016 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package internal

import (
	"fmt"
	"unsafe"
)

/*
#include <stdlib.h>
#include "sdl.h"

static inline void SwapWindow(SDL_Window* w) {SDL_GL_SwapWindow(w);}
*/
import "C"

//------------------------------------------------------------------------------

// Window is the game window.
var Window struct {
	window  *C.SDL_Window
	context C.SDL_GLContext
	Width   int
	Height  int
}

// Focus state
var (
	HasFocus      bool
	HasMouseFocus bool
)

//------------------------------------------------------------------------------

// OpenWindow creates the game window and its associated OpenGL context.
func OpenWindow(
	title string,
	resolution [2]int,
	display int,
	fullscreen bool,
	fullscreenMode string,
	vsync bool,
	debug bool,
) error {
	C.SDL_GL_SetAttribute(C.SDL_GL_CONTEXT_MAJOR_VERSION, 4)
	C.SDL_GL_SetAttribute(C.SDL_GL_CONTEXT_MINOR_VERSION, 5)
	C.SDL_GL_SetAttribute(C.SDL_GL_CONTEXT_PROFILE_MASK, C.SDL_GL_CONTEXT_PROFILE_CORE)
	C.SDL_GL_SetAttribute(C.SDL_GL_DOUBLEBUFFER, 1)
	C.SDL_GL_SetAttribute(C.SDL_GL_MULTISAMPLEBUFFERS, 1)
	C.SDL_GL_SetAttribute(C.SDL_GL_MULTISAMPLESAMPLES, 8)

	if debug {
		C.SDL_GL_SetAttribute(C.SDL_GL_CONTEXT_FLAGS, C.SDL_GL_CONTEXT_DEBUG_FLAG)
	}

	var si C.int
	if vsync {
		si = 1
	}
	C.SDL_GL_SetSwapInterval(si)

	t := C.CString(title)
	defer C.free(unsafe.Pointer(t))

	Window.Width, Window.Height = resolution[0], resolution[1]

	var fs uint32
	if fullscreen {
		if fullscreenMode == "Desktop" {
			fs = C.SDL_WINDOW_FULLSCREEN_DESKTOP
		} else {
			fs = C.SDL_WINDOW_FULLSCREEN
		}
	}
	fl := C.SDL_WINDOW_OPENGL | C.SDL_WINDOW_RESIZABLE | C.Uint32(fs)

	Window.window = C.SDL_CreateWindow(
		t,
		C.int(C.SDL_WINDOWPOS_CENTERED_MASK|display),
		C.int(C.SDL_WINDOWPOS_CENTERED_MASK|display),
		C.int(Window.Width),
		C.int(Window.Height),
		fl,
	)
	if Window.window == nil {
		err := GetSDLError()
		return fmt.Errorf("could not open window: %s", err)
	}

	ctx := C.SDL_GL_CreateContext(Window.window)
	if ctx == nil {
		err := GetSDLError()
		return fmt.Errorf("could not create OpenGL context: %s", err)
	}
	Window.context = ctx

	//TODO: logOpenGLInfos()

	//Send a fake resize event, with window initial size
	var e C.SDL_WindowEvent
	e._type = C.SDL_WINDOWEVENT
	e.event = C.SDL_WINDOWEVENT_RESIZED
	e.data1 = C.Sint32(Window.Width)
	e.data2 = C.Sint32(Window.Height)
	C.SDL_PushEvent((*C.SDL_Event)(unsafe.Pointer(&e)))

	return nil
}

// logOpenGLInfos displays information about the OpenGL context
func logOpenGLInfos() {
	s := "OpenGL: "
	maj, err1 := sdlGLAttribute(C.SDL_GL_CONTEXT_MAJOR_VERSION)
	min, err2 := sdlGLAttribute(C.SDL_GL_CONTEXT_MINOR_VERSION)
	if err1 == nil && err2 == nil {
		s += fmt.Sprintf("%d.%d", maj, min)
	}

	db, err1 := sdlGLAttribute(C.SDL_GL_DOUBLEBUFFER)
	if err1 == nil {
		if db != 0 {
			s += ", double buffer"
		} else {
			s += ", NO double buffer"
		}
	}

	av, err1 := sdlGLAttribute(C.SDL_GL_ACCELERATED_VISUAL)
	if err1 == nil {
		if av != 0 {
			s += ", accelerated"
		} else {
			s += ", NOT accelerated"
		}
	}

	sw := C.SDL_GL_GetSwapInterval()
	if sw > 0 {
		if sw != 0 {
			s += ", vsync"
		} else {
			s += ", NO vsync"
		}
	}
	//TODO: log.Print(s)
}

func sdlGLAttribute(attr C.SDL_GLattr) (int, error) {
	var v C.int
	errcode := C.SDL_GL_GetAttribute(attr, &v)
	if errcode < 0 {
		return 0, GetSDLError()
	}
	return int(v), nil
}

//------------------------------------------------------------------------------

// SwapWindow swaps the double-buffer.
func SwapWindow() {
	C.SwapWindow(Window.window)
}

//------------------------------------------------------------------------------

// DestroyWindow closes the game window and delete the OpenGL context
func DestroyWindow() {
	C.SDL_GL_DeleteContext(Window.context)
	C.SDL_DestroyWindow(Window.window)
}

//------------------------------------------------------------------------------
