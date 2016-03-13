package mouse

// #include "../internal/internal.h"
import "C"

import (
	"github.com/drakmaniso/glam/internal"
)

//------------------------------------------------------------------------------

// A Button on the mouse
type Button uint8

// Button constants
const (
	Left   Button = C.SDL_BUTTON_LEFT
	Middle Button = C.SDL_BUTTON_MIDDLE
	Right  Button = C.SDL_BUTTON_RIGHT
	Extra1 Button = C.SDL_BUTTON_X1
	Extra2 Button = C.SDL_BUTTON_X2
)

// IsPressed returns true if a specific button is held down.
func IsPressed(b Button) bool {
	var m uint32 = 1 << (b - 1)
	return internal.MouseButtons&m != 0
}

//------------------------------------------------------------------------------
