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

static inline void PipelineSampler(GLuint binding, GLuint sampler) {
	glBindSampler(binding, sampler);
}

*/
import "C"

//------------------------------------------------------------------------------

// A Sampler stores the sampling parameters for a texture inside a shader.
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
func (sa *Sampler) Filter(min, mag filter) {
	C.SamplerFilter(sa.object, C.GLenum(min), C.GLenum(mag))
}

type filter C.GLuint

// Sampler filtering functions.
const (
	Nearest             filter = C.GL_NEAREST
	Linear              filter = C.GL_LINEAR
	NearestMimapNearest filter = C.GL_NEAREST_MIPMAP_NEAREST
	LinearMipmapNearest filter = C.GL_LINEAR_MIPMAP_NEAREST
	NearestMipmapLinear filter = C.GL_NEAREST_MIPMAP_LINEAR
	LinearMipmapLinear  filter = C.GL_LINEAR_MIPMAP_LINEAR
)

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
func (sa *Sampler) Wrap(s, t, p wrap) {
	C.SamplerWrap(sa.object, C.GLenum(s), C.GLenum(t), C.GLenum(p))
}

type wrap C.GLuint

// Sampler wrapping modes.
const (
	ClampToBorder     wrap = C.GL_CLAMP_TO_BORDER
	ClampToEdge       wrap = C.GL_CLAMP_TO_EDGE
	MirrorClampToEdge wrap = C.GL_MIRROR_CLAMP_TO_EDGE
	Repeat            wrap = C.GL_REPEAT
	MirroredRepeat    wrap = C.GL_MIRRORED_REPEAT
)

// BorderColor sets the color used for texture filtering when ClampToborder
// wrapping mode is used.
func (sa *Sampler) BorderColor(c color.RGBA) {
	C.SamplerBorderColor(sa.object, (*C.GLfloat)(unsafe.Pointer(&c)))
}

//------------------------------------------------------------------------------

// CompareMode specifies the texture comparison mode.
func (sa *Sampler) CompareMode(m compareMode) {
	C.SamplerCompareMode(sa.object, C.GLenum(m))
}

type compareMode C.GLuint

// Sampler comparison modes.
const (
	None                compareMode = C.GL_NONE
	CompareRefToTexture compareMode = C.GL_COMPARE_REF_TO_TEXTURE
)

// CompareFunc specifies the comparison operator.
func (sa *Sampler) CompareFunc(f compareFunc) {
	C.SamplerCompareFunc(sa.object, C.GLenum(f))
}

type compareFunc C.GLuint

// Sampelr comparison operators.
const (
	LessOrEqual    compareFunc = C.GL_LEQUAL
	GreaterOrEqual compareFunc = C.GL_GEQUAL
	Less           compareFunc = C.GL_LESS
	Greater        compareFunc = C.GL_GREATER
	Equal          compareFunc = C.GL_EQUAL
	NotEqual       compareFunc = C.GL_NOTEQUAL
	Always         compareFunc = C.GL_ALWAYS
	Never          compareFunc = C.GL_NEVER
)

//------------------------------------------------------------------------------

// Sampler binds a sampler to a texture unit index.
func (p *Pipeline) Sampler(binding uint32, sa Sampler) {
	C.PipelineSampler(C.GLuint(binding), sa.object)
}

//------------------------------------------------------------------------------
