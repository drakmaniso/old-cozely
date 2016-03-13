// Package key provides keyboard support
package key

// #include "../internal/internal.h"
import "C"

import (
	"unsafe"
)

//------------------------------------------------------------------------------

var handler Handler

// SetHandler sets the handler for Mouse events.
func SetHandler(h Handler) {
	handler = h
}

// GetHandler returns the current handler for mouse events.
func GetHandler() Handler {
	return handler
}

// A Handler reacts to key events.
type Handler interface {
	KeyDown(l Label, p Position, time uint32)
	KeyUp(l Label, p Position, time uint32)
}

//------------------------------------------------------------------------------

// IsPressed returns true if the corresponding key position is currently
// held down.
func IsPressed(pos Position) bool {
	return (*[C.SDL_NUM_SCANCODES]uint8)(unsafe.Pointer(C.keystate))[pos] == 1
}

// Modifiers returns the current state of the keyboard modifiers (e.g. Shift,
// Ctrl, CapsLock...). The value is one or more Modifier constants OR'd
// together.
func Modifiers() Modifier {
	return Modifier(C.keymod)
}

//------------------------------------------------------------------------------
