// Copyright (c) 2013-2016 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package glam

import (
	"time"

	"github.com/drakmaniso/glam/internal"
	"github.com/drakmaniso/glam/internal/events"
)

//------------------------------------------------------------------------------

// Handler implements the game loop.
var Handler interface {
	Update()
	Draw()
}

//------------------------------------------------------------------------------

// Run opens the game window and runs the game loop, until Stop() is called.
//
// Important: must be called from main.main, or at least from a function that is
// known to run on the main OS thread.
func Run() error {
	if internal.InitError != nil {
		return internal.InitError
	}
	defer internal.SDLQuit()
	defer internal.DestroyWindow()

	// Main Loop

	then := internal.GetTime() * time.Millisecond
	remain := time.Duration(0)

	for !internal.QuitRequested {
		now = internal.GetTime() * time.Millisecond
		remain += now - then
		for remain >= TimeStep {
			// Fixed time step for logic and physics updates.
			events.Process()
			Handler.Update()
			remain -= TimeStep
		}
		Handler.Draw()
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
