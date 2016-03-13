package internal

import (
	"github.com/drakmaniso/glam/geom"
)

var KeyState [512]bool

var MouseDelta geom.IVec2
var MousePosition geom.IVec2
var MouseButtons uint32
