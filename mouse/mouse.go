// Package mouse provides support for the mouse.
package mouse

// #cgo windows LDFLAGS: -lSDL2
// #cgo linux freebsd darwin pkg-config: sdl2
// #include "../engine/engine.h"
import "C"

import (
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
