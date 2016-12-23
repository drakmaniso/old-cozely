// Copyright (c) 2013-2016 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package gfx

//------------------------------------------------------------------------------

import (
	"errors"
	"fmt"
	"unsafe"

	"github.com/drakmaniso/glam/color"
)

/*
#include <stdlib.h>
#include "glad.h"

GLuint NewPipeline() {
	GLuint p;
	glCreateProgramPipelines(1, &p);
	return p;
}

GLuint CreateVAO() {
	GLuint vao;
	glCreateVertexArrays(1, &vao);
	return vao;
}

char* PipelineLinkError(GLuint pr) {
    GLint ok = GL_TRUE;
    glGetProgramiv (pr, GL_LINK_STATUS, &ok);
    if (ok != GL_TRUE)
    {
        GLint l = 0;
        glGetProgramiv (pr, GL_INFO_LOG_LENGTH, &l);
        char *m = calloc(l + 1, sizeof(char));
        glGetProgramInfoLog (pr, l, &l, m);
        return m;
    }

	return NULL;
}

void PipelineUseShader(GLuint p, GLenum stages, GLuint shader) {
	glUseProgramStages(p, stages, shader);
}

void ClosePipeline(GLuint p, GLuint vao) {
	glDeleteVertexArrays(1, &vao);
	glDeleteProgramPipelines(1, &p);
}

static inline void BindPipeline(GLuint p, GLuint vao) {

    glDepthFunc (GL_LEQUAL);

    glCullFace (GL_BACK);

    glBlendFunc (GL_SRC_ALPHA, GL_ONE_MINUS_SRC_ALPHA);

	glBindProgramPipeline(p);
	glBindVertexArray(vao);
}

static inline void ClearColorBuffer(GLfloat *c) {
	glClearBufferfv(GL_COLOR, 0, c);
}

static inline void ClearDepthBuffer(GLfloat d) {
	glClearBufferfv(GL_DEPTH, 0, &d);
}

static inline void ClearStencilBuffer(GLint m) {
	glClearBufferiv(GL_STENCIL, 0, &m);
}
*/
import "C"

//------------------------------------------------------------------------------

// A Pipeline consists of shaders and state for the GPU.
type Pipeline struct {
	object C.GLuint
	vao    C.GLuint
}

//------------------------------------------------------------------------------

// A PipelineOption represents a configuration option used when creating a new
// pipeline.
type PipelineOption func(*Pipeline)

//------------------------------------------------------------------------------

// NewPipeline returns a pipeline with created from a specific set of shaders.
func NewPipeline(o ...PipelineOption) Pipeline {
	var p Pipeline
	p.object = C.NewPipeline() //TODO: Error Handling
	p.vao = C.CreateVAO()      //TODO: Error Handling
	for _, f := range o {
		f(&p)
	}
	return p
}

func (p *Pipeline) useShader(s shader) {
	C.PipelineUseShader(p.object, s.stages, s.shader)
	if errm := C.PipelineLinkError(p.object); errm != nil {
		defer C.free(unsafe.Pointer(errm))
		setErr(
			fmt.Errorf("shader link error:\n    %s", errors.New(C.GoString(errm))),
		)
	}
	return
}

//------------------------------------------------------------------------------

// ClearColorBuffer clears the color buffer with c.
func ClearColorBuffer(c color.RGBA) {
	C.ClearColorBuffer((*C.GLfloat)(unsafe.Pointer(&c)))
}

// ClearDepthBuffer clears the depth buffer with d.
func ClearDepthBuffer(d float32) {
	C.ClearDepthBuffer((C.GLfloat)(d))
}

// ClearStencilBuffer clears the stencil buffer with m.
func ClearStencilBuffer(m int32) {
	C.ClearStencilBuffer((C.GLint)(m))
}

//------------------------------------------------------------------------------

// Bind the pipeline for use by the GPU in all following draw commands.
func (p *Pipeline) Bind() {
	C.BindPipeline(
		p.object,
		p.vao,
	)
}

//------------------------------------------------------------------------------

// Close the pipeline.
func (p *Pipeline) Close() {
	C.ClosePipeline(p.object, p.vao)
}

//------------------------------------------------------------------------------
