// Copyright (c) 2013-2016 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package gfx

//------------------------------------------------------------------------------

import (
	"errors"
	"unsafe"

	"github.com/drakmaniso/carol/color"
)

/*
#include <stdlib.h>
#include "glad.h"

GLuint NewPipeline() {
	return glCreateProgram();
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

void PipelineAttachShader(GLuint p, GLenum stages, GLuint shader) {
	glAttachShader(p, shader);
}

void PipelineLinkProgram(GLuint p) {
	glLinkProgram(p);
}

typedef struct {
	// Input Assembly State
	GLenum    topology;
	GLboolean primitiveRestart;

	// Rasterization State
	GLboolean depthClamp;
	GLboolean  rasterizerDiscard;
	GLenum    cullMode;
	GLenum    frontFace;

	// Depth and Stencil State
	GLboolean	depthTest;
	GLenum    depthComparison;
	GLboolean	stencilTest;
} PipelineState;

PipelineState currentState;

static inline void BindPipeline(GLuint p, GLuint vao, PipelineState *state) {
	glUseProgram(p);
	glBindVertexArray(vao);

	// Input Assembly State

	if (state->topology != currentState.topology) {
		currentState.topology = state->topology;
	}

	if (state->primitiveRestart != currentState.primitiveRestart) {
		if (state->primitiveRestart) {
			glEnable(GL_PRIMITIVE_RESTART);
			glEnable(GL_PRIMITIVE_RESTART_FIXED_INDEX);
		} else {
			glDisable(GL_PRIMITIVE_RESTART);
		}
		currentState.primitiveRestart = state->primitiveRestart;
	}

	// Rasterization State

	if (state->depthClamp != currentState.depthClamp) {
		if (state->depthClamp) {
			glDisable(GL_DEPTH_CLAMP);
		} else {
			glEnable(GL_DEPTH_CLAMP);
		}
		currentState.depthClamp = state->depthClamp;
	}

	if (state->rasterizerDiscard != currentState.rasterizerDiscard) {
		if (state->rasterizerDiscard) {
			glEnable(GL_RASTERIZER_DISCARD);
		} else {
			glDisable(GL_RASTERIZER_DISCARD);
		}
		currentState.rasterizerDiscard = state->rasterizerDiscard;
	}

	if (state->cullMode != currentState.cullMode) {
		switch(state->cullMode) {
		case GL_BACK:
			glCullFace(GL_BACK);
			glEnable(GL_CULL_FACE);
			break;
		case GL_NONE:
			glDisable(GL_CULL_FACE);
			break;
		case GL_FRONT:
			glCullFace(GL_FRONT);
			glEnable(GL_CULL_FACE);
			break;
		case GL_FRONT_AND_BACK:
			glCullFace(GL_FRONT_AND_BACK);
			glEnable(GL_CULL_FACE);
			break;
		}
		currentState.cullMode = state->cullMode;
	}

	if (state->frontFace != currentState.frontFace) {
		glFrontFace(state->frontFace);
		currentState.frontFace = state->frontFace;
	}

	// Depth and Stencil State

	if (state->depthTest != currentState.depthTest) {
		if (state->depthTest) {
			glEnable(GL_DEPTH_TEST);
		} else {
			glDisable(GL_DEPTH_TEST);
		}
		currentState.depthTest = state->depthTest;
	}

	if (state->depthComparison != currentState.depthComparison) {
		glDepthFunc(state->depthComparison);
		currentState.depthComparison = state->depthComparison;
	}

	if (state->stencilTest != currentState.stencilTest) {
		if (state->stencilTest) {
			glEnable(GL_STENCIL_TEST);
		} else {
			glDisable(GL_STENCIL_TEST);
		}
		currentState.stencilTest = state->stencilTest;
	}
}

static inline void UnbindPipeline() {
	glUseProgram(0);
	glBindVertexArray(0);
}

void ClosePipeline(GLuint p, GLuint vao) {
	glDeleteVertexArrays(1, &vao);
	glDeleteProgramPipelines(1, &p);
}

void DeletePipelineProgram(GLuint p) {
	glDeleteProgramPipelines(1, &p);
}

void DeletePipelineVAO(GLuint vao) {
	glDeleteVertexArrays(1, &vao);
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

func init() {
	// VkPipelineInputAssemblyStateCreateInfo
	C.currentState.topology = C.GL_TRIANGLES
	C.currentState.primitiveRestart = C.GL_FALSE
	// VkPipelineRasterizationStateCreateInfo
	C.currentState.depthClamp = C.GL_TRUE
	C.currentState.rasterizerDiscard = C.GL_FALSE
	C.currentState.cullMode = C.GL_NONE
	C.currentState.frontFace = C.GL_CCW
	// VkPipelineDepthStencilStateCreateInfo
	C.currentState.depthTest = C.GL_FALSE
	C.currentState.depthComparison = C.GL_LESS
	C.currentState.stencilTest = C.GL_FALSE
}

//------------------------------------------------------------------------------

// A Pipeline consists of shaders and state for the GPU.
type Pipeline struct {
	object C.GLuint
	vao    C.GLuint
	state  C.PipelineState
}

type pipelineState struct {
	cullMode C.GLuint
}

//------------------------------------------------------------------------------

// A PipelineConfig represents a configuration option used when creating a new
// pipeline.
type PipelineConfig func(*Pipeline)

//------------------------------------------------------------------------------

// NewPipeline returns a pipeline with created from a specific set of shaders.
func NewPipeline(o ...PipelineConfig) *Pipeline {
	var p Pipeline

	// Input Assembly State
	p.state.topology = C.GL_TRIANGLES
	p.state.primitiveRestart = C.GL_FALSE
	// Rasterization State
	p.state.depthClamp = C.GL_TRUE
	p.state.rasterizerDiscard = C.GL_FALSE
	p.state.cullMode = C.GL_NONE
	p.state.frontFace = C.GL_CCW
	// DepthStencil State
	p.state.depthTest = C.GL_FALSE
	p.state.depthComparison = C.GL_LESS
	p.state.stencilTest = C.GL_FALSE

	p.object = C.NewPipeline() //TODO: Error Handling
	p.vao = C.CreateVAO()      //TODO: Error Handling
	oObj, oVao := p.object, p.vao
	for _, f := range o {
		f(&p)
	}
	C.PipelineLinkProgram(p.object)
	if errm := C.PipelineLinkError(p.object); errm != nil {
		defer C.free(unsafe.Pointer(errm))
		setErr("linking shaders", errors.New(C.GoString(errm)))
	}
	// A bit inelegant, but makes the API easier
	if oObj != p.object {
		C.DeletePipelineProgram(oObj)
	}
	if oVao != p.vao {
		C.DeletePipelineVAO(oVao)
	}

	return &p
}

func (p *Pipeline) attachShader(s shader) {
	C.PipelineAttachShader(p.object, s.stages, s.shader)
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
	if currentPipeline != p {
		C.BindPipeline(p.object, p.vao, &p.state)
		currentPipeline = p
	}
}

// Unbind the pipeline.
func (p *Pipeline) Unbind() {
	C.UnbindPipeline()
	currentPipeline = nil
}

var currentPipeline *Pipeline

//------------------------------------------------------------------------------

// Close the pipeline.
func (p *Pipeline) Close() {
	C.ClosePipeline(p.object, p.vao)
}

//------------------------------------------------------------------------------

// Input Assembly State

func Topology(m Primitive) PipelineConfig {
	return func(p *Pipeline) {
		p.state.topology = C.GLenum(m)
	}
}

// A Primitive specifies what kind of object to draw.
type Primitive C.GLenum

// Used in `Topology`.
const (
	Points               Primitive = C.GL_POINTS
	Lines                Primitive = C.GL_LINES
	LineLoop             Primitive = C.GL_LINE_LOOP
	LineStrip            Primitive = C.GL_LINE_STRIP
	Triangles            Primitive = C.GL_TRIANGLES
	TriangleStrip        Primitive = C.GL_TRIANGLE_STRIP
	TriangleFan          Primitive = C.GL_TRIANGLE_FAN
	LinesAdjency         Primitive = C.GL_LINES_ADJACENCY
	LineStripAdjency     Primitive = C.GL_LINE_STRIP_ADJACENCY
	TrianglesAdjency     Primitive = C.GL_TRIANGLES_ADJACENCY
	TriangleStripAdjency Primitive = C.GL_TRIANGLE_STRIP_ADJACENCY
	Patches              Primitive = C.GL_PATCHES
)

func PrimitiveRestart(enable bool) PipelineConfig {
	if enable {
		return func(p *Pipeline) {
			p.state.primitiveRestart = C.GL_TRUE
		}
	}
	return func(p *Pipeline) {
		p.state.primitiveRestart = C.GL_FALSE
	}
}

//------------------------------------------------------------------------------

// Rasterization State

func DepthClamp(enable bool) PipelineConfig {
	if enable {
		return func(p *Pipeline) {
			p.state.depthClamp = C.GL_TRUE
		}
	}
	return func(p *Pipeline) {
		p.state.depthClamp = C.GL_FALSE
	}
}

func RasterizerDiscard(enable bool) PipelineConfig {
	if enable {
		return func(p *Pipeline) {
			p.state.rasterizerDiscard = C.GL_TRUE
		}
	}
	return func(p *Pipeline) {
		p.state.rasterizerDiscard = C.GL_FALSE
	}
}

// CullFace specifies if front and/or back faces are culled.
//
// See also `FrontFace`.
func CullFace(front, back bool) PipelineConfig {
	switch {
	case front && back:
		return func(p *Pipeline) {
			p.state.cullMode = C.GL_FRONT_AND_BACK
		}
	case front:
		return func(p *Pipeline) {
			p.state.cullMode = C.GL_FRONT
		}
	case back:
		return func(p *Pipeline) {
			p.state.cullMode = C.GL_BACK
		}
	default:
		return func(p *Pipeline) {
			p.state.cullMode = C.GL_NONE
		}
	}
}

// FrontFace specifies which winding direction is considered front.
//
// See also `CullFace`.
func FrontFace(d WindingDirection) PipelineConfig {
	return func(p *Pipeline) {
		p.state.frontFace = C.GLenum(d)
	}
}

// A WindingDirection specifies a rotation direction.
type WindingDirection C.GLenum

// Used in `FrontFace`.
const (
	Clockwise        WindingDirection = C.GL_CW
	CounterClockwise WindingDirection = C.GL_CCW
)

//------------------------------------------------------------------------------

// Depth and Stencil State

func DepthTest(enable bool) PipelineConfig {
	if enable {
		return func(p *Pipeline) {
			p.state.depthTest = C.GL_TRUE
		}
	}
	return func(p *Pipeline) {
		p.state.depthTest = C.GL_FALSE
	}
}

// DepthComparison specifies the function used to compare pixel depth.
//
// Note that you must also `Enable(DepthTest)`. The default value is `Less`.
func DepthComparison(op ComparisonOp) PipelineConfig {
	return func(p *Pipeline) {
		p.state.depthComparison = C.GLenum(op)
	}
}

func StencilTest(enable bool) PipelineConfig {
	if enable {
		return func(p *Pipeline) {
			p.state.stencilTest = C.GL_TRUE
		}
	}
	return func(p *Pipeline) {
		p.state.stencilTest = C.GL_FALSE
	}
}

//------------------------------------------------------------------------------

func ShareShadersWith(other *Pipeline) PipelineConfig {
	return func(p *Pipeline) {
		p.object = other.object
	}
}

func ShareVertexFormatsWith(other *Pipeline) PipelineConfig {
	return func(p *Pipeline) {
		p.vao = other.vao
	}
}

//------------------------------------------------------------------------------
