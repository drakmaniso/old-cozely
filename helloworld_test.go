package glam_test

import (
	"github.com/drakmaniso/glam"
	"github.com/drakmaniso/glam/palette"
	"github.com/drakmaniso/glam/pixel"
)

// Initialization //////////////////////////////////////////////////////////////

var (
	canvas = pixel.Canvas()
	cursor = pixel.Cursor{Canvas: canvas}
)

func Example_helloWorld() {
	glam.Run(loop{})
	// Output:
}

// Game Loop ///////////////////////////////////////////////////////////////////

type loop struct{}

func (loop) Enter() error {
	palette.Load("C64")

	return nil
}

func (loop) Leave() error { return nil }

func (loop) React() error { return nil }

func (loop) Update() error { return nil }

func (loop) Render() error {
	canvas.Clear(0)

	cursor.Locate(8, 12)
	cursor.Color = 3
	cursor.Print("Hello, World!")

	canvas.Display()
	return nil
}
