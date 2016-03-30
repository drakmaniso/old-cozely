// Copyright (c) 2013-2016 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package gfx

//------------------------------------------------------------------------------

import (
	"errors"
	"fmt"
	"unsafe"

	"github.com/drakmaniso/glam/geom"
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

static inline void BindPipeline(GLuint p, GLuint vao, GLfloat *c) {

    glEnable (GL_DEPTH_TEST);
    glDepthFunc (GL_LEQUAL);

    glEnable (GL_CULL_FACE);
    glCullFace (GL_BACK);

    glBlendFunc (GL_SRC_ALPHA, GL_ONE_MINUS_SRC_ALPHA);

    glEnable(GL_FRAMEBUFFER_SRGB);
	glClearBufferfv(GL_COLOR, 0, c);
	GLfloat d = 1.0;
	glClearBufferfv(GL_DEPTH, 0, &d);
	glBindProgramPipeline(p);
	glBindVertexArray(vao);
}

static inline void UniformBuffer(GLuint binding, GLuint buffer) {
	glBindBufferBase(GL_UNIFORM_BUFFER, binding, buffer);
}
*/
import "C"

//------------------------------------------------------------------------------

type Pipeline struct {
	pipeline     C.GLuint
	vao          C.GLuint
	clearColor   [4]float32
	attribStride map[uint32]uintptr
}

//------------------------------------------------------------------------------

func NewPipeline(s ...Shader) (Pipeline, error) {
	var p Pipeline
	p.pipeline = C.NewPipeline() //TODO: Error Handling
	p.vao = C.CreateVAO()        //TODO: Error Handling
	p.attribStride = make(map[uint32]uintptr)
	for _, s := range s {
		if err := p.useShader(s); err != nil {
			return p, err
		}
	}
	return p, nil
}

func (p *Pipeline) useShader(s Shader) error {
	C.PipelineUseShader(p.pipeline, s.stages, s.shader)
	if errm := C.PipelineLinkError(p.pipeline); errm != nil {
		defer C.free(unsafe.Pointer(errm))
		return fmt.Errorf("shader link error:\n    %s", errors.New(C.GoString(errm)))
	}
	return nil
}

//------------------------------------------------------------------------------

func (p *Pipeline) ClearColor(color geom.Vec4) {
	p.clearColor[0] = color.X
	p.clearColor[1] = color.Y
	p.clearColor[2] = color.Z
	p.clearColor[3] = color.W
}

//------------------------------------------------------------------------------

func (p *Pipeline) UniformBuffer(binding uint32, b Buffer) {
	C.UniformBuffer(C.GLuint(binding), b.buffer)
}

//------------------------------------------------------------------------------

func (p *Pipeline) Bind() {
	C.BindPipeline(
		p.pipeline,
		p.vao,
		(*C.GLfloat)(unsafe.Pointer(&p.clearColor[0])),
	)
}

//------------------------------------------------------------------------------

func (p *Pipeline) Close() {
	C.ClosePipeline(p.pipeline, p.vao)
}

//------------------------------------------------------------------------------
