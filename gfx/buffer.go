// Copyright (c) 2013-2016 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package gfx

//------------------------------------------------------------------------------

import (
	"fmt"
	"reflect"
	"unsafe"
)

/*
#include "glad.h"

GLuint NewBuffer(GLsizeiptr size, void* data, GLenum flags) {
	GLuint b;
	glCreateBuffers(1, &b);
	glNamedBufferStorage(b, size, data, flags);
	return b;
}

static inline void UpdateBuffer(GLuint buffer, GLintptr offset, GLsizei size, void *data) {
	glNamedBufferSubData(buffer, offset, size, data);
}

static inline void BindUniform(GLuint binding, GLuint buffer) {
	glBindBufferBase(GL_UNIFORM_BUFFER, binding, buffer);
}

static inline void BindVertex(GLuint binding, GLuint buffer, GLintptr offset, GLsizei stride) {
	glBindVertexBuffer(binding, buffer, offset, stride);
}
*/
import "C"

//------------------------------------------------------------------------------

// A Buffer is a block of memory owned by the GPU.
type Buffer struct {
	object C.GLuint
	stride uintptr
}

//------------------------------------------------------------------------------

// NewBuffer asks the GPU to allocate a new block of memory. If data is a
// uinptr, it is interpreted as the desired size for the buffer (in bytes), and
// the content is not initialized. Otherwise data must be a slice or a pointer
// to pure values (no nested references). In all cases the size of the buffer is
// fixed at creation.
//
// The flags must be a combination of the following:
//     DynamicStorage
//     MapRead
//     MapWrite
//     MapPersistent
//     MapCoherent
func NewBuffer(data interface{}, f bufferFlags) (Buffer, error) {
	s, p, st, err := sizeAndPointerOf(data)
	if err != nil {
		return Buffer{}, err
	}
	var b Buffer
	b.object = C.NewBuffer(C.GLsizeiptr(s), p, C.GLenum(f))
	//TODO: error handling
	b.stride = st
	return b, nil
}

// Update a buffer with data, starting at a specified offset. It is your
// responsability to ensure that the size of data plus the offset does not
// exceed the buffer size.
func (b *Buffer) Update(data interface{}, atOffset uintptr) error {
	s, p, st, err := sizeAndPointerOf(data)
	if err != nil {
		return err
	}
	C.UpdateBuffer(b.object, C.GLintptr(atOffset), C.GLsizei(s), p)
	if b.stride == 0 {
		// In case the stride was not specified at buffer creation.
		b.stride = st
	}
	return nil
}

func sizeAndPointerOf(data interface{}) (size uintptr, ptr unsafe.Pointer, stride uintptr, err error) {
	var s uintptr
	var st uintptr
	var p unsafe.Pointer
	v := reflect.ValueOf(data)
	k := v.Kind()
	switch k {
	case reflect.Uintptr:
		s = uintptr(v.Uint())
		st = 0
		p = nil
	case reflect.Slice:
		l := v.Len()
		if l == 0 {
			return 0, nil, 0, fmt.Errorf("buffer data cannot be an empty slice")
		}
		p = unsafe.Pointer(v.Pointer())
		st = v.Index(0).Type().Size()
		s = uintptr(l) * st
	case reflect.Ptr:
		p = unsafe.Pointer(v.Pointer())
		st = 0
		s = v.Elem().Type().Size()
	default:
		return 0, nil, 0, fmt.Errorf("buffer data must be a slice or a pointer, not a %s", reflect.TypeOf(data).Kind())
	}
	return s, p, st, nil
}

//------------------------------------------------------------------------------

type bufferFlags uint32

// Flags for buffer creation.
const (
	MapRead        bufferFlags = 0x0001 // Data store will be mapped for reading
	MapWrite       bufferFlags = 0x0002 // Data store will be mapped for writing
	MapPersistent  bufferFlags = 0x0040 // Data store will be accessed by both application and GPU while mapped
	MapCoherent    bufferFlags = 0x0080 // No synchronization needed when persistently mapped
	DynamicStorage bufferFlags = 0x0100 // Content will be updated
	ClientStorage  bufferFlags = 0x0200 // Prefer storage on application side
)

//------------------------------------------------------------------------------

// BindUniform binds the buffer to a uniform binding index. This index should
// correspond to one indicated by a layout qualifier in the shaders.
func (b *Buffer) BindUniform(binding uint32) {
	C.BindUniform(C.GLuint(binding), b.object)
}

//------------------------------------------------------------------------------

// BindVertex binds the buffer to a vertex buffer binding index.
//
// The buffer should use the same struct type than the one used in the
// corresponding call to Pipeline.VertexFormat.
func (b *Buffer) BindVertex(binding uint32, offset uintptr) {
	C.BindVertex(C.GLuint(binding), b.object, C.GLintptr(offset), C.GLsizei(b.stride))
}

//------------------------------------------------------------------------------
