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
*/
import "C"

//------------------------------------------------------------------------------

// Enable an OpenGL capability
func Enable(c capability) {
	C.Enable((C.GLenum)(c))
}

//------------------------------------------------------------------------------

type capability C.GLenum

// OpenGL Capabilities
const (
	Blend                      capability = C.GL_BLEND
	ColorLogicOp               capability = C.GL_COLOR_LOGIC_OP
	CullFace                   capability = C.GL_CULL_FACE
	DebugOutput                capability = C.GL_DEBUG_OUTPUT
	DebugOutputSynchronous     capability = C.GL_DEBUG_OUTPUT_SYNCHRONOUS
	DepthClamp                 capability = C.GL_DEPTH_CLAMP
	DepthTest                  capability = C.GL_DEPTH_TEST
	Dither                     capability = C.GL_DITHER
	FramebufferSRGB            capability = C.GL_FRAMEBUFFER_SRGB
	LineSmooth                 capability = C.GL_LINE_SMOOTH
	Multisample                capability = C.GL_MULTISAMPLE
	PrimitiveRestart           capability = C.GL_PRIMITIVE_RESTART
	PrimitiveRestartFixedIndex capability = C.GL_PRIMITIVE_RESTART_FIXED_INDEX
	RasterixerDiscard          capability = C.GL_RASTERIZER_DISCARD
	SampleAlphaToCoverage      capability = C.GL_SAMPLE_ALPHA_TO_COVERAGE
	SampleAlphaToOne           capability = C.GL_SAMPLE_ALPHA_TO_ONE
	SampleCoverage             capability = C.GL_SAMPLE_COVERAGE
	SampleShading              capability = C.GL_SAMPLE_SHADING
	SampleMask                 capability = C.GL_SAMPLE_MASK
	ScissorTest                capability = C.GL_SCISSOR_TEST
	StencilTest                capability = C.GL_STENCIL_TEST
	TextureCubeMapSeamless     capability = C.GL_TEXTURE_CUBE_MAP_SEAMLESS
)

//------------------------------------------------------------------------------
