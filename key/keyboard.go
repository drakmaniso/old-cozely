// Package key provides keyboard support
package key

// #cgo windows LDFLAGS: -lSDL2
// #cgo linux freebsd darwin pkg-config: sdl2
// #include "../internal/sdl.h"
import "C"

import (
	"time"

	"github.com/drakmaniso/glam/internal"
)

//------------------------------------------------------------------------------

// Handler receives the key events.
var Handler interface {
	KeyDown(l Label, p Position, timestamp time.Duration)
	KeyUp(l Label, p Position, timestamp time.Duration)
} = DefaultHandler{}

// DefaultHandler implements default behavior for all keyboard events.
type DefaultHandler struct{}

func (dh DefaultHandler) KeyDown(l Label, p Position, timestamp time.Duration) {
	if l == LabelEscape {
		internal.QuitRequested = true
	}
}

func (dh DefaultHandler) KeyUp(l Label, p Position, timestamp time.Duration) {
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

//------------------------------------------------------------------------------
