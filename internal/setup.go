// Copyright (c) 2013-2016 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package internal

//------------------------------------------------------------------------------

/*
#cgo windows LDFLAGS: -lSDL2
#cgo linux freebsd darwin pkg-config: sdl2

#include "sdl.h"
*/
import "C"

import (
	"encoding/json"
	"os"
	"path/filepath"
	"runtime"
)

//------------------------------------------------------------------------------

// Path of the executable (uses slash separators, and ends with one).
var Path string

var Config = struct {
	Title          string
	WindowSize     [2]int32
	ScreenSize     [2]int32
	PixelSize      int
	Display        int
	Fullscreen     bool
	FullscreenMode string
	VSync          bool
	Debug          bool
}{
	Title:          "Carol",
	WindowSize:     [2]int32{1280, 720},
	ScreenSize:     [2]int32{64, 64},
	PixelSize:      8,
	Display:        0,
	Fullscreen:     false,
	FullscreenMode: "Desktop",
	VSync:          true,
	Debug:          false,
}

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
		return Error(`opening configuration file "init.json"`, err)
	}
	d := json.NewDecoder(f)
	if err := d.Decode(&Config); err != nil {
		return Error("decoding configuration file", err)
	}

	// Initialize SDL

	if errcode := C.SDL_Init(C.SDL_INIT_EVERYTHING); errcode != 0 {
		return Error("initializing SDL", GetSDLError())
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
		return Error("opening window", err)
	}

	return nil
}

//------------------------------------------------------------------------------
