// Copyright (c) 2013-2016 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package internal

//------------------------------------------------------------------------------

/*
#include "glad.h"

GLuint CompileShader(GLenum t, const GLchar* b);
char* CompileShaderError(GLuint s);
GLuint CreatePipeline();
GLuint CreateVAO();
void PipelineUseShader(GLuint p, GLenum stages, GLuint shader);
char* ShaderLinkError(GLuint s);
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
GLuint CreateBuffer(GLsizeiptr size, void* data, GLenum flags);

static inline void UsePipeline(GLuint p, GLuint vao, GLfloat *c) {
	glClearBufferfv(GL_COLOR, 0, c);
	GLfloat d = 1.0;
	glClearBufferfv(GL_DEPTH, 0, &d);
	glBindProgramPipeline(p);
	glBindVertexArray(vao);
};

static inline void UniformBuffer(GLuint binding, GLuint buffer) {glBindBufferBase(GL_UNIFORM_BUFFER, binding, buffer);};

static inline void DrawArrays(GLenum m, GLuint f, GLuint c) {glDrawArrays(m, f, c);};

static inline void VertexBuffer(GLuint vao, GLuint binding, GLuint buffer, GLintptr offset, GLsizei stride) {
	glVertexArrayVertexBuffer(vao, binding, buffer, offset, stride);
};

static inline void UpdateBuffer(GLuint buffer, GLintptr offset, GLsizei size, void *data) {
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
	pipeline C.GLuint
	vao      C.GLuint
}

func (p *Pipeline) Create() error {
	p.pipeline = C.CreatePipeline()
	p.vao = C.CreateVAO()
	return nil
}

func (p *Pipeline) UseShader(s *Shader) error {
	C.PipelineUseShader(C.GLuint(p.pipeline), C.GLenum(s.stages), C.GLuint(s.shader))
	if errm := C.ShaderLinkError(p.pipeline); errm != nil {
		defer C.free(unsafe.Pointer(errm))
		return fmt.Errorf("shader link error:\n    %s", errors.New(C.GoString(errm)))
	}
	return nil
}

func (p *Pipeline) Use(clearColor [4]float32) {
	C.UsePipeline(
		p.pipeline,
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
	C.ClosePipeline(p.pipeline, p.vao)
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

type Shader struct {
	shader C.GLuint
	stages C.GLenum
}

func (s *Shader) Create(t uint32, r io.Reader) error {
	b, err := ioutil.ReadAll(r)
	if err != nil {
		return fmt.Errorf("failed to read shader: %s", err)
	}
	cb := C.CString(string(b))
	defer C.free(unsafe.Pointer(cb))

	s.shader = C.CompileShader(C.GLenum(t), (*C.GLchar)(unsafe.Pointer(cb)))
	if errm := C.ShaderLinkError(s.shader); errm != nil {
		defer C.free(unsafe.Pointer(errm))
		return errors.New(C.GoString(errm))
	}

	switch t {
	case C.GL_VERTEX_SHADER:
		s.stages = C.GL_VERTEX_SHADER_BIT
	case C.GL_FRAGMENT_SHADER:
		s.stages = C.GL_FRAGMENT_SHADER_BIT
	case C.GL_GEOMETRY_SHADER:
		s.stages = C.GL_GEOMETRY_SHADER_BIT
	case C.GL_TESS_CONTROL_SHADER:
		s.stages = C.GL_TESS_CONTROL_SHADER_BIT
	case C.GL_TESS_EVALUATION_SHADER:
		s.stages = C.GL_TESS_EVALUATION_SHADER_BIT
	case C.GL_COMPUTE_SHADER:
		s.stages = C.GL_COMPUTE_SHADER_BIT
	}
	return nil
}

//------------------------------------------------------------------------------

type Buffer struct {
	buffer C.GLuint
}

func (b *Buffer) Create(size uintptr, data uintptr, flags uint32) error {
	b.buffer = C.CreateBuffer(C.GLsizeiptr(size), unsafe.Pointer(data), C.GLenum(flags))
	return nil
}

func (b *Buffer) Update(offset uintptr, size uintptr, data uintptr) {
	C.UpdateBuffer(b.buffer, C.GLintptr(offset), C.GLsizei(size), unsafe.Pointer(data))
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
