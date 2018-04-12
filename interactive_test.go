package cozely_test

import (
	"github.com/cozely/cozely"
	"github.com/cozely/cozely/coord"
	"github.com/cozely/cozely/input"
	"github.com/cozely/cozely/palette"
	"github.com/cozely/cozely/pixel"
)

// Input Bindings //////////////////////////////////////////////////////////////

var (
	quit    = input.Bool("Quit")
	context = input.Context("Default", quit)
)

var bindings = input.Bindings{
	"Default": {
		"Quit": {"Escape"},
	},
}

// Initialization //////////////////////////////////////////////////////////////

var (
	screen = pixel.Canvas(pixel.Zoom(3))
)

func Example_interactive() {
	cozely.Run(interactive{})
	// Output:
}

// Game Loop ///////////////////////////////////////////////////////////////////

type interactive struct{}

func (interactive) Enter() error {
	input.Load(bindings)
	context.Activate(1)

	palette.Load("C64")

	return nil
}

func (interactive) Leave() error {
	return nil
}

func (interactive) React() error {
	if quit.JustPressed(1) {
		cozely.Stop()
	}
	return nil
}

func (interactive) Update() error {
	return nil
}

func (interactive) Render() error {
	screen.Clear(0)

	margin := coord.CR{16, 12}
	screen.Box(4, 9, 4, 0, margin, screen.Size().Minus(margin))

	screen.Display()
	return nil
}
