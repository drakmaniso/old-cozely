// Copyright (c) 2013-2016 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package gpu

//------------------------------------------------------------------------------

/*
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

	"github.com/drakmaniso/carol/pixel"
)

//------------------------------------------------------------------------------

func Setup(debug bool, screenSize pixel.Coord, pixelSize int) error {

	// Initialize OpenGL

	var dbg C.int
	if debug {
		dbg = 1
	}
	if C.InitOpenGL(dbg) != 0 {
		return errors.New("impossible to initialize OpenGL")
	}

	CreateFramebuffer(screenSize, pixelSize)

	//

	SetupQuadPipeline() //TODO: elsewhere?

	return nil
}

//------------------------------------------------------------------------------
