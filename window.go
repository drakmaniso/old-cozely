// Copyright (c) 2013-2017 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package carol

//------------------------------------------------------------------------------

import (
	"github.com/drakmaniso/carol/internal"
	"github.com/drakmaniso/carol/screen"
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

// WindowSize returns the size of the window in (screen) pixels.
func WindowSize() screen.Coord {
	return internal.Window.Size
}

//TODO
// FramebufferSize returns the size of the framebuffer in (framebuffer) pixels.
// func FramebufferSize() screen.Coord {
// 	return gpu.Framebuffer.Size
// }

//TODO
// PixelSize returns the size of framebuffer pixels, in screen pixels.
// func PixelSize() screen.Coord {
// 	return gpu.Framebuffer.PixelSize
// }

//------------------------------------------------------------------------------
