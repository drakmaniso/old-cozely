// Copyright (c) 2013-2016 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package glam

//------------------------------------------------------------------------------

import (
	"fmt"
	"os"
	"time"

	"github.com/drakmaniso/glam/basic"
	"github.com/drakmaniso/glam/internal"
	"github.com/drakmaniso/glam/internal/events"
	"github.com/drakmaniso/glam/key"
	"github.com/drakmaniso/glam/mouse"
	"github.com/drakmaniso/glam/window"
)

//------------------------------------------------------------------------------

// A Looper implements a game loop.
type Looper interface {
	Update()
	Draw()
}

// Loop is the current game loop.
//
// It can be changed while the loop is running, but must never be nil.
var Loop Looper

//------------------------------------------------------------------------------

// Run opens the game window and runs the Loop, until Stop() is called.
//
// Important: must be called from main.main, or at least from a function that is
// known to run on the main OS thread.
func Run() error {
	if internal.InitError != nil {
		return internal.InitError
	}
	defer internal.SDLQuit()
	defer internal.DestroyWindow()

	// Setup Fallback Handlers
	if window.Handle == nil {
		window.Handle = basic.WindowHandler{}
	}
	if mouse.Handle == nil {
		mouse.Handle = basic.MouseHandler{}
	}
	if key.Handle == nil {
		key.Handle = basic.KeyHandler{}
	}

	// Main Loop

	then := internal.GetTime() * time.Millisecond
	remain := time.Duration(0)

	for !internal.QuitRequested {
		now = internal.GetTime() * time.Millisecond
		remain += now - then
		for remain >= TimeStep {
			// Fixed time step for logic and physics updates.
			events.Process()
			Loop.Update()
			remain -= TimeStep
		}
		Loop.Draw()
		internal.SwapWindow()
		if now-then < 10*time.Millisecond {
			// Prevent using too much CPU on empty loops.
			<-time.After(10 * time.Millisecond)
		}
		then = now
	}
	return nil
}

// now is the current time
var now time.Duration

// TimeStep is the fixed interval between each call to Update.
var TimeStep = 1 * time.Second / 50

//------------------------------------------------------------------------------

// Stop request the game loop to stop.
func Stop() {
	internal.QuitRequested = true
}

//------------------------------------------------------------------------------

// Path returns the executable path.
func Path() string {
	return internal.Path
}

//------------------------------------------------------------------------------

// ErrorDialog displays an error dialog box.
func ErrorDialog(e error) {
	err := internal.ErrorDialog(e.Error())
	if err != nil {
		fmt.Fprintln(os.Stderr, e.Error())
	}
}

//------------------------------------------------------------------------------
