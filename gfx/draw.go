// Copyright (c) 2013-2016 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package gfx

//------------------------------------------------------------------------------

/*
#include "glad.h"

static inline void DrawArrays(GLenum m, GLuint f, GLuint c) {
	glDrawArrays(m, f, c);
}

static inline void DrawArraysInstanced(GLenum m, GLuint f, GLuint c, GLuint ic) {
	glDrawArraysInstanced(m, f, c, ic);
}

static inline void DrawElements(GLenum m, GLsizei c, GLenum t, GLsizeiptr i) {
	glDrawElements(m, c, t, (const void *)i);
}
*/
import "C"

//------------------------------------------------------------------------------

// A Primitive specifies what kind of object to draw.
type Primitive C.GLenum

// Used in `Draw`.
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

// Draw count primitives, starting at first.
func DrawInstanced(mode Primitive, first int32, count int32, instances int32) {
	C.DrawArraysInstanced(C.GLenum(mode), C.GLuint(first), C.GLuint(count), C.GLuint(instances))
}

//------------------------------------------------------------------------------

// DrawElements draws count indexed primitives, starting at first.
func DrawElements(mode Primitive, first int32, count int32) {
	var s int32
	switch boundElement.gltype {
	case C.GL_UNSIGNED_BYTE:
		s = 1
	case C.GL_UNSIGNED_SHORT:
		s = 2
	case C.GL_UNSIGNED_INT:
		s = 4
	}
	C.DrawElements(C.GLenum(mode), C.GLsizei(count), boundElement.gltype, C.GLsizeiptr(first*s))
}

//------------------------------------------------------------------------------
