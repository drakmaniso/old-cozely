// Copyright (c) 2013-2016 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package gfx

//------------------------------------------------------------------------------

import (
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"unsafe"
)

/*
#include <stdlib.h>
#include "glad.h"

GLuint CompileShader(GLenum t, const GLchar* b) {
	GLuint s = glCreateShader(t);
	const GLchar*bb[] = {b};
	glShaderSource(s, 1, bb, NULL);
	glCompileShader(s);

	return s;
}

char* ShaderCompileError(GLuint p) {
    GLint ok = GL_TRUE;
    glGetShaderiv (p, GL_COMPILE_STATUS, &ok);
    if (ok != GL_TRUE)
    {
        GLint l = 0;
        glGetShaderiv (p, GL_INFO_LOG_LENGTH, &l);
        char *m = calloc(l + 1, sizeof(char));
        glGetShaderInfoLog (p, l, &l, m);
        return m;
    }

	return NULL;
}
*/
import "C"

//------------------------------------------------------------------------------

// A shader is a compiled program run by the GPU.
type shader struct {
	shader C.GLuint
	stages C.GLenum
}

//------------------------------------------------------------------------------

// VertexShader compiles a vertex shader.
func VertexShader(r io.Reader) PipelineOption {
	var s shader
	var err error
	s.stages = C.GL_VERTEX_SHADER_BIT
	s.shader, err = newShader(C.GL_VERTEX_SHADER, r)
	if err != nil {
		setErr(fmt.Errorf("unable to compile vertex shader: %s", err))
	}
	return func(p *Pipeline) {
		p.attachShader(s)
	}
}

// FragmentShader compiles a fragment shader.
func FragmentShader(r io.Reader) PipelineOption {
	var s shader
	var err error
	s.stages = C.GL_FRAGMENT_SHADER_BIT
	s.shader, err = newShader(C.GL_FRAGMENT_SHADER, r)
	if err != nil {
		setErr(fmt.Errorf("unable to compile fragment shader: %s", err))
	}
	return func(p *Pipeline) {
		p.attachShader(s)
	}
}

// GeometryShader compiles a geometry shader.
func GeometryShader(r io.Reader) PipelineOption {
	var s shader
	var err error
	s.stages = C.GL_GEOMETRY_SHADER_BIT
	s.shader, err = newShader(C.GL_GEOMETRY_SHADER, r)
	if err != nil {
		setErr(fmt.Errorf("unable to compile geometry shader: %s", err))
	}
	return func(p *Pipeline) {
		p.attachShader(s)
	}
}

// TessControlShader compiles a tesselation control shader.
func TessControlShader(r io.Reader) PipelineOption {
	var s shader
	var err error
	s.stages = C.GL_TESS_CONTROL_SHADER_BIT
	s.shader, err = newShader(C.GL_TESS_CONTROL_SHADER, r)
	if err != nil {
		setErr(fmt.Errorf("unable to compile tesselation control shader: %s", err))
	}
	return func(p *Pipeline) {
		p.attachShader(s)
	}
}

// TessEvaluationShader compiles a tesselation evaluation shader.
func TessEvaluationShader(r io.Reader) PipelineOption {
	var s shader
	var err error
	s.stages = C.GL_TESS_EVALUATION_SHADER_BIT
	s.shader, err = newShader(C.GL_TESS_EVALUATION_SHADER, r)
	if err != nil {
		setErr(fmt.Errorf("unable to compile tesselation evaluation shader: %s", err))
	}
	return func(p *Pipeline) {
		p.attachShader(s)
	}
}

// ComputeShader compiles a comput shader.
func ComputeShader(r io.Reader) PipelineOption {
	var s shader
	var err error
	s.stages = C.GL_COMPUTE_SHADER_BIT
	s.shader, err = newShader(C.GL_COMPUTE_SHADER, r)
	if err != nil {
		setErr(fmt.Errorf("unable to compile compute shader: %s", err))
	}
	return func(p *Pipeline) {
		p.attachShader(s)
	}
}

func newShader(t uint32, r io.Reader) (C.GLuint, error) {
	b, err := ioutil.ReadAll(r)
	if err != nil {
		return 0, fmt.Errorf("failed to read shader: %s", err)
	}
	cb := C.CString(string(b))
	defer C.free(unsafe.Pointer(cb))

	s := C.CompileShader(C.GLenum(t), (*C.GLchar)(unsafe.Pointer(cb)))
	if errm := C.ShaderCompileError(s); errm != nil {
		defer C.free(unsafe.Pointer(errm))
		return 0, errors.New(C.GoString(errm))
	}

	return s, nil
}

//------------------------------------------------------------------------------
