// Copyright (c) 2013-2017 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package pixel

//------------------------------------------------------------------------------

import (
	"github.com/drakmaniso/glam/x/gl"
)

//------------------------------------------------------------------------------

type stamp struct {
	//  word
	mode, mapping int16
	//  word
	x, y int16
}

var stamps []stamp

//------------------------------------------------------------------------------

var stampPipeline *gl.Pipeline

var stampSSBO gl.StorageBuffer

//------------------------------------------------------------------------------
