package cozely_test

import (
	"math/rand"

	"github.com/cozely/cozely"
	"github.com/cozely/cozely/color"
	"github.com/cozely/cozely/input"
	"github.com/cozely/cozely/pixel"
)

// Input Bindings //////////////////////////////////////////////////////////////

var (
	quit    = input.Bool("Quit")
	start   = input.Bool("Start")
	context = input.Context("Default", start, quit)
)

var bindings = input.Bindings{
	"Default": {
		"Start": {"Space", "Mouse Left"},
		"Quit":  {"Escape"},
	},
}

// Initialization //////////////////////////////////////////////////////////////

var (
	canv = pixel.Canvas(pixel.Resolution(160, 100))
	logo = pixel.Picture("graphics/cozely")
	pal1 = color.PaletteFrom("graphics/cozely")
	pal2 = color.Palette()
)

var started = false

var inv = false

func Example_interactive() {
	cozely.Configure(cozely.UpdateStep(1.0 / 3))
	cozely.Run(interactive{})
	// Output:
}

type interactive struct{}

func (interactive) Enter() error {
	bindings.Load()
	context.Activate(1)
	pal1.Activate()
	shufflecolors()
	return nil
}

func (interactive) Leave() error {
	return nil
}

// Game Loop ///////////////////////////////////////////////////////////////////

func (interactive) React() error {
	if start.JustPressed(1) {
		started = !started
	}
	if quit.JustPressed(1) {
		cozely.Stop()
	}
	return nil
}

func (interactive) Update() error {
	if started {
		shufflecolors()
	}
	return nil
}

func shufflecolors() {
	dark := [12]bool{
		true, false, true, false, true, false,
		false, true, false, true, false, true,
	}
	inv = !inv
	for i := 2; i < 14; i++ {
		r := .2+.8*rand.Float32()
		g := .2+.8*rand.Float32()
		b := .2+.8*rand.Float32()
		if dark[i-2] != inv {
			pal2.Set(uint8(i), color.SRGB{r, g, b})
		} else {
			pal2.Set(uint8(i), color.LRGB{r, g, b})
		}
	}
}

func (interactive) Render() error {
	canv.Clear(0)

	if started {
		pal2.Activate()
	} else {
		pal1.Activate()
	}

	o := canv.Size().Minus(logo.Size()).Slash(2)
	canv.Picture(logo, 0, o)

	canv.Display()
	return nil
}
