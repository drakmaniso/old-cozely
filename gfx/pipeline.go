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
	internal     internal.Pipeline
	attribStride map[uint32]uintptr
	isCompiled   bool
	isClosed     bool
}

//------------------------------------------------------------------------------

func (p *Pipeline) CompileShaders(
	vertexShader io.Reader,
	fragmentShader io.Reader,
) error {
	var err error
	err = p.internal.CompileShaders(vertexShader, fragmentShader)
	err = p.internal.SetupVAO()
	p.attribStride = make(map[uint32]uintptr)
	p.isCompiled = true
	return err
}

//------------------------------------------------------------------------------

func (p *Pipeline) Bind() {
	p.internal.Bind()
}

//------------------------------------------------------------------------------

func (p *Pipeline) BindAttributes(binding uint32, b *Buffer) {
	p.internal.BindAttributes(binding, &b.internal, p.attribStride[binding])
}

//------------------------------------------------------------------------------

func (p *Pipeline) Close() {
	p.internal.Close()
	p.isClosed = true
}

//------------------------------------------------------------------------------
