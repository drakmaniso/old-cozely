// Copyright (c) 2013-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package gl

import (
	"strconv"
	"unsafe"
)

////////////////////////////////////////////////////////////////////////////////

/*
#include "glad.h"

static inline GLuint NewFramebuffer() {
	GLuint fbo;
	glCreateFramebuffers(1, &fbo);
	return fbo;
}

static inline void FramebufferSetEmptySize(GLuint fbo, GLint width, GLint height) {
	glNamedFramebufferParameteri(fbo, GL_FRAMEBUFFER_DEFAULT_WIDTH, width);
	glNamedFramebufferParameteri(fbo, GL_FRAMEBUFFER_DEFAULT_HEIGHT, height);
	glNamedFramebufferParameteri(fbo, GL_FRAMEBUFFER_DEFAULT_LAYERS, 0);
	glNamedFramebufferParameteri(fbo, GL_FRAMEBUFFER_DEFAULT_SAMPLES, 0);
	glNamedFramebufferParameteri(fbo, GL_FRAMEBUFFER_DEFAULT_FIXED_SAMPLE_LOCATIONS, 0);
}

static inline void FramebufferTexture(GLuint fbo, GLenum a, GLuint t, GLint l) {
	glNamedFramebufferTexture(fbo, a, t, l);
}

static inline void FramebufferRenderbuffer(GLuint fbo, GLenum a, GLuint t) {
	glNamedFramebufferRenderbuffer(fbo, a, GL_RENDERBUFFER, t);
}

static inline void FramebufferReadBuffer(GLuint fbo, GLenum a) {
	glNamedFramebufferReadBuffer(fbo, a);
}

static inline void FramebufferDrawBuffer(GLuint fbo, GLenum a) {
	glNamedFramebufferDrawBuffer(fbo, a);
}

static inline void FramebufferDrawBuffers(GLuint fbo, GLsizei n, GLenum *a) {
	glNamedFramebufferDrawBuffers(fbo, n, a);
}

static inline GLenum FramebufferCheckStatus(GLuint fbo, GLenum t) {
	return glCheckNamedFramebufferStatus(fbo, t);
}

static inline void FramebufferBind(GLuint fbo, GLenum t) {
	glBindFramebuffer(t, fbo);
}

static inline void FramebufferClearColorUint(GLuint fbo, const GLuint *v) {
	glClearNamedFramebufferuiv(fbo, GL_COLOR, 0, v);
}

static inline void FramebufferClearDepth(GLuint fbo, const GLfloat v) {
	glClearNamedFramebufferfv(fbo, GL_DEPTH, 0, &v);
}

static inline void FramebufferBlit(GLuint fbo, GLuint dstFbo, GLint srcX1, GLint srcY1, GLint srcX2, GLint srcY2, GLint dstX1, GLint dstY1, GLint dstX2, GLint dstY2, GLbitfield m, GLenum f) {
	glBlitNamedFramebuffer(fbo, dstFbo, srcX1, srcY1, srcX2, srcY2, dstX1, dstY1, dstX2, dstY2, m, f);
}

static inline void FramebufferDelete(GLuint fbo) {
	glDeleteFramebuffers(1, &fbo);
}

*/
import (
	"C"
)

////////////////////////////////////////////////////////////////////////////////

type Framebuffer struct {
	object C.GLuint
}

var DefaultFramebuffer = Framebuffer{
	object: C.GLuint(0),
}

////////////////////////////////////////////////////////////////////////////////

func NewFramebuffer() Framebuffer {
	var f Framebuffer
	f.object = C.NewFramebuffer()
	return f
}

func (fb Framebuffer) SetEmptySize(w, h int16) {
	C.FramebufferSetEmptySize(fb.object, C.GLint(w), C.GLint(h))
}

////////////////////////////////////////////////////////////////////////////////

func (fb Framebuffer) Texture(a FramebufferAttachment, t Texture2D, level int32) {
	C.FramebufferTexture(fb.object, C.GLenum(a), t.object, C.GLint(level))
}

func (fb Framebuffer) Renderbuffer(a FramebufferAttachment, r Renderbuffer) {
	C.FramebufferRenderbuffer(fb.object, C.GLenum(a), r.object)
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
	NoAttachment           FramebufferAttachment = C.GL_NONE
)

////////////////////////////////////////////////////////////////////////////////

func (fb Framebuffer) ReadBuffer(a FramebufferAttachment) {
	C.FramebufferReadBuffer(fb.object, C.GLenum(a))
}

func (fb Framebuffer) DrawBuffer(a FramebufferAttachment) {
	C.FramebufferDrawBuffer(fb.object, C.GLenum(a))
}

func (fb Framebuffer) DrawBuffers(a []FramebufferAttachment) {
	C.FramebufferDrawBuffers(fb.object, C.GLsizei(len(a)), (*C.GLenum)(&a[0]))
}

////////////////////////////////////////////////////////////////////////////////

func (fb Framebuffer) Bind(t FramebufferTarget) {
	C.FramebufferBind(fb.object, C.GLenum(t))
}

type FramebufferTarget C.GLenum

const (
	DrawFramebuffer     FramebufferTarget = C.GL_DRAW_FRAMEBUFFER
	ReadFramebuffer     FramebufferTarget = C.GL_READ_FRAMEBUFFER
	DrawReadFramebuffer FramebufferTarget = C.GL_FRAMEBUFFER
)

////////////////////////////////////////////////////////////////////////////////

func (fb Framebuffer) CheckStatus(t FramebufferTarget) FramebufferStatus {
	e := C.FramebufferCheckStatus(fb.object, C.GLenum(t))
	return FramebufferStatus(e)
}

type FramebufferStatus C.GLenum

const (
	FramebufferComplete                    FramebufferStatus = C.GL_FRAMEBUFFER_COMPLETE
	FramebufferUndefined                   FramebufferStatus = C.GL_FRAMEBUFFER_UNDEFINED
	FramebufferIncompleteAttachment        FramebufferStatus = C.GL_FRAMEBUFFER_INCOMPLETE_ATTACHMENT
	FramebufferIncompleteMissingAttachment FramebufferStatus = C.GL_FRAMEBUFFER_INCOMPLETE_MISSING_ATTACHMENT
	FramebufferIncompleteDrawBuffer        FramebufferStatus = C.GL_FRAMEBUFFER_INCOMPLETE_DRAW_BUFFER
	FramebufferIncompleteReadBuffer        FramebufferStatus = C.GL_FRAMEBUFFER_INCOMPLETE_READ_BUFFER
	FramebufferUnsupported                 FramebufferStatus = C.GL_FRAMEBUFFER_UNSUPPORTED
	FramebufferIncompleteMultisample       FramebufferStatus = C.GL_FRAMEBUFFER_INCOMPLETE_MULTISAMPLE
	FramebufferIncompleteLayerTargets      FramebufferStatus = C.GL_FRAMEBUFFER_INCOMPLETE_LAYER_TARGETS
)

func (fbs FramebufferStatus) String() string {
	switch fbs {
	case FramebufferStatus(0):
		return "framebuffer check status error"
	case FramebufferComplete:
		return "framebuffer complete"
	case FramebufferUndefined:
		return "framebuffer undefined"
	case FramebufferIncompleteAttachment:
		return "framebuffer incomplete attachment"
	case FramebufferIncompleteMissingAttachment:
		return "framebuffer incomplete missing attachment"
	case FramebufferIncompleteDrawBuffer:
		return "framebuffer incomplete draw buffer"
	case FramebufferIncompleteReadBuffer:
		return "framebuffer incomplete read buffer"
	case FramebufferUnsupported:
		return "framebuffer unsupported"
	case FramebufferIncompleteMultisample:
		return "framebuffer incomplete multisample"
	case FramebufferIncompleteLayerTargets:
		return "framebuffer incomplete layer targets"
	}
	return "(unknown framebuffer status: " + strconv.Itoa(int(fbs)) + ")"
}

////////////////////////////////////////////////////////////////////////////////

func (fb Framebuffer) ClearColorUint(r, g, b, a uint32) {
	//TODO: other variants
	var c struct{ R, G, B, A uint32 }
	c.R = r
	c.G = g
	c.B = b
	c.A = a
	C.FramebufferClearColorUint(fb.object, (*C.GLuint)(unsafe.Pointer(&c)))
}

func (fb Framebuffer) ClearDepth(d float32) {
	C.FramebufferClearDepth(fb.object, (C.GLfloat)(d))
}

////////////////////////////////////////////////////////////////////////////////

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

////////////////////////////////////////////////////////////////////////////////

func (fb Framebuffer) Delete() {
	C.FramebufferDelete(fb.object)
}

////////////////////////////////////////////////////////////////////////////////
