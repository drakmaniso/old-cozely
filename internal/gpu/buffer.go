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

static inline void BufferSubData(GLuint buffer, GLintptr offset, GLsizeiptr size, const void* data) {
	glNamedBufferSubData(buffer, offset, size, data);
}

*/
import "C"

//------------------------------------------------------------------------------

import (
	"unsafe"
)

//------------------------------------------------------------------------------

func createStampBuffer(size int) {
	b := C.NewBuffer(C.GLsizeiptr(size), unsafe.Pointer(nil), C.GL_DYNAMIC_STORAGE_BIT|C.GL_MAP_WRITE_BIT)
	C.BindBufferBase(C.GL_SHADER_STORAGE_BUFFER, 0, b)
	stampSSBO = b
}

var stampSSBO C.GLuint

func updateStampBuffer(data []Stamp) {
	l := len(data) * int(unsafe.Sizeof(Stamp{}))
	if l > 0 {
		C.BufferSubData(stampSSBO, 0, C.GLsizeiptr(l), unsafe.Pointer(&data[0]))
	}
}

//------------------------------------------------------------------------------

func createPaletteBuffer() {
	b := C.NewBuffer(C.GLsizeiptr(256*256*4*4), unsafe.Pointer(nil), C.GL_DYNAMIC_STORAGE_BIT|C.GL_MAP_WRITE_BIT)
	C.BindBufferBase(C.GL_SHADER_STORAGE_BUFFER, 2, b)
	paletteSSBO = b
}

var paletteSSBO C.GLuint

const sizeofColor = unsafe.Sizeof(struct{ R, G, B, A float32 }{})
const sizeofPalette = unsafe.Sizeof([256]struct{ R, G, B, A float32 }{})

func UpdatePaletteBuffer(palette uint8, data []struct{ R, G, B, A float32 }) {
	l := uintptr(len(data)) * sizeofColor
	if l > 0 {
		C.BufferSubData(paletteSSBO, C.GLintptr(uintptr(palette)*sizeofPalette), C.GLsizeiptr(l), unsafe.Pointer(&data[0]))
	}
}

//------------------------------------------------------------------------------

func CreatePictureBuffer(data []uint8) {
	b := C.NewBuffer(C.GLsizeiptr(len(data)), unsafe.Pointer(&data[0]), 0)
	C.BindBufferBase(C.GL_SHADER_STORAGE_BUFFER, 1, b)
}

//------------------------------------------------------------------------------
