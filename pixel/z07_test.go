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
)

// Declarations ////////////////////////////////////////////////////////////////

type loop7 struct {
	// context                      input.ContextID
	pict                                             pixel.PictureID
	palette, c64, cpc, msx2                          pixel.PaletteID
	msx, msxIdeal, msxCV, msxCheap, msxTrim, msxLazy pixel.PaletteID
	orange, cyan, black                              pixel.Color
	mode                                             int
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
	a.msx2 = pixel.PaletteColors(msx2.Colors)
	a.msxIdeal = pixel.PaletteColors(msx.ColorsIdealized)
	a.msx = pixel.PaletteColors(msx.Colors)
	a.msxCV = pixel.PaletteColors(msx.ColorsCVtoRGB)
	a.msxCheap = pixel.PaletteColors(msx.ColorsCheapRGB)
	a.msxTrim = pixel.PaletteColors(msx.ColorsCheapRGBTrim)
	a.msxLazy = pixel.PaletteColors(msx.ColorsLazyRGB)

	a.pict = pixel.Picture("graphics/paletteswatch")

	// a.context = input.Context("Default", quit, next, previous,
	// 	scenes[1], scenes[2], scenes[3], scenes[4], scenes[5],
	// 	scenes[6], scenes[7], scenes[8], scenes[9], scenes[0])

	pixel.SetResolution(160, 160)

	a.mode = 0
}

func (a *loop7) Enter() {
	// a.context.Activate(1)

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

	a.palette.Use()
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
		if a.mode > 10 {
			a.mode = 1
		}
		a.palette.Use()
	}

	for i := range scenes {
		if scenes[i].Started(1) {
			a.mode = i + 1
		}
	}
}

func (a *loop7) Update() {
	switch a.mode {
	case 1:
		a.palette.Use()
	case 2:
		a.c64.Use()
	case 3:
		a.cpc.Use()
	case 4:
		a.msx2.Use()
	case 5:
		a.msxIdeal.Use()
	case 6:
		a.msx.Use()
	case 7:
		a.msxCV.Use()
	case 8:
		a.msxCheap.Use()
	case 9:
		a.msxTrim.Use()
	case 10:
		a.msxLazy.Use()
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
		pixel.Print("Custom Palette")
	case 2:
		pixel.Print("C64 Palette")
	case 3:
		pixel.Print("CPC Palette")
	case 4:
		pixel.Print("MSX2 Palette")
	case 5:
		pixel.Print("MSX Palette (Idealized)")
	case 6:
		pixel.Print("MSX Palette (ITU-R BT601)")
	case 7:
		pixel.Print("MSX Palette (CV to RGB)")
	case 8:
		pixel.Print("MSX Palette (Cheap RGB)")
	case 9:
		pixel.Print("MSX Palette (Cheap RGB + trimpots)")
	case 10:
		pixel.Print("MSX Palette (Lazy RGB)")
	}
}
