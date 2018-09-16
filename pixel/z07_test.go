// Copyright (c) 2018-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package pixel_test

import (
	"testing"

	"github.com/cozely/cozely"
	"github.com/cozely/cozely/color"
	"github.com/cozely/cozely/color/c64"
	"github.com/cozely/cozely/color/cpc"
	"github.com/cozely/cozely/color/msx"
	"github.com/cozely/cozely/color/msx2"
	"github.com/cozely/cozely/color/pico8"
	"github.com/cozely/cozely/pixel"
)

// Declarations ////////////////////////////////////////////////////////////////

type loop7 struct {
	pict                pixel.PictureID
	palette             color.Palette
	orange, cyan, black color.Index
	mode                int
}

// Initialization //////////////////////////////////////////////////////////////

func TestTest7(t *testing.T) {
	do(func() {
		defer cozely.Recover()

		l := loop7{}
		l.setup()

		err := cozely.Run(&l)
		if err != nil {
			panic(err)
		}
	})
}

func (a *loop7) setup() {
	a.pict = pixel.Picture("graphics/paletteswatch")

	pixel.SetResolution(pixel.XY{160, 160})

	a.mode = 0
}

func (a *loop7) Enter() {
	a.mode = 1
	a.palette = color.Palette{
		ByName: map[string]color.Index{},
		Colors: []color.LRGBA{
			{0.1, 0.1, 0.1, 0},
			{0.2, 0.2, 0.2, 0},
			{0.3, 0.3, 0.3, 0},
			{0.4, 0.4, 0.4, 0},
			{0.5, 0.5, 0.5, 0},
			{0.6, 0.6, 0.6, 0},
			{0.7, 0.7, 0.7, 0},
			{0.8, 0.8, 0.8, 0},
			{0.9, 0.9, 0.9, 0},
			{1.0, 1.0, 1.0, 0},
		},
	}
}

func (a *loop7) Leave() {
}

// Game Loop ///////////////////////////////////////////////////////////////////

func (a *loop7) React() {
	if quit.Pushed() {
		cozely.Stop(nil)
	}

	if next.Pushed() {
		a.mode++
		if a.mode > 9 {
			a.mode = 0
		}
	}

	for i := range scenes {
		if scenes[i].Pushed() {
			a.mode = i
		}
	}
}

func (a *loop7) Update() {
	switch a.mode {
	case 1:
		pixel.SetPalette(pixel.DefaultPalette)
	case 2:
		pixel.SetPalette(c64.Palette)
	case 3:
		pixel.SetPalette(cpc.Palette)
	case 4:
		pixel.SetPalette(msx.Palette)
	case 5:
		pixel.SetPalette(msx2.Palette)
	case 6:
		pixel.SetPalette(pico8.Palette)
	case 7:
		pixel.SetPalette(a.palette)
	case 8:
		pixel.SetPalette(a.palette)
	case 9:
		pixel.SetPalette(a.palette)
	case 0:
		pixel.SetPalette(a.palette)
		a.orange = pixel.SetColor(255, color.SRGB{1, 0.6, 0})
		a.cyan = pixel.SetColor(254, color.SRGB{0, 0.9, 1})
		a.black = pixel.SetColor(253, color.SRGB{0, 0, 0})
	}
}

func (a *loop7) Render() {
	pixel.Clear(0)

	cs := pixel.Resolution()

	ps := a.pict.Size()
	p := cs.Minus(ps).Slash(2)
	_ = p
	a.pict.Paint(0, p)

	cur := pixel.Cursor{}
	cur.Style(15, pixel.Monozela10)
	cur.Locate(0, p.Minus(pixel.XY{0, 8}))
	switch a.mode {
	case 1:
		cur.Print("Default Palette")
	case 2:
		cur.Print("C64 Palette")
	case 3:
		cur.Print("CPC Palette")
	case 4:
		cur.Print("MSX Palette")
	case 5:
		cur.Print("MSX2 Palette")
	case 6:
		cur.Print("PICO-8 Palette")
	case 7:
		cur.Print("Palette")
	case 8:
		cur.Print("Palette")
	case 9:
		cur.Print("Palette")
	case 0:
		cur.Print("Custom Palette")
	}
}
