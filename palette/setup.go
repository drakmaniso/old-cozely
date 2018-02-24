// Copyright (c) 2013-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package palette

import (
	"github.com/drakmaniso/glam/internal"
	"github.com/drakmaniso/glam/x/gl"
)

//------------------------------------------------------------------------------

var ssbo gl.StorageBuffer

//------------------------------------------------------------------------------

func init() {
	internal.PaletteSetup = setupHook
}

func setupHook() error {
	Clear()
	ssbo = gl.NewStorageBuffer(uintptr(256*4*4), gl.DynamicStorage|gl.MapWrite)
	return gl.Err()
}

//------------------------------------------------------------------------------

func init() {
	internal.PaletteUpload = uploadHook
}

func uploadHook() error {
	if changed {
		ssbo.SubData(colours[:], 0)
		changed = false
	}
	ssbo.Bind(0)

	return gl.Err()
}

//------------------------------------------------------------------------------
