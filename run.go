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

// Update is called by Loop and LoopStable. It should update the (physics and
// logic) state of the game.
var Update func(dt, now float64)

// Interpolate is called by LoopStable. It should interpolate between the current
// and previous state.
var Interpolate func(a float64)

// Draw is called by Loop and LoopStable. It should draw the current or
// interpolated state.
var Draw func()

//------------------------------------------------------------------------------

// Loop starts the game loop. There is one iteration for each frame displayed.
// Each iteration starts by calling the event handlers, then the Update
// callback, and finally the Draw callback. The loop runs until Stop() is
// called.
//
// Important: must be called from main.main, or at least from a function that is
// known to run on the main OS thread.
func Loop() error {
	defer internal.SDLQuit()
	defer internal.DestroyWindow()

	if !isSetUp {
		return errors.New("glam.Setup must be called before glam.Loop")
	}

	setupFallbackHandlers()

	// First, send a fake resize window event
	{
		s := pixel.Coord{internal.Window.Width, internal.Window.Height}
		window.Handle.WindowResized(s, 0)
	}

	// Main Loop

	internal.DrawInterpolation = 1.0
	then = internal.GetSeconds()

	for !internal.QuitRequested {
		now = internal.GetSeconds()
		updateFrameTime()

		events.Process()
		Interpolate(1.0)
		Update(frameTime, now)
		Draw()

		microtext.Draw()
		internal.SwapWindow()

		then = now
	}
	return nil
}

//------------------------------------------------------------------------------

// LoopStable starts the game loop, whith a fixed time step for the Update
// callback. Event handlers and the Draw callback are called once for each frame
// displayed. The loop runs until Stop() is called.
//
// Important: must be called from main.main, or at least from a function that is
// known to run on the main OS thread.
func LoopStable(timeStep float64) error {
	defer internal.SDLQuit()
	defer internal.DestroyWindow()

	if !isSetUp {
		return errors.New("glam.Setup must be called before glam.LoopStable")
	}

	setupFallbackHandlers()

	// First, send a fake resize window event
	{
		s := pixel.Coord{internal.Window.Width, internal.Window.Height}
		window.Handle.WindowResized(s, 0)
	}

	// Main Loop

	then = internal.GetSeconds()
	stepNow := then
	remain := 0.0

	for !internal.QuitRequested {
		now = internal.GetSeconds()
		updateFrameTime()

		// Fixed time step for logic and physics updates
		remain += frameTime
		for remain >= timeStep {
			events.Process()
			Update(timeStep, stepNow)
			remain -= timeStep
			stepNow += timeStep
		}

		// Interpolate and draw
		internal.DrawInterpolation = remain / timeStep
		Interpolate(internal.DrawInterpolation)
		Draw()

		microtext.Draw()
		internal.SwapWindow()

		then = now
	}
	return nil
}

//------------------------------------------------------------------------------

func setupFallbackHandlers() {
	if window.Handle == nil {
		window.Handle = basic.WindowHandler{}
	}
	if mouse.Handle == nil {
		mouse.Handle = basic.MouseHandler{}
	}
	if key.Handle == nil {
		key.Handle = basic.KeyHandler{}
	}
	if Update == nil {
		Update = func(_, _ float64) {}
	}
	if Interpolate == nil {
		Interpolate = func(_ float64) {}
	}
	if Draw == nil {
		Draw = func() {}
	}
}

//------------------------------------------------------------------------------

// now is the current time
var now, then float64

//------------------------------------------------------------------------------

func updateFrameTime() {
	frameTime = now - then
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
}

var frameTime float64

// FrameTime returns the duration of the last frame
func FrameTime() float64 {
	return frameTime
}

// AverageFrameTime returns the average durations of frames; it is updated 4
// times per second. It also returns the number of overruns (i.e. frame time
// longer than the threshold) during the last measurment interval.
func AverageFrameTime() (t float64, overruns int) {
	return ftAverage, xrunPrevious
}

const ftInterval = 1.0 / 4.0

var ftAverage float64
var ftSum float64
var ftCount int

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
