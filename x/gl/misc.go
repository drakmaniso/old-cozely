// Copyright (c) 2013-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package gl

//------------------------------------------------------------------------------

/*
#include "glad.h"

static void Viewport(GLint x,  GLint y,  GLsizei width,  GLsizei height) {
	glViewport(x, y, width, height);
}

static void DepthRange(GLdouble n, GLdouble f) {
	glDepthRange(n, f);
}
*/
import "C"

//------------------------------------------------------------------------------

// Viewport sets the size in pixels of the GL viewport.
//
// Note that this function is automatically called each time the window is
// resized.
func Viewport(ox, oy, width, height int32) {
	C.Viewport(C.GLint(ox), C.GLint(oy), C.GLsizei(width), C.GLsizei(height))
}

//------------------------------------------------------------------------------

// DepthRange specifies the mapping of depth values from normalized device
// coordinates to window coordinates.
func DepthRange(near, far float64) {
	C.DepthRange(C.GLdouble(near), C.GLdouble(far))
}

//------------------------------------------------------------------------------
