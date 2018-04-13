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

// Palette /////////////////////////////////////////////////////////////////////

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

// Input Bindings //////////////////////////////////////////////////////////////

var (
	quit     = input.Bool("Quit")
	next     = input.Bool("Next")
	previous = input.Bool("Previous")
	scenes   = []input.BoolID{
		input.Bool("Scene1"),
		input.Bool("Scene2"),
		input.Bool("Scene3"),
		input.Bool("Scene4"),
		input.Bool("Scene5"),
		input.Bool("Scene6"),
		input.Bool("Scene7"),
		input.Bool("Scene8"),
		input.Bool("Scene9"),
		input.Bool("Scene10"),
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

// Globals /////////////////////////////////////////////////////////////////////

var (
	canvas = pixel.Canvas(pixel.Resolution(160, 160))
	pict   = pixel.Picture("graphics/paletteswatch")
)

var mode int

// Initialization //////////////////////////////////////////////////////////////

func Example_paletteSwatch() {
	err := cozely.Run(loop{})
	if err != nil {
		cozely.ShowError(err)
	}
	//Output:
}

type loop struct{}

func (loop) Enter() error {
	bindings.Load()
	context.Activate(1)

	mode = 1
	palette.Activate()

	return nil
}

func (loop) Leave() error { return nil }

// React to User Inputs ////////////////////////////////////////////////////////

func (loop) React() error {
	if quit.JustPressed(1) {
		cozely.Stop()
	}

	if next.JustPressed(1) {
		mode++
		if mode > 5 {
			mode = 1
		}
		palette.Activate()
	}

	for i := range scenes {
		if scenes[i].JustPressed(1) {
			mode = i + 1
		}
	}

	return nil
}

func (loop) Update() error {
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
	return nil
}

// Render //////////////////////////////////////////////////////////////////////

func (loop) Render() error {
	canvas.Clear(0)

	cs := canvas.Size()

	ps := pict.Size()
	p := cs.Minus(ps).Slash(2)
	canvas.Picture(pict, 0, p)

	canvas.Text(253, pixel.Monozela10)
	canvas.Locate(-1, p.Minus(coord.CR{0, 8}))
	switch mode {
	case 1:
		canvas.Print("Custom Palette")
	case 2:
		canvas.Print("C64 Palette")
	case 3:
		canvas.Print("CPC Palette")
	case 4:
		canvas.Print("MSX Palette")
	case 5:
		canvas.Print("MSX2 Palette")
	}

	canvas.Display()
	return pixel.Err()
}
