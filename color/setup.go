// Copyright (c) 2013-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package color

import (
	"github.com/cozely/cozely/internal"
	"github.com/cozely/cozely/x/gl"
)

////////////////////////////////////////////////////////////////////////////////

var ssbo gl.StorageBuffer

////////////////////////////////////////////////////////////////////////////////

func init() {
	internal.PaletteSetup = setup
	internal.PaletteCleanup = cleanup
}

func setup() error {
	Clear()
	ssbo = gl.NewStorageBuffer(uintptr(256*4*4), gl.DynamicStorage|gl.MapWrite)
	return gl.Err()
}

func cleanup() error {
	Clear()
	for n := range palettes {
		if n != "MSX" &&
			n != "MSX2" &&
			n != "C64" &&
			n != "CPC" {
			delete(palettes, n)
		}
	}
	ssbo.Delete()
	return gl.Err()
}

////////////////////////////////////////////////////////////////////////////////

func init() {
	internal.PaletteUpload = upload
}

func upload() error {
	if changed {
		ssbo.SubData(colours[:], 0)
		changed = false
	}
	ssbo.Bind(0)

	return gl.Err()
}

////////////////////////////////////////////////////////////////////////////////
