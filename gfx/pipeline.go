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
}

//------------------------------------------------------------------------------

func NewPipeline(s ...Shader) (Pipeline, error) {
	var p Pipeline
	var err error
	p.internal, err = internal.NewPipeline()
	if err != nil {
		return Pipeline{}, err
	}
	p.attribStride = make(map[uint32]uintptr)
	for _, s := range s {
		if err := p.internal.UseShader(s.internal); err != nil {
			return p, err
		}
	}
	return p, nil
}

//------------------------------------------------------------------------------

func (p *Pipeline) ClearColor(color geom.Vec4) {
	p.clearColor[0] = color.X
	p.clearColor[1] = color.Y
	p.clearColor[2] = color.Z
	p.clearColor[3] = color.W
}

//------------------------------------------------------------------------------

func (p *Pipeline) UniformBuffer(binding uint32, b Buffer) {
	p.internal.UniformBuffer(binding, b.internal)
}

//------------------------------------------------------------------------------

func (p *Pipeline) Bind() {
	p.internal.Bind(p.clearColor)
}

//------------------------------------------------------------------------------

func (p *Pipeline) Close() {
	p.internal.Close()
}

//------------------------------------------------------------------------------
