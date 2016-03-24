// Copyright (c) 2013-2016 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package internal

//------------------------------------------------------------------------------

import "github.com/drakmaniso/glam/geom"

//------------------------------------------------------------------------------

var InitError error
var QuitRequested = false

//------------------------------------------------------------------------------

var KeyState [512]bool

var MouseDelta geom.IVec2
var MousePosition geom.IVec2
var MouseButtons uint32

//------------------------------------------------------------------------------
