// Copyright (c) 2013-2016 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package gfx

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
*/
import "C"

//------------------------------------------------------------------------------

// A Buffer is a block of memory owned by the GPU.
type Buffer struct {
	buffer C.GLuint
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
	s, p, err := sizeAndPointerOf(data)
	if err != nil {
		return Buffer{}, err
	}
	var b Buffer
	b.buffer = C.NewBuffer(C.GLsizeiptr(s), unsafe.Pointer(p), C.GLenum(f))
	//TODO: error handling
	return b, nil
}

// Update a buffer with data, starting at a specified offset. It is your
// responsability to ensure that the size of data plus the offset does not
// exceed the buffer size.
func (b *Buffer) Update(data interface{}, atOffset uintptr) error {
	s, p, err := sizeAndPointerOf(data)
	if err != nil {
		return err
	}
	C.UpdateBuffer(b.buffer, C.GLintptr(atOffset), C.GLsizei(s), unsafe.Pointer(p))
	return nil
}

func sizeAndPointerOf(data interface{}) (size uintptr, ptr uintptr, err error) {
	var s uintptr
	var p uintptr
	v := reflect.ValueOf(data)
	k := v.Kind()
	switch k {
	case reflect.Uintptr:
		s = uintptr(v.Uint())
		p = 0
	case reflect.Slice:
		l := v.Len()
		if l == 0 {
			return 0, 0, fmt.Errorf("buffer data cannot be an empty slice")
		}
		p = v.Pointer()
		s = uintptr(l) * v.Index(0).Type().Size()
	case reflect.Ptr:
		p = v.Pointer()
		s = v.Elem().Type().Size()
	default:
		return 0, 0, fmt.Errorf("buffer data must be a slice or a pointer, not a %s", reflect.TypeOf(data).Kind())
	}
	return s, p, nil
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
