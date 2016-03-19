package internal

// #include "sdl.h"
// #include "render.h"
import "C"

import (
	"log"
)

func initRender() {
	if C.InitOpenGL() < 0 {
		log.Panic("Failed to load OpenGL")
	}
}

func Render() {
	C.Render(Window.window)
}
