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
	quit = input.Digital("Quit")
	play = input.Digital("Play")
)

var bindings = input.Bindings{
	"Default": {
		"Play": {"Space", "Mouse Left", "Button A"},
		"Quit": {"Escape", "Button Back"},
	},
}

var (
	canvas2    = pixel.Canvas(pixel.Resolution(160, 100))
	logo       = pixel.Picture("graphics/cozely")
	monochrome = color.PaletteFrom("graphics/cozely")
	colorful   = color.Palette()
)

type loop2 struct {
	playing bool
}

// Initialization //////////////////////////////////////////////////////////////

func Example_interactive() {
	defer cozely.Recover()

	input.Load(bindings)
	cozely.Configure(cozely.UpdateStep(1.0 / 3))
	err := cozely.Run(&loop2{})
	if err != nil {
		panic(err)
	}
	// Output:
}

func (l *loop2) Enter() {
	monochrome.Activate()
	l.shufflecolors()
}

func (loop2) Leave() {
}

// Game Loop ///////////////////////////////////////////////////////////////////

func (l *loop2) React() {
	if play.Started(input.Any) {
		l.playing = !l.playing
	}
	if quit.Started(input.Any) {
		cozely.Stop(nil)
	}
}

func (l *loop2) Update() {
	if l.playing {
		l.shufflecolors()
	}
}

func (l *loop2) shufflecolors() {
	for i := 2; i < 14; i++ {
		g := 0.2 + 0.7*rand.Float32()
		r := 0.2 + 0.7*rand.Float32()
		b := 0.2 + 0.7*rand.Float32()
		colorful.Set(uint8(i), color.SRGB{r, g, b})
	}
}

func (l *loop2) Render() {
	canvas2.Clear(0)

	if l.playing {
		colorful.Activate()
	} else {
		monochrome.Activate()
	}

	o := canvas2.Size().Minus(logo.Size()).Slash(2)
	canvas2.Picture(logo, o)

	canvas2.Display()
}
