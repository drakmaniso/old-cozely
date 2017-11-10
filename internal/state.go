// Copyright (c) 2013-2016 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package internal

//------------------------------------------------------------------------------

import "github.com/drakmaniso/carol/pixel"

//------------------------------------------------------------------------------

type Looper interface {
	// Window events
	WindowShown()
	WindowHidden()
	WindowResized(newSize pixel.Coord)
	WindowMinimized()
	WindowMaximized()
	WindowRestored()
	WindowMouseEnter()
	WindowMouseLeave()
	WindowFocusGained()
	WindowFocusLost()
	WindowQuit()

	// Keyboard events
	KeyDown(l KeyLabel, p KeyPosition)
	KeyUp(l KeyLabel, p KeyPosition)

	// Mouse events
	MouseMotion(motion pixel.Coord, position pixel.Coord)
	MouseButtonDown(b MouseButton, clicks int)
	MouseButtonUp(b MouseButton, clicks int)
	MouseWheel(motion pixel.Coord)

	// Update and Draw
	Update()
	Draw(dt, interpolation float64)
}

var Loop Looper

//------------------------------------------------------------------------------

// VisibleNow is the current time (elapsed since program start).
//
// If called during the update callback, it corresponds to the current time
// step. If called during the draw callback, it corresponds to the current
// frame. And if called during an event callback, it corresponds to the event
// time stamp.
//
// It shouldn't be used outside of these three contexts.
var VisibleNow float64

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
