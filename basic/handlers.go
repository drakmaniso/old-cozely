package basic

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

func (dh WindowHandler) WindowShown(timestamp time.Duration)                {}
func (dh WindowHandler) WindowHidden(timestamp time.Duration)               {}
func (dh WindowHandler) WindowResized(s geom.Vec2, timestamp time.Duration) {}
func (dh WindowHandler) WindowMinimized(timestamp time.Duration)            {}
func (dh WindowHandler) WindowMaximized(timestamp time.Duration)            {}
func (dh WindowHandler) WindowRestored(timestamp time.Duration)             {}
func (dh WindowHandler) WindowMouseEnter(timestamp time.Duration)           {}
func (dh WindowHandler) WindowMouseLeave(timestamp time.Duration)           {}
func (dh WindowHandler) WindowFocusGained(timestamp time.Duration)          {}
func (dh WindowHandler) WindowFocusLost(timestamp time.Duration)            {}
func (dh WindowHandler) WindowQuit(timestamp time.Duration) {
	internal.QuitRequested = true
}

//------------------------------------------------------------------------------

// MouseHandler implements default behavior for all mouse events.
type MouseHandler struct{}

func (dh MouseHandler) MouseMotion(rel geom.Vec2, pos geom.Vec2, timestamp time.Duration)   {}
func (dh MouseHandler) MouseButtonDown(b mouse.Button, clicks int, timestamp time.Duration) {}
func (dh MouseHandler) MouseButtonUp(b mouse.Button, clicks int, timestamp time.Duration)   {}
func (dh MouseHandler) MouseWheel(w geom.Vec2, timestamp time.Duration)                     {}

//------------------------------------------------------------------------------

// KeyHandler implements default behavior for all keyboard events.
type KeyHandler struct{}

func (dh KeyHandler) KeyDown(l key.Label, p key.Position, timestamp time.Duration) {
	if l == key.LabelEscape {
		internal.QuitRequested = true
	}
}

func (dh KeyHandler) KeyUp(l key.Label, p key.Position, timestamp time.Duration) {
}

//------------------------------------------------------------------------------
