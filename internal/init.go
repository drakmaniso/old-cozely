// Copyright (c) 2013-2016 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package internal

/*
#cgo windows LDFLAGS: -lSDL2
#cgo linux freebsd darwin pkg-config: sdl2

#include "SDL.h"
#include "glad.h"

int InitOpenGL(int debug);
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

var Path = filepath.Dir(os.Args[0])

var config = struct {
	Title          string
	Resolution     [2]int
	Display        int
	Fullscreen     bool
	FullscreenMode string
	VSync          bool
	Debug          bool
}{
	Title:          "Glam",
	Resolution:     [2]int{1280, 720},
	Display:        0,
	Fullscreen:     false,
	FullscreenMode: "Desktop",
	VSync:          true,
	Debug:          false,
}

//------------------------------------------------------------------------------

func init() {
	log.SetFlags(log.Ltime | log.Lshortfile)

	//TODO: log.Printf("Path = \"%s\"", Path)

	loadConfig()

	runtime.LockOSThread()

	if errcode := C.SDL_Init(C.SDL_INIT_EVERYTHING); errcode != 0 {
		log.Panic(GetSDLError())
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
		log.Panic(err)
	}

	var d C.int
	if config.Debug {
		d = 1
	}
	if C.InitOpenGL(d) != 0 {
		log.Panic("failed to load OpenGL")
	}
}

func loadConfig() {
	f, err := os.Open(Path + "/init.json")
	if err != nil {
		log.Print(err)
		return
	}
	d := json.NewDecoder(f)
	err = d.Decode(&config)
	if err != nil {
		log.Panic(err)
	}
	//TODO: log.Printf("config = %+v", config)
}

//------------------------------------------------------------------------------
