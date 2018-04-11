// Copyright (c) 2013-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package poly

import (
	"github.com/drakmaniso/cozely/colour"
	"github.com/drakmaniso/cozely/x/gl"
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

var materialSSBO gl.StorageBuffer

type Material struct {
	diffuse colour.LRGB
}

//------------------------------------------------------------------------------
