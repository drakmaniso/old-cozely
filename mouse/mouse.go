// Package mouse provides support for the mouse.
package mouse

// #cgo windows LDFLAGS: -lSDL2
// #cgo linux freebsd darwin pkg-config: sdl2
// #include "../engine/engine.h"
import "C"

import (
	"errors"
	"log"

	"github.com/drakmaniso/glam/geom"
)

// Position returns the current mouse position, relative to the game window.
// Updated at the start of each game loop iteration.
func Position() geom.IVec2 {
	return geom.IVec2{X: int32(C.mouseX), Y: int32(C.mouseY)}
}

var delta geom.IVec2

// Delta returns the mouse position relative to the last call of Delta.
func Delta() geom.IVec2 {
	result := delta
	delta.X, delta.Y = 0, 0
	return result
}

// AddToDelta is used internally by the engine.
func AddToDelta(rel geom.IVec2) {
	delta = delta.Plus(rel)
}

// SetRelativeMode enables or disables the relative mode, where the mouse is
// hidden and mouse motions are continuously reported.
func SetRelativeMode(enabled bool) error {
	var err error
	var m C.SDL_bool
	if enabled {
		m = 1
	}
	if C.SDL_SetRelativeMouseMode(m) != 0 {
		err = getError()
		log.Print(err)
	}
	return err
}

// GetRelativeMode returns true if the relative mode is enabled.
func GetRelativeMode() bool {
	return C.SDL_GetRelativeMouseMode() == C.SDL_TRUE
}

func getError() error {
	if s := C.SDL_GetError(); s != nil {
		return errors.New(C.GoString(s))
	}
	return nil
}
