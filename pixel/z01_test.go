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
		input.Load(bindings)
		err := cozely.Run(&l)
		if err != nil {
			t.Error(err)
		}
	})
}

func (a *loop1) declare() {
	pixel.SetResolution(320, 180)

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

	pixel.Palette(a.palmire)
}

func (loop1) Leave() {
}

////////////////////////////////////////////////////////////////////////////////

func (a *loop1) React() {
	if quit.Started(0) {
		cozely.Stop(nil)
	}

	if next.Started(0) {
		a.mode++
		if a.mode > 1 {
			a.mode = 0
		}
		switch a.mode {
		case 0:
			pixel.Palette(a.palmire)
		case 1:
			pixel.Palette(a.palsrgb)
		}
	}
}

func (loop1) Update() {
}

func (a *loop1) Render() {
	pixel.Clear(0)
	sz := pixel.Resolution()
	switch a.mode {
	case 0:
		pz := a.mire.Size()
		pixel.Paint(a.mire, pixel.XY{0, 0})
		pixel.Paint(a.mire, pixel.XY{0, sz.Y - pz.Y})
		pixel.Paint(a.mire, pixel.XY{sz.X - pz.X, 0})
		pixel.Paint(a.mire, sz.Minus(pz))
	case 1:
		pz := a.srgbGray.Size()
		pixel.Paint(a.srgbGray, pixel.XY{sz.X/2 - pz.X/2, 32})
		pixel.Paint(a.srgbRed, pixel.XY{sz.X/4 - pz.X/2, 96})
		pixel.Paint(a.srgbGreen, pixel.XY{sz.X/2 - pz.X/2, 96})
		pixel.Paint(a.srgbBlue, pixel.XY{3*sz.X/4 - pz.X/2, 96})
	}
}
