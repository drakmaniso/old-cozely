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
	WindowQuit(timestamp time.Duration)
} = DefaultWindowHandler{}

// DefaultWindowHandler implements default behavior for all window events.
type DefaultWindowHandler struct{}

func (dh DefaultWindowHandler) WindowShown(timestamp time.Duration)                 {}
func (dh DefaultWindowHandler) WindowHidden(timestamp time.Duration)                {}
func (dh DefaultWindowHandler) WindowResized(s geom.IVec2, timestamp time.Duration) {}
func (dh DefaultWindowHandler) WindowMinimized(timestamp time.Duration)             {}
func (dh DefaultWindowHandler) WindowMaximized(timestamp time.Duration)             {}
func (dh DefaultWindowHandler) WindowRestored(timestamp time.Duration)              {}
func (dh DefaultWindowHandler) WindowMouseEnter(timestamp time.Duration)            {}
func (dh DefaultWindowHandler) WindowMouseLeave(timestamp time.Duration)            {}
func (dh DefaultWindowHandler) WindowFocusGained(timestamp time.Duration)           {}
func (dh DefaultWindowHandler) WindowFocusLost(timestamp time.Duration)             {}
func (dh DefaultWindowHandler) WindowQuit(timestamp time.Duration) {
	internal.QuitRequested = true
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

// Size returns the size of the window in pixels.
func Size() geom.IVec2 {
	return geom.IVec2{X: internal.Window.Width, Y: internal.Window.Height}
}

//------------------------------------------------------------------------------
