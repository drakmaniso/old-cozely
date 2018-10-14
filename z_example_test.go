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
	play = input.Button("Play")
)

type loop struct {
	logo       pixel.PictureID
	monochrome color.Palette
	colorful   color.Palette

	playing bool
}

// Initialization //////////////////////////////////////////////////////////////

func Example() {
	defer cozely.Recover()

	l := loop{}
	l.setup()

	cozely.Configure(cozely.UpdateStep(1.0 / 3))
	err := cozely.Run(&l)
	if err != nil {
		panic(err)
	}
	// Output:
}

func (l *loop) setup() {
	pixel.SetResolution(pixel.XY{160, 100})
	l.logo = pixel.Picture("graphics/cozely")
	l.monochrome = color.PaletteFrom("graphics/cozely")
	l.colorful = color.PaletteFrom("")
}

func (l *loop) Enter() {
	pixel.SetPalette(l.monochrome)
}

func (loop) Leave() {
}

// Game Loop ///////////////////////////////////////////////////////////////////

func (l *loop) React() {
	if play.Pressed() {
		l.playing = !l.playing
		if l.playing {
			pixel.SetPalette(l.colorful)
			l.shufflecolors()
		} else {
			pixel.SetPalette(l.monochrome)
		}
	}
	if input.MenuBack.Pressed() {
		cozely.Stop(nil)
	}
}

func (l *loop) Update() {
	if l.playing {
		l.shufflecolors()
	}
}

func (l *loop) shufflecolors() {
	for i := 2; i < 14; i++ {
		g := 0.2 + 0.7*rand.Float32()
		r := 0.2 + 0.7*rand.Float32()
		b := 0.2 + 0.7*rand.Float32()
		pixel.SetColor(color.Index(i), color.SRGB{r, g, b})
	}
}

func (l *loop) Render() {
	pixel.Clear(0)

	o := pixel.Resolution().Minus(l.logo.Size()).Slash(2)
	l.logo.Paint(0, o)
}
