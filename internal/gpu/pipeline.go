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

static inline void CreateStampPipeline(GLuint *program, GLuint*vao, GLuint vso, GLuint fso) {
	*program = glCreateProgram();
	glCreateVertexArrays(1, vao);
	glBindVertexArray(*vao);

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

static inline void BindStampPipeline(GLuint program, GLuint vao, unsigned int nbStamps) {
	glEnable(GL_BLEND);
	glBlendEquationSeparate(GL_FUNC_ADD, GL_FUNC_ADD);
	glBlendFuncSeparate(GL_SRC_ALPHA, GL_ONE_MINUS_SRC_ALPHA, GL_ONE, GL_ZERO);
	glUseProgram(program);
	glBindVertexArray(vao);
	glDrawArrays(GL_TRIANGLES, 0, nbStamps*6);
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
	Depth     int16
	Tint      uint8
	Transform byte
}

//------------------------------------------------------------------------------

var StampPipeline struct {
	program C.GLuint
	vao     C.GLuint
}

func createStampPipeline() error {
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

	C.CreateStampPipeline(
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
	updateStampBuffer(stamps)
	C.BindStampPipeline(StampPipeline.program, StampPipeline.vao, C.uint(len(stamps)))
	stamps = stamps[:0]
}

//------------------------------------------------------------------------------

var stamps = []Stamp{}

func Paint(addr uint32, w, h int16, x, y int16) {
	s := Stamp{Address: addr, W: w, H: h, X: x, Y: y, Tint: 0}
	stamps = append(stamps, s)
}

//------------------------------------------------------------------------------

const vertexShader = `#version 450 core

const vec2 PixelSize = vec2(1.0/320.0, 1.0/180.0);

struct Stamp {
	uint Address;
	uint WH;
	uint XY;
	uint DepthTintTrans;
};
layout(std430, binding = 0) buffer StampBuffer {
	Stamp []Stamps;
};

out gl_PerVertex {
	vec4 gl_Position;
};

out PerVertex {
	layout(location=0) vec2 UV;
	layout(location=1) flat uint Address;
	layout(location=2) flat uint Stride;
	layout(location=3) flat uint Depth;
	layout(location=4) flat uint Tint;
};

void main(void)
{
	// Calculate index in face buffer
	uint stampIndex = gl_VertexID / 6;

	int w = int(Stamps[stampIndex].WH & 0xFFFF);
	int h = int(Stamps[stampIndex].WH >> 16);
	//TODO: is it useful to extend sign on width and height?
	vec2 WH = vec2(
		w | (((w & 0x8000) >> 15) * 0xFFFF0000),
		h | (((h & 0x8000) >> 15) * 0xFFFF0000)
	);
	int x = int(Stamps[stampIndex].XY & 0xFFFF);
	int y = int(Stamps[stampIndex].XY >> 16);
	vec2 XY = vec2(
		x | (((x & 0x8000) >> 15) * 0xFFFF0000),
		y | (((y & 0x8000) >> 15) * 0xFFFF0000)
	);

	// Determine which corner of the stamp this is
	const uint [6]triangulate = {0, 1, 2, 0, 2, 3};
	uint currVert = triangulate[gl_VertexID - (6 * stampIndex)];

	const vec2 corners[4] = vec2[4](
		vec2(0, 0),
		vec2(1, 0),
		vec2(1, 1),
		vec2(0, 1)
	);
	vec2 p = (XY + corners[currVert] * WH) * PixelSize;
	gl_Position = vec4(p * vec2(2, -2) + vec2(-1,1), 0.5, 1);

	UV = corners[currVert] * WH;
	Address = Stamps[stampIndex].Address;
	Stride = uint(WH.x);
	Depth = Stamps[stampIndex].DepthTintTrans & 0xFFFF;
	Tint = (Stamps[stampIndex].DepthTintTrans >> 16) & 0xFF;
}
`

//------------------------------------------------------------------------------

const fragmentShader = `#version 450 core

in PerVertex {
	layout(location=0) vec2 UV;
	layout(location=1) flat uint Address;
	layout(location=2) flat uint Stride;
	layout(location=3) flat uint Depth;
	layout(location=4) flat uint Tint;
};

layout(std430, binding = 1) buffer PictureBuffer {
	uint []Pixels;
};

layout(std430, binding = 2) buffer PaletteBuffer {
	vec4 Colours[256];
};

out vec4 color;

uint getByte(uint addr) {
	uint waddr = addr >> 2;
	uint w = Pixels[waddr];
	w = w >> (8 * (addr & 0x3));
	return w & 0xFF;
}

uint getPixel(uint addr, uint stride, uint x, uint y) {
	return getByte(addr + x + y*stride);
}

float rand(vec2 c){
	return fract(sin(dot(c ,vec2(12.9898,78.233))) * 43758.5453);
}

void main(void)
{
	uint p = getPixel(Address, Stride, uint(UV.x), uint(UV.y));

	uint c;
	if (p == 0) {
		c = 0;
	} else {
		c = p + Tint;
		if (c > 255) {
			c -= 255;
		}
	}

	color = Colours[c];
}
`

//------------------------------------------------------------------------------
