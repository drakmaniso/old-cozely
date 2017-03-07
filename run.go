// Copyright (c) 2013-2016 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package glam

//------------------------------------------------------------------------------

import (
	"fmt"
	"os"
	"time"

	"errors"
	"github.com/drakmaniso/glam/basic"
	"github.com/drakmaniso/glam/gfx"
	"github.com/drakmaniso/glam/internal"
	"github.com/drakmaniso/glam/internal/events"
	"github.com/drakmaniso/glam/internal/microtext"
	"github.com/drakmaniso/glam/key"
	"github.com/drakmaniso/glam/mouse"
	"github.com/drakmaniso/glam/pixel"
	"github.com/drakmaniso/glam/window"
)

//------------------------------------------------------------------------------

func Setup() {
	s := pixel.Coord{internal.Window.Width, internal.Window.Height}
	gfx.Viewport(pixel.Coord{X: 0, Y: 0}, s)

	// Setup mtx
	microtext.Setup()
	microtext.WindowResized(s, internal.GetTime())

	isSetUp = true

	//TODO: error handling?
}

var isSetUp bool

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

	if !isSetUp {
		return errors.New("glam.Setup must be called before glam.Run")
	}

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

	// Process events once before the first time step
	{
		s := pixel.Coord{internal.Window.Width, internal.Window.Height}
		window.Handle.WindowResized(s, internal.GetTime())
		events.Process()
	}

	// Main Loop

	then := internal.GetTime() * time.Millisecond
	remain := time.Duration(0)

	for !internal.QuitRequested {
		now = internal.GetTime() * time.Millisecond
		frameTime = now - then

		// Compute smoothed frame time
		ftSmoothed = (ftSmoothed * ftSmoothing) + float64(frameTime)*(1.0-ftSmoothing)
		// Compute average frame time
		ftSum += frameTime
		ftCount++
		if ftSum >= ftInterval {
			ftAverage = float64(ftSum) / float64(ftCount)
			ftSum -= ftInterval
			ftCount = 0
		}

		// Fixed time step for logic and physics updates
		remain += frameTime
		for remain >= TimeStep {
			remain -= TimeStep
			events.Process()
			Loop.Update()
		}

		Loop.Draw()
		microtext.Draw()
		internal.SwapWindow()

		then = now
	}
	return nil
}

// now is the current time
var now time.Duration

// TimeStep is the fixed interval between each call to Update.
var TimeStep = 1 * time.Second / 50

//------------------------------------------------------------------------------

var frameTime time.Duration

// FrameTime returns the duration of the last frame
func FrameTime() time.Duration {
	return frameTime
}

// FrameAverage returns the average durations of frames; it is updated every 100
// frames.
func FrameAverage() float64 {
	return ftAverage / float64(time.Millisecond)
}

var ftAverage float64
var ftSum time.Duration
var ftCount int

const ftInterval = (time.Second / 4)

// FrameTimeSmoothed returns the frame time smoothed over time.
func FrameTimeSmoothed() float64 {
	return ftSmoothed / float64(time.Millisecond)
}

var ftSmoothing = 0.99
var ftSmoothed float64

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
	internal.Log("ErrorDialog: %s", e)
	err := internal.ErrorDialog(e.Error())
	if err != nil {
		fmt.Fprintln(os.Stderr, e.Error())
	}
}

//------------------------------------------------------------------------------
