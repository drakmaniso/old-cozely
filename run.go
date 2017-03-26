// Copyright (c) 2013-2016 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package glam

//------------------------------------------------------------------------------

import (
	"errors"
	"fmt"
	"os"

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

// Setup initializes all subsystems and open the game window.
func Setup() error {
	err := internal.Setup()
	if err != nil {
		return internal.Error("setting up internal", err)
	}

	gfx.Setup()
	if err != nil {
		return internal.Error("setting up gfx", err)
	}

	microtext.Setup()

	isSetUp = true
	return nil
}

var isSetUp bool

//------------------------------------------------------------------------------

// A Looper implements a game loop.
type Looper interface {
	Update(t, dt float64)
	Draw(interpolation float64)
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

	// First, send a fake resize window event
	{
		s := pixel.Coord{internal.Window.Width, internal.Window.Height}
		window.Handle.WindowResized(s, 0)
	}

	// Main Loop

	then := internal.GetSeconds()
	stepNow := then
	remain := 0.0

	for !internal.QuitRequested {
		now = internal.GetSeconds()
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
			microtext.PrintFrameTime(ftAverage, xrunCount)
			ftSum = 0
			ftCount = 0
			xrunCount = 0
		}

		//TODO: clamp frameTime ?

		// Process events
		events.Process()

		// Fixed time step for logic and physics updates
		remain += frameTime
		for remain >= TimeStep {
			remain -= TimeStep
			stepNow += TimeStep
			Loop.Update(stepNow, TimeStep)
		}

		internal.DrawInterpolation = remain / TimeStep
		Loop.Draw(internal.DrawInterpolation)
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

// Error returns nil if err is nil, or a wrapped error otherwise.
func Error(source string, err error) error {
	if err == nil {
		return nil
	}
	return wrappedError{source, err}
}

type wrappedError struct {
	source string
	err    error
}

func (e wrappedError) Error() string {
	msg := e.source + ":\n\t"
	a := e.err
	for b, ok := a.(wrappedError); ok; {
		msg += b.source + ":\n\t"
		a = b.err
		b, ok = a.(wrappedError)
	}
	return msg + a.Error()
}

func ShowError(source string, err error) {
	e := Error("setting up", err)
	Log("ERROR:\n\t%s", e)
}

// Log logs a formated message.
func Log(format string, v ...interface{}) {
	internal.Log(format, v...)
}

// ErrorDialog displays a message in an error dialog box.
func ErrorDialog(format string, v ...interface{}) {
	internal.DebugLog(format, v...)
	err := internal.ErrorDialog(format, v...)
	if err != nil {
		fmt.Fprintf(os.Stderr, format, v...)
	}
}

//------------------------------------------------------------------------------
