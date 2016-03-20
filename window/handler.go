// Package window provides support for window events
package window

import (
	"time"

	"github.com/drakmaniso/glam/geom"
	"github.com/drakmaniso/glam/internal"
)

//------------------------------------------------------------------------------

// Handler receives window events.
var Handler interface {
	WindowShown(timestamp time.Duration)
	WindowHidden(timestamp time.Duration)
	WindowResized(newSize geom.IVec2, timestamp time.Duration)
	WindowMinimized(timestamp time.Duration)
	WindowMaximized(timestamp time.Duration)
	WindowRestored(timestamp time.Duration)
	WindowMouseEnter(timestamp time.Duration)
	WindowMouseLeave(timestamp time.Duration)
	WindowFocusGained(timestamp time.Duration)
	WindowFocusLost(timestamp time.Duration)
}

//------------------------------------------------------------------------------

// HasFocus returns true if the game windows has focus.
func HasFocus() bool {
	return internal.HasFocus
}

// HasMouseFocus returns true if the mouse is currently inside the game window.
func HasMouseFocus() bool {
	return internal.HasMouseFocus
}

//------------------------------------------------------------------------------
