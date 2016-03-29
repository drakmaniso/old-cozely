// Copyright (c) 2013-2016 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package internal

//------------------------------------------------------------------------------

/*
#include "glad.h"

GLuint CompileShader(GLenum t, const GLchar* b);
char* CompileShaderError(GLuint s);
GLuint NewPipeline();
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
GLuint NewBuffer(GLsizeiptr size, void* data, GLenum flags);

static inline void BindPipeline(GLuint p, GLuint vao, GLfloat *c) {
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

func NewPipeline() (Pipeline, error) {
	var p Pipeline
	p.pipeline = C.NewPipeline()
	p.vao = C.CreateVAO()
	return p, nil //TODO: Error Handling
}

func (p *Pipeline) UseShader(s Shader) error {
	C.PipelineUseShader(C.GLuint(p.pipeline), C.GLenum(s.stages), C.GLuint(s.shader))
	if errm := C.ShaderLinkError(p.pipeline); errm != nil {
		defer C.free(unsafe.Pointer(errm))
		return fmt.Errorf("shader link error:\n    %s", errors.New(C.GoString(errm)))
	}
	return nil
}

func (p *Pipeline) Bind(clearColor [4]float32) {
	C.BindPipeline(
		p.pipeline,
		p.vao,
		(*C.GLfloat)(unsafe.Pointer(&clearColor[0])),
	)
}

func (p *Pipeline) UniformBuffer(binding uint32, b Buffer) {
	C.UniformBuffer(C.GLuint(binding), C.GLuint(b.buffer))
}

func (p *Pipeline) VertexBuffer(binding uint32, b Buffer, offset uintptr, stride uintptr) {
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

func NewVertexShader(r io.Reader) (Shader, error) {
	var s Shader
	var err error
	s.stages = C.GL_VERTEX_SHADER_BIT
	s.shader, err = newShader(C.GL_VERTEX_SHADER, r)
	return s, err
}

func NewFragmentShader(r io.Reader) (Shader, error) {
	var s Shader
	var err error
	s.stages = C.GL_FRAGMENT_SHADER_BIT
	s.shader, err = newShader(C.GL_FRAGMENT_SHADER, r)
	return s, err
}

func NewGeometryShader(r io.Reader) (Shader, error) {
	var s Shader
	var err error
	s.stages = C.GL_GEOMETRY_SHADER_BIT
	s.shader, err = newShader(C.GL_GEOMETRY_SHADER, r)
	return s, err
}

func NewTessControlShader(r io.Reader) (Shader, error) {
	var s Shader
	var err error
	s.stages = C.GL_TESS_CONTROL_SHADER_BIT
	s.shader, err = newShader(C.GL_TESS_CONTROL_SHADER, r)
	return s, err
}

func NewTessEvaluationShader(r io.Reader) (Shader, error) {
	var s Shader
	var err error
	s.stages = C.GL_TESS_EVALUATION_SHADER_BIT
	s.shader, err = newShader(C.GL_TESS_EVALUATION_SHADER, r)
	return s, err
}

func NewComputeShader(r io.Reader) (Shader, error) {
	var s Shader
	var err error
	s.stages = C.GL_COMPUTE_SHADER_BIT
	s.shader, err = newShader(C.GL_COMPUTE_SHADER, r)
	return s, err
}

func newShader(t uint32, r io.Reader) (C.GLuint, error) {
	b, err := ioutil.ReadAll(r)
	if err != nil {
		return 0, fmt.Errorf("failed to read shader: %s", err)
	}
	cb := C.CString(string(b))
	defer C.free(unsafe.Pointer(cb))

	s := C.CompileShader(C.GLenum(t), (*C.GLchar)(unsafe.Pointer(cb)))
	if errm := C.ShaderLinkError(s); errm != nil {
		defer C.free(unsafe.Pointer(errm))
		return 0, errors.New(C.GoString(errm))
	}

	return s, nil
}

//------------------------------------------------------------------------------

type Buffer struct {
	buffer C.GLuint
}

func NewBuffer(size uintptr, data uintptr, flags uint32) (Buffer, error) {
	var b Buffer
	b.buffer = C.NewBuffer(C.GLsizeiptr(size), unsafe.Pointer(data), C.GLenum(flags))
	return b, nil
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
