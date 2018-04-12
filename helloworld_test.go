package cozely_test

import (
	"github.com/cozely/cozely"
	"github.com/cozely/cozely/palette"
	"github.com/cozely/cozely/pixel"
)

// Initialization //////////////////////////////////////////////////////////////

var canvas = pixel.Canvas()

func Example_helloWorld() {
	cozely.Run(loop{})
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

	canvas.Text(3, pixel.Monozela10)
	canvas.Locate(8, 12, 0)
	canvas.Print("Hello, World!")

	canvas.Display()
	return nil
}