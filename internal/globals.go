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
	"log"
	"os"
)

//------------------------------------------------------------------------------

// Path of the executable (uses os-dependant separator).
var FilePath string

// Path of the executable (uses slash separators, and ends with one).
var Path string

//------------------------------------------------------------------------------

// Config holds the initial configuration of the game.
var Config = struct {
	Debug          bool
	Title          string
	WindowSize     [2]int32
	ScreenSize     [2]int16
	PixelSize      int32
	ScreenMode     string // "Extend", "Zoom" or "Fixed"
	Multisample    int
	Display        int
	Fullscreen     bool
	FullscreenMode string
	VSync          bool
	PaletteAuto    bool
}{
	Debug:          false,
	Title:          "Carol",
	WindowSize:     [2]int32{1280, 720},
	ScreenSize:     [2]int16{320, 200},
	PixelSize:      4,
	ScreenMode:     "Extend",
	Multisample:    0,
	Display:        0,
	Fullscreen:     false,
	FullscreenMode: "Desktop",
	VSync:          true,
	PaletteAuto:    true,
}

//------------------------------------------------------------------------------

var (
	Log   logger = log.New(os.Stderr, "", log.Ltime|log.Lmicroseconds)
	Debug logger = nolog{}
)

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

// Window is the game window.
var Window struct {
	window        *C.SDL_Window
	context       C.SDL_GLContext
	Width, Height int32
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
var Loop GameLoop

//------------------------------------------------------------------------------

// KeyState holds the pressed state of all keys, indexed by position.
var KeyState [512]bool

//------------------------------------------------------------------------------

// MouseDelta holds the delta from last mouse position.
var MouseDeltaX, MouseDeltaY int32

// MousePosition holds the current mouse position.
var MousePositionX, MousePositionY int32

// MouseButtons holds the state of the mouse buttons.
var MouseButtons uint32

//------------------------------------------------------------------------------

var PixelSetup = func() error { return nil }
var PixelDraw = func() error { return nil }

var ResizeScreen = func() {}

//------------------------------------------------------------------------------
