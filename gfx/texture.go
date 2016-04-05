// Copyright (c) 2013-2016 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package gfx

//------------------------------------------------------------------------------

import (
	"image"
	"unsafe"

	"github.com/drakmaniso/glam/geom"
)

/*
#include "glad.h"

static inline GLuint NewTexture2D(GLsizei levels, GLenum format, GLsizei width, GLsizei height) {
	GLuint t;
	glCreateTextures(GL_TEXTURE_2D, 1, &t);
	glTextureStorage2D(t, levels, format, width, height);
	return t;
}

static inline void TextureSubImage2D(
	GLuint texture,
  	GLint level,
  	GLint xoffset,
  	GLint yoffset,
  	GLsizei width,
  	GLsizei height,
  	GLenum format,
  	GLenum type,
  	const void *pixels
) {
	glTextureSubImage2D(texture, level, xoffset, yoffset, width, height, format, type, pixels);
}

static inline void BindTextureUnit(GLuint unit, GLuint texture) {
	glBindTextureUnit(unit, texture);
}
*/
import "C"

//------------------------------------------------------------------------------

type Texture struct {
	object C.GLuint
	format textureFormat
}

//------------------------------------------------------------------------------

func NewTexture2D(levels int32, size geom.IVec2, f textureFormat) Texture {
	var t Texture
	t.format = f
	t.object = C.NewTexture2D(C.GLsizei(levels), C.GLenum(f), C.GLsizei(size.X), C.GLsizei(size.Y))
	return t
}

type textureFormat C.GLenum

const (
	RGBA8 textureFormat = C.GL_RGBA8
)

//------------------------------------------------------------------------------

func (t *Texture) Data(img image.Image, offset geom.IVec2, level int32) {
	var p unsafe.Pointer
	var pf, pt C.GLenum
	switch img := img.(type) {
	case *image.RGBA:
		p = unsafe.Pointer(&img.Pix[0])
		pf = C.GL_RGBA
		pt = C.GL_UNSIGNED_BYTE
	}
	C.TextureSubImage2D(t.object, C.GLint(level), C.GLint(offset.X), C.GLint(offset.Y), C.GLsizei(img.Bounds().Dx()), C.GLsizei(img.Bounds().Dy()), pf, pt, p)
}

//------------------------------------------------------------------------------

func (p *Pipeline) Texture(index uint32, t Texture) {
	C.BindTextureUnit(C.GLuint(index), t.object)
}

//------------------------------------------------------------------------------
