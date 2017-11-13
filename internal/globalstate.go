// Copyright (c) 2013-2017 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package internal

//------------------------------------------------------------------------------

/*
#include "sdl.h"
*/
import "C"

//------------------------------------------------------------------------------

import (
	"github.com/drakmaniso/carol/pixel"
)

//------------------------------------------------------------------------------

// Path of the executable (uses slash separators, and ends with one).
var Path string

// Config holds the initial configuration of the game.
var Config = struct {
	Title           string
	WindowSize      pixel.Coord
	FramebufferSize pixel.Coord
	PixelSize       int
	Display         int
	Fullscreen      bool
	FullscreenMode  string
	VSync           bool
	Debug           bool
}{
	Title:           "Carol",
	WindowSize:      pixel.Coord{1280, 720},
	FramebufferSize: pixel.Coord{64, 64},
	PixelSize:       8,
	Display:         0,
	Fullscreen:      false,
	FullscreenMode:  "Desktop",
	VSync:           true,
	Debug:           false,
}

//------------------------------------------------------------------------------

// Window is the game window.
var Window struct {
	window  *C.SDL_Window
	context C.SDL_GLContext
	Size    pixel.Coord
}

// Focus state
var (
	HasFocus      bool
	HasMouseFocus bool
)

//------------------------------------------------------------------------------

// Loop holds the active looper.
//
// Note: The variable is set with carol.Loop.
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
