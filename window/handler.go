// Package window porvides support for window events
package window

import (
	"time"

	"github.com/drakmaniso/glam/geom"
)

//------------------------------------------------------------------------------

// Handler receives window events.
var Handler interface {
	WindowShown(timestamp time.Duration)
	WindowHidden(timestamp time.Duration)
	WindowResized(s geom.IVec2, timestamp time.Duration)
	WindowMinimized(timestamp time.Duration)
	WindowMaximized(timestamp time.Duration)
	WindowRestored(timestamp time.Duration)
	WindowMouseEnter(timestamp time.Duration)
	WindowMouseLeave(timestamp time.Duration)
	WindowFocusGained(timestamp time.Duration)
	WindowFocusLost(timestamp time.Duration)
}

//------------------------------------------------------------------------------
