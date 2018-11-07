// Copyright (c) 2013-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package gl

import (
	"image"
	"unsafe"
)

////////////////////////////////////////////////////////////////////////////////

/*
#include "glad.h"

static inline GLuint NewTexture2D(GLsizei levels, GLenum format, GLsizei width, GLsizei height) {
	GLuint t;
	glCreateTextures(GL_TEXTURE_2D, 1, &t);
	glTextureStorage2D(t, levels, format, width, height);
	return t;
}

static inline void Texture2DSubImage(
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

static inline void BindImageUnit(GLuint unit, GLuint texture, GLint level, GLboolean layered, GLint layer, GLenum access, GLenum format) {
	glBindImageTexture(unit, texture, level, layered, layer, access, format);
}

static inline void ClearTextureImage(GLuint texture, GLint level, GLenum format, GLenum type, const void *data) {
	glClearTexImage(texture, level, format, type, data);
}

static inline void DeleteTexture(GLuint texture) {
	glDeleteTextures(1, &texture);
}

*/
import "C"

////////////////////////////////////////////////////////////////////////////////

// A Texture2D is a two-dimensional texture.
type Texture2D struct {
	object C.GLuint
	format TextureFormat
}

// NewTexture2D returns a new two-dimensional texture.
func NewTexture2D(levels int32, f TextureFormat, width, height int32) Texture2D {
	var t Texture2D
	t.format = f
	t.object = C.NewTexture2D(C.GLsizei(levels), C.GLenum(f), C.GLsizei(width), C.GLsizei(height))
	//TODO: error handling?
	return t
}

// SubImage loads an image into a texture at a specific position offset and
// mipmap level.
func (t *Texture2D) SubImage(level int32, ox, oy int32, img image.Image) {
	p, pf, pt := pointerFormatAndTypeOf(img)
	C.Texture2DSubImage(
		t.object,
		C.GLint(level),
		C.GLint(ox), C.GLint(oy),
		C.GLsizei(img.Bounds().Dx()), C.GLsizei(img.Bounds().Dy()),
		pf, pt, p,
	)
}

// SubData loads data into a texture at a specific position offset, array
// index and mipmap level.
func (t *Texture2D) SubData(level int32, ox, oy int32, sx, sy int32, data interface{}) {
	p, pf, pt := pointerFormatAndTypeOfData(data)
	C.Texture2DSubImage(
		t.object,
		C.GLint(level),
		C.GLint(ox), C.GLint(oy),
		C.GLsizei(sx), C.GLsizei(sy),
		pf, pt, p,
	)
}

// GenerateMipmap generates mipmaps for the texture.
func (t *Texture2D) GenerateMipmap() {
	C.TextureGenerateMipmap(t.object)
}

// Bind to a texture unit.
func (t *Texture2D) Bind(index uint32) {
	C.BindTextureUnit(C.GLuint(index), t.object)
}

// BindImage to an image unit.
func (t *Texture2D) BindImage(index uint32, level int32, a Access, f TextureFormat) {
	C.BindImageUnit(C.GLuint(index), t.object, C.GLint(level), C.GL_FALSE, C.GLint(0), C.GLenum(a), C.GLenum(f))
}

func (t *Texture2D) ClearByte(level int32, r uint8) {
	var c struct{ R uint8 }
	c.R = r
	C.ClearTextureImage(t.object, C.GLint(level), C.GL_RED_INTEGER, C.GL_UNSIGNED_BYTE, unsafe.Pointer(&c))
}

// Delete frees the texture
func (t *Texture2D) Delete() {
	C.DeleteTexture(t.object)
}

////////////////////////////////////////////////////////////////////////////////
