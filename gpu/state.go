// Copyright (c) 2013-2016 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package gpu

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

static inline void Blending(GLenum src, GLenum dst) {
	glBlendFunc(src, dst);
}

static inline void PointSize(GLfloat s) {
	glPointSize(s);
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
	Blend                  Capability = C.GL_BLEND
	ColorLogicOp           Capability = C.GL_COLOR_LOGIC_OP
	DebugOutput            Capability = C.GL_DEBUG_OUTPUT
	DebugOutputSynchronous Capability = C.GL_DEBUG_OUTPUT_SYNCHRONOUS
	Dither                 Capability = C.GL_DITHER
	FramebufferSRGB        Capability = C.GL_FRAMEBUFFER_SRGB
	LineSmooth             Capability = C.GL_LINE_SMOOTH
	Multisample            Capability = C.GL_MULTISAMPLE
	SampleAlphaToCoverage  Capability = C.GL_SAMPLE_ALPHA_TO_COVERAGE
	SampleAlphaToOne       Capability = C.GL_SAMPLE_ALPHA_TO_ONE
	SampleCoverage         Capability = C.GL_SAMPLE_COVERAGE
	SampleShading          Capability = C.GL_SAMPLE_SHADING
	SampleMask             Capability = C.GL_SAMPLE_MASK
	ScissorTest            Capability = C.GL_SCISSOR_TEST
	TextureCubeMapSeamless Capability = C.GL_TEXTURE_CUBE_MAP_SEAMLESS
)

//------------------------------------------------------------------------------

// Blending specifies the formula used for blending pixels.
//
// Note that you must also `Enable(Blend)`. The default values are `One` and
// `Zero`. For alpha-blending and antialiasing, the most useful choice is
// `Blending(SrcAlpha, OneMinusSrcAlpha)`.
func Blending(src, dst BlendFactor) {
	C.Blending(C.GLenum(src), C.GLenum(dst))
}

// A BlendFactor is a formula used when blending pixels.
type BlendFactor C.GLenum

// Used in `Blending`.
const (
	Zero                  BlendFactor = C.GL_ZERO
	One                   BlendFactor = C.GL_ONE
	SrcColor              BlendFactor = C.GL_SRC_COLOR
	OneMinusSrcColor      BlendFactor = C.GL_ONE_MINUS_SRC_COLOR
	DstColor              BlendFactor = C.GL_DST_COLOR
	OneMinusDstColor      BlendFactor = C.GL_ONE_MINUS_DST_COLOR
	SrcAlpha              BlendFactor = C.GL_SRC_ALPHA
	OneMinusSrcAlpha      BlendFactor = C.GL_ONE_MINUS_SRC_ALPHA
	DstAlpha              BlendFactor = C.GL_DST_ALPHA
	OneMinusDstAlpha      BlendFactor = C.GL_ONE_MINUS_DST_ALPHA
	ConstantColor         BlendFactor = C.GL_CONSTANT_COLOR
	OneMinusConstantColor BlendFactor = C.GL_ONE_MINUS_CONSTANT_COLOR
	ConstantAlpha         BlendFactor = C.GL_CONSTANT_ALPHA
	OneMinusConstantAlpha BlendFactor = C.GL_ONE_MINUS_CONSTANT_ALPHA
	AlphaSaturate         BlendFactor = C.GL_SRC_ALPHA_SATURATE
	Src1Color             BlendFactor = C.GL_SRC1_COLOR
	OneMinusSrc1Color     BlendFactor = C.GL_ONE_MINUS_SRC1_COLOR
	Src1Alpha             BlendFactor = C.GL_SRC1_ALPHA
	OneMinuxSrc1Alpha     BlendFactor = C.GL_ONE_MINUS_SRC1_ALPHA
)

//------------------------------------------------------------------------------

// PointSize sets the size used when drawing points.
func PointSize(s float32) {
	C.PointSize(C.GLfloat(s))
}

//------------------------------------------------------------------------------
