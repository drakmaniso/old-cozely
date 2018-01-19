// Copyright (c) 2013-2017 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package pixel

//------------------------------------------------------------------------------

import (
	"github.com/drakmaniso/glam/x/gl"
)

//------------------------------------------------------------------------------

var (
	commands   []gl.DrawIndirectCommand
	parameters []int16
)

var (
	pipeline      *gl.Pipeline
	commandsICBO  gl.IndirectBuffer
	parametersTBO gl.BufferTexture
)

const (
	maxCommandCount = 1024
	maxParamCount   = 8
)

//------------------------------------------------------------------------------
