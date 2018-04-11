// Copyright (c) 2013-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package glam

import (
	"github.com/drakmaniso/glam/colour"
	"github.com/drakmaniso/glam/internal"
	"github.com/drakmaniso/glam/x/gl"
)

//------------------------------------------------------------------------------

// GameLoop methods are called in a loop to React to user inputs, Update the
// game state, and Render it.
type GameLoop interface {
	// Initialization and cleanup
	Enter() error
	Leave() error

	// The loop itself
	React() error
	Update() error
	Render() error
}

//------------------------------------------------------------------------------

// Window holds the callbacks for window events
var Window = struct {
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

//------------------------------------------------------------------------------

// Run initializes the framework, load the assets and starts the game loop.
//
// The Update callback is called with a fixed time step, while the Draw callback
// is tied to the framerate. Event callbacks are called before each Update, but
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
			err = internal.Error("in vector cleanup", derr)
			return
		}

		derr = internal.PixelCleanup()
		if err == nil && derr != nil {
			err = internal.Error("in pixel cleanup", derr)
			return
		}

		derr = internal.PaletteCleanup()
		if err == nil && derr != nil {
			err = internal.Error("in palette cleanup", derr)
			return
		}

		derr = internal.Cleanup()
		if err == nil && derr != nil {
			err = internal.Error("in internal cleanup", derr)
			return
		}
	}()

	internal.Loop = loop

	// Setup

	err = internal.Setup()
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

	err = internal.VectorSetup()
	if err != nil {
		return internal.Error("in vector setup", err)
	}

	// First, send a fake resize window event
	internal.PixelResize()
	Window.Resize()

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
			// Prevent "spiral of death" when Draw can't keep up with Update
			internal.RenderTime = 4 * internal.UpdateStep
		}

		// Update and Events

		internal.UpdateLag += internal.RenderTime
		if internal.UpdateLag < internal.UpdateStep {
			// Process events even if there is no Update this frame
			internal.ProcessEvents(Window)
			internal.ActionNewFrame() //TODO: error handling?
			internal.Loop.React()
		}
		for internal.UpdateLag >= internal.UpdateStep {
			// Do the Time Step
			internal.UpdateLag -= internal.UpdateStep
			gametime += internal.UpdateStep
			internal.GameTime = gametime
			// Events
			internal.ProcessEvents(Window)
			internal.ActionNewFrame() //TODO: error handling?
			internal.Loop.React()
			// Update
			err = internal.Loop.Update()
			if err != nil {
				return internal.Error("in Update callback", err)
			}
		}

		// Draw

		gl.DefaultFramebuffer.Bind(gl.DrawFramebuffer)
		gl.ClearColorBuffer(colour.LRGBA{0, 0, 0, 0}) //TODO: ...

		internal.GameTime = gametime + internal.UpdateLag
		err = internal.Loop.Render()
		if err != nil {
			return internal.Error("in Draw callback", err)
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

//------------------------------------------------------------------------------

// GameTime returns the time elapsed in the game. It is updated before each call
// to Update and before each call to Draw.
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
// being drawn.
//
// See also UpdateTime and UpdateLag.
func RenderTime() float64 {
	return internal.RenderTime
}

// UpdateLag returns the time elapsed between the last update and the frame
// being drawn. It should be used during Draw to extrapolate (or interpolate)
// the game state.
//
// Note: if called during Update (or an event callback), it returns the time
// between the current update and the next Draw call.
//
// See also UpdateTime and RenderTime.
func UpdateLag() float64 {
	return internal.UpdateLag
}

//------------------------------------------------------------------------------

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

// RenderStats returns the average durations of frames; it is updated 4
// times per second. It also returns the number of overruns (i.e. frame time
// longer than the threshold) during the last measurment interval.
func RenderStats() (t float64, overruns int) {
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
