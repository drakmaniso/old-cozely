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
		Palette []color.Color
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
		Palette []color.Color
	}{
		{"PICO8 Palette", pico8.Palette},
		{"C64 Palette", c64.Palette},
		{"CPC Palette", cpc.Palette},
		{"MSX Palette", msx.Palette},
		{"MSX2 Palette", msx2.Palette},
	}

	pixel.SetResolution(pixel.XY{160, 160})

	l.current = 0
}

func (l *loop7) Enter() {
	setPalette(l.palettes[l.current].Palette)
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

func setPalette(p []color.Color) {
	color.Clear()
	for i := 1; i < 251 && i < len(p); i++ {
		color.Set(color.Index(i), p[i-1])
	}
}

func (l *loop7) Render() {
	pixel.Clear(color.Black)

	cs := pixel.Resolution()

	ps := l.pict.Size()
	p := cs.Minus(ps).Slash(2)
	_ = p
	l.pict.Paint(p, 0)

	cur := pixel.Cursor{}
	cur.Color = 254
	cur.Position = p.Minus(pixel.XY{0, 8})
	cur.Print(l.palettes[l.current].string)
}
