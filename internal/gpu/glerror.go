// Copyright (c) 2013-2016 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package gpu

//------------------------------------------------------------------------------

/*
#include "glad.h"
*/
import "C"
import (
	"log"
	"os"
)

//------------------------------------------------------------------------------

//export logGLError
func logGLError(
	source C.GLenum,
	typ C.GLenum,
	id C.GLuint,
	severity C.GLenum,
	length C.GLsizei,
	m *C.char,
) {
	var sev string
	switch severity {
	case C.GL_DEBUG_SEVERITY_HIGH:
		sev = "ERROR"
	case C.GL_DEBUG_SEVERITY_MEDIUM:
		sev = "WARNING"
	case C.GL_DEBUG_SEVERITY_LOW:
		sev = "warning"
	case C.GL_DEBUG_SEVERITY_NOTIFICATION:
		sev = "info"
		// return //TODO
	}
	var sou string
	switch source {
	case C.GL_DEBUG_SOURCE_API:
		sou = "OpenGL"
	case C.GL_DEBUG_SOURCE_WINDOW_SYSTEM:
		sou = "Window-system API"
	case C.GL_DEBUG_SOURCE_SHADER_COMPILER:
		sou = "Shader compiler"
	case C.GL_DEBUG_SOURCE_THIRD_PARTY:
		sou = "Third party"
	case C.GL_DEBUG_SOURCE_APPLICATION:
		sou = "Application"
	case C.GL_DEBUG_SOURCE_OTHER:
		sou = "Other source"
	}
	var ty string
	switch typ {
	case C.GL_DEBUG_TYPE_ERROR:
		ty = " (error)"
	case C.GL_DEBUG_TYPE_DEPRECATED_BEHAVIOR:
		ty = " (deprecated behavior)"
	case C.GL_DEBUG_TYPE_UNDEFINED_BEHAVIOR:
		ty = " (undefined behavior)"
	case C.GL_DEBUG_TYPE_PORTABILITY:
		ty = " (portability)"
	case C.GL_DEBUG_TYPE_PERFORMANCE:
		ty = " (performance)"
	case C.GL_DEBUG_TYPE_MARKER:
		ty = " (marker)"
	case C.GL_DEBUG_TYPE_PUSH_GROUP:
		ty = " (push group)"
	case C.GL_DEBUG_TYPE_POP_GROUP:
		ty = " (pop group)"
	case C.GL_DEBUG_TYPE_OTHER:
		ty = ""
	}

	logger.Printf("%s %s%s: %s\n", sou, sev, ty, C.GoString(m))
}

var logger = log.New(os.Stderr, "*** ", log.Ltime)

//------------------------------------------------------------------------------
