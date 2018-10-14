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
	mire                                   pixel.PictureID
	srgbGray, srgbRed, srgbGreen, srgbBlue pixel.PictureID
	mode                                   int
}

////////////////////////////////////////////////////////////////////////////////

func TestTest1(t *testing.T) {
	do(func() {
		defer cozely.Recover()
		l := loop1{}
		l.setup()
		err := cozely.Run(&l)
		if err != nil {
			t.Error(err)
		}
	})
}

func (l *loop1) setup() {
	pixel.SetResolution(pixel.XY{320, 180})

	l.srgbGray = pixel.Picture("graphics/srgb-gray")
	l.srgbRed = pixel.Picture("graphics/srgb-red")
	l.srgbGreen = pixel.Picture("graphics/srgb-green")
	l.srgbBlue = pixel.Picture("graphics/srgb-blue")
	l.mire = pixel.Picture("graphics/mire")
}

func (l *loop1) Enter() {
	l.mode = 0
}

func (l *loop1) Leave() {
}

////////////////////////////////////////////////////////////////////////////////

func (l *loop1) React() {
	if input.Close.Pressed() {
		cozely.Stop(nil)
	}
}

func (l *loop1) Update() {
}

func (l *loop1) Render() {
	pixel.Clear(0)
	sz := pixel.Resolution()

	pz := l.mire.Size()
	l.mire.Paint(pixel.XY{0, 0}, 0)
	l.mire.Paint(pixel.XY{0, sz.Y - pz.Y}, 0)
	l.mire.Paint(pixel.XY{sz.X - pz.X, 0}, 0)
	l.mire.Paint(sz.Minus(pz), 0)

	pixel.Box(pz, sz.Minus(pz.Times(2)).MinusS(1), -1, 0, color.DarkGray, color.Black)

	for i := int16(0); i < 6; i++ {
		pixel.Box(
			pixel.XY{sz.X/2 - 3*10 + i*10, sz.Y - 20},
			pixel.XY{8, 8},
			0,
			i,
			254, 252,
		)
	}

	cur := pixel.Cursor{
		Position: pixel.XY{pz.X + 28, 108},
	}
	cur.Margin = cur.Position.X
	cur.Color = color.MidGray
	cur.Println("  sRGB:")
	cur.Println("Linear:")

	pz = l.srgbGray.Size()
	l.srgbGray.Paint(pixel.XY{pz.X + 44 + 1*(pz.X+4), 96}, 0)
	l.srgbRed.Paint(pixel.XY{pz.X + 44 + 2*(pz.X+4), 96}, 0)
	l.srgbGreen.Paint(pixel.XY{pz.X + 44 + 3*(pz.X+4), 96}, 0)
	l.srgbBlue.Paint(pixel.XY{pz.X + 44 + 4*(pz.X+4), 96}, 0)
}
