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
*/
import "C"

//------------------------------------------------------------------------------

// A Pipeline consists of shaders and state for the GPU.
type Pipeline struct {
	object       C.GLuint
	vao          C.GLuint
	clearColor   [4]float32
	attribStride map[uint32]uintptr
}

//------------------------------------------------------------------------------

// NewPipeline returns a pipeline with created from a specific set of shaders.
func NewPipeline(s ...Shader) (Pipeline, error) {
	var p Pipeline
	p.object = C.NewPipeline() //TODO: Error Handling
	p.vao = C.CreateVAO()      //TODO: Error Handling
	p.attribStride = make(map[uint32]uintptr)
	for _, s := range s {
		if err := p.useShader(s); err != nil {
			return p, err
		}
	}
	return p, nil
}

func (p *Pipeline) useShader(s Shader) error {
	C.PipelineUseShader(p.object, s.stages, s.shader)
	if errm := C.PipelineLinkError(p.object); errm != nil {
		defer C.free(unsafe.Pointer(errm))
		return fmt.Errorf("shader link error:\n    %s", errors.New(C.GoString(errm)))
	}
	return nil
}

//------------------------------------------------------------------------------

// ClearColor sets the color used to clear the framebuffer.
func (p *Pipeline) ClearColor(color geom.Vec4) {
	p.clearColor[0] = color.X
	p.clearColor[1] = color.Y
	p.clearColor[2] = color.Z
	p.clearColor[3] = color.W
}

//------------------------------------------------------------------------------

// Bind the pipeline for use by the GPU in all following draw commands.
func (p *Pipeline) Bind() {
	C.BindPipeline(
		p.object,
		p.vao,
		(*C.GLfloat)(unsafe.Pointer(&p.clearColor[0])),
	)
}

//------------------------------------------------------------------------------

// Close the pipeline.
func (p *Pipeline) Close() {
	C.ClosePipeline(p.object, p.vao)
}

//------------------------------------------------------------------------------
