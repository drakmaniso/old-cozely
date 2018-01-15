// Copyright (c) 2013-2017 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package glam

import (
	"github.com/drakmaniso/glam/internal"
	"github.com/drakmaniso/glam/key"
	"github.com/drakmaniso/glam/mouse"
	"github.com/drakmaniso/glam/x/gl"
)

//------------------------------------------------------------------------------

// GameLoop methods are called to setup the game, and during the main loop to
// process events, Update the game state and Draw it.
type GameLoop interface {
	// Loop setup
	Setup() error

	// The loop
	Update() error
	Draw(delta float64, lerp float64) error

	// Window events
	WindowShown()
	WindowHidden()
	WindowResized(width, height int32)
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
	MouseMotion(deltaX, deltaY int32, posX, posY int32)
	MouseButtonDown(b mouse.Button, clicks int)
	MouseButtonUp(b mouse.Button, clicks int)
	MouseWheel(deltaX, deltaY int32)

	// Pixel events
	ScreenResized(width, height int16, pixel int32)
}

// Handlers implements default handlers for all events.
//
// It's an empty struct intended to be embedded in the user-defined GameLoop:
//
//  type loop struct {
//    glam.Handlers
//  }
//
// This way it's possible to implement the GameLoop interface without writing a
// method for each event.
type Handlers = internal.Handlers

//------------------------------------------------------------------------------

func SetTimeStep(t float64) {
	timeStep = t
}

func TimeStep() float64 {
	return timeStep
}

var timeStep = float64(1.0 / 60)

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
	defer internal.SDLQuit()
	defer internal.DestroyWindow()

	internal.Loop = loop

	// Setup

	err := internal.Setup()
	if err != nil {
		return internal.Error("in internal setup", err)
	}

	err = gl.Setup(internal.Config.Debug)
	if err != nil {
		return internal.Error("in OpenGL setup", err)
	}

	err = internal.PaletteSetup()
	if err != nil {
		return internal.Error("in palette setup", err)
	}

	err = internal.PixelSetup()
	if err != nil {
		return internal.Error("in pixel Setup", err)
	}

	err = internal.Loop.Setup()
	if err != nil {
		return internal.Error("in game loop Setup", err)
	}

	// First, send a fake resize window event
	internal.Loop.WindowResized(internal.Window.Width, internal.Window.Height)
	internal.ResizeScreen()

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

		internal.ProcessEvents() //TODO: Should it be in the physisc loop?

		// Update with fixed time step

		remain += delta
		// Cap remain to avoid "spiral of death"
		for remain > 8*timeStep {
			remain -= timeStep
			stepNow += timeStep
		}
		for remain >= timeStep {
			internal.VisibleNow = stepNow
			err = internal.Loop.Update()
			if err != nil {
				return internal.Error("in Update callback", err)
			}
			remain -= timeStep
			stepNow += timeStep
		}

		// Draw

		internal.VisibleNow = now
		err = internal.Loop.Draw(delta, remain/timeStep)
		if err != nil {
			return internal.Error("in Draw callback", err)
		}

		err = internal.PixelDraw()
		if err != nil {
			return internal.Error("in pixel Draw", err)
		}

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
	return internal.VisibleNow
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

// FrameTimeAverage returns the average durations of frames; it is updated 4
// times per second. It also returns the number of overruns (i.e. frame time
// longer than the threshold) during the last measurment interval.
func FrameTimeAverage() (t float64, overruns int) {
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
