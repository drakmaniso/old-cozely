// Copyright (c) 2013-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package gl

//------------------------------------------------------------------------------

/*
#include "glad.h"

static inline GLuint NewFramebuffer() {
	GLuint fbo;
	glCreateFramebuffers(1, &fbo);
	return fbo;
}

static inline void FramebufferTexture(GLuint fbo, GLenum a, GLuint t, GLint l) {
	glNamedFramebufferTexture(fbo, a, t, l);
}

static inline void FramebufferRenderBuffer(GLuint fbo, GLenum a, GLuint t) {
	glNamedFramebufferRenderbuffer(fbo, a, GL_RENDERBUFFER, t);
}

static inline void FramebufferDrawBuffer(GLuint fbo, GLenum a) {
	glNamedFramebufferDrawBuffer(fbo, a);
}

static inline void FramebufferBind(GLuint fbo, GLenum t) {
	glBindFramebuffer(t, fbo);
}

static inline void FramebufferClearColorUint(GLuint fbo, const GLuint *v) {
	glClearNamedFramebufferuiv(fbo, GL_COLOR, GL_NONE, v);
}

static inline void FramebufferBlit(GLuint fbo, GLuint dstFbo, GLint srcX1, GLint srcY1, GLint srcX2, GLint srcY2, GLint dstX1, GLint dstY1, GLint dstX2, GLint dstY2, GLbitfield m, GLenum f) {
	glBlitNamedFramebuffer(fbo, dstFbo, srcX1, srcY1, srcX2, srcY2, dstX1, dstY1, dstX2, dstY2, m, f);
}

*/
import "C"
import "unsafe"

//------------------------------------------------------------------------------

type Framebuffer struct {
	object C.GLuint
}

var DefaultFramebuffer = Framebuffer{
	object: C.GLuint(0),
}

//------------------------------------------------------------------------------

func NewFramebuffer() Framebuffer {
	var f Framebuffer
	f.object = C.NewFramebuffer()
	return f
}

//------------------------------------------------------------------------------

func (fb Framebuffer) Texture(a FramebufferAttachment, t Texture2D, level int32) {
	C.FramebufferTexture(fb.object, C.GLenum(a), t.object, C.GLint(level))
}

func (fb Framebuffer) RenderBuffer(a FramebufferAttachment, r RenderBuffer) {
	C.FramebufferRenderBuffer(fb.object, C.GLenum(a), r.object)
}

type FramebufferAttachment C.GLenum

const (
	ColorAttachment0       FramebufferAttachment = C.GL_COLOR_ATTACHMENT0
	ColorAttachment1       FramebufferAttachment = C.GL_COLOR_ATTACHMENT1
	ColorAttachment2       FramebufferAttachment = C.GL_COLOR_ATTACHMENT2
	ColorAttachment3       FramebufferAttachment = C.GL_COLOR_ATTACHMENT3
	DepthAttachment        FramebufferAttachment = C.GL_DEPTH_ATTACHMENT
	StencilAttachment      FramebufferAttachment = C.GL_STENCIL_ATTACHMENT
	DepthStencilAttachment FramebufferAttachment = C.GL_DEPTH_STENCIL_ATTACHMENT
)

//------------------------------------------------------------------------------

func (fb Framebuffer) DrawBuffer(a FramebufferAttachment) {
	C.FramebufferDrawBuffer(fb.object, C.GLenum(a))
}

//------------------------------------------------------------------------------

func (fb Framebuffer) Bind(t FramebufferTarget) {
	C.FramebufferBind(fb.object, C.GLenum(t))
}

type FramebufferTarget C.GLenum

const (
	DrawFramebuffer     FramebufferTarget = C.GL_DRAW_FRAMEBUFFER
	ReadFramebuffer     FramebufferTarget = C.GL_READ_FRAMEBUFFER
	DrawReadFramebuffer FramebufferTarget = C.GL_FRAMEBUFFER
)

//------------------------------------------------------------------------------

func (fb Framebuffer) ClearColorUint(r, g, b, a uint32) {
	//TODO: other variants
	var c struct{ R, G, B, A uint32 }
	c.R = r
	c.G = g
	c.B = b
	c.A = a
	C.FramebufferClearColorUint(fb.object, (*C.GLuint)(unsafe.Pointer(&c)))
}

//------------------------------------------------------------------------------

func (fb Framebuffer) Blit(dst Framebuffer, srcX1, srcY1, srcX2, srcY2, dstX1, dstY1, dstX2, dstY2 int32, m BufferMask, f FilterMode) {
	C.FramebufferBlit(
		fb.object,
		dst.object,
		C.GLint(srcX1), C.GLint(srcY1), C.GLint(srcX2), C.GLint(srcY2),
		C.GLint(dstX1), C.GLint(dstY1), C.GLint(dstX2), C.GLint(dstY2),
		C.GLbitfield(m),
		C.GLenum(f),
	)
}

type BufferMask C.GLbitfield

const (
	ColorBufferBit   BufferMask = C.GL_COLOR_BUFFER_BIT
	DepthBufferBit   BufferMask = C.GL_DEPTH_BUFFER_BIT
	StencilBufferBit BufferMask = C.GL_STENCIL_BUFFER_BIT
)

//------------------------------------------------------------------------------
