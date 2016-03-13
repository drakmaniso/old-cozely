package mouse

// #include "../internal/internal.h"
import "C"

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

// A ButtonState describes the state of all mouse buttons.
type ButtonState uint32

// IsPressed returns true if a specific button is held down in the
// specified button state.
func (s ButtonState) IsPressed(b Button) bool {
	var m uint32 = 1 << (b - 1)
	return uint32(s)&m != 0
}

// Buttons returns the current state of all mouse buttons.
// Updated at the start of each game loop iteration.
func Buttons() ButtonState {
	return ButtonState(C.mouseButtons)
}

//------------------------------------------------------------------------------
