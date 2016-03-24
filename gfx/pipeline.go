// Copyright (c) 2013-2016 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package gfx

//------------------------------------------------------------------------------

import (
	"io"

	"github.com/drakmaniso/glam/geom"
	"github.com/drakmaniso/glam/internal"
)

//------------------------------------------------------------------------------

type Pipeline struct {
	internal     internal.Pipeline
	clearColor   [4]float32
	attribStride map[uint32]uintptr
	isCompiled   bool
	isClosed     bool
}

//------------------------------------------------------------------------------

func (p *Pipeline) CompileShaders(
	vertexShader io.Reader,
	fragmentShader io.Reader,
) error {
	if err := p.internal.CompileShaders(vertexShader, fragmentShader); err != nil {
		return err
	}
	if err := p.internal.SetupVAO(); err != nil {
		return err
	}
	p.attribStride = make(map[uint32]uintptr)
	p.isCompiled = true
	return nil
}

//------------------------------------------------------------------------------

func (p *Pipeline) SetClearColor(color geom.Vec4) {
	p.clearColor[0] = color.X
	p.clearColor[1] = color.Y
	p.clearColor[2] = color.Z
	p.clearColor[3] = color.W
}

//------------------------------------------------------------------------------

func (p *Pipeline) Bind() {
	p.internal.Bind(p.clearColor)
}

//------------------------------------------------------------------------------

func (p *Pipeline) Close() {
	p.internal.Close()
	p.isClosed = true
}

//------------------------------------------------------------------------------
