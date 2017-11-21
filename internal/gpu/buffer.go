// Copyright (c) 2013-2017 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package gpu

//------------------------------------------------------------------------------

/*
#include "glad.h"

static inline GLuint NewBuffer(GLsizeiptr size, void* data, GLbitfield flags) {
	GLuint b;
	glCreateBuffers(1, &b);
	glNamedBufferStorage(b, size, data, flags);
	return b;
}

static inline void BindBufferBase(GLuint target, GLuint binding, GLuint buffer) {
	glBindBufferBase(target, binding, buffer);
}

*/
import "C"

//------------------------------------------------------------------------------

import (
	"unsafe"
)

//------------------------------------------------------------------------------

type (
	PictureBuffer C.GLuint
	StampBuffer   C.GLuint
)

//------------------------------------------------------------------------------

func CreatePictureBuffer(data []uint8) PictureBuffer {
	b := C.NewBuffer(C.GLsizeiptr(len(data)), unsafe.Pointer(&data[0]), 0)
	pictureBuffer = PictureBuffer(b)
	return PictureBuffer(b)
}

var pictureBuffer PictureBuffer

func CreateStampBuffer(size int) StampBuffer {
	b := C.NewBuffer(C.GLsizeiptr(size), unsafe.Pointer(nil), C.GL_MAP_WRITE_BIT)
	return StampBuffer(b)
}

//------------------------------------------------------------------------------
