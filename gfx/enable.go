// Copyright (c) 2013-2016 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package gfx

//------------------------------------------------------------------------------

/*
#include <stdlib.h>
#include "glad.h"

static inline void Enable(GLenum c) {
	glEnable(c);
}

static inline void Disable(GLenum c) {
	glDisable(c);
}

static inline void CullFace(GLenum f) {
	glCullFace(f);
}

static inline void FrontFace(GLenum d) {
	glFrontFace(d);
}
*/
import "C"

//------------------------------------------------------------------------------

// Enable an OpenGL Capability
func Enable(c Capability) {
	C.Enable((C.GLenum)(c))
}

// Disable an OpenGL Capability
func Disable(c Capability) {
	C.Disable((C.GLenum)(c))
}

//------------------------------------------------------------------------------

// A Capability is an OpenGL functionality that can be enabled or disabled.
type Capability C.GLenum

// Used in `Enable` and `Disable`.
const (
	Blend                      Capability = C.GL_BLEND
	ColorLogicOp               Capability = C.GL_COLOR_LOGIC_OP
	DebugOutput                Capability = C.GL_DEBUG_OUTPUT
	DebugOutputSynchronous     Capability = C.GL_DEBUG_OUTPUT_SYNCHRONOUS
	DepthClamp                 Capability = C.GL_DEPTH_CLAMP
	DepthTest                  Capability = C.GL_DEPTH_TEST
	Dither                     Capability = C.GL_DITHER
	FramebufferSRGB            Capability = C.GL_FRAMEBUFFER_SRGB
	LineSmooth                 Capability = C.GL_LINE_SMOOTH
	Multisample                Capability = C.GL_MULTISAMPLE
	PrimitiveRestart           Capability = C.GL_PRIMITIVE_RESTART
	PrimitiveRestartFixedIndex Capability = C.GL_PRIMITIVE_RESTART_FIXED_INDEX
	RasterixerDiscard          Capability = C.GL_RASTERIZER_DISCARD
	SampleAlphaToCoverage      Capability = C.GL_SAMPLE_ALPHA_TO_COVERAGE
	SampleAlphaToOne           Capability = C.GL_SAMPLE_ALPHA_TO_ONE
	SampleCoverage             Capability = C.GL_SAMPLE_COVERAGE
	SampleShading              Capability = C.GL_SAMPLE_SHADING
	SampleMask                 Capability = C.GL_SAMPLE_MASK
	ScissorTest                Capability = C.GL_SCISSOR_TEST
	StencilTest                Capability = C.GL_STENCIL_TEST
	TextureCubeMapSeamless     Capability = C.GL_TEXTURE_CUBE_MAP_SEAMLESS
)

//------------------------------------------------------------------------------

// CullFace specifies if front and/or back faces are culled.
//
// See also `FrontFace`.
func CullFace(front, back bool) {
	switch {
	case front && back:
		C.CullFace(C.GL_FRONT_AND_BACK)
		C.Enable(C.GL_CULL_FACE)
	case front:
		C.CullFace(C.GL_FRONT)
		C.Enable(C.GL_CULL_FACE)
	case back:
		C.CullFace(C.GL_BACK)
		C.Enable(C.GL_CULL_FACE)
	default:
		C.Disable(C.GL_CULL_FACE)
	}
}

// FrontFace specifies which winding direction is considered front.
//
// See also `CullFace`.
func FrontFace(d WindingDirection) {
	C.FrontFace(C.GLenum(d))
}

// A WindingDirection specifies a rotation direction.
type WindingDirection C.GLenum

// Used in `FrontFace`.
const (
	Clockwise        WindingDirection = C.GL_CW
	CounterClockwise WindingDirection = C.GL_CCW
)

//------------------------------------------------------------------------------
