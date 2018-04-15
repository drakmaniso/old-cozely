// Copyright (c) 2013-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package cozely

import (
	"github.com/cozely/cozely/color"
	"github.com/cozely/cozely/internal"
	"github.com/cozely/cozely/x/gl"
)

////////////////////////////////////////////////////////////////////////////////

// GameLoop methods are called in a loop to React to player actions, Update the
// game state, and Render it.
type GameLoop interface {
	// Enter is called once, after the framework initialization, but before the
	// loop is started.
	Enter() error

	// Leave is called when the loop is stopped.
	Leave() error

	// React is called as often as possible, before Update and Render, to react to
	// the player's actions. This is the only method that is guaranteed to run at
	// least once per frame.
	React() error

	// Update is called at fixed intervals, to update the game state (e.g. logic,
	// physics, AI...).
	Update() error

	// Render is called to display the game state to the player.
	//
	// Note that the framerate of Update and Render is independent, so the game
	// state might need to be interpolated (see UpdateLag).
	Render() error
}

////////////////////////////////////////////////////////////////////////////////

// Events holds the callbacks for each window events.
//
// They can be modified at anytime, but should always contain valid callbacks
// (i.e., non nil). The change will take effect at the next call to the React
// method of the game loop.
var Events = struct {
	Resize  func()
	Hide    func()
	Show    func()
	Focus   func()
	Unfocus func()
	Quit    func()
}{
	Resize:  func() {},
	Hide:    func() {},
	Show:    func() {},
	Focus:   func() {},
	Unfocus: func() {},
	Quit:    func() { Stop() },
}

////////////////////////////////////////////////////////////////////////////////

// Run initializes the framework, load the assets and starts the game loop.
//
// The Update method is called with a fixed time step, while the Render method
// is tied to the framerate. The React method is called before each Update, but
// at least once for every frame. The loop runs until Stop() is called.
//
// Important: must be called from main.main, or at least from a function that is
// known to run on the main OS thread.
func Run(loop GameLoop) (err error) {
	defer func() {
		internal.Running = false
		internal.QuitRequested = false

		derr := internal.VectorCleanup()
		if err == nil && derr != nil {
			err = internal.Error("vector cleanup", derr)
			return
		}

		derr = internal.PolyCleanup()
		if err == nil && derr != nil {
			err = internal.Error("poly cleanup", derr)
			return
		}

		derr = internal.PixelCleanup()
		if err == nil && derr != nil {
			err = internal.Error("pixel cleanup", derr)
			return
		}

		derr = internal.ColorCleanup()
		if err == nil && derr != nil {
			err = internal.Error("color cleanup", derr)
			return
		}

		derr = internal.GLCleanup()
		if err == nil && derr != nil {
			err = internal.Error("gl cleanup", derr)
			return
		}

		derr = internal.Cleanup()
		if err == nil && derr != nil {
			err = internal.Error("internal cleanup", derr)
			return
		}
	}()

	if internal.Running {
		//TODO:
		return nil
	}

	internal.Loop = loop

	// Setup

	err = internal.Setup()
	if err != nil {
		return internal.Error("internal setup", err)
	}

	err = internal.GLSetup()
	if err != nil {
		return internal.Error("gl setup", err)
	}

	err = internal.ColorSetup()
	if err != nil {
		return internal.Error("color setup", err)
	}

	err = internal.PixelSetup()
	if err != nil {
		return internal.Error("pixel setup", err)
	}

	err = internal.PolySetup()
	if err != nil {
		return internal.Error("poly setup", err)
	}

	err = internal.VectorSetup()
	if err != nil {
		return internal.Error("vector setup", err)
	}

	// First, send a fake resize window event
	internal.PixelResize()
	Events.Resize()

	// Main Loop

	internal.Running = true

	internal.RenderTime = 0.0
	internal.UpdateLag = 0.0

	then := internal.GetSeconds()
	now := then
	gametime := 0.0
	internal.GameTime = gametime

	err = internal.Loop.Enter()
	if err != nil {
		return err
	}

	for !internal.QuitRequested {
		internal.RenderTime = now - then
		countFrames()
		if internal.RenderTime > 4*internal.UpdateStep {
			// Prevent "spiral of death" when Render can't keep up with Update
			internal.RenderTime = 4 * internal.UpdateStep
		}

		// Update and Events

		internal.UpdateLag += internal.RenderTime
		//TODO: ProcessEvents should always be called with GameTime = now!
		if internal.UpdateLag < internal.UpdateStep {
			// Process events even if there is no Update this frame
			internal.GameTime = now //TODO: check if correct
			internal.ProcessEvents(Events)
			internal.InputNewFrame() //TODO: error handling?
			internal.Loop.React()
		}
		for internal.UpdateLag >= internal.UpdateStep {
			// Do the Time Step
			internal.UpdateLag -= internal.UpdateStep
			gametime += internal.UpdateStep
			internal.GameTime = gametime
			// Events
			internal.ProcessEvents(Events)
			internal.InputNewFrame() //TODO: error handling?
			internal.Loop.React()
			// Update
			err = internal.Loop.Update()
			if err != nil {
				return internal.Error("in Update callback", err)
			}
		}

		// Render

		gl.DefaultFramebuffer.Bind(gl.DrawFramebuffer)
		gl.ClearColorBuffer(color.LRGBA{0, 0, 0, 0}) //TODO: ...

		internal.GameTime = gametime + internal.UpdateLag //TODO: check if correct
		err = internal.Loop.Render()
		if err != nil {
			return internal.Error("in Render callback", err)
		}

		err = internal.VectorDraw()
		if err != nil {
			return internal.Error("in vector draw", err)
		}

		internal.SwapWindow()

		then = now
		now = internal.GetSeconds()
	}

	err = internal.Loop.Leave()
	if err != nil {
		return err
	}

	return nil
}

////////////////////////////////////////////////////////////////////////////////

// GameTime returns the time elapsed in the game. It is updated before each call
// to Update and before each call to Render.
func GameTime() float64 {
	return internal.GameTime
}

// UpdateTime returns the time between previous update and current one. It is a
// fixed value, that only changes when configured with UpdateStep.
//
// See also UpdateLag.
func UpdateTime() float64 {
	return internal.UpdateStep
}

// RenderTime returns the time elapsed between the previous frame and the one
// being rendered.
//
// See also UpdateTime and UpdateLag.
func RenderTime() float64 {
	return internal.RenderTime
}

// UpdateLag returns the time elapsed between the last update and the frame
// being rendered. It should be used during Render to extrapolate (or
// interpolate) the game state.
//
// Note: if called during Update (or an event callback), it returns the time
// between the current update and the next Render call.
//
// See also UpdateTime and RenderTime.
func UpdateLag() float64 {
	return internal.UpdateLag
}

////////////////////////////////////////////////////////////////////////////////

// RenderStats returns the average durations of frames; it is updated 4
// times per second. It also returns the number of overruns (i.e. frame time
// longer than the threshold) during the last measurment interval.
func RenderStats() (t float64, overruns int) {
	return frAverage, xrunPrevious
}

func countFrames() {
	frCount++
	frSum += internal.RenderTime
	if internal.RenderTime > xrunThreshold {
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

const frInterval = 1.0 / 4.0

var frAverage float64
var frSum float64
var frCount int

const xrunThreshold float64 = 17 / 1000.0

var xrunCount, xrunPrevious int

////////////////////////////////////////////////////////////////////////////////

// Stop request the game loop to stop.
func Stop() {
	internal.QuitRequested = true
}

////////////////////////////////////////////////////////////////////////////////

// Path returns the (slash-separated) path of the executable, with a trailing
// slash.
func Path() string {
	return internal.Path
}

////////////////////////////////////////////////////////////////////////////////
