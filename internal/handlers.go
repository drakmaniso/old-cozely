// Copyright (c) 2013-2017 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package internal

//------------------------------------------------------------------------------

// Handlers implements default handlers for all events.
//
// It's an empty struct intended to be embedded in the user-defined GameLoop.
type Handlers struct{}

//------------------------------------------------------------------------------

// WindowShown does nothing.
func (h Handlers) WindowShown() {}

// WindowHidden does nothing.
func (h Handlers) WindowHidden() {}

// WindowResized does nothing.
func (h Handlers) WindowResized(width, height int32) {}

// WindowMinimized does nothing.
func (h Handlers) WindowMinimized() {}

// WindowMaximized does nothing.
func (h Handlers) WindowMaximized() {}

// WindowRestored does nothing.
func (h Handlers) WindowRestored() {}

// WindowMouseEnter does nothing.
func (h Handlers) WindowMouseEnter() {}

// WindowMouseLeave does nothing.
func (h Handlers) WindowMouseLeave() {}

// WindowFocusGained does nothing.
func (h Handlers) WindowFocusGained() {}

// WindowFocusLost does nothing.
func (h Handlers) WindowFocusLost() {}

// WindowQuit requests the game loop to stop.
func (h Handlers) WindowQuit() {
	QuitRequested = true
}

//------------------------------------------------------------------------------

// MouseMotion does nothing.
func (h Handlers) MouseMotion(dx, dy int32, x, y int32) {}

// MouseButtonDown does nothing.
func (h Handlers) MouseButtonDown(b MouseButton, clicks int) {}

// MouseButtonUp does nothing.
func (h Handlers) MouseButtonUp(b MouseButton, clicks int) {}

// MouseWheel does nothing.
func (h Handlers) MouseWheel(dx, dy int32) {}

//------------------------------------------------------------------------------

// KeyDown requests the game loop to stop if Escape is pressed.
func (h Handlers) KeyDown(l KeyLabel, p KeyPosition) {
	switch l {
	case '\033': // key.LabelEscape
		QuitRequested = true
	case (1 << 30) | 68: // key.LabelF11
		ToggleFullscreen()
	}
}

// KeyUp does nothing.
func (h Handlers) KeyUp(l KeyLabel, p KeyPosition) {
}

//------------------------------------------------------------------------------
