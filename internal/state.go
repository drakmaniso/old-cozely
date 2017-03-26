// Copyright (c) 2013-2016 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package internal

//------------------------------------------------------------------------------

import "github.com/drakmaniso/glam/pixel"

//------------------------------------------------------------------------------

// QuitRequested makes the game loop stop if true.
var QuitRequested = false

//------------------------------------------------------------------------------

// KeyState holds the pressed state of all keys, indexed by position.
var KeyState [512]bool

//------------------------------------------------------------------------------

// MouseDelta holds the delta from last mouse position.
var MouseDelta pixel.Coord

// MousePosition holds the current mouse position.
var MousePosition pixel.Coord

// MouseButtons holds the state of the mouse buttons.
var MouseButtons uint32

//------------------------------------------------------------------------------

// DrawInterpolation is blend factor to use between previous and current
// (physics) state when rendering.
var DrawInterpolation float64

//------------------------------------------------------------------------------
