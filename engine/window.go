// Copyright (c) 2013-2016 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package engine

import (
	"log"
	"unsafe"

	"github.com/drakmaniso/glam/internal"
)

// #include "../internal/internal.h"
import "C"

//------------------------------------------------------------------------------

// Window represents an OS window associated with an OpenGL context.
type Window struct {
	window  *C.SDL_Window
	context C.SDL_GLContext
	width   int
	height  int
}

var window Window

//------------------------------------------------------------------------------

// open creates the window and its associated OpenGL context.
func (w *Window) open(
	title string,
	resolution [2]int,
	display int,
	fullscreen bool,
	fullscreenMode string,
	vsync bool,
) (err error) {
	C.SDL_GL_SetAttribute(C.SDL_GL_CONTEXT_MAJOR_VERSION, 4)
	C.SDL_GL_SetAttribute(C.SDL_GL_CONTEXT_MINOR_VERSION, 5)
	C.SDL_GL_SetAttribute(C.SDL_GL_CONTEXT_PROFILE_MASK,
		C.SDL_GL_CONTEXT_PROFILE_CORE)
	C.SDL_GL_SetAttribute(C.SDL_GL_DOUBLEBUFFER, 1)
	C.SDL_GL_SetAttribute(C.SDL_GL_MULTISAMPLESAMPLES, 8)

	var si C.int
	if vsync {
		si = 1
	}
	C.SDL_GL_SetSwapInterval(si)

	t := C.CString(title)
	defer C.free(unsafe.Pointer(t))

	w.width, w.height = resolution[0], resolution[1]

	var fs uint32
	if fullscreen {
		if fullscreenMode == "Desktop" {
			fs = C.SDL_WINDOW_FULLSCREEN_DESKTOP
		} else {
			fs = C.SDL_WINDOW_FULLSCREEN
		}
	}
	fl := C.SDL_WINDOW_OPENGL | C.SDL_WINDOW_RESIZABLE | C.Uint32(fs)

	w.window = C.SDL_CreateWindow(
		t,
		C.int(C.SDL_WINDOWPOS_CENTERED_MASK|display),
		C.int(C.SDL_WINDOWPOS_CENTERED_MASK|display),
		C.int(w.width),
		C.int(w.height),
		fl,
	)
	if w.window == nil {
		err = internal.GetSDLError()
		log.Print(err)
		return
	}

	w.context, err = C.SDL_GL_CreateContext(w.window)
	if err != nil {
		log.Print(err)
		return
	}

	w.logOpenGLInfos()

	//TODO: Send a fake resize event (for the renderer)

	return
}

// logOpenGLInfos displays information about the OpenGL context
func (w *Window) logOpenGLInfos() {
	maj, err1 := sdlGLAttribute(C.SDL_GL_CONTEXT_MAJOR_VERSION)
	min, err2 := sdlGLAttribute(C.SDL_GL_CONTEXT_MINOR_VERSION)
	if err1 == nil && err2 == nil {
		log.Printf("OpenGL version: %d, %d\n", maj, min)
	}

	db, err1 := sdlGLAttribute(C.SDL_GL_DOUBLEBUFFER)
	if err1 == nil {
		log.Printf("OpenGL Double Buffer: %t\n", db != 0)
	}

	av, err1 := sdlGLAttribute(C.SDL_GL_ACCELERATED_VISUAL)
	if err1 == nil {
		log.Printf("OpenGL Accelerated Visual: %t\n", av != 0)
	}

	sw := C.SDL_GL_GetSwapInterval()
	if sw > 0 {
		log.Printf("OpenGL Vertical Sync: %t\n", sw != 0)
	} else {
		err1 = internal.GetSDLError()
		log.Print(err1)
	}
}

func sdlGLAttribute(attr C.SDL_GLattr) (value int, err error) {
	var v C.int
	errcode := C.SDL_GL_GetAttribute(attr, &v)
	if errcode < 0 {
		err = internal.GetSDLError()
	}
	value = int(v)
	return
}

//------------------------------------------------------------------------------

// Destroy closes the window and delete the OpenGL context
func (w *Window) destroy() {
	C.SDL_GL_DeleteContext(w.context)
	C.SDL_DestroyWindow(w.window)
}

//------------------------------------------------------------------------------
