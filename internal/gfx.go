// Copyright (c) 2013-2016 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package internal

//------------------------------------------------------------------------------

/*
#include "glad.h"

GLuint CompileShader(const GLchar* b, GLenum t);
char* CompileShaderError(GLuint s);
GLuint LinkProgram(GLuint vs, GLuint fs);
char* LinkProgramError(GLuint s);
GLuint SetupVAO();
void ClosePipeline(GLuint p, GLuint vao);
void VertexAttribute(
	GLuint vao,
	GLuint index,
	GLuint binding,
	GLint size,
	GLenum type,
	GLboolean normalized,
	GLuint relativeOffset
);
GLuint CreateBufferFrom(GLsizeiptr size, void* data, GLenum flags);

static inline void UsePipeline(GLuint p, GLuint vao, GLfloat *c) {
	glClearBufferfv(GL_COLOR, 0, c);
	GLfloat d = 1.0;
	glClearBufferfv(GL_DEPTH, 0, &d);
	glUseProgram(p);
	glBindVertexArray(vao);
};

static inline void UniformBuffer(GLuint binding, GLuint buffer) {glBindBufferBase(GL_UNIFORM_BUFFER, binding, buffer);};

static inline void DrawArrays(GLenum m, GLuint f, GLuint c) {glDrawArrays(m, f, c);};

static inline void VertexBuffer(GLuint vao, GLuint binding, GLuint buffer, GLintptr offset, GLsizei stride) {
	glVertexArrayVertexBuffer(vao, binding, buffer, offset, stride);
};

static inline void UpdateBufferWith(GLuint buffer, GLintptr offset, GLsizei size, void *data) {
	glNamedBufferSubData(buffer, offset, size, data);
};
*/
import "C"

import (
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"unsafe"
)

//------------------------------------------------------------------------------

type Pipeline struct {
	program C.GLuint
	vao     C.GLuint
}

func (p *Pipeline) Shaders(
	vertexShader io.Reader,
	fragmentShader io.Reader,
) error {
	vs, err := compileShader(vertexShader, C.GL_VERTEX_SHADER)
	if err != nil {
		return fmt.Errorf("vertex shader compile error:\n    %s", err)
	}
	fs, err := compileShader(fragmentShader, C.GL_FRAGMENT_SHADER)
	if err != nil {
		return fmt.Errorf("fragment shader compile error:\n    %s", err)
	}

	p.program = C.LinkProgram(vs, fs)
	if errm := C.LinkProgramError(p.program); errm != nil {
		defer C.free(unsafe.Pointer(errm))
		return fmt.Errorf("shader link error:\n    %s", errors.New(C.GoString(errm)))
	}
	return nil
}

func compileShader(r io.Reader, t C.GLenum) (C.GLuint, error) {
	b, err := ioutil.ReadAll(r)
	if err != nil {
		return 0, fmt.Errorf("failed to read shader: %s", err)
	}
	cb := C.CString(string(b))
	defer C.free(unsafe.Pointer(cb))

	s := C.CompileShader((*C.GLchar)(unsafe.Pointer(cb)), t)
	if errm := C.CompileShaderError(s); errm != nil {
		defer C.free(unsafe.Pointer(errm))
		return s, errors.New(C.GoString(errm))
	}

	return s, nil
}

func (p *Pipeline) SetupVAO() error {
	p.vao = C.SetupVAO()
	return nil
}

func (p *Pipeline) Use(clearColor [4]float32) {
	C.UsePipeline(
		p.program,
		p.vao,
		(*C.GLfloat)(unsafe.Pointer(&clearColor[0])),
	)
}

func (p *Pipeline) UniformBuffer(binding uint32, b *Buffer) {
	C.UniformBuffer(C.GLuint(binding), C.GLuint(b.buffer))
}

func (p *Pipeline) VertexBuffer(binding uint32, b *Buffer, offset uintptr, stride uintptr) {
	C.VertexBuffer(C.GLuint(p.vao), C.GLuint(binding), C.GLuint(b.buffer), C.GLintptr(offset), C.GLsizei(stride))
}

func (p *Pipeline) Close() {
	C.ClosePipeline(p.program, p.vao)
}

//------------------------------------------------------------------------------

func (p *Pipeline) VertexAttribute(
	index uint32,
	binding uint32,
	size int32,
	typ uint32,
	normalized byte,
	relativeOffset uint32,
) {
	C.VertexAttribute(
		C.GLuint(p.vao),
		C.GLuint(index),
		C.GLuint(binding),
		C.GLint(size),
		C.GLenum(typ),
		C.GLboolean(normalized),
		C.GLuint(relativeOffset),
	)
}

//------------------------------------------------------------------------------

type Buffer struct {
	buffer C.GLuint
}

func (b *Buffer) CreateFrom(size uintptr, data uintptr, flags uint32) error {
	b.buffer = C.CreateBufferFrom(C.GLsizeiptr(size), unsafe.Pointer(data), C.GLenum(flags))
	return nil
}

func (b *Buffer) UpdateWith(offset uintptr, size uintptr, data uintptr) {
	C.UpdateBufferWith(b.buffer, C.GLintptr(offset), C.GLsizei(size), unsafe.Pointer(data))
}

//------------------------------------------------------------------------------

func Draw(mode uint32, first, count int32) {
	C.DrawArrays(C.GLenum(mode), C.GLuint(first), C.GLuint(count))
}

//------------------------------------------------------------------------------

const (
	GlByteEnum          C.GLenum = C.GL_BYTE
	GlUnsignedByteEnum  C.GLenum = C.GL_UNSIGNED_BYTE
	GlShortEnum         C.GLenum = C.GL_SHORT
	GlUnsignedShortEnum C.GLenum = C.GL_UNSIGNED_SHORT
	GlIntEnum           C.GLenum = C.GL_INT
	GlUnsignedIntEnum   C.GLenum = C.GL_UNSIGNED_INT
	GlFloatEnum         C.GLenum = C.GL_FLOAT
	GlDoubleEnum        C.GLenum = C.GL_DOUBLE
)

//------------------------------------------------------------------------------
