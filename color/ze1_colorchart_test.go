// Copyright (c) 2018-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package color_test

import (
	"github.com/cozely/cozely"
	"github.com/cozely/cozely/color"
	"github.com/cozely/cozely/color/palettes/c64"
	"github.com/cozely/cozely/color/palettes/cpc"
	"github.com/cozely/cozely/color/palettes/msx"
	"github.com/cozely/cozely/color/palettes/msx2"
	"github.com/cozely/cozely/coord"
	"github.com/cozely/cozely/input"
	"github.com/cozely/cozely/pixel"
)

// Declarations ////////////////////////////////////////////////////////////////

// Palette

var (
	palette = color.Palette(
		color.LRGB{0.1, 0.1, 0.1},
		color.LRGB{0.2, 0.2, 0.2},
		color.LRGB{0.3, 0.3, 0.3},
		color.LRGB{0.4, 0.4, 0.4},
		color.LRGB{0.5, 0.5, 0.5},
		color.LRGB{0.6, 0.6, 0.6},
		color.LRGB{0.7, 0.7, 0.7},
		color.LRGB{0.8, 0.8, 0.8},
		color.LRGB{0.9, 0.9, 0.9},
		color.LRGB{1.0, 1.0, 1.0},
	)
	orange = palette.Entry(color.SRGB{1, 0.6, 0})
	cyan   = palette.Entry(color.SRGB{0, 0.9, 1})
	black  = palette.Entry(color.SRGB{0, 0, 0})
)

// Input Bindings

var (
	quit     = input.Digital("Quit")
	next     = input.Digital("Next")
	previous = input.Digital("Previous")
	scenes   = []input.DigitalID{
		input.Digital("Scene1"),
		input.Digital("Scene2"),
		input.Digital("Scene3"),
		input.Digital("Scene4"),
		input.Digital("Scene5"),
		input.Digital("Scene6"),
		input.Digital("Scene7"),
		input.Digital("Scene8"),
		input.Digital("Scene9"),
		input.Digital("Scene10"),
	}
)

var context = input.Context("Default", quit, next, previous,
	scenes[1], scenes[2], scenes[3], scenes[4], scenes[5],
	scenes[6], scenes[7], scenes[8], scenes[9], scenes[0])

var bindings = input.Bindings{
	"Default": {
		"Quit":     {"Escape"},
		"Next":     {"Mouse Left", "Space"},
		"Previous": {"Mouse Right", "U"},
		"Scene1":   {"1"},
		"Scene2":   {"2"},
		"Scene3":   {"3"},
		"Scene4":   {"4"},
		"Scene5":   {"5"},
		"Scene6":   {"6"},
		"Scene7":   {"7"},
		"Scene8":   {"8"},
		"Scene9":   {"9"},
		"Scene10":  {"10"},
	},
}

// Globals

var (
	canvas = pixel.Canvas(pixel.Resolution(160, 160))
	scene  = pixel.Scene()
	pict   = pixel.Picture("graphics/paletteswatch")
)

var mode int

// Initialization //////////////////////////////////////////////////////////////

func Example_colorChart() {
	defer cozely.Recover()

	input.Load(bindings)
	err := cozely.Run(loop{})
	if err != nil {
		panic(err)
	}
	//Output:
}

type loop struct{}

func (loop) Enter() {
	context.Activate(1)

	mode = 1
	palette.Activate()
}

func (loop) Leave() {
}

// Game Loop ///////////////////////////////////////////////////////////////////

func (loop) React() {
	if quit.Started(1) {
		cozely.Stop(nil)
	}

	if next.Started(1) {
		mode++
		if mode > 5 {
			mode = 1
		}
		palette.Activate()
	}

	for i := range scenes {
		if scenes[i].Started(1) {
			mode = i + 1
		}
	}
}

func (loop) Update() {
	switch mode {
	case 1:
		palette.Activate()
	case 2:
		c64.Palette.Activate()
	case 3:
		cpc.Palette.Activate()
	case 4:
		msx.Palette.Activate()
	case 5:
		msx2.Palette.Activate()
	}
}

func (loop) Render() {
	canvas.Clear(0)

	cs := canvas.Size()

	ps := pict.Size()
	p := cs.Minus(ps).Slash(2)
	scene.Picture(pict, p)

	scene.Text(253, pixel.Monozela10)
	scene.Locate(p.Minus(coord.CR{0, 8}))
	switch mode {
	case 1:
		scene.Print("Custom Palette")
	case 2:
		scene.Print("C64 Palette")
	case 3:
		scene.Print("CPC Palette")
	case 4:
		scene.Print("MSX Palette")
	case 5:
		scene.Print("MSX2 Palette")
	}

	canvas.Display(scene)
}
