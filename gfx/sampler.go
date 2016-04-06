// Copyright (c) 2013-2016 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package gfx

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

static inline void PipelineSampler(GLuint binding, GLuint sampler) {
	glBindSampler(binding, sampler);
}

*/
import "C"

//------------------------------------------------------------------------------

type Sampler struct {
	object C.GLuint
}

//------------------------------------------------------------------------------

func NewSampler() Sampler {
	var s Sampler
	s.object = C.NewSampler()
	return s
}

//------------------------------------------------------------------------------

func (s *Sampler) Filter(min, mag filter) {
	C.SamplerFilter(s.object, C.GLenum(min), C.GLenum(mag))
}

type filter C.GLuint

const (
	Nearest             filter = C.GL_NEAREST
	Linear              filter = C.GL_LINEAR
	NearestMimapNearest filter = C.GL_NEAREST_MIPMAP_NEAREST
	LinearMipmapNearest filter = C.GL_LINEAR_MIPMAP_NEAREST
	NearestMipmapLinear filter = C.GL_NEAREST_MIPMAP_LINEAR
	LinearMipmapLinear  filter = C.GL_LINEAR_MIPMAP_LINEAR
)

//------------------------------------------------------------------------------

func (s *Sampler) LevelOfDetail(min, max float32) {
	C.SamplerLevelOfDetail(s.object, C.GLfloat(min), C.GLfloat(max))
}

//------------------------------------------------------------------------------

func (s *Sampler) Anisotropy(max float32) {
	C.SamplerAnisotropy(s.object, C.GLfloat(max))
}

//------------------------------------------------------------------------------

func (p *Pipeline) Sampler(binding uint32, s Sampler) {
	C.PipelineSampler(C.GLuint(binding), s.object)
}

//------------------------------------------------------------------------------
