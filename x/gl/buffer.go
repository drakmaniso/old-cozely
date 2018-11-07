// Copyright (c) 2013-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package gl

import (
	"fmt"
	"reflect"
	"unsafe"

	"github.com/cozely/cozely/internal"
)

////////////////////////////////////////////////////////////////////////////////

/*
#include "glad.h"

static inline GLuint NewBuffer(GLsizeiptr size, void* data, GLbitfield flags) {
	GLuint b;
	glCreateBuffers(1, &b);
	glNamedBufferStorage(b, size, data, flags);
	return b;
}

static inline GLuint NewBufferTexture(GLuint buffer, GLenum format) {
	GLuint t;
	glCreateTextures(GL_TEXTURE_BUFFER, 1, &t);
	glTextureBuffer(t, format, buffer);
	return t;
}

void DeleteBuffer(GLuint b) {
	glDeleteBuffers(1, &b);
}

static inline void BufferSubData(GLuint buffer, GLintptr offset, GLsizei size, void *data) {
	glNamedBufferSubData(buffer, offset, size, data);
}

static inline void BindUniform(GLuint binding, GLuint buffer) {
	glBindBufferBase(GL_UNIFORM_BUFFER, binding, buffer);
}

static inline void BindStorage(GLuint binding, GLuint buffer) {
	glBindBufferBase(GL_SHADER_STORAGE_BUFFER, binding, buffer);
}

static inline void ClearBuffer(GLuint buffer, GLenum ifo, GLenum fo, GLenum t, const void *data) {
	glClearNamedBufferData(buffer, ifo, fo, t, data);
}

static inline void BindCounter(GLuint binding, GLuint buffer) {
	glBindBufferBase(GL_ATOMIC_COUNTER_BUFFER, binding, buffer);
}

static inline void BindVertex(GLuint binding, GLuint buffer, GLintptr offset, GLsizei stride) {
	glBindVertexBuffer(binding, buffer, offset, stride);
}

static inline void BindElement(GLuint buffer) {
	glBindBuffer(GL_ELEMENT_ARRAY_BUFFER, buffer);
}

static inline void BindIndirect(GLuint buffer) {
	glBindBuffer(GL_DRAW_INDIRECT_BUFFER, buffer);
}

static inline void BindTextureUnit(GLuint unit, GLuint texture) {
	glBindTextureUnit(unit, texture);
}

*/
import "C"

////////////////////////////////////////////////////////////////////////////////

// A UniformBuffer is a block of memory owned by the GPU.
type UniformBuffer struct {
	object C.GLuint
}

// NewUniformBuffer asks the GPU to allocate a new block of memory.
//
// If data is a uinptr, it is interpreted as the desired size for the buffer (in
// bytes), and the content is not initialized. Otherwise data must be a pointer
// to a struct of pure values (no nested references). In all cases the size of
// the buffer is fixed at creation.
func NewUniformBuffer(data interface{}, f BufferFlags) UniformBuffer {
	p, s, err := pointerAndSizeOf(data)
	if err != nil {
		setErr(internal.Wrap("gl uniform buffer creation", err))
		return UniformBuffer{}
	}
	var ub UniformBuffer
	ub.object = C.NewBuffer(C.GLsizeiptr(s), p, C.GLbitfield(f))
	//TODO: error handling
	return ub
}

// SubData updates the buffer with data, starting at a specified offset.
//
// It is your responsibility to ensure that the size of data plus the offset
// does not exceed the buffer size.
func (ub *UniformBuffer) SubData(data interface{}, atOffset uintptr) {
	p, s, err := pointerAndSizeOf(data)
	if err != nil {
		setErr(internal.Wrap("gl uniform buffer update", err))
		return
	}
	C.BufferSubData(ub.object, C.GLintptr(atOffset), C.GLsizei(s), p)
}

// Bind to a uniform binding index.
//
// This index should correspond to one indicated by a layout qualifier in the
// shaders.
func (ub *UniformBuffer) Bind(binding uint32) {
	C.BindUniform(C.GLuint(binding), ub.object)
}

// Delete frees the buffer.
func (ub *UniformBuffer) Delete() {
	C.DeleteBuffer(C.GLuint(ub.object))
}

////////////////////////////////////////////////////////////////////////////////

// A StorageBuffer is a block of memory owned by the GPU.
type StorageBuffer struct {
	object C.GLuint
}

// NewStorageBuffer asks the GPU to allocate a new block of memory.
//
// If data is a uinptr, it is interpreted as the desired size for the buffer (in
// bytes), and the content is not initialized. Otherwise data must be a pointer
// to a struct of pure values (no nested references). In all cases the size of
// the buffer is fixed at creation.
func NewStorageBuffer(data interface{}, f BufferFlags) StorageBuffer {
	p, s, err := pointerAndSizeOf(data)
	if err != nil {
		setErr(internal.Wrap("gl storage buffer creation", err))
		return StorageBuffer{}
	}
	var sb StorageBuffer
	sb.object = C.NewBuffer(C.GLsizeiptr(s), p, C.GLbitfield(f))
	//TODO: error handling
	return sb
}

// SubData updates the buffer with data, starting at a specified offset.
//
// It is your responsibility to ensure that the size of data plus the offset
// does not exceed the buffer size.
func (sb *StorageBuffer) SubData(data interface{}, atOffset uintptr) {
	p, s, err := pointerAndSizeOf(data)
	if err != nil {
		setErr(internal.Wrap("gl storage buffer update", err))
		return
	}
	C.BufferSubData(sb.object, C.GLintptr(atOffset), C.GLsizei(s), p)
}

func (sb *StorageBuffer) ClearUint32(data uint32) {
	C.ClearBuffer(C.GLuint(sb.object), C.GL_R32UI, C.GL_RED, C.GL_UNSIGNED_INT, unsafe.Pointer(&data))
}

// Bind to a storage binding index.
//
// This index should correspond to one indicated by a layout qualifier in the
// shaders.
func (sb *StorageBuffer) Bind(binding uint32) {
	C.BindStorage(C.GLuint(binding), sb.object)
}

// Delete frees the buffer.
func (sb *StorageBuffer) Delete() {
	if sb != nil {
		C.DeleteBuffer(C.GLuint(sb.object))
	}
}

////////////////////////////////////////////////////////////////////////////////

// A CounterBuffer is a block of memory owned by the GPU.
type CounterBuffer struct {
	object C.GLuint
}

// NewCounterBuffer asks the GPU to allocate a new block of memory.
func NewCounterBuffer(nb int, f BufferFlags) CounterBuffer {
	var sb CounterBuffer
	sb.object = C.NewBuffer(C.GLsizeiptr(4*nb), nil, C.GLbitfield(f))
	//TODO: error handling
	return sb
}

// SubData updates the buffer with data, starting at a specified offset.
//
// It is your responsibility to ensure that the size of data plus the offset
// does not exceed the buffer size.
func (sb *CounterBuffer) SubData(data []uint32, atOffset uintptr) {
	C.BufferSubData(sb.object, C.GLintptr(atOffset), C.GLsizei(4*len(data)), unsafe.Pointer(&data[0]))
}

// Bind to a storage binding index.
//
// This index should correspond to one indicated by a layout qualifier in the
// shaders.
func (sb *CounterBuffer) Bind(binding uint32) {
	C.BindCounter(C.GLuint(binding), sb.object)
}

// Delete frees the buffer.
func (sb *CounterBuffer) Delete() {
	if sb != nil {
		C.DeleteBuffer(C.GLuint(sb.object))
	}
}

////////////////////////////////////////////////////////////////////////////////

// A VertexBuffer is a block of memory owned by the GPU.
type VertexBuffer struct {
	object C.GLuint
	stride uintptr
}

// NewVertexBuffer asks the GPU to allocate a new block of memory.
//
// If data is a uinptr, it is interpreted as the desired size for the buffer (in
// bytes), and the content is not initialized. Otherwise data must be a slice of
// pure values (no nested references). In all cases the size of the buffer is
// fixed at creation.
func NewVertexBuffer(data interface{}, f BufferFlags) VertexBuffer {
	p, s, st, err := pointerSizeAndStrideOf(data)
	if err != nil {
		setErr(internal.Wrap("gl vertex buffer creation", err))
		return VertexBuffer{}
	}
	var vb VertexBuffer
	vb.object = C.NewBuffer(C.GLsizeiptr(s), p, C.GLbitfield(f))
	//TODO: error handling
	vb.stride = st
	return vb
}

// SubData updates the buffer with data, starting at a specified offset.
//
// It is your responsibility to ensure that the size of data plus the offset
// does not exceed the buffer size.
func (vb *VertexBuffer) SubData(data interface{}, atOffset uintptr) {
	p, s, st, err := pointerSizeAndStrideOf(data)
	if err != nil {
		setErr(internal.Wrap("gl vertex buffer update", err))
		return
	}
	if st != 0 {
		vb.stride = st
	}
	C.BufferSubData(vb.object, C.GLintptr(atOffset), C.GLsizei(s), p)
}

// Bind to a vertex buffer binding index.
//
// The buffer should use the same struct type than the one used in the
// corresponding call to Pipeline.VertexFormat.
func (vb *VertexBuffer) Bind(binding uint32, offset uintptr) {
	C.BindVertex(C.GLuint(binding), vb.object, C.GLintptr(offset), C.GLsizei(vb.stride))
}

func (vb *VertexBuffer) AsStorage() StorageBuffer {
	return StorageBuffer{object: vb.object}
}

// Delete frees the buffer.
func (vb *VertexBuffer) Delete() {
	C.DeleteBuffer(C.GLuint(vb.object))
}

////////////////////////////////////////////////////////////////////////////////

// An IndexBuffer is a block of memory owned by the GPU, used to store vertex
// indices.
type IndexBuffer struct {
	object C.GLuint
	gltype C.GLenum
}

var boundElement IndexBuffer

// NewIndexBuffer asks the GPU to allocate a new block of memory.
//
// If data is a uinptr, it is interpreted as the desired size for the buffer (in
// bytes), and the content is not initialized. Otherwise data must be a slice of
// uint8, uint16 or uin32. In all cases the size of the buffer is fixed at
// creation.
func NewIndexBuffer(data interface{}, f BufferFlags) IndexBuffer {
	p, s, t, err := pointerSizeAndUintTypeOf(data)
	if err != nil {
		setErr(internal.Wrap("gl index buffer creation", err))
		return IndexBuffer{}
	}
	var eb IndexBuffer
	eb.object = C.NewBuffer(C.GLsizeiptr(s), p, C.GLbitfield(f))
	//TODO: error handling
	eb.gltype = t
	return eb
}

// SubData updates the buffer with data, starting at a specified offset.
//
// It is your responsibility to ensure that the size of data plus the offset
// does not exceed the buffer size.
func (eb *IndexBuffer) SubData(data interface{}, atOffset uintptr) {
	p, s, t, err := pointerSizeAndUintTypeOf(data)
	if err != nil {
		setErr(internal.Wrap("gl index buffer update", err))
		return
	}
	if t != 0 {
		eb.gltype = t
	}
	C.BufferSubData(eb.object, C.GLintptr(atOffset), C.GLsizei(s), p)
}

// Bind the element buffer.
func (eb *IndexBuffer) Bind() {
	C.BindElement(eb.object)
	boundElement = *eb
}

// Delete frees the buffer.
func (eb *IndexBuffer) Delete() {
	C.DeleteBuffer(C.GLuint(eb.object))
}

////////////////////////////////////////////////////////////////////////////////

// An IndirectBuffer is a block of memory owned by the GPU.
type IndirectBuffer struct {
	object C.GLuint
}

// NewIndirectBuffer asks the GPU to allocate a new block of memory.
//
// If data is a
//
// In all cases the size of the buffer is
// fixed at creation.
func NewIndirectBuffer(data interface{}, f BufferFlags) IndirectBuffer {
	p, s, err := pointerAndSizeOf(data)
	if err != nil {
		setErr(internal.Wrap("gl indirect buffer creation", err))
		return IndirectBuffer{}
	}
	var ib IndirectBuffer
	ib.object = C.NewBuffer(C.GLsizeiptr(s), p, C.GLbitfield(f))
	//TODO: error handling
	return ib
}

// SubData updates the buffer with data, starting at a specified offset.
//
// It is your responsibility to ensure that the size of data plus the offset
// does not exceed the buffer size.
func (ib *IndirectBuffer) SubData(data interface{}, atOffset uintptr) {
	p, s, err := pointerAndSizeOf(data)
	if err != nil {
		setErr(internal.Wrap("gl indirect buffer update", err))
		return
	}
	C.BufferSubData(ib.object, C.GLintptr(atOffset), C.GLsizei(s), p)
}

// Bind the indirect buffer.
//
// The buffer should use the same struct type than the one used in the
// corresponding call to Pipeline.IndirectFormat.
func (ib *IndirectBuffer) Bind() {
	C.BindIndirect(ib.object)
}

// Delete frees the buffer.
func (ib *IndirectBuffer) Delete() {
	C.DeleteBuffer(C.GLuint(ib.object))
}

////////////////////////////////////////////////////////////////////////////////

// BufferTexture is a block of memory own by the GPU.
type BufferTexture struct {
	object  C.GLuint
	texture C.GLuint
}

// NewBufferTexture asks the GPU to allocate a new block of memory.
func NewBufferTexture(data interface{}, fmt TextureFormat, f BufferFlags) BufferTexture {
	p, s, err := pointerAndSizeOf(data)
	if err != nil {
		setErr(internal.Wrap("gl storage buffer creation", err))
		return BufferTexture{}
	}
	var bt BufferTexture
	bt.object = C.NewBuffer(C.GLsizeiptr(s), p, C.GLbitfield(f))
	bt.texture = C.NewBufferTexture(bt.object, C.GLenum(fmt))
	//TODO: error handling
	return bt
}

// Bind the buffer texture.
func (bt BufferTexture) Bind(index uint32) {
	C.BindTextureUnit(C.GLuint(index), bt.texture)
}

// SubData updates the buffer with data, starting at a specified offset.
//
// It is your responsibility to ensure that the size of data plus the offset
// does not exceed the buffer size.
func (bt *BufferTexture) SubData(data interface{}, atOffset uintptr) {
	p, s, err := pointerAndSizeOf(data)
	if err != nil {
		setErr(internal.Wrap("gl buffer texture update", err))
		return
	}
	C.BufferSubData(bt.object, C.GLintptr(atOffset), C.GLsizei(s), p)
}

// Delete frees the buffer.
func (bt *BufferTexture) Delete() {
	C.DeleteBuffer(C.GLuint(bt.object))
	//TODO: delete texture
}

////////////////////////////////////////////////////////////////////////////////

// BufferFlags specify which settings to use when creating a new buffer. Values
// can be ORed together.
type BufferFlags C.GLbitfield

// Used in `NewUniformBuffer` and `NewVertexBuffer`.
const (
	StaticStorage  BufferFlags = C.GL_NONE                // Content will not be updated
	MapRead        BufferFlags = C.GL_MAP_READ_BIT        // Data store will be mapped for reading
	MapWrite       BufferFlags = C.GL_MAP_WRITE_BIT       // Data store will be mapped for writing
	MapPersistent  BufferFlags = C.GL_MAP_PERSISTENT_BIT  // Data store will be accessed by both application and GPU while mapped
	MapCoherent    BufferFlags = C.GL_MAP_COHERENT_BIT    // No synchronization needed when persistently mapped
	DynamicStorage BufferFlags = C.GL_DYNAMIC_STORAGE_BIT // Content will be updated
	ClientStorage  BufferFlags = C.GL_CLIENT_STORAGE_BIT  // Prefer storage on application side
)

////////////////////////////////////////////////////////////////////////////////

func pointerSizeAndUintTypeOf(data interface{}) (ptr unsafe.Pointer, size uintptr, gltype C.GLenum, err error) {
	var p unsafe.Pointer
	var s uintptr
	var t C.GLenum
	v := reflect.ValueOf(data)
	k := v.Kind()
	switch k {
	case reflect.Uintptr:
		p = nil
		s = uintptr(v.Uint())
		t = 0
	case reflect.Slice:
		l := v.Len()
		if l == 0 {
			return nil, 0, 0, fmt.Errorf("buffer data cannot be an empty slice")
		}
		p = unsafe.Pointer(v.Pointer())
		switch v.Index(0).Kind() {
		case reflect.Uint8:
			t = C.GL_UNSIGNED_BYTE
			s = uintptr(1 * l)
		case reflect.Uint16:
			t = C.GL_UNSIGNED_SHORT
			s = uintptr(2 * l)
		case reflect.Uint32:
			t = C.GL_UNSIGNED_INT
			s = uintptr(4 * l)
		default:
			return nil, 0, 0, fmt.Errorf("buffer data must be a slice of uint8, uint16 or uint32")
		}
	default:
		return nil, 0, 0, fmt.Errorf("%s instead of slice or uintptr", reflect.TypeOf(data).Kind())
	}
	return p, s, t, nil
}

func pointerAndSizeOf(data interface{}) (ptr unsafe.Pointer, size uintptr, err error) {
	var p unsafe.Pointer
	var s uintptr
	v := reflect.ValueOf(data)
	k := v.Kind()
	switch k {
	case reflect.Uintptr:
		p = nil
		s = uintptr(v.Uint())
	case reflect.Ptr:
		p = unsafe.Pointer(v.Pointer())
		//TODO: check if pointer refer to a struct
		s = v.Elem().Type().Size()
	case reflect.Slice:
		l := v.Len()
		if l == 0 {
			return nil, 0, fmt.Errorf("buffer data cannot be an empty slice")
		}
		p = unsafe.Pointer(v.Pointer())
		st := v.Index(0).Type().Size()
		s = uintptr(l) * st
	default:
		return nil, 0, fmt.Errorf("%s instead of point-to-struct or uinptr", reflect.TypeOf(data).Kind())
	}
	return p, s, nil
}

func pointerSizeAndStrideOf(data interface{}) (ptr unsafe.Pointer, size uintptr, stride uintptr, err error) {
	var p unsafe.Pointer
	var s uintptr
	var st uintptr
	v := reflect.ValueOf(data)
	k := v.Kind()
	switch k {
	case reflect.Uintptr:
		p = nil
		s = uintptr(v.Uint())
		st = 0
	case reflect.Slice:
		l := v.Len()
		if l == 0 {
			return nil, 0, 0, fmt.Errorf("buffer data cannot be an empty slice")
		}
		p = unsafe.Pointer(v.Pointer())
		st = v.Index(0).Type().Size()
		s = uintptr(l) * st
	default:
		return nil, 0, 0, fmt.Errorf("%s instead of slice or uintptr", reflect.TypeOf(data).Kind())
	}
	return p, s, st, nil
}

////////////////////////////////////////////////////////////////////////////////
