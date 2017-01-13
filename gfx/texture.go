// Copyright (c) 2013-2016 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package gfx

//------------------------------------------------------------------------------

import (
	"image"
	"unsafe"

	"github.com/drakmaniso/glam/pixel"
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

static inline void TextureGenerateMipmap(GLuint texture) {
	glGenerateTextureMipmap(texture);
}

static inline void BindTextureUnit(GLuint unit, GLuint texture) {
	glBindTextureUnit(unit, texture);
}
*/
import "C"

//------------------------------------------------------------------------------

// A Texture2D contains one or more images that all have the same format.
type Texture2D struct {
	object C.GLuint
	format textureFormat
}

// NewTexture2D returns a new 2-dimensional texture.
func NewTexture2D(levels int32, size pixel.Coord, f textureFormat) Texture2D {
	var t Texture2D
	t.format = f
	t.object = C.NewTexture2D(C.GLsizei(levels), C.GLenum(f), C.GLsizei(size.X), C.GLsizei(size.Y))
	//TODO: error handling?
	return t
}

// Data loads an image into a texture at a specific position offset and level.
func (t *Texture2D) Data(img image.Image, offset pixel.Coord, level int32) {
	p, pf, pt := pointerFormatAndTypeOf(img)
	C.TextureSubImage2D(t.object, C.GLint(level), C.GLint(offset.X), C.GLint(offset.Y), C.GLsizei(img.Bounds().Dx()), C.GLsizei(img.Bounds().Dy()), pf, pt, p)
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

type textureFormat C.GLenum

// Texture image formats.
const (
	RGBA8  textureFormat = C.GL_RGBA8
	SRGBA8 textureFormat = C.GL_SRGB8_ALPHA8
)

func pointerFormatAndTypeOf(img image.Image) (p unsafe.Pointer, pformat C.GLenum, ptype C.GLenum) {
	switch img := img.(type) {
	//TODO: other formats
	case *image.RGBA:
		p = unsafe.Pointer(&img.Pix[0])
		pformat = C.GL_RGBA
		ptype = C.GL_UNSIGNED_BYTE
	case *image.NRGBA:
		p = unsafe.Pointer(&img.Pix[0])
		pformat = C.GL_RGBA
		ptype = C.GL_UNSIGNED_BYTE
	}
	return p, pformat, ptype
}

//------------------------------------------------------------------------------
