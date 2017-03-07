// Copyright (c) 2013-2016 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package glam

//------------------------------------------------------------------------------

import (
	"fmt"
	"os"

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
	microtext.WindowResized(s, 0)

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
		window.Handle.WindowResized(s, 0)
		events.Process()
	}

	// Main Loop

	then := internal.GetTime()
	remain := 0.0

	for !internal.QuitRequested {
		now = internal.GetTime()
		frameTime = now - then

		// Compute smoothed frame time
		ftSmoothed = (ftSmoothed * ftSmoothing) + frameTime*(1.0-ftSmoothing)
		// Compute average frame time
		ftCount++
		ftSum += frameTime
		if frameTime > xrunThreshold {
			xrunCount++
		}
		if ftSum >= ftInterval {
			ftAverage = ftSum / float64(ftCount)
			xrunPrevious = xrunCount
			ftSum = 0
			ftCount = 0
			xrunCount = 0
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
var now float64

// TimeStep is the fixed interval between each call to Update.
var TimeStep = 1.0 / 50.0

//------------------------------------------------------------------------------

var frameTime float64

// FrameTime returns the duration of the last frame
func FrameTime() float64 {
	return frameTime
}

// AverageFrameTime returns the average durations of frames; it is updated 4
// times per second.
func AverageFrameTime() float64 {
	return ftAverage
}

var ftAverage float64
var ftSum float64
var ftCount int

const ftInterval = 1.0 / 4.0

// Overruns returns the number of overruns (i.e. frame time longer than the
// threshold) during the last measurment interval.
func Overruns() int {
	return xrunPrevious
}

var xrunCount, xrunPrevious int

const xrunThreshold float64 = 17 / 1000.0

// SmoothedFrameTime returns the frame time smoothed over time.
func SmoothedFrameTime() float64 {
	return ftSmoothed
}

var ftSmoothing = 0.995
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
