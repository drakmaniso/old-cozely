// Copyright (c) 2013-2017 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package core

//------------------------------------------------------------------------------

/*
#cgo windows LDFLAGS: -lSDL2
#cgo linux freebsd darwin pkg-config: sdl2

#include "sdl.h"
*/
import "C"

//------------------------------------------------------------------------------

import (
	"encoding/json"
	"os"
	"path/filepath"
	"runtime"
)

//------------------------------------------------------------------------------

func init() {
	runtime.LockOSThread()

	e1, err := os.Executable()
	if err == nil {
		e2, err := filepath.EvalSymlinks(e1)
		if err == nil {
			Path = filepath.ToSlash(filepath.Dir(e2)) + "/"
		}
	}
	if Path == "" {
		Path = filepath.ToSlash(filepath.Dir(os.Args[0])) + "/"
	}
}

//------------------------------------------------------------------------------

func Setup() error {
	// Load config file

	f, err := os.Open(Path + "init.json")
	if err != nil {
		return Error(`in configuration file "init.json" opening`, err)
	}
	d := json.NewDecoder(f)
	if err := d.Decode(&Config); err != nil {
		return Error(`in configuration file "init.json" parsing`, err)
	}

	// Initialize SDL

	if errcode := C.SDL_Init(C.SDL_INIT_EVERYTHING); errcode != 0 {
		return Error("in SDL initalization", GetSDLError())
	}

	C.SDL_StopTextInput()

	// Open the window

	err = OpenWindow(
		Config.Title,
		Config.WindowSize,
		Config.Display,
		Config.Fullscreen,
		Config.FullscreenMode,
		Config.VSync,
		Config.Debug,
	)
	if err != nil {
		return Error("in window opening", err)
	}

	return nil
}

//------------------------------------------------------------------------------
