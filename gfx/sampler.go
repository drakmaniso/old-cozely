// Copyright (c) 2013-2016 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package gfx

//------------------------------------------------------------------------------

import (
	"unsafe"

	"github.com/drakmaniso/glam/color"
)

//------------------------------------------------------------------------------

/*
#include "glad.h"

static inline GLuint NewSampler() {
	GLuint s;
	glCreateSamplers(1, &s);
	return s;
}

static inline void SamplerFilter(GLuint object, GLenum min, GLenum mag) {
	glSamplerParameteri(object, GL_TEXTURE_MIN_FILTER, min);
	glSamplerParameteri(object, GL_TEXTURE_MAG_FILTER, mag);
}

static inline void SamplerLevelOfDetail(GLuint object, GLfloat min, GLfloat max) {
	glSamplerParameterf(object, GL_TEXTURE_MIN_LOD, min);
	glSamplerParameterf(object, GL_TEXTURE_MAX_LOD, max);
}

static inline void SamplerAnisotropy(GLuint object, GLfloat max) {
	glSamplerParameterf(object, 0x84FE, max); //GL_TEXTURE_MAX_ANISOTROPY_EXT
}

static inline void SamplerWrap(GLuint object, GLenum s, GLenum t, GLenum r) {
	glSamplerParameteri(object, GL_TEXTURE_WRAP_S, s);
	glSamplerParameteri(object, GL_TEXTURE_WRAP_T, t);
	glSamplerParameteri(object, GL_TEXTURE_WRAP_R, r);
}

static inline void SamplerBorderColor(GLuint object, GLfloat* c) {
	glSamplerParameterfv(object, GL_TEXTURE_BORDER_COLOR, c);
}

static inline void SamplerCompareMode(GLuint object, GLenum m) {
	glSamplerParameteri(object, GL_TEXTURE_COMPARE_MODE, m);
}

static inline void SamplerCompareFunc(GLuint object, GLenum f) {
	glSamplerParameteri(object, GL_TEXTURE_COMPARE_MODE, f);
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

//------------------------------------------------------------------------------

// NewSampler returns a new sampler.
func NewSampler() Sampler {
	var s Sampler
	s.object = C.NewSampler()
	return s
}

//------------------------------------------------------------------------------

// Filter specifies which function is used when minifying and magnifying
// the texture.
func (sa *Sampler) Filter(min, mag FilterMode) {
	C.SamplerFilter(sa.object, C.GLenum(min), C.GLenum(mag))
}

// A FilterMode specifies how to filter textures when minifying or magnifying.
type FilterMode C.GLuint

// Used in `Sampler.Filter`.
const (
	Nearest             FilterMode = C.GL_NEAREST
	Linear              FilterMode = C.GL_LINEAR
	NearestMimapNearest FilterMode = C.GL_NEAREST_MIPMAP_NEAREST
	LinearMipmapNearest FilterMode = C.GL_LINEAR_MIPMAP_NEAREST
	NearestMipmapLinear FilterMode = C.GL_NEAREST_MIPMAP_LINEAR
	LinearMipmapLinear  FilterMode = C.GL_LINEAR_MIPMAP_LINEAR
)

// func MinFilterNearest() SamplerOption {
//   return func(sa *Sampler) {
//     C.SamplerMinFilter(sa.object, C.GL_NEAREST)
//   }
// }

//------------------------------------------------------------------------------

// LevelOfDetail specifies the minimum and maximum LOD to use.
func (sa *Sampler) LevelOfDetail(min, max float32) {
	C.SamplerLevelOfDetail(sa.object, C.GLfloat(min), C.GLfloat(max))
}

//------------------------------------------------------------------------------

// Anisotropy specifies the maximum anisotropy level.
func (sa *Sampler) Anisotropy(max float32) {
	C.SamplerAnisotropy(sa.object, C.GLfloat(max))
}

//------------------------------------------------------------------------------

// Wrap sets the wrapping modes for texture coordinates.
func (sa *Sampler) Wrap(s, t, p WrapMode) {
	C.SamplerWrap(sa.object, C.GLenum(s), C.GLenum(t), C.GLenum(p))
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
func (sa *Sampler) BorderColor(c color.RGBA) {
	C.SamplerBorderColor(sa.object, (*C.GLfloat)(unsafe.Pointer(&c)))
}

//------------------------------------------------------------------------------

// CompareMode specifies the texture comparison mode.
func (sa *Sampler) CompareMode(m CompareMode) {
	C.SamplerCompareMode(sa.object, C.GLenum(m))
}

// A CompareMode specifies a mode of texture comparison.
type CompareMode C.GLuint

// Used in `Sampler.CompareMode`.
const (
	None                CompareMode = C.GL_NONE
	CompareRefToTexture CompareMode = C.GL_COMPARE_REF_TO_TEXTURE
)

//------------------------------------------------------------------------------

// CompareFunc specifies the comparison operator.
func (sa *Sampler) CompareFunc(f CompareFunc) {
	C.SamplerCompareFunc(sa.object, C.GLenum(f))
}

// A CompareFunc specifies an operator for texture comparison.
type CompareFunc C.GLuint

// Used in `Sampler.CompareFunc`.
const (
	LessOrEqual    CompareFunc = C.GL_LEQUAL
	GreaterOrEqual CompareFunc = C.GL_GEQUAL
	Less           CompareFunc = C.GL_LESS
	Greater        CompareFunc = C.GL_GREATER
	Equal          CompareFunc = C.GL_EQUAL
	NotEqual       CompareFunc = C.GL_NOTEQUAL
	Always         CompareFunc = C.GL_ALWAYS
	Never          CompareFunc = C.GL_NEVER
)

//------------------------------------------------------------------------------

// Bind a sampler to a texture unit index.
func (sa Sampler) Bind(binding uint32) {
	C.SamplerBind(C.GLuint(binding), sa.object)
}

//------------------------------------------------------------------------------
