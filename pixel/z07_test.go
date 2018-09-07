// Copyright (c) 2018-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package pixel_test

import (
	"testing"

	"github.com/cozely/cozely"
	"github.com/cozely/cozely/color"
	"github.com/cozely/cozely/input"
	"github.com/cozely/cozely/pixel"
	"github.com/cozely/cozely/pixel/palettes/c64"
	"github.com/cozely/cozely/pixel/palettes/cpc"
	"github.com/cozely/cozely/pixel/palettes/msx"
	"github.com/cozely/cozely/pixel/palettes/msx2"
	"github.com/cozely/cozely/pixel/palettes/pico8"
)

// Declarations ////////////////////////////////////////////////////////////////

type loop7 struct {
	pict                         pixel.PictureID
	palette, c64, cpc, msx, msx2 pixel.PaletteID
	pico8                        pixel.PaletteID
	orange, cyan, black          pixel.Color
	mode                         int
}

// Initialization //////////////////////////////////////////////////////////////

func TestTest7(t *testing.T) {
	do(func() {
		defer cozely.Recover()

		l := loop7{}
		l.setup()

		input.Load(bindings)

		err := cozely.Run(&l)
		if err != nil {
			panic(err)
		}
	})
}

func (a *loop7) setup() {
	a.c64 = pixel.PaletteColors(c64.Colors)
	a.cpc = pixel.PaletteColors(cpc.Colors)
	a.msx = pixel.PaletteColors(msx.Colors)
	a.msx2 = pixel.PaletteColors(msx2.Colors)
	a.pico8 = pixel.PaletteColors(pico8.Colors)

	a.pict = pixel.Picture("graphics/paletteswatch")

	pixel.SetResolution(160, 160)

	a.mode = 0
}

func (a *loop7) Enter() {
	a.mode = 1
	a.palette = pixel.PaletteColors([256]color.Color{
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
	})
	a.orange = a.palette.Set(255, color.SRGB{1, 0.6, 0})
	a.cyan = a.palette.Set(254, color.SRGB{0, 0.9, 1})
	a.black = a.palette.Set(253, color.SRGB{0, 0, 0})
}

func (a *loop7) Leave() {
}

// Game Loop ///////////////////////////////////////////////////////////////////

func (a *loop7) React() {
	if quit.Started(1) {
		cozely.Stop(nil)
	}

	if next.Started(1) {
		a.mode++
		if a.mode > 9 {
			a.mode = 0
		}
		a.palette.Use()
	}

	for i := range scenes {
		if scenes[i].Started(1) {
			a.mode = i
		}
	}
}

func (a *loop7) Update() {
	switch a.mode {
	case 1:
		pixel.DefaultPalette.Use()
	case 2:
		a.c64.Use()
	case 3:
		a.cpc.Use()
	case 4:
		a.msx.Use()
	case 5:
		a.msx2.Use()
	case 6:
		a.pico8.Use()
	case 7:
		a.palette.Use()
	case 8:
		a.palette.Use()
	case 9:
		a.palette.Use()
	case 0:
		a.palette.Use()
	}
}

func (a *loop7) Render() {
	pixel.Clear(0)

	cs := pixel.Resolution()

	ps := a.pict.Size()
	p := cs.Minus(ps).Slash(2)
	_ = p
	pixel.Paint(a.pict, p)

	pixel.Text(15, pixel.Monozela10)
	pixel.Locate(p.Minus(pixel.XY{0, 8}))
	switch a.mode {
	case 1:
		pixel.Print("Default Palette")
	case 2:
		pixel.Print("C64 Palette")
	case 3:
		pixel.Print("CPC Palette")
	case 4:
		pixel.Print("MSX Palette")
	case 5:
		pixel.Print("MSX2 Palette")
	case 6:
		pixel.Print("PICO-8 Palette")
	case 7:
		pixel.Print("Palette")
	case 8:
		pixel.Print("Palette")
	case 9:
		pixel.Print("Palette")
	case 0:
		pixel.Print("Custom Palette")
	}
}
