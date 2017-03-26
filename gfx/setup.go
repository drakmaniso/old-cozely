// Copyright (c) 2013-2016 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package gfx

//------------------------------------------------------------------------------

/*
#cgo linux LDFLAGS: -ldl
#include "glad.h"

void errorCallback(
    GLenum source,
    GLenum type,
    GLuint id,
    GLenum severity,
    GLsizei length,
    const GLchar *message,
    const void *userParam);

static inline int InitOpenGL(int debug){
	if(!gladLoadGL()) {
		return -1;
	}

	if(debug) {
		glEnable(GL_DEBUG_OUTPUT);
		glDebugMessageCallback(errorCallback, NULL);
	}

	return 0;
}
*/
import "C"

import (
	"errors"

	"github.com/drakmaniso/glam/internal"
	"github.com/drakmaniso/glam/pixel"
)

//------------------------------------------------------------------------------

// Setup is called by glam.Setup, and should not be called manually.
func Setup() error {
	var d C.int
	if internal.Config.Debug {
		d = 1
	}
	if C.InitOpenGL(d) != 0 {
		return errors.New("impossible to initialize OpenGL")
	}

	s := pixel.Coord{internal.Window.Width, internal.Window.Height}
	Viewport(pixel.Coord{X: 0, Y: 0}, s)

	return nil
}

//------------------------------------------------------------------------------
