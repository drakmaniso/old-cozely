// Copyright (c) 2013-2017 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package poly

//------------------------------------------------------------------------------

import (
	"github.com/drakmaniso/glam/color"
	"github.com/drakmaniso/glam/gfx"
)

//------------------------------------------------------------------------------

func SetupMaterialBuffer(p []Material) error {
	materialSSBO.Delete()
	return nil
}

//------------------------------------------------------------------------------

func BindMaterialBuffer() {
	materialSSBO.Bind(0)
}

//------------------------------------------------------------------------------

func DeleteMaterialBuffer() {
	materialSSBO.Delete()
}

//------------------------------------------------------------------------------

var materialSSBO gfx.StorageBuffer

type Material struct {
	diffuse color.RGB
}

//------------------------------------------------------------------------------
