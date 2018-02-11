// Copyright (c) 2013-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package gl

import (
	"image"
)

//------------------------------------------------------------------------------

/*
#include "glad.h"

static inline GLuint NewTextureArray2D(GLsizei levels, GLenum format, GLsizei width, GLsizei height, GLsizei depth) {
	GLuint t;
	glCreateTextures(GL_TEXTURE_2D_ARRAY, 1, &t);
	glTextureStorage3D(t, levels, format, width, height, depth);
	return t;
}

static inline void TextureArray2DSubImage(
	GLuint texture,
  	GLint level,
  	GLint xoffset,
  	GLint yoffset,
  	GLint zoffset,
  	GLsizei width,
  	GLsizei height,
  	GLsizei depth,
  	GLenum format,
  	GLenum type,
  	const void *pixels
) {
	glTextureSubImage3D(texture, level, xoffset, yoffset, zoffset, width, height, depth, format, type, pixels);
}

static inline void TextureGenerateMipmap(GLuint texture) {
	glGenerateTextureMipmap(texture);
}

static inline void BindTextureUnit(GLuint unit, GLuint texture) {
	glBindTextureUnit(unit, texture);
}

static inline void DeleteTexture(GLuint texture) {
	glDeleteTextures(1, &texture);
}

*/
import "C"

//------------------------------------------------------------------------------

// A TextureArray2D is an array of two-dimensional textures.
type TextureArray2D struct {
	object C.GLuint
	format TextureFormat
}

// NewTextureArray2D returns a new array of two-dimensional textures.
func NewTextureArray2D(levels int32, f TextureFormat, width, height int32, count int32) TextureArray2D {
	var t TextureArray2D
	t.format = f
	t.object = C.NewTextureArray2D(C.GLsizei(levels), C.GLenum(f), C.GLsizei(width), C.GLsizei(height), C.GLsizei(count))
	//TODO: error handling?
	return t
}

// SubImage loads an image into a texture at a specific position offset, array
// index and mipmap level.
func (t *TextureArray2D) SubImage(level int32, ox, oy int32, index int32, img image.Image) {
	p, pf, pt := pointerFormatAndTypeOf(img)
	C.TextureArray2DSubImage(
		t.object,
		C.GLint(level),
		C.GLint(ox), C.GLint(oy), C.GLint(index),
		C.GLsizei(img.Bounds().Dx()), C.GLsizei(img.Bounds().Dy()), C.GLsizei(1),
		pf, pt, p,
	)
}

// GenerateMipmap generates mipmaps for the texture.
func (t *TextureArray2D) GenerateMipmap() {
	C.TextureGenerateMipmap(t.object)
}

// Bind to a texture unit.
func (t *TextureArray2D) Bind(index uint32) {
	C.BindTextureUnit(C.GLuint(index), t.object)
}

// Delete frees the texture
func (t *TextureArray2D) Delete() {
	C.DeleteTexture(t.object)
}

//------------------------------------------------------------------------------
