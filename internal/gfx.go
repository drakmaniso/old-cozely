// Copyright (c) 2013-2016 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package internal

//------------------------------------------------------------------------------

//#include "gfx.h"
import "C"

import (
	"errors"
	"io"
	"io/ioutil"
	"log"
	"unsafe"
)

//------------------------------------------------------------------------------

func CompileShaders(
	vertexShader io.Reader,
	fragmentShader io.Reader,
) (GLuint, error) {
	vs, err := compileShader(vertexShader, C.GL_VERTEX_SHADER)
	if err != nil {
		log.Print("vertex shader compile error: ", err)
	}
	fs, err := compileShader(fragmentShader, C.GL_FRAGMENT_SHADER)
	if err != nil {
		log.Print("fragment shader compile error: ", err)
	}

	p := C.LinkProgram(vs, fs)
	if errm := C.LinkProgramError(p); errm != nil {
		defer C.free(unsafe.Pointer(errm))
		err = errors.New(C.GoString(errm))
		log.Print("shader link error: ", err)
		return GLuint(p), err
	}
	return GLuint(p), err
}

func compileShader(r io.Reader, t C.GLenum) (C.GLuint, error) {
	b, err := ioutil.ReadAll(r)
	if err != nil {
		log.Print("failed to read shader: ", err)
		return 0, err
	}
	cb := C.CString(string(b))
	defer C.free(unsafe.Pointer(cb))

	s := C.CompileShader((*C.GLchar)(unsafe.Pointer(cb)), t)
	if errm := C.CompileShaderError(s); errm != nil {
		defer C.free(unsafe.Pointer(errm))
		return s, errors.New(C.GoString(errm))
	}

	return s, nil
}

//------------------------------------------------------------------------------

func CloseProgram(p GLuint) {
	C.CloseProgram((C.GLuint)(p))
}

//------------------------------------------------------------------------------
