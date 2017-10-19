// Copyright (c) 2013-2016 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package carol

//------------------------------------------------------------------------------

import (
	"errors"

	"github.com/drakmaniso/carol/gfx"
	"github.com/drakmaniso/carol/internal"
	"github.com/drakmaniso/carol/internal/microtext" //TODO: remove
	"github.com/drakmaniso/carol/internal/overl"
	"github.com/drakmaniso/carol/key"
	"github.com/drakmaniso/carol/mouse"
	"github.com/drakmaniso/carol/pixel"
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
	overl.Setup()

	isSetUp = true
	return nil
}

var isSetUp bool

//------------------------------------------------------------------------------

type Looper interface {
	// Window events
	WindowShown()
	WindowHidden()
	WindowResized(newSize pixel.Coord)
	WindowMinimized()
	WindowMaximized()
	WindowRestored()
	WindowMouseEnter()
	WindowMouseLeave()
	WindowFocusGained()
	WindowFocusLost()
	WindowQuit()

	// Keyboard events
	KeyDown(l key.Label, p key.Position)
	KeyUp(l key.Label, p key.Position)

	// Mouse events
	MouseMotion(motion pixel.Coord, position pixel.Coord)
	MouseButtonDown(b mouse.Button, clicks int)
	MouseButtonUp(b mouse.Button, clicks int)
	MouseWheel(motion pixel.Coord)

	// Update and Draw
	Update()
	Draw(dt, interpolation float64)
}

func Loop(l Looper) {
	loop = l
}

var loop Looper

//------------------------------------------------------------------------------

func SetTimeStep(t float64) {
	timeStep = t
}

func TimeStep() float64 {
	return timeStep
}

var timeStep float64 = 1.0 / 60

//------------------------------------------------------------------------------

// Run starts the game loop.
//
// The update callback is called with a fixed time step, while event handlers
// and the draw callback are called once for each frame displayed. The loop runs
// until Stop() is called.
//
// Important: must be called from main.main, or at least from a function that is
// known to run on the main OS thread.
func Run() error {
	defer internal.SDLQuit()
	defer internal.DestroyWindow()

	if !isSetUp {
		return errors.New("carol.Setup must be called before carol.LoopStable")
	}

	// Setup fallback handlers

	// First, send a fake resize window event
	{
		s := pixel.Coord{internal.Window.Width, internal.Window.Height}
		overl.WindowResized(s)
		loop.WindowResized(s)
	}

	// Main Loop

	then := internal.GetSeconds()
	now := then
	stepNow := now
	remain := 0.0

	for !internal.QuitRequested {
		now = internal.GetSeconds()
		delta = now - then
		//TODO: clamp delta ?
		countFrames()

		processEvents()

		// Update with fixed time step

		remain += delta
		// Cap remain to avoid "spiral of death"
		for remain > 8*timeStep {
			remain -= timeStep
			stepNow += timeStep
		}
		for remain >= timeStep {
			visibleNow = stepNow
			loop.Update()
			remain -= timeStep
			stepNow += timeStep
		}

		// Draw

		visibleNow = now
		loop.Draw(delta, remain/timeStep)
		microtext.Draw()
		overl.Draw()
		internal.SwapWindow()

		then = now
	}
	return nil
}

//------------------------------------------------------------------------------

// delta is the time elapsed between current and previous frames
var delta float64

//------------------------------------------------------------------------------

// Now returns the current time (elapsed since program start).
//
// If called during the update callback, it corresponds to the current time
// step. If called during the draw callback, it corresponds to the current
// frame. And if called during an event callback, it corresponds to the event
// time stamp.
//
// It shouldn't be used outside of these three contexts.
func Now() float64 {
	return visibleNow
}

var visibleNow float64

//------------------------------------------------------------------------------

func countFrames() {
	frCount++
	frSum += delta
	if delta > xrunThreshold {
		xrunCount++
	}
	if frSum >= frInterval {
		frAverage = frSum / float64(frCount)
		xrunPrevious = xrunCount
		microtext.PrintFrameTime(frAverage, xrunCount)
		frSum = 0
		frCount = 0
		xrunCount = 0
	}
}

// FrameTime returns the duration of the last frame
func FrameTime() float64 {
	return delta
}

// AverageFrameTime returns the average durations of frames; it is updated 4
// times per second. It also returns the number of overruns (i.e. frame time
// longer than the threshold) during the last measurment interval.
func AverageFrameTime() (t float64, overruns int) {
	return frAverage, xrunPrevious
}

const frInterval = 1.0 / 4.0

var frAverage float64
var frSum float64
var frCount int

const xrunThreshold float64 = 17 / 1000.0

var xrunCount, xrunPrevious int

//------------------------------------------------------------------------------

// Stop request the game loop to stop.
func Stop() {
	internal.QuitRequested = true
}

//------------------------------------------------------------------------------

// Path returns the (slash-separated) path of the executable, with a trailing
// slash.
func Path() string {
	return internal.Path
}

//------------------------------------------------------------------------------
