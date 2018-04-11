// Copyright (c) 2013-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package internal

import (
	"unsafe"
)

//------------------------------------------------------------------------------

/*
#include "sdl.h"

#define PEEP_SIZE 128

SDL_Event Events[PEEP_SIZE];

int PeepEvents()
{
  SDL_PumpEvents();
  int n = SDL_PeepEvents(Events, PEEP_SIZE, SDL_GETEVENT, SDL_FIRSTEVENT, SDL_LASTEVENT);
  return n;
}
*/
import "C"

//------------------------------------------------------------------------------

// GameLoop methods are called to setup the game, and during the main loop to
// process events, Update the game state and Draw it.
type GameLoop interface {
	// The loop
	Enter() error
	React() error
	Update() error
	Draw() error
	Leave() error

	// Window events
	Resize()
	Hide()
	Show()
	Focus()
	Unfocus()
	Quit()

	// Keyboard events
	KeyDown(l KeyLabel, p KeyCode)
	KeyUp(l KeyLabel, p KeyCode)

	// Mouse events
	MouseMotion(deltaX, deltaY int32, posX, posY int32)
	MouseButtonDown(b MouseButton, clicks int)
	MouseButtonUp(b MouseButton, clicks int)
	MouseWheel(deltaX, deltaY int32)
}

//------------------------------------------------------------------------------

// ProcessEvents processes and dispatches all events.
func ProcessEvents() {
	more := true
	for more && !QuitRequested {
		n := peepEvents()

		var mx, my C.int
		btn := C.SDL_GetRelativeMouseState(&mx, &my)
		MouseDeltaX += int32(mx)
		MouseDeltaY += int32(my)
		C.SDL_GetMouseState(&mx, &my)
		MousePositionX = int32(mx)
		MousePositionY = int32(my)
		MouseButtons = uint32(btn)

		for i := 0; i < n && !QuitRequested; i++ {
			e := eventAt(i)
			dispatch(e)
		}
		more = n >= C.PEEP_SIZE
	}
}

func dispatch(e unsafe.Pointer) {
	switch ((*C.SDL_CommonEvent)(e))._type {
	case C.SDL_QUIT:
		Loop.Quit()
	// Window Events
	case C.SDL_WINDOWEVENT:
		e := (*C.SDL_WindowEvent)(e)
		switch e.event {
		case C.SDL_WINDOWEVENT_NONE:
			// Ignore
		case C.SDL_WINDOWEVENT_SHOWN:
			Loop.Show()
		case C.SDL_WINDOWEVENT_HIDDEN:
			Loop.Hide()
		case C.SDL_WINDOWEVENT_EXPOSED:
			// Ignore
		case C.SDL_WINDOWEVENT_MOVED:
			// Ignore
		case C.SDL_WINDOWEVENT_RESIZED:
			Window.Width = int16(e.data1)
			Window.Height = int16(e.data2)
			PixelResize()
			Loop.Resize()
		case C.SDL_WINDOWEVENT_SIZE_CHANGED:
			//TODO
		case C.SDL_WINDOWEVENT_MINIMIZED:
			//TODO: check that Hide is enough
		case C.SDL_WINDOWEVENT_MAXIMIZED:
			// Ingnore
		case C.SDL_WINDOWEVENT_RESTORED:
			//TODO: check that Show is enough
		case C.SDL_WINDOWEVENT_ENTER:
			HasMouseFocus = true
		case C.SDL_WINDOWEVENT_LEAVE:
			HasMouseFocus = false
		case C.SDL_WINDOWEVENT_FOCUS_GAINED:
			HasFocus = true
			Loop.Focus()
		case C.SDL_WINDOWEVENT_FOCUS_LOST:
			HasFocus = false
			Loop.Unfocus()
		case C.SDL_WINDOWEVENT_CLOSE:
			// Ignore
		default:
			//TODO: log.Print("unkown window event")
		}
	// Mouse Events
	case C.SDL_MOUSEWHEEL:
		e := (*C.SDL_MouseWheelEvent)(e)
		var d int32 = 1
		if e.direction == C.SDL_MOUSEWHEEL_FLIPPED {
			d = -1
		}
		Loop.MouseWheel(
			int32(e.x)*d, int32(e.y)*d,
		)
	//TODO: Joystick Events
	case C.SDL_JOYAXISMOTION:
	case C.SDL_JOYBALLMOTION:
	case C.SDL_JOYHATMOTION:
	case C.SDL_JOYBUTTONDOWN:
	case C.SDL_JOYBUTTONUP:
	case C.SDL_JOYDEVICEADDED:
	case C.SDL_JOYDEVICEREMOVED:
	//TODO: Controller Events
	case C.SDL_CONTROLLERAXISMOTION:
	case C.SDL_CONTROLLERBUTTONDOWN:
	case C.SDL_CONTROLLERBUTTONUP:
	case C.SDL_CONTROLLERDEVICEADDED:
	case C.SDL_CONTROLLERDEVICEREMOVED:
	case C.SDL_CONTROLLERDEVICEREMAPPED:
	//TODO: Audio Device Events
	case C.SDL_AUDIODEVICEADDED:
	case C.SDL_AUDIODEVICEREMOVED:
	default:
		//TODO: log.Print("unknown SDL event:", ((*C.SDL_CommonEvent)(e))._type)
	}
}

// peepEvents fill the event buffer and returns the number of events fetched.
func peepEvents() int {
	return int(C.PeepEvents())
}

// EventAt returns a pointer to an event in the event buffer.
func eventAt(i int) unsafe.Pointer {
	return unsafe.Pointer(&C.Events[i])
}

//------------------------------------------------------------------------------

// SDLQuit is called when the game loop stops.
func SDLQuit() {
	C.SDL_Quit()
}

//------------------------------------------------------------------------------
