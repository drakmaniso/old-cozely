// Copyright (c) 2013-2016 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package gfx

//------------------------------------------------------------------------------

import (
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

func (p *Pipeline) Create(s ...*Shader) error {
	if err := p.internal.Create(); err != nil {
		return err
	}
	for _, s := range s {
		if err := p.internal.UseShader(&s.internal); err != nil {
			return err
		}
	}
	p.attribStride = make(map[uint32]uintptr)
	p.isCompiled = true
	return nil
}

//------------------------------------------------------------------------------

func (p *Pipeline) ClearColor(color geom.Vec4) {
	p.clearColor[0] = color.X
	p.clearColor[1] = color.Y
	p.clearColor[2] = color.Z
	p.clearColor[3] = color.W
}

//------------------------------------------------------------------------------

func (p *Pipeline) UniformBuffer(binding uint32, b *Buffer) {
	p.internal.UniformBuffer(binding, &b.internal)
}

//------------------------------------------------------------------------------

func (p *Pipeline) Use() {
	p.internal.Use(p.clearColor)
}

//------------------------------------------------------------------------------

func (p *Pipeline) Close() {
	p.internal.Close()
	p.isClosed = true
}

//------------------------------------------------------------------------------
