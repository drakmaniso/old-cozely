// Copyright (c) 2013-2017 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package carol

//------------------------------------------------------------------------------

import (
	"github.com/drakmaniso/carol/internal/core"
	"github.com/drakmaniso/carol/internal/gfx"
	"github.com/drakmaniso/carol/internal/gpu"
)

//------------------------------------------------------------------------------

// GameLoop methods are called to setup the game, and during the main loop to
// process events, Update the game state and Draw it.
//
// Methods to implement:
//
// 	 Setup() error
// 	 Update() error
//	 Draw(delta float64, lerp float64) error
//
// Plus all the event handlers: see Handlers.
type GameLoop = core.GameLoop

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
func Run(loop GameLoop) error {
	defer core.SDLQuit()
	defer core.DestroyWindow()

	core.Loop = loop

	// Setup

	err := core.Setup()
	if err != nil {
		return core.Error("in internal setup", err)
	}

	err = gpu.Setup(
		core.Config.Debug,
		// core.Window.Size,
		core.Config.FramebufferSize,
		core.Config.PixelSize,
	)
	if err != nil {
		return core.Error("in gpu setup", err)
	}

	err = gfx.ScanPictures()
	if err != nil {
		return core.Error("while scanning images", err)
	}
	err = gfx.LoadPictures()
	if err != nil {
		return core.Error("while loading images", err)
	}

	err = core.Loop.Setup()
	if err != nil {
		return core.Error("in Setup callback", err)
	}

	// First, send a fake resize window event
	core.Loop.WindowResized(core.Window.Size)

	// Main Loop

	then := core.GetSeconds()
	now := then
	stepNow := now
	remain := 0.0

	for !core.QuitRequested {
		now = core.GetSeconds()
		delta = now - then
		//TODO: clamp delta ?
		countFrames()

		core.ProcessEvents() //TODO: Should it be in the physisc loop?

		// Update with fixed time step

		remain += delta
		// Cap remain to avoid "spiral of death"
		for remain > 8*timeStep {
			remain -= timeStep
			stepNow += timeStep
		}
		for remain >= timeStep {
			core.VisibleNow = stepNow
			err = core.Loop.Update()
			if err != nil {
				return core.Error("in Update callback", err)
			}
			remain -= timeStep
			stepNow += timeStep
		}

		// Draw

		core.VisibleNow = now
		err = core.Loop.Draw(delta, remain/timeStep)
		if err != nil {
			return core.Error("in Draw callback", err)
		}

		gpu.BindStampPipeline()
		gpu.BlitFramebuffer(core.Window.Size)
		core.SwapWindow()

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
	return core.VisibleNow
}

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
		//TODO: microtext.PrintFrameTime(frAverage, xrunCount)
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
	core.QuitRequested = true
}

//------------------------------------------------------------------------------

// Path returns the (slash-separated) path of the executable, with a trailing
// slash.
func Path() string {
	return core.Path
}

//------------------------------------------------------------------------------
