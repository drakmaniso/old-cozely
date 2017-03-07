// Copyright (c) 2013-2016 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package basic

//------------------------------------------------------------------------------

import (
	"github.com/drakmaniso/glam/internal"
	"github.com/drakmaniso/glam/internal/microtext"
	"github.com/drakmaniso/glam/key"
	"github.com/drakmaniso/glam/mouse"
	"github.com/drakmaniso/glam/pixel"
)

//------------------------------------------------------------------------------

// WindowHandler implements default behavior for all window events.
type WindowHandler struct{}

// WindowShown does nothing.
func (dh WindowHandler) WindowShown(timestamp uint32) {}

// WindowHidden does nothing.
func (dh WindowHandler) WindowHidden(timestamp uint32) {}

// WindowResized does nothing.
func (dh WindowHandler) WindowResized(s pixel.Coord, timestamp uint32) {}

// WindowMinimized does nothing.
func (dh WindowHandler) WindowMinimized(timestamp uint32) {}

// WindowMaximized does nothing.
func (dh WindowHandler) WindowMaximized(timestamp uint32) {}

// WindowRestored does nothing.
func (dh WindowHandler) WindowRestored(timestamp uint32) {}

// WindowMouseEnter does nothing.
func (dh WindowHandler) WindowMouseEnter(timestamp uint32) {}

// WindowMouseLeave does nothing.
func (dh WindowHandler) WindowMouseLeave(timestamp uint32) {}

// WindowFocusGained does nothing.
func (dh WindowHandler) WindowFocusGained(timestamp uint32) {}

// WindowFocusLost does nothing.
func (dh WindowHandler) WindowFocusLost(timestamp uint32) {}

// WindowQuit requests the game loop to stop.
func (dh WindowHandler) WindowQuit(timestamp uint32) {
	internal.QuitRequested = true
}

//------------------------------------------------------------------------------

// MouseHandler implements default behavior for all mouse events.
type MouseHandler struct{}

// MouseMotion does nothing.
func (dh MouseHandler) MouseMotion(rel pixel.Coord, pos pixel.Coord, timestamp uint32) {}

// MouseButtonDown does nothing.
func (dh MouseHandler) MouseButtonDown(b mouse.Button, clicks int, timestamp uint32) {}

// MouseButtonUp does nothing.
func (dh MouseHandler) MouseButtonUp(b mouse.Button, clicks int, timestamp uint32) {}

// MouseWheel does nothing.
func (dh MouseHandler) MouseWheel(w pixel.Coord, timestamp uint32) {}

//------------------------------------------------------------------------------

// KeyHandler implements default behavior for all keyboard events.
type KeyHandler struct{}

// KeyDown requests the game loop to stop if Escape is pressed.
func (dh KeyHandler) KeyDown(l key.Label, p key.Position, timestamp uint32) {
	if l == key.LabelEscape {
		internal.QuitRequested = true
	}
	if (key.IsPressed(key.PositionLAlt) || key.IsPressed(key.PositionRAlt)) &&
		(key.IsPressed(key.PositionLCtrl) || key.IsPressed(key.PositionRCtrl)) {
		switch l {
		case key.LabelKPEnter:
			microtext.ToggleBgAlpha()
		case key.LabelF11, key.LabelReturn:
			internal.ToggleFullscreen()
		}
	}
}

// KeyUp does nothing.
func (dh KeyHandler) KeyUp(l key.Label, p key.Position, timestamp uint32) {
}

//------------------------------------------------------------------------------
