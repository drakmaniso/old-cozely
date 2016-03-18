// Copyright (c) 2013-2016 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package engine

import (
	"encoding/json"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"time"
	"unsafe"

	"github.com/drakmaniso/glam/geom"
	"github.com/drakmaniso/glam/internal"
	"github.com/drakmaniso/glam/key"
	"github.com/drakmaniso/glam/mouse"
	"github.com/drakmaniso/glam/window"
)

// #cgo windows LDFLAGS: -lSDL2
// #cgo linux freebsd darwin pkg-config: sdl2
// #include "engine.h"
import "C"

//------------------------------------------------------------------------------

var path = filepath.Dir(os.Args[0])

var config = struct {
	Title          string
	Resolution     [2]int
	Display        int
	Fullscreen     bool
	FullscreenMode string
	VSync          bool
}{
	Title:          "Glam",
	Resolution:     [2]int{1280, 720},
	Display:        0,
	Fullscreen:     false,
	FullscreenMode: "Desktop",
	VSync:          true,
}

//------------------------------------------------------------------------------

func init() {
	log.SetFlags(log.Lshortfile | log.Ltime)

	log.Printf("path = \"%s\"", path)

	loadConfig()

	runtime.LockOSThread()

	if errcode := C.SDL_Init(C.SDL_INIT_EVERYTHING); errcode != 0 {
		panic(internal.GetSDLError())
	}

	C.SDL_StopTextInput()
}

func loadConfig() {
	f, err := os.Open(path + "/init.json")
	if err != nil {
		log.Print(err)
		return
	}
	d := json.NewDecoder(f)
	err = d.Decode(&config)
	if err != nil {
		log.Panic(err)
	}
	log.Printf("config = %v\n", config)
}

//------------------------------------------------------------------------------

// Run opens the game window and runs the main loop. It returns only once the
// user quits or closes the window.
//
// Important: must be called from main.main, or at least from a function that is
// known to run on the main OS thread.
func Run() error {
	defer C.SDL_Quit()

	err := internal.OpenWindow(
		config.Title,
		config.Resolution,
		config.Display,
		config.Fullscreen,
		config.FullscreenMode,
		config.VSync,
	)
	if err != nil {
		log.Print(err)
		return err
	}
	defer internal.DestroyWindow()

	// Main Loop

	then := time.Duration(C.SDL_GetTicks()) * time.Millisecond
	remain := time.Duration(0)

	for !quit {
		now = time.Duration(C.SDL_GetTicks()) * time.Millisecond
		remain += now - then
		for remain >= TimeStep {
			// Fixed time step for logic and physics updates.
			processEvents()
			Handler.Update()
			remain -= TimeStep
		}
		doMainthread()
		Handler.Draw()
		if now-then < 10*time.Millisecond {
			// Prevent using too much CPU on empty loops.
			<-time.After(10 * time.Millisecond)
		}
		then = now
	}
	return nil
}

// now is the current time
var now time.Duration

// TimeStep is the fixed interval between each call to Update.
var TimeStep = 1 * time.Second / 50

var quit = false

func processEvents() {
	more := true
	for more && !quit {
		n := int(C.peepEvents())
		for i := 0; i < n && !quit; i++ {
			e := unsafe.Pointer(&C.events[i])
			dispatchEvent(e)
		}
		more = n >= C.PEEP_SIZE
	}
}

func dispatchEvent(e unsafe.Pointer) {
	ts := time.Duration(((*C.SDL_CommonEvent)(e)).timestamp) * time.Millisecond
	switch ((*C.SDL_CommonEvent)(e))._type {
	case C.SDL_QUIT:
		Handler.Quit()
	// Window Events
	case C.SDL_WINDOWEVENT:
		e := (*C.SDL_WindowEvent)(e)
		switch e.event {
		case C.SDL_WINDOWEVENT_NONE:
			// Ignore
		case C.SDL_WINDOWEVENT_SHOWN:
			window.Handler.WindowShown(ts)
		case C.SDL_WINDOWEVENT_HIDDEN:
			window.Handler.WindowHidden(ts)
		case C.SDL_WINDOWEVENT_EXPOSED:
			// Ignore
		case C.SDL_WINDOWEVENT_MOVED:
			// Ignore
		case C.SDL_WINDOWEVENT_RESIZED:
			window.Handler.WindowResized(geom.IVec2{X: int32(e.data1), Y: int32(e.data2)}, ts)
		case C.SDL_WINDOWEVENT_SIZE_CHANGED:
			//TODO
		case C.SDL_WINDOWEVENT_MINIMIZED:
			window.Handler.WindowMinimized(ts)
		case C.SDL_WINDOWEVENT_MAXIMIZED:
			window.Handler.WindowMaximized(ts)
		case C.SDL_WINDOWEVENT_RESTORED:
			window.Handler.WindowRestored(ts)
		case C.SDL_WINDOWEVENT_ENTER:
			internal.HasMouseFocus = true
			window.Handler.WindowMouseEnter(ts)
		case C.SDL_WINDOWEVENT_LEAVE:
			internal.HasMouseFocus = false
			window.Handler.WindowMouseLeave(ts)
		case C.SDL_WINDOWEVENT_FOCUS_GAINED:
			internal.HasFocus = true
			window.Handler.WindowFocusGained(ts)
		case C.SDL_WINDOWEVENT_FOCUS_LOST:
			internal.HasFocus = false
			window.Handler.WindowFocusLost(ts)
		case C.SDL_WINDOWEVENT_CLOSE:
			// Ignore
		default:
			log.Printf("Unkown window event.")
		}
	// Keyboard Events
	case C.SDL_KEYDOWN:
		e := (*C.SDL_KeyboardEvent)(e)
		if e.repeat == 0 {
			internal.KeyState[e.keysym.scancode] = true
			key.Handler.KeyDown(
				key.Label(e.keysym.sym),
				key.Position(e.keysym.scancode),
				ts,
			)
		}
	case C.SDL_KEYUP:
		e := (*C.SDL_KeyboardEvent)(e)
		internal.KeyState[e.keysym.scancode] = false
		key.Handler.KeyUp(
			key.Label(e.keysym.sym),
			key.Position(e.keysym.scancode),
			ts,
		)
	// Mouse Events
	case C.SDL_MOUSEMOTION:
		e := (*C.SDL_MouseMotionEvent)(e)
		rel := geom.IVec2{X: int32(e.xrel), Y: int32(e.yrel)}
		internal.MouseDelta = internal.MouseDelta.Plus(rel)
		internal.MousePosition = geom.IVec2{X: int32(e.x), Y: int32(e.y)}
		internal.MouseButtons = uint32(e.state)
		mouse.Handler.MouseMotion(
			rel,
			internal.MousePosition,
			ts,
		)
	case C.SDL_MOUSEBUTTONDOWN:
		e := (*C.SDL_MouseButtonEvent)(e)
		mouse.Handler.MouseButtonDown(
			mouse.Button(e.button),
			int(e.clicks),
			ts,
		)
	case C.SDL_MOUSEBUTTONUP:
		e := (*C.SDL_MouseButtonEvent)(e)
		mouse.Handler.MouseButtonUp(
			mouse.Button(e.button),
			int(e.clicks),
			ts,
		)
	case C.SDL_MOUSEWHEEL:
		e := (*C.SDL_MouseWheelEvent)(e)
		var d int32 = 1
		if e.direction == C.SDL_MOUSEWHEEL_FLIPPED {
			d = -1
		}
		mouse.Handler.MouseWheel(
			geom.IVec2{X: int32(e.x) * d, Y: int32(e.y) * d},
			ts,
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
		//TODO: remove
		log.Println("Unknown", ((*C.SDL_CommonEvent)(e))._type)
	}
}

func doMainthread() {
	more := true
	for more {
		select {
		case f := <-mainthread:
			f()
		default:
			more = false
		}
	}
}

//------------------------------------------------------------------------------

// From a post by Russ Cox on go-nuts.
// See https://github.com/golang/go/wiki/LockOSThread

var mainthread = make(chan func())

// Do runs a function on the rendering thread.
func Do(f func()) {
	done := make(chan bool, 1)
	mainthread <- func() {
		f()
		done <- true
	}
	<-done
}

// Go runs a function on the rendering thread, without blocking.
func Go(f func()) {
	mainthread <- f
}

//------------------------------------------------------------------------------

// Stop request the engine to stop. No more events will be processed,
// and at most one Update and one Draw will be called.
func Stop() {
	quit = true
}

//------------------------------------------------------------------------------
