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
) GLuint {
	vs, err := compileShader(vertexShader, C.GL_VERTEX_SHADER)
	if err != nil {
		log.Print("compile error in vertex shader:", err)
	}
	fs, err := compileShader(fragmentShader, C.GL_FRAGMENT_SHADER)
	if err != nil {
		log.Print("compile error in fragment shader:", err)
	}

	p := C.LinkProgram(vs, fs)
	return GLuint(p)
}

func compileShader(r io.Reader, t C.GLenum) (C.GLuint, error) {
	b, err := ioutil.ReadAll(r)
	if err != nil {

	}
	cb := C.CString(string(b))
	defer C.free(unsafe.Pointer(cb))
	s := C.CompileShader((*C.GLchar)(unsafe.Pointer(cb)), t)

	errmsg := C.CheckCompileShaderError(s)
	if errmsg != nil {
		defer C.free(unsafe.Pointer(errmsg))
		return 0, errors.New(C.GoString(errmsg))
	}

	return s, nil
}

//------------------------------------------------------------------------------
