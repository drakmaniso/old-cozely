// Copyright (c) 2013-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package gl

import (
	"image"
)

//------------------------------------------------------------------------------

/*
#include "glad.h"

static inline GLuint NewTexture3D(GLsizei levels, GLenum format, GLsizei width, GLsizei height, GLsizei depth) {
	GLuint t;
	glCreateTextures(GL_TEXTURE_3D, 1, &t);
	glTextureStorage3D(t, levels, format, width, height, depth);
	return t;
}

static inline void Texture3DSubImage(
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

// A Texture3D is a three-dimensional texture.
type Texture3D struct {
	object C.GLuint
	format TextureFormat
}

// NewTexture3D returns a new three-dimensional texture.
func NewTexture3D(levels int32, f TextureFormat, width, height, depth int32) Texture3D {
	var t Texture3D
	t.format = f
	t.object = C.NewTexture3D(C.GLsizei(levels), C.GLenum(f), C.GLsizei(width), C.GLsizei(height), C.GLsizei(depth))
	//TODO: error handling?
	return t
}

// SubImage loads an image into a texture at a specific position offset and
// mipmap level.
func (t *Texture3D) SubImage(level int32, ox, oy, oz int32, img image.Image) {
	p, pf, pt := pointerFormatAndTypeOf(img)
	C.Texture3DSubImage(
		t.object,
		C.GLint(level),
		C.GLint(ox), C.GLint(oy), C.GLint(oz),
		C.GLsizei(img.Bounds().Dx()), C.GLsizei(img.Bounds().Dy()), C.GLsizei(1),
		pf, pt, p,
	)
}

// GenerateMipmap generates mipmaps for the texture.
func (t *Texture3D) GenerateMipmap() {
	C.TextureGenerateMipmap(t.object)
}

// Bind to a texture unit.
func (t *Texture3D) Bind(index uint32) {
	C.BindTextureUnit(C.GLuint(index), t.object)
}

// Delete frees the texture
func (t *Texture3D) Delete() {
	C.DeleteTexture(t.object)
}

//------------------------------------------------------------------------------
