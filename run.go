// Copyright (c) 2013-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package cozely

import (
	"github.com/cozely/cozely/internal"
)

////////////////////////////////////////////////////////////////////////////////

// GameLoop methods are called in a loop to React to player actions, Update the
// game state, and Render it.
type GameLoop interface {
	// Enter is called once, after the framework initialization, but before the
	// loop is started.
	Enter()

	// Leave is called when the loop is stopped.
	Leave()

	// React is called as often as possible, before Update and Render, to react to
	// the player's actions. This is the only method that is guaranteed to run at
	// least once per frame.
	React()

	// Update is called at fixed intervals, to update the game state (e.g. logic,
	// physics, AI...).
	Update()

	// Render is called to display the game state to the player.
	//
	// Note that the framerate of Update and Render is independent, so the game
	// state might need to be interpolated (see UpdateLag).
	Render()
}

////////////////////////////////////////////////////////////////////////////////

// Events holds the callbacks for each window events.
//
// These callbacks can be modified at anytime, but should always contain valid
// functions (i.e., non nil). The change will take effect at the next frame.
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
	Quit:    func() { Stop(nil) },
}

////////////////////////////////////////////////////////////////////////////////

// Run initializes the framework, creates the ressources and loads all assets,
// and finally starts the game loop. The loop will run until Stop is called.
//
// Important: Run must be called from main.main, or at least from a function
// that is known to run on the main OS thread.
func Run(loop GameLoop) (err error) {
	defer func() {
		internal.Running = false
		internal.QuitRequested = false

		derr := internal.VectorCleanup()
		if err == nil && derr != nil {
			err = internal.Wrap("vector cleanup", derr)
			return
		}
		derr = internal.PolyCleanup()
		if err == nil && derr != nil {
			err = internal.Wrap("poly cleanup", derr)
			return
		}
		derr = internal.PixelCleanup()
		if err == nil && derr != nil {
			err = internal.Wrap("pixel cleanup", derr)
			return
		}
		derr = internal.ColorCleanup()
		if err == nil && derr != nil {
			err = internal.Wrap("color cleanup", derr)
			return
		}
		derr = internal.InputCleanup()
		if err == nil && derr != nil {
			err = internal.Wrap("input cleanup", derr)
			return
		}
		derr = internal.GLCleanup()
		if err == nil && derr != nil {
			err = internal.Wrap("gl cleanup", derr)
			return
		}
		derr = internal.Cleanup()
		if err == nil && derr != nil {
			err = internal.Wrap("internal cleanup", derr)
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
		return internal.Wrap("internal setup", err)
	}
	err = internal.GLSetup()
	if err != nil {
		return internal.Wrap("gl setup", err)
	}
	err = internal.InputSetup()
	if err != nil {
		return internal.Wrap("input setup", err)
	}
	err = internal.ColorSetup()
	if err != nil {
		return internal.Wrap("color setup", err)
	}
	err = internal.PixelSetup()
	if err != nil {
		return internal.Wrap("pixel setup", err)
	}
	err = internal.PolySetup()
	if err != nil {
		return internal.Wrap("poly setup", err)
	}
	err = internal.VectorSetup()
	if err != nil {
		return internal.Wrap("vector setup", err)
	}

	// First, send a fake resize window event
	internal.PixelResize()
	Events.Resize()

	// Main Loop

	internal.MouseShow(false)
	internal.Running = true

	internal.RenderDelta = 0.0
	internal.UpdateLag = 0.0

	then := internal.GetSeconds()
	now := then
	gametime := 0.0
	internal.GameTime = gametime

	internal.Loop.Enter()

	for !internal.QuitRequested {
		internal.RenderDelta = now - then
		countFrames()
		if internal.RenderDelta > 4*internal.UpdateStep {
			// Prevent "spiral of death" when Render can't keep up with Update
			internal.RenderDelta = 4 * internal.UpdateStep
		}

		// Update and Events

		internal.UpdateLag += internal.RenderDelta
		//TODO: ProcessEvents should always be called with GameTime = now!
		if internal.UpdateLag < internal.UpdateStep {
			// Process events even if there is no Update this frame
			internal.GameTime = now //TODO: check if correct
			internal.ProcessEvents(Events)
			internal.InputNewFrame()
			internal.Loop.React()
		}
		for internal.UpdateLag >= internal.UpdateStep {
			// Do the Time Step
			internal.UpdateLag -= internal.UpdateStep
			gametime += internal.UpdateStep
			internal.GameTime = gametime
			// Events
			internal.ProcessEvents(Events)
			internal.InputNewFrame()
			internal.Loop.React()
			// Update
			internal.Loop.Update()
		}

		// Render

		//TODO: render before react and update?
		err = internal.GLPrerender()
		if err != nil {
			return err
		}
		internal.GameTime = gametime + internal.UpdateLag //TODO: check if correct
		internal.Loop.Render()

		err = internal.VectorDraw()
		//TODO:
		if err != nil {
			return internal.Wrap("vector draw", err)
		}

		internal.SwapWindow()

		then = now
		now = internal.GetSeconds()
	}

	internal.Loop.Leave()
	return stopErr
}

////////////////////////////////////////////////////////////////////////////////

// GameTime returns the time elapsed in the game. It is updated before each call
// to Update and before each call to Render.
func GameTime() float64 {
	return internal.GameTime
}

// UpdateDelta returns the time between previous update and current one. It is a
// fixed value, that only changes when configured with UpdateStep.
//
// See also UpdateLag.
func UpdateDelta() float64 {
	return internal.UpdateStep
}

// RenderDelta returns the time elapsed between the previous frame and the one
// being rendered.
//
// See also UpdateDelta and UpdateLag.
func RenderDelta() float64 {
	return internal.RenderDelta
}

// UpdateLag returns the time elapsed between the last Update and the frame
// being rendered. It should be used during Render to extrapolate (or
// interpolate) the game state.
//
// See also UpdateTime and RenderDelta.
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
	frSum += internal.RenderDelta
	if internal.RenderDelta > xrunThreshold {
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

// Path returns the (slash-separated) path of the executable, with a trailing
// slash.
func Path() string {
	return internal.Path
}

////////////////////////////////////////////////////////////////////////////////
