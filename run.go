// Copyright (c) 2013-2016 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package glam

//------------------------------------------------------------------------------

import (
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

func UpdateWith(u func()) {
	update = u
}

var update func()

func DrawWith(d func(dt, interpolation float64)) {
	draw = d
}

var draw func(dt, interpolation float64)

//------------------------------------------------------------------------------

func SetTimeStep(t float64) {
	timeStep = t
}

func TimeStep() float64 {
	return timeStep
}

var timeStep float64 = 1 / 60

//------------------------------------------------------------------------------

// Loop starts the game loop.
//
// The update callback is called with a fixed time step, while event handlers
// and the draw callback are called once for each frame displayed. The loop runs
// until Stop() is called.
//
// Important: must be called from main.main, or at least from a function that is
// known to run on the main OS thread.
func Loop() error {
	defer internal.SDLQuit()
	defer internal.DestroyWindow()

	if !isSetUp {
		return errors.New("glam.Setup must be called before glam.LoopStable")
	}

	// Setup fallback handlers
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
	now := then
	stepNow := now
	remain := 0.0

	for !internal.QuitRequested {
		now = internal.GetSeconds()
		delta = now - then
		//TODO: clamp delta ?
		countFrames()

		events.Process()

		// Fixed time step for logic and physics updates
		remain += delta
		//TODO: add cap to avoid "spiral of death"
		for remain >= timeStep {
			visibleNow = stepNow
			update()
			remain -= timeStep
			stepNow += timeStep
		}

		visibleNow = now
		draw(delta, remain/timeStep)
		microtext.Draw()
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
// frame. It shouldn't be used outside of these two callbacks.
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
