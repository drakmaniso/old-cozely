// Copyright (c) 2013-2017 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package gpu

//------------------------------------------------------------------------------

/*
#include "glad.h"

GLuint CompileShader(GLenum t, const GLchar* b) {
	GLuint s = glCreateShader(t);
	const GLchar*bb[] = {b};
	glShaderSource(s, 1, bb, NULL);
	glCompileShader(s);

	return s;
}

static inline void SetupStampPipeline(GLuint *program, GLuint*vao, GLuint vso, GLuint fso) {
	*program = glCreateProgram();
	glCreateVertexArrays(1, vao);

	glAttachShader(*program, vso);
	glAttachShader(*program, fso);
	glLinkProgram(*program);
	//TODO: error handling
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

static inline void BindStampPipeline(GLuint program, GLuint vao) {
	glUseProgram(program);
	glBindVertexArray(vao);
	glDrawArrays(GL_TRIANGLES, 0, 3);
}

*/
import "C"

//------------------------------------------------------------------------------

import (
	"errors"
	"unsafe"

	"github.com/drakmaniso/carol/internal/core"
)

//------------------------------------------------------------------------------

type Stamp struct {
	//  word
	Address uint32

	//  word
	W, H int16

	//  word
	X, Y int16

	// word
	OrientAndUser  uint16
	ZoomAndPalette uint16
}

//------------------------------------------------------------------------------

var StampPipeline struct {
	program C.GLuint
	vao     C.GLuint
}

func SetupStampPipeline() error {
	vs := C.CString(string(vertexShader))
	defer C.free(unsafe.Pointer(vs))
	vso := C.CompileShader(C.GL_VERTEX_SHADER, (*C.GLchar)(vs))
	if errm := C.ShaderCompileError(vso); errm != nil {
		defer C.free(unsafe.Pointer(errm))
		return core.Error("while compiling vertex shader for stamp pipeline", errors.New(C.GoString(errm)))
	}

	fs := C.CString(string(fragmentShader))
	defer C.free(unsafe.Pointer(fs))
	fso := C.CompileShader(C.GL_FRAGMENT_SHADER, (*C.GLchar)(fs))
	if errm := C.ShaderCompileError(fso); errm != nil {
		defer C.free(unsafe.Pointer(errm))
		return core.Error("while compiling fragment shader for stamp pipeline", errors.New(C.GoString(errm)))
	}

	C.SetupStampPipeline(
		&StampPipeline.program,
		&StampPipeline.vao,
		vso,
		fso,
	)
	if errm := C.PipelineLinkError(StampPipeline.program); errm != nil {
		defer C.free(unsafe.Pointer(errm))
		return core.Error("while linking shaders for stamp pipeline", errors.New(C.GoString(errm)))
	}

	return nil
}

//------------------------------------------------------------------------------

func BindStampPipeline() {
	C.BindStampPipeline(StampPipeline.program, StampPipeline.vao)
}

//------------------------------------------------------------------------------

const vertexShader = `#version 450 core

struct Stamp {
	uint Address;
	uint WH;
	uint XY;
	uint TransformPalette;
};
layout(std430, binding = 0) buffer StampBuffer {
	Stamp []Stamps;
};

out gl_PerVertex {
	vec4 gl_Position;
};

void main(void)
{
	const vec4 triangle[3] = vec4[3](
		vec4(0, 0.4, 0.5, 1),
		vec4(-0.8, -0.4, 0.5, 1),
		vec4(0.8, -0.4, 0.5, 1)
	);
	gl_Position = triangle[gl_VertexID];
}
`

//------------------------------------------------------------------------------

const fragmentShader = `#version 450 core

// in vec4 gl_FragCoord;

out vec4 color;

float rand(vec2 c){
	return fract(sin(dot(c ,vec2(12.9898,78.233))) * 43758.5453);
}

void main(void)
{
	color = vec4(
		0.5 + 0.25*rand(vec2(0.3, rand(gl_FragCoord.xy))),
		0.5 + 0.25*rand(vec2(0.1, rand(gl_FragCoord.xy))),
		0.5 + 0.25*rand(vec2(0.6, rand(gl_FragCoord.xy))),
		1.0
	);
}
`

//------------------------------------------------------------------------------
