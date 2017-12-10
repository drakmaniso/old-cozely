// Copyright (c) 2013-2017 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package gfx

//------------------------------------------------------------------------------

import (
	"strings"

	"github.com/drakmaniso/carol/core/gl"
	"github.com/drakmaniso/carol/internal/core"
)

//------------------------------------------------------------------------------

func init() {
	var c core.Hook

	c = core.Hook{
		Callback: preSetupHook,
		Context:  "in gfx pre-Setup hook",
	}
	core.PreSetupHooks = append(core.PreSetupHooks, c)

	c = core.Hook{
		Callback: postDrawHook,
		Context:  "in gfx post-Draw hook",
	}
	core.PostDrawHooks = append(core.PostDrawHooks, c)
}

//------------------------------------------------------------------------------

func preSetupHook() error {
	var err error

	createScreen(core.Config.FramebufferSize, core.Config.PixelSize)

	stampPipeline = gl.NewPipeline(
		gl.VertexShader(strings.NewReader(vertexShader)),
		gl.FragmentShader(strings.NewReader(fragmentShader)),
		gl.Topology(gl.Triangles),
	)

	paletteBuffer = gl.NewStorageBuffer(uintptr(256*4*4), gl.DynamicStorage|gl.MapWrite)
	paletteBuffer.Bind(2)

	stampBuffer = gl.NewStorageBuffer(uintptr(1024), gl.DynamicStorage|gl.MapWrite)
	stampBuffer.Bind(0)

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
	if colChanged {
		paletteBuffer.SubData(colours[:], 0)
		colChanged = false
	}

	stampPipeline.Bind()
	gl.Blending(gl.SrcAlpha, gl.OneMinusSrcAlpha)
	gl.Enable(gl.Blend)

	if true {
		if len(stamps) > 0 {
			stampBuffer.SubData(stamps, 0)
			gl.Draw(0, int32(len(stamps)*6))
			stamps = stamps[:0]
		}
	}

	blitScreen(core.Window.Size)
	return nil
}

//------------------------------------------------------------------------------
