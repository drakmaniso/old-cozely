// Copyright (c) 2013-2016 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package glam

import (
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

// Run opens the game window and runs a game loop, until Stop() is called.
//
// If the Looper is also a window, mouse or key Handler, it is set as the
// corresponding Handle.
//
// Important: must be called from main.main, or at least from a function that is
// known to run on the main OS thread.
func Run(l Looper) error {
	if internal.InitError != nil {
		return internal.InitError
	}
	defer internal.SDLQuit()
	defer internal.DestroyWindow()

	// Setup Handlers
	Loop = l
	if h, ok := l.(window.Handler); ok {
		window.Handle = h
	} else {
		window.Handle = basic.WindowHandler{}
	}
	if h, ok := l.(mouse.Handler); ok {
		mouse.Handle = h
	} else {
		mouse.Handle = basic.MouseHandler{}
	}
	if h, ok := l.(key.Handler); ok {
		key.Handle = h
	} else {
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
