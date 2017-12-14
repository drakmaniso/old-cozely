// Copyright (c) 2013-2017 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package pixel

//------------------------------------------------------------------------------

import (
	"strings"

	"github.com/drakmaniso/carol/colour"
	"github.com/drakmaniso/carol/core/gl"
	"github.com/drakmaniso/carol/internal"
)

//------------------------------------------------------------------------------

func init() {
	var c internal.Hook

	c = internal.Hook{
		Callback: preSetupHook,
		Context:  "in gfx pre-Setup hook",
	}
	internal.PreSetupHooks = append(internal.PreSetupHooks, c)

	c = internal.Hook{
		Callback: postDrawHook,
		Context:  "in gfx post-Draw hook",
	}
	internal.PostDrawHooks = append(internal.PostDrawHooks, c)
}

//------------------------------------------------------------------------------

func preSetupHook() error {
	var err error

	createScreen()

	stampPipeline = gl.NewPipeline(
		gl.VertexShader(strings.NewReader(vertexShader)),
		gl.FragmentShader(strings.NewReader(fragmentShader)),
		gl.Topology(gl.TriangleStrip),
	)

	screenUBO = gl.NewUniformBuffer(&screenUniforms, gl.DynamicStorage|gl.MapWrite)
	screenUBO.Bind(0)

	paletteSSBO = gl.NewStorageBuffer(uintptr(256*4*4), gl.DynamicStorage|gl.MapWrite)
	paletteSSBO.Bind(2)

	stampSSBO = gl.NewStorageBuffer(uintptr(1024), gl.DynamicStorage|gl.MapWrite)
	stampSSBO.Bind(0)

	err = gl.Err()
	if err != nil {
		return err
	}

	err = loadAllPictures()
	if err != nil {
		return err
	}

	return nil
}

//------------------------------------------------------------------------------

func postDrawHook() error {
	if palette.changed {
		paletteSSBO.SubData(colours[:], 0)
		palette.changed = false
	}

	screenUniforms.PixelSize.X = 1.0 / float32(screen.size.X)
	screenUniforms.PixelSize.Y = 1.0 / float32(screen.size.Y)
	screenUBO.SubData(&screenUniforms, 0)

	screen.buffer.Bind(gl.DrawReadFramebuffer)
	gl.Viewport(0, 0, int32(screen.size.X), int32(screen.size.Y))
	stampPipeline.Bind()
	gl.ClearColorBuffer(colour.RGBA{0, 0, 0, 0}) //TODO
	gl.Blending(gl.SrcAlpha, gl.OneMinusSrcAlpha)
	gl.Enable(gl.Blend)

	if true {
		if len(stamps) > 0 {
			stampSSBO.SubData(stamps, 0)
			gl.DrawInstanced(0, 4, int32(len(stamps)))
			stamps = stamps[:0]
		}
	}

	blitScreen()

	return nil
}

//------------------------------------------------------------------------------
