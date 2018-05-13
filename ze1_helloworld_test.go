package cozely_test

import (
	"github.com/cozely/cozely"
	"github.com/cozely/cozely/color"
	"github.com/cozely/cozely/coord"
	"github.com/cozely/cozely/pixel"
)

// Declarations ////////////////////////////////////////////////////////////////

var (
	canvas = pixel.Canvas(pixel.Resolution(320, 200))
	scene  = pixel.Scene()

	palette = color.Palette()
	fg      = palette.Entry(color.SRGB{0.75, 0.98, 0.52})
	bg      = palette.Entry(color.SRGB{0.06, 0.18, 0.12})
)

type loop struct{}

// Initialization //////////////////////////////////////////////////////////////

func Example_helloWorld() {
	defer cozely.Recover()

	err := cozely.Run(loop{})
	if err != nil {
		panic(err)
	}
	// Output:
}

func (loop) Enter() {
	palette.Activate()
}

func (loop) Leave() {
}

// Game Loop ///////////////////////////////////////////////////////////////////

func (loop) React() {
}

func (loop) Update() {
}

func (loop) Render() {
	canvas.Clear(bg)
	scene.Clear()

	scene.Text(fg, pixel.Monozela10)
	scene.Locate(coord.CR{16, 32})
	scene.Print("Hello, World!")

	canvas.Display(scene)
}
