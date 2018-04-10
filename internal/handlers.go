// Copyright (c) 2013-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package internal

//------------------------------------------------------------------------------

// Handlers implements default handlers for all events.
//
// It's an empty struct intended to be embedded in the user-defined GameLoop.
type Handlers struct{}

//------------------------------------------------------------------------------

// Enter does nothing
func (Handlers) Enter() error { return nil }

// React does nothing
func (Handlers) React() error { return nil }

// Update does nothing
func (Handlers) Update() error { return nil }

// Draw does nothing
func (Handlers) Draw() error { return nil }

// Leave does nothing
func (Handlers) Leave() error { return nil }

//------------------------------------------------------------------------------

// Resize does nothing.
func (Handlers) Resize() {}

// Hide does nothing.
func (Handlers) Hide() {}

// Show does nothing.
func (Handlers) Show() {}

// Focus does nothing.
func (Handlers) Focus() {}

// Unfocus does nothing.
func (Handlers) Unfocus() {}

// Quit requests the game loop to stop.
func (Handlers) Quit() {
	QuitRequested = true
}

//------------------------------------------------------------------------------

// MouseMotion does nothing.
func (Handlers) MouseMotion(dx, dy int32, x, y int32) {}

// MouseButtonDown does nothing.
func (Handlers) MouseButtonDown(b MouseButton, clicks int) {}

// MouseButtonUp does nothing.
func (Handlers) MouseButtonUp(b MouseButton, clicks int) {}

// MouseWheel does nothing.
func (Handlers) MouseWheel(dx, dy int32) {}

//------------------------------------------------------------------------------

// KeyDown requests the game loop to stop if Escape is pressed.
func (h Handlers) KeyDown(l KeyLabel, p KeyCode) {
	switch l {
	case '\033': // key.LabelEscape
		QuitRequested = true
	case (1 << 30) | 68: // key.LabelF11
		ToggleFullscreen()
	}
}

// KeyUp does nothing.
func (Handlers) KeyUp(l KeyLabel, p KeyCode) {}

//------------------------------------------------------------------------------

// ScreenResized does nothing
func (Handlers) ScreenResized(width, height int16, pixel int32) {}

//------------------------------------------------------------------------------
