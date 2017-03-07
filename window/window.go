// Copyright (c) 2013-2016 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

// Package window provides support for window events
package window

//------------------------------------------------------------------------------

import (
	"github.com/drakmaniso/glam/internal"
	"github.com/drakmaniso/glam/pixel"
)

//------------------------------------------------------------------------------

// Handler receives window events.
type Handler interface {
	WindowShown(timestamp uint32)
	WindowHidden(timestamp uint32)
	WindowResized(newSize pixel.Coord, timestamp uint32)
	WindowMinimized(timestamp uint32)
	WindowMaximized(timestamp uint32)
	WindowRestored(timestamp uint32)
	WindowMouseEnter(timestamp uint32)
	WindowMouseLeave(timestamp uint32)
	WindowFocusGained(timestamp uint32)
	WindowFocusLost(timestamp uint32)
	WindowQuit(timestamp uint32)
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
func Size() pixel.Coord {
	return pixel.Coord{X: internal.Window.Width, Y: internal.Window.Height}
}

//------------------------------------------------------------------------------
