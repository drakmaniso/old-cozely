// Copyright (c) 2013-2016 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package gpu

//------------------------------------------------------------------------------

import (
	"github.com/drakmaniso/carol/pixel"
)

/*
#include <stdlib.h>
#include "glad.h"

static inline GLuint CreateFramebuffer(GLsizei width, GLsizei height) {
	//TODO: delte previous FBO and textures

	GLuint fbo;
	glCreateFramebuffers(1, &fbo);

	GLuint ct; // Color texture
	glCreateTextures(GL_TEXTURE_2D, 1, &ct);
	glTextureStorage2D(ct, 1, GL_RGB8, width, height);
	glTextureParameteri(ct, GL_TEXTURE_MIN_FILTER, GL_NEAREST);
	glTextureParameteri(ct, GL_TEXTURE_MAG_FILTER, GL_NEAREST);

	// GLuint dt; // Depth texture
	// glCreateTextures(GL_TEXTURE_2D, 1, &dt);
	// glTextureStorage2D(dt, 1, GL_DEPTH_COMPONENT16, width, height);
	// glTextureParameteri(dt, GL_TEXTURE_MIN_FILTER, GL_NEAREST); //TODO: remove?
	// glTextureParameteri(dt, GL_TEXTURE_MAG_FILTER, GL_NEAREST); //TODO: remove?

	glNamedFramebufferTexture(fbo, GL_COLOR_ATTACHMENT0, ct, 0);
	// glNamedFramebufferTexture(fbo, GL_DEPTH_ATTACHMENT, dt, 0);

	glNamedFramebufferDrawBuffer(fbo, GL_COLOR_ATTACHMENT0);

	glViewport(0, 0, width, height);
	glBindFramebuffer(GL_FRAMEBUFFER, fbo);

	return fbo;
}

static inline void BlitFramebuffer(GLint winWidth, GLint winHeight, GLint scrWidth, GLint scrHeight, GLint pixel, GLuint fbo) {
	glBindFramebuffer(GL_READ_FRAMEBUFFER, fbo);
	glBindFramebuffer(GL_DRAW_FRAMEBUFFER, 0);
	glClearColor(0.2,0.2,0.2,1);
	glClear(GL_COLOR_BUFFER_BIT);
	GLint w = scrWidth*pixel;
	GLint h = scrHeight*pixel;
	GLint ox = (winWidth - w) / 2;
	GLint oy = (winHeight - h) / 2;
	glBlitFramebuffer(
		0, 0, scrWidth, scrHeight, // fbo
		ox, oy, ox+w, oy+h,
		GL_COLOR_BUFFER_BIT, GL_NEAREST
	);
	//TODO: bind FBO back (but where?)
	glBindFramebuffer(GL_DRAW_FRAMEBUFFER, fbo);
	glClearColor(0,0,0,1);
	glClear(GL_COLOR_BUFFER_BIT);
}
*/
import "C"

//------------------------------------------------------------------------------

// CreateFramebuffer prepares the framebuffer.
func CreateFramebuffer(screenSize pixel.Coord, pixelSize int) {
	Framebuffer.fbo = C.CreateFramebuffer(C.GLsizei(screenSize.X), C.GLsizei(screenSize.Y))
	Framebuffer.size = screenSize
	Framebuffer.pixel = pixelSize
}

var Framebuffer struct {
	fbo   C.GLuint
	size  pixel.Coord
	pixel int
}

//------------------------------------------------------------------------------

// BlitFramebuffer blits the framebuffer onto the window backbuffer.
func BlitFramebuffer(windowSize pixel.Coord) {
	C.BlitFramebuffer(
		C.GLint(windowSize.X), C.GLint(windowSize.Y),
		C.GLint(Framebuffer.size.X), C.GLint(Framebuffer.size.Y),
		C.GLint(Framebuffer.pixel),
		C.GLuint(Framebuffer.fbo),
	)
}

//------------------------------------------------------------------------------
