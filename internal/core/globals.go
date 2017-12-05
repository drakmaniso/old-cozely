// Copyright (c) 2013-2017 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package core

//------------------------------------------------------------------------------

/*
#include "sdl.h"
*/
import "C"

//------------------------------------------------------------------------------

import (
	"log"
	"os"

	"github.com/drakmaniso/carol/screen"
)

//------------------------------------------------------------------------------

// Path of the executable (uses slash separators, and ends with one).
var Path string

// Config holds the initial configuration of the game.
var Config = struct {
	Title           string
	WindowSize      screen.Coord
	FramebufferSize screen.Coord
	PixelSize       screen.Coord
	Display         int
	Fullscreen      bool
	FullscreenMode  string
	VSync           bool
	Debug           bool
}{
	Title:           "Carol",
	WindowSize:      screen.Coord{X: 1280, Y: 720},
	FramebufferSize: screen.Coord{X: 64, Y: 64},
	PixelSize:       screen.Coord{X: 8, Y: 8},
	Display:         0,
	Fullscreen:      false,
	FullscreenMode:  "Desktop",
	VSync:           true,
	Debug:           false,
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
	window  *C.SDL_Window
	context C.SDL_GLContext
	Size    screen.Coord
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
var MouseDelta screen.Coord

// MousePosition holds the current mouse position.
var MousePosition screen.Coord

// MouseButtons holds the state of the mouse buttons.
var MouseButtons uint32

//------------------------------------------------------------------------------

type Hook struct {
	Callback func() error
	Context  string
}

var PostSetupHooks = []Hook{}
var PostDrawHooks = []Hook{}

//------------------------------------------------------------------------------
