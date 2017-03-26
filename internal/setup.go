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
	"log"
	"os"
	"path/filepath"
	"runtime"
)

//------------------------------------------------------------------------------

// Path of the executable.
var Path = filepath.Dir(os.Args[0]) + "/"

var Config = struct {
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
	runtime.LockOSThread()
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
		Config.Resolution,
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

var logger = log.New(os.Stderr, "glam: ", log.Ltime)

// Log logs a formated message.
func Log(format string, v ...interface{}) {
	logger.Printf(format, v...)
}

// DebugLog logs a formated message if Debug mode is enabled.
func DebugLog(format string, v ...interface{}) {
	if Config.Debug {
		logger.Printf(format, v...)
	}
}

//------------------------------------------------------------------------------

// Error returns nil if err is nil, or a wrapped error otherwise.
func Error(source string, err error) error {
	if err == nil {
		return nil
	}
	return wrappedError{source, err}
}

type wrappedError struct {
	source string
	err    error
}

func (e wrappedError) Error() string {
	msg := e.source + ":\n\t"
	a := e.err
	for b, ok := a.(wrappedError); ok; {
		msg += b.source + ":\n\t"
		a = b.err
		b, ok = a.(wrappedError)
	}
	return msg + a.Error()
}

//------------------------------------------------------------------------------
