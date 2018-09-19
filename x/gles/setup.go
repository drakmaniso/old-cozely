// Copyright (c) 2013-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package gles

import (
	"errors"

	"github.com/cozely/cozely/internal"
)

////////////////////////////////////////////////////////////////////////////////

/*
#cgo linux LDFLAGS: -ldl
#cgo windows LDFLAGS: -lSDL2
#cgo linux freebsd darwin pkg-config: sdl2
#include "glad.h"

#if defined(__WIN32)
#include <SDL2/SDL.h>
#else
#include <SDL2/SDL.h>
#endif

void errorCallback(
    GLenum source,
    GLenum type,
    GLuint id,
    GLenum severity,
    GLsizei length,
    const GLchar *message,
    const void *userParam);

static inline int InitOpenGL(int debug){
	if(!gladLoadGLES2Loader(SDL_GL_GetProcAddress)) {
		return -1;
	}

	if(debug) {
		//glEnable(GL_DEBUG_OUTPUT);
		// glDebugMessageCallback(errorCallback, NULL);
	}

	return 0;
}


static inline const char* GetString(GLenum name) {
	return glGetString(name);
}
*/
import "C"

////////////////////////////////////////////////////////////////////////////////

func init() {
	internal.GLSetup = setup
}

func setup() error {
	var d C.int
	if internal.Config.Debug {
		d = 1
	}
	if C.InitOpenGL(d) != 0 {
		return errors.New("gl setup: impossible to initialize OpenGL")
	}

	// clearPipelineCurrentState()

	ver := C.GoString(C.GetString(C.GL_VERSION))
	println(ver)
	return nil
}

////////////////////////////////////////////////////////////////////////////////

func init() {
	internal.GLPrerender = prerender
}

func prerender() error {
	// DefaultFramebuffer.Bind(DrawFramebuffer)
	// if !noclear {
	// 	ClearColorBuffer(struct{ R, G, B, A float32 }{0, 0, 0, 0})
	// }
	return nil
}

////////////////////////////////////////////////////////////////////////////////

func init() {
	internal.GLCleanup = cleanup
}

func cleanup() error {
	return nil
}

////////////////////////////////////////////////////////////////////////////////
