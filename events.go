// Copyright (c) 2013-2016 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package glam

//------------------------------------------------------------------------------

/*
#cgo windows LDFLAGS: -lSDL2
#cgo linux freebsd darwin pkg-config: sdl2

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

import (
	"unsafe"

	"github.com/drakmaniso/glam/gfx"
	"github.com/drakmaniso/glam/internal"
	"github.com/drakmaniso/glam/internal/microtext"
	"github.com/drakmaniso/glam/key"
	"github.com/drakmaniso/glam/mouse"
	"github.com/drakmaniso/glam/pixel"
)

//------------------------------------------------------------------------------

// processEvents processes and dispatches all events.
func processEvents() {
	more := true
	for more && !internal.QuitRequested {
		n := peepEvents()
		for i := 0; i < n && !internal.QuitRequested; i++ {
			e := eventAt(i)
			dispatch(e)
		}
		more = n >= C.PEEP_SIZE
	}
}

func dispatch(e unsafe.Pointer) {
	ts := uint32(((*C.SDL_CommonEvent)(e)).timestamp)
	visibleNow = float64(ts) / 1000.0
	switch ((*C.SDL_CommonEvent)(e))._type {
	case C.SDL_QUIT:
		loop.WindowQuit()
	// Window Events
	case C.SDL_WINDOWEVENT:
		e := (*C.SDL_WindowEvent)(e)
		switch e.event {
		case C.SDL_WINDOWEVENT_NONE:
			// Ignore
		case C.SDL_WINDOWEVENT_SHOWN:
			loop.WindowShown()
		case C.SDL_WINDOWEVENT_HIDDEN:
			loop.WindowHidden()
		case C.SDL_WINDOWEVENT_EXPOSED:
			// Ignore
		case C.SDL_WINDOWEVENT_MOVED:
			// Ignore
		case C.SDL_WINDOWEVENT_RESIZED:
			internal.Window.Width = int32(e.data1)
			internal.Window.Height = int32(e.data2)
			s := pixel.Coord{X: int32(e.data1), Y: int32(e.data2)}
			gfx.Viewport(pixel.Coord{X: 0, Y: 0}, s)
			microtext.WindowResized(s)
			loop.WindowResized(s)
		case C.SDL_WINDOWEVENT_SIZE_CHANGED:
			//TODO
		case C.SDL_WINDOWEVENT_MINIMIZED:
			loop.WindowMinimized()
		case C.SDL_WINDOWEVENT_MAXIMIZED:
			loop.WindowMaximized()
		case C.SDL_WINDOWEVENT_RESTORED:
			loop.WindowRestored()
		case C.SDL_WINDOWEVENT_ENTER:
			internal.HasMouseFocus = true
			loop.WindowMouseEnter()
		case C.SDL_WINDOWEVENT_LEAVE:
			internal.HasMouseFocus = false
			loop.WindowMouseLeave()
		case C.SDL_WINDOWEVENT_FOCUS_GAINED:
			internal.HasFocus = true
			loop.WindowFocusGained()
		case C.SDL_WINDOWEVENT_FOCUS_LOST:
			internal.HasFocus = false
			loop.WindowFocusLost()
		case C.SDL_WINDOWEVENT_CLOSE:
			// Ignore
		default:
			//TODO: log.Print("unkown window event")
		}
	// Keyboard Events
	case C.SDL_KEYDOWN:
		e := (*C.SDL_KeyboardEvent)(e)
		if e.repeat == 0 {
			internal.KeyState[e.keysym.scancode] = true
			loop.KeyDown(
				key.Label(e.keysym.sym),
				key.Position(e.keysym.scancode),
			)
		}
	case C.SDL_KEYUP:
		e := (*C.SDL_KeyboardEvent)(e)
		internal.KeyState[e.keysym.scancode] = false
		loop.KeyUp(
			key.Label(e.keysym.sym),
			key.Position(e.keysym.scancode),
		)
	// Mouse Events
	case C.SDL_MOUSEMOTION:
		e := (*C.SDL_MouseMotionEvent)(e)
		rel := pixel.Coord{X: int32(e.xrel), Y: int32(e.yrel)}
		internal.MouseDelta = internal.MouseDelta.Plus(rel)
		internal.MousePosition = pixel.Coord{X: int32(e.x), Y: int32(e.y)}
		internal.MouseButtons = uint32(e.state)
		loop.MouseMotion(
			rel,
			internal.MousePosition,
		)
	case C.SDL_MOUSEBUTTONDOWN:
		e := (*C.SDL_MouseButtonEvent)(e)
		internal.MouseButtons |= 1 << (e.button - 1)
		loop.MouseButtonDown(
			mouse.Button(e.button),
			int(e.clicks),
		)
	case C.SDL_MOUSEBUTTONUP:
		e := (*C.SDL_MouseButtonEvent)(e)
		internal.MouseButtons &= ^(1 << (e.button - 1))
		loop.MouseButtonUp(
			mouse.Button(e.button),
			int(e.clicks),
		)
	case C.SDL_MOUSEWHEEL:
		e := (*C.SDL_MouseWheelEvent)(e)
		var d int32 = 1
		if e.direction == C.SDL_MOUSEWHEEL_FLIPPED {
			d = -1
		}
		loop.MouseWheel(
			pixel.Coord{X: int32(e.x) * d, Y: int32(e.y) * d},
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
