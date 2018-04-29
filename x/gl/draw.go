// Copyright (c) 2013-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package gl

////////////////////////////////////////////////////////////////////////////////

/*
#include "glad.h"

static inline void DrawArrays(GLenum m, GLuint f, GLuint c) {
	glDrawArrays(m, f, c);
}

static inline void DrawArraysInstanced(GLenum m, GLuint f, GLuint c, GLuint ic) {
	glDrawArraysInstanced(m, f, c, ic);
}

static inline void DrawIndexed(GLenum m, GLsizei c, GLenum t, GLsizeiptr i) {
	glDrawElements(m, c, t, (const void *)i);
}


static inline void DrawIndirect(GLenum m, GLsizeiptr f, GLsizei dc, GLsizei s) {
	glMultiDrawArraysIndirect(m, (void*)(f*16), dc, s);
}

static inline void MemoryBarrier() {
	glMemoryBarrier( GL_ALL_BARRIER_BITS);
}

*/
import "C"

////////////////////////////////////////////////////////////////////////////////

// Draw asks the GPU to draw a sequence of primitives.
func Draw(first int32, count int32) {
	C.DrawArrays(currentPipeline.state.topology, C.GLuint(first), C.GLuint(count))
}

////////////////////////////////////////////////////////////////////////////////

// DrawInstanced asks the GPU to draw several instances of a sequence of
// primitives.
func DrawInstanced(first int32, count int32, instances int32) {
	C.DrawArraysInstanced(currentPipeline.state.topology, C.GLuint(first), C.GLuint(count), C.GLuint(instances))
}

////////////////////////////////////////////////////////////////////////////////

// DrawIndexed asks the GPU to draw a sequence of primitives with indexed
// vertices.
func DrawIndexed(first int32, count int32) {
	var s int32
	switch boundElement.gltype {
	case C.GL_UNSIGNED_BYTE:
		s = 1
	case C.GL_UNSIGNED_SHORT:
		s = 2
	case C.GL_UNSIGNED_INT:
		s = 4
	}
	C.DrawIndexed(currentPipeline.state.topology, C.GLsizei(count), boundElement.gltype, C.GLsizeiptr(first*s))
}

////////////////////////////////////////////////////////////////////////////////

// A DrawIndirectCommand describes a single draw call. A slice of these is used
// to fill indirect buffers.
type DrawIndirectCommand struct {
	VertexCount   uint32
	InstanceCount uint32
	FirstVertex   uint32
	BaseInstance  uint32
}

// DrawIndirect asks the GPU to read the Indirect Buffer starting at firstdraw,
// and make drawcount draw calls.
func DrawIndirect(firstdraw uintptr, drawcount int32) {
	for i := int32(0); i < drawcount; i++ {
		C.MemoryBarrier()
		C.DrawIndirect(currentPipeline.state.topology, C.GLsizeiptr(firstdraw+uintptr(i)), C.GLsizei(1), C.GLsizei(0))
	}
	// C.DrawIndirect(currentPipeline.state.topology, C.GLsizeiptr(firstdraw), C.GLsizei(drawcount), C.GLsizei(0))
}

////////////////////////////////////////////////////////////////////////////////
