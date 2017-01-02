// Copyright (c) 2013-2016 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package basic

//------------------------------------------------------------------------------

import (
	"time"

	"github.com/drakmaniso/glam/geom"
	"github.com/drakmaniso/glam/internal"
	"github.com/drakmaniso/glam/key"
	"github.com/drakmaniso/glam/mouse"
)

//------------------------------------------------------------------------------

// WindowHandler implements default behavior for all window events.
type WindowHandler struct{}

// WindowShown does nothing.
func (dh WindowHandler) WindowShown(timestamp time.Duration) {}

// WindowHidden does nothing.
func (dh WindowHandler) WindowHidden(timestamp time.Duration) {}

// WindowResized does nothing.
func (dh WindowHandler) WindowResized(s geom.IVec2, timestamp time.Duration) {}

// WindowMinimized does nothing.
func (dh WindowHandler) WindowMinimized(timestamp time.Duration) {}

// WindowMaximized does nothing.
func (dh WindowHandler) WindowMaximized(timestamp time.Duration) {}

// WindowRestored does nothing.
func (dh WindowHandler) WindowRestored(timestamp time.Duration) {}

// WindowMouseEnter does nothing.
func (dh WindowHandler) WindowMouseEnter(timestamp time.Duration) {}

// WindowMouseLeave does nothing.
func (dh WindowHandler) WindowMouseLeave(timestamp time.Duration) {}

// WindowFocusGained does nothing.
func (dh WindowHandler) WindowFocusGained(timestamp time.Duration) {}

// WindowFocusLost does nothing.
func (dh WindowHandler) WindowFocusLost(timestamp time.Duration) {}

// WindowQuit requests the game loop to stop.
func (dh WindowHandler) WindowQuit(timestamp time.Duration) {
	internal.QuitRequested = true
}

//------------------------------------------------------------------------------

// MouseHandler implements default behavior for all mouse events.
type MouseHandler struct{}

// MouseMotion does nothing.
func (dh MouseHandler) MouseMotion(rel geom.IVec2, pos geom.IVec2, timestamp time.Duration) {}

// MouseButtonDown does nothing.
func (dh MouseHandler) MouseButtonDown(b mouse.Button, clicks int, timestamp time.Duration) {}

// MouseButtonUp does nothing.
func (dh MouseHandler) MouseButtonUp(b mouse.Button, clicks int, timestamp time.Duration) {}

// MouseWheel does nothing.
func (dh MouseHandler) MouseWheel(w geom.IVec2, timestamp time.Duration) {}

//------------------------------------------------------------------------------

// KeyHandler implements default behavior for all keyboard events.
type KeyHandler struct{}

// KeyDown requests the game loop to stop if Escape is pressed.
func (dh KeyHandler) KeyDown(l key.Label, p key.Position, timestamp time.Duration) {
	if l == key.LabelEscape {
		internal.QuitRequested = true
	}
}

// KeyUp does nothing.
func (dh KeyHandler) KeyUp(l key.Label, p key.Position, timestamp time.Duration) {
}

//------------------------------------------------------------------------------
