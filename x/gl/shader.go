// Copyright (c) 2013-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package gl

import (
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"strings"
	"unsafe"

	"github.com/cozely/cozely/internal"
	"github.com/cozely/cozely/resource"
)

////////////////////////////////////////////////////////////////////////////////

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

void BindSubroutines(GLenum st, GLsizei c, void *ind) {
	glUniformSubroutinesuiv(st, c, (const GLuint *)ind);
}
*/
import "C"

////////////////////////////////////////////////////////////////////////////////

// A shader is a compiled program run by the GPU.
type shader struct {
	shader C.GLuint
	stages C.GLenum
}

////////////////////////////////////////////////////////////////////////////////

// Shader compiles a shader. The path is slash-separated, and the file extension
// determine the type of shader:
//
// - ".vert" for a vertex shader
// - ".frag" for a fragment shader
// - ".comp" for a compute shader
// - ".geom" for a geometry shader
// - ".tesc" for a tesselation control shader
// - ".tese" for a tesselation evaluation shader
func Shader(path string) PipelineConfig {
	f, err := resource.Open(path)
	if err != nil {
		setErr(internal.Wrap("gl shader file opening", err))
		return func(*Pipeline) {}
	}
	defer f.Close()
	switch {
	case strings.HasSuffix(path, ".vert"):
		return VertexShader(f)
	case strings.HasSuffix(path, ".frag"):
		return FragmentShader(f)
	case strings.HasSuffix(path, ".tesc"):
		return TessControlShader(f)
	case strings.HasSuffix(path, ".tese"):
		return TessEvaluationShader(f)
	case strings.HasSuffix(path, ".geom"):
		return GeometryShader(f)
	case strings.HasSuffix(path, ".comp"):
		return ComputeShader(f)
	}
	setErr(errors.New("gl shader file opening: unknown file extension"))
	return func(*Pipeline) {}
}

////////////////////////////////////////////////////////////////////////////////

// VertexShader compiles a vertex shader.
func VertexShader(r io.Reader) PipelineConfig {
	var s shader
	var err error
	s.stages = C.GL_VERTEX_SHADER_BIT
	s.shader, err = newShader(C.GL_VERTEX_SHADER, r)
	if err != nil {
		setErr(internal.Wrap("gl vertex shader compiling", err))
	}
	return func(p *Pipeline) {
		p.attachShader(s)
	}
}

// FragmentShader compiles a fragment shader.
func FragmentShader(r io.Reader) PipelineConfig {
	var s shader
	var err error
	s.stages = C.GL_FRAGMENT_SHADER_BIT
	s.shader, err = newShader(C.GL_FRAGMENT_SHADER, r)
	if err != nil {
		setErr(internal.Wrap("gl fragment shader compiling", err))
	}
	return func(p *Pipeline) {
		p.attachShader(s)
	}
}

// GeometryShader compiles a geometry shader.
func GeometryShader(r io.Reader) PipelineConfig {
	var s shader
	var err error
	s.stages = C.GL_GEOMETRY_SHADER_BIT
	s.shader, err = newShader(C.GL_GEOMETRY_SHADER, r)
	if err != nil {
		setErr(internal.Wrap("gl geometry shader compiling", err))
	}
	return func(p *Pipeline) {
		p.attachShader(s)
	}
}

// TessControlShader compiles a tesselation control shader.
func TessControlShader(r io.Reader) PipelineConfig {
	var s shader
	var err error
	s.stages = C.GL_TESS_CONTROL_SHADER_BIT
	s.shader, err = newShader(C.GL_TESS_CONTROL_SHADER, r)
	if err != nil {
		setErr(internal.Wrap("gl tesselation control shader compiling", err))
	}
	return func(p *Pipeline) {
		p.attachShader(s)
	}
}

// TessEvaluationShader compiles a tesselation evaluation shader.
func TessEvaluationShader(r io.Reader) PipelineConfig {
	var s shader
	var err error
	s.stages = C.GL_TESS_EVALUATION_SHADER_BIT
	s.shader, err = newShader(C.GL_TESS_EVALUATION_SHADER, r)
	if err != nil {
		setErr(internal.Wrap("gl tesselation evaluation shader compiling", err))
	}
	return func(p *Pipeline) {
		p.attachShader(s)
	}
}

// ComputeShader compiles a comput shader.
func ComputeShader(r io.Reader) PipelineConfig {
	var s shader
	var err error
	s.stages = C.GL_COMPUTE_SHADER_BIT
	s.shader, err = newShader(C.GL_COMPUTE_SHADER, r)
	if err != nil {
		setErr(internal.Wrap("gl compute shader compiling", err))
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

////////////////////////////////////////////////////////////////////////////////

// BindVertexSubroutines binds the vertex shader subroutines to indices in s.
func BindVertexSubroutines(s []uint32) {
	C.BindSubroutines(C.GL_VERTEX_SHADER, C.GLsizei(len(s)), unsafe.Pointer(&s[0]))
}

// BindFragmentSubroutines binds the fragment shader subroutines to indices in
// s.
func BindFragmentSubroutines(s []uint32) {
	C.BindSubroutines(C.GL_FRAGMENT_SHADER, C.GLsizei(len(s)), unsafe.Pointer(&s[0]))
}

// BindGeometrySubroutines binds the geometry shader subroutines to indices in
// s.
func BindGeometrySubroutines(s []uint32) {
	C.BindSubroutines(C.GL_GEOMETRY_SHADER, C.GLsizei(len(s)), unsafe.Pointer(&s[0]))
}

// BindTessControlSubroutines binds the tesselation control shader subroutines
// to indices in s.
func BindTessControlSubroutines(s []uint32) {
	C.BindSubroutines(C.GL_TESS_CONTROL_SHADER, C.GLsizei(len(s)), unsafe.Pointer(&s[0]))
}

// BindTessEvaluationSubroutines binds the tesselation evaluation shader
// subroutines to indices in s.
func BindTessEvaluationSubroutines(s []uint32) {
	C.BindSubroutines(C.GL_TESS_EVALUATION_SHADER, C.GLsizei(len(s)), unsafe.Pointer(&s[0]))
}

// BindComputeSubroutines binds the compute shader subroutines to indices in s.
func BindComputeSubroutines(s []uint32) {
	C.BindSubroutines(C.GL_COMPUTE_SHADER, C.GLsizei(len(s)), unsafe.Pointer(&s[0]))
}

////////////////////////////////////////////////////////////////////////////////
