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
	GLuint s = glCreateShaderProgramv(t, 1, &b);
	if (s == 0) {
		return 0;
	}

	return s;
}

char* ShaderLinkError(GLuint p) {
    GLint ok = GL_TRUE;
    glGetProgramiv (p, GL_LINK_STATUS, &ok);
    if (ok != GL_TRUE)
    {
        GLint l = 0;
        glGetProgramiv (p, GL_INFO_LOG_LENGTH, &l);
        char *m = calloc(l + 1, sizeof(char));
        glGetProgramInfoLog (p, l, &l, m);
        return m;
    }

	return NULL;
}
*/
import "C"

//------------------------------------------------------------------------------

// A Shader is a compiled program run by the GPU.
type Shader struct {
	shader C.GLuint
	stages C.GLenum
}

//------------------------------------------------------------------------------

// NewVertexShader compiles a vertex shader.
func NewVertexShader(r io.Reader) Shader {
	var s Shader
	var err error
	s.stages = C.GL_VERTEX_SHADER_BIT
	s.shader, err = newShader(C.GL_VERTEX_SHADER, r)
	if err != nil {
		setErr(fmt.Errorf("error in vertex shader: %s", err))
	}
	return s
}

// NewFragmentShader compiles a fragment shader.
func NewFragmentShader(r io.Reader) Shader {
	var s Shader
	var err error
	s.stages = C.GL_FRAGMENT_SHADER_BIT
	s.shader, err = newShader(C.GL_FRAGMENT_SHADER, r)
	if err != nil {
		setErr(fmt.Errorf("error in fragment shader: %s", err))
	}
	return s
}

// NewGeometryShader compiles a geometry shader.
func NewGeometryShader(r io.Reader) Shader {
	var s Shader
	var err error
	s.stages = C.GL_GEOMETRY_SHADER_BIT
	s.shader, err = newShader(C.GL_GEOMETRY_SHADER, r)
	if err != nil {
		setErr(fmt.Errorf("error in geometry shader: %s", err))
	}
	return s
}

// NewTessControlShader compiles a tesselation control shader.
func NewTessControlShader(r io.Reader) Shader {
	var s Shader
	var err error
	s.stages = C.GL_TESS_CONTROL_SHADER_BIT
	s.shader, err = newShader(C.GL_TESS_CONTROL_SHADER, r)
	if err != nil {
		setErr(fmt.Errorf("error in tesselation control shader: %s", err))
	}
	return s
}

// NewTessEvaluationShader compiles a tesselation evaluation shader.
func NewTessEvaluationShader(r io.Reader) Shader {
	var s Shader
	var err error
	s.stages = C.GL_TESS_EVALUATION_SHADER_BIT
	s.shader, err = newShader(C.GL_TESS_EVALUATION_SHADER, r)
	if err != nil {
		setErr(fmt.Errorf("error in tesselation evaluation shader: %s", err))
	}
	return s
}

// NewComputeShader compiles a comput shader.
func NewComputeShader(r io.Reader) Shader {
	var s Shader
	var err error
	s.stages = C.GL_COMPUTE_SHADER_BIT
	s.shader, err = newShader(C.GL_COMPUTE_SHADER, r)
	if err != nil {
		setErr(fmt.Errorf("error in compute shader: %s", err))
	}
	return s
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
