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
	"github.com/cozely/cozely/input"
	"github.com/cozely/cozely/pixel"
)

// Declarations ////////////////////////////////////////////////////////////////

type loop7 struct {
	pict     pixel.PictureID
	palettes []struct {
		string
		color.Palette
	}
	current int
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

func (l *loop7) setup() {
	l.pict = pixel.Picture("graphics/paletteswatch")

	l.palettes = []struct {
		string
		color.Palette
	}{
		{"Default Palette", pixel.DefaultPalette},
		{"C64 Palette", c64.Palette},
		{"CPC Palette", cpc.Palette},
		{"MSX Palette", msx.Palette},
		{"MSX2 Palette", msx2.Palette},
		{"PICO8 Palette", pico8.Palette},
	}

	pixel.SetResolution(pixel.XY{160, 160})

	l.current = 0
}

func (loop7) Enter() {
}

func (loop7) Leave() {
}

// Game Loop ///////////////////////////////////////////////////////////////////

func (l *loop7) React() {
	if input.MenuBack.Pushed() {
		cozely.Stop(nil)
	}

	if input.MenuRight.Pushed() {
		l.current++
		if l.current >= len(l.palettes) {
			l.current = len(l.palettes) - 1
		}
		setPalette(l.palettes[l.current].Palette)
	}
	if input.MenuLeft.Pushed() {
		l.current--
		if l.current < 0 {
			l.current = 0
		}
		setPalette(l.palettes[l.current].Palette)
	}
}

func (loop7) Update() {
}

func setPalette(p color.Palette) {
	for i := 1; i < 256; i++ {
		switch {
		case i <= len(p.Colors):
			pixel.SetColor(color.Index(i), p.Colors[i-1])
		case i == 255:
			pixel.SetColor(color.Index(i), color.LRGBA{1, 1, 1, 1})
		default:
			pixel.SetColor(color.Index(i), color.LRGBA{0, 0, 0, 1})
		}
	}
}

func (l *loop7) Render() {
	pixel.Clear(0)

	cs := pixel.Resolution()

	ps := l.pict.Size()
	p := cs.Minus(ps).Slash(2)
	_ = p
	l.pict.Paint(p, 0)

	cur := pixel.Cursor{}
	cur.Color = 15
	cur.Position = p.Minus(pixel.XY{0, 8})
	cur.Print(l.palettes[l.current].string)
}
