// Copyright (c) 2013-2016 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package gl

//------------------------------------------------------------------------------

import (
	"unsafe"
)

//------------------------------------------------------------------------------

/*
#include "glad.h"

static inline GLuint NewSampler() {
	GLuint s;
	glCreateSamplers(1, &s);
	return s;
}

static inline void SamplerMinification(GLuint object, GLenum fm) {
	glSamplerParameteri(object, GL_TEXTURE_MIN_FILTER, fm);
}

static inline void SamplerMagnification(GLuint object, GLenum fm) {
	glSamplerParameteri(object, GL_TEXTURE_MAG_FILTER, fm);
}

static inline void SamplerLevelOfDetail(GLuint object, GLfloat min, GLfloat max) {
	glSamplerParameterf(object, GL_TEXTURE_MIN_LOD, min);
	glSamplerParameterf(object, GL_TEXTURE_MAX_LOD, max);
}

static inline void SamplerAnisotropy(GLuint object, GLfloat max) {
	glSamplerParameterf(object, 0x84FE, max); //GL_TEXTURE_MAX_ANISOTROPY_EXT
}

static inline void SamplerWrapping(GLuint object, GLenum s, GLenum t, GLenum r) {
	glSamplerParameteri(object, GL_TEXTURE_WRAP_S, s);
	glSamplerParameteri(object, GL_TEXTURE_WRAP_T, t);
	glSamplerParameteri(object, GL_TEXTURE_WRAP_R, r);
}

static inline void SamplerBorderColor(GLuint object, GLfloat* c) {
	glSamplerParameterfv(object, GL_TEXTURE_BORDER_COLOR, c);
}

static inline void SamplerComparison(GLuint object, GLenum cf) {
	glSamplerParameteri(object, GL_TEXTURE_COMPARE_MODE, GL_COMPARE_REF_TO_TEXTURE);
	glSamplerParameteri(object, GL_TEXTURE_COMPARE_MODE, cf);
}

static inline void SamplerBind(GLuint binding, GLuint sampler) {
	glBindSampler(binding, sampler);
}

*/
import "C"

//------------------------------------------------------------------------------

// A Sampler describes a way to sample textures inside shaders.
type Sampler struct {
	object C.GLuint
}

// A SamplerOption is a setting used when creating a new `Sampler`.
type SamplerOption func(sa *Sampler)

//------------------------------------------------------------------------------

// NewSampler returns a new sampler.
func NewSampler(o ...SamplerOption) Sampler {
	var sa Sampler
	sa.object = C.NewSampler()
	for _, f := range o {
		f(&sa)
	}
	return sa
}

//------------------------------------------------------------------------------

// Minification specifies which filter is used when minifying the texture.
//
// The default value is `NearestMipmapLinear`.
func Minification(fm FilterMode) SamplerOption {
	return func(sa *Sampler) {
		C.SamplerMinification(sa.object, C.GLenum(fm))
	}
}

// Magnification specifies which filter is used when minifying the texture.
//
// The default value is `Linear`.
func Magnification(fm FilterMode) SamplerOption {
	return func(sa *Sampler) {
		C.SamplerMagnification(sa.object, C.GLenum(fm))
	}
}

// A FilterMode specifies how to filter textures when minifying or magnifying.
type FilterMode C.GLuint

// Used in `Sampler.Minification` and `Sampler.Magnification` (only `Nearest`
// and `Linear` are valid for magnification).
const (
	Nearest             FilterMode = C.GL_NEAREST
	Linear              FilterMode = C.GL_LINEAR
	NearestMimapNearest FilterMode = C.GL_NEAREST_MIPMAP_NEAREST
	LinearMipmapNearest FilterMode = C.GL_LINEAR_MIPMAP_NEAREST
	NearestMipmapLinear FilterMode = C.GL_NEAREST_MIPMAP_LINEAR
	LinearMipmapLinear  FilterMode = C.GL_LINEAR_MIPMAP_LINEAR
)

//------------------------------------------------------------------------------

// LevelOfDetail specifies the minimum and maximum LOD to use.
//
// The default values are -1000 and 1000.
func LevelOfDetail(min, max float32) SamplerOption {
	return func(sa *Sampler) {
		C.SamplerLevelOfDetail(sa.object, C.GLfloat(min), C.GLfloat(max))
	}
}

//------------------------------------------------------------------------------

// Anisotropy specifies the maximum anisotropy level.
func Anisotropy(max float32) SamplerOption {
	return func(sa *Sampler) {
		C.SamplerAnisotropy(sa.object, C.GLfloat(max))
	}
}

//------------------------------------------------------------------------------

// Wrapping sets the wrapping modes for texture coordinates.
//
// The default value is `Repeat` for all coordinates.
func Wrapping(s, t, p WrapMode) SamplerOption {
	return func(sa *Sampler) {
		C.SamplerWrapping(sa.object, C.GLenum(s), C.GLenum(t), C.GLenum(p))
	}
}

// A WrapMode specifies the way a texture wraps.
type WrapMode C.GLuint

// Used in `Sampler.Wrap`.
const (
	ClampToBorder     WrapMode = C.GL_CLAMP_TO_BORDER
	ClampToEdge       WrapMode = C.GL_CLAMP_TO_EDGE
	MirrorClampToEdge WrapMode = C.GL_MIRROR_CLAMP_TO_EDGE
	Repeat            WrapMode = C.GL_REPEAT
	MirroredRepeat    WrapMode = C.GL_MIRRORED_REPEAT
)

// BorderColor sets the color used for texture filtering when ClampToborder
// wrapping mode is used.
func BorderColor(color struct{ R, G, B, A float32 }) SamplerOption {
	return func(sa *Sampler) {
		C.SamplerBorderColor(sa.object, (*C.GLfloat)(unsafe.Pointer(&color)))
	}
}

//------------------------------------------------------------------------------

// Comparison specifies the mode and operator used when comparing depth
// textures.
func Comparison(op ComparisonOp) SamplerOption {
	return func(sa *Sampler) {
		C.SamplerComparison(sa.object, C.GLenum(op))
	}
}

// A ComparisonOp specifies an operator for depth texture comparison.
type ComparisonOp C.GLuint

// Used in `Sampler.Comparison` and `DepthComparison`.
const (
	LessOrEqual    ComparisonOp = C.GL_LEQUAL
	GreaterOrEqual ComparisonOp = C.GL_GEQUAL
	Less           ComparisonOp = C.GL_LESS
	Greater        ComparisonOp = C.GL_GREATER
	Equal          ComparisonOp = C.GL_EQUAL
	NotEqual       ComparisonOp = C.GL_NOTEQUAL
	Always         ComparisonOp = C.GL_ALWAYS
	Never          ComparisonOp = C.GL_NEVER
)

//------------------------------------------------------------------------------

// Bind a sampler to a texture unit index.
func (sa Sampler) Bind(binding uint32) {
	C.SamplerBind(C.GLuint(binding), sa.object)
}

//------------------------------------------------------------------------------
