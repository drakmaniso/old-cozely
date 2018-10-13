// Copyright (c) 2013-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package color

import (
	"github.com/cozely/cozely/internal"
	"github.com/cozely/cozely/x/gl"
)

////////////////////////////////////////////////////////////////////////////////

var renderer = glRenderer{}

type glRenderer struct {
	SSBO gl.StorageBuffer
}

////////////////////////////////////////////////////////////////////////////////

func init() {
	internal.ColorSetup = renderer.setup
	internal.ColorCleanup = renderer.cleanup
	internal.ColorRender = renderer.render
}

func (r *glRenderer) setup() error {
	r.SSBO = gl.NewStorageBuffer(uintptr(256*4*4), gl.DynamicStorage|gl.MapWrite)
	return nil
}

func (r *glRenderer) cleanup() error {
	Clear()
	dirty = true

	return nil
}

func (r *glRenderer) render() error {
	if dirty {
		r.SSBO.SubData(colors[:], 0)
		dirty = false
	}
	r.SSBO.Bind(0)

	return nil
}
