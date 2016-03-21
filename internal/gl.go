// Copyright (c) 2013-2016 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package internal

//------------------------------------------------------------------------------

/*
#include "glad.h"

GLenum GetOpenGLError() {
	return glGetError();
}
*/
import "C"
import "errors"

//------------------------------------------------------------------------------

type GLuint C.GLuint

type GLenum C.GLenum

//------------------------------------------------------------------------------

// CheckGLError checks for any OpenGL error.
func CheckGLError() error {
	switch C.GetOpenGLError() {
	case C.GL_NO_ERROR:
		return nil
	case C.GL_INVALID_ENUM:
		return ErrGLInvalidEnum
	case C.GL_INVALID_VALUE:
		return ErrGLInvalidValue
	case C.GL_INVALID_OPERATION:
		return ErrGLInvalidOperation
	case C.GL_INVALID_FRAMEBUFFER_OPERATION:
		return ErrGLInvalidFramebufferOperation
	case C.GL_OUT_OF_MEMORY:
		return ErrGLOutOfMemory
	}
	return errors.New("Unkown OpenGL Error")
}

// OpenGL errors.
var (
	ErrGLInvalidEnum                 = errors.New("Invalid OpenGL enum")
	ErrGLInvalidValue                = errors.New("Invalid OpenGL value")
	ErrGLInvalidOperation            = errors.New("Invalid OpenGL operation")
	ErrGLInvalidFramebufferOperation = errors.New("Invalid OpenGL framebuffer operation")
	ErrGLOutOfMemory                 = errors.New("Out of memory for OpenGL")
)

//------------------------------------------------------------------------------
