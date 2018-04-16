// Copyright (c) 2013-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package gl

import (
	"errors"

	"github.com/cozely/cozely/internal"
)

////////////////////////////////////////////////////////////////////////////////

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

	clearPipelineCurrentState()

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
