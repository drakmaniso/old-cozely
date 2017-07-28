// Copyright (c) 2013-2016 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package key

//------------------------------------------------------------------------------

// #cgo windows LDFLAGS: -lSDL2
// #cgo linux freebsd darwin pkg-config: sdl2
// #include "../sdl.h"
import "C"

import (
	"github.com/drakmaniso/glam/internal"
)

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
