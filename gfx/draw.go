// Copyright (c) 2013-2016 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package gfx

//------------------------------------------------------------------------------

/*
#include "glad.h"

static inline void DrawArrays(GLenum m, GLuint f, GLuint c) {
	glDrawArrays(m, f, c);
}
*/
import "C"

//------------------------------------------------------------------------------

// A Primitive specifies what kind of object to draw.
type Primitive C.GLenum

// Used in 'Draw'.
const (
	Points               Primitive = C.GL_POINTS
	Lines                Primitive = C.GL_LINES
	LineLoop             Primitive = C.GL_LINE_LOOP
	LineStrip            Primitive = C.GL_LINE_STRIP
	Triangles            Primitive = C.GL_TRIANGLES
	TriangleStrip        Primitive = C.GL_TRIANGLE_STRIP
	TriangleFan          Primitive = C.GL_TRIANGLE_FAN
	LinesAdjency         Primitive = C.GL_LINES_ADJACENCY
	LineStripAdjency     Primitive = C.GL_LINE_STRIP_ADJACENCY
	TrianglesAdjency     Primitive = C.GL_TRIANGLES_ADJACENCY
	TriangleStripAdjency Primitive = C.GL_TRIANGLE_STRIP_ADJACENCY
	Patches              Primitive = C.GL_PATCHES
)

//------------------------------------------------------------------------------

// Draw count primitives, starting at first.
func Draw(mode Primitive, first int32, count int32) {
	C.DrawArrays(C.GLenum(mode), C.GLuint(first), C.GLuint(count))
}

//------------------------------------------------------------------------------
