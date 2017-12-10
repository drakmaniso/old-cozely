// Copyright (c) 2013-2017 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package gfx

//------------------------------------------------------------------------------

import (
	"strings"

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

	createScreen(internal.Config.FramebufferSize, internal.Config.PixelSize)

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
	if palette.changed {
		paletteBuffer.SubData(colours[:], 0)
		palette.changed = false
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

	blitScreen(internal.Window.Size)
	return nil
}

//------------------------------------------------------------------------------
