package cozely_test

import (
	"math/rand"

	"github.com/cozely/cozely"
	"github.com/cozely/cozely/color"
	"github.com/cozely/cozely/input"
	"github.com/cozely/cozely/pixel"
)

// Declarations ////////////////////////////////////////////////////////////////

var (
	quit  = input.Digital("Quit")
	start = input.Digital("Start")
)

var bindings = input.Bindings{
	"Default": {
		"Start": {"Space", "Mouse Left", "Button A"},
		"Quit":  {"Escape", "Button Back"},
	},
}

var (
	canv       = pixel.Canvas(pixel.Resolution(160, 100))
	logo       = pixel.Picture("graphics/cozely")
	monochrome = color.PaletteFrom("graphics/cozely")
	colorful   = color.Palette()
)

var started = false

var reverse = false

type loop2 struct{}

// Initialization //////////////////////////////////////////////////////////////

func Example_interactive() {
	defer cozely.Recover()

	input.Load(bindings)
	cozely.Configure(cozely.UpdateStep(1.0 / 3))
	err := cozely.Run(loop2{})
	if err != nil {
		panic(err)
	}
	// Output:
}

func (loop2) Enter() {
	monochrome.Activate()
	shufflecolors()
}

func (loop2) Leave() {
}

// Game Loop ///////////////////////////////////////////////////////////////////

func (loop2) React() {
	if start.Started(input.Any) {
		started = !started
	}
	if quit.Started(input.Any) {
		cozely.Stop(nil)
	}
}

func (loop2) Update() {
	if started {
		shufflecolors()
	}
}

func shufflecolors() {
	dark := [12]bool{
		true, false, true, false, true, false,
		false, true, false, true, false, true,
	}
	reverse = !reverse
	for i := 2; i < 14; i++ {
		g := 0.2 + 0.8*rand.Float32()
		r := 0.2 + 0.8*rand.Float32()
		b := 0.2 + 0.8*rand.Float32()
		if dark[i-2] != reverse {
			colorful.Set(uint8(i), color.SRGB{r, g, b})
		} else {
			colorful.Set(uint8(i), color.LRGB{r, g, b})
		}
	}
}

func (loop2) Render() {
	canv.Clear(0)

	if started {
		colorful.Activate()
	} else {
		monochrome.Activate()
	}

	o := canv.Size().Minus(logo.Size()).Slash(2)
	canv.Picture(logo, 0, o)

	canv.Display()
}
