package glam_test

import (
	"github.com/drakmaniso/glam/plane"
	"github.com/drakmaniso/glam"
	"github.com/drakmaniso/glam/input"
	"github.com/drakmaniso/glam/palette"
	"github.com/drakmaniso/glam/pixel"
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
	glam.Run(interactive{})
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
		glam.Stop()
	}
	return nil
}

func (interactive) Update() error {
	return nil
}

func (interactive) Render() error {
	screen.Clear(0)

	b := plane.Pixel{16, 12}
	screen.Box(4, 9, 4, 0, b, screen.Size().Minus(b))

	screen.Display()
	return nil
}
