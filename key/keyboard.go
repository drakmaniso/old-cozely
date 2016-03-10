// Package key provides keyboard support
package key

// #include "../engine/engine.h"
import "C"

import (
	"unsafe"
)

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
