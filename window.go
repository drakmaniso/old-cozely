// Copyright (c) 2013-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package cozely

import (
	"github.com/cozely/cozely/internal"
	"github.com/cozely/cozely/window"
)

////////////////////////////////////////////////////////////////////////////////

// HasFocus returns true if the game windows has focus.
func HasFocus() bool {
	return internal.HasFocus
}

// HasMouseFocus returns true if the mouse is currently inside the game window.
func HasMouseFocus() bool {
	return internal.HasMouseFocus
}

// WindowSize returns the size of the window in (screen) pixels.
func WindowSize() window.XY {
	//TODO: move to package window
	return window.XY{internal.Window.Width, internal.Window.Height}
}

////////////////////////////////////////////////////////////////////////////////
