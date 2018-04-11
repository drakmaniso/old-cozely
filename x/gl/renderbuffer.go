// Copyright (c) 2018-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package gl

////////////////////////////////////////////////////////////////////////////////

/*
#include "glad.h"

static inline GLuint NewRenderbuffer(GLenum format, GLsizei width, GLsizei height) {
	GLuint r;
	glCreateRenderbuffers(1, &r);
	glNamedRenderbufferStorage(r, format, width, height);
	return r;
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

static inline void DeleteRenderbuffer(GLuint r) {
	glDeleteRenderbuffers(1, &r);
}

*/
import "C"

////////////////////////////////////////////////////////////////////////////////

// A Renderbuffer is a two-dimensional texture that can only be used for
// rendering (attached a Framebuffer)
type Renderbuffer struct {
	object C.GLuint
	format TextureFormat
}

// NewRenderbuffer returns a new render buffer.
func NewRenderbuffer(f TextureFormat, width, height int32) Renderbuffer {
	var r Renderbuffer
	r.format = f
	r.object = C.NewRenderbuffer(C.GLenum(f), C.GLsizei(width), C.GLsizei(height))
	//TODO: error handling?
	return r
}

// Delete frees the render buffer
func (r *Renderbuffer) Delete() {
	C.DeleteRenderbuffer(r.object)
}

////////////////////////////////////////////////////////////////////////////////
