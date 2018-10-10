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

func (a *loop1) declare() {
	pixel.SetResolution(pixel.XY{320, 180})

	a.palmire = color.PaletteFrom("graphics/mire")
	a.palsrgb = color.PaletteFrom("graphics/srgb-gray")

	a.mire = pixel.Picture("graphics/mire")
	a.srgbGray = pixel.Picture("graphics/srgb-gray")
	a.srgbRed = pixel.Picture("graphics/srgb-red")
	a.srgbGreen = pixel.Picture("graphics/srgb-green")
	a.srgbBlue = pixel.Picture("graphics/srgb-blue")
}

func (a *loop1) Enter() {
	a.mode = 0

	pixel.SetPalette(a.palsrgb)
}

func (loop1) Leave() {
}

////////////////////////////////////////////////////////////////////////////////

func (a *loop1) React() {
	if input.MenuBack.Pushed() {
		cozely.Stop(nil)
	}
}

func (loop1) Update() {
}

func (a *loop1) Render() {
	pixel.Clear(0)
	sz := pixel.Resolution()

	pz := a.mire.Size()
	a.mire.Paint(pixel.XY{0, 0}, 0)
	a.mire.Paint(pixel.XY{0, sz.Y - pz.Y}, 0)
	a.mire.Paint(pixel.XY{sz.X - pz.X, 0}, 0)
	a.mire.Paint(sz.Minus(pz), 0)

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

	pz = a.srgbGray.Size()
	a.srgbGray.Paint(pixel.XY{sz.X/2 - pz.X/2, 48}, 0)
	a.srgbRed.Paint(pixel.XY{sz.X/4 - pz.X/2, 96}, 0)
	a.srgbGreen.Paint(pixel.XY{sz.X/2 - pz.X/2, 96}, 0)
	a.srgbBlue.Paint(pixel.XY{3*sz.X/4 - pz.X/2, 96}, 0)

	cur := pixel.Cursor{
		Position: pixel.XY{sz.X/2 - pz.X/2 - 29, 66},
	}
	cur.Color = 2
	cur.Print("sRGB       ")
	cur.Color = 3
	cur.Print("Linear")
}
