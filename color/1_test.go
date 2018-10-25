package color_test

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
	"github.com/cozely/cozely/resource"
)

// Declarations ////////////////////////////////////////////////////////////////

type loop7 struct {
	pict     pixel.PictureID
	palettes []struct {
		string
		*color.Palette
	}
	current int
}

// Initialization //////////////////////////////////////////////////////////////

func TestTest7(t *testing.T) {
	do(func() {
		defer cozely.Recover()

		err := resource.Path("testdata/")
		if err != nil {
			t.Error(err)
			return
		}
		pixel.SetResolution(pixel.XY{160, 160})

		err = cozely.Run(&loop7{})
		if err != nil {
			panic(err)
		}
	})
}

func (l *loop7) Enter() {
	l.palettes = []struct {
		string
		*color.Palette
	}{
		{"PICO8 Palette", &pico8.Palette},
		{"C64 Palette", &c64.Palette},
		{"CPC Palette", &cpc.Palette},
		{"MSX Palette", &msx.Palette},
		{"MSX2 Palette", &msx2.Palette},
	}
	l.current = 0
	color.Load(l.palettes[l.current].Palette)

	l.pict = pixel.Picture("graphics/paletteswatch")
}

func (loop7) Leave() {
}

// Game Loop ///////////////////////////////////////////////////////////////////

func (l *loop7) React() {
	if input.Close.Pressed() {
		cozely.Stop(nil)
	}

	if input.Right.Pressed() {
		l.current++
		if l.current >= len(l.palettes) {
			l.current = len(l.palettes) - 1
		}
		color.Load(l.palettes[l.current].Palette)
	}
	if input.Left.Pressed() {
		l.current--
		if l.current < 0 {
			l.current = 0
		}
		color.Load(l.palettes[l.current].Palette)
	}
}

func (loop7) Update() {
}

func (l *loop7) Render() {
	pixel.Clear(color.Transparent)

	cs := pixel.Resolution()

	ps := l.pict.Size()
	p := cs.Minus(ps).Slash(2)
	_ = p
	l.pict.Paint(p, 0)

	cur := pixel.Cursor{}
	cur.Color = color.MidGray
	cur.Position = p.Minus(pixel.XY{0, 8})
	cur.Print(l.palettes[l.current].string)
}

// Copyright (c) 2018-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).
