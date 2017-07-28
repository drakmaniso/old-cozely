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

// Handlers implements default behavior for all events.
type Handlers struct {
	WindowHandlers
	MouseHandlers
	KeyHandlers
}

//------------------------------------------------------------------------------

// WindowHandlers implements default behavior for all window events.
type WindowHandlers struct{}

// WindowShown does nothing.
func (dh WindowHandlers) WindowShown() {}

// WindowHidden does nothing.
func (dh WindowHandlers) WindowHidden() {}

// WindowResized does nothing.
func (dh WindowHandlers) WindowResized(s pixel.Coord) {}

// WindowMinimized does nothing.
func (dh WindowHandlers) WindowMinimized() {}

// WindowMaximized does nothing.
func (dh WindowHandlers) WindowMaximized() {}

// WindowRestored does nothing.
func (dh WindowHandlers) WindowRestored() {}

// WindowMouseEnter does nothing.
func (dh WindowHandlers) WindowMouseEnter() {}

// WindowMouseLeave does nothing.
func (dh WindowHandlers) WindowMouseLeave() {}

// WindowFocusGained does nothing.
func (dh WindowHandlers) WindowFocusGained() {}

// WindowFocusLost does nothing.
func (dh WindowHandlers) WindowFocusLost() {}

// WindowQuit requests the game loop to stop.
func (dh WindowHandlers) WindowQuit() {
	internal.QuitRequested = true
}

//------------------------------------------------------------------------------

// MouseHandlers implements default behavior for all mouse events.
type MouseHandlers struct{}

// MouseMotion does nothing.
func (dh MouseHandlers) MouseMotion(rel pixel.Coord, pos pixel.Coord) {}

// MouseButtonDown does nothing.
func (dh MouseHandlers) MouseButtonDown(b mouse.Button, clicks int) {}

// MouseButtonUp does nothing.
func (dh MouseHandlers) MouseButtonUp(b mouse.Button, clicks int) {}

// MouseWheel does nothing.
func (dh MouseHandlers) MouseWheel(w pixel.Coord) {}

//------------------------------------------------------------------------------

// KeyHandlers implements default behavior for all keyboard events.
type KeyHandlers struct{}

// KeyDown requests the game loop to stop if Escape is pressed.
func (dh KeyHandlers) KeyDown(l key.Label, p key.Position) {
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
func (dh KeyHandlers) KeyUp(l key.Label, p key.Position) {
}

//------------------------------------------------------------------------------
