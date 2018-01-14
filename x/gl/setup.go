// Copyright (c) 2013-2016 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package gl

//------------------------------------------------------------------------------

import (
	"errors"
	"log"
	"os"
)

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

//------------------------------------------------------------------------------

type logger interface {
	Print(v ...interface{})
	Println(v ...interface{})
	Printf(format string, v ...interface{})
}

type nolog struct{}

func (nolog) Print(v ...interface{})                 {}
func (nolog) Println(v ...interface{})               {}
func (nolog) Printf(format string, v ...interface{}) {}

var debug logger = nolog{}

//------------------------------------------------------------------------------

// Setup is called by glam.Run, and should not be called manually.
func Setup(dbg bool) error {
	var d C.int
	if dbg {
		debug = log.New(os.Stderr, "", log.Ltime|log.Lmicroseconds)
		d = 1
	}
	if C.InitOpenGL(d) != 0 {
		return errors.New("impossible to initialize OpenGL")
	}

	return nil
}

//------------------------------------------------------------------------------
