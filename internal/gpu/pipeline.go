// Copyright (c) 2013-2017 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package gpu

//------------------------------------------------------------------------------

/*
#include "glad.h"

static inline void SetupImagePipeline(GLuint *program, GLuint*vao, const GLchar* vs, const GLchar* fs) {
	*program = glCreateProgram();
	glCreateVertexArrays(1, vao);

	GLuint vso = glCreateShader(GL_VERTEX_SHADER);
	const GLchar*vsb[] = {vs};
	glShaderSource(vso, 1, vsb, NULL);
	glCompileShader(vso);
	//TODO: error handling

	GLuint fso = glCreateShader(GL_FRAGMENT_SHADER);
	const GLchar*fsb[] = {fs};
	glShaderSource(fso, 1, fsb, NULL);
	glCompileShader(fso);
	//TODO: error handling

	glAttachShader(*program, vso);
	glAttachShader(*program, fso);
	glLinkProgram(*program);
	//TODO: error handling
}

static inline void BindImagePipeline(GLuint program, GLuint vao) {
	glUseProgram(program);
	glBindVertexArray(vao);
	glDrawArrays(GL_TRIANGLES, 0, 3);
}
*/
import "C"

//------------------------------------------------------------------------------

import (
	"unsafe"
)

//------------------------------------------------------------------------------

type ImagePaint struct {
	//  word
	X1, Y1 int16

	//  word
	X2, Y2 int16

	//  word
	Xtex, Ytex int16

	//  word
	format, index uint16

	// word
	user           uint16
	ZoomAndPalette uint16
}

//------------------------------------------------------------------------------

var ImagePipeline struct {
	program C.GLuint
	vao     C.GLuint
}

func SetupImagePipeline() {
	vs := C.CString(string(vertexShader))
	defer C.free(unsafe.Pointer(vs))
	fs := C.CString(string(fragmentShader))
	defer C.free(unsafe.Pointer(fs))

	C.SetupImagePipeline(
		&ImagePipeline.program,
		&ImagePipeline.vao,
		(*C.GLchar)(vs),
		(*C.GLchar)(fs),
	)
}

//------------------------------------------------------------------------------

func BindImagePipeline() {
	C.BindImagePipeline(ImagePipeline.program, ImagePipeline.vao)
}

//------------------------------------------------------------------------------

const vertexShader = `#version 450 core

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
