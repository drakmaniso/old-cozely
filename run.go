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

// Run opens the game window and runs the main loop. It returns only once the
// user quits or closes the window.
//
// Important: must be called from main.main, or at least from a function that is
// known to run on the main OS thread.
func Run() error {
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
		doMainthread()
		Handler.Draw()
		internal.Render()
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

func doMainthread() {
	more := true
	for more {
		select {
		case f := <-mainthread:
			f()
		default:
			more = false
		}
	}
}

//------------------------------------------------------------------------------

// From a post by Russ Cox on go-nuts.
// See https://github.com/golang/go/wiki/LockOSThread

var mainthread = make(chan func())

// Do runs a function on the rendering thread.
func Do(f func()) {
	done := make(chan bool, 1)
	mainthread <- func() {
		f()
		done <- true
	}
	<-done
}

// Go runs a function on the rendering thread, without blocking.
func Go(f func()) {
	mainthread <- f
}

//------------------------------------------------------------------------------

// Stop request the engine to stop. No more events will be processed,
// and at most one Update and one Draw will be called.
func Stop() {
	internal.QuitRequested = true
}

//------------------------------------------------------------------------------
