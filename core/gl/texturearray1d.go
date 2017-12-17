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

static inline GLuint NewTextureArray1D(GLsizei levels, GLenum format, GLsizei width, GLsizei height) {
	GLuint t;
	glCreateTextures(GL_TEXTURE_1D_ARRAY, 1, &t);
	glTextureStorage2D(t, levels, format, width, height);
	return t;
}

static inline void TextureArray1DSubImage(
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

static inline void TextureGenerateMipmap(GLuint texture) {
	glGenerateTextureMipmap(texture);
}

static inline void BindTextureUnit(GLuint unit, GLuint texture) {
	glBindTextureUnit(unit, texture);
}

*/
import "C"

//------------------------------------------------------------------------------

// A TextureArray1D is an array of one-dimensional textures.
type TextureArray1D struct {
	object C.GLuint
	format TextureFormat
}

// NewTextureArray1D returns a new array of one-dimensional textures.
func NewTextureArray1D(levels int32, f TextureFormat, width int32, count int32) TextureArray1D {
	var t TextureArray1D
	t.format = f
	t.object = C.NewTextureArray1D(C.GLsizei(levels), C.GLenum(f), C.GLsizei(width), C.GLsizei(count))
	//TODO: error handling?
	return t
}

// SubImage loads an image into a texture at a specific position offset, index
// and mipmap level.
func (t *TextureArray1D) SubImage(level int32, ox int32, index int32, img image.Image) {
	p, pf, pt := pointerFormatAndTypeOf(img)
	C.TextureArray1DSubImage(t.object, C.GLint(level), C.GLint(ox), C.GLint(index), C.GLsizei(img.Bounds().Dx()), C.GLsizei(img.Bounds().Dy()), pf, pt, p)
}

// GenerateMipmap generates mipmaps for the texture.
func (t *TextureArray1D) GenerateMipmap() {
	C.TextureGenerateMipmap(t.object)
}

// Bind to a texture unit.
func (t *TextureArray1D) Bind(index uint32) {
	C.BindTextureUnit(C.GLuint(index), t.object)
}

//------------------------------------------------------------------------------
