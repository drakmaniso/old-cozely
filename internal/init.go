// Copyright (c) 2013-2016 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package internal

//------------------------------------------------------------------------------

/*
#cgo windows LDFLAGS: -lSDL2
#cgo linux freebsd darwin pkg-config: sdl2

#include "SDL.h"
*/
import "C"

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
)

//------------------------------------------------------------------------------

// Path of the executable.
var Path = filepath.Dir(os.Args[0]) + "/"

var config = struct {
	Title          string
	Resolution     [2]int32
	Display        int
	Fullscreen     bool
	FullscreenMode string
	VSync          bool
	Debug          bool
}{
	Title:          "Glam",
	Resolution:     [2]int32{1280, 720},
	Display:        0,
	Fullscreen:     false,
	FullscreenMode: "Desktop",
	VSync:          true,
	Debug:          false,
}

//------------------------------------------------------------------------------

func init() {
	loadConfig()

	if config.Debug {
		Debug = true
	}

	// Log("Path = \"%s\"", Path)

	runtime.LockOSThread()

	if errcode := C.SDL_Init(C.SDL_INIT_EVERYTHING); errcode != 0 {
		InitError = fmt.Errorf("impossible to initialize SDL: %s", GetSDLError())
		Log("%s", InitError)
		return
	}

	C.SDL_StopTextInput()

	err := OpenWindow(
		config.Title,
		config.Resolution,
		config.Display,
		config.Fullscreen,
		config.FullscreenMode,
		config.VSync,
		config.Debug,
	)
	if err != nil {
		InitError = err
		Log("%s", InitError)
		return
	}
}

//------------------------------------------------------------------------------

func loadConfig() {
	f, err := os.Open(Path + "init.json")
	if err != nil {
		Log(`No configuration file ("init.json") found: %s`, err)
		return
	}
	d := json.NewDecoder(f)
	if err := d.Decode(&config); err != nil {
		InitError = fmt.Errorf(`impossible to decode configuration file "init.json": %s`, err)
		Log("%s", InitError)
		return
	}
}

//------------------------------------------------------------------------------

var logger = log.New(os.Stderr, "glam: ", log.Ltime)

// Log logs a formated message if Debug mode is enabled.
func Log(format string, v ...interface{}) {
	if Debug {
		logger.Printf(format, v...)
	}
}

//------------------------------------------------------------------------------
