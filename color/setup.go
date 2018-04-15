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
	internal.ColorSetup = setup
	internal.ColorCleanup = cleanup
}

func setup() error {
	ssbo = gl.NewStorageBuffer(uintptr(256*4*4), gl.DynamicStorage|gl.MapWrite)

	for id, pp := range palettes.path {
		if len(pp) > 0 {
			PaletteID(id).load(pp[0])
			//TODO: load remaing paths
		}
	}

	//TODO: PaletteID(0).Activate()

	return gl.Err()
}

func cleanup() error {
	for id := range palettes.path {
		palettes.changed[id] = true
	}
	ssbo.Delete()
	return gl.Err()
}

////////////////////////////////////////////////////////////////////////////////

func init() {
	internal.ColorUpload = upload
}

func upload() error {
	if activated || palettes.changed[active] {
		ssbo.SubData(colours[:], 0)
		activated = false
		palettes.changed[active] = false
	}
	ssbo.Bind(0)

	return gl.Err()
}

////////////////////////////////////////////////////////////////////////////////
