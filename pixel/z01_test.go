// Copyright (c) 2018-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package pixel_test

import (
	"testing"

	"github.com/cozely/cozely"
	"github.com/cozely/cozely/color"
	"github.com/cozely/cozely/input"
	"github.com/cozely/cozely/pixel"
)

////////////////////////////////////////////////////////////////////////////////

var ()

type loop1 struct {
	palmire, palsrgb                       color.Palette
	mire                                   pixel.PictureID
	srgbGray, srgbRed, srgbGreen, srgbBlue pixel.PictureID
	mode                                   int
}

////////////////////////////////////////////////////////////////////////////////

func TestTest1(t *testing.T) {
	do(func() {
		defer cozely.Recover()
		l := loop1{}
		l.declare()
		err := cozely.Run(&l)
		if err != nil {
			t.Error(err)
		}
	})
}

func (l *loop1) declare() {
	pixel.SetResolution(pixel.XY{320, 180})

	l.palmire = color.PaletteFrom("graphics/mire")
	l.palsrgb = color.PaletteFrom("graphics/srgb-gray")

	l.mire = pixel.Picture("graphics/mire")
	l.srgbGray = pixel.Picture("graphics/srgb-gray")
	l.srgbRed = pixel.Picture("graphics/srgb-red")
	l.srgbGreen = pixel.Picture("graphics/srgb-green")
	l.srgbBlue = pixel.Picture("graphics/srgb-blue")
}

func (l *loop1) Enter() {
	l.mode = 0

	pixel.SetPalette(l.palsrgb)
}

func (loop1) Leave() {
}

////////////////////////////////////////////////////////////////////////////////

func (l *loop1) React() {
	if input.MenuBack.Pushed() {
		cozely.Stop(nil)
	}
}

func (loop1) Update() {
}

func (l *loop1) Render() {
	pixel.Clear(0)
	sz := pixel.Resolution()

	pz := l.mire.Size()
	l.mire.Paint(pixel.XY{0, 0}, 0)
	l.mire.Paint(pixel.XY{0, sz.Y - pz.Y}, 0)
	l.mire.Paint(pixel.XY{sz.X - pz.X, 0}, 0)
	l.mire.Paint(sz.Minus(pz), 0)

	pixel.Box(pz, sz.Minus(pz.Times(2)).MinusS(1), -1, 0, 3, 4)

	for i := int16(0); i < 6; i++ {
		pixel.Box(
			pixel.XY{sz.X/2 - 3*10 + i*10, sz.Y - 20},
			pixel.XY{8, 8},
			0,
			i,
			1, 3,
		)
	}

	pz = l.srgbGray.Size()
	l.srgbGray.Paint(pixel.XY{sz.X/2 - pz.X/2, 48}, 0)
	l.srgbRed.Paint(pixel.XY{sz.X/4 - pz.X/2, 96}, 0)
	l.srgbGreen.Paint(pixel.XY{sz.X/2 - pz.X/2, 96}, 0)
	l.srgbBlue.Paint(pixel.XY{3*sz.X/4 - pz.X/2, 96}, 0)

	cur := pixel.Cursor{
		Position: pixel.XY{sz.X/2 - pz.X/2 - 29, 66},
	}
	cur.Color = 2
	cur.Print("sRGB       ")
	cur.Color = 3
	cur.Print("Linear")
}
