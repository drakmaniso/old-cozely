// Copyright (c) 2013-2016 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

// Package window provides support for window events
package window

//------------------------------------------------------------------------------

import (
	"time"

	"github.com/drakmaniso/glam/geom"
	"github.com/drakmaniso/glam/internal"
)

//------------------------------------------------------------------------------

// Handler receives window events.
type Handler interface {
	WindowShown(timestamp time.Duration)
	WindowHidden(timestamp time.Duration)
	WindowResized(newSize geom.Vec2, timestamp time.Duration)
	WindowMinimized(timestamp time.Duration)
	WindowMaximized(timestamp time.Duration)
	WindowRestored(timestamp time.Duration)
	WindowMouseEnter(timestamp time.Duration)
	WindowMouseLeave(timestamp time.Duration)
	WindowFocusGained(timestamp time.Duration)
	WindowFocusLost(timestamp time.Duration)
	WindowQuit(timestamp time.Duration)
}

// Handle is the current handlers for window events
//
// It can be changed while the loop is running, but must never be nil.
var Handle Handler

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
func Size() geom.Vec2 {
	return geom.Vec2{X: internal.Window.Width, Y: internal.Window.Height}
}

//------------------------------------------------------------------------------
