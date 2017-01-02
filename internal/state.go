// Copyright (c) 2013-2016 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package internal

//------------------------------------------------------------------------------

import "github.com/drakmaniso/glam/geom"

//------------------------------------------------------------------------------

// Debug specifies if the game runs in debug mode or not.
var Debug = false

// InitError is non-nil if an error occured during initialisation.
var InitError error

// QuitRequested makes the game loop stop if true.
var QuitRequested = false

//------------------------------------------------------------------------------

// KeyState holds the pressed state of all keys, indexed by position.
var KeyState [512]bool

//------------------------------------------------------------------------------

// MouseDelta holds the delta from last mouse position.
var MouseDelta geom.IVec2

// MousePosition holds the current mouse position.
var MousePosition geom.IVec2

// MouseButtons holds the state of the mouse buttons.
var MouseButtons uint32

//------------------------------------------------------------------------------
