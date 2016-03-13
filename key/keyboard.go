// Package key provides keyboard support
package key

// #cgo windows LDFLAGS: -lSDL2
// #cgo linux freebsd darwin pkg-config: sdl2
// #include "../internal/internal.h"
import "C"

import (
	"github.com/drakmaniso/glam/internal"
)

//------------------------------------------------------------------------------

// Handler receives the key events.
var Handler interface {
	KeyDown(l Label, p Position, time uint32)
	KeyUp(l Label, p Position, time uint32)
}

//------------------------------------------------------------------------------

// IsPressed returns true if the corresponding key position is currently
// held down.
func IsPressed(pos Position) bool {
	return internal.KeyState[pos]
}

// LabelOf returns the key label at the specified position in the current
// layout.
func LabelOf(pos Position) Label {
	return Label(C.SDL_GetKeyFromScancode(C.SDL_Scancode(pos)))
}

// SearchPositionOf searches the current position of label in the current
// layout.
func SearchPositionOf(l Label) Position {
	return Position(C.SDL_GetScancodeFromKey(C.SDL_Keycode(l)))
}

// Modifiers returns the current state of the keyboard modifiers (e.g. Shift,
// Ctrl, CapsLock...). The value is one or more Modifier constants OR'd
// together.
func Modifiers() Modifier {
	return Modifier(C.keymod)
}

//------------------------------------------------------------------------------
