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

static inline void TextureGenerateMipmap(GLuint texture) {
	glGenerateTextureMipmap(texture);
}

static inline void BindTextureUnit(GLuint unit, GLuint texture) {
	glBindTextureUnit(unit, texture);
}

*/
import "C"

//------------------------------------------------------------------------------

// A Texture2D is a two-dimensional texture.
type Texture2D struct {
	object C.GLuint
	format TextureFormat
}

// NewTexture2D returns a new two-dimensional texture.
func NewTexture2D(levels int32, width int32, height int32, f TextureFormat) Texture2D {
	var t Texture2D
	t.format = f
	t.object = C.NewTexture2D(C.GLsizei(levels), C.GLenum(f), C.GLsizei(width), C.GLsizei(height))
	//TODO: error handling?
	return t
}

// Load loads an image into a texture at a specific position offset and level.
func (t *Texture2D) Load(img image.Image, ox int32, oy int32, level int32) {
	p, pf, pt := pointerFormatAndTypeOf(img)
	C.TextureSubImage2D(t.object, C.GLint(level), C.GLint(ox), C.GLint(oy), C.GLsizei(img.Bounds().Dx()), C.GLsizei(img.Bounds().Dy()), pf, pt, p)
}

// GenerateMipmap generates mipmaps for the texture.
func (t *Texture2D) GenerateMipmap() {
	C.TextureGenerateMipmap(t.object)
}

// Bind to a texture unit.
func (t *Texture2D) Bind(index uint32) {
	C.BindTextureUnit(C.GLuint(index), t.object)
}

//------------------------------------------------------------------------------
