package cozely_test

import (
	"github.com/cozely/cozely"
	"github.com/cozely/cozely/color"
	"github.com/cozely/cozely/pixel"
)

// Declarations ////////////////////////////////////////////////////////////////

var (
	palette = color.Palette()
	fg      = palette.Entry(color.SRGB{0.75, 0.98, 0.52})
	bg      = palette.Entry(color.SRGB{0.06, 0.18, 0.12})
)

type loop struct{}

// Initialization //////////////////////////////////////////////////////////////

func Example_helloWorld() {
	defer cozely.Recover()

	pixel.SetResolution(320, 200)

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
	pixel.Clear(bg)

	pixel.Text(fg, pixel.Monozela10)
	pixel.Locate(pixel.XY{16, 32})
	pixel.Print("Hello, World!")
}
