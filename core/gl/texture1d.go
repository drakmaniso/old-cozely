// Copyright (c) 2013-2016 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package gl

//------------------------------------------------------------------------------

import (
	"image"
)

//------------------------------------------------------------------------------

/*
#include "glad.h"

static inline GLuint NewTexture1D(GLsizei levels, GLenum format, GLsizei width) {
	GLuint t;
	glCreateTextures(GL_TEXTURE_1D, 1, &t);
	glTextureStorage1D(t, levels, format, width);
	return t;
}

static inline void TextureSubImage1D(
	GLuint texture,
  	GLint level,
  	GLint xoffset,
  	GLsizei width,
  	GLenum format,
  	GLenum type,
  	const void *pixels
) {
	glTextureSubImage1D(texture, level, xoffset, width, format, type, pixels);
}

static inline void TextureGenerateMipmap(GLuint texture) {
	glGenerateTextureMipmap(texture);
}

static inline void BindTextureUnit(GLuint unit, GLuint texture) {
	glBindTextureUnit(unit, texture);
}

*/
import "C"

//------------------------------------------------------------------------------

// A Texture1D is a one-dimensional texture.
type Texture1D struct {
	object C.GLuint
	format TextureFormat
}

// NewTexture1D returns a new one-dimensional texture.
func NewTexture1D(levels int32, width int32, f TextureFormat) Texture1D {
	var t Texture1D
	t.format = f
	t.object = C.NewTexture1D(C.GLsizei(levels), C.GLenum(f), C.GLsizei(width))
	//TODO: error handling?
	return t
}

// Load loads an image into a texture at a specific position offset and level.
func (t *Texture1D) Load(img image.Image, ox int32, level int32) {
	p, pf, pt := pointerFormatAndTypeOf(img)
	C.TextureSubImage1D(t.object, C.GLint(level), C.GLint(ox), C.GLsizei(img.Bounds().Dx()), pf, pt, p)
}

// GenerateMipmap generates mipmaps for the texture.
func (t *Texture1D) GenerateMipmap() {
	C.TextureGenerateMipmap(t.object)
}

// Bind to a texture unit.
func (t *Texture1D) Bind(index uint32) {
	C.BindTextureUnit(C.GLuint(index), t.object)
}

//------------------------------------------------------------------------------
