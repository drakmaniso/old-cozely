// Copyright (c) 2013-2016 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package gfx

import (
	"fmt"
	"io"

	"github.com/drakmaniso/glam/internal"
)

//------------------------------------------------------------------------------

// A Shader is a compiled program run by the GPU.
type Shader struct {
	internal internal.Shader
}

//------------------------------------------------------------------------------

func NewVertexShader(r io.Reader) (Shader, error) {
	var s Shader
	var err error
	s.internal, err = internal.NewVertexShader(r)
	if err != nil {
		return s, fmt.Errorf("error in vertex shader: %s", err)
	}
	return s, nil
}

func NewFragmentShader(r io.Reader) (Shader, error) {
	var s Shader
	var err error
	s.internal, err = internal.NewFragmentShader(r)
	if err != nil {
		return s, fmt.Errorf("error in fragment shader: %s", err)
	}
	return s, nil
}

func NewGeometryShader(r io.Reader) (Shader, error) {
	var s Shader
	var err error
	s.internal, err = internal.NewGeometryShader(r)
	if err != nil {
		return s, fmt.Errorf("error in geometry shader: %s", err)
	}
	return s, nil
}

func NewTessControlShader(r io.Reader) (Shader, error) {
	var s Shader
	var err error
	s.internal, err = internal.NewTessControlShader(r)
	if err != nil {
		return s, fmt.Errorf("error in tesselation control shader: %s", err)
	}
	return s, nil
}

func NewTessEvaluationShader(r io.Reader) (Shader, error) {
	var s Shader
	var err error
	s.internal, err = internal.NewTessEvaluationShader(r)
	if err != nil {
		return s, fmt.Errorf("error in tesselation evaluation shader: %s", err)
	}
	return s, nil
}

func NewComputeShader(r io.Reader) (Shader, error) {
	var s Shader
	var err error
	s.internal, err = internal.NewComputeShader(r)
	if err != nil {
		return s, fmt.Errorf("error in compute shader: %s", err)
	}
	return s, nil
}

//------------------------------------------------------------------------------
