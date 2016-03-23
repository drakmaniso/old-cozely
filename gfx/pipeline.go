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
	internal internal.Pipeline
	isClosed bool
}

//------------------------------------------------------------------------------

func NewPipeline(
	vertexShader io.Reader,
	fragmentShader io.Reader,
) (*Pipeline, error) {
	var p Pipeline
	var err error
	err = p.internal.CompileShaders(vertexShader, fragmentShader)
	err = p.internal.SetupVAO()
	return &p, err
}

//------------------------------------------------------------------------------

func (p *Pipeline) Bind() {
	p.internal.Bind()
}

//------------------------------------------------------------------------------

func (p *Pipeline) Close() {
	p.internal.Close()
	p.isClosed = true
}

//------------------------------------------------------------------------------
