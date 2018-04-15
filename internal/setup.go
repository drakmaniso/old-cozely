// Copyright (c) 2013-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package internal

import (
	"encoding/json"
	"log"
	"os"
	"path/filepath"
	"runtime"
)

////////////////////////////////////////////////////////////////////////////////

/*
#cgo windows LDFLAGS: -lSDL2
#cgo linux freebsd darwin pkg-config: sdl2

#include "sdl.h"
*/
import "C"

////////////////////////////////////////////////////////////////////////////////

func init() {
	runtime.LockOSThread()

	_, err := os.Stat("testdata")
	if err == nil {
		FilePath = "testdata"
		Path = "testdata/"
	} else {
		FilePath, err = os.Executable()
		if err != nil {
			FilePath = filepath.Dir(os.Args[0])
		}
		FilePath = filepath.Dir(FilePath)
		FilePath, _ = filepath.EvalSymlinks(FilePath)
		FilePath = filepath.Clean(FilePath)
		Path = filepath.ToSlash(FilePath) + "/"
	}
}

////////////////////////////////////////////////////////////////////////////////

func Setup() error {
	// Load config file

	//TODO: test if file present.
	f, err := os.Open(Path + "init.json")
	if !os.IsNotExist(err) {
		if err != nil {
			return Wrap(`in configuration file "init.json" opening`, err)
		}
		d := json.NewDecoder(f)
		if err := d.Decode(&Config); err != nil {
			return Wrap(`in configuration file "init.json" parsing`, err)
		}
	}

	// Setup logger

	if Config.Debug {
		Debug = log.New(os.Stdout, "", log.Ltime|log.Lmicroseconds)
	}

	// Initialize SDL

	if errcode := C.SDL_Init(C.SDL_INIT_VIDEO | C.SDL_INIT_JOYSTICK); errcode != 0 {
		return Wrap("in SDL initialization", GetSDLError())
	}

	C.SDL_EventState(C.SDL_MOUSEMOTION, C.SDL_IGNORE)
	C.SDL_EventState(C.SDL_MOUSEBUTTONDOWN, C.SDL_IGNORE)
	C.SDL_EventState(C.SDL_MOUSEBUTTONUP, C.SDL_IGNORE)
	C.SDL_EventState(C.SDL_MOUSEBUTTONDOWN, C.SDL_IGNORE)

	C.SDL_EventState(C.SDL_KEYDOWN, C.SDL_IGNORE)
	C.SDL_EventState(C.SDL_KEYUP, C.SDL_IGNORE)

	C.SDL_StopTextInput()

	getKeyboardArray()

	// Open the window

	err = OpenWindow(
		Title,
		Config.WindowSize[0], Config.WindowSize[1],
		Config.Display,
		Config.Fullscreen,
		Config.FullscreenMode,
		Config.VSync,
		Config.Debug,
	)
	if err != nil {
		return Wrap("in window opening", err)
	}

	return nil
}

func Cleanup() error {
	destroyWindow()
	SDLQuit()
	return nil
}

////////////////////////////////////////////////////////////////////////////////
