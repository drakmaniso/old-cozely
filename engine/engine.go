// Copyright (c) 2013-2016 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package engine

//------------------------------------------------------------------------------

import (
	"encoding/json"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"time"
	"unsafe"
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
		panic(getError())
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
func Run() (err error) {
	defer C.SDL_Quit()

	err = window.open(
		config.Title,
		config.Resolution,
		config.Display,
		config.Fullscreen,
		config.FullscreenMode,
		config.VSync,
	)
	if err != nil {
		log.Print(err)
		return
	}
	defer window.destroy()

	quit := false
	for ! quit {
		quit = processEvents()
		handleUpdate()
		doMainthread()
		handleDraw()
		<-time.After(10 * time.Millisecond)
	}

	return
}

func processEvents() bool {
	more := true
	for more {
		n := int(C.peepEvents())
		for i := 0; i < n; i++ {
			e := unsafe.Pointer(&C.events[i])
			t := dispatchEvent(e)
			if t == C.SDL_QUIT {
				return true
			}
		}
		more = n >= C.PEEP_SIZE
	}
	return false
}

func dispatchEvent(e unsafe.Pointer) (t C.Uint32) {
	t = ((*C.SDL_CommonEvent)(e))._type
	switch t {
	case C.SDL_QUIT:
		handleQuit()
	//TODO: Window Events
	case C.SDL_WINDOWEVENT:
	// Keybord Events
	case C.SDL_KEYDOWN:
		handleKeyDown()
	case C.SDL_KEYUP:
		handleKeyUp()
	// Mouse Events
	case C.SDL_MOUSEMOTION:
		handleMouseMotion()
	case C.SDL_MOUSEBUTTONDOWN:
		handleMouseButtonDown()
	case C.SDL_MOUSEBUTTONUP:
		handleMouseButtonUp()
	case C.SDL_MOUSEWHEEL:
		handleMouseWheel()
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
	return
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
