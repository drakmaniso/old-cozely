// Copyright (c) 2013-2016 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package gfx

//------------------------------------------------------------------------------

import (
	"io"

	"github.com/drakmaniso/glam/internal"
)

//------------------------------------------------------------------------------

type Pipeline struct {
	program  uint32
	isClosed bool
}

//------------------------------------------------------------------------------

func NewPipeline(
	vertexShader io.Reader,
	fragmentShader io.Reader,
) (*Pipeline, error) {
	var p Pipeline
	var err error
	p.program, err = internal.CompileShaders(vertexShader, fragmentShader)
	return &p, err
}

//------------------------------------------------------------------------------

func (p *Pipeline) Close() {
	internal.CloseProgram(p.program)
	p.isClosed = true
}

//------------------------------------------------------------------------------
