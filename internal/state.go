// Copyright (c) 2013-2016 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package internal

//------------------------------------------------------------------------------

import "github.com/drakmaniso/glam/geom"

//------------------------------------------------------------------------------

var Debug = false
var InitError error
var QuitRequested = false

//------------------------------------------------------------------------------

var KeyState [512]bool

var MouseDelta geom.Vec2
var MousePosition geom.Vec2
var MouseButtons uint32

//------------------------------------------------------------------------------
