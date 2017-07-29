// Copyright (c) 2013-2016 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package glam

//------------------------------------------------------------------------------

import (
	"github.com/drakmaniso/glam/internal"
	"github.com/drakmaniso/glam/internal/microtext"
	"github.com/drakmaniso/glam/key"
	"github.com/drakmaniso/glam/mouse"
	"github.com/drakmaniso/glam/pixel"
)

//------------------------------------------------------------------------------

// DefaultHandlers implements default behavior for all events.
type DefaultHandlers struct{}

//------------------------------------------------------------------------------

// WindowShown does nothing.
func (dh DefaultHandlers) WindowShown() {}

// WindowHidden does nothing.
func (dh DefaultHandlers) WindowHidden() {}

// WindowResized does nothing.
func (dh DefaultHandlers) WindowResized(s pixel.Coord) {}

// WindowMinimized does nothing.
func (dh DefaultHandlers) WindowMinimized() {}

// WindowMaximized does nothing.
func (dh DefaultHandlers) WindowMaximized() {}

// WindowRestored does nothing.
func (dh DefaultHandlers) WindowRestored() {}

// WindowMouseEnter does nothing.
func (dh DefaultHandlers) WindowMouseEnter() {}

// WindowMouseLeave does nothing.
func (dh DefaultHandlers) WindowMouseLeave() {}

// WindowFocusGained does nothing.
func (dh DefaultHandlers) WindowFocusGained() {}

// WindowFocusLost does nothing.
func (dh DefaultHandlers) WindowFocusLost() {}

// WindowQuit requests the game loop to stop.
func (dh DefaultHandlers) WindowQuit() {
	internal.QuitRequested = true
}

//------------------------------------------------------------------------------

// MouseMotion does nothing.
func (dh DefaultHandlers) MouseMotion(rel pixel.Coord, pos pixel.Coord) {}

// MouseButtonDown does nothing.
func (dh DefaultHandlers) MouseButtonDown(b mouse.Button, clicks int) {}

// MouseButtonUp does nothing.
func (dh DefaultHandlers) MouseButtonUp(b mouse.Button, clicks int) {}

// MouseWheel does nothing.
func (dh DefaultHandlers) MouseWheel(w pixel.Coord) {}

//------------------------------------------------------------------------------

// KeyDown requests the game loop to stop if Escape is pressed.
func (dh DefaultHandlers) KeyDown(l key.Label, p key.Position) {
	switch l {
	case key.LabelEscape:
		internal.QuitRequested = true
	case key.LabelF11:
		internal.ToggleFullscreen()
	case key.LabelF12:
		microtext.ToggleReverseVideo()
	}
}

// KeyUp does nothing.
func (dh DefaultHandlers) KeyUp(l key.Label, p key.Position) {
}

//------------------------------------------------------------------------------
