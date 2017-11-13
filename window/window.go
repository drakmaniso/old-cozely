// Copyright (c) 2013-2017 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

// Package window provides support for window events
package window

//------------------------------------------------------------------------------

import (
	"github.com/drakmaniso/carol/internal"
	"github.com/drakmaniso/carol/pixel"
)

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
